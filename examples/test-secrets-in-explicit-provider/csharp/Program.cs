using System;
using System.Collections.Generic;
using System.Text;

using Pulumi;
using Pulumi.Random;

return await Deployment.RunAsync(() =>
{

   // Test passing values marked secret by the user program to the provider.
   var provider1 = new Pulumi.Docker.Provider("provider-with-sensitive-address", new Pulumi.Docker.ProviderArgs
   {
      RegistryAuth = new List<Pulumi.Docker.Inputs.ProviderRegistryAuthArgs>
      {
         new Pulumi.Docker.Inputs.ProviderRegistryAuthArgs
         {
            Address = Output.CreateSecret("secret-address"),
            Username = "some-user",
         }
      }
   });

   // Test setting values marked secret in the provider schema.
   var provider2 = new Pulumi.Docker.Provider("provider-with-password", new Pulumi.Docker.ProviderArgs
   {
      RegistryAuth = new List<Pulumi.Docker.Inputs.ProviderRegistryAuthArgs>
      {
         new Pulumi.Docker.Inputs.ProviderRegistryAuthArgs
         {
            Address = "some-address",
            Username = "some-user",
            Password = "secret-password",
         }
      }
   });

   // Test setting values that start as unknown and resolve to secrets.
   var password = new Pulumi.Random.RandomPassword("password", new()
   {
       Length = 16,
       Special = false,
   });

   var provider3 = new Pulumi.Docker.Provider("provider-with-random-sensitive-username", new Pulumi.Docker.ProviderArgs
   {
      RegistryAuth = new List<Pulumi.Docker.Inputs.ProviderRegistryAuthArgs>
      {
         new Pulumi.Docker.Inputs.ProviderRegistryAuthArgs
         {
            Address = "some-address",
            Username = password.Result,
         }
      }
   });

   var pw = Output.Unsecret(password.Result).Apply(s => Convert.ToBase64String(Encoding.UTF8.GetBytes(s)));

   return new Dictionary<string, object?>{
       ["password"] = pw,
   };
});
