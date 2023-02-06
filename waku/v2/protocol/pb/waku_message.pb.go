// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: waku_message.proto

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

type RateLimitProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof         []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	MerkleRoot    []byte `protobuf:"bytes,2,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	Epoch         []byte `protobuf:"bytes,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	ShareX        []byte `protobuf:"bytes,4,opt,name=share_x,json=shareX,proto3" json:"share_x,omitempty"`
	ShareY        []byte `protobuf:"bytes,5,opt,name=share_y,json=shareY,proto3" json:"share_y,omitempty"`
	Nullifier     []byte `protobuf:"bytes,6,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	RlnIdentifier []byte `protobuf:"bytes,7,opt,name=rln_identifier,json=rlnIdentifier,proto3" json:"rln_identifier,omitempty"`
}

func (x *RateLimitProof) Reset() {
	*x = RateLimitProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitProof) ProtoMessage() {}

func (x *RateLimitProof) ProtoReflect() protoreflect.Message {
	mi := &file_waku_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitProof.ProtoReflect.Descriptor instead.
func (*RateLimitProof) Descriptor() ([]byte, []int) {
	return file_waku_message_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitProof) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *RateLimitProof) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *RateLimitProof) GetEpoch() []byte {
	if x != nil {
		return x.Epoch
	}
	return nil
}

func (x *RateLimitProof) GetShareX() []byte {
	if x != nil {
		return x.ShareX
	}
	return nil
}

func (x *RateLimitProof) GetShareY() []byte {
	if x != nil {
		return x.ShareY
	}
	return nil
}

func (x *RateLimitProof) GetNullifier() []byte {
	if x != nil {
		return x.Nullifier
	}
	return nil
}

func (x *RateLimitProof) GetRlnIdentifier() []byte {
	if x != nil {
		return x.RlnIdentifier
	}
	return nil
}

type WakuMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload        []byte          `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ContentTopic   string          `protobuf:"bytes,2,opt,name=contentTopic,proto3" json:"contentTopic,omitempty"`
	Version        uint32          `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp      int64           `protobuf:"zigzag64,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RateLimitProof *RateLimitProof `protobuf:"bytes,21,opt,name=rate_limit_proof,json=rateLimitProof,proto3" json:"rate_limit_proof,omitempty"`
	Ephemeral      bool            `protobuf:"varint,31,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
}

func (x *WakuMessage) Reset() {
	*x = WakuMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WakuMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WakuMessage) ProtoMessage() {}

func (x *WakuMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waku_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WakuMessage.ProtoReflect.Descriptor instead.
func (*WakuMessage) Descriptor() ([]byte, []int) {
	return file_waku_message_proto_rawDescGZIP(), []int{1}
}

func (x *WakuMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *WakuMessage) GetContentTopic() string {
	if x != nil {
		return x.ContentTopic
	}
	return ""
}

func (x *WakuMessage) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WakuMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *WakuMessage) GetRateLimitProof() *RateLimitProof {
	if x != nil {
		return x.RateLimitProof
	}
	return nil
}

func (x *WakuMessage) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

var File_waku_message_proto protoreflect.FileDescriptor

var file_waku_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x77, 0x61, 0x6b, 0x75, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xd4, 0x01, 0x0a, 0x0e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x5f, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x58, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x59, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75,
	0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e,
	0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x6c, 0x6e, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x72, 0x6c, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22,
	0xdf, 0x01, 0x0a, 0x0b, 0x57, 0x61, 0x6b, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x10, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x52, 0x0e, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waku_message_proto_rawDescOnce sync.Once
	file_waku_message_proto_rawDescData = file_waku_message_proto_rawDesc
)

func file_waku_message_proto_rawDescGZIP() []byte {
	file_waku_message_proto_rawDescOnce.Do(func() {
		file_waku_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_waku_message_proto_rawDescData)
	})
	return file_waku_message_proto_rawDescData
}

var file_waku_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waku_message_proto_goTypes = []interface{}{
	(*RateLimitProof)(nil), // 0: pb.RateLimitProof
	(*WakuMessage)(nil),    // 1: pb.WakuMessage
}
var file_waku_message_proto_depIdxs = []int32{
	0, // 0: pb.WakuMessage.rate_limit_proof:type_name -> pb.RateLimitProof
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_waku_message_proto_init() }
func file_waku_message_proto_init() {
	if File_waku_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waku_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitProof); i {
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
		file_waku_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WakuMessage); i {
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
			RawDescriptor: file_waku_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waku_message_proto_goTypes,
		DependencyIndexes: file_waku_message_proto_depIdxs,
		MessageInfos:      file_waku_message_proto_msgTypes,
	}.Build()
	File_waku_message_proto = out.File
	file_waku_message_proto_rawDesc = nil
	file_waku_message_proto_goTypes = nil
	file_waku_message_proto_depIdxs = nil
}
