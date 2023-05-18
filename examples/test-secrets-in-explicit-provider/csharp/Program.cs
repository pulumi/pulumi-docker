using Pulumi;
using System.Collections.Generic;

return await Deployment.RunAsync(() =>
{
   var provider = new Pulumi.Docker.Provider("docker", new Pulumi.Docker.ProviderArgs
   {
      Host = "host",
      RegistryAuth = new List<Pulumi.Docker.Inputs.ProviderRegistryAuthArgs>
      {
         new Pulumi.Docker.Inputs.ProviderRegistryAuthArgs
         {
            Address = "somewhere.org",
            Username = "some-user",
            Password = "some-password"
         }
      }
   });

   return new Dictionary<string, object?>
   {
      ["outputKey"] = "outputValue"
   };
});
