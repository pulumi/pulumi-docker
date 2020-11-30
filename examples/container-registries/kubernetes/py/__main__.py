import base64
import json
import pulumi
import pulumi_docker as docker
from pulumi_kubernetes.apps.v1 import Deployment
from pulumi_kubernetes.core.v1 import Secret, Service

# Fetch the Docker Hub auth info from config.
config = pulumi.Config()
username = config.require('dockerUsername')
password = config.require_secret('dockerPassword')

# Build and publish the image.
image = docker.Image('my-image',
    build='app',
    image_name=f'{username}/myapp',
    registry=password.apply(lambda pwd: docker.ImageRegistry(
        server='docker.io',
        username=username,
        password=pwd,
    )),
)

# Ensure we can pull from the Docker Hub.
pull_secret = Secret('my-regcred',
    type='kubernetes.io/dockerconfigjson',
    string_data={
        '.dockerconfigjson': password.apply(lambda pwd: json.dumps({
            'auths': {
                'https://index.docker.io/v1/': {
                    'username': username,
                    'password': pwd,
                    'auth': str(base64.b64encode(f'{username}:{pwd}'.encode('utf-8')), 'utf-8'),
                },
            },
        })),
    },
)

# Deploy a load-balanced service that uses this image.
labels = { 'app': 'my-app' }
dep = Deployment('my-app-dep', spec={
    'selector': { 'matchLabels': labels },
    'replicas': 1,
    'template': {
        'metadata': { 'labels': labels },
        'spec': {
            'containers': [{
                'name': labels['app'],
                'image': image.image_name,
            }],
            'image_pull_secrets': [{
                'name': pull_secret.metadata['name'],
            }],
        },
    },
})
svc = Service('my-app-svc', spec={
    'selector': labels,
    'type': 'LoadBalancer',
    'ports': [{ 'port': 80 }],
})

# Export the resulting image name.
pulumi.export('imageName', image.image_name)
# Export the k8s ingress IP to access the service.
pulumi.export('serviceIp', svc.status['load_balancer']['ingress'][0]['ip'])
