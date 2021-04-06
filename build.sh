#!/bin/sh

go get github.com/xlab/c-for-go

cd ./extern/filecoin-ffi/ && make
cd ../solana-ffi/ && make
cd ../../ && go build main.go
