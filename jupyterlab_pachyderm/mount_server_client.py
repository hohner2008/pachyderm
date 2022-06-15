import os
import subprocess
import time
import platform
import json
from base64 import b64decode

import python_pachyderm
from python_pachyderm.service import health_proto
from tornado.httpclient import AsyncHTTPClient, HTTPClientError
from tornado import locks

from .pachyderm import MountInterface
from .log import get_logger
from .env import ENSURE_MOUNT_SERVER, JUPYTERLAB_CI_TESTS

lock = locks.Lock()
MOUNT_SERVER_PORT = 9002
PACH_CONFIG = "~/.pachyderm/config.json"


class MountServerClient(MountInterface):
    """Client interface for the pachctl mount-server backend."""

    def __init__(
        self,
        mount_dir: str,
    ):
        self.client = AsyncHTTPClient()
        self.mount_dir = mount_dir
        self.address = f"http://localhost:{MOUNT_SERVER_PORT}"
        if not JUPYTERLAB_CI_TESTS:
            subprocess.run(
                ["bash", "-c", "pkill -f mount-server"],
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL
            )


    async def _is_mount_server_running(self):
        get_logger().debug("Checking if mount server running...")
        try:
            await self.health()
        except Exception as e:
            get_logger().debug(f"Unable to hit server at {self.address}")
            get_logger().debug(e)
            return False
        get_logger().debug(f"Able to hit server at {self.address}")
        return True

    
    def _unmount(self):
        if platform.system() == "Linux":
            subprocess.run(["bash", "-c", f"fusermount -uzq {self.mount_dir}"])
        else:
            subprocess.run(
                ["bash", "-c", f"umount {self.mount_dir}"],
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL
            )


    async def _ensure_mount_server(self):
        """
        When we first start up, we might not have auth configured. So just try
        re-launching the mount-server on every command! If the port is already
        bound, it will exit straight away. If it's not, it might start up
        successfully with the updated config.
        """
        # TODO: add --socket and --log-file stdout args
        # TODO: add better error handling        
        if await self._is_mount_server_running():
            return True

        if not ENSURE_MOUNT_SERVER:
            get_logger().debug("Kubernetes is responsible for running mount server")
            return False

        get_logger().info("Starting mount server...")
        async with lock:
            if not await self._is_mount_server_running():                
                self._unmount()
                subprocess.Popen(
                    [
                        "bash", "-c",
                        "set -o pipefail; "
                        +f"pachctl mount-server --mount-dir {self.mount_dir}"
                        +" >> /tmp/pachctl-mount-server.log 2>&1",
                    ],
                    env={
                        "KUBECONFIG": os.path.expanduser('~/.kube/config')
                    }
                )
                
                tries = 0
                get_logger().debug("Waiting for mount server...")
                while not await self._is_mount_server_running():
                    time.sleep(1)
                    tries += 1

                    if tries == 10:
                        get_logger().debug("Unable to start mount server...")
                        return False

        return True


    async def list(self):
        await self._ensure_mount_server()
        response = await self.client.fetch(f"{self.address}/repos")
        return response.body
        

    async def mount(self, repo, branch, mode, name):
        await self._ensure_mount_server()
        response = await self.client.fetch(
            f"{self.address}/repos/{repo}/{branch}/_mount?name={name}&mode={mode}",
            method="PUT",
            body="{}",
        )
        return response.body

    async def unmount(self, repo, branch, name):
        await self._ensure_mount_server()
        response = await self.client.fetch(
            f"{self.address}/repos/{repo}/{branch}/_unmount?name={name}",
            method="PUT",
            body="{}",
        )
        return response.body

    async def unmount_all(self):
        await self._ensure_mount_server()
        response = await self.client.fetch(
            f"{self.address}/repos/_unmount",
            method="PUT",
            body="{}"
        )
        return response.body

    async def commit(self, repo, branch, name, message):
        await self._ensure_mount_server()
        pass

    async def config(self, request=None):
        await self._ensure_mount_server()
        if request is None:
            try:
                response = await self.client.fetch(f"{self.address}/config")
            except HTTPClientError as e:
                if e.code == 404:
                    return json.dumps({"cluster_status": "INVALID"})
                raise e
        else:
            response = await self.client.fetch(f"{self.address}/config", method="PUT", body=json.dumps(request))
            
        return response.body

    async def auth_login(self):
        await self._ensure_mount_server()
        response = await self.client.fetch(f"{self.address}/auth/_login", method="PUT", body="{}")
        return response.body

    async def auth_logout(self):
        await self._ensure_mount_server()
        return await self.client.fetch(f"{self.address}/auth/_logout", method="PUT", body="{}")

    async def health(self):
        response = await self.client.fetch(f"{self.address}/health")
        return response.body
