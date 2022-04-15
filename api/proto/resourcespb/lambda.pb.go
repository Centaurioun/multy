// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/proto/resourcespb/lambda.proto

package resourcespb

import (
	commonpb "github.com/multycloud/multy/api/proto/commonpb"
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

type CreateLambdaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *LambdaArgs `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CreateLambdaRequest) Reset() {
	*x = CreateLambdaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLambdaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLambdaRequest) ProtoMessage() {}

func (x *CreateLambdaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLambdaRequest.ProtoReflect.Descriptor instead.
func (*CreateLambdaRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_lambda_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLambdaRequest) GetResource() *LambdaArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ReadLambdaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadLambdaRequest) Reset() {
	*x = ReadLambdaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLambdaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLambdaRequest) ProtoMessage() {}

func (x *ReadLambdaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLambdaRequest.ProtoReflect.Descriptor instead.
func (*ReadLambdaRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_lambda_proto_rawDescGZIP(), []int{1}
}

func (x *ReadLambdaRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateLambdaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string      `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resource   *LambdaArgs `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateLambdaRequest) Reset() {
	*x = UpdateLambdaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLambdaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLambdaRequest) ProtoMessage() {}

func (x *UpdateLambdaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLambdaRequest.ProtoReflect.Descriptor instead.
func (*UpdateLambdaRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_lambda_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateLambdaRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateLambdaRequest) GetResource() *LambdaArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteLambdaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteLambdaRequest) Reset() {
	*x = DeleteLambdaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLambdaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLambdaRequest) ProtoMessage() {}

func (x *DeleteLambdaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLambdaRequest.ProtoReflect.Descriptor instead.
func (*DeleteLambdaRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_lambda_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteLambdaRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type LambdaArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters   *commonpb.ResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name               string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Runtime            string                       `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SourceCodeObjectId string                       `protobuf:"bytes,4,opt,name=source_code_object_id,json=sourceCodeObjectId,proto3" json:"source_code_object_id,omitempty"`
	// outputs
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *LambdaArgs) Reset() {
	*x = LambdaArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaArgs) ProtoMessage() {}

func (x *LambdaArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LambdaArgs.ProtoReflect.Descriptor instead.
func (*LambdaArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_lambda_proto_rawDescGZIP(), []int{4}
}

func (x *LambdaArgs) GetCommonParameters() *commonpb.ResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *LambdaArgs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LambdaArgs) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *LambdaArgs) GetSourceCodeObjectId() string {
	if x != nil {
		return x.SourceCodeObjectId
	}
	return ""
}

func (x *LambdaArgs) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type LambdaResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters   *commonpb.CommonResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name               string                             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Runtime            string                             `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SourceCodeObjectId string                             `protobuf:"bytes,4,opt,name=source_code_object_id,json=sourceCodeObjectId,proto3" json:"source_code_object_id,omitempty"`
}

func (x *LambdaResource) Reset() {
	*x = LambdaResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaResource) ProtoMessage() {}

func (x *LambdaResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_lambda_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LambdaResource.ProtoReflect.Descriptor instead.
func (*LambdaResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_lambda_proto_rawDescGZIP(), []int{5}
}

func (x *LambdaResource) GetCommonParameters() *commonpb.CommonResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *LambdaResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LambdaResource) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *LambdaResource) GetSourceCodeObjectId() string {
	if x != nil {
		return x.SourceCodeObjectId
	}
	return ""
}

var File_api_proto_resourcespb_lambda_proto protoreflect.FileDescriptor

var file_api_proto_resourcespb_lambda_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61,
	0x41, 0x72, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x34,
	0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x6d, 0x62, 0x64, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x41, 0x72, 0x67, 0x73, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x22, 0xd2, 0x01, 0x0a, 0x0a, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x41, 0x72, 0x67, 0x73,
	0x12, 0x51, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x72, 0x67,
	0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x4c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x31, 0x0a, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x42, 0x5e, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_resourcespb_lambda_proto_rawDescOnce sync.Once
	file_api_proto_resourcespb_lambda_proto_rawDescData = file_api_proto_resourcespb_lambda_proto_rawDesc
)

func file_api_proto_resourcespb_lambda_proto_rawDescGZIP() []byte {
	file_api_proto_resourcespb_lambda_proto_rawDescOnce.Do(func() {
		file_api_proto_resourcespb_lambda_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resourcespb_lambda_proto_rawDescData)
	})
	return file_api_proto_resourcespb_lambda_proto_rawDescData
}

var file_api_proto_resourcespb_lambda_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_resourcespb_lambda_proto_goTypes = []interface{}{
	(*CreateLambdaRequest)(nil),               // 0: dev.multy.resources.CreateLambdaRequest
	(*ReadLambdaRequest)(nil),                 // 1: dev.multy.resources.ReadLambdaRequest
	(*UpdateLambdaRequest)(nil),               // 2: dev.multy.resources.UpdateLambdaRequest
	(*DeleteLambdaRequest)(nil),               // 3: dev.multy.resources.DeleteLambdaRequest
	(*LambdaArgs)(nil),                        // 4: dev.multy.resources.LambdaArgs
	(*LambdaResource)(nil),                    // 5: dev.multy.resources.LambdaResource
	(*commonpb.ResourceCommonArgs)(nil),       // 6: dev.multy.common.ResourceCommonArgs
	(*commonpb.CommonResourceParameters)(nil), // 7: dev.multy.common.CommonResourceParameters
}
var file_api_proto_resourcespb_lambda_proto_depIdxs = []int32{
	4, // 0: dev.multy.resources.CreateLambdaRequest.resource:type_name -> dev.multy.resources.LambdaArgs
	4, // 1: dev.multy.resources.UpdateLambdaRequest.resource:type_name -> dev.multy.resources.LambdaArgs
	6, // 2: dev.multy.resources.LambdaArgs.common_parameters:type_name -> dev.multy.common.ResourceCommonArgs
	7, // 3: dev.multy.resources.LambdaResource.common_parameters:type_name -> dev.multy.common.CommonResourceParameters
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_resourcespb_lambda_proto_init() }
func file_api_proto_resourcespb_lambda_proto_init() {
	if File_api_proto_resourcespb_lambda_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resourcespb_lambda_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLambdaRequest); i {
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
		file_api_proto_resourcespb_lambda_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLambdaRequest); i {
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
		file_api_proto_resourcespb_lambda_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLambdaRequest); i {
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
		file_api_proto_resourcespb_lambda_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLambdaRequest); i {
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
		file_api_proto_resourcespb_lambda_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LambdaArgs); i {
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
		file_api_proto_resourcespb_lambda_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LambdaResource); i {
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
			RawDescriptor: file_api_proto_resourcespb_lambda_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resourcespb_lambda_proto_goTypes,
		DependencyIndexes: file_api_proto_resourcespb_lambda_proto_depIdxs,
		MessageInfos:      file_api_proto_resourcespb_lambda_proto_msgTypes,
	}.Build()
	File_api_proto_resourcespb_lambda_proto = out.File
	file_api_proto_resourcespb_lambda_proto_rawDesc = nil
	file_api_proto_resourcespb_lambda_proto_goTypes = nil
	file_api_proto_resourcespb_lambda_proto_depIdxs = nil
}
