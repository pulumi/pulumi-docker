# SSH Connection Integration Test

## Setup

This test takes a bit of manual setup to work properly.

In order to verify that a remote Docker host via SSH works correctly, we are setting up the following test stack:

```bash
 +   pulumi:pulumi:Stack
 +   ├─ digitalocean:index:Droplet
 +   ├─ pulumi:providers:docker
 +   ├─ docker:index:RemoteImage
 +   └─ docker:index:Container
```

We are imaging a Docker/Ubuntu Digitalocean Droplet as a remote docker host to pull a RemoteImage and run a Container. 
We declare an explicit Docker provider so that we can configure it to use the IP output of the Droplet as our Docker host.

1. Generate an ssh rsa key pair [according to GitHub's instructions](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent), using your github email.
1. Add the private key to the repository as `PRIVATE_SSH_KEY_FOR_DIGITALOCEAN`
1. Upload the _public key_ to pulumi-bot's [digitalocean account](https://docs.digitalocean.com/products/droplets/how-to/add-ssh-keys/to-team/)
1. Digitalocean will give you a fingerprint. Use this fingerprint to populate the `sshKeys` field of the DigitalOcean Droplet. This will ensure that the new Droplet will be created with the public key in its `authorized_keys` file.

Caveats:

- We connect as `root` because the Droplet does not come preconfigured with users.
