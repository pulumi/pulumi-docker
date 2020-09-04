using System.Collections.Generic;
using System.Threading.Tasks;

using Pulumi;
using Pulumi.Docker;

class Program
{
    static Task<int> Main()
    {
        return Deployment.RunAsync(() => {
            var image = new Image("my-image", new ImageArgs
            {
                ImageName = "pulumi-user/example:v1.0.0",
                Build = new DockerBuild
                {
                    Env = new Dictionary<string, string>
                    { 
                        {"TEST_ENV", "42"}, 
                    },
                    Target = "dependencies",
                },
                SkipPush = true,
            });

            return new Dictionary<string, object?>
            {
                { "image", image.ImageName },
            };
        });
    }
}
