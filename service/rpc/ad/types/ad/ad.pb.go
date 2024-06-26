// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: service/rpc/ad/ad.proto

package ad

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

// 空消息
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{0}
}

// 广告结构体
type AdInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	SiteUrl  string `protobuf:"bytes,6,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
}

func (x *AdInfo) Reset() {
	*x = AdInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfo) ProtoMessage() {}

func (x *AdInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfo.ProtoReflect.Descriptor instead.
func (*AdInfo) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{1}
}

func (x *AdInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AdInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AdInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *AdInfo) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

type AdminSetAdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	SiteUrl  string `protobuf:"bytes,6,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
}

func (x *AdminSetAdRequest) Reset() {
	*x = AdminSetAdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminSetAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminSetAdRequest) ProtoMessage() {}

func (x *AdminSetAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminSetAdRequest.ProtoReflect.Descriptor instead.
func (*AdminSetAdRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{2}
}

func (x *AdminSetAdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdminSetAdRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AdminSetAdRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminSetAdRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AdminSetAdRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *AdminSetAdRequest) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

type CheckCostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId int64 `protobuf:"varint,1,opt,name=AdId,proto3" json:"AdId,omitempty"`
}

func (x *CheckCostRequest) Reset() {
	*x = CheckCostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCostRequest) ProtoMessage() {}

func (x *CheckCostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCostRequest.ProtoReflect.Descriptor instead.
func (*CheckCostRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{3}
}

func (x *CheckCostRequest) GetAdId() int64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

type CheckCostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cost int64 `protobuf:"varint,1,opt,name=Cost,proto3" json:"Cost,omitempty"`
}

func (x *CheckCostReply) Reset() {
	*x = CheckCostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCostReply) ProtoMessage() {}

func (x *CheckCostReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCostReply.ProtoReflect.Descriptor instead.
func (*CheckCostReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{4}
}

func (x *CheckCostReply) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type AdListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdList []*AdInfo `protobuf:"bytes,1,rep,name=AdList,proto3" json:"AdList,omitempty"`
}

func (x *AdListReply) Reset() {
	*x = AdListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdListReply) ProtoMessage() {}

func (x *AdListReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdListReply.ProtoReflect.Descriptor instead.
func (*AdListReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{5}
}

func (x *AdListReply) GetAdList() []*AdInfo {
	if x != nil {
		return x.AdList
	}
	return nil
}

type AdInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ad_Id int64 `protobuf:"varint,1,opt,name=Ad_Id,json=AdId,proto3" json:"Ad_Id,omitempty"`
}

func (x *AdInfoRequest) Reset() {
	*x = AdInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfoRequest) ProtoMessage() {}

func (x *AdInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfoRequest.ProtoReflect.Descriptor instead.
func (*AdInfoRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{6}
}

func (x *AdInfoRequest) GetAd_Id() int64 {
	if x != nil {
		return x.Ad_Id
	}
	return 0
}

type AdInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ad *AdInfo `protobuf:"bytes,1,opt,name=Ad,proto3" json:"Ad,omitempty"`
}

func (x *AdInfoReply) Reset() {
	*x = AdInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfoReply) ProtoMessage() {}

func (x *AdInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfoReply.ProtoReflect.Descriptor instead.
func (*AdInfoReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{7}
}

func (x *AdInfoReply) GetAd() *AdInfo {
	if x != nil {
		return x.Ad
	}
	return nil
}

type ClickAdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId int64 `protobuf:"varint,1,opt,name=AdId,proto3" json:"AdId,omitempty"`
}

func (x *ClickAdRequest) Reset() {
	*x = ClickAdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickAdRequest) ProtoMessage() {}

func (x *ClickAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickAdRequest.ProtoReflect.Descriptor instead.
func (*ClickAdRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{8}
}

func (x *ClickAdRequest) GetAdId() int64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

type SendAdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	District string `protobuf:"bytes,1,opt,name=District,proto3" json:"District,omitempty"`
	Sex      string `protobuf:"bytes,2,opt,name=Sex,proto3" json:"Sex,omitempty"`
	Collage  string `protobuf:"bytes,3,opt,name=Collage,proto3" json:"Collage,omitempty"`
}

func (x *SendAdRequest) Reset() {
	*x = SendAdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAdRequest) ProtoMessage() {}

func (x *SendAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAdRequest.ProtoReflect.Descriptor instead.
func (*SendAdRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{9}
}

func (x *SendAdRequest) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *SendAdRequest) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *SendAdRequest) GetCollage() string {
	if x != nil {
		return x.Collage
	}
	return ""
}

type SendAdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ad *AdInfo `protobuf:"bytes,1,opt,name=Ad,proto3" json:"Ad,omitempty"`
}

func (x *SendAdReply) Reset() {
	*x = SendAdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_ad_ad_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAdReply) ProtoMessage() {}

