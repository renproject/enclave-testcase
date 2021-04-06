# Install
Clone repo and submodules
```
https://github.com/renproject/enclave-testcase.git
cd enclave-testcase
git submodule update --init --recursive
```

Make sure that Go and Rust are installed. To build the executable, run
```
./build.sh
```
If local dependencies are missing (see list below), you will need to install them first. If you are using Ubuntu, the script `install-deps.sh` is provided to download the needed dependencies.

# Requirements
- make
- jq
- rust
- curl
- build-essential
- openssl
- pkg-config
- go
- c-for-go
- OpenCL (ocl-icd-opencl-dev on ubuntu)
