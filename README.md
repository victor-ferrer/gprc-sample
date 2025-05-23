Sample microservice built in Go 

## Running a local K8s cluster

```
# Create the ctlptl registry for your images
ctlptl create cluster kind --registry=ctlptl-registry

# Spin up tilt
tilt up

```