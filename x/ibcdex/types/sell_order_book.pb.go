// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/sell_order_book.proto

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

type SellOrderBook struct {
	Creator     string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index       string     `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	AmountDenom string     `protobuf:"bytes,3,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom  string     `protobuf:"bytes,4,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Book        *OrderBook `protobuf:"bytes,5,opt,name=book,proto3" json:"book,omitempty"`
}

func (m *SellOrderBook) Reset()         { *m = SellOrderBook{} }
func (m *SellOrderBook) String() string { return proto.CompactTextString(m) }
func (*SellOrderBook) ProtoMessage()    {}
func (*SellOrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cfec8ad1959141, []int{0}
}
func (m *SellOrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOrderBook.Merge(m, src)
}
func (m *SellOrderBook) XXX_Size() int {
	return m.Size()
}
func (m *SellOrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_SellOrderBook proto.InternalMessageInfo

func (m *SellOrderBook) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SellOrderBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SellOrderBook) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *SellOrderBook) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *SellOrderBook) GetBook() *OrderBook {
	if m != nil {
		return m.Book
	}
	return nil
}

func init() {
	proto.RegisterType((*SellOrderBook)(nil), "code0xff.interchange.ibcdex.SellOrderBook")
}

func init() { proto.RegisterFile("ibcdex/sell_order_book.proto", fileDescriptor_e1cfec8ad1959141) }

var fileDescriptor_e1cfec8ad1959141 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x6b, 0x68, 0x41, 0xb8, 0x62, 0xb1, 0x3a, 0x44, 0x05, 0x59, 0x11, 0x03, 0xea, 0x64,
	0x23, 0xd8, 0x18, 0x2b, 0x06, 0x36, 0xa4, 0xb2, 0xb1, 0x54, 0x89, 0x73, 0x4d, 0xad, 0x26, 0xbe,
	0xc8, 0x75, 0xa5, 0xf0, 0x2f, 0xf8, 0x43, 0xec, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x47, 0x50, 0xec,
	0x16, 0x65, 0x62, 0xf3, 0x7d, 0x7e, 0xef, 0xe9, 0xee, 0xd1, 0x6b, 0x9d, 0xaa, 0x0c, 0x6a, 0xb9,
	0x85, 0xa2, 0x58, 0xa2, 0xcd, 0xc0, 0x2e, 0x53, 0xc4, 0x8d, 0xa8, 0x2c, 0x3a, 0x64, 0x57, 0x0a,
	0x33, 0xb8, 0xab, 0x57, 0x2b, 0xa1, 0x8d, 0x03, 0xab, 0xd6, 0x89, 0xc9, 0x41, 0x04, 0xcb, 0x74,
	0x92, 0x63, 0x8e, 0x5e, 0x27, 0xbb, 0x57, 0xb0, 0x4c, 0xd9, 0x21, 0xd0, 0x67, 0x05, 0x76, 0xf3,
	0x49, 0xe8, 0xe5, 0x2b, 0x14, 0xc5, 0x4b, 0xc7, 0xe6, 0x88, 0x1b, 0x16, 0xd1, 0x73, 0x65, 0x21,
	0x71, 0x68, 0x23, 0x12, 0x93, 0xd9, 0xc5, 0xe2, 0x38, 0xb2, 0x09, 0x1d, 0x69, 0x93, 0x41, 0x1d,
	0x9d, 0x78, 0x1e, 0x06, 0x16, 0xd3, 0x71, 0x52, 0xe2, 0xce, 0xb8, 0x27, 0x30, 0x58, 0x46, 0xa7,
	0xfe, 0xaf, 0x8f, 0x18, 0xa7, 0xb4, 0xb2, 0x5a, 0x41, 0x10, 0x0c, 0xbd, 0xa0, 0x47, 0xd8, 0x23,
	0x1d, 0x76, 0x87, 0x45, 0xa3, 0x98, 0xcc, 0xc6, 0xf7, 0xb7, 0xe2, 0x9f, 0xcb, 0xc4, 0xdf, 0x9e,
	0x0b, 0xef, 0x99, 0x3f, 0x7f, 0x35, 0x9c, 0xec, 0x1b, 0x4e, 0x7e, 0x1a, 0x4e, 0x3e, 0x5a, 0x3e,
	0xd8, 0xb7, 0x7c, 0xf0, 0xdd, 0xf2, 0xc1, 0x9b, 0xc8, 0xb5, 0x5b, 0xef, 0x52, 0xa1, 0xb0, 0x94,
	0xc7, 0x44, 0xd9, 0x4b, 0x94, 0xb5, 0x3c, 0xf4, 0xe1, 0xde, 0x2b, 0xd8, 0xa6, 0x67, 0xbe, 0x90,
	0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x2d, 0x06, 0xd2, 0x77, 0x01, 0x00, 0x00,
}

func (m *SellOrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Book != nil {
		{
			size, err := m.Book.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSellOrderBook(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSellOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovSellOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SellOrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	if m.Book != nil {
		l = m.Book.Size()
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	return n
}

func sovSellOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSellOrderBook(x uint64) (n int) {
	return sovSellOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SellOrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSellOrderBook
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
			return fmt.Errorf("proto: SellOrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
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
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Book", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Book == nil {
				m.Book = &OrderBook{}
			}
			if err := m.Book.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSellOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSellOrderBook
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
func skipSellOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSellOrderBook
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
					return 0, ErrIntOverflowSellOrderBook
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
					return 0, ErrIntOverflowSellOrderBook
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
				return 0, ErrInvalidLengthSellOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSellOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSellOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSellOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSellOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSellOrderBook = fmt.Errorf("proto: unexpected end of group")
)
