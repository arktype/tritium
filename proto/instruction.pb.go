// Code generated by protoc-gen-go from "instruction.proto"
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto1.GetString
var _ = math.Inf

type Instruction_InstructionType int32

const (
	Instruction_BLOCK         Instruction_InstructionType = 0
	Instruction_FUNCTION_CALL Instruction_InstructionType = 1
	Instruction_IMPORT        Instruction_InstructionType = 2
	Instruction_TEXT          Instruction_InstructionType = 3
	Instruction_LOCAL_VAR     Instruction_InstructionType = 4
	Instruction_POSITION      Instruction_InstructionType = 5
	Instruction_COMMENT       Instruction_InstructionType = 6
)

var Instruction_InstructionType_name = map[int32]string{
	0: "BLOCK",
	1: "FUNCTION_CALL",
	2: "IMPORT",
	3: "TEXT",
	4: "LOCAL_VAR",
	5: "POSITION",
	6: "COMMENT",
}
var Instruction_InstructionType_value = map[string]int32{
	"BLOCK":         0,
	"FUNCTION_CALL": 1,
	"IMPORT":        2,
	"TEXT":          3,
	"LOCAL_VAR":     4,
	"POSITION":      5,
	"COMMENT":       6,
}

// NewInstruction_InstructionType is deprecated. Use x.Enum() instead.
func NewInstruction_InstructionType(x Instruction_InstructionType) *Instruction_InstructionType {
	e := Instruction_InstructionType(x)
	return &e
}
func (x Instruction_InstructionType) Enum() *Instruction_InstructionType {
	p := new(Instruction_InstructionType)
	*p = x
	return p
}
func (x Instruction_InstructionType) String() string {
	return proto1.EnumName(Instruction_InstructionType_name, int32(x))
}

type Instruction struct {
	Type             *Instruction_InstructionType `protobuf:"varint,1,req,name=type,enum=proto.Instruction_InstructionType" json:"type,omitempty"`
	Value            *string                      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	ObjectId         *int32                       `protobuf:"varint,3,opt,name=object_id" json:"object_id,omitempty"`
	Children         []*Instruction               `protobuf:"bytes,4,rep,name=children" json:"children,omitempty"`
	Arguments        []*Instruction               `protobuf:"bytes,5,rep,name=arguments" json:"arguments,omitempty"`
	FunctionId       *int32                       `protobuf:"varint,6,opt,name=function_id" json:"function_id,omitempty"`
	LineNumber       *int32                       `protobuf:"varint,7,opt,name=line_number" json:"line_number,omitempty"`
	YieldTypeId      *int32                       `protobuf:"varint,8,opt,name=yield_type_id" json:"yield_type_id,omitempty"`
	IsValid          *bool                        `protobuf:"varint,9,opt,name=is_valid" json:"is_valid,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (this *Instruction) Reset()         { *this = Instruction{} }
func (this *Instruction) String() string { return proto1.CompactTextString(this) }

func init() {
	proto1.RegisterEnum("proto.Instruction_InstructionType", Instruction_InstructionType_name, Instruction_InstructionType_value)
}
