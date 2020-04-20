#!/bin/bash
  sudo apt-get update && \
  sudo apt install iproute2 && \
  sudo apt install net-tools && \
  sudo apt-get install network-manager &&\
  wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz && \
  tar -xvf go1.13.linux-amd64.tar.gz && \
  mv go /usr/local && \
  echo export PATH=/usr/local/go/bin:$PATH >> ~/bashrc && \
  go run cmd/main.go