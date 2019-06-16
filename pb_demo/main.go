package main

import (
    "github.com/golang/protobuf/proto"
    "github.com/kataras/golog"
    "github.com/m01i0ng/minx-demo/pb_demo/pb"
)

func main() {
    person := &pb.Person{
        Name:   "minx",
        Age:    18,
        Emails: []string{"a@b.com"},
        Phones: []*pb.PhoneNumber{
            {
                Number: "1222",
                Type:   pb.PhoneType_HOME,
            },
        },
    }

    //编码
    bytes, _ := proto.Marshal(person)
    golog.Infof("%s", bytes)

    //解码
    i := &pb.Person{}
    _ = proto.Unmarshal(bytes, i)
    golog.Infof("%v", i)
}
