# CLI for managing a cluster

## Setup
From a linux host:
```
# Install gcloud CLI
curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-460.0.0-linux-x86_64.tar.gz
tar -xf google-cloud-cli-460.0.0-linux-x86_64.tar.gz
./google-cloud-sdk/install.sh
# Reload .bashrc

# Authenticate and configure gcloud
gcloud auth login --no-launch-browser
# Do the needful with the resulting URL
gcloud config set project saas-test-sre-interview-410711
gcloud components install kubectl

# Install go
curl -O https://dl.google.com/go/go1.21.6.linux-amd64.tar.gz
rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.*
export PATH=$PATH:/usr/local/go/bin

# Clone sources and build binaries
git clone https://github.com/scottkaplan/srecli.git
cd srecli
go mod init srecli
go mod tidy
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
