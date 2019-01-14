// Code generated by protoc-gen-gogo.
// source: definitions.proto
// DO NOT EDIT!

package dialog

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/golang/protobuf/protoc-gen-go/descriptor"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UUIDValue struct {
	Msb int64 `protobuf:"varint,1,opt,name=msb,proto3" json:"msb,omitempty"`
	Lsb int64 `protobuf:"varint,2,opt,name=lsb,proto3" json:"lsb,omitempty"`
}

func (m *UUIDValue) Reset()                    { *m = UUIDValue{} }
func (*UUIDValue) ProtoMessage()               {}
func (*UUIDValue) Descriptor() ([]byte, []int) { return fileDescriptorDefinitions, []int{0} }

func (m *UUIDValue) GetMsb() int64 {
	if m != nil {
		return m.Msb
	}
	return 0
}

func (m *UUIDValue) GetLsb() int64 {
	if m != nil {
		return m.Lsb
	}
	return 0
}

type DialogOptions struct {
	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *DialogOptions) Reset()                    { *m = DialogOptions{} }
func (*DialogOptions) ProtoMessage()               {}
func (*DialogOptions) Descriptor() ([]byte, []int) { return fileDescriptorDefinitions, []int{1} }

func (m *DialogOptions) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

// / server timestamp when state was created or modified
type DataClock struct {
	Clock int64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *DataClock) Reset()                    { *m = DataClock{} }
func (*DataClock) ProtoMessage()               {}
func (*DataClock) Descriptor() ([]byte, []int) { return fileDescriptorDefinitions, []int{2} }

func (m *DataClock) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

var E_Dlg = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf1.FieldOptions)(nil),
	ExtensionType: (*DialogOptions)(nil),
	Field:         100001,
	Name:          "dialog.dlg",
	Tag:           "bytes,100001,opt,name=dlg",
	Filename:      "definitions.proto",
}

