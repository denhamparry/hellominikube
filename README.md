# Hello Minikube

This is to show how to deploy a simple Go website into Minikube.

## Description

Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop for users looking to try out Kubernetes or develop with it day-to-day.

This tutorial looks at how to spin up a simple `Hello Minikube` Go web application inside a Minikube Kubernetes cluster.
We will look at a simple way to run an application using `Kubectl`, a command line interface for running commands against Kubernetes clusters.
Once we have the application running, we'll look at another way we can declare what is running in the cluster by using `yaml` files that can be version controlled.
The last step will be to update the application without any downtime.

## Prerequisites

* [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

## App

Details of the application can be found [here](app/README.md).

## Minikube

### Start

`minikube start` will initialise a Kubernetes cluster on your machine.  The cluster is created within a single Virtual Machine (vm).
Minikube has its own Docker Daemon located within the VM, and launches a Kubernetes cluster using `kubeadm`.  More details about `kubeadm` can be found [here](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/).
One the cluster is initialised, we can connect the Kubernetes Control Plane via `kubectl`.

```bash
$  minikube start
😄  minikube v1.0.0 on darwin (amd64)
🤹  Downloading Kubernetes v1.14.0 images in the background ...
🔥  Creating virtualbox VM (CPUs=2, Memory=2048MB, Disk=20000MB) ...
📶  "minikube" IP address is 192.168.99.100
🐳  Configuring Docker as the container runtime ...
🐳  Version of container runtime is 18.06.2-ce
⌛  Waiting for image downloads to complete ...
✨  Preparing Kubernetes environment ...
🚜  Pulling images required by Kubernetes v1.14.0 ...
🚀  Launching Kubernetes v1.14.0 using kubeadm ...
⌛  Waiting for pods: apiserver proxy etcd scheduler controller dns
🔑  Configuring cluster permissions ...
🤔  Verifying component health .....
💗  kubectl is now configured to use "minikube"
🏄  Done! Thank you for using minikube!
```

### Status

If we want to check the status of Minikube, we can run `minikube status` to confirm that the environment is setup correctly.

```bash
$ minikube status
host: Running
kubelet: Running
apiserver: Running
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.99.100
```

### Run Application

Lets start by deploying our application to the Kubernetes Cluster using `kubectl run`.

We are stating that

* The deployment name is __hello-minikube__.
* That the Docker image to run in the deployment is __denhamparry/hellominikube:0.1__.
* To allow traffic onto __port 5000__ within the container.
* To always pull the image.

```bash
$ kubectl run hello-minikube --image=denhamparry/hellominikube:0.1 --port=5000 --image-pull-policy=Always
deployment.apps/hello-minikube created
```

`kubectl run` has created a __deployment__ within the Kubernetes cluster.

> NOTE: Deployments represent a set of multiple, identical Pods.  More information can be found [here](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/).

We have declared to Kubernetes that we want a deployment called __hello-minikube__, lets check the status of the deployment.

```bash
$ kubectl get deployments
NAME                       READY     UP-TO-DATE   AVAILABLE   AGE
hello-minikube             0/1       1            0           6s
```

__0/1__ pods are ready, lets check the status of the pod:

> NOTE: Pods are the smallest deployable units of computing that can be created and managed in Kubernetes. More information can be found [here](https://kubernetes.io/docs/concepts/workloads/pods/pod/).

```bash
$ kubectl get pods
kubectl get pods
NAME                                        READY     STATUS              RESTARTS   AGE
hello-minikube-8667d876d7-fwr72             0/1       ContainerCreating   0          5s
```

The pod has a __STATUS__ of __ContainerCreating__, this means that the pod is still initialising (e.g. pulling a Docker image).

> TIP: Use the `-w` flag to watch the pods, e.g. `kubectl get pods -w`.  To exit, _ctl-c_

Lets check to see if the pod is ready...

```bash
$ kubectl get pods
kubectl get pods
NAME                                        READY     STATUS              RESTARTS   AGE
hello-minikube-8667d876d7-fwr72             1/1       Running             0          57s
```

We can now see that __1/1__ pods are ready and the __STATUS__ is __Running__.

Now we want to be able to access the Go web application from a browser, to do that we need to make the website publicly accessible.

```bash
$ kubectl expose deployment hello-minikube --type="NodePort"
service/hello-minikube exposed
```

To be able to access the deployment outside of the cluster, one option is to use a __NodePort__.

> NOTE: A NodePort exposes the service on each Node's IP at a static port.  More information can be found [here](https://kubernetes.io/docs/concepts/services-networking/service/#nodeport).

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

> NOTE: The port assigned is automatically generated, check your `NodePort` via `kubectl get service`.

```bash
$ kubectl delete server hello-minikube
service "hello-minikube" deleted
$ kubectl delete deployment hello-minikube
deployment.extensions "hello-minikube" deleted
$ kubectl get pods
No resources found.
```

### A better way to run the application

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-minikube
  labels:
    app: hello-minikube
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello-minikube
  template:
    metadata:
      labels:
        app: hello-minikube
    spec:
      containers:
      - name: hello-minikube
        image: denhamparry/hellominikube:0.1
        ports:
        - containerPort: 5000
```

```bash
$  kubectl apply -f kubernetes/deployment.yaml
deployment.apps/hello-minikube created
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello-minikube
spec:
  selector:
    app: hello-minikube
  ports:
  - protocol: TCP
    port: 80
    targetPort: 5000
```

```bash
$ kubectl apply -f kubernetes/service.yaml
service/hello-minikube created
```

```bash
$  minikube addons enable ingress
✅  ingress was successfully enabled
```

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: hello-minikube
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /
        backend:
          serviceName: hello-minikube
          servicePort: 80
```

```bash
$ kubectl apply -f kubernetes/ingress.yaml
ingress.extensions/hello-minikube created
```

```bash
$ curl https://$(minikube ip)
Hello Minikube! (v0.1)
```

### Lets update the application with zero downtime

```bash
$ watch curl $(minikube ip)
Every 2.0s: curl 192.168.99.100
Hello Minikube! (v0.1)
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-minikube
  labels:
    app: hello-minikube
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello-minikube
  template:
    metadata:
      labels:
        app: hello-minikube
    spec:
      containers:
      - name: hello-minikube
        image: denhamparry/hellominikube:0.2
        ports:
        - containerPort: 5000
```

```bash
$ kubectl apply -f kubernetes/deployment.v2.yaml
deployment.apps/hello-minikube configured
```

```bash
Every 2.0s: curl 192.168.99.100
Hello Minikube! (v0.1)
```

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

* Could show use of namespaces:

```bash
$ kubectl create namespace hello-minikube
namespace/hello-minikube created
$ kubectl delete namespace hello-minikube
namespace "hello-minikube" deleted
```

* Could look into Kube context management:

```bash
$ ls ~/.kube
cache            config           http-cache       kind-config-kind kubectx          kubens           ~
```

* Talk about different kinds of deployments:
  * Rolling.
  * Blue / Green.
  * Canary.
* Look to define Minikube Architecture.
  * Its a single node compromising of master and worker roles.
  * Where should we add a definition of what a node is?