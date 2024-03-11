// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: spacemesh/v1/debug_types.proto

package spacemeshv1

import (
	_ "google.golang.org/genproto/googleapis/rpc/status"
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

type NetworkInfoResponse_NATType int32

const (
	NetworkInfoResponse_NATTypeUnknown NetworkInfoResponse_NATType = 0
	NetworkInfoResponse_Cone           NetworkInfoResponse_NATType = 1
	NetworkInfoResponse_Symmetric      NetworkInfoResponse_NATType = 2
)

// Enum value maps for NetworkInfoResponse_NATType.
var (
	NetworkInfoResponse_NATType_name = map[int32]string{
		0: "NATTypeUnknown",
		1: "Cone",
		2: "Symmetric",
	}
	NetworkInfoResponse_NATType_value = map[string]int32{
		"NATTypeUnknown": 0,
		"Cone":           1,
		"Symmetric":      2,
	}
)

func (x NetworkInfoResponse_NATType) Enum() *NetworkInfoResponse_NATType {
	p := new(NetworkInfoResponse_NATType)
	*p = x
	return p
}

func (x NetworkInfoResponse_NATType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkInfoResponse_NATType) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v1_debug_types_proto_enumTypes[0].Descriptor()
}

func (NetworkInfoResponse_NATType) Type() protoreflect.EnumType {
	return &file_spacemesh_v1_debug_types_proto_enumTypes[0]
}

func (x NetworkInfoResponse_NATType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkInfoResponse_NATType.Descriptor instead.
func (NetworkInfoResponse_NATType) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{2, 0}
}

type NetworkInfoResponse_Reachability int32

const (
	NetworkInfoResponse_ReachabilityUnknown NetworkInfoResponse_Reachability = 0
	NetworkInfoResponse_Public              NetworkInfoResponse_Reachability = 1
	NetworkInfoResponse_Private             NetworkInfoResponse_Reachability = 2
)

// Enum value maps for NetworkInfoResponse_Reachability.
var (
	NetworkInfoResponse_Reachability_name = map[int32]string{
		0: "ReachabilityUnknown",
		1: "Public",
		2: "Private",
	}
	NetworkInfoResponse_Reachability_value = map[string]int32{
		"ReachabilityUnknown": 0,
		"Public":              1,
		"Private":             2,
	}
)

func (x NetworkInfoResponse_Reachability) Enum() *NetworkInfoResponse_Reachability {
	p := new(NetworkInfoResponse_Reachability)
	*p = x
	return p
}

func (x NetworkInfoResponse_Reachability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkInfoResponse_Reachability) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v1_debug_types_proto_enumTypes[1].Descriptor()
}

func (NetworkInfoResponse_Reachability) Type() protoreflect.EnumType {
	return &file_spacemesh_v1_debug_types_proto_enumTypes[1]
}

func (x NetworkInfoResponse_Reachability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkInfoResponse_Reachability.Descriptor instead.
func (NetworkInfoResponse_Reachability) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{2, 1}
}

type Proposal_Status int32

const (
	Proposal_Created  Proposal_Status = 0
	Proposal_Included Proposal_Status = 1
)

// Enum value maps for Proposal_Status.
var (
	Proposal_Status_name = map[int32]string{
		0: "Created",
		1: "Included",
	}
	Proposal_Status_value = map[string]int32{
		"Created":  0,
		"Included": 1,
	}
)

func (x Proposal_Status) Enum() *Proposal_Status {
	p := new(Proposal_Status)
	*p = x
	return p
}

func (x Proposal_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Proposal_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spacemesh_v1_debug_types_proto_enumTypes[2].Descriptor()
}

func (Proposal_Status) Type() protoreflect.EnumType {
	return &file_spacemesh_v1_debug_types_proto_enumTypes[2]
}

func (x Proposal_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Proposal_Status.Descriptor instead.
func (Proposal_Status) EnumDescriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{7, 0}
}

type AccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layer uint32 `protobuf:"varint,1,opt,name=layer,proto3" json:"layer,omitempty"`
}

func (x *AccountsRequest) Reset() {
	*x = AccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsRequest) ProtoMessage() {}

func (x *AccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsRequest.ProtoReflect.Descriptor instead.
func (*AccountsRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{0}
}

func (x *AccountsRequest) GetLayer() uint32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

type AccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountWrapper []*Account `protobuf:"bytes,1,rep,name=account_wrapper,json=accountWrapper,proto3" json:"account_wrapper,omitempty"`
}

func (x *AccountsResponse) Reset() {
	*x = AccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsResponse) ProtoMessage() {}