func (x *SendAdReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_ad_ad_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAdReply.ProtoReflect.Descriptor instead.
func (*SendAdReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_ad_ad_proto_rawDescGZIP(), []int{10}
}

func (x *SendAdReply) GetAd() *AdInfo {
	if x != nil {
		return x.Ad
	}
	return nil
}

var File_service_rpc_ad_ad_proto protoreflect.FileDescriptor

var file_service_rpc_ad_ad_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x64,
	0x2f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x61, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x41, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x9f, 0x01,
	0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x22,
	0x26, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x41, 0x64, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x22, 0x31, 0x0a,
	0x0b, 0x41, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x06,
	0x41, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x64, 0x2e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x41, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x24, 0x0a, 0x0d, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x13, 0x0a, 0x05, 0x41, 0x64, 0x5f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x41, 0x64, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0b, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x02, 0x41, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x41,
	0x64, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x41, 0x64, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x67, 0x65,
	0x22, 0x29, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1a, 0x0a, 0x02, 0x41, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x64,
	0x2e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x41, 0x64, 0x32, 0x97, 0x02, 0x0a, 0x02,
	0x41, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x41, 0x64,
	0x12, 0x15, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x41, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x61, 0x64, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x35, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x14, 0x2e, 0x61, 0x64, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x64, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x61, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f,
	0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2c, 0x0a, 0x06, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x61, 0x64, 0x2e, 0x41,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61,
	0x64, 0x2e, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a,
	0x07, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x41, 0x64, 0x12, 0x12, 0x2e, 0x61, 0x64, 0x2e, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x61,
	0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x64, 0x12, 0x11, 0x2e, 0x61, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x61, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_rpc_ad_ad_proto_rawDescOnce sync.Once
	file_service_rpc_ad_ad_proto_rawDescData = file_service_rpc_ad_ad_proto_rawDesc
)

func file_service_rpc_ad_ad_proto_rawDescGZIP() []byte {
	file_service_rpc_ad_ad_proto_rawDescOnce.Do(func() {
		file_service_rpc_ad_ad_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rpc_ad_ad_proto_rawDescData)
	})
	return file_service_rpc_ad_ad_proto_rawDescData
}

var file_service_rpc_ad_ad_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_rpc_ad_ad_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: ad.Empty
	(*AdInfo)(nil),            // 1: ad.AdInfo
	(*AdminSetAdRequest)(nil), // 2: ad.AdminSetAdRequest
	(*CheckCostRequest)(nil),  // 3: ad.CheckCostRequest
	(*CheckCostReply)(nil),    // 4: ad.CheckCostReply
	(*AdListReply)(nil),       // 5: ad.AdListReply
	(*AdInfoRequest)(nil),     // 6: ad.AdInfoRequest
	(*AdInfoReply)(nil),       // 7: ad.AdInfoReply
	(*ClickAdRequest)(nil),    // 8: ad.ClickAdRequest
	(*SendAdRequest)(nil),     // 9: ad.SendAdRequest
	(*SendAdReply)(nil),       // 10: ad.SendAdReply
}
var file_service_rpc_ad_ad_proto_depIdxs = []int32{
	1,  // 0: ad.AdListReply.AdList:type_name -> ad.AdInfo
	1,  // 1: ad.AdInfoReply.Ad:type_name -> ad.AdInfo
	1,  // 2: ad.SendAdReply.Ad:type_name -> ad.AdInfo
	2,  // 3: ad.Ad.AdminSetAd:input_type -> ad.AdminSetAdRequest
	3,  // 4: ad.Ad.CheckCost:input_type -> ad.CheckCostRequest
	0,  // 5: ad.Ad.AdList:input_type -> ad.Empty
	6,  // 6: ad.Ad.AdInfo:input_type -> ad.AdInfoRequest
	8,  // 7: ad.Ad.ClickAd:input_type -> ad.ClickAdRequest
	9,  // 8: ad.Ad.SendAd:input_type -> ad.SendAdRequest
	0,  // 9: ad.Ad.AdminSetAd:output_type -> ad.Empty
	4,  // 10: ad.Ad.CheckCost:output_type -> ad.CheckCostReply
	5,  // 11: ad.Ad.AdList:output_type -> ad.AdListReply
	7,  // 12: ad.Ad.AdInfo:output_type -> ad.AdInfoReply
	0,  // 13: ad.Ad.ClickAd:output_type -> ad.Empty
	10, // 14: ad.Ad.SendAd:output_type -> ad.SendAdReply
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_rpc_ad_ad_proto_init() }
func file_service_rpc_ad_ad_proto_init() {
	if File_service_rpc_ad_ad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_rpc_ad_ad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfo); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminSetAdRequest); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCostRequest); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCostReply); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdListReply); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfoRequest); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfoReply); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickAdRequest); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAdRequest); i {
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
		file_service_rpc_ad_ad_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAdReply); i {
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
			RawDescriptor: file_service_rpc_ad_ad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_rpc_ad_ad_proto_goTypes,
		DependencyIndexes: file_service_rpc_ad_ad_proto_depIdxs,
		MessageInfos:      file_service_rpc_ad_ad_proto_msgTypes,
	}.Build()
	File_service_rpc_ad_ad_proto = out.File
	file_service_rpc_ad_ad_proto_rawDesc = nil
	file_service_rpc_ad_ad_proto_goTypes = nil
	file_service_rpc_ad_ad_proto_depIdxs = nil
}
