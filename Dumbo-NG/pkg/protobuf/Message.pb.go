// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: Message.proto

package protobuf

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // instance id
	Id     []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Sender uint32 `protobuf:"varint,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Message) GetSender() uint32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// provable broadcast
type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`           // paload
	Validation []byte `protobuf:"bytes,2,opt,name=validation,proto3" json:"validation,omitempty"` // for external validating
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Value) GetValidation() []byte {
	if x != nil {
		return x.Validation
	}
	return nil
}

type Echo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sigshare []byte `protobuf:"bytes,1,opt,name=sigshare,proto3" json:"sigshare,omitempty"`
}

func (x *Echo) Reset() {
	*x = Echo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Echo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Echo) ProtoMessage() {}

func (x *Echo) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Echo.ProtoReflect.Descriptor instead.
func (*Echo) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{2}
}

func (x *Echo) GetSigshare() []byte {
	if x != nil {
		return x.Sigshare
	}
	return nil
}

// smvba
type Lock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Validation []byte `protobuf:"bytes,2,opt,name=validation,proto3" json:"validation,omitempty"`
	Sig        []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *Lock) Reset() {
	*x = Lock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lock) ProtoMessage() {}

func (x *Lock) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lock.ProtoReflect.Descriptor instead.
func (*Lock) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{3}
}

func (x *Lock) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Lock) GetValidation() []byte {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *Lock) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Finish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Validation []byte `protobuf:"bytes,2,opt,name=validation,proto3" json:"validation,omitempty"`
	Sig        []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *Finish) Reset() {
	*x = Finish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Finish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Finish) ProtoMessage() {}

func (x *Finish) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Finish.ProtoReflect.Descriptor instead.
func (*Finish) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{4}
}

func (x *Finish) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Finish) GetValidation() []byte {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *Finish) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Done struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinShare []byte `protobuf:"bytes,1,opt,name=coinShare,proto3" json:"coinShare,omitempty"`
}

func (x *Done) Reset() {
	*x = Done{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Done) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Done) ProtoMessage() {}

func (x *Done) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Done.ProtoReflect.Descriptor instead.
func (*Done) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{5}
}

func (x *Done) GetCoinShare() []byte {
	if x != nil {
		return x.CoinShare
	}
	return nil
}

type Halt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Validation []byte `protobuf:"bytes,2,opt,name=validation,proto3" json:"validation,omitempty"`
	Sig        []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *Halt) Reset() {
	*x = Halt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Halt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Halt) ProtoMessage() {}

func (x *Halt) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Halt.ProtoReflect.Descriptor instead.
func (*Halt) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{6}
}

func (x *Halt) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Halt) GetValidation() []byte {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *Halt) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type PreVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote       bool   `protobuf:"varint,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Validation []byte `protobuf:"bytes,3,opt,name=validation,proto3" json:"validation,omitempty"`
	Sig        []byte `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *PreVote) Reset() {
	*x = PreVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreVote) ProtoMessage() {}

func (x *PreVote) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreVote.ProtoReflect.Descriptor instead.
func (*PreVote) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{7}
}

func (x *PreVote) GetVote() bool {
	if x != nil {
		return x.Vote
	}
	return false
}

func (x *PreVote) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *PreVote) GetValidation() []byte {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *PreVote) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote       bool   `protobuf:"varint,1,opt,name=vote,proto3" json:"vote,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Validation []byte `protobuf:"bytes,3,opt,name=validation,proto3" json:"validation,omitempty"`
	Sig        []byte `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
	Sigshare   []byte `protobuf:"bytes,5,opt,name=sigshare,proto3" json:"sigshare,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{8}
}

func (x *Vote) GetVote() bool {
	if x != nil {
		return x.Vote
	}
	return false
}

func (x *Vote) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Vote) GetValidation() []byte {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *Vote) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *Vote) GetSigshare() []byte {
	if x != nil {
		return x.Sigshare
	}
	return nil
}

//Dumbo-NG
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx   []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"` // hash of previous tx
	Sig  []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`   // signature on previous tx
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{9}
}

func (x *Proposal) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *Proposal) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Proposal) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Received struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sigshare []byte `protobuf:"bytes,1,opt,name=sigshare,proto3" json:"sigshare,omitempty"`
}

func (x *Received) Reset() {
	*x = Received{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Received) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Received) ProtoMessage() {}

func (x *Received) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Received.ProtoReflect.Descriptor instead.
func (*Received) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{10}
}

func (x *Received) GetSigshare() []byte {
	if x != nil {
		return x.Sigshare
	}
	return nil
}

type BLockSetValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  []uint32 `protobuf:"varint,1,rep,packed,name=pid,proto3" json:"pid,omitempty"`
	Slot []uint32 `protobuf:"varint,2,rep,packed,name=slot,proto3" json:"slot,omitempty"`
	Hash [][]byte `protobuf:"bytes,3,rep,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BLockSetValue) Reset() {
	*x = BLockSetValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLockSetValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLockSetValue) ProtoMessage() {}

func (x *BLockSetValue) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLockSetValue.ProtoReflect.Descriptor instead.
func (*BLockSetValue) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{11}
}

