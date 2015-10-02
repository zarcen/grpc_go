# grpc_go
grpc evaluation using go language

# Graph, Tables, Snapshots
![1](/img/1.png)
![2](/img/2.png)
![3](/img/3.png)
![4](/img/4.png)
![5](/img/5.png)
![6](/img/6.png)
![7](/img/7.png)
![8](/img/8.png)
![9](/img/9.png)
![10](/img/10.png)
![11](/img/11.png)

# Notes
- xxx.proto:16:12: Explicit 'optional' labels are disallowed in the Proto3 syntax. To define 'optional' fields in Proto3, simply remove the 'optional' label, as fields│
 are 'optional' by default.

- [How to avoid annoying error “declared and not used” from golang] (http://stackoverflow.com/questions/21743841/how-to-avoid-annoying-error-declared-and-not-used-from-golang)

- Under the $GOPATH/src/github.com/zarcen/grpc_go/, run ``$ protoc -I greeter greeter/greeter.proto --go_out=plugins=grpc:greeter``
  It can parse the protobuf and generate the related source code.

# Getting Started (Install, Run)
To install and run it, change the working directory to $GOPATH and do the following:
  - Build and install client: ``go install github.com/zarcen/grpc_go/greeter_client``
  - Build and install server: ``go install github.com/zarcen/grpc_go/greeter_server``
  - If you update anything in `greeter/greeter.proto`, run ``$ protoc -I greeter greeter/greeter.proto --go_out=plugins=grpc:greeter`` again
  - Run the server on one machine (default port:50051), ``$ greeter_server``
  - Run the client on another machine (or the same one), ``$ greeter_client``

# TODO
1. Streaming interface implementation
2. Modify the client code such that it evaluates and logs information for plotting the diagram
