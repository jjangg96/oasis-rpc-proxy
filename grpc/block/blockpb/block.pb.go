// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/block/blockpb/block.proto

package blockpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Block struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	LastCommit           *Commit  `protobuf:"bytes,2,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetLastCommit() *Commit {
	if m != nil {
		return m.LastCommit
	}
	return nil
}

type Header struct {
	Version *Version             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ChainId string               `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height  int64                `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	// prev block info
	LastBlockId *BlockID `protobuf:"bytes,7,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id,omitempty"`
	// hashes of block data
	LastCommitHash string `protobuf:"bytes,8,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	DataHash       string `protobuf:"bytes,9,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	// hashes from the app output from the prev block
	ValidatorsHash     string `protobuf:"bytes,10,opt,name=validators_hash,json=validatorsHash,proto3" json:"validators_hash,omitempty"`
	NextValidatorsHash string `protobuf:"bytes,11,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"next_validators_hash,omitempty"`
	ConsensusHash      string `protobuf:"bytes,12,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`
	AppHash            string `protobuf:"bytes,13,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	// root hash of all results from the txs from the previous block
	LastResultsHash string `protobuf:"bytes,14,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// consensus info
	EvidenceHash         string   `protobuf:"bytes,15,opt,name=evidence_hash,json=evidenceHash,proto3" json:"evidence_hash,omitempty"`
	ProposerAddress      string   `protobuf:"bytes,16,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{1}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Header) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Header) GetLastBlockId() *BlockID {
	if m != nil {
		return m.LastBlockId
	}
	return nil
}

func (m *Header) GetLastCommitHash() string {
	if m != nil {
		return m.LastCommitHash
	}
	return ""
}

func (m *Header) GetDataHash() string {
	if m != nil {
		return m.DataHash
	}
	return ""
}

func (m *Header) GetValidatorsHash() string {
	if m != nil {
		return m.ValidatorsHash
	}
	return ""
}

func (m *Header) GetNextValidatorsHash() string {
	if m != nil {
		return m.NextValidatorsHash
	}
	return ""
}

func (m *Header) GetConsensusHash() string {
	if m != nil {
		return m.ConsensusHash
	}
	return ""
}

func (m *Header) GetAppHash() string {
	if m != nil {
		return m.AppHash
	}
	return ""
}

func (m *Header) GetLastResultsHash() string {
	if m != nil {
		return m.LastResultsHash
	}
	return ""
}

func (m *Header) GetEvidenceHash() string {
	if m != nil {
		return m.EvidenceHash
	}
	return ""
}

func (m *Header) GetProposerAddress() string {
	if m != nil {
		return m.ProposerAddress
	}
	return ""
}

type Version struct {
	Block                uint64   `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	App                  uint64   `protobuf:"varint,2,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{2}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Version) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type BlockID struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{3}
}

func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Commit struct {
	BlockId              *BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Height               int64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round                int64    `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	Hash                 string   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Votes                []*Vote  `protobuf:"bytes,5,rep,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{4}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetBlockId() *BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *Commit) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Commit) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Commit) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Commit) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

// Vote represents a prevote, precommit, or commit vote from validators for consensus.
type Vote struct {
	BlockIdFlag          int64    `protobuf:"varint,1,opt,name=block_id_flag,json=blockIdFlag,proto3" json:"block_id_flag,omitempty"`
	ValidatorAddress     string   `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	ValidatorIndex       int64    `protobuf:"varint,3,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Signature            string   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{5}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetBlockIdFlag() int64 {
	if m != nil {
		return m.BlockIdFlag
	}
	return 0
}

func (m *Vote) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *Vote) GetValidatorIndex() int64 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

func (m *Vote) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type GetByHeightRequest struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByHeightRequest) Reset()         { *m = GetByHeightRequest{} }
func (m *GetByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetByHeightRequest) ProtoMessage()    {}
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{6}
}

func (m *GetByHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByHeightRequest.Unmarshal(m, b)
}
func (m *GetByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightRequest.Merge(m, src)
}
func (m *GetByHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetByHeightRequest.Size(m)
}
func (m *GetByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightRequest proto.InternalMessageInfo

func (m *GetByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetByHeightResponse struct {
	Block                *Block   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByHeightResponse) Reset()         { *m = GetByHeightResponse{} }
func (m *GetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetByHeightResponse) ProtoMessage()    {}
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa3e78aae19c108, []int{7}
}

func (m *GetByHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByHeightResponse.Unmarshal(m, b)
}
func (m *GetByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByHeightResponse.Marshal(b, m, deterministic)
}
func (m *GetByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightResponse.Merge(m, src)
}
func (m *GetByHeightResponse) XXX_Size() int {
	return xxx_messageInfo_GetByHeightResponse.Size(m)
}
func (m *GetByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightResponse proto.InternalMessageInfo

func (m *GetByHeightResponse) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "block.Block")
	proto.RegisterType((*Header)(nil), "block.Header")
	proto.RegisterType((*Version)(nil), "block.Version")
	proto.RegisterType((*BlockID)(nil), "block.BlockID")
	proto.RegisterType((*Commit)(nil), "block.Commit")
	proto.RegisterType((*Vote)(nil), "block.Vote")
	proto.RegisterType((*GetByHeightRequest)(nil), "block.GetByHeightRequest")
	proto.RegisterType((*GetByHeightResponse)(nil), "block.GetByHeightResponse")
}

func init() { proto.RegisterFile("grpc/block/blockpb/block.proto", fileDescriptor_6fa3e78aae19c108) }

var fileDescriptor_6fa3e78aae19c108 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x8d, 0xf3, 0x35, 0x4e, 0xd2, 0x74, 0x89, 0x90, 0x1b, 0xbe, 0x8a, 0x51, 0x45, 0x0a,
	0xc8, 0x85, 0x70, 0xe2, 0x48, 0x41, 0xa5, 0xbd, 0x2e, 0xa8, 0x07, 0x0e, 0x58, 0x1b, 0x7b, 0xeb,
	0x58, 0x38, 0xde, 0xc5, 0xbb, 0x89, 0xca, 0x3f, 0xe1, 0xc4, 0x9d, 0x7f, 0x89, 0x3c, 0xbb, 0xce,
	0x07, 0xed, 0xa5, 0xf5, 0xbc, 0x79, 0x9e, 0x79, 0x33, 0xf3, 0x62, 0x78, 0x92, 0x96, 0x32, 0x3e,
	0x9d, 0xe5, 0x22, 0xfe, 0x61, 0xfe, 0xca, 0x99, 0xf9, 0x1f, 0xca, 0x52, 0x68, 0x41, 0x9a, 0x18,
	0x8c, 0x9f, 0xa6, 0x42, 0xa4, 0x39, 0x3f, 0x45, 0x70, 0xb6, 0xbc, 0x3e, 0xd5, 0xd9, 0x82, 0x2b,
	0xcd, 0x16, 0xd2, 0xf0, 0x82, 0xef, 0xd0, 0x3c, 0xab, 0x98, 0xe4, 0x18, 0x5a, 0x73, 0xce, 0x12,
	0x5e, 0xfa, 0xce, 0x91, 0x33, 0xf1, 0xa6, 0xfd, 0xd0, 0x94, 0xbb, 0x40, 0x90, 0xda, 0x24, 0x09,
	0xc1, 0xcb, 0x99, 0xd2, 0x51, 0x2c, 0x16, 0x8b, 0x4c, 0xfb, 0x7b, 0x3b, 0xdc, 0x8f, 0x08, 0x52,
	0xa8, 0x18, 0xe6, 0x39, 0xf8, 0xeb, 0x42, 0xcb, 0x94, 0x20, 0x13, 0x68, 0xaf, 0x78, 0xa9, 0x32,
	0x51, 0xd8, 0x16, 0x03, 0xfb, 0xda, 0x95, 0x41, 0x69, 0x9d, 0x26, 0x87, 0xd0, 0x89, 0xe7, 0x2c,
	0x2b, 0xa2, 0x2c, 0xc1, 0x0e, 0x5d, 0xda, 0xc6, 0xf8, 0x32, 0x21, 0x0f, 0x2a, 0x99, 0x59, 0x3a,
	0xd7, 0x7e, 0xe3, 0xc8, 0x99, 0x34, 0xa8, 0x8d, 0x48, 0x08, 0x6e, 0x35, 0x9a, 0xef, 0x62, 0xe5,
	0x71, 0x68, 0xe6, 0x0e, 0xeb, 0xb9, 0xc3, 0xaf, 0xf5, 0xdc, 0x14, 0x79, 0x64, 0x0a, 0x7d, 0x9c,
	0x03, 0x15, 0x54, 0x7d, 0xda, 0x3b, 0x92, 0x70, 0x27, 0x97, 0x9f, 0x28, 0x0e, 0x6b, 0x82, 0x84,
	0x4c, 0x60, 0xb8, 0x35, 0x7b, 0x34, 0x67, 0x6a, 0xee, 0x77, 0x50, 0xde, 0x60, 0x33, 0xf1, 0x05,
	0x53, 0x73, 0xf2, 0x10, 0xba, 0x09, 0xd3, 0xcc, 0x50, 0xba, 0x48, 0xe9, 0x54, 0x00, 0x26, 0x5f,
	0xc0, 0xfe, 0x8a, 0xe5, 0x59, 0xc2, 0xb4, 0x28, 0x95, 0xa1, 0x80, 0xa9, 0xb2, 0x81, 0x91, 0xf8,
	0x06, 0x46, 0x05, 0xbf, 0xd1, 0xd1, 0xff, 0x6c, 0x0f, 0xd9, 0xa4, 0xca, 0x5d, 0xed, 0xbe, 0x71,
	0x0c, 0x83, 0x58, 0x14, 0x8a, 0x17, 0x6a, 0x69, 0xb9, 0x3d, 0xe4, 0xf6, 0xd7, 0x28, 0xd2, 0x0e,
	0xa1, 0xc3, 0xa4, 0x34, 0x84, 0xbe, 0xd9, 0x2f, 0x93, 0x12, 0x53, 0x2f, 0xe1, 0x00, 0x67, 0x2c,
	0xb9, 0x5a, 0xe6, 0xda, 0x16, 0x19, 0x20, 0x67, 0xbf, 0x4a, 0x50, 0x83, 0x23, 0xf7, 0x39, 0xf4,
	0xf9, 0x2a, 0x4b, 0x78, 0x11, 0x73, 0xc3, 0xdb, 0x47, 0x5e, 0xaf, 0x06, 0x91, 0x74, 0x02, 0x43,
	0x59, 0x0a, 0x29, 0x14, 0x2f, 0x23, 0x96, 0x24, 0x25, 0x57, 0xca, 0x1f, 0x9a, 0x7a, 0x35, 0xfe,
	0xc1, 0xc0, 0xc1, 0x5b, 0x68, 0x5b, 0x2b, 0x90, 0x11, 0x18, 0x03, 0xa3, 0x53, 0x5c, 0x6a, 0x02,
	0x32, 0x84, 0x06, 0x93, 0x12, 0x2d, 0xe1, 0xd2, 0xea, 0x31, 0x78, 0x0c, 0x6d, 0x7b, 0x2a, 0x42,
	0xc0, 0x45, 0x11, 0x0e, 0x16, 0xc7, 0xe7, 0xe0, 0xb7, 0x03, 0x2d, 0x73, 0x16, 0x72, 0x02, 0x9d,
	0xf5, 0xad, 0x9d, 0x3b, 0x6f, 0xdd, 0x9e, 0xd9, 0x3b, 0x6f, 0x3c, 0xb6, 0xb7, 0xe3, 0xb1, 0x11,
	0x34, 0x4b, 0xb1, 0x2c, 0x12, 0x6b, 0x3d, 0x13, 0xac, 0xfb, 0xba, 0x9b, 0xbe, 0xe4, 0x19, 0x34,
	0x57, 0x42, 0x73, 0xe5, 0x37, 0x8f, 0x1a, 0x13, 0x6f, 0xea, 0xd5, 0x46, 0x17, 0x9a, 0x53, 0x93,
	0x09, 0xfe, 0x38, 0xe0, 0x56, 0x31, 0x09, 0xa0, 0x5f, 0x0b, 0x8b, 0xae, 0x73, 0x96, 0xa2, 0xba,
	0x06, 0xf5, 0xac, 0x9a, 0xf3, 0x9c, 0xa5, 0xe4, 0x15, 0x1c, 0xac, 0x4d, 0xb0, 0xde, 0xa2, 0xf9,
	0x65, 0x0c, 0xd7, 0x09, 0xbb, 0xc6, 0x1d, 0x7f, 0x45, 0x59, 0x91, 0xf0, 0x1b, 0x2b, 0x78, 0xe3,
	0xaf, 0xcb, 0x0a, 0x25, 0x8f, 0xa0, 0xab, 0xb2, 0xb4, 0x60, 0x7a, 0x59, 0x72, 0x2b, 0x7f, 0x03,
	0x04, 0xaf, 0x81, 0x7c, 0xe6, 0xfa, 0xec, 0xd7, 0x05, 0x0e, 0x4f, 0xf9, 0xcf, 0x25, 0x57, 0x7a,
	0x6b, 0x37, 0xce, 0xf6, 0x6e, 0x82, 0xf7, 0x70, 0x7f, 0x87, 0xad, 0x64, 0xe5, 0x38, 0x12, 0x6c,
	0xdf, 0xd1, 0x9b, 0xf6, 0xb6, 0x57, 0x6e, 0xaf, 0x3a, 0xbd, 0x82, 0x1e, 0xc6, 0x5f, 0x78, 0xb9,
	0xca, 0x62, 0x4e, 0xce, 0xc1, 0xdb, 0x2a, 0x45, 0x0e, 0xed, 0x3b, 0xb7, 0xc5, 0x8c, 0xc7, 0x77,
	0xa5, 0x4c, 0xe7, 0xe0, 0xde, 0xd9, 0xe8, 0x1b, 0xb9, 0xfd, 0x91, 0x9c, 0xb5, 0xf0, 0x93, 0xf0,
	0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xf3, 0x52, 0xf4, 0x41, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockServiceClient interface {
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
}

type blockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockServiceClient(cc grpc.ClientConnInterface) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/block.BlockService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServiceServer is the server API for BlockService service.
type BlockServiceServer interface {
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
}

// UnimplementedBlockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (*UnimplementedBlockServiceServer) GetByHeight(ctx context.Context, req *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterBlockServiceServer(s *grpc.Server, srv BlockServiceServer) {
	s.RegisterService(&_BlockService_serviceDesc, srv)
}

func _BlockService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/block.BlockService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "block.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _BlockService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/block/blockpb/block.proto",
}
