simple gRPC server and client for calculating simple math tasks

# system requirements
* Unix-like operating system with bash
* go1.9.4 but should work with any version past 1.5
* dep (get with curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh)
* protoc-gen-go (alternatively installed by installer script)

# installation instructions
* run ```go get github.com/rolandktud/calculator```.
* run ```./install.sh``` from the directory ```$GOPATH/src/github.com/rolandktud/calculator/```. This script compiles the protobuf file and builds both binaries. Set ```GOBIN``` to set output directories. Example: ```cd $GOPATH/src/github.com/rolandktud/calculator/; GOBIN=/tmp; ./install.sh```

# run instructions
* run server with ```$GOBIN/server``` (runs in foreground)
* run client with ```$GOBIN/client```.
Listening ports and addresses can be set. Use the -help parameter of the programs for details
* alternatively, run with ```go run client/main.go``` and ```go run server/main.go```, from the directory, respectively


# open issues
* Add the protobuf compiler as a dependency and save it in the workspace. Currently it's saved in the global GOBIN directory. To my best knowledge, godep does not handle this use case

(disclaimer: I learned go for this in about half a workday)