func (x *AccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsResponse.ProtoReflect.Descriptor instead.
func (*AccountsResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{1}
}

func (x *AccountsResponse) GetAccountWrapper() []*Account {
	if x != nil {
		return x.AccountWrapper
	}
	return nil
}

type NetworkInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ListenAddresses  []string                         `protobuf:"bytes,2,rep,name=listen_addresses,json=listenAddresses,proto3" json:"listen_addresses,omitempty"`
	KnownAddresses   []string                         `protobuf:"bytes,3,rep,name=known_addresses,json=knownAddresses,proto3" json:"known_addresses,omitempty"`
	NatTypeUdp       NetworkInfoResponse_NATType      `protobuf:"varint,4,opt,name=nat_type_udp,json=natTypeUdp,proto3,enum=spacemesh.v1.NetworkInfoResponse_NATType" json:"nat_type_udp,omitempty"`
	NatTypeTcp       NetworkInfoResponse_NATType      `protobuf:"varint,5,opt,name=nat_type_tcp,json=natTypeTcp,proto3,enum=spacemesh.v1.NetworkInfoResponse_NATType" json:"nat_type_tcp,omitempty"`
	Reachability     NetworkInfoResponse_Reachability `protobuf:"varint,6,opt,name=reachability,proto3,enum=spacemesh.v1.NetworkInfoResponse_Reachability" json:"reachability,omitempty"`
	DhtServerEnabled bool                             `protobuf:"varint,7,opt,name=dht_server_enabled,json=dhtServerEnabled,proto3" json:"dht_server_enabled,omitempty"`
}

func (x *NetworkInfoResponse) Reset() {
	*x = NetworkInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfoResponse) ProtoMessage() {}

func (x *NetworkInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfoResponse.ProtoReflect.Descriptor instead.
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkInfoResponse) GetListenAddresses() []string {
	if x != nil {
		return x.ListenAddresses
	}
	return nil
}

func (x *NetworkInfoResponse) GetKnownAddresses() []string {
	if x != nil {
		return x.KnownAddresses
	}
	return nil
}

func (x *NetworkInfoResponse) GetNatTypeUdp() NetworkInfoResponse_NATType {
	if x != nil {
		return x.NatTypeUdp
	}
	return NetworkInfoResponse_NATTypeUnknown
}

func (x *NetworkInfoResponse) GetNatTypeTcp() NetworkInfoResponse_NATType {
	if x != nil {
		return x.NatTypeTcp
	}
	return NetworkInfoResponse_NATTypeUnknown
}

func (x *NetworkInfoResponse) GetReachability() NetworkInfoResponse_Reachability {
	if x != nil {
		return x.Reachability
	}
	return NetworkInfoResponse_ReachabilityUnknown
}

func (x *NetworkInfoResponse) GetDhtServerEnabled() bool {
	if x != nil {
		return x.DhtServerEnabled
	}
	return false
}

type EpochData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beacon []byte `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
}

func (x *EpochData) Reset() {
	*x = EpochData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochData) ProtoMessage() {}

func (x *EpochData) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochData.ProtoReflect.Descriptor instead.
func (*EpochData) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{3}
}

func (x *EpochData) GetBeacon() []byte {
	if x != nil {
		return x.Beacon
	}
	return nil
}

type Eligibility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	J         uint32 `protobuf:"varint,1,opt,name=j,proto3" json:"j,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Eligibility) Reset() {
	*x = Eligibility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Eligibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Eligibility) ProtoMessage() {}

func (x *Eligibility) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Eligibility.ProtoReflect.Descriptor instead.
func (*Eligibility) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{4}
}

func (x *Eligibility) GetJ() uint32 {
	if x != nil {
		return x.J
	}
	return 0
}

func (x *Eligibility) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ActiveSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch uint32 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (x *ActiveSetRequest) Reset() {
	*x = ActiveSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveSetRequest) ProtoMessage() {}

