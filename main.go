package main

import (
    "bytes"
    "compress/gzip"
    "encoding/json"
    "fmt"

    "io/ioutil"

    proto "github.com/gogo/protobuf/proto"
    dpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
    _ "github.com/ludete/testpb/pb"
)

func main() {
    // here write the path that is used in the generated file
    // in init(), as an argument to proto.RegisterFile 
    // (or just copypaste the bytes instead of using proto.FileDescriptor)
    bytes := proto.FileDescriptor("student.proto")
    fd, err := decodeFileDesc(bytes)
    if err != nil {
        panic(err)
    }
    b, err := json.MarshalIndent(fd,"","  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}

func decodeFileDesc(enc []byte) (*dpb.FileDescriptorProto, error) {
    raw, err := decompress(enc)
    if err != nil {
        return nil, fmt.Errorf("failed to decompress enc: %v", err)
    }

    fd := new(dpb.FileDescriptorProto)
    if err := proto.Unmarshal(raw, fd); err != nil {
        return nil, fmt.Errorf("bad descriptor: %v", err)
    }
    return fd, nil
}

// decompress does gzip decompression.
func decompress(b []byte) ([]byte, error) {
    r, err := gzip.NewReader(bytes.NewReader(b))
    if err != nil {
        return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
    }
    out, err := ioutil.ReadAll(r)
    if err != nil {
        return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
    }
    return out, nil
}


