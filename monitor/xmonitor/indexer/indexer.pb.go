// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2-devel
// 	protoc        (unknown)
// source: monitor/xmonitor/indexer/indexer.proto

package indexer

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      // Auto-incremented ID
	ChainId     uint64 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`             // Source chain ID as per https://chainlist.org
	BlockHeight uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"` // Height of the source-chain block
	BlockHash   []byte `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`        // Hash of the source-chain block
	BlockJson   []byte `protobuf:"bytes,5,opt,name=block_json,json=blockJson,proto3" json:"block_json,omitempty"`        // xchain.Block JSON
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Block) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Block) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Block) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *Block) GetBlockJson() []byte {
	if x != nil {
		return x.BlockJson
	}
	return nil
}

type MsgLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdHash         []byte `protobuf:"bytes,1,opt,name=id_hash,json=idHash,proto3" json:"id_hash,omitempty"` // RouteScan IDHash of the MsgID
	MsgBlockId     uint64 `protobuf:"varint,2,opt,name=msg_block_id,json=msgBlockId,proto3" json:"msg_block_id,omitempty"`
	ReceiptBlockId uint64 `protobuf:"varint,3,opt,name=receipt_block_id,json=receiptBlockId,proto3" json:"receipt_block_id,omitempty"`
}

func (x *MsgLink) Reset() {
	*x = MsgLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgLink) ProtoMessage() {}

func (x *MsgLink) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgLink.ProtoReflect.Descriptor instead.
func (*MsgLink) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *MsgLink) GetIdHash() []byte {
	if x != nil {
		return x.IdHash
	}
	return nil
}

func (x *MsgLink) GetMsgBlockId() uint64 {
	if x != nil {
		return x.MsgBlockId
	}
	return 0
}

func (x *MsgLink) GetReceiptBlockId() uint64 {
	if x != nil {
		return x.ReceiptBlockId
	}
	return 0
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId     uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConfLevel   uint32 `protobuf:"varint,2,opt,name=conf_level,json=confLevel,proto3" json:"conf_level,omitempty"`
	BlockHeight uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_indexer_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *Cursor) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Cursor) GetConfLevel() uint32 {
	if x != nil {
		return x.ConfLevel
	}
	return 0
}

func (x *Cursor) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

var File_monitor_xmonitor_indexer_indexer_proto protoreflect.FileDescriptor

var file_monitor_xmonitor_indexer_indexer_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x78, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x78, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x72, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x05,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6a, 0x73, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4a, 0x73, 0x6f,
	0x6e, 0x3a, 0x38, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x32, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10,
	0x01, 0x12, 0x26, 0x0a, 0x20, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x10, 0x02, 0x18, 0x01, 0x18, 0x01, 0x22, 0x83, 0x01, 0x0a, 0x07,
	0x4d, 0x73, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x73, 0x67, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x3a, 0x13, 0xf2, 0x9e,
	0xd3, 0x8e, 0x03, 0x0d, 0x0a, 0x09, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x22, 0x86, 0x01, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x1f, 0xf2, 0x9e, 0xd3, 0x8e, 0x03,
	0x19, 0x0a, 0x15, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x63, 0x6f,
	0x6e, 0x66, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x42, 0xe5, 0x01, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x78, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x42, 0x0c, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x2f, 0x78, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x72, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x49, 0xaa, 0x02, 0x18, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x58, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0xca, 0x02, 0x18, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x58, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0xe2, 0x02,
	0x24, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x58, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x5c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x3a,
	0x3a, 0x58, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_xmonitor_indexer_indexer_proto_rawDescOnce sync.Once
	file_monitor_xmonitor_indexer_indexer_proto_rawDescData = file_monitor_xmonitor_indexer_indexer_proto_rawDesc
)

func file_monitor_xmonitor_indexer_indexer_proto_rawDescGZIP() []byte {
	file_monitor_xmonitor_indexer_indexer_proto_rawDescOnce.Do(func() {
		file_monitor_xmonitor_indexer_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_xmonitor_indexer_indexer_proto_rawDescData)
	})
	return file_monitor_xmonitor_indexer_indexer_proto_rawDescData
}

var file_monitor_xmonitor_indexer_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_monitor_xmonitor_indexer_indexer_proto_goTypes = []any{
	(*Block)(nil),   // 0: monitor.xmonitor.indexer.Block
	(*MsgLink)(nil), // 1: monitor.xmonitor.indexer.MsgLink
	(*Cursor)(nil),  // 2: monitor.xmonitor.indexer.Cursor
}
var file_monitor_xmonitor_indexer_indexer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitor_xmonitor_indexer_indexer_proto_init() }
func file_monitor_xmonitor_indexer_indexer_proto_init() {
	if File_monitor_xmonitor_indexer_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_xmonitor_indexer_indexer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Block); i {
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
		file_monitor_xmonitor_indexer_indexer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MsgLink); i {
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
		file_monitor_xmonitor_indexer_indexer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Cursor); i {
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
			RawDescriptor: file_monitor_xmonitor_indexer_indexer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_monitor_xmonitor_indexer_indexer_proto_goTypes,
		DependencyIndexes: file_monitor_xmonitor_indexer_indexer_proto_depIdxs,
		MessageInfos:      file_monitor_xmonitor_indexer_indexer_proto_msgTypes,
	}.Build()
	File_monitor_xmonitor_indexer_indexer_proto = out.File
	file_monitor_xmonitor_indexer_indexer_proto_rawDesc = nil
	file_monitor_xmonitor_indexer_indexer_proto_goTypes = nil
	file_monitor_xmonitor_indexer_indexer_proto_depIdxs = nil
}