func (x *ActiveSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveSetRequest.ProtoReflect.Descriptor instead.
func (*ActiveSetRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{5}
}

func (x *ActiveSetRequest) GetEpoch() uint32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

type ActiveSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []*ActivationId `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ActiveSetResponse) Reset() {
	*x = ActiveSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveSetResponse) ProtoMessage() {}

func (x *ActiveSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveSetResponse.ProtoReflect.Descriptor instead.
func (*ActiveSetResponse) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{6}
}

func (x *ActiveSetResponse) GetIds() []*ActivationId {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Epoch   *EpochNumber `protobuf:"bytes,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Layer   *LayerNumber `protobuf:"bytes,3,opt,name=layer,proto3" json:"layer,omitempty"`
	Smesher *SmesherId   `protobuf:"bytes,4,opt,name=smesher,proto3" json:"smesher,omitempty"`
	// Types that are assignable to EpochData:
	//
	//	*Proposal_Reference
	//	*Proposal_Data
	EpochData     isProposal_EpochData `protobuf_oneof:"epoch_data"`
	Ballot        []byte               `protobuf:"bytes,7,opt,name=ballot,proto3" json:"ballot,omitempty"`
	Eligibilities []*Eligibility       `protobuf:"bytes,8,rep,name=eligibilities,proto3" json:"eligibilities,omitempty"`
	Status        Proposal_Status      `protobuf:"varint,9,opt,name=status,proto3,enum=spacemesh.v1.Proposal_Status" json:"status,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[7]
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
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{7}
}

func (x *Proposal) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Proposal) GetEpoch() *EpochNumber {
	if x != nil {
		return x.Epoch
	}
	return nil
}

func (x *Proposal) GetLayer() *LayerNumber {
	if x != nil {
		return x.Layer
	}
	return nil
}

func (x *Proposal) GetSmesher() *SmesherId {
	if x != nil {
		return x.Smesher
	}
	return nil
}

func (m *Proposal) GetEpochData() isProposal_EpochData {
	if m != nil {
		return m.EpochData
	}
	return nil
}

func (x *Proposal) GetReference() []byte {
	if x, ok := x.GetEpochData().(*Proposal_Reference); ok {
		return x.Reference
	}
	return nil
}

func (x *Proposal) GetData() *EpochData {
	if x, ok := x.GetEpochData().(*Proposal_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Proposal) GetBallot() []byte {
	if x != nil {
		return x.Ballot
	}
	return nil
}

func (x *Proposal) GetEligibilities() []*Eligibility {
	if x != nil {
		return x.Eligibilities
	}
	return nil
}

func (x *Proposal) GetStatus() Proposal_Status {
	if x != nil {
		return x.Status
	}
	return Proposal_Created
}

type isProposal_EpochData interface {
	isProposal_EpochData()
}

type Proposal_Reference struct {
	Reference []byte `protobuf:"bytes,5,opt,name=reference,proto3,oneof"`
}

type Proposal_Data struct {
	Data *EpochData `protobuf:"bytes,6,opt,name=data,proto3,oneof"`
}

func (*Proposal_Reference) isProposal_EpochData() {}

func (*Proposal_Data) isProposal_EpochData() {}

type ChangeLogLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Level  string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *ChangeLogLevelRequest) Reset() {
	*x = ChangeLogLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacemesh_v1_debug_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLogLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLogLevelRequest) ProtoMessage() {}

func (x *ChangeLogLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spacemesh_v1_debug_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLogLevelRequest.ProtoReflect.Descriptor instead.
func (*ChangeLogLevelRequest) Descriptor() ([]byte, []int) {
	return file_spacemesh_v1_debug_types_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeLogLevelRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *ChangeLogLevelRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

var File_spacemesh_v1_debug_types_proto protoreflect.FileDescriptor

var file_spacemesh_v1_debug_types_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x25,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x22, 0x52, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x22, 0x8f, 0x04, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a,
	0x10, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x6e, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x64,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x41, 0x54, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x55, 0x64, 0x70, 0x12, 0x4b,
	0x0a, 0x0c, 0x6e, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x74, 0x63, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x41, 0x54, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x6e, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x54, 0x63, 0x70, 0x12, 0x52, 0x0a, 0x0c, 0x72,
	0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x2c, 0x0a, 0x12, 0x64, 0x68, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x68, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x36, 0x0a,
	0x07, 0x4e, 0x41, 0x54, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x41, 0x54, 0x54,
	0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x43, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x10, 0x02, 0x22, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x10, 0x02, 0x22, 0x23, 0x0a, 0x09, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x0b,
	0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x6a,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6a, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xc1, 0x03, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x05, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x12, 0x2f, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x05, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x52, 0x07, 0x73,
	0x6d, 0x65, 0x73, 0x68, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x12, 0x3f, 0x0a,
	0x0d, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0d, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x23, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x10, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42,
	0xb4, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x0c,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacemesh_v1_debug_types_proto_rawDescOnce sync.Once
	file_spacemesh_v1_debug_types_proto_rawDescData = file_spacemesh_v1_debug_types_proto_rawDesc
)

func file_spacemesh_v1_debug_types_proto_rawDescGZIP() []byte {
	file_spacemesh_v1_debug_types_proto_rawDescOnce.Do(func() {
		file_spacemesh_v1_debug_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacemesh_v1_debug_types_proto_rawDescData)
	})
	return file_spacemesh_v1_debug_types_proto_rawDescData
}

var file_spacemesh_v1_debug_types_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_spacemesh_v1_debug_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spacemesh_v1_debug_types_proto_goTypes = []interface{}{
	(NetworkInfoResponse_NATType)(0),      // 0: spacemesh.v1.NetworkInfoResponse.NATType
	(NetworkInfoResponse_Reachability)(0), // 1: spacemesh.v1.NetworkInfoResponse.Reachability
	(Proposal_Status)(0),                  // 2: spacemesh.v1.Proposal.Status
	(*AccountsRequest)(nil),               // 3: spacemesh.v1.AccountsRequest
	(*AccountsResponse)(nil),              // 4: spacemesh.v1.AccountsResponse
	(*NetworkInfoResponse)(nil),           // 5: spacemesh.v1.NetworkInfoResponse
	(*EpochData)(nil),                     // 6: spacemesh.v1.EpochData
	(*Eligibility)(nil),                   // 7: spacemesh.v1.Eligibility
	(*ActiveSetRequest)(nil),              // 8: spacemesh.v1.ActiveSetRequest
	(*ActiveSetResponse)(nil),             // 9: spacemesh.v1.ActiveSetResponse
	(*Proposal)(nil),                      // 10: spacemesh.v1.Proposal
	(*ChangeLogLevelRequest)(nil),         // 11: spacemesh.v1.ChangeLogLevelRequest
	(*Account)(nil),                       // 12: spacemesh.v1.Account
	(*ActivationId)(nil),                  // 13: spacemesh.v1.ActivationId
	(*EpochNumber)(nil),                   // 14: spacemesh.v1.EpochNumber
	(*LayerNumber)(nil),                   // 15: spacemesh.v1.LayerNumber
	(*SmesherId)(nil),                     // 16: spacemesh.v1.SmesherId
}
var file_spacemesh_v1_debug_types_proto_depIdxs = []int32{
	12, // 0: spacemesh.v1.AccountsResponse.account_wrapper:type_name -> spacemesh.v1.Account
	0,  // 1: spacemesh.v1.NetworkInfoResponse.nat_type_udp:type_name -> spacemesh.v1.NetworkInfoResponse.NATType
	0,  // 2: spacemesh.v1.NetworkInfoResponse.nat_type_tcp:type_name -> spacemesh.v1.NetworkInfoResponse.NATType
	1,  // 3: spacemesh.v1.NetworkInfoResponse.reachability:type_name -> spacemesh.v1.NetworkInfoResponse.Reachability
	13, // 4: spacemesh.v1.ActiveSetResponse.ids:type_name -> spacemesh.v1.ActivationId
	14, // 5: spacemesh.v1.Proposal.epoch:type_name -> spacemesh.v1.EpochNumber
	15, // 6: spacemesh.v1.Proposal.layer:type_name -> spacemesh.v1.LayerNumber
	16, // 7: spacemesh.v1.Proposal.smesher:type_name -> spacemesh.v1.SmesherId
	6,  // 8: spacemesh.v1.Proposal.data:type_name -> spacemesh.v1.EpochData
	7,  // 9: spacemesh.v1.Proposal.eligibilities:type_name -> spacemesh.v1.Eligibility
	2,  // 10: spacemesh.v1.Proposal.status:type_name -> spacemesh.v1.Proposal.Status
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_debug_types_proto_init() }
func file_spacemesh_v1_debug_types_proto_init() {
	if File_spacemesh_v1_debug_types_proto != nil {
		return
	}
	file_spacemesh_v1_global_state_types_proto_init()
	file_spacemesh_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spacemesh_v1_debug_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsRequest); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsResponse); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfoResponse); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochData); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Eligibility); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveSetRequest); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveSetResponse); i {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_spacemesh_v1_debug_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLogLevelRequest); i {
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
	file_spacemesh_v1_debug_types_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Proposal_Reference)(nil),
		(*Proposal_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_debug_types_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacemesh_v1_debug_types_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_debug_types_proto_depIdxs,
		EnumInfos:         file_spacemesh_v1_debug_types_proto_enumTypes,
		MessageInfos:      file_spacemesh_v1_debug_types_proto_msgTypes,
	}.Build()
	File_spacemesh_v1_debug_types_proto = out.File
	file_spacemesh_v1_debug_types_proto_rawDesc = nil
	file_spacemesh_v1_debug_types_proto_goTypes = nil
	file_spacemesh_v1_debug_types_proto_depIdxs = nil
}
