// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/fault/v2/fault.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _type "github.com/envoyproxy/go-control-plane/envoy/type"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import time "time"
import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	// Fixed delay (step function).
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}
var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}
func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fault_a089bb335ffda0ce, []int{0, 0}
}

// Delay specification is used to inject latency into the
// HTTP/gRPC/Mongo/Redis operation or delay proxying of TCP connections.
type FaultDelay struct {
	// Delay type to use (fixed|exponential|..). Currently, only fixed delay (step function) is
	// supported.
	Type FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.config.filter.fault.v2.FaultDelay_FaultDelayType" json:"type,omitempty"`
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	FaultDelaySecifier isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	// The percentage of operations/connection requests on which the delay will be injected.
	Percentage           *_type.FractionalPercent `protobuf:"bytes,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultDelay) Reset()         { *m = FaultDelay{} }
func (m *FaultDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()    {}
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_fault_a089bb335ffda0ce, []int{0}
}
func (m *FaultDelay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultDelay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FaultDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay.Merge(dst, src)
}
func (m *FaultDelay) XXX_Size() int {
	return m.Size()
}
func (m *FaultDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay proto.InternalMessageInfo

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultDelay_FixedDelay struct {
	FixedDelay *time.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof,stdduration"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.Type
	}
	return FaultDelay_FIXED
}

func (m *FaultDelay) GetFixedDelay() *time.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultDelay_OneofMarshaler, _FaultDelay_OneofUnmarshaler, _FaultDelay_OneofSizer, []interface{}{
		(*FaultDelay_FixedDelay)(nil),
	}
}

func _FaultDelay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		dAtA, err := github_com_gogo_protobuf_types.StdDurationMarshal(*x.FixedDelay)
		if err != nil {
			return err
		}
		if err := b.EncodeRawBytes(dAtA); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FaultDelay.FaultDelaySecifier has unexpected type %T", x)
	}
	return nil
}

func _FaultDelay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultDelay)
	switch tag {
	case 3: // fault_delay_secifier.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		if err != nil {
			return true, err
		}
		c := new(time.Duration)
		if err2 := github_com_gogo_protobuf_types.StdDurationUnmarshal(c, x); err2 != nil {
			return true, err
		}
		m.FaultDelaySecifier = &FaultDelay_FixedDelay{c}
		return true, err
	default:
		return false, nil
	}
}

func _FaultDelay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		s := github_com_gogo_protobuf_types.SizeOfStdDuration(*x.FixedDelay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FaultDelay)(nil), "envoy.config.filter.fault.v2.FaultDelay")
	proto.RegisterEnum("envoy.config.filter.fault.v2.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
}
func (m *FaultDelay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultDelay) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Type))
	}
	if m.FaultDelaySecifier != nil {
		nn1, err := m.FaultDelaySecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Percentage != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Percentage.Size()))
		n2, err := m.Percentage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FaultDelay_FixedDelay) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.FixedDelay != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFault(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.FixedDelay)))
		n3, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.FixedDelay, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FaultDelay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovFault(uint64(m.Type))
	}
	if m.FaultDelaySecifier != nil {
		n += m.FaultDelaySecifier.Size()
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultDelay_FixedDelay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedDelay != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.FixedDelay)
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}

func sovFault(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultDelay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FaultDelay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultDelay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (FaultDelay_FaultDelayType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := new(time.Duration)
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(v, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.FaultDelaySecifier = &FaultDelay_FixedDelay{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Percentage == nil {
				m.Percentage = &_type.FractionalPercent{}
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFault(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/fault/v2/fault.proto", fileDescriptor_fault_a089bb335ffda0ce)
}

var fileDescriptor_fault_a089bb335ffda0ce = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xbf, 0x4a, 0xfb, 0x40,
	0x1c, 0xef, 0x25, 0xed, 0x8f, 0x9f, 0x57, 0x28, 0x25, 0x14, 0x8c, 0xd5, 0xc6, 0xd2, 0xa9, 0x74,
	0xb8, 0x83, 0x38, 0x38, 0xe9, 0x10, 0x6a, 0x51, 0x71, 0x28, 0x45, 0x50, 0x5c, 0xca, 0x35, 0xf9,
	0x26, 0x1c, 0x84, 0x5c, 0x88, 0x69, 0x68, 0x56, 0x9f, 0xc2, 0x67, 0x70, 0x76, 0x10, 0xa7, 0x8e,
	0x8e, 0xbe, 0x81, 0xd2, 0xad, 0x6f, 0x21, 0xb9, 0x4b, 0xb1, 0x2e, 0xdd, 0x3e, 0xc9, 0x7d, 0xfe,
	0x7e, 0x71, 0x1f, 0xa2, 0x4c, 0xe4, 0xd4, 0x15, 0x91, 0xcf, 0x03, 0xea, 0xf3, 0x30, 0x85, 0x84,
	0xfa, 0x6c, 0x1e, 0xa6, 0x34, 0xb3, 0x15, 0x20, 0x71, 0x22, 0x52, 0x61, 0x1c, 0x49, 0x26, 0x51,
	0x4c, 0xa2, 0x98, 0x44, 0x11, 0x32, 0xbb, 0x6d, 0x2a, 0x9f, 0x34, 0x8f, 0x81, 0xc6, 0x90, 0xb8,
	0x10, 0x95, 0xba, 0xb6, 0x15, 0x08, 0x11, 0x84, 0x40, 0xe5, 0xd7, 0x6c, 0xee, 0x53, 0x6f, 0x9e,
	0xb0, 0x94, 0x8b, 0xa8, 0x7c, 0xdf, 0xcf, 0x58, 0xc8, 0x3d, 0x96, 0x02, 0xdd, 0x80, 0xf2, 0xa1,
	0x15, 0x88, 0x40, 0x48, 0x48, 0x0b, 0xa4, 0xfe, 0xf6, 0x5e, 0x35, 0x8c, 0x47, 0x45, 0xea, 0x10,
	0x42, 0x96, 0x1b, 0x77, 0xb8, 0x5a, 0x64, 0x9a, 0xa8, 0x8b, 0xfa, 0x0d, 0xfb, 0x94, 0xec, 0x2a,
	0x49, 0x7e, 0x75, 0x5b, 0xf0, 0x36, 0x8f, 0xc1, 0xc1, 0xef, 0xeb, 0xa5, 0x5e, 0x7b, 0x42, 0x5a,
	0x13, 0x4d, 0xa4, 0xa1, 0x71, 0x83, 0xeb, 0x3e, 0x5f, 0x80, 0x37, 0xf5, 0x0a, 0x92, 0xa9, 0x77,
	0x51, 0xbf, 0x6e, 0x1f, 0x10, 0x35, 0x86, 0x6c, 0xc6, 0x90, 0x61, 0x39, 0xc6, 0x69, 0x3c, 0x7f,
	0x1d, 0x23, 0xe9, 0xf2, 0x82, 0xb4, 0x41, 0xe5, 0xb2, 0x32, 0xc1, 0x52, 0xaf, 0x6a, 0x9e, 0x61,
	0x5c, 0x5e, 0x85, 0x05, 0x60, 0x56, 0xa5, 0x59, 0xa7, 0x2c, 0x5b, 0xc4, 0x91, 0x51, 0xc2, 0xdc,
	0xc2, 0x87, 0x85, 0x63, 0xc5, 0x9b, 0x6c, 0x09, 0x7a, 0x87, 0xb8, 0xf1, 0xb7, 0xb0, 0xb1, 0x87,
	0x6b, 0xa3, 0xab, 0xfb, 0x8b, 0x61, 0xb3, 0xe2, 0x74, 0x70, 0x4b, 0x2e, 0x54, 0x4d, 0xa7, 0x8f,
	0xe0, 0x72, 0x9f, 0x43, 0x62, 0xd4, 0xde, 0xd6, 0x4b, 0x1d, 0x5d, 0x57, 0xff, 0x6b, 0x4d, 0xdd,
	0x39, 0xff, 0x58, 0x59, 0xe8, 0x73, 0x65, 0xa1, 0xef, 0x95, 0x85, 0xf0, 0x80, 0x0b, 0x15, 0x1e,
	0x27, 0x62, 0x91, 0xef, 0x3c, 0xda, 0x18, 0x3d, 0x68, 0x99, 0x3d, 0xfb, 0x27, 0x17, 0x9f, 0xfc,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x22, 0x71, 0xdd, 0x30, 0x02, 0x00, 0x00,
}
