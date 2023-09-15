// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/administration_event.proto

package storage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// AdministrationEventType is the storage representation of the event type.
//
// Refer to v1.AdministrationEventType for a more detailed doc.
type AdministrationEventType int32

const (
	AdministrationEventType_ADMINISTRATION_EVENT_TYPE_UNKNOWN     AdministrationEventType = 0
	AdministrationEventType_ADMINISTRATION_EVENT_TYPE_GENERIC     AdministrationEventType = 1
	AdministrationEventType_ADMINISTRATION_EVENT_TYPE_LOG_MESSAGE AdministrationEventType = 2
)

var AdministrationEventType_name = map[int32]string{
	0: "ADMINISTRATION_EVENT_TYPE_UNKNOWN",
	1: "ADMINISTRATION_EVENT_TYPE_GENERIC",
	2: "ADMINISTRATION_EVENT_TYPE_LOG_MESSAGE",
}

var AdministrationEventType_value = map[string]int32{
	"ADMINISTRATION_EVENT_TYPE_UNKNOWN":     0,
	"ADMINISTRATION_EVENT_TYPE_GENERIC":     1,
	"ADMINISTRATION_EVENT_TYPE_LOG_MESSAGE": 2,
}

func (x AdministrationEventType) String() string {
	return proto.EnumName(AdministrationEventType_name, int32(x))
}

func (AdministrationEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b517965aab4cb00c, []int{0}
}

// AdministrationEventLevel is the storage representation of the event level.
//
// Refer to v1.AdministrationEventLevel for a more detailed doc.
type AdministrationEventLevel int32

const (
	AdministrationEventLevel_ADMINISTRATION_EVENT_LEVEL_UNKNOWN AdministrationEventLevel = 0
	AdministrationEventLevel_ADMINISTRATION_EVENT_LEVEL_INFO    AdministrationEventLevel = 1
	AdministrationEventLevel_ADMINISTRATION_EVENT_LEVEL_SUCCESS AdministrationEventLevel = 2
	AdministrationEventLevel_ADMINISTRATION_EVENT_LEVEL_WARNING AdministrationEventLevel = 3
	AdministrationEventLevel_ADMINISTRATION_EVENT_LEVEL_ERROR   AdministrationEventLevel = 4
)

var AdministrationEventLevel_name = map[int32]string{
	0: "ADMINISTRATION_EVENT_LEVEL_UNKNOWN",
	1: "ADMINISTRATION_EVENT_LEVEL_INFO",
	2: "ADMINISTRATION_EVENT_LEVEL_SUCCESS",
	3: "ADMINISTRATION_EVENT_LEVEL_WARNING",
	4: "ADMINISTRATION_EVENT_LEVEL_ERROR",
}

var AdministrationEventLevel_value = map[string]int32{
	"ADMINISTRATION_EVENT_LEVEL_UNKNOWN": 0,
	"ADMINISTRATION_EVENT_LEVEL_INFO":    1,
	"ADMINISTRATION_EVENT_LEVEL_SUCCESS": 2,
	"ADMINISTRATION_EVENT_LEVEL_WARNING": 3,
	"ADMINISTRATION_EVENT_LEVEL_ERROR":   4,
}

func (x AdministrationEventLevel) String() string {
	return proto.EnumName(AdministrationEventLevel_name, int32(x))
}

func (AdministrationEventLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b517965aab4cb00c, []int{1}
}

