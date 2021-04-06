#!/bin/sh

cd ./extern/filecoin-ffi/ && make
cd ../solana-ffi/ && make
cd ../../ && go build main.go
