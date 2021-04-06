#!/bin/sh

if ! grep -q Ubuntu /etc/os-release; then
	echo "detected distribution is not ubuntu, exiting..."
	exit 1
fi

apt-get update && apt-get install -y \
	build-essential \
	jq \
	libssl-dev \
	libudev-dev \
	make \
	ocl-icd-opencl-dev \
	pkg-config
