// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: author.proto

package genproto

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

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuthorRequest) Reset() {
	*x = GetAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorRequest) ProtoMessage() {}

func (x *GetAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetAuthorResponse) Reset() {
	*x = GetAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorResponse) ProtoMessage() {}

func (x *GetAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{2}
}

func (x *GetAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type GetAuthorByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAuthorByNameRequest) Reset() {
	*x = GetAuthorByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorByNameRequest) ProtoMessage() {}

func (x *GetAuthorByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorByNameRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorByNameRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthorByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAuthorByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetAuthorByNameResponse) Reset() {
	*x = GetAuthorByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorByNameResponse) ProtoMessage() {}

func (x *GetAuthorByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorByNameResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorByNameResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuthorByNameResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type CreateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateAuthorRequest) Reset() {
	*x = CreateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorRequest) ProtoMessage() {}

func (x *CreateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateAuthorResponse) Reset() {
	*x = CreateAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorResponse) ProtoMessage() {}

func (x *CreateAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type UpdateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateAuthorRequest) Reset() {
	*x = UpdateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorRequest) ProtoMessage() {}

func (x *UpdateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAuthorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *UpdateAuthorResponse) Reset() {
	*x = UpdateAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorResponse) ProtoMessage() {}

func (x *UpdateAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorResponse.ProtoReflect.Descriptor instead.
func (*UpdateAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type DeleteAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAuthorRequest) Reset() {
	*x = DeleteAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorRequest) ProtoMessage() {}

func (x *DeleteAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAuthorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteAuthorResponse) Reset() {
	*x = DeleteAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorResponse) ProtoMessage() {}

func (x *DeleteAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAuthorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteAuthorByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteAuthorByNameRequest) Reset() {
	*x = DeleteAuthorByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorByNameRequest) ProtoMessage() {}

func (x *DeleteAuthorByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorByNameRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthorByNameRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteAuthorByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteAuthorByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteAuthorByNameResponse) Reset() {
	*x = DeleteAuthorByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorByNameResponse) ProtoMessage() {}

func (x *DeleteAuthorByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorByNameResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthorByNameResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteAuthorByNameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAuthorsRequest) Reset() {
	*x = ListAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthorsRequest) ProtoMessage() {}

func (x *ListAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthorsRequest.ProtoReflect.Descriptor instead.
func (*ListAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{13}
}

type ListAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *ListAuthorsResponse) Reset() {
	*x = ListAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthorsResponse) ProtoMessage() {}

func (x *ListAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthorsResponse.ProtoReflect.Descriptor instead.
func (*ListAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{14}
}

func (x *ListAuthorsResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

var File_author_proto protoreflect.FileDescriptor

var file_author_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x29, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2f, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x36, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x32, 0xc7, 0x04, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x68, 0x6d, 0x61, 0x64,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_author_proto_rawDescOnce sync.Once
	file_author_proto_rawDescData = file_author_proto_rawDesc
)

func file_author_proto_rawDescGZIP() []byte {
	file_author_proto_rawDescOnce.Do(func() {
		file_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_author_proto_rawDescData)
	})
	return file_author_proto_rawDescData
}

var file_author_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_author_proto_goTypes = []interface{}{
	(*Author)(nil),                     // 0: genproto.Author
	(*GetAuthorRequest)(nil),           // 1: genproto.GetAuthorRequest
	(*GetAuthorResponse)(nil),          // 2: genproto.GetAuthorResponse
	(*GetAuthorByNameRequest)(nil),     // 3: genproto.GetAuthorByNameRequest
	(*GetAuthorByNameResponse)(nil),    // 4: genproto.GetAuthorByNameResponse
	(*CreateAuthorRequest)(nil),        // 5: genproto.CreateAuthorRequest
	(*CreateAuthorResponse)(nil),       // 6: genproto.CreateAuthorResponse
	(*UpdateAuthorRequest)(nil),        // 7: genproto.UpdateAuthorRequest
	(*UpdateAuthorResponse)(nil),       // 8: genproto.UpdateAuthorResponse
	(*DeleteAuthorRequest)(nil),        // 9: genproto.DeleteAuthorRequest
	(*DeleteAuthorResponse)(nil),       // 10: genproto.DeleteAuthorResponse
	(*DeleteAuthorByNameRequest)(nil),  // 11: genproto.DeleteAuthorByNameRequest
	(*DeleteAuthorByNameResponse)(nil), // 12: genproto.DeleteAuthorByNameResponse
	(*ListAuthorsRequest)(nil),         // 13: genproto.ListAuthorsRequest
	(*ListAuthorsResponse)(nil),        // 14: genproto.ListAuthorsResponse
}
var file_author_proto_depIdxs = []int32{
	0,  // 0: genproto.GetAuthorResponse.author:type_name -> genproto.Author
	0,  // 1: genproto.GetAuthorByNameResponse.author:type_name -> genproto.Author
	0,  // 2: genproto.CreateAuthorResponse.author:type_name -> genproto.Author
	0,  // 3: genproto.UpdateAuthorResponse.author:type_name -> genproto.Author
	0,  // 4: genproto.ListAuthorsResponse.authors:type_name -> genproto.Author
	1,  // 5: genproto.AuthorService.GetAuthor:input_type -> genproto.GetAuthorRequest
	3,  // 6: genproto.AuthorService.GetAuthorByName:input_type -> genproto.GetAuthorByNameRequest
	5,  // 7: genproto.AuthorService.CreateAuthor:input_type -> genproto.CreateAuthorRequest
	7,  // 8: genproto.AuthorService.UpdateAuthor:input_type -> genproto.UpdateAuthorRequest
	9,  // 9: genproto.AuthorService.DeleteAuthor:input_type -> genproto.DeleteAuthorRequest
	11, // 10: genproto.AuthorService.DeleteAuthorByName:input_type -> genproto.DeleteAuthorByNameRequest
	13, // 11: genproto.AuthorService.ListAuthors:input_type -> genproto.ListAuthorsRequest
	2,  // 12: genproto.AuthorService.GetAuthor:output_type -> genproto.GetAuthorResponse
	4,  // 13: genproto.AuthorService.GetAuthorByName:output_type -> genproto.GetAuthorByNameResponse
	6,  // 14: genproto.AuthorService.CreateAuthor:output_type -> genproto.CreateAuthorResponse
	8,  // 15: genproto.AuthorService.UpdateAuthor:output_type -> genproto.UpdateAuthorResponse
	10, // 16: genproto.AuthorService.DeleteAuthor:output_type -> genproto.DeleteAuthorResponse
	12, // 17: genproto.AuthorService.DeleteAuthorByName:output_type -> genproto.DeleteAuthorByNameResponse
	14, // 18: genproto.AuthorService.ListAuthors:output_type -> genproto.ListAuthorsResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_author_proto_init() }
func file_author_proto_init() {
	if File_author_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_author_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_author_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorRequest); i {
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
		file_author_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorResponse); i {
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
		file_author_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorByNameRequest); i {
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
		file_author_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorByNameResponse); i {
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
		file_author_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorRequest); i {
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
		file_author_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorResponse); i {
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
		file_author_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorRequest); i {
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
		file_author_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorResponse); i {
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
		file_author_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorRequest); i {
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
		file_author_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorResponse); i {
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
		file_author_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorByNameRequest); i {
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
		file_author_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorByNameResponse); i {
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
		file_author_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthorsRequest); i {
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
		file_author_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthorsResponse); i {
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
			RawDescriptor: file_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_author_proto_goTypes,
		DependencyIndexes: file_author_proto_depIdxs,
		MessageInfos:      file_author_proto_msgTypes,
	}.Build()
	File_author_proto = out.File
	file_author_proto_rawDesc = nil
	file_author_proto_goTypes = nil
	file_author_proto_depIdxs = nil
}
