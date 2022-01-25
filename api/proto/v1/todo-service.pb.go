// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: todo-service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 操作实体
type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reminder    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetReminder() *timestamppb.Timestamp {
	if x != nil {
		return x.Reminder
	}
	return nil
}

// 创建记录请求，插入记录
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api  string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Todo *ToDo  `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"` //插入对应的记录
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *CreateRequest) GetTodo() *ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

// 创建记录响应，回执
type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"` //创建请求的响应回执，返回插入记录的id
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *CreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 读取记录请求，读取记录
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"` // 读取该主键对应的记录
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReadRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 读取记录响应，返回该条记录
type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api  string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Todo *ToDo  `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"` //读取请求的响应，返回该条记录
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{4}
}

func (x *ReadResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReadResponse) GetTodo() *ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

// 读取全部记录请求，读取全部记录
type ReadAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
}

func (x *ReadAllRequest) Reset() {
	*x = ReadAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllRequest) ProtoMessage() {}

func (x *ReadAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllRequest.ProtoReflect.Descriptor instead.
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{5}
}

func (x *ReadAllRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

// 读取全部记录响应，返回所有记录
type ReadAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api   string  `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Todos []*ToDo `protobuf:"bytes,2,rep,name=todos,proto3" json:"todos,omitempty"` //读取全部记录请求的响应，返回所有记录
}

func (x *ReadAllResponse) Reset() {
	*x = ReadAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllResponse) ProtoMessage() {}

func (x *ReadAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllResponse.ProtoReflect.Descriptor instead.
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReadAllResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReadAllResponse) GetTodos() []*ToDo {
	if x != nil {
		return x.Todos
	}
	return nil
}

// 删除记录请求，删除记录
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id  int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"` //删除记录的对应主键
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 删除记录响应，返回删除状态
type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api    string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"` //删除请求的响应回执，返回删除是否成功的状态
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *DeleteResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// 更新记录请求，更新记录
type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api  string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Todo *ToDo  `protobuf:"bytes,2,opt,name=todo,proto3" json:"todo,omitempty"` // 需要更新的记录
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *UpdateRequest) GetTodo() *ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

// 更新记录响应，返回更新状态
type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api    string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"` //更新请求的响应，返回更新状态
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *UpdateResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_todo_service_proto protoreflect.FileDescriptor

var file_todo_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x04, 0x54, 0x6f,
	0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x22, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x22, 0x43, 0x0a, 0x0f,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70,
	0x69, 0x12, 0x1e, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x22, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x3f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x70, 0x69, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64,
	0x6f, 0x22, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xff, 0x01,
	0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_service_proto_rawDescOnce sync.Once
	file_todo_service_proto_rawDescData = file_todo_service_proto_rawDesc
)

func file_todo_service_proto_rawDescGZIP() []byte {
	file_todo_service_proto_rawDescOnce.Do(func() {
		file_todo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_service_proto_rawDescData)
	})
	return file_todo_service_proto_rawDescData
}

var file_todo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_todo_service_proto_goTypes = []interface{}{
	(*ToDo)(nil),                  // 0: v1.ToDo
	(*CreateRequest)(nil),         // 1: v1.CreateRequest
	(*CreateResponse)(nil),        // 2: v1.CreateResponse
	(*ReadRequest)(nil),           // 3: v1.ReadRequest
	(*ReadResponse)(nil),          // 4: v1.ReadResponse
	(*ReadAllRequest)(nil),        // 5: v1.ReadAllRequest
	(*ReadAllResponse)(nil),       // 6: v1.ReadAllResponse
	(*DeleteRequest)(nil),         // 7: v1.DeleteRequest
	(*DeleteResponse)(nil),        // 8: v1.DeleteResponse
	(*UpdateRequest)(nil),         // 9: v1.UpdateRequest
	(*UpdateResponse)(nil),        // 10: v1.UpdateResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_todo_service_proto_depIdxs = []int32{
	11, // 0: v1.ToDo.reminder:type_name -> google.protobuf.Timestamp
	0,  // 1: v1.CreateRequest.todo:type_name -> v1.ToDo
	0,  // 2: v1.ReadResponse.todo:type_name -> v1.ToDo
	0,  // 3: v1.ReadAllResponse.todos:type_name -> v1.ToDo
	0,  // 4: v1.UpdateRequest.todo:type_name -> v1.ToDo
	1,  // 5: v1.ToDoService.Create:input_type -> v1.CreateRequest
	7,  // 6: v1.ToDoService.Delete:input_type -> v1.DeleteRequest
	9,  // 7: v1.ToDoService.Update:input_type -> v1.UpdateRequest
	3,  // 8: v1.ToDoService.Read:input_type -> v1.ReadRequest
	5,  // 9: v1.ToDoService.ReadAll:input_type -> v1.ReadAllRequest
	2,  // 10: v1.ToDoService.Create:output_type -> v1.CreateResponse
	8,  // 11: v1.ToDoService.Delete:output_type -> v1.DeleteResponse
	10, // 12: v1.ToDoService.Update:output_type -> v1.UpdateResponse
	4,  // 13: v1.ToDoService.Read:output_type -> v1.ReadResponse
	6,  // 14: v1.ToDoService.ReadAll:output_type -> v1.ReadAllResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_todo_service_proto_init() }
func file_todo_service_proto_init() {
	if File_todo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_todo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_todo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_todo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_todo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_todo_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllRequest); i {
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
		file_todo_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllResponse); i {
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
		file_todo_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_todo_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_todo_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_todo_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_todo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_service_proto_goTypes,
		DependencyIndexes: file_todo_service_proto_depIdxs,
		MessageInfos:      file_todo_service_proto_msgTypes,
	}.Build()
	File_todo_service_proto = out.File
	file_todo_service_proto_rawDesc = nil
	file_todo_service_proto_goTypes = nil
	file_todo_service_proto_depIdxs = nil
}
