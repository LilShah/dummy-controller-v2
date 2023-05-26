# Dummy Controller

## Overview

Dummy Controller is a simple example Kubernetes controller that echoes the message written in its spec to the status.

```yaml
apiVersion: interview.com/v1alpha1
kind: Dummy
metadata:
  name: dummy1
  namespace: example
spec:
  message: "I'm just a dummy"
```

The status gets updated as such:

```yaml
status:
  specEcho: "I'm just a dummy"
```

## How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/)
which provides a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster

## Installing Kind

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) ([Installation](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)) to get a local cluster for testing, or run against a remote cluster.

After installing kind, you can start it by using `kind create cluster`

**Note:** The kubernetes context should automatically switch to using `Kind`, but make sure to check with `kubectl config current-context`.

## Test the controller

The simplest method to install the controller for testing is to clone the [dummy-controller repo](https://github.com/LilShah/dummy-controller-v2) and from within it, run the following command:

```sh
make deploy
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

Give it a few minutes to download and create the pod. You can check pod status by using:

```sh
kubectl get pods -n dummy-controller-v2-system
```

A sample Dummy CR is also provided within the repo:

```sh
kubectl apply -f config/samples/interview_v1alpha1_dummy.yaml
```

**NOTE:**: You will need [kubectl](https://kubernetes.io/docs/tasks/tools/) installed to run the above commands.

## Uninstall

### Uninstall CRDs

To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller

Undeploy the controller from the cluster:

```sh
make undeploy
```

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

```bash
http://www.apache.org/licenses/LICENSE-2.0
```

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
