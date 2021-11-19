// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: GetGlobalState.proto

package getglobalstate

import (
	common "github.com/gandalf1024/go-futu-api/pb/common"
	_ "github.com/gandalf1024/go-futu-api/pb/qotcommon"
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

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID *uint64 `protobuf:"varint,1,req,name=userID" json:"userID,omitempty"` //历史原因，目前已废弃，填0即可
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetGlobalState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_GetGlobalState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_GetGlobalState_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetUserID() uint64 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketHK       *int32                `protobuf:"varint,1,req,name=marketHK" json:"marketHK,omitempty"`              //Qot_Common.QotMarketState,港股主板市场状态
	MarketUS       *int32                `protobuf:"varint,2,req,name=marketUS" json:"marketUS,omitempty"`              //Qot_Common.QotMarketState,美股Nasdaq市场状态
	MarketSH       *int32                `protobuf:"varint,3,req,name=marketSH" json:"marketSH,omitempty"`              //Qot_Common.QotMarketState,沪市状态
	MarketSZ       *int32                `protobuf:"varint,4,req,name=marketSZ" json:"marketSZ,omitempty"`              //Qot_Common.QotMarketState,深市状态
	MarketHKFuture *int32                `protobuf:"varint,5,req,name=marketHKFuture" json:"marketHKFuture,omitempty"`  //Qot_Common.QotMarketState,港股期货市场状态
	MarketUSFuture *int32                `protobuf:"varint,15,opt,name=marketUSFuture" json:"marketUSFuture,omitempty"` //Qot_Common.QotMarketState,美国期货市场状态
	QotLogined     *bool                 `protobuf:"varint,6,req,name=qotLogined" json:"qotLogined,omitempty"`          //是否登陆行情服务器
	TrdLogined     *bool                 `protobuf:"varint,7,req,name=trdLogined" json:"trdLogined,omitempty"`          //是否登陆交易服务器
	ServerVer      *int32                `protobuf:"varint,8,req,name=serverVer" json:"serverVer,omitempty"`            //版本号
	ServerBuildNo  *int32                `protobuf:"varint,9,req,name=serverBuildNo" json:"serverBuildNo,omitempty"`    //buildNo
	Time           *int64                `protobuf:"varint,10,req,name=time" json:"time,omitempty"`                     //当前服务器时间
	LocalTime      *float64              `protobuf:"fixed64,11,opt,name=localTime" json:"localTime,omitempty"`          //当前本地时间
	ProgramStatus  *common.ProgramStatus `protobuf:"bytes,12,opt,name=programStatus" json:"programStatus,omitempty"`    //当前程序状态
	QotSvrIpAddr   *string               `protobuf:"bytes,13,opt,name=qotSvrIpAddr" json:"qotSvrIpAddr,omitempty"`
	TrdSvrIpAddr   *string               `protobuf:"bytes,14,opt,name=trdSvrIpAddr" json:"trdSvrIpAddr,omitempty"`
	ConnID         *uint64               `protobuf:"varint,16,opt,name=connID" json:"connID,omitempty"` //此连接的连接ID，连接的唯一标识
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetGlobalState_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_GetGlobalState_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_GetGlobalState_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetMarketHK() int32 {
	if x != nil && x.MarketHK != nil {
		return *x.MarketHK
	}
	return 0
}

func (x *S2C) GetMarketUS() int32 {
	if x != nil && x.MarketUS != nil {
		return *x.MarketUS
	}
	return 0
}

func (x *S2C) GetMarketSH() int32 {
	if x != nil && x.MarketSH != nil {
		return *x.MarketSH
	}
	return 0
}

func (x *S2C) GetMarketSZ() int32 {
	if x != nil && x.MarketSZ != nil {
		return *x.MarketSZ
	}
	return 0
}

func (x *S2C) GetMarketHKFuture() int32 {
	if x != nil && x.MarketHKFuture != nil {
		return *x.MarketHKFuture
	}
	return 0
}

func (x *S2C) GetMarketUSFuture() int32 {
	if x != nil && x.MarketUSFuture != nil {
		return *x.MarketUSFuture
	}
	return 0
}

func (x *S2C) GetQotLogined() bool {
	if x != nil && x.QotLogined != nil {
		return *x.QotLogined
	}
	return false
}

func (x *S2C) GetTrdLogined() bool {
	if x != nil && x.TrdLogined != nil {
		return *x.TrdLogined
	}
	return false
}

func (x *S2C) GetServerVer() int32 {
	if x != nil && x.ServerVer != nil {
		return *x.ServerVer
	}
	return 0
}

func (x *S2C) GetServerBuildNo() int32 {
	if x != nil && x.ServerBuildNo != nil {
		return *x.ServerBuildNo
	}
	return 0
}

func (x *S2C) GetTime() int64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *S2C) GetLocalTime() float64 {
	if x != nil && x.LocalTime != nil {
		return *x.LocalTime
	}
	return 0
}

