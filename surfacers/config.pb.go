// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/config.proto

/*
Package surfacers is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/surfacers/config.proto

It has these top-level messages:
	SurfacerDef
*/
package surfacers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_surfacer_prometheus "github.com/google/cloudprober/surfacers/prometheus"
import cloudprober_surfacer_stackdriver "github.com/google/cloudprober/surfacers/stackdriver"
import cloudprober_surfacer_file "github.com/google/cloudprober/surfacers/file"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enumeration for each type of surfacer we can parse and create
type Type int32

const (
	Type_NONE         Type = 0
	Type_PROMETHEUS   Type = 1
	Type_STACKDRIVER  Type = 2
	Type_FILE         Type = 3
	Type_USER_DEFINED Type = 99
)

var Type_name = map[int32]string{
	0:  "NONE",
	1:  "PROMETHEUS",
	2:  "STACKDRIVER",
	3:  "FILE",
	99: "USER_DEFINED",
}
var Type_value = map[string]int32{
	"NONE":         0,
	"PROMETHEUS":   1,
	"STACKDRIVER":  2,
	"FILE":         3,
	"USER_DEFINED": 99,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SurfacerDef struct {
	// This name is used for logging. If not defined, it's derived from the type.
	// Note that this field is required for the USER_DEFINED surfacer type and
	// should match with the name that you used while registering the user defined
	// surfacer.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type *Type   `protobuf:"varint,2,opt,name=type,enum=cloudprober.surfacer.Type" json:"type,omitempty"`
	// Matching surfacer specific configuration (one for each type in the above
	// enum)
	//
	// Types that are valid to be assigned to Surfacer:
	//	*SurfacerDef_PrometheusSurfacer
	//	*SurfacerDef_StackdriverSurfacer
	//	*SurfacerDef_FileSurfacer
	Surfacer         isSurfacerDef_Surfacer `protobuf_oneof:"surfacer"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SurfacerDef) Reset()                    { *m = SurfacerDef{} }
func (m *SurfacerDef) String() string            { return proto.CompactTextString(m) }
func (*SurfacerDef) ProtoMessage()               {}
func (*SurfacerDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isSurfacerDef_Surfacer interface {
	isSurfacerDef_Surfacer()
}

type SurfacerDef_PrometheusSurfacer struct {
	PrometheusSurfacer *cloudprober_surfacer_prometheus.SurfacerConf `protobuf:"bytes,10,opt,name=prometheus_surfacer,json=prometheusSurfacer,oneof"`
}
type SurfacerDef_StackdriverSurfacer struct {
	StackdriverSurfacer *cloudprober_surfacer_stackdriver.SurfacerConf `protobuf:"bytes,11,opt,name=stackdriver_surfacer,json=stackdriverSurfacer,oneof"`
}
type SurfacerDef_FileSurfacer struct {
	FileSurfacer *cloudprober_surfacer_file.SurfacerConf `protobuf:"bytes,12,opt,name=file_surfacer,json=fileSurfacer,oneof"`
}

func (*SurfacerDef_PrometheusSurfacer) isSurfacerDef_Surfacer()  {}
func (*SurfacerDef_StackdriverSurfacer) isSurfacerDef_Surfacer() {}
func (*SurfacerDef_FileSurfacer) isSurfacerDef_Surfacer()        {}

func (m *SurfacerDef) GetSurfacer() isSurfacerDef_Surfacer {
	if m != nil {
		return m.Surfacer
	}
	return nil
}

func (m *SurfacerDef) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SurfacerDef) GetType() Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Type_NONE
}

func (m *SurfacerDef) GetPrometheusSurfacer() *cloudprober_surfacer_prometheus.SurfacerConf {
	if x, ok := m.GetSurfacer().(*SurfacerDef_PrometheusSurfacer); ok {
		return x.PrometheusSurfacer
	}
	return nil
}

func (m *SurfacerDef) GetStackdriverSurfacer() *cloudprober_surfacer_stackdriver.SurfacerConf {
	if x, ok := m.GetSurfacer().(*SurfacerDef_StackdriverSurfacer); ok {
		return x.StackdriverSurfacer
	}
	return nil
}

func (m *SurfacerDef) GetFileSurfacer() *cloudprober_surfacer_file.SurfacerConf {
	if x, ok := m.GetSurfacer().(*SurfacerDef_FileSurfacer); ok {
		return x.FileSurfacer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SurfacerDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SurfacerDef_OneofMarshaler, _SurfacerDef_OneofUnmarshaler, _SurfacerDef_OneofSizer, []interface{}{
		(*SurfacerDef_PrometheusSurfacer)(nil),
		(*SurfacerDef_StackdriverSurfacer)(nil),
		(*SurfacerDef_FileSurfacer)(nil),
	}
}

func _SurfacerDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SurfacerDef)
	// surfacer
	switch x := m.Surfacer.(type) {
	case *SurfacerDef_PrometheusSurfacer:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrometheusSurfacer); err != nil {
			return err
		}
	case *SurfacerDef_StackdriverSurfacer:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StackdriverSurfacer); err != nil {
			return err
		}
	case *SurfacerDef_FileSurfacer:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileSurfacer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SurfacerDef.Surfacer has unexpected type %T", x)
	}
	return nil
}

func _SurfacerDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SurfacerDef)
	switch tag {
	case 10: // surfacer.prometheus_surfacer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_surfacer_prometheus.SurfacerConf)
		err := b.DecodeMessage(msg)
		m.Surfacer = &SurfacerDef_PrometheusSurfacer{msg}
		return true, err
	case 11: // surfacer.stackdriver_surfacer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_surfacer_stackdriver.SurfacerConf)
		err := b.DecodeMessage(msg)
		m.Surfacer = &SurfacerDef_StackdriverSurfacer{msg}
		return true, err
	case 12: // surfacer.file_surfacer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_surfacer_file.SurfacerConf)
		err := b.DecodeMessage(msg)
		m.Surfacer = &SurfacerDef_FileSurfacer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SurfacerDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SurfacerDef)
	// surfacer
	switch x := m.Surfacer.(type) {
	case *SurfacerDef_PrometheusSurfacer:
		s := proto.Size(x.PrometheusSurfacer)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SurfacerDef_StackdriverSurfacer:
		s := proto.Size(x.StackdriverSurfacer)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SurfacerDef_FileSurfacer:
		s := proto.Size(x.FileSurfacer)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*SurfacerDef)(nil), "cloudprober.surfacer.SurfacerDef")
	proto.RegisterEnum("cloudprober.surfacer.Type", Type_name, Type_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/surfacers/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x6f, 0xb2, 0x40,
	0x18, 0xc4, 0xc5, 0x97, 0x83, 0xef, 0x83, 0xb5, 0x64, 0xf5, 0x60, 0x3c, 0x99, 0x5e, 0x6a, 0x9a,
	0x74, 0x49, 0x4c, 0x2f, 0x3d, 0xf5, 0x8f, 0xac, 0xd1, 0xb4, 0x62, 0xb3, 0x68, 0xaf, 0x56, 0xd7,
	0x05, 0x49, 0xd5, 0x25, 0x2b, 0x34, 0xf1, 0x7b, 0xf7, 0x03, 0x34, 0x90, 0x6e, 0x41, 0xc3, 0x81,
	0x1b, 0x79, 0x66, 0xe6, 0x37, 0xec, 0xc0, 0x9d, 0x1f, 0x44, 0x9b, 0x78, 0x85, 0x99, 0xd8, 0x59,
	0xbe, 0x10, 0xfe, 0x96, 0x5b, 0x6c, 0x2b, 0xe2, 0x75, 0x28, 0xc5, 0x8a, 0x4b, 0xeb, 0x10, 0x4b,
	0x6f, 0xc9, 0xb8, 0x3c, 0x58, 0x4c, 0xec, 0xbd, 0xc0, 0xc7, 0xa1, 0x14, 0x91, 0x40, 0xad, 0x9c,
	0x07, 0x2b, 0x4f, 0xe7, 0xa1, 0x2c, 0x2b, 0x94, 0x62, 0xc7, 0xa3, 0x0d, 0x8f, 0x4f, 0xb1, 0x9d,
	0xc7, 0xb2, 0x80, 0x43, 0xb4, 0x64, 0x9f, 0x6b, 0x19, 0x7c, 0x71, 0x79, 0x4a, 0xb8, 0x2f, 0x4b,
	0xf0, 0x82, 0x44, 0xca, 0x45, 0xaf, 0xbe, 0xab, 0x60, 0xb8, 0xbf, 0xba, 0xcd, 0x3d, 0x84, 0x40,
	0xdf, 0x2f, 0x77, 0xbc, 0xad, 0x75, 0xb5, 0xde, 0x7f, 0x9a, 0x7e, 0x23, 0x0c, 0x7a, 0x74, 0x0c,
	0x79, 0xbb, 0xda, 0xd5, 0x7a, 0x8d, 0x7e, 0x07, 0x17, 0xcd, 0x80, 0x67, 0xc7, 0x90, 0xd3, 0xd4,
	0x87, 0x3e, 0xa0, 0x99, 0xbd, 0x75, 0xa1, 0x1c, 0x6d, 0xe8, 0x6a, 0x3d, 0xa3, 0x7f, 0x5b, 0x1c,
	0xcf, 0x02, 0x58, 0xfd, 0xce, 0x40, 0xec, 0xbd, 0x51, 0x85, 0xa2, 0x4c, 0x52, 0x0a, 0x62, 0xd0,
	0xca, 0x8d, 0x91, 0x55, 0x18, 0x69, 0x05, 0x2e, 0xae, 0xc8, 0x25, 0xce, 0x3b, 0x9a, 0x39, 0xed,
	0xaf, 0xc4, 0x81, 0x8b, 0x64, 0xaf, 0x8c, 0x5e, 0x4f, 0xe9, 0xd7, 0xc5, 0xf4, 0xc4, 0x7a, 0x8e,
	0xad, 0x27, 0x47, 0x75, 0x7b, 0x06, 0xa8, 0x29, 0xf7, 0xcd, 0x04, 0xf4, 0x64, 0x30, 0x54, 0x03,
	0xdd, 0x99, 0x3a, 0xc4, 0xac, 0xa0, 0x06, 0xc0, 0x1b, 0x9d, 0x4e, 0xc8, 0x6c, 0x44, 0xe6, 0xae,
	0xa9, 0xa1, 0x4b, 0x30, 0xdc, 0xd9, 0xd3, 0xe0, 0xc5, 0xa6, 0xe3, 0x77, 0x42, 0xcd, 0x6a, 0x62,
	0x1d, 0x8e, 0x5f, 0x89, 0xf9, 0x0f, 0x99, 0x50, 0x9f, 0xbb, 0x84, 0x2e, 0x6c, 0x32, 0x1c, 0x3b,
	0xc4, 0x36, 0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb4, 0x7c, 0x99, 0xd0, 0x02, 0x00,
	0x00,
}