// AdministrationEvent is the storage representation of administrative events in Central.
//
// Refer to v1.AdministrationEvent for a more detailed doc.
type AdministrationEvent struct {
	// The id is a UUIDv5 generated deterministically from the tuple (type, level, domain,
	// resource_type, resource_id, message). It is used for deduplication of events.
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`
	Type                 AdministrationEventType  `protobuf:"varint,2,opt,name=type,proto3,enum=storage.AdministrationEventType" json:"type,omitempty" search:"Event Type,hidden"`
	Level                AdministrationEventLevel `protobuf:"varint,3,opt,name=level,proto3,enum=storage.AdministrationEventLevel" json:"level,omitempty" search:"Event Level,hidden"`
	Message              string                   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Hint                 string                   `protobuf:"bytes,5,opt,name=hint,proto3" json:"hint,omitempty"`
	Domain               string                   `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty" search:"Event Domain,hidden"`
	ResourceType         string                   `protobuf:"bytes,7,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty" search:"Resource Type,hidden"`
	ResourceId           string                   `protobuf:"bytes,8,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	NumOccurrences       int64                    `protobuf:"varint,9,opt,name=num_occurrences,json=numOccurrences,proto3" json:"num_occurrences,omitempty"`
	LastOccurredAt       *types.Timestamp         `protobuf:"bytes,10,opt,name=last_occurred_at,json=lastOccurredAt,proto3" json:"last_occurred_at,omitempty" search:"Last Updated,hidden"`
	CreatedAt            *types.Timestamp         `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" search:"Created Time,hidden"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AdministrationEvent) Reset()         { *m = AdministrationEvent{} }
func (m *AdministrationEvent) String() string { return proto.CompactTextString(m) }
func (*AdministrationEvent) ProtoMessage()    {}
func (*AdministrationEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b517965aab4cb00c, []int{0}
}
func (m *AdministrationEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdministrationEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdministrationEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdministrationEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdministrationEvent.Merge(m, src)
}
func (m *AdministrationEvent) XXX_Size() int {
	return m.Size()
}
func (m *AdministrationEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AdministrationEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AdministrationEvent proto.InternalMessageInfo

func (m *AdministrationEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AdministrationEvent) GetType() AdministrationEventType {
	if m != nil {
		return m.Type
	}
	return AdministrationEventType_ADMINISTRATION_EVENT_TYPE_UNKNOWN
}

func (m *AdministrationEvent) GetLevel() AdministrationEventLevel {
	if m != nil {
		return m.Level
	}
	return AdministrationEventLevel_ADMINISTRATION_EVENT_LEVEL_UNKNOWN
}

func (m *AdministrationEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AdministrationEvent) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *AdministrationEvent) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *AdministrationEvent) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *AdministrationEvent) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AdministrationEvent) GetNumOccurrences() int64 {
	if m != nil {
		return m.NumOccurrences
	}
	return 0
}

func (m *AdministrationEvent) GetLastOccurredAt() *types.Timestamp {
	if m != nil {
		return m.LastOccurredAt
	}
	return nil
}

func (m *AdministrationEvent) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *AdministrationEvent) MessageClone() proto.Message {
	return m.Clone()
}
func (m *AdministrationEvent) Clone() *AdministrationEvent {
	if m == nil {
		return nil
	}
	cloned := new(AdministrationEvent)
	*cloned = *m

	cloned.LastOccurredAt = m.LastOccurredAt.Clone()
	cloned.CreatedAt = m.CreatedAt.Clone()
	return cloned
}

func init() {
	proto.RegisterEnum("storage.AdministrationEventType", AdministrationEventType_name, AdministrationEventType_value)
	proto.RegisterEnum("storage.AdministrationEventLevel", AdministrationEventLevel_name, AdministrationEventLevel_value)
	proto.RegisterType((*AdministrationEvent)(nil), "storage.AdministrationEvent")
}

func init() {
	proto.RegisterFile("storage/administration_event.proto", fileDescriptor_b517965aab4cb00c)
}