func (x *S2C) GetProgramStatus() *common.ProgramStatus {
	if x != nil {
		return x.ProgramStatus
	}
	return nil
}

func (x *S2C) GetQotSvrIpAddr() string {
	if x != nil && x.QotSvrIpAddr != nil {
		return *x.QotSvrIpAddr
	}
	return ""
}

func (x *S2C) GetTrdSvrIpAddr() string {
	if x != nil && x.TrdSvrIpAddr != nil {
		return *x.TrdSvrIpAddr
	}
	return ""
}

func (x *S2C) GetConnID() uint64 {
	if x != nil && x.ConnID != nil {
		return *x.ConnID
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetGlobalState_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_GetGlobalState_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_GetGlobalState_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetGlobalState_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_GetGlobalState_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_GetGlobalState_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_GetGlobalState_proto protoreflect.FileDescriptor

var file_GetGlobalState_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x98, 0x04, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x4b, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x4b, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x55, 0x53, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x55, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53,
	0x48, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53,
	0x48, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x5a, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x5a, 0x12, 0x26, 0x0a,
	0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x4b, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x4b, 0x46,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55,
	0x53, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x53, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x71, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x02, 0x28,
	0x08, 0x52, 0x0a, 0x71, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x07, 0x20, 0x02, 0x28,
	0x08, 0x52, 0x0a, 0x74, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x18, 0x08, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x6f, 0x18, 0x09, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4e,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x71, 0x6f, 0x74, 0x53, 0x76, 0x72, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x6f, 0x74, 0x53, 0x76, 0x72, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x64, 0x53, 0x76, 0x72, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x64, 0x53,
	0x76, 0x72, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x6e,
	0x49, 0x44, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x44,
	0x22, 0x30, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x63,
	0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63,
	0x32, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05,
	0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x25, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x49, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69,
	0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x74, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x74, 0x65,
}

var (
	file_GetGlobalState_proto_rawDescOnce sync.Once
	file_GetGlobalState_proto_rawDescData = file_GetGlobalState_proto_rawDesc
)

func file_GetGlobalState_proto_rawDescGZIP() []byte {
	file_GetGlobalState_proto_rawDescOnce.Do(func() {
		file_GetGlobalState_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetGlobalState_proto_rawDescData)
	})
	return file_GetGlobalState_proto_rawDescData
}

var file_GetGlobalState_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_GetGlobalState_proto_goTypes = []interface{}{
	(*C2S)(nil),                  // 0: GetGlobalState.C2S
	(*S2C)(nil),                  // 1: GetGlobalState.S2C
	(*Request)(nil),              // 2: GetGlobalState.Request
	(*Response)(nil),             // 3: GetGlobalState.Response
	(*common.ProgramStatus)(nil), // 4: Common.ProgramStatus
}
var file_GetGlobalState_proto_depIdxs = []int32{
	4, // 0: GetGlobalState.S2C.programStatus:type_name -> Common.ProgramStatus
	0, // 1: GetGlobalState.Request.c2s:type_name -> GetGlobalState.C2S
	1, // 2: GetGlobalState.Response.s2c:type_name -> GetGlobalState.S2C
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetGlobalState_proto_init() }
func file_GetGlobalState_proto_init() {
	if File_GetGlobalState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetGlobalState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_GetGlobalState_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_GetGlobalState_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_GetGlobalState_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_GetGlobalState_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetGlobalState_proto_goTypes,
		DependencyIndexes: file_GetGlobalState_proto_depIdxs,
		MessageInfos:      file_GetGlobalState_proto_msgTypes,
	}.Build()
	File_GetGlobalState_proto = out.File
	file_GetGlobalState_proto_rawDesc = nil
	file_GetGlobalState_proto_goTypes = nil
	file_GetGlobalState_proto_depIdxs = nil
}
