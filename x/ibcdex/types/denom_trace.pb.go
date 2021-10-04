// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/denom_trace.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DenomTrace struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Port    string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	Channel string `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	Origin  string `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (m *DenomTrace) Reset()         { *m = DenomTrace{} }
func (m *DenomTrace) String() string { return proto.CompactTextString(m) }
func (*DenomTrace) ProtoMessage()    {}
func (*DenomTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_d98bc4572f53ce26, []int{0}
}
func (m *DenomTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomTrace.Merge(m, src)
}
func (m *DenomTrace) XXX_Size() int {
	return m.Size()
}
func (m *DenomTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomTrace.DiscardUnknown(m)
}

var xxx_messageInfo_DenomTrace proto.InternalMessageInfo

func (m *DenomTrace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DenomTrace) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *DenomTrace) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *DenomTrace) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *DenomTrace) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func init() {
	proto.RegisterType((*DenomTrace)(nil), "code0xff.interchange.ibcdex.DenomTrace")
}

func init() { proto.RegisterFile("ibcdex/denom_trace.proto", fileDescriptor_d98bc4572f53ce26) }

var fileDescriptor_d98bc4572f53ce26 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x37, 0x7a, 0x77, 0x62, 0xca, 0x70, 0x48, 0x50, 0x08, 0x62, 0x65, 0x95, 0x08, 0xbe,
	0x81, 0x58, 0x58, 0x8b, 0x95, 0x8d, 0x64, 0x93, 0xb9, 0x6c, 0xc0, 0xcb, 0x2c, 0x31, 0x42, 0x6c,
	0x7d, 0x02, 0x1f, 0xcb, 0xf2, 0x4a, 0x4b, 0xd9, 0x7d, 0x11, 0xd9, 0x64, 0x17, 0xec, 0xe6, 0x9b,
	0xfc, 0xf9, 0xe0, 0x1f, 0xca, 0x7d, 0x6b, 0x2c, 0x64, 0x65, 0x21, 0xe0, 0xfe, 0x25, 0x45, 0x6d,
	0x40, 0xf6, 0x11, 0x13, 0xb2, 0x0b, 0x83, 0x16, 0x6e, 0xf2, 0x6e, 0x27, 0x7d, 0x48, 0x10, 0x4d,
	0xa7, 0x83, 0x03, 0x59, 0xe3, 0xe7, 0x5b, 0x87, 0x0e, 0x4b, 0x4e, 0x4d, 0x53, 0xfd, 0x72, 0xf5,
	0x49, 0x28, 0xbd, 0x9f, 0x44, 0x4f, 0x93, 0x87, 0x71, 0x7a, 0x62, 0x22, 0xe8, 0x84, 0x91, 0x93,
	0x4b, 0x72, 0x7d, 0xfa, 0xb8, 0x20, 0xdb, 0xd2, 0xb5, 0x0f, 0x16, 0x32, 0x3f, 0x2a, 0xfb, 0x0a,
	0x8c, 0xd1, 0x55, 0x8f, 0x31, 0xf1, 0xe3, 0xb2, 0x2c, 0x73, 0x71, 0x74, 0x3a, 0x04, 0x78, 0xe5,
	0xab, 0xd9, 0x51, 0x91, 0x9d, 0xd1, 0x0d, 0x46, 0xef, 0x7c, 0xe0, 0xeb, 0xf2, 0x30, 0xd3, 0xdd,
	0xc3, 0xf7, 0x20, 0xc8, 0x61, 0x10, 0xe4, 0x77, 0x10, 0xe4, 0x6b, 0x14, 0xcd, 0x61, 0x14, 0xcd,
	0xcf, 0x28, 0x9a, 0x67, 0xe9, 0x7c, 0xea, 0xde, 0x5b, 0x69, 0x70, 0xaf, 0x96, 0x72, 0xea, 0x5f,
	0x39, 0x95, 0xd5, 0x7c, 0x8d, 0xf4, 0xd1, 0xc3, 0x5b, 0xbb, 0x29, 0xad, 0x6e, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x25, 0x60, 0x79, 0x86, 0x24, 0x01, 0x00, 0x00,
}

func (m *DenomTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDenomTrace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenomTrace(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenomTrace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovDenomTrace(uint64(l))
	}
	return n
}

func sovDenomTrace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenomTrace(x uint64) (n int) {
	return sovDenomTrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenomTrace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DenomTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDenomTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenomTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenomTrace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDenomTrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenomTrace
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
					return 0, ErrIntOverflowDenomTrace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDenomTrace
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
			if length < 0 {
				return 0, ErrInvalidLengthDenomTrace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenomTrace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenomTrace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenomTrace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenomTrace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenomTrace = fmt.Errorf("proto: unexpected end of group")
)
