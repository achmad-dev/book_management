// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: category.proto

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

// The request message containing the category details.
type CreateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the created category details.
type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCategoryResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

// The request message for getting a category by ID.
type GetCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{2}
}

func (x *GetCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The response message containing the category details.
type GetCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetCategoryResponse) Reset() {
	*x = GetCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryResponse) ProtoMessage() {}

func (x *GetCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetCategoryResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type GetCategoryByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCategoryByNameRequest) Reset() {
	*x = GetCategoryByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByNameRequest) ProtoMessage() {}

func (x *GetCategoryByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByNameRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryByNameRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetCategoryByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCategoryByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetCategoryByNameResponse) Reset() {
	*x = GetCategoryByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByNameResponse) ProtoMessage() {}

func (x *GetCategoryByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByNameResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryByNameResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{5}
}

func (x *GetCategoryByNameResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

// The request message for updating a category.
type UpdateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the updated category details.
type UpdateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *UpdateCategoryResponse) Reset() {
	*x = UpdateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryResponse) ProtoMessage() {}

func (x *UpdateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCategoryResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

// The request message for deleting a category.
type DeleteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The response message for deleting a category.
type DeleteCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteCategoryResponse) Reset() {
	*x = DeleteCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryResponse) ProtoMessage() {}

func (x *DeleteCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCategoryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// The request message for listing all categories.
type ListCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCategoriesRequest) Reset() {
	*x = ListCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesRequest) ProtoMessage() {}

func (x *ListCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesRequest.ProtoReflect.Descriptor instead.
func (*ListCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{10}
}

// The response message containing the list of categories.
type ListCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *ListCategoriesResponse) Reset() {
	*x = ListCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesResponse) ProtoMessage() {}

func (x *ListCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesResponse.ProtoReflect.Descriptor instead.
func (*ListCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{11}
}

func (x *ListCategoriesResponse) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

// The category message containing the category details.
type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{12}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_category_proto protoreflect.FileDescriptor

var file_category_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x2e,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x3b, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8f, 0x04, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x68, 0x6d, 0x61, 0x64, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_category_proto_rawDescOnce sync.Once
	file_category_proto_rawDescData = file_category_proto_rawDesc
)

func file_category_proto_rawDescGZIP() []byte {
	file_category_proto_rawDescOnce.Do(func() {
		file_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_proto_rawDescData)
	})
	return file_category_proto_rawDescData
}

var file_category_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_category_proto_goTypes = []interface{}{
	(*CreateCategoryRequest)(nil),     // 0: genproto.CreateCategoryRequest
	(*CreateCategoryResponse)(nil),    // 1: genproto.CreateCategoryResponse
	(*GetCategoryRequest)(nil),        // 2: genproto.GetCategoryRequest
	(*GetCategoryResponse)(nil),       // 3: genproto.GetCategoryResponse
	(*GetCategoryByNameRequest)(nil),  // 4: genproto.GetCategoryByNameRequest
	(*GetCategoryByNameResponse)(nil), // 5: genproto.GetCategoryByNameResponse
	(*UpdateCategoryRequest)(nil),     // 6: genproto.UpdateCategoryRequest
	(*UpdateCategoryResponse)(nil),    // 7: genproto.UpdateCategoryResponse
	(*DeleteCategoryRequest)(nil),     // 8: genproto.DeleteCategoryRequest
	(*DeleteCategoryResponse)(nil),    // 9: genproto.DeleteCategoryResponse
	(*ListCategoriesRequest)(nil),     // 10: genproto.ListCategoriesRequest
	(*ListCategoriesResponse)(nil),    // 11: genproto.ListCategoriesResponse
	(*Category)(nil),                  // 12: genproto.Category
}
var file_category_proto_depIdxs = []int32{
	12, // 0: genproto.CreateCategoryResponse.category:type_name -> genproto.Category
	12, // 1: genproto.GetCategoryResponse.category:type_name -> genproto.Category
	12, // 2: genproto.GetCategoryByNameResponse.category:type_name -> genproto.Category
	12, // 3: genproto.UpdateCategoryResponse.category:type_name -> genproto.Category
	12, // 4: genproto.ListCategoriesResponse.categories:type_name -> genproto.Category
	0,  // 5: genproto.CategoryService.CreateCategory:input_type -> genproto.CreateCategoryRequest
	2,  // 6: genproto.CategoryService.GetCategory:input_type -> genproto.GetCategoryRequest
	4,  // 7: genproto.CategoryService.GetCategoryByName:input_type -> genproto.GetCategoryByNameRequest
	6,  // 8: genproto.CategoryService.UpdateCategory:input_type -> genproto.UpdateCategoryRequest
	8,  // 9: genproto.CategoryService.DeleteCategory:input_type -> genproto.DeleteCategoryRequest
	10, // 10: genproto.CategoryService.ListCategories:input_type -> genproto.ListCategoriesRequest
	1,  // 11: genproto.CategoryService.CreateCategory:output_type -> genproto.CreateCategoryResponse
	3,  // 12: genproto.CategoryService.GetCategory:output_type -> genproto.GetCategoryResponse
	5,  // 13: genproto.CategoryService.GetCategoryByName:output_type -> genproto.GetCategoryByNameResponse
	7,  // 14: genproto.CategoryService.UpdateCategory:output_type -> genproto.UpdateCategoryResponse
	9,  // 15: genproto.CategoryService.DeleteCategory:output_type -> genproto.DeleteCategoryResponse
	11, // 16: genproto.CategoryService.ListCategories:output_type -> genproto.ListCategoriesResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_category_proto_init() }
func file_category_proto_init() {
	if File_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryRequest); i {
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
		file_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
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
		file_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRequest); i {
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
		file_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryResponse); i {
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
		file_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByNameRequest); i {
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
		file_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByNameResponse); i {
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
		file_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryRequest); i {
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
		file_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryResponse); i {
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
		file_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryRequest); i {
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
		file_category_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryResponse); i {
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
		file_category_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoriesRequest); i {
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
		file_category_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoriesResponse); i {
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
		file_category_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_category_proto_goTypes,
		DependencyIndexes: file_category_proto_depIdxs,
		MessageInfos:      file_category_proto_msgTypes,
	}.Build()
	File_category_proto = out.File
	file_category_proto_rawDesc = nil
	file_category_proto_goTypes = nil
	file_category_proto_depIdxs = nil
}