func (x *BLockSetValue) GetPid() []uint32 {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *BLockSetValue) GetSlot() []uint32 {
	if x != nil {
		return x.Slot
	}
	return nil
}

func (x *BLockSetValue) GetHash() [][]byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type BLockSetValidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig [][]byte `protobuf:"bytes,1,rep,name=sig,proto3" json:"sig,omitempty"`
}

func (x *BLockSetValidation) Reset() {
	*x = BLockSetValidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLockSetValidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLockSetValidation) ProtoMessage() {}

func (x *BLockSetValidation) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLockSetValidation.ProtoReflect.Descriptor instead.
func (*BLockSetValidation) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{12}
}

func (x *BLockSetValidation) GetSig() [][]byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

//callHelp
type CallHelp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Slot uint32 `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	Sig  []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *CallHelp) Reset() {
	*x = CallHelp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallHelp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallHelp) ProtoMessage() {}

func (x *CallHelp) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallHelp.ProtoReflect.Descriptor instead.
func (*CallHelp) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{13}
}

func (x *CallHelp) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *CallHelp) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *CallHelp) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Help struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid    uint32   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Slot   uint32   `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	Shard  []byte   `protobuf:"bytes,3,opt,name=shard,proto3" json:"shard,omitempty"`
	Root   []byte   `protobuf:"bytes,4,opt,name=root,proto3" json:"root,omitempty"`
	Proof1 [][]byte `protobuf:"bytes,5,rep,name=proof1,proto3" json:"proof1,omitempty"`
	Proof2 []int64  `protobuf:"varint,6,rep,packed,name=proof2,proto3" json:"proof2,omitempty"`
}

func (x *Help) Reset() {
	*x = Help{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Help) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Help) ProtoMessage() {}

func (x *Help) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Help.ProtoReflect.Descriptor instead.
func (*Help) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{14}
}

func (x *Help) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Help) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *Help) GetShard() []byte {
	if x != nil {
		return x.Shard
	}
	return nil
}

func (x *Help) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *Help) GetProof1() [][]byte {
	if x != nil {
		return x.Proof1
	}
	return nil
}

func (x *Help) GetProof2() []int64 {
	if x != nil {
		return x.Proof2
	}
	return nil
}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x59, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x04, 0x45, 0x63, 0x68,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x67, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x4e, 0x0a,
	0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x50, 0x0a,
	0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22,
	0x24, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x69, 0x6e,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x4e, 0x0a, 0x04, 0x48, 0x61, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x65, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x76, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x7e, 0x0a, 0x04,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x73, 0x69, 0x67, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x40, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x26,
	0x0a, 0x08, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69,
	0x67, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x69,
	0x67, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x42, 0x4c, 0x6f, 0x63, 0x6b, 0x53,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x22, 0x26, 0x0a, 0x12, 0x42, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x42, 0x0a, 0x08, 0x43, 0x61, 0x6c,
	0x6c, 0x48, 0x65, 0x6c, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0x86, 0x01,
	0x0a, 0x04, 0x48, 0x65, 0x6c, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x31,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x31, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x32, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x32, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Message_proto_rawDescOnce sync.Once
	file_Message_proto_rawDescData = file_Message_proto_rawDesc
)

func file_Message_proto_rawDescGZIP() []byte {
	file_Message_proto_rawDescOnce.Do(func() {
		file_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Message_proto_rawDescData)
	})
	return file_Message_proto_rawDescData
}

var file_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_Message_proto_goTypes = []interface{}{
	(*Message)(nil),            // 0: Message
	(*Value)(nil),              // 1: Value
	(*Echo)(nil),               // 2: Echo
	(*Lock)(nil),               // 3: Lock
	(*Finish)(nil),             // 4: Finish
	(*Done)(nil),               // 5: Done
	(*Halt)(nil),               // 6: Halt
	(*PreVote)(nil),            // 7: PreVote
	(*Vote)(nil),               // 8: Vote
	(*Proposal)(nil),           // 9: Proposal
	(*Received)(nil),           // 10: Received
	(*BLockSetValue)(nil),      // 11: BLockSetValue
	(*BLockSetValidation)(nil), // 12: BLockSetValidation
	(*CallHelp)(nil),           // 13: CallHelp
	(*Help)(nil),               // 14: Help
}
var file_Message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Message_proto_init() }
func file_Message_proto_init() {
	if File_Message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_Message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_Message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Echo); i {
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
		file_Message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lock); i {
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
		file_Message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Finish); i {
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
		file_Message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Done); i {
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
		file_Message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Halt); i {
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
		file_Message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreVote); i {
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
		file_Message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_Message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_Message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Received); i {
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
		file_Message_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLockSetValue); i {
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
		file_Message_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLockSetValidation); i {
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
		file_Message_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallHelp); i {
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
		file_Message_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Help); i {
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
			RawDescriptor: file_Message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Message_proto_goTypes,
		DependencyIndexes: file_Message_proto_depIdxs,
		MessageInfos:      file_Message_proto_msgTypes,
	}.Build()
	File_Message_proto = out.File
	file_Message_proto_rawDesc = nil
	file_Message_proto_goTypes = nil
	file_Message_proto_depIdxs = nil
}
