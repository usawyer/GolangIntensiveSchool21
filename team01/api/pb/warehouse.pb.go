// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: warehouse.proto

package pb

import (
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

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host              string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port              int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ReplicationFactor int32  `protobuf:"varint,3,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HeartbeatRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *HeartbeatRequest) GetReplicationFactor() int32 {
	if x != nil {
		return x.ReplicationFactor
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KnownNodes        []string `protobuf:"bytes,1,rep,name=known_nodes,json=knownNodes,proto3" json:"known_nodes,omitempty"`
	ReplicationFactor int32    `protobuf:"varint,2,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
	Error             string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatResponse) GetKnownNodes() []string {
	if x != nil {
		return x.KnownNodes
	}
	return nil
}

func (x *HeartbeatResponse) GetReplicationFactor() int32 {
	if x != nil {
		return x.ReplicationFactor
	}
	return 0
}

func (x *HeartbeatResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SetDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Document string `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *SetDocumentRequest) Reset() {
	*x = SetDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDocumentRequest) ProtoMessage() {}

func (x *SetDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDocumentRequest.ProtoReflect.Descriptor instead.
func (*SetDocumentRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *SetDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetDocumentRequest) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

type SetDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicasCreated int32  `protobuf:"varint,1,opt,name=replicas_created,json=replicasCreated,proto3" json:"replicas_created,omitempty"`
	Error           string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SetDocumentResponse) Reset() {
	*x = SetDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDocumentResponse) ProtoMessage() {}

func (x *SetDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDocumentResponse.ProtoReflect.Descriptor instead.
func (*SetDocumentResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *SetDocumentResponse) GetReplicasCreated() int32 {
	if x != nil {
		return x.ReplicasCreated
	}
	return 0
}

func (x *SetDocumentResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetDocumentRequest) Reset() {
	*x = GetDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentRequest) ProtoMessage() {}

func (x *GetDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentRequest.ProtoReflect.Descriptor instead.
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{4}
}

func (x *GetDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetDocumentResponse) Reset() {
	*x = GetDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentResponse) ProtoMessage() {}

func (x *GetDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentResponse.ProtoReflect.Descriptor instead.
func (*GetDocumentResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{5}
}

func (x *GetDocumentResponse) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

func (x *GetDocumentResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteDocumentRequest) Reset() {
	*x = DeleteDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentRequest) ProtoMessage() {}

func (x *DeleteDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicasDeleted int32  `protobuf:"varint,1,opt,name=replicas_deleted,json=replicasDeleted,proto3" json:"replicas_deleted,omitempty"`
	Error           string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteDocumentResponse) Reset() {
	*x = DeleteDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentResponse) ProtoMessage() {}

func (x *DeleteDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentResponse.ProtoReflect.Descriptor instead.
func (*DeleteDocumentResponse) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDocumentResponse) GetReplicasDeleted() int32 {
	if x != nil {
		return x.ReplicasDeleted
	}
	return 0
}

func (x *DeleteDocumentResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_warehouse_proto protoreflect.FileDescriptor

var file_warehouse_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2d, 0x0a,
	0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x79, 0x0a, 0x11,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x13, 0x53,
	0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x59, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xfe, 0x01, 0x0a, 0x09, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e,
	0x53, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_proto_rawDescOnce sync.Once
	file_warehouse_proto_rawDescData = file_warehouse_proto_rawDesc
)

func file_warehouse_proto_rawDescGZIP() []byte {
	file_warehouse_proto_rawDescOnce.Do(func() {
		file_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_proto_rawDescData)
	})
	return file_warehouse_proto_rawDescData
}

var file_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_warehouse_proto_goTypes = []interface{}{
	(*HeartbeatRequest)(nil),       // 0: HeartbeatRequest
	(*HeartbeatResponse)(nil),      // 1: HeartbeatResponse
	(*SetDocumentRequest)(nil),     // 2: SetDocumentRequest
	(*SetDocumentResponse)(nil),    // 3: SetDocumentResponse
	(*GetDocumentRequest)(nil),     // 4: GetDocumentRequest
	(*GetDocumentResponse)(nil),    // 5: GetDocumentResponse
	(*DeleteDocumentRequest)(nil),  // 6: DeleteDocumentRequest
	(*DeleteDocumentResponse)(nil), // 7: DeleteDocumentResponse
}
var file_warehouse_proto_depIdxs = []int32{
	0, // 0: Warehouse.Heartbeat:input_type -> HeartbeatRequest
	2, // 1: Warehouse.SetDocument:input_type -> SetDocumentRequest
	4, // 2: Warehouse.GetDocument:input_type -> GetDocumentRequest
	6, // 3: Warehouse.DeleteDocument:input_type -> DeleteDocumentRequest
	1, // 4: Warehouse.Heartbeat:output_type -> HeartbeatResponse
	3, // 5: Warehouse.SetDocument:output_type -> SetDocumentResponse
	5, // 6: Warehouse.GetDocument:output_type -> GetDocumentResponse
	7, // 7: Warehouse.DeleteDocument:output_type -> DeleteDocumentResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_warehouse_proto_init() }
func file_warehouse_proto_init() {
	if File_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_warehouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_warehouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDocumentRequest); i {
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
		file_warehouse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDocumentResponse); i {
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
		file_warehouse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocumentRequest); i {
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
		file_warehouse_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocumentResponse); i {
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
		file_warehouse_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentRequest); i {
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
		file_warehouse_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentResponse); i {
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
			RawDescriptor: file_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouse_proto_goTypes,
		DependencyIndexes: file_warehouse_proto_depIdxs,
		MessageInfos:      file_warehouse_proto_msgTypes,
	}.Build()
	File_warehouse_proto = out.File
	file_warehouse_proto_rawDesc = nil
	file_warehouse_proto_goTypes = nil
	file_warehouse_proto_depIdxs = nil
}
