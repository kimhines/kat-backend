// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/range.proto

package envoy_type

import (
	bytes "bytes"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the int64 start and end of the range using half-open interval semantics [start,
// end).
type Int64Range struct {
	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Range) Reset()         { *m = Int64Range{} }
func (m *Int64Range) String() string { return proto.CompactTextString(m) }
func (*Int64Range) ProtoMessage()    {}
func (*Int64Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d5354588716e6e7, []int{0}
}
func (m *Int64Range) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Int64Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Int64Range.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Int64Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Range.Merge(m, src)
}
func (m *Int64Range) XXX_Size() int {
	return m.Size()
}
func (m *Int64Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Range.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Range proto.InternalMessageInfo

func (m *Int64Range) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Int64Range) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End                  float64  `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRange) Reset()         { *m = DoubleRange{} }
func (m *DoubleRange) String() string { return proto.CompactTextString(m) }
func (*DoubleRange) ProtoMessage()    {}
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d5354588716e6e7, []int{1}
}
func (m *DoubleRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DoubleRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DoubleRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DoubleRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRange.Merge(m, src)
}
func (m *DoubleRange) XXX_Size() int {
	return m.Size()
}
func (m *DoubleRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRange.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRange proto.InternalMessageInfo

func (m *DoubleRange) GetStart() float64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DoubleRange) GetEnd() float64 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*Int64Range)(nil), "envoy.type.Int64Range")
	proto.RegisterType((*DoubleRange)(nil), "envoy.type.DoubleRange")
}

func init() { proto.RegisterFile("envoy/type/range.proto", fileDescriptor_6d5354588716e6e7) }

var fileDescriptor_6d5354588716e6e7 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x4a, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x8b, 0xeb, 0x81, 0xc4, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1,
	0xc2, 0xfa, 0x20, 0x16, 0x44, 0x85, 0x92, 0x09, 0x17, 0x97, 0x67, 0x5e, 0x89, 0x99, 0x49, 0x10,
	0x48, 0x97, 0x90, 0x08, 0x17, 0x6b, 0x71, 0x49, 0x62, 0x51, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x73, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x9a, 0x97, 0x22, 0xc1, 0x04, 0x16, 0x03, 0x31,
	0x95, 0x4c, 0xb9, 0xb8, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0xb1, 0x68, 0x63, 0xc4, 0xa2, 0x8d,
	0x11, 0xac, 0xcd, 0x49, 0x66, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x0a, 0xe2, 0xc0, 0x78, 0x90, 0x03, 0x93, 0xd8, 0xc0, 0x2e,
	0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xae, 0xd4, 0x26, 0x8b, 0xcd, 0x00, 0x00, 0x00,
}

func (this *Int64Range) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Int64Range)
	if !ok {
		that2, ok := that.(Int64Range)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DoubleRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DoubleRange)
	if !ok {
		that2, ok := that.(DoubleRange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Int64Range) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Int64Range) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRange(dAtA, i, uint64(m.Start))
	}
	if m.End != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRange(dAtA, i, uint64(m.End))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DoubleRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DoubleRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Start))))
		i += 8
	}
	if m.End != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.End))))
		i += 8
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRange(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Int64Range) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovRange(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovRange(uint64(m.End))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DoubleRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 9
	}
	if m.End != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRange(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRange(x uint64) (n int) {
	return sovRange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Int64Range) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRange
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
			return fmt.Errorf("proto: Int64Range: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Int64Range: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRange
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
func (m *DoubleRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRange
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
			return fmt.Errorf("proto: DoubleRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DoubleRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Start = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.End = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRange
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
func skipRange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRange
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
					return 0, ErrIntOverflowRange
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
					return 0, ErrIntOverflowRange
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
				return 0, ErrInvalidLengthRange
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRange
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRange
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
				next, err := skipRange(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRange
				}
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
	ErrInvalidLengthRange = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRange   = fmt.Errorf("proto: integer overflow")
)