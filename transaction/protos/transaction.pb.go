// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: protos/transaction.proto

package protos

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

type TransactionDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	Amount  int32  `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *TransactionDetails) Reset() {
	*x = TransactionDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionDetails) ProtoMessage() {}

func (x *TransactionDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionDetails.ProtoReflect.Descriptor instead.
func (*TransactionDetails) Descriptor() ([]byte, []int) {
	return file_protos_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionDetails) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TransactionDetails) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *TransactionDetails) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddPointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPointsResponse) Reset() {
	*x = AddPointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPointsResponse) ProtoMessage() {}

func (x *AddPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPointsResponse.ProtoReflect.Descriptor instead.
func (*AddPointsResponse) Descriptor() ([]byte, []int) {
	return file_protos_transaction_proto_rawDescGZIP(), []int{1}
}

type UsePointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionPointsUsed bool                `protobuf:"varint,1,opt,name=transactionPointsUsed,proto3" json:"transactionPointsUsed,omitempty"`
	TransactionDetails    *TransactionDetails `protobuf:"bytes,2,opt,name=transactionDetails,proto3" json:"transactionDetails,omitempty"`
}

func (x *UsePointsResponse) Reset() {
	*x = UsePointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsePointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsePointsResponse) ProtoMessage() {}

func (x *UsePointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsePointsResponse.ProtoReflect.Descriptor instead.
func (*UsePointsResponse) Descriptor() ([]byte, []int) {
	return file_protos_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *UsePointsResponse) GetTransactionPointsUsed() bool {
	if x != nil {
		return x.TransactionPointsUsed
	}
	return false
}

func (x *UsePointsResponse) GetTransactionDetails() *TransactionDetails {
	if x != nil {
		return x.TransactionDetails
	}
	return nil
}

var File_protos_transaction_proto protoreflect.FileDescriptor

var file_protos_transaction_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x22, 0x5e, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x55,
	0x73, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x12, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32,
	0xb1, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4d, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_transaction_proto_rawDescOnce sync.Once
	file_protos_transaction_proto_rawDescData = file_protos_transaction_proto_rawDesc
)

func file_protos_transaction_proto_rawDescGZIP() []byte {
	file_protos_transaction_proto_rawDescOnce.Do(func() {
		file_protos_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_transaction_proto_rawDescData)
	})
	return file_protos_transaction_proto_rawDescData
}

var file_protos_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_transaction_proto_goTypes = []interface{}{
	(*TransactionDetails)(nil), // 0: protos.TransactionDetails
	(*AddPointsResponse)(nil),  // 1: protos.AddPointsResponse
	(*UsePointsResponse)(nil),  // 2: protos.UsePointsResponse
}
var file_protos_transaction_proto_depIdxs = []int32{
	0, // 0: protos.UsePointsResponse.transactionDetails:type_name -> protos.TransactionDetails
	0, // 1: protos.TransactionPoints.AddTransactionPoints:input_type -> protos.TransactionDetails
	0, // 2: protos.TransactionPoints.UseTransactionPoints:input_type -> protos.TransactionDetails
	1, // 3: protos.TransactionPoints.AddTransactionPoints:output_type -> protos.AddPointsResponse
	2, // 4: protos.TransactionPoints.UseTransactionPoints:output_type -> protos.UsePointsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_transaction_proto_init() }
func file_protos_transaction_proto_init() {
	if File_protos_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionDetails); i {
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
		file_protos_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPointsResponse); i {
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
		file_protos_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsePointsResponse); i {
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
			RawDescriptor: file_protos_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_transaction_proto_goTypes,
		DependencyIndexes: file_protos_transaction_proto_depIdxs,
		MessageInfos:      file_protos_transaction_proto_msgTypes,
	}.Build()
	File_protos_transaction_proto = out.File
	file_protos_transaction_proto_rawDesc = nil
	file_protos_transaction_proto_goTypes = nil
	file_protos_transaction_proto_depIdxs = nil
}