var fileDescriptor_b517965aab4cb00c = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x71, 0x08, 0xf0, 0x71, 0xf9, 0x4a, 0xad, 0xa1, 0x12, 0x6e, 0xda, 0xc6, 0xc6, 0x2d,
	0x25, 0x20, 0xe4, 0x48, 0xb4, 0x0f, 0x15, 0x6f, 0x4e, 0x30, 0x91, 0xd5, 0xe0, 0x54, 0x8e, 0x81,
	0xfe, 0x79, 0xb0, 0x8c, 0x67, 0x08, 0x16, 0xb1, 0x27, 0xb5, 0xc7, 0xa8, 0x6c, 0xa3, 0x4f, 0xdd,
	0x44, 0xf7, 0xd1, 0xc7, 0x6e, 0xa0, 0x51, 0x45, 0x77, 0x90, 0x15, 0x54, 0x9e, 0xc4, 0x69, 0x23,
	0x91, 0xa0, 0xbe, 0x8d, 0xcf, 0x9c, 0xf9, 0x9d, 0x3b, 0xd7, 0xd7, 0x06, 0x35, 0x61, 0x34, 0xf6,
	0x3a, 0xa4, 0xea, 0xe1, 0x30, 0x88, 0x82, 0x84, 0xc5, 0x1e, 0x0b, 0x68, 0xe4, 0x92, 0x2b, 0x12,
	0x31, 0xad, 0x17, 0x53, 0x46, 0xd1, 0xd2, 0xc8, 0x53, 0x92, 0x3b, 0x94, 0x76, 0xba, 0xa4, 0xca,
	0xe5, 0xb3, 0xf4, 0xbc, 0xca, 0x82, 0x90, 0x24, 0xcc, 0x0b, 0x7b, 0x43, 0x67, 0xe9, 0x41, 0x87,
	0x76, 0x28, 0x5f, 0x56, 0xb3, 0xd5, 0x50, 0x55, 0xbf, 0x2e, 0xc0, 0x9a, 0x3e, 0x81, 0x37, 0x32,
	0x3a, 0xda, 0x82, 0x42, 0x80, 0x25, 0x41, 0x11, 0x2a, 0xcb, 0xb5, 0xf5, 0x41, 0x5f, 0x5e, 0x4b,
	0x3e, 0x76, 0xf7, 0xd5, 0xde, 0xe5, 0x2e, 0xbb, 0xee, 0x91, 0x4a, 0x9a, 0x06, 0x78, 0x5b, 0xb5,
	0x0b, 0x01, 0x46, 0x0e, 0x14, 0x33, 0x49, 0x2a, 0x28, 0x42, 0x65, 0x75, 0x4f, 0xd1, 0x46, 0xf5,
	0x68, 0xb7, 0x40, 0x9d, 0xeb, 0x1e, 0xa9, 0x95, 0x07, 0x7d, 0xb9, 0x94, 0x10, 0x2f, 0xf6, 0x2f,
	0xf6, 0x55, 0x2e, 0x2b, 0x99, 0xbe, 0x7b, 0x11, 0x60, 0x4c, 0x22, 0xd5, 0xe6, 0x34, 0xf4, 0x16,
	0x16, 0xba, 0xe4, 0x8a, 0x74, 0xa5, 0x79, 0x8e, 0xdd, 0x98, 0x85, 0x6d, 0x66, 0xc6, 0x9a, 0x3c,
	0xe8, 0xcb, 0x8f, 0x26, 0xb9, 0x7c, 0x63, 0x0c, 0x1e, 0x02, 0x91, 0x04, 0x4b, 0x21, 0x49, 0x12,
	0xaf, 0x43, 0xa4, 0x62, 0x76, 0x3b, 0x3b, 0x7f, 0x44, 0x08, 0x8a, 0x17, 0x41, 0xc4, 0xa4, 0x05,
	0x2e, 0xf3, 0x35, 0x7a, 0x05, 0x8b, 0x98, 0x86, 0x5e, 0x10, 0x49, 0x8b, 0xbc, 0x15, 0xca, 0xa0,
	0x2f, 0x3f, 0x9e, 0x4c, 0x39, 0xe0, 0xfb, 0xe3, 0x98, 0x91, 0x1f, 0x1d, 0xc2, 0xbd, 0x98, 0x24,
	0x34, 0x8d, 0x7d, 0xe2, 0xf2, 0x06, 0x2d, 0x71, 0xc0, 0xc6, 0xa0, 0x2f, 0x3f, 0xc9, 0x01, 0xf6,
	0xc8, 0x30, 0xd9, 0x81, 0xff, 0xf3, 0x73, 0x99, 0x8a, 0x64, 0x58, 0x19, 0x73, 0x02, 0x2c, 0xfd,
	0xc7, 0x8b, 0x83, 0x5c, 0x32, 0x31, 0xda, 0x82, 0xfb, 0x51, 0x1a, 0xba, 0xd4, 0xf7, 0xd3, 0x38,
	0x26, 0x91, 0x4f, 0x12, 0x69, 0x59, 0x11, 0x2a, 0xf3, 0xf6, 0x6a, 0x94, 0x86, 0xad, 0x3f, 0x2a,
	0x3a, 0x07, 0xb1, 0xeb, 0x25, 0x2c, 0x77, 0x62, 0xd7, 0x63, 0x12, 0x28, 0x42, 0x65, 0x65, 0xaf,
	0xa4, 0x0d, 0x87, 0x47, 0xcb, 0x87, 0x47, 0x73, 0xf2, 0xe1, 0x99, 0xbc, 0x71, 0xd3, 0x4b, 0x98,
	0x72, 0xdc, 0xc3, 0x1e, 0x23, 0x78, 0x5c, 0xef, 0x6a, 0x46, 0x1d, 0x05, 0x61, 0x9d, 0xa1, 0x0f,
	0x00, 0x7e, 0x4c, 0x32, 0x4b, 0x96, 0xb0, 0xf2, 0x6f, 0x09, 0xf5, 0xe1, 0x49, 0x25, 0xdb, 0x1e,
	0x27, 0x2c, 0x8f, 0x78, 0x3a, 0xdb, 0xf9, 0x2c, 0xc0, 0xfa, 0x94, 0xd1, 0x42, 0x9b, 0xb0, 0xa1,
	0x1f, 0x1c, 0x99, 0x96, 0xd9, 0x76, 0x6c, 0xdd, 0x31, 0x5b, 0x96, 0x6b, 0x9c, 0x18, 0x96, 0xe3,
	0x3a, 0xef, 0xde, 0x18, 0xee, 0xb1, 0xf5, 0xda, 0x6a, 0x9d, 0x5a, 0xe2, 0xdc, 0x6c, 0x5b, 0xc3,
	0xb0, 0x0c, 0xdb, 0xac, 0x8b, 0x02, 0xda, 0x86, 0xcd, 0xe9, 0xb6, 0x66, 0xab, 0xe1, 0x1e, 0x19,
	0xed, 0xb6, 0xde, 0x30, 0xc4, 0xc2, 0xce, 0x0f, 0x01, 0xa4, 0x69, 0x83, 0x89, 0x9e, 0x83, 0x7a,
	0x2b, 0xa7, 0x69, 0x9c, 0x18, 0xcd, 0xbf, 0xca, 0x7a, 0x0a, 0xf2, 0x0c, 0x9f, 0x69, 0x1d, 0xb6,
	0x44, 0xe1, 0x0e, 0x58, 0xfb, 0xb8, 0x5e, 0x37, 0xda, 0x6d, 0xb1, 0x70, 0x87, 0xef, 0x54, 0xb7,
	0x2d, 0xd3, 0x6a, 0x88, 0xf3, 0xe8, 0x19, 0x28, 0x33, 0x7c, 0x86, 0x6d, 0xb7, 0x6c, 0xb1, 0x58,
	0x7b, 0xf9, 0xed, 0xa6, 0x2c, 0x7c, 0xbf, 0x29, 0x0b, 0x3f, 0x6f, 0xca, 0xc2, 0x97, 0x5f, 0xe5,
	0x39, 0x78, 0x18, 0x50, 0x2d, 0x61, 0x9e, 0x7f, 0x19, 0xd3, 0x4f, 0xc3, 0x77, 0x9a, 0x7f, 0xa1,
	0xef, 0xf3, 0x3f, 0xd2, 0xd9, 0x22, 0xd7, 0x5f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x60, 0x9f,
	0x91, 0x0c, 0xc7, 0x04, 0x00, 0x00,
}

