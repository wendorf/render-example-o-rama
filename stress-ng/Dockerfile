FROM ubuntu

RUN apt-get update && apt-get install -y stress-ng
RUN stress-ng --cpu 1 --vm 1 --vm-bytes 512g --timeout 10s --vm-keep
