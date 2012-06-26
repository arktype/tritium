// Code generated by protoc-gen-go from "tritium_test.proto"
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto1.GetString
var _ = math.Inf

type TritiumTest struct {
	Script           *string             `protobuf:"bytes,1,req,name=script" json:"script,omitempty"`
	Input            *string             `protobuf:"bytes,2,req,name=input" json:"input,omitempty"`
	Output           *string             `protobuf:"bytes,3,req,name=output" json:"output,omitempty"`
	EnvProto         []*TritiumTest_Hash `protobuf:"bytes,4,rep,name=env_proto" json:"env_proto,omitempty"`
	ExportsProto     []*TritiumTest_Hash `protobuf:"bytes,5,rep,name=exports_proto" json:"exports_proto,omitempty"`
	Logs             []string            `protobuf:"bytes,6,rep,name=logs" json:"logs,omitempty"`
	Transformer      *Transform          `protobuf:"bytes,7,opt,name=transformer" json:"transformer,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (this *TritiumTest) Reset()         { *this = TritiumTest{} }
func (this *TritiumTest) String() string { return proto1.CompactTextString(this) }

type TritiumTest_Hash struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *TritiumTest_Hash) Reset()         { *this = TritiumTest_Hash{} }
func (this *TritiumTest_Hash) String() string { return proto1.CompactTextString(this) }

func init() {
}
