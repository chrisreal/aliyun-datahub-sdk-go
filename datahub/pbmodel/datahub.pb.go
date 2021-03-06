// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.2.0
// source: datahub.proto

package pbmodel

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StringPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *StringPair) Reset() {
	*x = StringPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringPair) ProtoMessage() {}

func (x *StringPair) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringPair.ProtoReflect.Descriptor instead.
func (*StringPair) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{0}
}

func (x *StringPair) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *StringPair) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type FieldData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (x *FieldData) Reset() {
	*x = FieldData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldData) ProtoMessage() {}

func (x *FieldData) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldData.ProtoReflect.Descriptor instead.
func (*FieldData) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{1}
}

func (x *FieldData) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type RecordAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*StringPair `protobuf:"bytes,1,rep,name=attributes" json:"attributes,omitempty"`
}

func (x *RecordAttributes) Reset() {
	*x = RecordAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordAttributes) ProtoMessage() {}

func (x *RecordAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordAttributes.ProtoReflect.Descriptor instead.
func (*RecordAttributes) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{2}
}

func (x *RecordAttributes) GetAttributes() []*StringPair {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type RecordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*FieldData `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (x *RecordData) Reset() {
	*x = RecordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordData) ProtoMessage() {}

func (x *RecordData) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordData.ProtoReflect.Descriptor instead.
func (*RecordData) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{3}
}

func (x *RecordData) GetData() []*FieldData {
	if x != nil {
		return x.Data
	}
	return nil
}

type RecordEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId      *string           `protobuf:"bytes,1,opt,name=shard_id,json=shardId" json:"shard_id,omitempty"`
	HashKey      *string           `protobuf:"bytes,2,opt,name=hash_key,json=hashKey" json:"hash_key,omitempty"`
	PartitionKey *string           `protobuf:"bytes,3,opt,name=partition_key,json=partitionKey" json:"partition_key,omitempty"`
	Cursor       *string           `protobuf:"bytes,4,opt,name=cursor" json:"cursor,omitempty"`
	NextCursor   *string           `protobuf:"bytes,5,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	Sequence     *int64            `protobuf:"varint,6,opt,name=sequence" json:"sequence,omitempty"`
	SystemTime   *int64            `protobuf:"varint,7,opt,name=system_time,json=systemTime" json:"system_time,omitempty"`
	Attributes   *RecordAttributes `protobuf:"bytes,8,opt,name=attributes" json:"attributes,omitempty"`
	Data         *RecordData       `protobuf:"bytes,9,req,name=data" json:"data,omitempty"`
}

func (x *RecordEntry) Reset() {
	*x = RecordEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordEntry) ProtoMessage() {}

func (x *RecordEntry) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordEntry.ProtoReflect.Descriptor instead.
func (*RecordEntry) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{4}
}

func (x *RecordEntry) GetShardId() string {
	if x != nil && x.ShardId != nil {
		return *x.ShardId
	}
	return ""
}

func (x *RecordEntry) GetHashKey() string {
	if x != nil && x.HashKey != nil {
		return *x.HashKey
	}
	return ""
}

func (x *RecordEntry) GetPartitionKey() string {
	if x != nil && x.PartitionKey != nil {
		return *x.PartitionKey
	}
	return ""
}

func (x *RecordEntry) GetCursor() string {
	if x != nil && x.Cursor != nil {
		return *x.Cursor
	}
	return ""
}

func (x *RecordEntry) GetNextCursor() string {
	if x != nil && x.NextCursor != nil {
		return *x.NextCursor
	}
	return ""
}

func (x *RecordEntry) GetSequence() int64 {
	if x != nil && x.Sequence != nil {
		return *x.Sequence
	}
	return 0
}

func (x *RecordEntry) GetSystemTime() int64 {
	if x != nil && x.SystemTime != nil {
		return *x.SystemTime
	}
	return 0
}

func (x *RecordEntry) GetAttributes() *RecordAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *RecordEntry) GetData() *RecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PutRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*RecordEntry `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (x *PutRecordsRequest) Reset() {
	*x = PutRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRecordsRequest) ProtoMessage() {}

func (x *PutRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRecordsRequest.ProtoReflect.Descriptor instead.
func (*PutRecordsRequest) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{5}
}

func (x *PutRecordsRequest) GetRecords() []*RecordEntry {
	if x != nil {
		return x.Records
	}
	return nil
}

type FailedRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index        *int32  `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	ErrorCode    *string `protobuf:"bytes,2,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ErrorMessage *string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
}

func (x *FailedRecord) Reset() {
	*x = FailedRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailedRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailedRecord) ProtoMessage() {}

func (x *FailedRecord) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailedRecord.ProtoReflect.Descriptor instead.
func (*FailedRecord) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{6}
}

func (x *FailedRecord) GetIndex() int32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *FailedRecord) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

func (x *FailedRecord) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

type PutRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailedCount   *int32          `protobuf:"varint,1,opt,name=failed_count,json=failedCount" json:"failed_count,omitempty"`
	FailedRecords []*FailedRecord `protobuf:"bytes,2,rep,name=failed_records,json=failedRecords" json:"failed_records,omitempty"`
}

func (x *PutRecordsResponse) Reset() {
	*x = PutRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRecordsResponse) ProtoMessage() {}

func (x *PutRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRecordsResponse.ProtoReflect.Descriptor instead.
func (*PutRecordsResponse) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{7}
}

func (x *PutRecordsResponse) GetFailedCount() int32 {
	if x != nil && x.FailedCount != nil {
		return *x.FailedCount
	}
	return 0
}

func (x *PutRecordsResponse) GetFailedRecords() []*FailedRecord {
	if x != nil {
		return x.FailedRecords
	}
	return nil
}

type GetRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor *string `protobuf:"bytes,1,req,name=cursor" json:"cursor,omitempty"`
	Limit  *int32  `protobuf:"varint,2,opt,name=limit,def=1" json:"limit,omitempty"`
}

// Default values for GetRecordsRequest fields.
const (
	Default_GetRecordsRequest_Limit = int32(1)
)

func (x *GetRecordsRequest) Reset() {
	*x = GetRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsRequest) ProtoMessage() {}

func (x *GetRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsRequest.ProtoReflect.Descriptor instead.
func (*GetRecordsRequest) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{8}
}

func (x *GetRecordsRequest) GetCursor() string {
	if x != nil && x.Cursor != nil {
		return *x.Cursor
	}
	return ""
}

func (x *GetRecordsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return Default_GetRecordsRequest_Limit
}

type GetRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextCursor    *string        `protobuf:"bytes,1,req,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	RecordCount   *int32         `protobuf:"varint,2,req,name=record_count,json=recordCount" json:"record_count,omitempty"`
	StartSequence *int64         `protobuf:"varint,3,opt,name=start_sequence,json=startSequence" json:"start_sequence,omitempty"`
	Records       []*RecordEntry `protobuf:"bytes,4,rep,name=records" json:"records,omitempty"`
}

