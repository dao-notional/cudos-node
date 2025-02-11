// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: marketplace/collection.proto

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

// Collection listed for sale in the marketplace
type Collection struct {
	Id              uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DenomId         string    `protobuf:"bytes,2,opt,name=denomId,proto3" json:"denomId,omitempty"`
	MintRoyalties   []Royalty `protobuf:"bytes,3,rep,name=mintRoyalties,proto3" json:"mintRoyalties"`
	ResaleRoyalties []Royalty `protobuf:"bytes,4,rep,name=resaleRoyalties,proto3" json:"resaleRoyalties"`
	Verified        bool      `protobuf:"varint,5,opt,name=verified,proto3" json:"verified,omitempty"`
	Owner           string    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_2874de85607b5723, []int{0}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Collection) GetDenomId() string {
	if m != nil {
		return m.DenomId
	}
	return ""
}

func (m *Collection) GetMintRoyalties() []Royalty {
	if m != nil {
		return m.MintRoyalties
	}
	return nil
}

func (m *Collection) GetResaleRoyalties() []Royalty {
	if m != nil {
		return m.ResaleRoyalties
	}
	return nil
}

func (m *Collection) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Collection) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Collection)(nil), "cudoventures.cudosnode.marketplace.Collection")
}

func init() { proto.RegisterFile("marketplace/collection.proto", fileDescriptor_2874de85607b5723) }

var fileDescriptor_2874de85607b5723 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x9b, 0xee, 0x8f, 0x33, 0xa2, 0x42, 0xd8, 0xa1, 0x0e, 0xa9, 0x65, 0xa7, 0x82, 0x98,
	0x80, 0x82, 0x1f, 0x60, 0x3b, 0x79, 0x93, 0x1e, 0x14, 0xf4, 0xd4, 0x35, 0xaf, 0x33, 0xd8, 0xe6,
	0x1d, 0x69, 0x3a, 0xed, 0xb7, 0xf0, 0x13, 0xf8, 0x79, 0x76, 0xdc, 0xd1, 0x93, 0xc8, 0xf6, 0x45,
	0x64, 0xad, 0xce, 0xe9, 0x45, 0xf0, 0xf6, 0x3e, 0x81, 0xe7, 0xf7, 0xfe, 0x92, 0xd0, 0xc3, 0x2c,
	0x36, 0x0f, 0x60, 0x27, 0x69, 0x9c, 0x80, 0x48, 0x30, 0x4d, 0x21, 0xb1, 0x0a, 0x35, 0x9f, 0x18,
	0xb4, 0xc8, 0xfa, 0x49, 0x21, 0x71, 0x0a, 0xda, 0x16, 0x06, 0x72, 0xbe, 0x0a, 0xb9, 0x46, 0x09,
	0x7c, 0xa3, 0xd4, 0x3b, 0xd8, 0x24, 0x18, 0x2c, 0xe3, 0xd4, 0x96, 0x75, 0xbd, 0xd7, 0x1d, 0xe3,
	0x18, 0xab, 0x51, 0xac, 0xa6, 0xfa, 0xb4, 0xff, 0xe2, 0x52, 0x3a, 0x5c, 0x6f, 0x62, 0x7b, 0xd4,
	0x55, 0xd2, 0x23, 0x01, 0x09, 0x9b, 0x91, 0xab, 0x24, 0xf3, 0xe8, 0x96, 0x04, 0x8d, 0xd9, 0x85,
	0xf4, 0xdc, 0x80, 0x84, 0xdb, 0xd1, 0x57, 0x64, 0xd7, 0x74, 0x37, 0x53, 0xda, 0x46, 0xd5, 0x0e,
	0x05, 0xb9, 0xd7, 0x08, 0x1a, 0xe1, 0xce, 0xe9, 0x31, 0xff, 0xdb, 0x92, 0xd7, 0xa5, 0x72, 0xd0,
	0x9c, 0xbd, 0x1d, 0x39, 0xd1, 0x4f, 0x0e, 0xbb, 0xa5, 0xfb, 0x06, 0xf2, 0x38, 0x85, 0x6f, 0x74,
	0xf3, 0xbf, 0xe8, 0xdf, 0x24, 0xd6, 0xa3, 0x9d, 0x29, 0x18, 0x75, 0xa7, 0x40, 0x7a, 0xad, 0x80,
	0x84, 0x9d, 0x68, 0x9d, 0x59, 0x97, 0xb6, 0xf0, 0x51, 0x83, 0xf1, 0xda, 0xd5, 0x4d, 0xeb, 0x30,
	0xb8, 0x9c, 0x2d, 0x7c, 0x32, 0x5f, 0xf8, 0xe4, 0x7d, 0xe1, 0x93, 0xe7, 0xa5, 0xef, 0xcc, 0x97,
	0xbe, 0xf3, 0xba, 0xf4, 0x9d, 0x9b, 0xf3, 0xb1, 0xb2, 0xf7, 0xc5, 0x88, 0x27, 0x98, 0x89, 0x61,
	0x21, 0xf1, 0xea, 0xd3, 0x4c, 0x54, 0x66, 0x27, 0x2b, 0x35, 0xf1, 0x24, 0x36, 0x3f, 0xc4, 0x96,
	0x13, 0xc8, 0x47, 0xed, 0xea, 0xe5, 0xcf, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x44, 0x8b,
	0x2f, 0xee, 0x01, 0x00, 0x00,
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x32
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.ResaleRoyalties) > 0 {
		for iNdEx := len(m.ResaleRoyalties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResaleRoyalties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MintRoyalties) > 0 {
		for iNdEx := len(m.MintRoyalties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintRoyalties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DenomId) > 0 {
		i -= len(m.DenomId)
		copy(dAtA[i:], m.DenomId)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.DenomId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCollection(uint64(m.Id))
	}
	l = len(m.DenomId)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.MintRoyalties) > 0 {
		for _, e := range m.MintRoyalties {
			l = e.Size()
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	if len(m.ResaleRoyalties) > 0 {
		for _, e := range m.ResaleRoyalties {
			l = e.Size()
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	if m.Verified {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	return n
}

func sovCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollection(x uint64) (n int) {
	return sovCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
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
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRoyalties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRoyalties = append(m.MintRoyalties, Royalty{})
			if err := m.MintRoyalties[len(m.MintRoyalties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResaleRoyalties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResaleRoyalties = append(m.ResaleRoyalties, Royalty{})
			if err := m.ResaleRoyalties[len(m.ResaleRoyalties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCollection
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
func skipCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
				return 0, ErrInvalidLengthCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollection = fmt.Errorf("proto: unexpected end of group")
)
