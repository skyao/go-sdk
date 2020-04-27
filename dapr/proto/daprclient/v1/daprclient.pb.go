// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/daprclient/v1/daprclient.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/dapr/go-sdk/dapr/proto/common/v1"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CloudEventEnvelope struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	SpecVersion          string   `protobuf:"bytes,4,opt,name=specVersion,proto3" json:"specVersion,omitempty"`
	DataContentType      string   `protobuf:"bytes,5,opt,name=data_content_type,json=dataContentType,proto3" json:"data_content_type,omitempty"`
	Topic                string   `protobuf:"bytes,6,opt,name=topic,proto3" json:"topic,omitempty"`
	Data                 *any.Any `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudEventEnvelope) Reset()         { *m = CloudEventEnvelope{} }
func (m *CloudEventEnvelope) String() string { return proto.CompactTextString(m) }
func (*CloudEventEnvelope) ProtoMessage()    {}
func (*CloudEventEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{0}
}

func (m *CloudEventEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudEventEnvelope.Unmarshal(m, b)
}
func (m *CloudEventEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudEventEnvelope.Marshal(b, m, deterministic)
}
func (m *CloudEventEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudEventEnvelope.Merge(m, src)
}
func (m *CloudEventEnvelope) XXX_Size() int {
	return xxx_messageInfo_CloudEventEnvelope.Size(m)
}
func (m *CloudEventEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudEventEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_CloudEventEnvelope proto.InternalMessageInfo

func (m *CloudEventEnvelope) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloudEventEnvelope) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *CloudEventEnvelope) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CloudEventEnvelope) GetSpecVersion() string {
	if m != nil {
		return m.SpecVersion
	}
	return ""
}

func (m *CloudEventEnvelope) GetDataContentType() string {
	if m != nil {
		return m.DataContentType
	}
	return ""
}

func (m *CloudEventEnvelope) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *CloudEventEnvelope) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindingEventEnvelope struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 *any.Any          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BindingEventEnvelope) Reset()         { *m = BindingEventEnvelope{} }
func (m *BindingEventEnvelope) String() string { return proto.CompactTextString(m) }
func (*BindingEventEnvelope) ProtoMessage()    {}
func (*BindingEventEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{1}
}

func (m *BindingEventEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingEventEnvelope.Unmarshal(m, b)
}
func (m *BindingEventEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingEventEnvelope.Marshal(b, m, deterministic)
}
func (m *BindingEventEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingEventEnvelope.Merge(m, src)
}
func (m *BindingEventEnvelope) XXX_Size() int {
	return xxx_messageInfo_BindingEventEnvelope.Size(m)
}
func (m *BindingEventEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingEventEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_BindingEventEnvelope proto.InternalMessageInfo

func (m *BindingEventEnvelope) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BindingEventEnvelope) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BindingEventEnvelope) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BindingResponseEnvelope struct {
	Data                 *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	To                   []string `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	State                []*State `protobuf:"bytes,3,rep,name=state,proto3" json:"state,omitempty"`
	Concurrency          string   `protobuf:"bytes,4,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindingResponseEnvelope) Reset()         { *m = BindingResponseEnvelope{} }
func (m *BindingResponseEnvelope) String() string { return proto.CompactTextString(m) }
func (*BindingResponseEnvelope) ProtoMessage()    {}
func (*BindingResponseEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{2}
}

func (m *BindingResponseEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingResponseEnvelope.Unmarshal(m, b)
}
func (m *BindingResponseEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingResponseEnvelope.Marshal(b, m, deterministic)
}
func (m *BindingResponseEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingResponseEnvelope.Merge(m, src)
}
func (m *BindingResponseEnvelope) XXX_Size() int {
	return xxx_messageInfo_BindingResponseEnvelope.Size(m)
}
func (m *BindingResponseEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingResponseEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_BindingResponseEnvelope proto.InternalMessageInfo

func (m *BindingResponseEnvelope) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BindingResponseEnvelope) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *BindingResponseEnvelope) GetState() []*State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *BindingResponseEnvelope) GetConcurrency() string {
	if m != nil {
		return m.Concurrency
	}
	return ""
}

type GetTopicSubscriptionsEnvelope struct {
	Topics               []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicSubscriptionsEnvelope) Reset()         { *m = GetTopicSubscriptionsEnvelope{} }
func (m *GetTopicSubscriptionsEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetTopicSubscriptionsEnvelope) ProtoMessage()    {}
func (*GetTopicSubscriptionsEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{3}
}

func (m *GetTopicSubscriptionsEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicSubscriptionsEnvelope.Unmarshal(m, b)
}
func (m *GetTopicSubscriptionsEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicSubscriptionsEnvelope.Marshal(b, m, deterministic)
}
func (m *GetTopicSubscriptionsEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicSubscriptionsEnvelope.Merge(m, src)
}
func (m *GetTopicSubscriptionsEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetTopicSubscriptionsEnvelope.Size(m)
}
func (m *GetTopicSubscriptionsEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicSubscriptionsEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicSubscriptionsEnvelope proto.InternalMessageInfo

func (m *GetTopicSubscriptionsEnvelope) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type GetBindingsSubscriptionsEnvelope struct {
	Bindings             []string `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBindingsSubscriptionsEnvelope) Reset()         { *m = GetBindingsSubscriptionsEnvelope{} }
func (m *GetBindingsSubscriptionsEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetBindingsSubscriptionsEnvelope) ProtoMessage()    {}
func (*GetBindingsSubscriptionsEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{4}
}

func (m *GetBindingsSubscriptionsEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBindingsSubscriptionsEnvelope.Unmarshal(m, b)
}
func (m *GetBindingsSubscriptionsEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBindingsSubscriptionsEnvelope.Marshal(b, m, deterministic)
}
func (m *GetBindingsSubscriptionsEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBindingsSubscriptionsEnvelope.Merge(m, src)
}
func (m *GetBindingsSubscriptionsEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetBindingsSubscriptionsEnvelope.Size(m)
}
func (m *GetBindingsSubscriptionsEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBindingsSubscriptionsEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetBindingsSubscriptionsEnvelope proto.InternalMessageInfo

func (m *GetBindingsSubscriptionsEnvelope) GetBindings() []string {
	if m != nil {
		return m.Bindings
	}
	return nil
}

type State struct {
	Key                  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *any.Any          `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Etag                 string            `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options              *StateOptions     `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{5}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *State) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *State) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *State) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *State) GetOptions() *StateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type StateOptions struct {
	Concurrency          string       `protobuf:"bytes,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Consistency          string       `protobuf:"bytes,2,opt,name=consistency,proto3" json:"consistency,omitempty"`
	RetryPolicy          *RetryPolicy `protobuf:"bytes,3,opt,name=retry_policy,json=retryPolicy,proto3" json:"retry_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StateOptions) Reset()         { *m = StateOptions{} }
func (m *StateOptions) String() string { return proto.CompactTextString(m) }
func (*StateOptions) ProtoMessage()    {}
func (*StateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{6}
}

func (m *StateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateOptions.Unmarshal(m, b)
}
func (m *StateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateOptions.Marshal(b, m, deterministic)
}
func (m *StateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateOptions.Merge(m, src)
}
func (m *StateOptions) XXX_Size() int {
	return xxx_messageInfo_StateOptions.Size(m)
}
func (m *StateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateOptions proto.InternalMessageInfo

func (m *StateOptions) GetConcurrency() string {
	if m != nil {
		return m.Concurrency
	}
	return ""
}

func (m *StateOptions) GetConsistency() string {
	if m != nil {
		return m.Consistency
	}
	return ""
}

func (m *StateOptions) GetRetryPolicy() *RetryPolicy {
	if m != nil {
		return m.RetryPolicy
	}
	return nil
}

type RetryPolicy struct {
	Threshold            int32              `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Pattern              string             `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Interval             *duration.Duration `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RetryPolicy) Reset()         { *m = RetryPolicy{} }
func (m *RetryPolicy) String() string { return proto.CompactTextString(m) }
func (*RetryPolicy) ProtoMessage()    {}
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb919fe08a3c35cb, []int{7}
}

