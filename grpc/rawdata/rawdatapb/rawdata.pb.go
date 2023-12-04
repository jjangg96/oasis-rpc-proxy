// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: grpc/rawdata/rawdatapb/rawdata.proto

package rawdatapb

import (
	blockpb "github.com/jjangg96/oasis-rpc-proxy/grpc/block/blockpb"
	eventpb "github.com/jjangg96/oasis-rpc-proxy/grpc/event/eventpb"
	statepb "github.com/jjangg96/oasis-rpc-proxy/grpc/state/statepb"
	transactionpb "github.com/jjangg96/oasis-rpc-proxy/grpc/transaction/transactionpb"
	validatorpb "github.com/jjangg96/oasis-rpc-proxy/grpc/validator/validatorpb"
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

type RawData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block          *blockpb.Block               `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	EscrowEvents   *eventpb.EscrowEvents        `protobuf:"bytes,2,opt,name=escrow_events,json=escrowEvents,proto3" json:"escrow_events,omitempty"`
	Staking        *statepb.Staking             `protobuf:"bytes,3,opt,name=staking,proto3" json:"staking,omitempty"`
	State          *statepb.State               `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	TransferEvents []*eventpb.TransferEvent     `protobuf:"bytes,5,rep,name=transfer_events,json=transferEvents,proto3" json:"transfer_events,omitempty"`
	Transactions   []*transactionpb.Transaction `protobuf:"bytes,6,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Validators     []*validatorpb.Validator     `protobuf:"bytes,7,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (x *RawData) Reset() {
	*x = RawData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rawdata_rawdatapb_rawdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawData) ProtoMessage() {}

func (x *RawData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rawdata_rawdatapb_rawdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawData.ProtoReflect.Descriptor instead.
func (*RawData) Descriptor() ([]byte, []int) {
	return file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescGZIP(), []int{0}
}

func (x *RawData) GetBlock() *blockpb.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *RawData) GetEscrowEvents() *eventpb.EscrowEvents {
	if x != nil {
		return x.EscrowEvents
	}
	return nil
}

func (x *RawData) GetStaking() *statepb.Staking {
	if x != nil {
		return x.Staking
	}
	return nil
}

func (x *RawData) GetState() *statepb.State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *RawData) GetTransferEvents() []*eventpb.TransferEvent {
	if x != nil {
		return x.TransferEvents
	}
	return nil
}

func (x *RawData) GetTransactions() []*transactionpb.Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *RawData) GetValidators() []*validatorpb.Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

var File_grpc_rawdata_rawdatapb_rawdata_proto protoreflect.FileDescriptor

var file_grpc_rawdata_rawdatapb_rawdata_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x30, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02,
	0x0a, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x38, 0x0a,
	0x0d, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x73, 0x63,
	0x72, 0x6f, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x65, 0x73, 0x63, 0x72, 0x6f,
	0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x6f, 0x61, 0x73, 0x69, 0x73, 0x2d, 0x72, 0x70,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x61, 0x77,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescOnce sync.Once
	file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescData = file_grpc_rawdata_rawdatapb_rawdata_proto_rawDesc
)

func file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescGZIP() []byte {
	file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescOnce.Do(func() {
		file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescData)
	})
	return file_grpc_rawdata_rawdatapb_rawdata_proto_rawDescData
}

var file_grpc_rawdata_rawdatapb_rawdata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_rawdata_rawdatapb_rawdata_proto_goTypes = []interface{}{
	(*RawData)(nil),                   // 0: rawdata.RawData
	(*blockpb.Block)(nil),             // 1: block.Block
	(*eventpb.EscrowEvents)(nil),      // 2: event.EscrowEvents
	(*statepb.Staking)(nil),           // 3: state.Staking
	(*statepb.State)(nil),             // 4: state.State
	(*eventpb.TransferEvent)(nil),     // 5: event.TransferEvent
	(*transactionpb.Transaction)(nil), // 6: transaction.Transaction
	(*validatorpb.Validator)(nil),     // 7: validator.Validator
}
var file_grpc_rawdata_rawdatapb_rawdata_proto_depIdxs = []int32{
	1, // 0: rawdata.RawData.block:type_name -> block.Block
	2, // 1: rawdata.RawData.escrow_events:type_name -> event.EscrowEvents
	3, // 2: rawdata.RawData.staking:type_name -> state.Staking
	4, // 3: rawdata.RawData.state:type_name -> state.State
	5, // 4: rawdata.RawData.transfer_events:type_name -> event.TransferEvent
	6, // 5: rawdata.RawData.transactions:type_name -> transaction.Transaction
	7, // 6: rawdata.RawData.validators:type_name -> validator.Validator
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_grpc_rawdata_rawdatapb_rawdata_proto_init() }
func file_grpc_rawdata_rawdatapb_rawdata_proto_init() {
	if File_grpc_rawdata_rawdatapb_rawdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_rawdata_rawdatapb_rawdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawData); i {
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
			RawDescriptor: file_grpc_rawdata_rawdatapb_rawdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_rawdata_rawdatapb_rawdata_proto_goTypes,
		DependencyIndexes: file_grpc_rawdata_rawdatapb_rawdata_proto_depIdxs,
		MessageInfos:      file_grpc_rawdata_rawdatapb_rawdata_proto_msgTypes,
	}.Build()
	File_grpc_rawdata_rawdatapb_rawdata_proto = out.File
	file_grpc_rawdata_rawdatapb_rawdata_proto_rawDesc = nil
	file_grpc_rawdata_rawdatapb_rawdata_proto_goTypes = nil
	file_grpc_rawdata_rawdatapb_rawdata_proto_depIdxs = nil
}
