# Minikube NFS

[![Build Status](https://travis-ci.org/mstrzele/minikube-nfs.svg?branch=master)](https://travis-ci.org/mstrzele/minikube-nfs)

Enables and configures NFS server daemon on macOS and defines remote mount points for the Minikube VM.

## Requirements

* macOS
* minikube v0.18.0

## Installation

```bash
$ curl -LOSs https://github.com/mstrzele/minikube-nfs/releases/download/v0.1.0/minikube-nfs
$ chmod +x minikube-nfs
$ sudo mv minikube-nfs /usr/local/bin
```


## Usage

```bash
$ minikube start
Starting local Kubernetes v1.6.4 cluster...
Starting VM...
Moving files into cluster...
Setting up certs...
Starting cluster components...
Connecting to cluster...
Setting up kubeconfig...
Kubectl is now configured to use the cluster.
$ sudo minikube-nfs -n"-alldirs -mapall=$(id -u):$(id -g)"
Password:
INFO[0000] machine presence ...                          clusterState=Running
INFO[0001] machine running ...                           clusterState=Running
INFO[0001] Lookup mandatory properties ...               clusterIP=192.168.99.100
INFO[0001] Configure NFS ...                             machineIP=192.168.99.100 nfsConfig="-alldirs -mapall=501:20" sharedFolders=[/Users] useIPRange=false
$ cat <<EOF | kubectl create -f -
apiVersion: v1
kind: PersistentVolume
metadata:
  name: users
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  nfs:
    path: /Users
    server: 192.168.99.1
EOF
persistentvolume "users" created
```

## Performance Comparison

NFS seems to be **10x faster** than VirtualBox filesystem.

Tested using [IOzone](http://www.iozone.org/) 3.434.

Logs are available in `iozone-vboxfs.log` and `iozone-nfs.log` files.

Configuraiton can be seen in `iozone-vboxfs-job.yaml` and `iozone-nfs-job.yaml` files.

1. Start Minikube.
1. Apply `iozone-vboxfs-job.yaml`.
1. See the logs from the pod.
1. Configure NFS using `minikube-nfs`.
1. Apply `iozone-nfs-job.yaml`.
1. See the logs from the pod.