func (m *RetryPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryPolicy.Unmarshal(m, b)
}
func (m *RetryPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryPolicy.Marshal(b, m, deterministic)
}
func (m *RetryPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryPolicy.Merge(m, src)
}
func (m *RetryPolicy) XXX_Size() int {
	return xxx_messageInfo_RetryPolicy.Size(m)
}
func (m *RetryPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RetryPolicy proto.InternalMessageInfo

func (m *RetryPolicy) GetThreshold() int32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *RetryPolicy) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *RetryPolicy) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterType((*CloudEventEnvelope)(nil), "dapr.proto.daprclient.v1.CloudEventEnvelope")
	proto.RegisterType((*BindingEventEnvelope)(nil), "dapr.proto.daprclient.v1.BindingEventEnvelope")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.daprclient.v1.BindingEventEnvelope.MetadataEntry")
	proto.RegisterType((*BindingResponseEnvelope)(nil), "dapr.proto.daprclient.v1.BindingResponseEnvelope")
	proto.RegisterType((*GetTopicSubscriptionsEnvelope)(nil), "dapr.proto.daprclient.v1.GetTopicSubscriptionsEnvelope")
	proto.RegisterType((*GetBindingsSubscriptionsEnvelope)(nil), "dapr.proto.daprclient.v1.GetBindingsSubscriptionsEnvelope")
	proto.RegisterType((*State)(nil), "dapr.proto.daprclient.v1.State")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.daprclient.v1.State.MetadataEntry")
	proto.RegisterType((*StateOptions)(nil), "dapr.proto.daprclient.v1.StateOptions")
	proto.RegisterType((*RetryPolicy)(nil), "dapr.proto.daprclient.v1.RetryPolicy")
}

func init() {
	proto.RegisterFile("dapr/proto/daprclient/v1/daprclient.proto", fileDescriptor_bb919fe08a3c35cb)
}

