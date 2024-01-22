# CLI for managing a cluster

## Setup
From a linux host:
```
gcloud auth login --no-launch-browser
gcloud config set project saas-test-sre-interview-410711
gcloud components install kubectl
```

## Run

```
sre list # List pods in the cluster
sre list --namespace scott #  List pods in namespace scott
sre info --namespace scott #  Print info about pods in namespace scott
sre info --pod healthy-pod #  Print info about healthy-pod
sre scale --replicas 3 --deployment prod-deploy # Run 3 pods in the prod-deploy deployment
```

## To Do
- Document scale
- README.md
