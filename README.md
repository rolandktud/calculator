simple gRPC server and client for calculating simple math tasks

# system requirements
* Unix-like operating system with bash
* go1.9.4 but should work with any version past 1.5
* dep (get with curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh)
* protoc-gen-go (alternatively installed by installer script)
# installation instructions
* run ```go get github.com/rolandktud/calculator``` 
* run ```./install.sh``` from the directory ```$GOPATH/src/github.com/rolandktud/calculator/```. This script compiles the protobuf file and builds both binaries. Set ```$GOBIN``` to set output directories.
* run server with ```$GOBIN/server``` (runs in foreground)
* run client with ```$GOBIN/client```.
Listening ports and addresses can be set. Use the -help parameter of the programs for details