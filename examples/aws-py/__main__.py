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

import base64
import pulumi_aws as aws
import pulumi_docker as docker
import pulumi


# Get registry info (creds and endpoint) so we can build/publish to it.
def get_registry_info(rid):
    creds = aws.ecr.get_credentials(registry_id=rid)
    decoded = base64.b64decode(creds.authorization_token).decode()
    parts = decoded.split(':')
    if len(parts) != 2:
        raise Exception("Invalid credentials")
    return docker.ImageRegistry(creds.proxy_endpoint, parts[0], parts[1])


for i in range(3):
    # Create a private ECR registry.
    repo = aws.ecr.Repository('my-repo-%i' % i)

    registry = repo.registry_id.apply(get_registry_info)

    # Build and publish the image.

    docker.Image(
        'my-image-%i' % i,
        build=docker.DockerBuild(context='app', args={'parameter': str(i)}),
        image_name=repo.repository_url,
        registry=registry,
    )
