// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: state.proto

package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              []byte   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce                uint64   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Staking              []byte   `protobuf:"bytes,11,opt,name=staking,proto3" json:"staking,omitempty"`
	VotedRootHash        []byte   `protobuf:"bytes,12,opt,name=voted_root_hash,json=votedRootHash,proto3" json:"voted_root_hash,omitempty"`
	CandidateId          []byte   `protobuf:"bytes,13,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Bandwidth            []byte   `protobuf:"bytes,14,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	LastBandwidthTs      int64    `protobuf:"varint,15,opt,name=last_bandwidth_ts,json=lastBandwidthTs,proto3" json:"last_bandwidth_ts,omitempty"`
	Unstaking            []byte   `protobuf:"bytes,16,opt,name=unstaking,proto3" json:"unstaking,omitempty"`
	LastUnstakingTs      int64    `protobuf:"varint,17,opt,name=last_unstaking_ts,json=lastUnstakingTs,proto3" json:"last_unstaking_ts,omitempty"`
	DataRootHash         []byte   `protobuf:"bytes,40,opt,name=data_root_hash,json=dataRootHash,proto3" json:"data_root_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_b735a4e0a688517c, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Account) GetStaking() []byte {
	if m != nil {
		return m.Staking
	}
	return nil
}

func (m *Account) GetVotedRootHash() []byte {
	if m != nil {
		return m.VotedRootHash
	}
	return nil
}

func (m *Account) GetCandidateId() []byte {
	if m != nil {
		return m.CandidateId
	}
	return nil
}

func (m *Account) GetBandwidth() []byte {
	if m != nil {
		return m.Bandwidth
	}
	return nil
}

func (m *Account) GetLastBandwidthTs() int64 {
	if m != nil {
		return m.LastBandwidthTs
	}
	return 0
}

func (m *Account) GetUnstaking() []byte {
	if m != nil {
		return m.Unstaking
	}
	return nil
}

func (m *Account) GetLastUnstakingTs() int64 {
	if m != nil {
		return m.LastUnstakingTs
	}
	return 0
}

func (m *Account) GetDataRootHash() []byte {
	if m != nil {
		return m.DataRootHash
	}
	return nil
}

type AliasAccount struct {
	Account              []byte   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Alias                string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliasAccount) Reset()         { *m = AliasAccount{} }
func (m *AliasAccount) String() string { return proto.CompactTextString(m) }
func (*AliasAccount) ProtoMessage()    {}
func (*AliasAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_b735a4e0a688517c, []int{1}
}
func (m *AliasAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliasAccount.Unmarshal(m, b)
}
func (m *AliasAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliasAccount.Marshal(b, m, deterministic)
}
func (dst *AliasAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliasAccount.Merge(dst, src)
}
func (m *AliasAccount) XXX_Size() int {
	return xxx_messageInfo_AliasAccount.Size(m)
}
func (m *AliasAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AliasAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AliasAccount proto.InternalMessageInfo

func (m *AliasAccount) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *AliasAccount) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

type DataState struct {
	TxStateRootHash            []byte   `protobuf:"bytes,1,opt,name=tx_state_root_hash,json=txStateRootHash,proto3" json:"tx_state_root_hash,omitempty"`
	RecordStateRootHash        []byte   `protobuf:"bytes,2,opt,name=record_state_root_hash,json=recordStateRootHash,proto3" json:"record_state_root_hash,omitempty"`
	CertificationStateRootHash []byte   `protobuf:"bytes,3,opt,name=certification_state_root_hash,json=certificationStateRootHash,proto3" json:"certification_state_root_hash,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *DataState) Reset()         { *m = DataState{} }
func (m *DataState) String() string { return proto.CompactTextString(m) }
func (*DataState) ProtoMessage()    {}
func (*DataState) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_b735a4e0a688517c, []int{2}
}
func (m *DataState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataState.Unmarshal(m, b)
}
func (m *DataState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataState.Marshal(b, m, deterministic)
}
func (dst *DataState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataState.Merge(dst, src)
}
func (m *DataState) XXX_Size() int {
	return xxx_messageInfo_DataState.Size(m)
}
func (m *DataState) XXX_DiscardUnknown() {
	xxx_messageInfo_DataState.DiscardUnknown(m)
}

var xxx_messageInfo_DataState proto.InternalMessageInfo

func (m *DataState) GetTxStateRootHash() []byte {
	if m != nil {
		return m.TxStateRootHash
	}
	return nil
}

func (m *DataState) GetRecordStateRootHash() []byte {
	if m != nil {
		return m.RecordStateRootHash
	}
	return nil
}

func (m *DataState) GetCertificationStateRootHash() []byte {
	if m != nil {
		return m.CertificationStateRootHash
	}
	return nil
}