var fileDescriptor_bb919fe08a3c35cb = []byte{
	// 835 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x8e, 0x9d, 0x64, 0x37, 0x3d, 0x59, 0x96, 0x32, 0x5a, 0x16, 0xd7, 0xe5, 0x27, 0x98, 0x1f,
	0x85, 0x8a, 0x3a, 0x4a, 0x50, 0x55, 0x54, 0x10, 0xa2, 0xbb, 0x8d, 0x96, 0x5e, 0xa0, 0x54, 0x6e,
	0x55, 0x10, 0x37, 0x2b, 0xc7, 0x1e, 0xb2, 0x66, 0x9d, 0x19, 0x33, 0x33, 0xb6, 0x64, 0x89, 0xa7,
	0xe0, 0x9a, 0x3b, 0xee, 0x10, 0xcf, 0xc4, 0x0b, 0xf0, 0x12, 0x68, 0x7e, 0x9c, 0x75, 0x93, 0x38,
	0x11, 0xea, 0xdd, 0xf9, 0xf9, 0xe6, 0xcc, 0x39, 0xdf, 0x39, 0x73, 0x6c, 0xf8, 0x2c, 0x0e, 0x33,
	0x36, 0xca, 0x18, 0x15, 0x74, 0x24, 0xc5, 0x28, 0x4d, 0x30, 0x11, 0xa3, 0x62, 0x5c, 0xd3, 0x7c,
	0xe5, 0x46, 0x8e, 0xb4, 0x68, 0xd9, 0xaf, 0x39, 0x8b, 0xb1, 0x7b, 0x67, 0x41, 0xe9, 0x22, 0xc5,
	0x3a, 0xcc, 0x3c, 0xff, 0x79, 0x14, 0x92, 0x52, 0x03, 0xdd, 0xbb, 0xeb, 0x2e, 0xbc, 0xcc, 0x44,
	0xe5, 0x7c, 0x7f, 0xdd, 0x19, 0xe7, 0x2c, 0x14, 0x09, 0x25, 0xc6, 0xff, 0x61, 0x2d, 0xb9, 0x88,
	0x2e, 0x97, 0x94, 0xc8, 0xc4, 0xb4, 0xa4, 0x21, 0xde, 0x3f, 0x16, 0xa0, 0xf3, 0x94, 0xe6, 0xf1,
	0xb4, 0xc0, 0x44, 0x4c, 0x49, 0x81, 0x53, 0x9a, 0x61, 0x74, 0x0c, 0x76, 0x12, 0x3b, 0xd6, 0xc0,
	0x1a, 0xde, 0x0a, 0xec, 0x24, 0x46, 0xa7, 0x70, 0xc0, 0x69, 0xce, 0x22, 0xec, 0xd8, 0xca, 0x66,
	0x34, 0x84, 0xa0, 0x23, 0xca, 0x0c, 0x3b, 0x6d, 0x65, 0x55, 0x32, 0x1a, 0x40, 0x9f, 0x67, 0x38,
	0x7a, 0x89, 0x19, 0x4f, 0x28, 0x71, 0x3a, 0xca, 0x55, 0x37, 0xa1, 0x7b, 0xf0, 0x56, 0x1c, 0x8a,
	0xf0, 0x32, 0xa2, 0x44, 0x60, 0x22, 0x2e, 0x55, 0x88, 0xae, 0xc2, 0xbd, 0x29, 0x1d, 0xe7, 0xda,
	0xfe, 0x42, 0x46, 0x3b, 0x81, 0xae, 0xa0, 0x59, 0x12, 0x39, 0x07, 0xca, 0xaf, 0x15, 0x34, 0x84,
	0x8e, 0x04, 0x3a, 0x87, 0x03, 0x6b, 0xd8, 0x9f, 0x9c, 0xf8, 0x9a, 0x08, 0xbf, 0x22, 0xc2, 0x7f,
	0x4c, 0xca, 0x40, 0x21, 0xbc, 0x7f, 0x2d, 0x38, 0x39, 0x4b, 0x48, 0x9c, 0x90, 0xc5, 0xab, 0x25,
	0x22, 0xe8, 0x90, 0x70, 0x89, 0x4d, 0x91, 0x4a, 0x5e, 0x85, 0xb5, 0xf7, 0x85, 0x45, 0x3f, 0x42,
	0x6f, 0x89, 0x45, 0xa8, 0xd0, 0xed, 0x41, 0x7b, 0xd8, 0x9f, 0x7c, 0xed, 0x37, 0xf5, 0xd7, 0xdf,
	0x76, 0xbf, 0xff, 0xbd, 0x39, 0x3e, 0x25, 0x82, 0x95, 0xc1, 0x2a, 0x9a, 0xfb, 0x15, 0xbc, 0xf1,
	0x8a, 0x0b, 0xdd, 0x86, 0xf6, 0x35, 0x2e, 0x4d, 0x9e, 0x52, 0x94, 0x9c, 0x14, 0x61, 0x9a, 0x57,
	0xcd, 0xd0, 0xca, 0x23, 0xfb, 0x4b, 0xcb, 0xfb, 0xdb, 0x82, 0x77, 0xcc, 0x6d, 0x01, 0xe6, 0x19,
	0x25, 0x1c, 0xaf, 0x0a, 0xae, 0x8a, 0xb3, 0xf6, 0x16, 0x77, 0x0c, 0xb6, 0xa0, 0x8e, 0x3d, 0x68,
	0xcb, 0xee, 0x0b, 0x8a, 0x1e, 0x40, 0x97, 0x8b, 0x50, 0x60, 0x53, 0xe9, 0x07, 0xcd, 0x95, 0x3e,
	0x97, 0xb0, 0x40, 0xa3, 0xe5, 0x20, 0x44, 0x94, 0x44, 0x39, 0x63, 0x98, 0x44, 0x65, 0x35, 0x08,
	0x35, 0x93, 0xf7, 0x10, 0xde, 0xbb, 0xc0, 0xe2, 0x85, 0x6c, 0xe9, 0xf3, 0x7c, 0xce, 0x23, 0x96,
	0x64, 0x72, 0x7c, 0xf9, 0x2a, 0xe7, 0x53, 0x38, 0x50, 0x0d, 0xe7, 0x8e, 0xa5, 0xb2, 0x31, 0x9a,
	0xf7, 0x0d, 0x0c, 0x2e, 0xb0, 0x30, 0x95, 0xf2, 0xed, 0x67, 0x5d, 0xe8, 0xcd, 0x0d, 0xc0, 0x9c,
	0x5e, 0xe9, 0xde, 0x9f, 0x36, 0x74, 0x55, 0xae, 0x5b, 0xd8, 0xbd, 0x57, 0x67, 0xb7, 0x89, 0x28,
	0x0d, 0x91, 0x43, 0x84, 0x45, 0xb8, 0xa8, 0xe6, 0x5f, 0xca, 0xe8, 0x69, 0x6d, 0x34, 0x3a, 0x8a,
	0xb0, 0xfb, 0x7b, 0x08, 0x6b, 0x9a, 0x05, 0xf4, 0x2d, 0x1c, 0x52, 0x5d, 0x95, 0x7a, 0x1e, 0xfd,
	0xc9, 0xa7, 0x7b, 0x22, 0xcd, 0x34, 0x3a, 0xa8, 0x8e, 0xbd, 0xde, 0x34, 0xfd, 0x61, 0xc1, 0x51,
	0x3d, 0xec, 0x7a, 0x47, 0xad, 0x8d, 0x8e, 0x1a, 0x04, 0x4f, 0xb8, 0x50, 0x08, 0x7b, 0x85, 0xa8,
	0x4c, 0xe8, 0x3b, 0x38, 0x62, 0x58, 0xb0, 0xf2, 0x32, 0xa3, 0x69, 0x12, 0x95, 0x8a, 0xba, 0xfe,
	0xe4, 0x93, 0xe6, 0xc2, 0x02, 0x89, 0x7e, 0xa6, 0xc0, 0x41, 0x9f, 0xdd, 0x28, 0xde, 0x6f, 0xd0,
	0xaf, 0xf9, 0xd0, 0xbb, 0x70, 0x4b, 0x5c, 0x31, 0xcc, 0xaf, 0x68, 0xaa, 0x57, 0x57, 0x37, 0xb8,
	0x31, 0x20, 0x07, 0x0e, 0xb3, 0x50, 0x08, 0xcc, 0x88, 0x49, 0xaa, 0x52, 0xd1, 0x03, 0xe8, 0x25,
	0x44, 0x60, 0x56, 0x84, 0xa9, 0x49, 0xe6, 0xce, 0x46, 0xcb, 0x9f, 0x98, 0xc5, 0x1a, 0xac, 0xa0,
	0x93, 0xdf, 0x3b, 0x00, 0x4f, 0xc2, 0x8c, 0x9d, 0xab, 0x44, 0xd1, 0x0f, 0xd0, 0x9b, 0x91, 0xa7,
	0xa4, 0xa0, 0xd7, 0x18, 0x7d, 0x54, 0x2f, 0xc6, 0xac, 0xdb, 0x62, 0xec, 0x6b, 0x6f, 0x80, 0x7f,
	0xcd, 0x31, 0x17, 0xee, 0xc7, 0xbb, 0x41, 0xfa, 0xf1, 0x7a, 0x2d, 0xf4, 0x0b, 0xbc, 0xbd, 0xf5,
	0x8d, 0xa0, 0xd3, 0x8d, 0x2c, 0xa7, 0xf2, 0xdb, 0xe0, 0x3e, 0x6c, 0xa6, 0x72, 0xe7, 0x63, 0xf3,
	0x5a, 0x28, 0x03, 0xa7, 0xe9, 0x59, 0x35, 0x5e, 0xf7, 0x68, 0xe7, 0x75, 0x3b, 0x9f, 0xa8, 0xd7,
	0x42, 0x39, 0x1c, 0xcf, 0x48, 0x7d, 0x3f, 0x22, 0xff, 0xff, 0xed, 0x51, 0x77, 0xbc, 0x17, 0xbf,
	0xbe, 0x09, 0xbd, 0x16, 0x7a, 0x09, 0x47, 0x33, 0xa2, 0xa8, 0xd0, 0x97, 0x7e, 0xde, 0x1c, 0x64,
	0xf3, 0xeb, 0xe8, 0x36, 0x50, 0xe1, 0xb5, 0xce, 0x96, 0x00, 0x89, 0x0e, 0xe0, 0x17, 0xe3, 0xb3,
	0xdb, 0x37, 0xf3, 0xf1, 0x4c, 0x22, 0xf9, 0x4f, 0xa3, 0x45, 0x22, 0xae, 0xf2, 0xb9, 0xec, 0xb7,
	0xfa, 0x41, 0x18, 0x2d, 0xe8, 0x7d, 0x1e, 0x5f, 0x8f, 0x9a, 0xfe, 0x22, 0xfe, 0xb2, 0xef, 0xca,
	0x18, 0xbe, 0x0e, 0xe2, 0x3f, 0xce, 0x05, 0x5d, 0x60, 0xe2, 0x5f, 0xb0, 0x2c, 0xf2, 0x8b, 0xf1,
	0xfc, 0x40, 0x1d, 0xf9, 0xe2, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x1d, 0x8a, 0x4e, 0x86,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaprClientClient is the client API for DaprClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaprClientClient interface {
	OnInvoke(ctx context.Context, in *v1.InvokeRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error)
	GetTopicSubscriptions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTopicSubscriptionsEnvelope, error)
	GetBindingsSubscriptions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBindingsSubscriptionsEnvelope, error)
	OnBindingEvent(ctx context.Context, in *BindingEventEnvelope, opts ...grpc.CallOption) (*BindingResponseEnvelope, error)
	OnTopicEvent(ctx context.Context, in *CloudEventEnvelope, opts ...grpc.CallOption) (*empty.Empty, error)
}

