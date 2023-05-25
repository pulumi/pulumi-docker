using System;
using System.Collections.Generic;
using System.Text;

using Pulumi;
using Pulumi.Random;

return await Deployment.RunAsync(() =>
{
   // Test handling secrets marked by the program.
   var providerWithSecretAddress = new Pulumi.Docker.Provider("provider-with-sensitive-address", new Pulumi.Docker.ProviderArgs
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

   var password = new Pulumi.Random.RandomPassword("password", new()
   {
       Length = 16,
       Special = false,
   });

   // Test handling dynamic secrets that start as unknown.
   var providerWithSecretUsername = new Pulumi.Docker.Provider("provider-with-sensitive-username", new Pulumi.Docker.ProviderArgs
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

   // Test handling secrets marked in the schema.
   var providerWithSecretPassword = new Pulumi.Docker.Provider("provider-with-password", new Pulumi.Docker.ProviderArgs
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

   return new Dictionary<string, object?>{
       ["password"] = Output.Unsecret(password.Result).Apply(s => Convert.ToBase64String(Encoding.UTF8.GetBytes(s))),
   };
});
