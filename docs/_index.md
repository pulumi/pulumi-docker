---
title: Docker Provider
meta_desc: Provides an overview on how to configure the Pulumi Docker provider.
layout: package
---
## Installation

The docker provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/docker`](https://www.npmjs.com/package/@pulumi/docker)
* Python: [`pulumi-docker`](https://pypi.org/project/pulumi-docker/)
* Go: [`github.com/pulumi/pulumi-docker/sdk/v4/go/docker`](https://github.com/pulumi/pulumi-docker)
* .NET: [`Pulumi.Docker`](https://www.nuget.org/packages/Pulumi.Docker)
* Java: [`com.pulumi/docker`](https://central.sonatype.com/artifact/com.pulumi/docker)
## Overview
The Docker provider is used to interact with Docker containers and images.
It uses the Docker API to manage the lifecycle of Docker containers. Because
the Docker provider uses the Docker API, it is immediately compatible not
only with single server Docker but Swarm and any additional Docker-compatible
API hosts.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

// Pulls the image
const ubuntu = new docker.RemoteImage("ubuntu", {name: "ubuntu:latest"});
// Create a container
const foo = new docker.Container("foo", {
    image: ubuntu.imageId,
    name: "foo",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```python
import pulumi
import pulumi_docker as docker

# Pulls the image
ubuntu = docker.RemoteImage("ubuntu", name="ubuntu:latest")
# Create a container
foo = docker.Container("foo",
    image=ubuntu.image_id,
    name="foo")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() =>
{
    // Pulls the image
    var ubuntu = new Docker.RemoteImage("ubuntu", new()
    {
        Name = "ubuntu:latest",
    });

    // Create a container
    var foo = new Docker.Container("foo", new()
    {
        Image = ubuntu.ImageId,
        Name = "foo",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Pulls the image
		ubuntu, err := docker.NewRemoteImage(ctx, "ubuntu", &docker.RemoteImageArgs{
			Name: pulumi.String("ubuntu:latest"),
		})
		if err != nil {
			return err
		}
		// Create a container
		_, err = docker.NewContainer(ctx, "foo", &docker.ContainerArgs{
			Image: ubuntu.ImageId,
			Name:  pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```yaml
resources:
  # Pulls the image
  ubuntu:
    type: docker:RemoteImage
    properties:
      name: ubuntu:latest
  # Create a container
  foo:
    type: docker:Container
    properties:
      image: ${ubuntu.imageId}
      name: foo
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.RemoteImage;
import com.pulumi.docker.RemoteImageArgs;
import com.pulumi.docker.Container;
import com.pulumi.docker.ContainerArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Pulls the image
        var ubuntu = new RemoteImage("ubuntu", RemoteImageArgs.builder()
            .name("ubuntu:latest")
            .build());

        // Create a container
        var foo = new Container("foo", ContainerArgs.builder()
            .image(ubuntu.imageId())
            .name("foo")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

// Pulls the image
const ubuntu = new docker.RemoteImage("ubuntu", {name: "ubuntu:latest"});
// Create a container
const foo = new docker.Container("foo", {
    image: ubuntu.imageId,
    name: "foo",
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```python
import pulumi
import pulumi_docker as docker

# Pulls the image
ubuntu = docker.RemoteImage("ubuntu", name="ubuntu:latest")
# Create a container
foo = docker.Container("foo",
    image=ubuntu.image_id,
    name="foo")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() =>
{
    // Pulls the image
    var ubuntu = new Docker.RemoteImage("ubuntu", new()
    {
        Name = "ubuntu:latest",
    });

    // Create a container
    var foo = new Docker.Container("foo", new()
    {
        Image = ubuntu.ImageId,
        Name = "foo",
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Pulls the image
		ubuntu, err := docker.NewRemoteImage(ctx, "ubuntu", &docker.RemoteImageArgs{
			Name: pulumi.String("ubuntu:latest"),
		})
		if err != nil {
			return err
		}
		// Create a container
		_, err = docker.NewContainer(ctx, "foo", &docker.ContainerArgs{
			Image: ubuntu.ImageId,
			Name:  pulumi.String("foo"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```yaml
resources:
  # Pulls the image
  ubuntu:
    type: docker:RemoteImage
    properties:
      name: ubuntu:latest
  # Create a container
  foo:
    type: docker:Container
    properties:
      image: ${ubuntu.imageId}
      name: foo
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    docker:host:
        value: unix:///var/run/docker.sock

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.RemoteImage;
import com.pulumi.docker.RemoteImageArgs;
import com.pulumi.docker.Container;
import com.pulumi.docker.ContainerArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Pulls the image
        var ubuntu = new RemoteImage("ubuntu", RemoteImageArgs.builder()
            .name("ubuntu:latest")
            .build());

        // Create a container
        var foo = new Container("foo", ContainerArgs.builder()
            .image(ubuntu.imageId())
            .name("foo")
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Remote Hosts
You can also use the `ssh` protocol to connect to the docker host on a remote machine.
The configuration would look as follows:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    docker:host:
        value: ssh://user@remote-host:22
    docker:sshOpts:
        value:
            - -o
            - StrictHostKeyChecking=no
            - -o
            - UserKnownHostsFile=/dev/null

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    docker:host:
        value: ssh://user@remote-host:22
    docker:sshOpts:
        value:
            - -o
            - StrictHostKeyChecking=no
            - -o
            - UserKnownHostsFile=/dev/null

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    docker:host:
        value: ssh://user@remote-host:22
    docker:sshOpts:
        value:
            - -o
            - StrictHostKeyChecking=no
            - -o
            - UserKnownHostsFile=/dev/null

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    docker:host:
        value: ssh://user@remote-host:22
    docker:sshOpts:
        value:
            - -o
            - StrictHostKeyChecking=no
            - -o
            - UserKnownHostsFile=/dev/null

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    docker:host:
        value: ssh://user@remote-host:22
    docker:sshOpts:
        value:
            - -o
            - StrictHostKeyChecking=no
            - -o
            - UserKnownHostsFile=/dev/null

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    docker:host:
        value: ssh://user@remote-host:22
    docker:sshOpts:
        value:
            - -o
            - StrictHostKeyChecking=no
            - -o
            - UserKnownHostsFile=/dev/null

```

{{% /choosable %}}
{{< /chooser >}}

When using a remote host, the daemon configuration on the remote host can apply default configuration to your resources when running `pulumi up`, for example by appling log options to containers. When running `pulumi preview` the next time, it will show up as a diff. In such cases it is recommended to use the `ignoreChanges` lifecycle meta-argument to ignore the changing attribute (See this issue for more information).
## Registry credentials

Registry credentials can be provided on a per-registry basis with the `registryAuth`
field, passing either a config file or the username/password directly.
If you want to use an insecure http registry, please explicitly specify the `address` with the `http` protocol.

> **Note**
The config file is loaded from the machine `pulumi` runs on. This also applies when the specified docker host is on another machine.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}

{{% /choosable %}}
{{% choosable language python %}}

{{% /choosable %}}
{{% choosable language csharp %}}

{{% /choosable %}}
{{% choosable language go %}}

{{% /choosable %}}
{{% choosable language yaml %}}

{{% /choosable %}}
{{% choosable language java %}}

{{% /choosable %}}
{{< /chooser >}}

> **Note**
When passing in a config file either the corresponding `auth` string of the repository is read or the os specific
[credential helpers](https://github.com/docker/docker-credential-helpers#available-programs) are
used to retrieve the authentication credentials.

> **Note**
`configFile` has predence over all other options. You can theoretically specify values for every attribute but the credentials obtained through the `configFile` will override the manually set `username`/`password`

You can still use the environment variables `DOCKER_REGISTRY_USER` and `DOCKER_REGISTRY_PASS`.

An example content of the file `~/.docker/config.json` on macOS may look like follows:

```json
{
    "auths": {
        "repo.mycompany:8181": {
            "auth": "dXNlcjpwYXNz="
        },
        "otherrepo.other-company:8181": {}
    },
    "credsStore": "osxkeychain"
}
```
## Certificate information

Specify certificate information either with a directory or
directly with the content of the files for connecting to the Docker host via TLS.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    docker:caMaterial:
        value: 'TODO: file(pathexpand("~/.docker/ca.pem"))'
    docker:certMaterial:
        value: 'TODO: file(pathexpand("~/.docker/cert.pem"))'
    docker:certPath:
        value: /Users/guin/.docker
    docker:host:
        value: tcp://your-host-ip:2376/
    docker:keyMaterial:
        value: 'TODO: file(pathexpand("~/.docker/key.pem"))'

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    docker:caMaterial:
        value: 'TODO: file(pathexpand("~/.docker/ca.pem"))'
    docker:certMaterial:
        value: 'TODO: file(pathexpand("~/.docker/cert.pem"))'
    docker:certPath:
        value: /Users/guin/.docker
    docker:host:
        value: tcp://your-host-ip:2376/
    docker:keyMaterial:
        value: 'TODO: file(pathexpand("~/.docker/key.pem"))'

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    docker:caMaterial:
        value: 'TODO: file(pathexpand("~/.docker/ca.pem"))'
    docker:certMaterial:
        value: 'TODO: file(pathexpand("~/.docker/cert.pem"))'
    docker:certPath:
        value: /Users/guin/.docker
    docker:host:
        value: tcp://your-host-ip:2376/
    docker:keyMaterial:
        value: 'TODO: file(pathexpand("~/.docker/key.pem"))'

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    docker:caMaterial:
        value: 'TODO: file(pathexpand("~/.docker/ca.pem"))'
    docker:certMaterial:
        value: 'TODO: file(pathexpand("~/.docker/cert.pem"))'
    docker:certPath:
        value: /Users/guin/.docker
    docker:host:
        value: tcp://your-host-ip:2376/
    docker:keyMaterial:
        value: 'TODO: file(pathexpand("~/.docker/key.pem"))'

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    docker:caMaterial:
        value: 'TODO: file(pathexpand("~/.docker/ca.pem"))'
    docker:certMaterial:
        value: 'TODO: file(pathexpand("~/.docker/cert.pem"))'
    docker:certPath:
        value: /Users/guin/.docker
    docker:host:
        value: tcp://your-host-ip:2376/
    docker:keyMaterial:
        value: 'TODO: file(pathexpand("~/.docker/key.pem"))'

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    docker:caMaterial:
        value: 'TODO: file(pathexpand("~/.docker/ca.pem"))'
    docker:certMaterial:
        value: 'TODO: file(pathexpand("~/.docker/cert.pem"))'
    docker:certPath:
        value: /Users/guin/.docker
    docker:host:
        value: tcp://your-host-ip:2376/
    docker:keyMaterial:
        value: 'TODO: file(pathexpand("~/.docker/key.pem"))'

```

{{% /choosable %}}
{{< /chooser >}}

<!-- schema generated by tfplugindocs -->
## Configuration Reference

- `caMaterial` (String) PEM-encoded content of Docker host CA certificate
- `certMaterial` (String) PEM-encoded content of Docker client certificate
- `certPath` (String) Path to directory with Docker TLS config
- `host` (String) The Docker daemon address
- `keyMaterial` (String) PEM-encoded content of Docker client private key
- `registryAuth` (Block Set) (see below for nested schema)
- `sshOpts` (List of String) Additional SSH option flags to be appended when using `ssh://` protocol

<a id="nestedblock--registry_auth"></a>
### Nested Configuration Reference for `registryAuth`

Required:

- `address` (String) Address of the registry

Optional:

- `authDisabled` (Boolean) Setting this to `true` will tell the provider that this registry does not need authentication. Due to the docker internals, the provider will use dummy credentials (see <https://github.com/kreuzwerker/pulumi-provider-docker/issues/470> for more information). Defaults to `false`.
- `configFile` (String) Path to docker json file for registry auth. Defaults to `~/.docker/config.json`. If `DOCKER_CONFIG` is set, the value of `DOCKER_CONFIG` is used as the path. `configFile` has predencen over all other options.
- `configFileContent` (String) Plain content of the docker json file for registry auth. `configFileContent` has precedence over username/password.
- `password` (String, Sensitive) Password for the registry. Defaults to `DOCKER_REGISTRY_PASS` env variable if set.
- `username` (String) Username for the registry. Defaults to `DOCKER_REGISTRY_USER` env variable if set.