type daprClientClient struct {
	cc *grpc.ClientConn
}

func NewDaprClientClient(cc *grpc.ClientConn) DaprClientClient {
	return &daprClientClient{cc}
}

func (c *daprClientClient) OnInvoke(ctx context.Context, in *v1.InvokeRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error) {
	out := new(v1.InvokeResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprclient.v1.DaprClient/OnInvoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClientClient) GetTopicSubscriptions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTopicSubscriptionsEnvelope, error) {
	out := new(GetTopicSubscriptionsEnvelope)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprclient.v1.DaprClient/GetTopicSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClientClient) GetBindingsSubscriptions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBindingsSubscriptionsEnvelope, error) {
	out := new(GetBindingsSubscriptionsEnvelope)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprclient.v1.DaprClient/GetBindingsSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClientClient) OnBindingEvent(ctx context.Context, in *BindingEventEnvelope, opts ...grpc.CallOption) (*BindingResponseEnvelope, error) {
	out := new(BindingResponseEnvelope)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprclient.v1.DaprClient/OnBindingEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClientClient) OnTopicEvent(ctx context.Context, in *CloudEventEnvelope, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprclient.v1.DaprClient/OnTopicEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaprClientServer is the server API for DaprClient service.
type DaprClientServer interface {
	OnInvoke(context.Context, *v1.InvokeRequest) (*v1.InvokeResponse, error)
	GetTopicSubscriptions(context.Context, *empty.Empty) (*GetTopicSubscriptionsEnvelope, error)
	GetBindingsSubscriptions(context.Context, *empty.Empty) (*GetBindingsSubscriptionsEnvelope, error)
	OnBindingEvent(context.Context, *BindingEventEnvelope) (*BindingResponseEnvelope, error)
	OnTopicEvent(context.Context, *CloudEventEnvelope) (*empty.Empty, error)
}

// UnimplementedDaprClientServer can be embedded to have forward compatible implementations.
type UnimplementedDaprClientServer struct {
}

func (*UnimplementedDaprClientServer) OnInvoke(ctx context.Context, req *v1.InvokeRequest) (*v1.InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnInvoke not implemented")
}
func (*UnimplementedDaprClientServer) GetTopicSubscriptions(ctx context.Context, req *empty.Empty) (*GetTopicSubscriptionsEnvelope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicSubscriptions not implemented")
}
func (*UnimplementedDaprClientServer) GetBindingsSubscriptions(ctx context.Context, req *empty.Empty) (*GetBindingsSubscriptionsEnvelope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBindingsSubscriptions not implemented")
}
func (*UnimplementedDaprClientServer) OnBindingEvent(ctx context.Context, req *BindingEventEnvelope) (*BindingResponseEnvelope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnBindingEvent not implemented")
}
func (*UnimplementedDaprClientServer) OnTopicEvent(ctx context.Context, req *CloudEventEnvelope) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnTopicEvent not implemented")
}

