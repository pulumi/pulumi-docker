{{% examples %}}
## Example Usage
{{% example %}}
### A Docker image build

```typescript
import * as pulumi from "@pulumi/pulumi";

export const imageName = demoImage.imageName;
```
```python
import pulumi

pulumi.export("imageName", demo_image["imageName"])
```
```csharp
using System.Collections.Generic;
using Pulumi;

return await Deployment.RunAsync(() => 
{
    return new Dictionary<string, object?>
    {
        ["imageName"] = demoImage.ImageName,
    };
});

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
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
{{% /example %}}
{{% /examples %}}
