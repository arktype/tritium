// Code generated by protoc-gen-go.
// source: package.proto
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto1.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Type struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Implements       *int32  `protobuf:"varint,2,opt,name=implements" json:"implements,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto1.CompactTextString(m) }
func (*Type) ProtoMessage()    {}

func (m *Type) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Type) GetImplements() int32 {
	if m != nil && m.Implements != nil {
		return *m.Implements
	}
	return 0
}

type Package struct {
	Name             *string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Functions        []*Function `protobuf:"bytes,2,rep,name=functions" json:"functions,omitempty"`
	Types            []*Type     `protobuf:"bytes,3,rep,name=types" json:"types,omitempty"`
	Dependencies     []string    `protobuf:"bytes,4,rep,name=dependencies" json:"dependencies,omitempty"`
	Path             *string     `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto1.CompactTextString(m) }
func (*Package) ProtoMessage()    {}

func (m *Package) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Package) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *Package) GetTypes() []*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Package) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Package) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func init() {
}