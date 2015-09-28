package main

import (
  "log"
  "os"
  "bytes"
  "fmt"
  "time"
  "reflect"

  proto "github.com/golang/protobuf/proto"
  pb "github.com/zarcen/grpc_go/greeter"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

const (
  defaultAddress     = "localhost:50051"
)

func main() {
  address := defaultAddress
  if len(os.Args) > 1 {
    var buffer bytes.Buffer
    buffer.WriteString(fmt.Sprint(os.Args[1], ":50051"))
    address = buffer.String()
  }
  // Set up a connection to the server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewGreeterClient(conn)

  // Contact the server and print out its response.

  // sending a single int32
  // invoke marshaling the first time
  randomMsg := &pb.Greeting{Id:10, Fraction: 0.5}
  data, err := proto.Marshal(randomMsg)
  //
  intMsg :=  &pb.SimpleInt {Num: 50}
  log.Printf("Packing a single int")
  data, err = proto.Marshal(intMsg)
  start := time.Now()
  data, err = proto.Marshal(intMsg)
  elasped := time.Since(start)
  log.Printf("After packing -> %i",data)
  log.Printf("Marshaling SimpleInt took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  fmt.Printf("\n")
  // sending the request to server
  intReply, err := c.SendInt(context.Background(), intMsg)
  _ = intReply
  if err != nil {
    log.Fatalf("could not greet with int32: %v", err)
  }

  // sending a single doulbe(float64)
  doubleMsg :=  &pb.SimpleDouble {Num: 0.12345}
  log.Printf("Packing a single double")
  start = time.Now()
  data, err = proto.Marshal(doubleMsg)
  elasped = time.Since(start)
  log.Printf("After packing -> %i", data)
  log.Printf("Marshaling SimpleDouble took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  fmt.Printf("\n")
  // sending the request to server
  doubleReply, err := c.SendDouble(context.Background(), doubleMsg)
  _ = doubleReply
  if err != nil {
    log.Fatalf("could not greet with double: %v", err)
  }

  // sending a single string
  strMsg :=  &pb.SimpleString {Msg: "Please give me an offer!"}
  log.Printf("Packing a single string(len=25)")
  start = time.Now()
  data, err = proto.Marshal(strMsg)
  elasped = time.Since(start)
  log.Printf("After packing -> %i",data)
  log.Printf("Marshaling SimpleString took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  fmt.Printf("\n")
  // sending the request to server
  strReply, err := c.SendString(context.Background(), strMsg)
  _ = strReply
  if err != nil {
    log.Fatalf("could not greet with string: %v", err)
  }

  // sending a more complicated structure: Greeting
  greeting := &pb.Greeting{Id: 4, Fraction: 0.9527, Name: "Superman"}
  log.Printf("Packing a Greeting message (%d, %f, %s)", greeting.Id, greeting.Fraction, greeting.Name)
  start = time.Now()
  data, err = proto.Marshal(greeting)
  elasped = time.Since(start)
  log.Printf("After packing -> %i",data)
  log.Printf("Marshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  newgreeting := &pb.Greeting{}
  start = time.Now()
  err = proto.Unmarshal(data, newgreeting)
  elasped = time.Since(start)
  log.Printf("Unmarshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("unmarshaling error: ", err)
  }
  // greeting and newgreeting should contain the same data.
  if !reflect.DeepEqual(greeting, newgreeting) {
    log.Fatalf("Greeting data mismatch")
  }

  // Try marshaling the same Greeting struct the second time
  log.Printf("Packing *again* a Greeting message (%d, %f, %s)", greeting.Id, greeting.Fraction, greeting.Name)
  start = time.Now()
  data, err = proto.Marshal(greeting)
  elasped = time.Since(start)
  log.Printf("After packing -> %i",data)
  log.Printf("Marshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  newgreeting = &pb.Greeting{}
  start = time.Now()
  err = proto.Unmarshal(data, newgreeting)
  elasped = time.Since(start)
  log.Printf("Unmarshaling Greeting took %s", elasped)
  if err != nil {
    log.Fatal("unmarshaling error: ", err)
  }
  // sending the greeting to server
  reply, err := c.SendGreeting(context.Background(), greeting)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  log.Printf("Received a Greeting message back (%d, %f, %s)", reply.Id, reply.Fraction, reply.Name)
  start = time.Now()
  data, err = proto.Marshal(reply)
  elasped = time.Since(start)
  log.Printf("After packing -> %i",data)
  log.Printf("Marshaling Reply-Greeting took %s", elasped)
  if err != nil {
    log.Fatal("marshaling error: ", err)
  }
  newReply := &pb.Greeting{}
  start = time.Now()
  err = proto.Unmarshal(data, newReply)
  elasped = time.Since(start)
  log.Printf("Unmarshaling Reply-Greeting took %s", elasped)
  if err != nil {
    log.Fatal("unmarshaling error: ", err)
  }
  fmt.Printf("\n")
}
