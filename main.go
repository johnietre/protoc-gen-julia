package main

import (
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
  log.SetFlags(0)
  // Read and unmarshal the request
  buf, err := io.ReadAll(os.Stdin)
  if err != nil {
    log.Fatal(err)
  }
  req := &pluginpb.CodeGeneratorRequest{}
  if err = proto.Unmarshal(buf, req); err != nil {
    log.Fatal(err)
  }
  resp := &pluginpb.CodeGeneratorResponse{}

  // TODO: Check compiler version for optional
  // TODO: Check parameters

  protoFiles := make(
    map[string]*descriptorpb.FileDescriptorProto, len(req.ProtoFile),
  )
  for _, file := range req.ProtoFile {
    // TODO: Handle files without names
    if name := file.GetName(); name != "" {
      protoFiles[name] = file
    }
  }
  for _, filename := range req.FileToGenerate {
  }

  // Write the response
  if buf, err = proto.Marshal(resp); err != nil {
    log.Fatal(err)
  }
  if _, err = os.Stdout.Write(buf); err != nil {
    log.Fatal(err)
  }
}

func printReq(req *pluginpb.CodeGeneratorRequest) {
  buf, err := (protojson.MarshalOptions{Multiline: true}).Marshal(req)
  if err != nil {
    log.Fatal(err)
  }
  if err = os.WriteFile("out.json", buf, 0755); err != nil {
    log.Fatal(err)
  }
  log.Fatal("done")
}
