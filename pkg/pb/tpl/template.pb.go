// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: template.proto

package tpl

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TemplateType int32

const (
	Template_Unknown TemplateType = 0
	Template_Sms     TemplateType = 1
	Template_WeChat  TemplateType = 2
	Template_Email   TemplateType = 3
)

var TemplateType_name = map[int32]string{
	0: "Template_Unknown",
	1: "Template_Sms",
	2: "Template_WeChat",
	3: "Template_Email",
}
var TemplateType_value = map[string]int32{
	"Template_Unknown": 0,
	"Template_Sms":     1,
	"Template_WeChat":  2,
	"Template_Email":   3,
}

func (x TemplateType) String() string {
	return proto.EnumName(TemplateType_name, int32(x))
}
func (TemplateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_template_4ceff9a5dc236d18, []int{0}
}

type TemplateAdder struct {
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// @inject_tag: json:"simple_id,omitempty"
	SimpleID             string   `protobuf:"bytes,2,opt,name=simpleID,proto3" json:"simple_id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateAdder) Reset()         { *m = TemplateAdder{} }
func (m *TemplateAdder) String() string { return proto.CompactTextString(m) }
func (*TemplateAdder) ProtoMessage()    {}
func (*TemplateAdder) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_4ceff9a5dc236d18, []int{0}
}
func (m *TemplateAdder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TemplateAdder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TemplateAdder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TemplateAdder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateAdder.Merge(dst, src)
}
func (m *TemplateAdder) XXX_Size() int {
	return m.Size()
}
func (m *TemplateAdder) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateAdder.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateAdder proto.InternalMessageInfo

func (m *TemplateAdder) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TemplateAdder) GetSimpleID() string {
	if m != nil {
		return m.SimpleID
	}
	return ""
}

func (m *TemplateAdder) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DBTemplate struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// @inject_tag: json:"simple_id,omitempty" db:"simple_id"
	SimpleID string `protobuf:"bytes,3,opt,name=simpleID,proto3" json:"simple_id,omitempty" db:"simple_id"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// @inject_tag: json:"created_at,omitempty" db:"created_at"
	CreatedAt string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: json:"updated_at,omitempty" db:"updated_at"
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBTemplate) Reset()         { *m = DBTemplate{} }
func (m *DBTemplate) String() string { return proto.CompactTextString(m) }
func (*DBTemplate) ProtoMessage()    {}
func (*DBTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_4ceff9a5dc236d18, []int{1}
}
func (m *DBTemplate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DBTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DBTemplate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DBTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBTemplate.Merge(dst, src)
}
func (m *DBTemplate) XXX_Size() int {
	return m.Size()
}
func (m *DBTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_DBTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_DBTemplate proto.InternalMessageInfo

func (m *DBTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DBTemplate) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *DBTemplate) GetSimpleID() string {
	if m != nil {
		return m.SimpleID
	}
	return ""
}

func (m *DBTemplate) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *DBTemplate) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DBTemplate) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*TemplateAdder)(nil), "tpl.TemplateAdder")
	proto.RegisterType((*DBTemplate)(nil), "tpl.DBTemplate")
	proto.RegisterEnum("tpl.TemplateType", TemplateType_name, TemplateType_value)
}
func (m *TemplateAdder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TemplateAdder) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(m.Type))
	}
	if len(m.SimpleID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.SimpleID)))
		i += copy(dAtA[i:], m.SimpleID)
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DBTemplate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DBTemplate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(m.Type))
	}
	if len(m.SimpleID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.SimpleID)))
		i += copy(dAtA[i:], m.SimpleID)
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if len(m.CreatedAt) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.CreatedAt)))
		i += copy(dAtA[i:], m.CreatedAt)
	}
	if len(m.UpdatedAt) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTemplate(dAtA, i, uint64(len(m.UpdatedAt)))
		i += copy(dAtA[i:], m.UpdatedAt)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTemplate(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TemplateAdder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTemplate(uint64(m.Type))
	}
	l = len(m.SimpleID)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DBTemplate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTemplate(uint64(m.Type))
	}
	l = len(m.SimpleID)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovTemplate(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTemplate(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplate(x uint64) (n int) {
	return sovTemplate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TemplateAdder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplate
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
			return fmt.Errorf("proto: TemplateAdder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TemplateAdder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SimpleID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SimpleID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplate
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
func (m *DBTemplate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplate
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
			return fmt.Errorf("proto: DBTemplate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DBTemplate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SimpleID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SimpleID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplate
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
				return ErrInvalidLengthTemplate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplate
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
func skipTemplate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplate
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
					return 0, ErrIntOverflowTemplate
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
					return 0, ErrIntOverflowTemplate
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
				return 0, ErrInvalidLengthTemplate
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplate
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
				next, err := skipTemplate(dAtA[start:])
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
	ErrInvalidLengthTemplate = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplate   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("template.proto", fileDescriptor_template_4ceff9a5dc236d18) }

var fileDescriptor_template_4ceff9a5dc236d18 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0xcd, 0x2d,
	0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x29, 0xc8, 0x91,
	0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xf3, 0xf5, 0x41, 0x2c, 0x88, 0x94, 0x52, 0x24, 0x17, 0x6f,
	0x08, 0x54, 0xb1, 0x63, 0x4a, 0x4a, 0x6a, 0x91, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x98, 0x2d, 0x24, 0xc5, 0xc5, 0x51, 0x9c, 0x99, 0x5b,
	0x90, 0x93, 0xea, 0xe9, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x0b, 0x49, 0x70,
	0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0x30, 0x83, 0xa5, 0x60, 0x5c, 0xa5, 0x45,
	0x8c, 0x5c, 0x5c, 0x2e, 0x4e, 0x30, 0xd3, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0xc0, 0xc6, 0x72,
	0x06, 0x31, 0x65, 0xa6, 0xc0, 0x2d, 0x62, 0xc2, 0x61, 0x11, 0x33, 0x6e, 0x8b, 0x58, 0x50, 0x2c,
	0x12, 0x92, 0xe1, 0xe2, 0x4c, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0x71, 0x2c, 0x91, 0x60, 0x05,
	0xcb, 0x21, 0x04, 0x40, 0xb2, 0xa5, 0x05, 0x29, 0x50, 0x59, 0x36, 0x88, 0x2c, 0x5c, 0x40, 0x2b,
	0x9e, 0x8b, 0x07, 0xe6, 0xc2, 0x10, 0x90, 0x0b, 0x44, 0xb8, 0x04, 0x60, 0xfc, 0xf8, 0xd0, 0xbc,
	0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x01, 0x06, 0x21, 0x01, 0x84, 0xaa, 0xf8, 0xe0, 0xdc, 0x62, 0x01,
	0x46, 0x21, 0x61, 0x2e, 0x7e, 0xb8, 0x48, 0x78, 0xaa, 0x73, 0x46, 0x62, 0x89, 0x00, 0x93, 0x90,
	0x10, 0x17, 0x1f, 0x5c, 0xd0, 0x35, 0x37, 0x31, 0x33, 0x47, 0x80, 0xd9, 0x49, 0xe4, 0xc2, 0x43,
	0x39, 0x86, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6,
	0x63, 0x39, 0x86, 0x24, 0x36, 0x70, 0xe8, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0x69,
	0x81, 0x26, 0xaa, 0x01, 0x00, 0x00,
}
