docker_build('localhost:37659/gprc-sample','.',dockerfile='Dockerfile')

k8s_yaml([
    'k8s/postgres-deployment.yaml',
    'k8s/app-deployment.yaml'
])

k8s_resource('grpc-sample', port_forwards=8080)