func (x *GetRecordsResponse) Reset() {
	*x = GetRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datahub_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsResponse) ProtoMessage() {}

func (x *GetRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datahub_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsResponse.ProtoReflect.Descriptor instead.
func (*GetRecordsResponse) Descriptor() ([]byte, []int) {
	return file_datahub_proto_rawDescGZIP(), []int{9}
}

func (x *GetRecordsResponse) GetNextCursor() string {
	if x != nil && x.NextCursor != nil {
		return *x.NextCursor
	}
	return ""
}

func (x *GetRecordsResponse) GetRecordCount() int32 {
	if x != nil && x.RecordCount != nil {
		return *x.RecordCount
	}
	return 0
}

func (x *GetRecordsResponse) GetStartSequence() int64 {
	if x != nil && x.StartSequence != nil {
		return *x.StartSequence
	}
	return 0
}

func (x *GetRecordsResponse) GetRecords() []*RecordEntry {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_datahub_proto protoreflect.FileDescriptor

var file_datahub_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x34, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21,
	0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x47, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xc2, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68,
	0x61, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x68, 0x0a, 0x0c, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x75, 0x0a, 0x12, 0x50, 0x75, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a,
	0x0e, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0d, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x44, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x31, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0xaf, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x39, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75,
	0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42,
	0x0d, 0x44, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
}

var (
	file_datahub_proto_rawDescOnce sync.Once
	file_datahub_proto_rawDescData = file_datahub_proto_rawDesc
)

func file_datahub_proto_rawDescGZIP() []byte {
	file_datahub_proto_rawDescOnce.Do(func() {
		file_datahub_proto_rawDescData = protoimpl.X.CompressGZIP(file_datahub_proto_rawDescData)
	})
	return file_datahub_proto_rawDescData
}

var file_datahub_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_datahub_proto_goTypes = []interface{}{
	(*StringPair)(nil),         // 0: pbmodel.StringPair
	(*FieldData)(nil),          // 1: pbmodel.FieldData
	(*RecordAttributes)(nil),   // 2: pbmodel.RecordAttributes
	(*RecordData)(nil),         // 3: pbmodel.RecordData
	(*RecordEntry)(nil),        // 4: pbmodel.RecordEntry
	(*PutRecordsRequest)(nil),  // 5: pbmodel.PutRecordsRequest
	(*FailedRecord)(nil),       // 6: pbmodel.FailedRecord
	(*PutRecordsResponse)(nil), // 7: pbmodel.PutRecordsResponse
	(*GetRecordsRequest)(nil),  // 8: pbmodel.GetRecordsRequest
	(*GetRecordsResponse)(nil), // 9: pbmodel.GetRecordsResponse
}
var file_datahub_proto_depIdxs = []int32{
	0, // 0: pbmodel.RecordAttributes.attributes:type_name -> pbmodel.StringPair
	1, // 1: pbmodel.RecordData.data:type_name -> pbmodel.FieldData
	2, // 2: pbmodel.RecordEntry.attributes:type_name -> pbmodel.RecordAttributes
	3, // 3: pbmodel.RecordEntry.data:type_name -> pbmodel.RecordData
	4, // 4: pbmodel.PutRecordsRequest.records:type_name -> pbmodel.RecordEntry
	6, // 5: pbmodel.PutRecordsResponse.failed_records:type_name -> pbmodel.FailedRecord
	4, // 6: pbmodel.GetRecordsResponse.records:type_name -> pbmodel.RecordEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_datahub_proto_init() }
func file_datahub_proto_init() {
	if File_datahub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datahub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordAttributes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailedRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRecordsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datahub_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datahub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datahub_proto_goTypes,
		DependencyIndexes: file_datahub_proto_depIdxs,
		MessageInfos:      file_datahub_proto_msgTypes,
	}.Build()
	File_datahub_proto = out.File
	file_datahub_proto_rawDesc = nil
	file_datahub_proto_goTypes = nil
	file_datahub_proto_depIdxs = nil
}