type Record struct {
	RecordHash           []byte   `protobuf:"bytes,1,opt,name=record_hash,json=recordHash,proto3" json:"record_hash,omitempty"`
	Owner                []byte   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_b735a4e0a688517c, []int{3}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (dst *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(dst, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetRecordHash() []byte {
	if m != nil {
		return m.RecordHash
	}
	return nil
}

func (m *Record) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Certification struct {
	CertificateHash      []byte   `protobuf:"bytes,1,opt,name=certificate_hash,json=certificateHash,proto3" json:"certificate_hash,omitempty"`
	Issuer               []byte   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Certified            []byte   `protobuf:"bytes,3,opt,name=certified,proto3" json:"certified,omitempty"`
	IssueTime            int64    `protobuf:"varint,4,opt,name=issue_time,json=issueTime,proto3" json:"issue_time,omitempty"`
	ExpirationTime       int64    `protobuf:"varint,5,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	RevocationTime       int64    `protobuf:"varint,6,opt,name=revocation_time,json=revocationTime,proto3" json:"revocation_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certification) Reset()         { *m = Certification{} }
func (m *Certification) String() string { return proto.CompactTextString(m) }
func (*Certification) ProtoMessage()    {}
func (*Certification) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_b735a4e0a688517c, []int{4}
}
func (m *Certification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certification.Unmarshal(m, b)
}
func (m *Certification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certification.Marshal(b, m, deterministic)
}
func (dst *Certification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certification.Merge(dst, src)
}
func (m *Certification) XXX_Size() int {
	return xxx_messageInfo_Certification.Size(m)
}
func (m *Certification) XXX_DiscardUnknown() {
	xxx_messageInfo_Certification.DiscardUnknown(m)
}

var xxx_messageInfo_Certification proto.InternalMessageInfo

func (m *Certification) GetCertificateHash() []byte {
	if m != nil {
		return m.CertificateHash
	}
	return nil
}

func (m *Certification) GetIssuer() []byte {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *Certification) GetCertified() []byte {
	if m != nil {
		return m.Certified
	}
	return nil
}

func (m *Certification) GetIssueTime() int64 {
	if m != nil {
		return m.IssueTime
	}
	return 0
}

func (m *Certification) GetExpirationTime() int64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *Certification) GetRevocationTime() int64 {
	if m != nil {
		return m.RevocationTime
	}
	return 0
}

type Alias struct {
	AliasName            string   `protobuf:"bytes,1,opt,name=alias_name,json=aliasName,proto3" json:"alias_name,omitempty"`
	AliasCollateral      []byte   `protobuf:"bytes,2,opt,name=alias_collateral,json=aliasCollateral,proto3" json:"alias_collateral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_b735a4e0a688517c, []int{5}
}
func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (dst *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(dst, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetAliasName() string {
	if m != nil {
		return m.AliasName
	}
	return ""
}

func (m *Alias) GetAliasCollateral() []byte {
	if m != nil {
		return m.AliasCollateral
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*AliasAccount)(nil), "corepb.AliasAccount")
	proto.RegisterType((*DataState)(nil), "corepb.DataState")
	proto.RegisterType((*Record)(nil), "corepb.Record")
	proto.RegisterType((*Certification)(nil), "corepb.Certification")
	proto.RegisterType((*Alias)(nil), "corepb.Alias")
}

func init() { proto.RegisterFile("state.proto", fileDescriptor_state_b735a4e0a688517c) }

var fileDescriptor_state_b735a4e0a688517c = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0xe5, 0xa6, 0x49, 0x95, 0x49, 0x5a, 0xb7, 0xfb, 0x45, 0x95, 0xf5, 0x89, 0x8a, 0x60,
	0x21, 0x08, 0x20, 0x71, 0xd2, 0x73, 0xa4, 0x50, 0x0e, 0xe0, 0x04, 0x09, 0x13, 0x0e, 0x91, 0x35,
	0xf1, 0x2e, 0x64, 0x85, 0xe3, 0x8d, 0x76, 0x27, 0x6d, 0x2e, 0x8b, 0x5b, 0xe2, 0x16, 0xb8, 0x02,
	0xb4, 0xb3, 0xfe, 0x8b, 0x7a, 0x96, 0x79, 0xe6, 0xdd, 0x77, 0xfe, 0xac, 0xc0, 0xc4, 0x11, 0x92,
	0x7a, 0xbb, 0xb3, 0x86, 0x8c, 0x18, 0x15, 0xc6, 0xaa, 0xdd, 0x3a, 0xfd, 0x7b, 0x02, 0x67, 0xcb,
	0xa2, 0x30, 0xfb, 0x8a, 0x44, 0x02, 0x67, 0x28, 0xa5, 0x55, 0xce, 0x25, 0xd1, 0x3c, 0x5a, 0x4c,
	0xb3, 0x26, 0xf4, 0x99, 0x35, 0x96, 0x58, 0x15, 0x2a, 0x39, 0x09, 0x99, 0x3a, 0x14, 0x33, 0x18,
	0x56, 0xc6, 0xf3, 0xc1, 0x3c, 0x5a, 0x9c, 0x66, 0x21, 0xf0, 0x7a, 0x47, 0xf8, 0x4b, 0x57, 0x3f,
	0x93, 0x49, 0xd0, 0xd7, 0xa1, 0x78, 0x01, 0xf1, 0xbd, 0x21, 0x25, 0x73, 0x6b, 0x0c, 0xe5, 0x1b,
	0x74, 0x9b, 0x64, 0xca, 0x8a, 0x73, 0xc6, 0x99, 0x31, 0xf4, 0x11, 0xdd, 0x46, 0x3c, 0x83, 0x69,
	0x81, 0x95, 0xd4, 0x12, 0x49, 0xe5, 0x5a, 0x26, 0xe7, 0x2c, 0x9a, 0xb4, 0xec, 0x93, 0x14, 0x4f,
	0x60, 0xbc, 0xc6, 0x4a, 0x3e, 0x68, 0x49, 0x9b, 0xe4, 0x82, 0xf3, 0x1d, 0x10, 0xaf, 0xe1, 0xaa,
	0x44, 0x47, 0x79, 0x4b, 0x72, 0x72, 0x49, 0x3c, 0x8f, 0x16, 0x83, 0x2c, 0xf6, 0x89, 0xf7, 0x0d,
	0x5f, 0x39, 0xef, 0xb4, 0xaf, 0x9a, 0x86, 0x2f, 0x83, 0x53, 0x0b, 0x5a, 0xa7, 0x96, 0x78, 0xa7,
	0xab, 0xce, 0xe9, 0x5b, 0xc3, 0x57, 0x4e, 0x3c, 0x87, 0x0b, 0x89, 0x84, 0xbd, 0xe9, 0x16, 0x6c,
	0x37, 0xf5, 0xb4, 0x19, 0x2e, 0x7d, 0x07, 0xd3, 0x65, 0xa9, 0xd1, 0xf5, 0x17, 0x1f, 0x7e, 0xb6,
	0x8b, 0xaf, 0x33, 0x33, 0x18, 0xa2, 0x57, 0xf2, 0xda, 0xc7, 0x59, 0x08, 0xd2, 0xdf, 0x11, 0x8c,
	0x3f, 0x20, 0xe1, 0x57, 0x7f, 0x50, 0xf1, 0x06, 0x04, 0x1d, 0x72, 0x3e, 0x6e, 0xaf, 0x6e, 0x30,
	0x8a, 0xe9, 0xc0, 0xa2, 0x76, 0xaf, 0xb7, 0x70, 0x6d, 0x55, 0x61, 0xac, 0x7c, 0xf4, 0x20, 0x1c,
	0xf6, 0xbf, 0x90, 0x3d, 0x7e, 0xb4, 0x84, 0x9b, 0x42, 0x59, 0xd2, 0x3f, 0x74, 0x81, 0xa4, 0x4d,
	0xf5, 0xe8, 0xed, 0x80, 0xdf, 0xfe, 0x7f, 0x24, 0x3a, 0xb2, 0x48, 0xbf, 0xc3, 0x28, 0x63, 0x67,
	0xf1, 0x14, 0x26, 0x75, 0x07, 0xbd, 0x3e, 0x21, 0x20, 0xae, 0x36, 0x83, 0xa1, 0x79, 0xa8, 0x94,
	0xad, 0x3b, 0x0a, 0x81, 0xbf, 0x11, 0xe9, 0xad, 0x72, 0x84, 0xdb, 0x1d, 0xd7, 0x1b, 0x64, 0x1d,
	0x48, 0xff, 0x44, 0x70, 0x7e, 0xd7, 0xaf, 0x2e, 0x5e, 0xc1, 0x65, 0xd7, 0x8e, 0x3a, 0xda, 0x49,
	0x8f, 0x73, 0xc1, 0x6b, 0x18, 0x69, 0xe7, 0xf6, 0x6d, 0xc5, 0x3a, 0xf2, 0x25, 0x6b, 0xa9, 0x92,
	0xf5, 0x88, 0x1d, 0x10, 0x37, 0x00, 0xac, 0xcb, 0x7d, 0x17, 0xc9, 0x69, 0xe8, 0x88, 0xc9, 0x4a,
	0x6f, 0x95, 0x78, 0x09, 0xb1, 0x3a, 0xec, 0xb4, 0x0d, 0x0b, 0x63, 0xcd, 0x90, 0x35, 0x17, 0x1d,
	0x6e, 0x84, 0x56, 0xdd, 0x9b, 0xa2, 0x27, 0x1c, 0x05, 0x61, 0x87, 0xbd, 0x30, 0xfd, 0x02, 0x43,
	0xfe, 0x6a, 0x7c, 0x65, 0xfe, 0x0e, 0xf2, 0x0a, 0xb7, 0x8a, 0x87, 0x1a, 0x67, 0x63, 0x26, 0x9f,
	0x71, 0xab, 0xfc, 0xe4, 0x21, 0x5d, 0x98, 0xb2, 0x44, 0x52, 0x16, 0xcb, 0x7a, 0xb0, 0x98, 0xf9,
	0x5d, 0x8b, 0xd7, 0x23, 0xfe, 0x33, 0xb8, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x93, 0x4e, 0x43,
	0x08, 0x1b, 0x04, 0x00, 0x00,
}
