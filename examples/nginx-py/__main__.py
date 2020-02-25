# Copyright 2016-2018, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import pulumi
from pulumi_docker import RemoteImage, Container

# Get a reference to the remote image "nginx:1.15.6". Without specifying the repository, the Docker provider will
# try to download it from the public Docker Hub.
image = RemoteImage("nginx-image", name="nginx", keep_locally=True)
container = Container("nginx", image=image.latest, ports=[{
    "internal": 80
}])