func init() {
	proto.RegisterType((*UUIDValue)(nil), "dialog.UUIDValue")
	proto.RegisterType((*DialogOptions)(nil), "dialog.DialogOptions")
	proto.RegisterType((*DataClock)(nil), "dialog.DataClock")
	proto.RegisterExtension(E_Dlg)
}
func (this *UUIDValue) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UUIDValue)
	if !ok {
		that2, ok := that.(UUIDValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Msb != that1.Msb {
		return false
	}
	if this.Lsb != that1.Lsb {
		return false
	}
	return true
}
func (this *DialogOptions) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DialogOptions)
	if !ok {
		that2, ok := that.(DialogOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	return true
}
func (this *DataClock) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DataClock)
	if !ok {
		that2, ok := that.(DataClock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Clock != that1.Clock {
		return false
	}
	return true
}
func (this *UUIDValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dialog.UUIDValue{")
	s = append(s, "Msb: "+fmt.Sprintf("%#v", this.Msb)+",\n")
	s = append(s, "Lsb: "+fmt.Sprintf("%#v", this.Lsb)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DialogOptions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dialog.DialogOptions{")
	s = append(s, "Log: "+fmt.Sprintf("%#v", this.Log)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DataClock) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dialog.DataClock{")
	s = append(s, "Clock: "+fmt.Sprintf("%#v", this.Clock)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDefinitions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UUIDValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UUIDValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Msb != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDefinitions(dAtA, i, uint64(m.Msb))
	}
	if m.Lsb != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDefinitions(dAtA, i, uint64(m.Lsb))
	}
	return i, nil
}

func (m *DialogOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DialogOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDefinitions(dAtA, i, uint64(len(m.Log)))
		i += copy(dAtA[i:], m.Log)
	}
	return i, nil
}

func (m *DataClock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataClock) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Clock != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDefinitions(dAtA, i, uint64(m.Clock))
	}
	return i, nil
}

func encodeFixed64Definitions(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Definitions(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDefinitions(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UUIDValue) Size() (n int) {
	var l int
	_ = l
	if m.Msb != 0 {
		n += 1 + sovDefinitions(uint64(m.Msb))
	}
	if m.Lsb != 0 {
		n += 1 + sovDefinitions(uint64(m.Lsb))
	}
	return n
}

func (m *DialogOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovDefinitions(uint64(l))
	}
	return n
}

func (m *DataClock) Size() (n int) {
	var l int
	_ = l
	if m.Clock != 0 {
		n += 1 + sovDefinitions(uint64(m.Clock))
	}
	return n
}

func sovDefinitions(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDefinitions(x uint64) (n int) {
	return sovDefinitions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UUIDValue) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UUIDValue{`,
		`Msb:` + fmt.Sprintf("%v", this.Msb) + `,`,
		`Lsb:` + fmt.Sprintf("%v", this.Lsb) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DialogOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DialogOptions{`,
		`Log:` + fmt.Sprintf("%v", this.Log) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DataClock) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DataClock{`,
		`Clock:` + fmt.Sprintf("%v", this.Clock) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDefinitions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UUIDValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefinitions
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
			return fmt.Errorf("proto: UUIDValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UUIDValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msb", wireType)
			}
			m.Msb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Msb |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lsb", wireType)
			}
			m.Lsb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lsb |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDefinitions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDefinitions
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
func (m *DialogOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefinitions
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
			return fmt.Errorf("proto: DialogOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DialogOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDefinitions
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDefinitions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDefinitions
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
func (m *DataClock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefinitions
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
			return fmt.Errorf("proto: DataClock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataClock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			m.Clock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefinitions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Clock |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDefinitions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDefinitions
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
func skipDefinitions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDefinitions
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
					return 0, ErrIntOverflowDefinitions
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
					return 0, ErrIntOverflowDefinitions
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
				return 0, ErrInvalidLengthDefinitions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDefinitions
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
				next, err := skipDefinitions(dAtA[start:])
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
	ErrInvalidLengthDefinitions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDefinitions   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("definitions.proto", fileDescriptorDefinitions) }

var fileDescriptorDefinitions = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xe3, 0x5b, 0xdd, 0x4a, 0x35, 0x42, 0x82, 0xa8, 0x45, 0x51, 0x25, 0xac, 0x34, 0x53,
	0x27, 0x47, 0x82, 0x8d, 0x05, 0x09, 0x02, 0x82, 0x09, 0x29, 0x52, 0xd9, 0xfd, 0xaf, 0x96, 0x85,
	0x5b, 0x47, 0x71, 0xca, 0xcc, 0x0b, 0x20, 0xf1, 0x0a, 0x6c, 0x3c, 0x0a, 0x63, 0x47, 0x46, 0x62,
	0x16, 0xc6, 0x3e, 0x02, 0x8a, 0x1d, 0x06, 0x26, 0x9f, 0xef, 0xf3, 0xf1, 0x39, 0x3f, 0xc3, 0x43,
	0x2e, 0x96, 0x6a, 0xad, 0x1a, 0x65, 0xd6, 0x16, 0x57, 0xb5, 0x69, 0x4c, 0x3c, 0xe4, 0x8a, 0x68,
	0x23, 0xa7, 0x13, 0xcb, 0x88, 0x26, 0x15, 0xcd, 0xfb, 0x33, 0x5c, 0x4f, 0x53, 0x69, 0x8c, 0xd4,
	0x22, 0xf7, 0x13, 0xdd, 0x2c, 0x73, 0x2e, 0x2c, 0xab, 0x55, 0xd5, 0x98, 0x3a, 0x38, 0xb2, 0x1c,
	0x8e, 0x16, 0x8b, 0xdb, 0xe2, 0x9e, 0xe8, 0x8d, 0x88, 0x0f, 0xe0, 0x60, 0x65, 0x69, 0x02, 0x52,
	0x30, 0x1f, 0x94, 0x9d, 0xec, 0x36, 0xda, 0xd2, 0xe4, 0x5f, 0xd8, 0x68, 0x4b, 0xb3, 0x19, 0xdc,
	0x2f, 0x7c, 0xe7, 0x5d, 0xe5, 0x41, 0xbc, 0xc5, 0x48, 0xff, 0x68, 0x54, 0x76, 0x32, 0x9b, 0xc1,
	0x51, 0x41, 0x1a, 0x72, 0xa9, 0x0d, 0x7b, 0x88, 0xc7, 0xf0, 0x3f, 0xeb, 0x44, 0x9f, 0x1a, 0x86,
	0xb3, 0x1b, 0x38, 0xe0, 0x5a, 0xc6, 0xc7, 0x38, 0x00, 0xe2, 0x5f, 0x40, 0x7c, 0xad, 0x84, 0xe6,
	0x7d, 0x74, 0xf2, 0xfa, 0x3c, 0x4c, 0xc1, 0x7c, 0xef, 0x64, 0x82, 0xc3, 0x37, 0xf1, 0x9f, 0xe6,
	0xb2, 0x8b, 0xb8, 0xb8, 0x72, 0xe7, 0x47, 0x70, 0xac, 0x56, 0x98, 0x6b, 0x89, 0x65, 0x5d, 0x31,
	0x6c, 0x45, 0xfd, 0xa8, 0x98, 0xb0, 0xdb, 0x16, 0x45, 0x1f, 0x2d, 0x8a, 0x76, 0x2d, 0x02, 0x4f,
	0x0e, 0x81, 0x37, 0x87, 0xc0, 0xbb, 0x43, 0x60, 0xeb, 0x10, 0xf8, 0x74, 0x08, 0x7c, 0x3b, 0x14,
	0xed, 0x1c, 0x02, 0x2f, 0x5f, 0x28, 0xa2, 0x43, 0x4f, 0x70, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0xf3, 0xd0, 0xe1, 0x64, 0x01, 0x00, 0x00,
}