func RegisterDaprClientServer(s *grpc.Server, srv DaprClientServer) {
	s.RegisterService(&_DaprClient_serviceDesc, srv)
}

func _DaprClient_OnInvoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprClientServer).OnInvoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprclient.v1.DaprClient/OnInvoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprClientServer).OnInvoke(ctx, req.(*v1.InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaprClient_GetTopicSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprClientServer).GetTopicSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprclient.v1.DaprClient/GetTopicSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprClientServer).GetTopicSubscriptions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaprClient_GetBindingsSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprClientServer).GetBindingsSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprclient.v1.DaprClient/GetBindingsSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprClientServer).GetBindingsSubscriptions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaprClient_OnBindingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindingEventEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprClientServer).OnBindingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprclient.v1.DaprClient/OnBindingEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprClientServer).OnBindingEvent(ctx, req.(*BindingEventEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaprClient_OnTopicEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudEventEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprClientServer).OnTopicEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprclient.v1.DaprClient/OnTopicEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprClientServer).OnTopicEvent(ctx, req.(*CloudEventEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _DaprClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.daprclient.v1.DaprClient",
	HandlerType: (*DaprClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnInvoke",
			Handler:    _DaprClient_OnInvoke_Handler,
		},
		{
			MethodName: "GetTopicSubscriptions",
			Handler:    _DaprClient_GetTopicSubscriptions_Handler,
		},
		{
			MethodName: "GetBindingsSubscriptions",
			Handler:    _DaprClient_GetBindingsSubscriptions_Handler,
		},
		{
			MethodName: "OnBindingEvent",
			Handler:    _DaprClient_OnBindingEvent_Handler,
		},
		{
			MethodName: "OnTopicEvent",
			Handler:    _DaprClient_OnTopicEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/daprclient/v1/daprclient.proto",
}