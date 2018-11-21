// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as docker from "@pulumi/docker";

// Get a reference to the remote image "nginx:1.15.6". Without specifying the repository, the Docker provider will
// try to download it from the public Docker Hub.
const image = new docker.RemoteImage("nginx-image", {
    name: "nginx:1.15.6",
    keepLocally: true, // don't delete the image from the local cache when deleting this resource
});

// Launch a container using the nginx image we just downloaded.
const container = new docker.Container("nginx", {
    image: image.name,
    ports: [{
        internal: 80,
        // external: defaults to an open ephemeral port
        // protocol: defaults to TCP
        // ip: defaults to 0.0.0.0
    }]
});

// Since the container is auto-named, export the name.
export const name = container.name;

// Since the provider picked a random ephemeral port for this container, export the endpoint.
export const endpoints = container.ports.apply(ports => `${ports![0].ip}:${ports![0].external}`);
