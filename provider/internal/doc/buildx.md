{{% examples %}}
## Example Usage
{{% example %}}
### Multi-platform image

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
    platforms: [
        docker.buildx.image.Platform.Plan9_amd64,
        docker.buildx.image.Platform.Plan9_386,
    ],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ),
    platforms=[
        docker.buildx/image.Platform.PLAN9_AMD64,
        docker.buildx/image.Platform.PLAN9_386,
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
        Platforms = new[]
        {
            Docker.Buildx.Image.Platform.Plan9_amd64,
            Docker.Buildx.Image.Platform.Plan9_386,
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
			Platforms: buildx.PlatformArray{
				buildx.Platform_Plan9_amd64,
				buildx.Platform_Plan9_386,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Multi-platform image
name: multi-platform
resources:
    image:
        properties:
            context:
                location: app
            dockerfile:
                location: app/Dockerfile
            platforms:
                - plan9/amd64
                - plan9/386
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .platforms(            
                "plan9/amd64",
                "plan9/386")
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Registry export

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
    exports: [{
        registry: {
            ociMediaTypes: true,
        },
    }],
    registries: [{
        address: "docker.io",
        password: dockerHubPassword,
        username: "pulumibot",
    }],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ),
    exports=[docker.buildx.ExportEntryArgs(
        registry=docker.buildx.ExportRegistryArgs(
            oci_media_types=True,
        ),
    )],
    registries=[docker.buildx.RegistryAuthArgs(
        address="docker.io",
        password=docker_hub_password,
        username="pulumibot",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
        Exports = new[]
        {
            new Docker.Buildx.Inputs.ExportEntryArgs
            {
                Registry = new Docker.Buildx.Inputs.ExportRegistryArgs
                {
                    OciMediaTypes = true,
                },
            },
        },
        Registries = new[]
        {
            new Docker.Buildx.Inputs.RegistryAuthArgs
            {
                Address = "docker.io",
                Password = dockerHubPassword,
                Username = "pulumibot",
            },
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
			Exports: buildx.ExportEntryArray{
				&buildx.ExportEntryArgs{
					Registry: &buildx.ExportRegistryArgs{
						OciMediaTypes: pulumi.Bool(true),
					},
				},
			},
			Registries: buildx.RegistryAuthArray{
				&buildx.RegistryAuthArgs{
					Address:  pulumi.String("docker.io"),
					Password: pulumi.Any(dockerHubPassword),
					Username: pulumi.String("pulumibot"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Registry export
name: registry
resources:
    image:
        properties:
            context:
                location: app
            dockerfile:
                location: app/Dockerfile
            exports:
                - registry:
                    ociMediaTypes: true
            registries:
                - address: docker.io
                  password: ${dockerHubPassword}
                  username: pulumibot
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
import com.pulumi.docker.buildx.inputs.ExportEntryArgs;
import com.pulumi.docker.buildx.inputs.ExportRegistryArgs;
import com.pulumi.docker.buildx.inputs.RegistryAuthArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .exports(ExportEntryArgs.builder()
                .registry(ExportRegistryArgs.builder()
                    .ociMediaTypes(true)
                    .build())
                .build())
            .registries(RegistryAuthArgs.builder()
                .address("docker.io")
                .password(dockerHubPassword)
                .username("pulumibot")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Caching

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    cacheFrom: [{
        local: {
            src: "tmp/cache",
        },
    }],
    cacheTo: [{
        local: {
            dest: "tmp/cache",
            mode: docker.buildx.image.CacheMode.Max,
        },
    }],
    context: {
        location: "app",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    cache_from=[docker.buildx.CacheFromEntryArgs(
        local=docker.buildx.CacheFromLocalArgs(
            src="tmp/cache",
        ),
    )],
    cache_to=[docker.buildx.CacheToEntryArgs(
        local=docker.buildx.CacheToLocalArgs(
            dest="tmp/cache",
            mode=docker.buildx/image.CacheMode.MAX,
        ),
    )],
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        CacheFrom = new[]
        {
            new Docker.Buildx.Inputs.CacheFromEntryArgs
            {
                Local = new Docker.Buildx.Inputs.CacheFromLocalArgs
                {
                    Src = "tmp/cache",
                },
            },
        },
        CacheTo = new[]
        {
            new Docker.Buildx.Inputs.CacheToEntryArgs
            {
                Local = new Docker.Buildx.Inputs.CacheToLocalArgs
                {
                    Dest = "tmp/cache",
                    Mode = Docker.Buildx.Image.CacheMode.Max,
                },
            },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			CacheFrom: buildx.CacheFromEntryArray{
				&buildx.CacheFromEntryArgs{
					Local: &buildx.CacheFromLocalArgs{
						Src: pulumi.String("tmp/cache"),
					},
				},
			},
			CacheTo: buildx.CacheToEntryArray{
				&buildx.CacheToEntryArgs{
					Local: &buildx.CacheToLocalArgs{
						Dest: pulumi.String("tmp/cache"),
						Mode: buildx.CacheModeMax,
					},
				},
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Caching
name: caching
resources:
    image:
        properties:
            cacheFrom:
                - local:
                    src: tmp/cache
            cacheTo:
                - local:
                    dest: tmp/cache
                    mode: max
            context:
                location: app
            dockerfile:
                location: app/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.CacheFromEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheFromLocalArgs;
import com.pulumi.docker.buildx.inputs.CacheToEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheToLocalArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .cacheFrom(CacheFromEntryArgs.builder()
                .local(CacheFromLocalArgs.builder()
                    .src("tmp/cache")
                    .build())
                .build())
            .cacheTo(CacheToEntryArgs.builder()
                .local(CacheToLocalArgs.builder()
                    .dest("tmp/cache")
                    .mode("max")
                    .build())
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Build arguments

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    buildArgs: {
        SET_ME_TO_TRUE: "true",
    },
    context: {
        location: "app",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    build_args={
        "SET_ME_TO_TRUE": "true",
    },
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        BuildArgs = 
        {
            { "SET_ME_TO_TRUE", "true" },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			BuildArgs: pulumi.StringMap{
				"SET_ME_TO_TRUE": pulumi.String("true"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Build arguments
name: build-args
resources:
    image:
        properties:
            buildArgs:
                SET_ME_TO_TRUE: "true"
            context:
                location: app
            dockerfile:
                location: app/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .buildArgs(Map.of("SET_ME_TO_TRUE", "true"))
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Build targets

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
    targets: [
        "build-me",
        "also-build-me",
    ],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ),
    targets=[
        "build-me",
        "also-build-me",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
        Targets = new[]
        {
            "build-me",
            "also-build-me",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
			Targets: pulumi.StringArray{
				pulumi.String("build-me"),
				pulumi.String("also-build-me"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Build targets
name: build-targets
resources:
    image:
        properties:
            context:
                location: app
            dockerfile:
                location: app/Dockerfile
            targets:
                - build-me
                - also-build-me
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .targets(            
                "build-me",
                "also-build-me")
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Named contexts

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
        named: {
            "golang:latest": {
                location: "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
            },
        },
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
        named={
            "golang:latest": docker.buildx.ContextArgs(
                location="docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
            ),
        },
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
            Named = 
            {
                { "golang:latest", new Docker.Buildx.Inputs.ContextArgs
                {
                    Location = "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
                } },
            },
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
				Named: buildx.ContextMap{
					"golang:latest": &buildx.ContextArgs{
						Location: pulumi.String("docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984"),
					},
				},
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Named contexts
name: named-contexts
resources:
    image:
        properties:
            context:
                location: app
                named:
                    golang:latest:
                        location: docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984
            dockerfile:
                location: app/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .named(Map.of("golang:latest", Map.of("location", "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984")))
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Remote context

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {context: {
    location: "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
}});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image", context=docker.buildx.BuildContextArgs(
    location="https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Remote context
name: remote-context
resources:
    image:
        properties:
            context:
                location: https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Inline Dockerfile

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    dockerfile: {
        inline: `FROM busybox
COPY hello.c ./
`,
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        inline="""FROM busybox
COPY hello.c ./
""",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Inline = @"FROM busybox
COPY hello.c ./
",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Inline: pulumi.String("FROM busybox\nCOPY hello.c ./\n"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Inline Dockerfile
name: inline
resources:
    image:
        properties:
            context:
                location: app
            dockerfile:
                inline: |
                    FROM busybox
                    COPY hello.c ./
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .inline("""
FROM busybox
COPY hello.c ./
                """)
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Remote context

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "https://github.com/docker-library/hello-world.git",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="https://github.com/docker-library/hello-world.git",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "https://github.com/docker-library/hello-world.git",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("https://github.com/docker-library/hello-world.git"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Remote context
name: remote-context
resources:
    image:
        properties:
            context:
                location: https://github.com/docker-library/hello-world.git
            dockerfile:
                location: app/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("https://github.com/docker-library/hello-world.git")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Local export

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
    exports: [{
        docker: {
            tar: true,
        },
    }],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ),
    exports=[docker.buildx.ExportEntryArgs(
        docker=docker.buildx.ExportDockerArgs(
            tar=True,
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
        Exports = new[]
        {
            new Docker.Buildx.Inputs.ExportEntryArgs
            {
                Docker = new Docker.Buildx.Inputs.ExportDockerArgs
                {
                    Tar = true,
                },
            },
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
			Exports: buildx.ExportEntryArray{
				&buildx.ExportEntryArgs{
					Docker: &buildx.ExportDockerArgs{
						Tar: pulumi.Bool(true),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Local export
name: docker-load
resources:
    image:
        properties:
            context:
                location: app
            dockerfile:
                location: app/Dockerfile
            exports:
                - docker:
                    tar: true
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
import com.pulumi.docker.buildx.inputs.ExportEntryArgs;
import com.pulumi.docker.buildx.inputs.ExportDockerArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .exports(ExportEntryArgs.builder()
                .docker(ExportDockerArgs.builder()
                    .tar(true)
                    .build())
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}