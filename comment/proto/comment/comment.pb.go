// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.3
// source: proto/comment/comment.proto

package go_micro_service_comment

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID    string `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`        // 订单id，即评价id
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`          // 评价用户id
	ProductID  string `protobuf:"bytes,3,opt,name=productID,proto3" json:"productID,omitempty"`    // 评价商品id
	Star       int32  `protobuf:"varint,4,opt,name=star,proto3" json:"star,omitempty"`             // 推荐等级（评价星级）
	Content    string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`        // 评价内容
	CreateTime int64  `protobuf:"varint,6,opt,name=createTime,proto3" json:"createTime,omitempty"` // 评价创建时间
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Comment) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Comment) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *Comment) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Resp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Resp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetCommentsByProductIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *GetCommentsByProductIdReq) Reset() {
	*x = GetCommentsByProductIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByProductIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByProductIdReq) ProtoMessage() {}

func (x *GetCommentsByProductIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByProductIdReq.ProtoReflect.Descriptor instead.
func (*GetCommentsByProductIdReq) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{2}
}

func (x *GetCommentsByProductIdReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type GetCommentsByProductIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Comments []*Comment `protobuf:"bytes,3,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetCommentsByProductIdResp) Reset() {
	*x = GetCommentsByProductIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByProductIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByProductIdResp) ProtoMessage() {}

func (x *GetCommentsByProductIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByProductIdResp.ProtoReflect.Descriptor instead.
func (*GetCommentsByProductIdResp) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentsByProductIdResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetCommentsByProductIdResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetCommentsByProductIdResp) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetCommentByOrderIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *GetCommentByOrderIdReq) Reset() {
	*x = GetCommentByOrderIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByOrderIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByOrderIdReq) ProtoMessage() {}

func (x *GetCommentByOrderIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByOrderIdReq.ProtoReflect.Descriptor instead.
func (*GetCommentByOrderIdReq) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentByOrderIdReq) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type GetCommentByOrderIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Comment *Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *GetCommentByOrderIdResp) Reset() {
	*x = GetCommentByOrderIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentByOrderIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentByOrderIdResp) ProtoMessage() {}

func (x *GetCommentByOrderIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentByOrderIdResp.ProtoReflect.Descriptor instead.
func (*GetCommentByOrderIdResp) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentByOrderIdResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetCommentByOrderIdResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetCommentByOrderIdResp) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type AddCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *AddCommentReq) Reset() {
	*x = AddCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentReq) ProtoMessage() {}

func (x *AddCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentReq.ProtoReflect.Descriptor instead.
func (*AddCommentReq) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{6}
}

func (x *AddCommentReq) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_proto_comment_comment_proto protoreflect.FileDescriptor

var file_proto_comment_comment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x40, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x81,
	0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x3d, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x32, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x7c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x32, 0xef, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x85,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x34,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comment_comment_proto_rawDescOnce sync.Once
	file_proto_comment_comment_proto_rawDescData = file_proto_comment_comment_proto_rawDesc
)

func file_proto_comment_comment_proto_rawDescGZIP() []byte {
	file_proto_comment_comment_proto_rawDescOnce.Do(func() {
		file_proto_comment_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comment_comment_proto_rawDescData)
	})
	return file_proto_comment_comment_proto_rawDescData
}

var file_proto_comment_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_comment_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),                    // 0: go.micro.service.comment.Comment
	(*Resp)(nil),                       // 1: go.micro.service.comment.Resp
	(*GetCommentsByProductIdReq)(nil),  // 2: go.micro.service.comment.GetCommentsByProductIdReq
	(*GetCommentsByProductIdResp)(nil), // 3: go.micro.service.comment.GetCommentsByProductIdResp
	(*GetCommentByOrderIdReq)(nil),     // 4: go.micro.service.comment.GetCommentByOrderIdReq
	(*GetCommentByOrderIdResp)(nil),    // 5: go.micro.service.comment.GetCommentByOrderIdResp
	(*AddCommentReq)(nil),              // 6: go.micro.service.comment.AddCommentReq
}
var file_proto_comment_comment_proto_depIdxs = []int32{
	0, // 0: go.micro.service.comment.GetCommentsByProductIdResp.comments:type_name -> go.micro.service.comment.Comment
	0, // 1: go.micro.service.comment.GetCommentByOrderIdResp.comment:type_name -> go.micro.service.comment.Comment
	0, // 2: go.micro.service.comment.AddCommentReq.comment:type_name -> go.micro.service.comment.Comment
	6, // 3: go.micro.service.comment.CommentService.AddComment:input_type -> go.micro.service.comment.AddCommentReq
	2, // 4: go.micro.service.comment.CommentService.GetCommentsByProductId:input_type -> go.micro.service.comment.GetCommentsByProductIdReq
	4, // 5: go.micro.service.comment.CommentService.GetCommentByOrderId:input_type -> go.micro.service.comment.GetCommentByOrderIdReq
	1, // 6: go.micro.service.comment.CommentService.AddComment:output_type -> go.micro.service.comment.Resp
	3, // 7: go.micro.service.comment.CommentService.GetCommentsByProductId:output_type -> go.micro.service.comment.GetCommentsByProductIdResp
	5, // 8: go.micro.service.comment.CommentService.GetCommentByOrderId:output_type -> go.micro.service.comment.GetCommentByOrderIdResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_comment_comment_proto_init() }
func file_proto_comment_comment_proto_init() {
	if File_proto_comment_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comment_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_proto_comment_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
		file_proto_comment_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByProductIdReq); i {
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
		file_proto_comment_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByProductIdResp); i {
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
		file_proto_comment_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentByOrderIdReq); i {
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
		file_proto_comment_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentByOrderIdResp); i {
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
		file_proto_comment_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentReq); i {
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
			RawDescriptor: file_proto_comment_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_comment_comment_proto_goTypes,
		DependencyIndexes: file_proto_comment_comment_proto_depIdxs,
		MessageInfos:      file_proto_comment_comment_proto_msgTypes,
	}.Build()
	File_proto_comment_comment_proto = out.File
	file_proto_comment_comment_proto_rawDesc = nil
	file_proto_comment_comment_proto_goTypes = nil
	file_proto_comment_comment_proto_depIdxs = nil
}
