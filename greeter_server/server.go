package main

import (
  "log"
  "net"
  "time"
  //"io"

  proto "github.com/golang/protobuf/proto"
  pb "github.com/zarcen/grpc_go/greeter"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

const (
  port = ":50051"
)

// greeterServer is used to implement cs739p1.GreeterServer.
type greeterServer struct{
	savedReplies []*pb.Greeting
	savedRequest []*pb.Greeting
}

func (s *greeterServer) SendInt(ctx context.Context, in *pb.SimpleInt) (*pb.SimpleInt, error) {
  reply := &pb.SimpleInt{}
  return reply, nil
}

func (s *greeterServer) SendDouble(ctx context.Context, in *pb.SimpleDouble) (*pb.SimpleDouble, error) {
  reply := &pb.SimpleDouble{}
  return reply, nil
}

func (s *greeterServer) SendString(ctx context.Context, in *pb.SimpleString) (*pb.SimpleString, error) {
  reply := &pb.SimpleString{}
  return reply, nil
}

// SendGreeting implements greeter.GreeterServer
func (s *greeterServer) SendGreeting(ctx context.Context, in *pb.Greeting) (*pb.Greeting, error) {
  reply := &pb.Greeting{Id: in.Id+1, Fraction: in.Fraction-0.1, Name: in.Name}
  start := time.Now()
  data, err := proto.Marshal(reply)
  elasped := time.Since(start)
  log.Printf("reply data -> %i",data)
  log.Printf("Marshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  newReply := &pb.Greeting{}
  start = time.Now()
  err = proto.Unmarshal(data, newReply)
  elasped = time.Since(start)
  log.Printf("Unmarshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("unmarshaling error: ", err)
  }
  start = time.Now()
  data, err = proto.Marshal(in)
  elasped = time.Since(start)
  log.Printf("request data -> %i",data)
  log.Printf("Marshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  newRequest := &pb.Greeting{}
  start = time.Now()
  err = proto.Unmarshal(data, newRequest)
  elasped = time.Since(start)
  log.Printf("Unmarshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("unmarshaling error: ", err)
  }
  return reply, nil
}

func (s *greeterServer) ServerStreaming(request *pb.Greeting, stream pb.Greeter_ServerStreamingServer) error {
  // to be implemented
  return nil
}

func (s *greeterServer) ClientStreaming(stream pb.Greeter_ClientStreamingServer) error {
  // to be implemented
  return nil
}

func (s *greeterServer) BidirectionalStreaming(stream pb.Greeter_BidirectionalStreamingServer) error {
  /*for {
    in, err := stream.Recv()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      return err
    }
    // do some processing for "in" here while streaming
  }*/
  return nil
}

func newServer() *greeterServer {
	s := new(greeterServer)
	return s
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, newServer())
  s.Serve(lis)
}