func (m *AdministrationEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdministrationEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdministrationEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CreatedAt != nil {
		{
			size, err := m.CreatedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdministrationEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.LastOccurredAt != nil {
		{
			size, err := m.LastOccurredAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdministrationEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.NumOccurrences != 0 {
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(m.NumOccurrences))
		i--
		dAtA[i] = 0x48
	}
	if len(m.ResourceId) > 0 {
		i -= len(m.ResourceId)
		copy(dAtA[i:], m.ResourceId)
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(len(m.ResourceId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ResourceType) > 0 {
		i -= len(m.ResourceType)
		copy(dAtA[i:], m.ResourceType)
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(len(m.ResourceType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Hint) > 0 {
		i -= len(m.Hint)
		copy(dAtA[i:], m.Hint)
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(len(m.Hint)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x22
	}
	if m.Level != 0 {
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if m.Type != 0 {
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAdministrationEvent(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdministrationEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdministrationEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AdministrationEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovAdministrationEvent(uint64(m.Type))
	}
	if m.Level != 0 {
		n += 1 + sovAdministrationEvent(uint64(m.Level))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	l = len(m.Hint)
	if l > 0 {
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	l = len(m.ResourceType)
	if l > 0 {
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	l = len(m.ResourceId)
	if l > 0 {
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	if m.NumOccurrences != 0 {
		n += 1 + sovAdministrationEvent(uint64(m.NumOccurrences))
	}
	if m.LastOccurredAt != nil {
		l = m.LastOccurredAt.Size()
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	if m.CreatedAt != nil {
		l = m.CreatedAt.Size()
		n += 1 + l + sovAdministrationEvent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAdministrationEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdministrationEvent(x uint64) (n int) {
	return sovAdministrationEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AdministrationEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdministrationEvent
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
			return fmt.Errorf("proto: AdministrationEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdministrationEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
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
					return ErrIntOverflowAdministrationEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= AdministrationEventType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= AdministrationEventLevel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOccurrences", wireType)
			}
			m.NumOccurrences = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOccurrences |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastOccurredAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastOccurredAt == nil {
				m.LastOccurredAt = &types.Timestamp{}
			}
			if err := m.LastOccurredAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationEvent
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
				return ErrInvalidLengthAdministrationEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = &types.Timestamp{}
			}
			if err := m.CreatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdministrationEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdministrationEvent
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
func skipAdministrationEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdministrationEvent
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
					return 0, ErrIntOverflowAdministrationEvent
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
					return 0, ErrIntOverflowAdministrationEvent
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
				return 0, ErrInvalidLengthAdministrationEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdministrationEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdministrationEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdministrationEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdministrationEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdministrationEvent = fmt.Errorf("proto: unexpected end of group")
)
