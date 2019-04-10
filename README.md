# Hello Minikube

This is to show how to deploy a simple Go website into Minikube.

## Prerequisites

* [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

## App

Details of the application can be found [here](app/README.md).

## Minikube

### Start

```bash
$ minikube start
😄  minikube v1.0.0 on darwin (amd64)
🤹  Downloading Kubernetes v1.14.0 images in the background ...
💡  Tip: Use 'minikube start -p <name>' to create a new cluster, or 'minikube delete' to delete this one.
🔄  Restarting existing virtualbox VM for "minikube" ...
⌛  Waiting for SSH access ...
📶  "minikube" IP address is 192.168.99.100
🐳  Configuring Docker as the container runtime ...
🐳  Version of container runtime is 18.06.2-ce
⌛  Waiting for image downloads to complete ...
✨  Preparing Kubernetes environment ...
🚜  Pulling images required by Kubernetes v1.14.0 ...
🔄  Relaunching Kubernetes v1.14.0 using kubeadm ...
⌛  Waiting for pods: apiserver proxy etcd scheduler controller dns
📯  Updating kube-proxy configuration ...
🤔  Verifying component health ......
💗  kubectl is now configured to use "minikube"
🏄  Done! Thank you for using minikube!
```

### Status

```bash
$ minikube status
host: Running
kubelet: Running
apiserver: Running
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.99.100
```

### Run Application

```bash
$ kubectl run hello-minikube --image=denhamparry/hellominikube:0.1 --port=5000 --image-pull-policy=Always
deployment.apps/hello-minikube created
```

```bash
$ kubectl get deployments
NAME                       READY     UP-TO-DATE   AVAILABLE   AGE
hello-minikube             0/1       1            0           6s
```

```bash
$ kubectl get pods
kubectl get pods
NAME                                        READY     STATUS              RESTARTS   AGE
hello-minikube-8667d876d7-fwr72             0/1       ContainerCreating   0          5s
```

```bash
$ kubectl get pods
kubectl get pods
NAME                                        READY     STATUS              RESTARTS   AGE
hello-minikube-8667d876d7-fwr72             1/1       Running             0          57s
```

```bash
$ kubectl expose deployment hello-minikube --type="NodePort"
service/hello-minikube exposed
```

```bash
$  kubectl get service
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-minikube   NodePort    10.110.126.57   <none>        5000:30762/TCP   19s
```

```bash
$ minikube service hello-minikube
🎉  Opening kubernetes service default/hello-minikube in default browser...
```

```bash
$ curl http://192.168.99.100:30762
Hello Minikube! (v0.1)
```

### Better way to run the application

> TODO: yaml

### Delete

```bash
minikube delete
🔥  Deleting "minikube" from virtualbox ...
💔  The "minikube" cluster has been deleted.
```

## Notes

* Could build Docker images on Minikube Docker Daemon rather than pull from Docker Hub:

```bash
$ minikube docker-env
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.100:2376"
export DOCKER_CERT_PATH="/Users/denhamparry/.minikube/certs"
export DOCKER_API_VERSION="1.35"
# Run this command to configure your shell:
# eval $(minikube docker-env)
$ eval $(minikube docker-env)
$ docker ps
```