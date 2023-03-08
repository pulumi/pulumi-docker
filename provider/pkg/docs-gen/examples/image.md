## Example Usage
{{% example %}}
### A Docker image build

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const demoImage = new docker.Image("demo-image", {
    build: {
        context: ".",
        dockerfile: "Dockerfile",
    },
    imageName: "username/image:tag1",
    skipPush: true,
});
export const imageName = demoImage.imageName;
```
```python
import pulumi
import pulumi_docker as docker

demo_image = docker.Image("demo-image",
    build={
        "context": ".",
        "dockerfile": "Dockerfile",
    },
    image_name="username/image:tag1",
    skip_push=True)
pulumi.export("imageName", demo_image.image_name)
```
```csharp
using System.Collections.Generic;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var demoImage = new Docker.Image("demo-image", new()
    {
        Build = 
        {
            { "context", "." },
            { "dockerfile", "Dockerfile" },
        },
        ImageName = "username/image:tag1",
        SkipPush = true,
    });

    return new Dictionary<string, object?>
    {
        ["imageName"] = demoImage.ImageName,
    };
});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		demoImage, err := docker.NewImage(ctx, "demo-image", &docker.ImageArgs{
			Build: pulumi.Any{
				Context:    pulumi.String("."),
				Dockerfile: pulumi.String("Dockerfile"),
			},
			ImageName: pulumi.String("username/image:tag1"),
			SkipPush:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("imageName", demoImage.ImageName)
		return nil
	})
}
```
```yaml
config: {}
description: A Docker image build
name: image-yaml
outputs:
    imageName: ${demo-image.imageName}
resources:
    demo-image:
        options:
            version: v4.0.0-alpha.1
        properties:
            build:
                context: .
                dockerfile: Dockerfile
            imageName: username/image:tag1
            skipPush: true
        type: docker:Image
runtime: yaml
variables: {}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.Image;
import com.pulumi.docker.ImageArgs;
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
        var demoImage = new Image("demoImage", ImageArgs.builder()        
            .build(%!v(PANIC=Format method: interface conversion: model.Expression is *model.TemplateExpression, not *model.LiteralValueExpression))
            .imageName("username/image:tag1")
            .skipPush(true)
            .build());

        ctx.export("imageName", demoImage.imageName());
    }
}
```
{{% /example %}}
