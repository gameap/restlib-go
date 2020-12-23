#/bin/bash

go build -buildmode=c-shared -o restclient.so ../restclient.go
