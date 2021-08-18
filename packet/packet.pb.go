package packet

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ResolveStatus encodes the status of an oracle request.
type ResolveStatus int32

const (
	// Open - the request is not yet resolved.
	RESOLVE_STATUS_OPEN ResolveStatus = 0
	// Success - the request has been resolved successfully with no errors.
	RESOLVE_STATUS_SUCCESS ResolveStatus = 1
	// Failure - an error occured during the request's resolve call.
	RESOLVE_STATUS_FAILURE ResolveStatus = 2
	// Expired - the request does not get enough reports from validator within the
	// timeframe.
	RESOLVE_STATUS_EXPIRED ResolveStatus = 3
)

var ResolveStatus_name = map[int32]string{
	0: "RESOLVE_STATUS_OPEN_UNSPECIFIED",
	1: "RESOLVE_STATUS_SUCCESS",
	2: "RESOLVE_STATUS_FAILURE",
	3: "RESOLVE_STATUS_EXPIRED",
}

var ResolveStatus_value = map[string]int32{
	"RESOLVE_STATUS_OPEN_UNSPECIFIED": 0,
	"RESOLVE_STATUS_SUCCESS":          1,
	"RESOLVE_STATUS_FAILURE":          2,
	"RESOLVE_STATUS_EXPIRED":          3,
}

func (x ResolveStatus) String() string {
	return proto.EnumName(ResolveStatus_name, int32(x))
}

func (ResolveStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{0}
}

// OracleRequestPacketData encodes an oracle request sent from other blockchains
// to BandChain.
type OracleRequestPacketData struct {
	// ClientID is the unique identifier of this oracle request, as specified by
	// the client. This same unique ID will be sent back to the requester with the
	// oracle response.
	ClientID string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// OracleScriptID is the unique identifier of the oracle script to be
	// executed.
	OracleScriptID uint64 `protobuf:"varint,2,opt,name=oracle_script_id,json=oracleScriptId,proto3,casttype=OracleScriptID" json:"oracle_script_id,omitempty"`
	// Calldata is the OBI-encoded calldata bytes available for oracle executor to
	// read.
	Calldata []byte `protobuf:"bytes,3,opt,name=calldata,proto3" json:"calldata,omitempty"`
	// AskCount is the number of validators that are requested to respond to this
	// oracle request. Higher value means more security, at a higher gas cost.
	AskCount uint64 `protobuf:"varint,4,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty"`
	// MinCount is the minimum number of validators necessary for the request to
	// proceed to the execution phase. Higher value means more security, at the
	// cost of liveness.
	MinCount uint64 `protobuf:"varint,5,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	// FeeLimit is the maximum tokens that will be paid to all data source
	// providers.
	FeeLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=fee_limit,json=feeLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_limit"`
	// PrepareGas is amount of gas to pay to prepare raw requests
	PrepareGas uint64 `protobuf:"varint,7,opt,name=prepare_gas,json=prepareGas,proto3" json:"prepare_gas,omitempty"`
	// ExecuteGas is amount of gas to reserve for executing
	ExecuteGas uint64 `protobuf:"varint,8,opt,name=execute_gas,json=executeGas,proto3" json:"execute_gas,omitempty"`
}

func (m *OracleRequestPacketData) Reset()         { *m = OracleRequestPacketData{} }
func (m *OracleRequestPacketData) String() string { return proto.CompactTextString(m) }
func (*OracleRequestPacketData) ProtoMessage()    {}
func (*OracleRequestPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{6}
}
func (m *OracleRequestPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleRequestPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleRequestPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleRequestPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleRequestPacketData.Merge(m, src)
}
func (m *OracleRequestPacketData) XXX_Size() int {
	return m.Size()
}
func (m *OracleRequestPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleRequestPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_OracleRequestPacketData proto.InternalMessageInfo

func (m *OracleRequestPacketData) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OracleRequestPacketData) GetOracleScriptID() uint64 {
	if m != nil {
		return m.OracleScriptID
	}
	return 0
}

func (m *OracleRequestPacketData) GetCalldata() []byte {
	if m != nil {
		return m.Calldata
	}
	return nil
}

func (m *OracleRequestPacketData) GetAskCount() uint64 {
	if m != nil {
		return m.AskCount
	}
	return 0
}

func (m *OracleRequestPacketData) GetMinCount() uint64 {
	if m != nil {
		return m.MinCount
	}
	return 0
}

func (m *OracleRequestPacketData) GetFeeLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeLimit
	}
	return nil
}

func (m *OracleRequestPacketData) GetPrepareGas() uint64 {
	if m != nil {
		return m.PrepareGas
	}
	return 0
}

func (m *OracleRequestPacketData) GetExecuteGas() uint64 {
	if m != nil {
		return m.ExecuteGas
	}
	return 0
}

// OracleRequestPacketAcknowledgement encodes an oracle request acknowledgement
// send back to requester chain.
type OracleRequestPacketAcknowledgement struct {
	// RequestID is BandChain's unique identifier for this oracle request.
	RequestID uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3,casttype=RequestID" json:"request_id,omitempty"`
}

func (m *OracleRequestPacketAcknowledgement) Reset()         { *m = OracleRequestPacketAcknowledgement{} }
func (m *OracleRequestPacketAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*OracleRequestPacketAcknowledgement) ProtoMessage()    {}
func (*OracleRequestPacketAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{7}
}
func (m *OracleRequestPacketAcknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleRequestPacketAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleRequestPacketAcknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleRequestPacketAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleRequestPacketAcknowledgement.Merge(m, src)
}
func (m *OracleRequestPacketAcknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *OracleRequestPacketAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleRequestPacketAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_OracleRequestPacketAcknowledgement proto.InternalMessageInfo

func (m *OracleRequestPacketAcknowledgement) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

// OracleResponsePacketData encodes an oracle response from BandChain to the
// requester.
type OracleResponsePacketData struct {
	// ClientID is the unique identifier matched with that of the oracle request
	// packet.
	ClientID string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// RequestID is BandChain's unique identifier for this oracle request.
	RequestID uint64 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3,casttype=RequestID" json:"request_id,omitempty"`
	// AnsCount is the number of validators among to the asked validators that
	// actually responded to this oracle request prior to this oracle request
	// being resolved.
	AnsCount uint64 `protobuf:"varint,3,opt,name=ans_count,json=ansCount,proto3" json:"ans_count,omitempty"`
	// RequestTime is the UNIX epoch time at which the request was sent to
	// BandChain.
	RequestTime int64 `protobuf:"varint,4,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// ResolveTime is the UNIX epoch time at which the request was resolved to the
	// final result.
	ResolveTime int64 `protobuf:"varint,5,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
	// ResolveStatus is the status of this oracle request, which can be OK,
	// FAILURE, or EXPIRED.
	ResolveStatus ResolveStatus `protobuf:"varint,6,opt,name=resolve_status,json=resolveStatus,proto3,enum=oracle.v1.ResolveStatus" json:"resolve_status,omitempty"`
	// Result is the final aggregated value encoded in OBI format. Only available
	// if status if OK.
	Result []byte `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *OracleResponsePacketData) Reset()         { *m = OracleResponsePacketData{} }
func (m *OracleResponsePacketData) String() string { return proto.CompactTextString(m) }
func (*OracleResponsePacketData) ProtoMessage()    {}
func (*OracleResponsePacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_652b57db11528d07, []int{8}
}
func (m *OracleResponsePacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleResponsePacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleResponsePacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleResponsePacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleResponsePacketData.Merge(m, src)
}
func (m *OracleResponsePacketData) XXX_Size() int {
	return m.Size()
}
func (m *OracleResponsePacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleResponsePacketData.DiscardUnknown(m)
}

var xxx_messageInfo_OracleResponsePacketData proto.InternalMessageInfo

func (m *OracleResponsePacketData) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *OracleResponsePacketData) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *OracleResponsePacketData) GetAnsCount() uint64 {
	if m != nil {
		return m.AnsCount
	}
	return 0
}

func (m *OracleResponsePacketData) GetRequestTime() int64 {
	if m != nil {
		return m.RequestTime
	}
	return 0
}

func (m *OracleResponsePacketData) GetResolveTime() int64 {
	if m != nil {
		return m.ResolveTime
	}
	return 0
}

func (m *OracleResponsePacketData) GetResolveStatus() ResolveStatus {
	if m != nil {
		return m.ResolveStatus
	}
	return RESOLVE_STATUS_OPEN
}

func (m *OracleResponsePacketData) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterEnum("oracle.v1.ResolveStatus", ResolveStatus_name, ResolveStatus_value)
	proto.RegisterType((*OracleRequestPacketData)(nil), "oracle.v1.OracleRequestPacketData")
	proto.RegisterType((*OracleRequestPacketAcknowledgement)(nil), "oracle.v1.OracleRequestPacketAcknowledgement")
	proto.RegisterType((*OracleResponsePacketData)(nil), "oracle.v1.OracleResponsePacketData")
}

func init() { proto.RegisterFile("oracle/v1/oracle.proto", fileDescriptor_652b57db11528d07) }

var fileDescriptor_652b57db11528d07 = []byte{
	// 1767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x58, 0xcd, 0x6f, 0x23, 0x59,
	0x11, 0x4f, 0xdb, 0x4e, 0x62, 0x97, 0x9d, 0xaf, 0x37, 0x99, 0x8c, 0xc7, 0xbb, 0xd8, 0x21, 0x5a,
	0xd0, 0x30, 0x5a, 0x6c, 0x32, 0x20, 0xc4, 0xcc, 0xf2, 0xa1, 0xd8, 0x71, 0x16, 0xa3, 0x68, 0xc6,
	0x7a, 0x4e, 0x46, 0x80, 0x84, 0x5a, 0xcf, 0xdd, 0x2f, 0xce, 0x53, 0xfa, 0x8b, 0xf7, 0xda, 0x89,
	0xb3, 0x37, 0x6e, 0x68, 0x4f, 0x7b, 0x41, 0xe2, 0xc0, 0x4a, 0x2b, 0x71, 0xe3, 0x8a, 0xf8, 0x07,
	0x38, 0xcd, 0x8d, 0x3d, 0x21, 0x24, 0x24, 0x2f, 0xf2, 0x5c, 0xe0, 0x0f, 0xe0, 0x02, 0x17, 0xf4,
	0x3e, 0xda, 0x6d, 0x9b, 0x2c, 0x43, 0x86, 0x8f, 0xc3, 0x9e, 0xe2, 0xfa, 0x55, 0x55, 0x77, 0xbd,
	0xaa, 0x5f, 0xd5, 0xab, 0x0e, 0xec, 0x84, 0x9c, 0x38, 0x1e, 0x6d, 0x5c, 0xee, 0x37, 0xf4, 0xaf,
	0x7a, 0xc4, 0xc3, 0x38, 0x44, 0x05, 0x23, 0x5d, 0xee, 0x57, 0xb6, 0x07, 0xe1, 0x20, 0x54, 0x68,
	0x43, 0xfe, 0xd2, 0x06, 0x95, 0xda, 0x20, 0x0c, 0x07, 0x1e, 0x6d, 0x28, 0xa9, 0x3f, 0x3c, 0x6b,
	0xc4, 0xcc, 0xa7, 0x22, 0x26, 0x7e, 0x64, 0x0c, 0xee, 0x2f, 0x1a, 0x90, 0xe0, 0xda, 0xa8, 0xaa,
	0x4e, 0x28, 0xfc, 0x50, 0x34, 0xfa, 0x44, 0xc8, 0x37, 0xf7, 0x69, 0x4c, 0xf6, 0x1b, 0x4e, 0xc8,
	0x02, 0xad, 0xdf, 0xfb, 0xab, 0x05, 0x70, 0x48, 0x62, 0xd2, 0x0b, 0x87, 0xdc, 0xa1, 0x68, 0x1b,
	0x96, 0xc3, 0xab, 0x80, 0xf2, 0xb2, 0xb5, 0x6b, 0x3d, 0x28, 0x60, 0x2d, 0x20, 0x04, 0xb9, 0x80,
	0xf8, 0xb4, 0x9c, 0x51, 0xa0, 0xfa, 0x8d, 0x76, 0xa1, 0xe8, 0x52, 0xe1, 0x70, 0x16, 0xc5, 0x2c,
	0x0c, 0xca, 0x59, 0xa5, 0x9a, 0x85, 0x50, 0x05, 0xf2, 0x67, 0xcc, 0xa3, 0xca, 0x33, 0xa7, 0xd4,
	0x53, 0x59, 0xea, 0x62, 0x4e, 0x89, 0x18, 0xf2, 0xeb, 0xf2, 0xb2, 0xd6, 0x25, 0x32, 0xfa, 0x11,
	0x64, 0xcf, 0x28, 0x2d, 0xaf, 0xec, 0x66, 0x1f, 0x14, 0x1f, 0xdd, 0xaf, 0xeb, 0x03, 0xd4, 0xe5,
	0x01, 0xea, 0xe6, 0x00, 0xf5, 0x56, 0xc8, 0x82, 0xe6, 0x57, 0x5e, 0x8c, 0x6b, 0x4b, 0xbf, 0xfa,
	0xa4, 0xf6, 0x60, 0xc0, 0xe2, 0xf3, 0x61, 0xbf, 0xee, 0x84, 0x7e, 0xc3, 0x9c, 0x56, 0xff, 0xf9,
	0xb2, 0x70, 0x2f, 0x1a, 0xf1, 0x75, 0x44, 0x85, 0x72, 0x10, 0x58, 0x3e, 0xf7, 0x49, 0xee, 0xcf,
	0x1f, 0xd5, 0xac, 0xbd, 0xdf, 0x59, 0x50, 0x7a, 0xa6, 0xf2, 0xde, 0x53, 0x01, 0xff, 0xdf, 0x4e,
	0xbe, 0x03, 0x2b, 0xc2, 0x39, 0xa7, 0x3e, 0x31, 0xe7, 0x36, 0x12, 0x7a, 0x0c, 0x1b, 0x42, 0xd5,
	0xc0, 0x76, 0x42, 0x97, 0xda, 0x43, 0xee, 0x95, 0x57, 0xa4, 0x41, 0x73, 0x6b, 0x32, 0xae, 0xad,
	0xe9, 0xf2, 0xb4, 0x42, 0x97, 0x9e, 0xe2, 0x63, 0xbc, 0x26, 0x52, 0x91, 0x7b, 0xe6, 0x44, 0xbf,
	0xb1, 0x00, 0x30, 0xb9, 0xc2, 0xf4, 0xc7, 0x43, 0x2a, 0x62, 0xf4, 0x2d, 0x28, 0xd2, 0x51, 0x4c,
	0x79, 0x40, 0x3c, 0x9b, 0xb9, 0xea, 0x54, 0xb9, 0xe6, 0x9b, 0x93, 0x71, 0x0d, 0xda, 0x06, 0xee,
	0x1c, 0xfe, 0x6d, 0x4e, 0xc2, 0x90, 0x38, 0x74, 0x5c, 0x74, 0x04, 0xeb, 0x2e, 0x89, 0x89, 0x6d,
	0x62, 0x62, 0xae, 0x4a, 0x41, 0xae, 0xb9, 0x3b, 0x19, 0xd7, 0x4a, 0x29, 0x61, 0xd4, 0x33, 0xe6,
	0x64, 0x5c, 0x72, 0x53, 0xc9, 0x95, 0xa9, 0x70, 0x88, 0xe7, 0x49, 0x4c, 0x65, 0xaa, 0x84, 0xa7,
	0xb2, 0x89, 0xfb, 0x27, 0x16, 0x14, 0x54, 0xdc, 0x51, 0xc8, 0xff, 0xe3, 0xb0, 0xdf, 0x80, 0x02,
	0x1d, 0xb1, 0x58, 0xe5, 0x50, 0x45, 0xbc, 0x86, 0xf3, 0x12, 0x90, 0xa9, 0x92, 0xc5, 0x9c, 0x89,
	0x23, 0x37, 0x13, 0xc3, 0x5f, 0xb2, 0xb0, 0x9a, 0x24, 0xee, 0x29, 0x6c, 0xea, 0x86, 0xb4, 0x75,
	0x41, 0xd3, 0x30, 0xde, 0x9a, 0x8c, 0x6b, 0xeb, 0xb3, 0xa4, 0x51, 0xa1, 0x2c, 0x20, 0x78, 0x3d,
	0x9c, 0x95, 0xe7, 0x33, 0x90, 0x99, 0xcf, 0x00, 0xda, 0x87, 0x6d, 0xae, 0x5f, 0x4b, 0x5d, 0xfb,
	0x92, 0x78, 0xcc, 0x25, 0x71, 0xc8, 0x45, 0x39, 0xbb, 0x9b, 0x7d, 0x50, 0xc0, 0x77, 0xa6, 0xba,
	0xe7, 0x53, 0x95, 0x3c, 0xa1, 0xcf, 0x02, 0xdb, 0x09, 0x87, 0x41, 0xac, 0xc8, 0x95, 0xc3, 0x79,
	0x9f, 0x05, 0x2d, 0x29, 0xa3, 0x2f, 0xc0, 0xba, 0xf1, 0xb1, 0xcf, 0x29, 0x1b, 0x9c, 0xc7, 0x8a,
	0x64, 0x59, 0xbc, 0x66, 0xd0, 0xef, 0x2a, 0x10, 0x7d, 0x1e, 0x4a, 0x89, 0x99, 0x1c, 0x25, 0x8a,
	0x68, 0x59, 0x5c, 0x34, 0xd8, 0x09, 0xf3, 0x29, 0xfa, 0x12, 0x14, 0x1c, 0x8f, 0xd1, 0x40, 0x1d,
	0x7f, 0x55, 0x11, 0xb1, 0x34, 0x19, 0xd7, 0xf2, 0x2d, 0x05, 0x76, 0x0e, 0x71, 0x5e, 0xab, 0x3b,
	0x2e, 0xfa, 0x36, 0x94, 0x38, 0xb9, 0xb2, 0x8d, 0xb7, 0x28, 0xe7, 0x55, 0xe3, 0xde, 0xad, 0x4f,
	0xc7, 0x5a, 0x3d, 0xa5, 0x65, 0x33, 0x27, 0x9b, 0x16, 0x17, 0xf9, 0x14, 0x11, 0xe8, 0x08, 0x8a,
	0xac, 0xef, 0xd8, 0xce, 0x39, 0x09, 0x02, 0xea, 0x95, 0x0b, 0xbb, 0xd6, 0x82, 0x7b, 0xa7, 0xd9,
	0x6a, 0x69, 0x65, 0x73, 0x5d, 0x32, 0x21, 0x95, 0x31, 0xb0, 0xbe, 0x63, 0x7e, 0xa3, 0x9a, 0xa4,
	0x0e, 0x75, 0x86, 0x31, 0xb5, 0x07, 0x44, 0x94, 0x41, 0xe5, 0x06, 0x0c, 0xf4, 0x2e, 0x11, 0xa6,
	0xd6, 0x3f, 0xb3, 0x60, 0xc5, 0x90, 0xed, 0x4d, 0x28, 0x4c, 0x93, 0x6e, 0xfa, 0x3e, 0x05, 0xd0,
	0x43, 0xd8, 0x62, 0x81, 0xdd, 0xa7, 0x67, 0x21, 0xa7, 0x36, 0xa7, 0x22, 0xf4, 0x2e, 0x35, 0xa7,
	0xf2, 0x78, 0x83, 0x05, 0x4d, 0x85, 0x63, 0x0d, 0xa3, 0x77, 0xa0, 0xa8, 0x73, 0x20, 0x9f, 0xab,
	0xeb, 0x57, 0x7c, 0xb4, 0xbd, 0x98, 0x02, 0xa9, 0x34, 0x19, 0x00, 0x9e, 0x00, 0x49, 0x5c, 0xbf,
	0xc8, 0xc2, 0x3d, 0x4d, 0x25, 0x93, 0x99, 0x2e, 0x71, 0x2e, 0x68, 0x2c, 0x7b, 0x6b, 0xbe, 0x1a,
	0xd6, 0xbf, 0xac, 0xc6, 0x4d, 0xf4, 0xcd, 0xfc, 0x97, 0xe8, 0xbb, 0xd0, 0xc0, 0x92, 0x8b, 0x44,
	0x5c, 0xcc, 0x73, 0x91, 0x88, 0x0b, 0xcd, 0xc5, 0x39, 0xa2, 0x2e, 0x2f, 0x10, 0xf5, 0x1c, 0x0a,
	0x67, 0x94, 0xda, 0x1e, 0xf3, 0x59, 0xfc, 0xbf, 0x98, 0xf4, 0xf9, 0x33, 0x4a, 0x8f, 0xe5, 0xc3,
	0x25, 0x2b, 0x22, 0x4e, 0x23, 0xc2, 0x35, 0x2b, 0x56, 0x35, 0x2b, 0x0c, 0xf4, 0x2e, 0x11, 0x8b,
	0xb4, 0xc9, 0x7f, 0x0a, 0x6d, 0x28, 0xec, 0xdd, 0x50, 0x9d, 0x03, 0xe7, 0x22, 0x08, 0xaf, 0x3c,
	0xea, 0x0e, 0xa8, 0x4f, 0x83, 0x18, 0x3d, 0x06, 0x48, 0x3a, 0x6b, 0x3a, 0x36, 0x2a, 0x93, 0x71,
	0xad, 0x60, 0xbc, 0x54, 0xca, 0x53, 0x01, 0x17, 0x8c, 0x75, 0xc7, 0x35, 0xaf, 0xf9, 0x6d, 0x06,
	0xca, 0xc9, 0x7b, 0x44, 0x14, 0x06, 0x82, 0xbe, 0x1e, 0x0d, 0xe6, 0x03, 0xc9, 0xdc, 0x22, 0x10,
	0x55, 0xd5, 0x40, 0x98, 0xc2, 0x65, 0x4d, 0x55, 0x03, 0xa1, 0x0b, 0xb7, 0x38, 0x3a, 0x72, 0xff,
	0x3c, 0x3a, 0x94, 0x89, 0x6a, 0x0b, 0x6d, 0xb2, 0x9c, 0x98, 0x28, 0x4c, 0x99, 0x7c, 0x47, 0xce,
	0x29, 0x6d, 0x22, 0x62, 0x12, 0x0f, 0x85, 0x1a, 0x41, 0xeb, 0x8f, 0xca, 0xb3, 0x1d, 0xa3, 0x0d,
	0x7a, 0x4a, 0x2f, 0x27, 0xd8, 0x8c, 0x28, 0x6f, 0x51, 0x4e, 0xc5, 0xd0, 0x8b, 0x55, 0x41, 0x4b,
	0xd8, 0x48, 0x26, 0x89, 0xbf, 0xcf, 0xca, 0x16, 0x97, 0xc0, 0x67, 0xaf, 0x73, 0xe6, 0x0b, 0xbb,
	0xf2, 0xda, 0x85, 0x5d, 0x7d, 0x45, 0x61, 0xf3, 0xaf, 0x2e, 0x6c, 0xe1, 0xdf, 0x29, 0x2c, 0xbc,
	0x6e, 0x61, 0x8b, 0x37, 0x14, 0x36, 0x82, 0x8d, 0xe9, 0x55, 0x68, 0x1c, 0xde, 0x80, 0x02, 0x13,
	0x36, 0x71, 0x62, 0x76, 0x49, 0x55, 0x81, 0xf3, 0x38, 0xcf, 0xc4, 0x81, 0x92, 0xd1, 0x13, 0x58,
	0x16, 0x2c, 0x70, 0xf4, 0xd8, 0x2e, 0x3e, 0xaa, 0xd4, 0xf5, 0xa2, 0x5c, 0x4f, 0x16, 0xe5, 0xfa,
	0x49, 0xb2, 0x49, 0x37, 0xf3, 0x72, 0xc6, 0x7c, 0xf0, 0x49, 0xcd, 0xc2, 0xda, 0xc5, 0xbc, 0xf1,
	0x00, 0x36, 0xf4, 0xb3, 0xa6, 0xef, 0x45, 0x65, 0x58, 0x25, 0xae, 0xcb, 0xa9, 0x10, 0xe6, 0xce,
	0x48, 0x44, 0xb9, 0x43, 0x46, 0xe1, 0x15, 0xe5, 0x9a, 0x36, 0x58, 0x0b, 0x7b, 0x2f, 0x72, 0xb0,
	0xd2, 0x25, 0x9c, 0xf8, 0x02, 0xed, 0xc3, 0x5d, 0x9f, 0x8c, 0xec, 0x99, 0xeb, 0xd2, 0x54, 0x43,
	0x4d, 0x0a, 0x8c, 0x7c, 0x32, 0x4a, 0xef, 0x4a, 0x5d, 0x97, 0x3d, 0x58, 0x93, 0x2e, 0x29, 0x5b,
	0xf4, 0xb3, 0x8b, 0x3e, 0x19, 0x1d, 0x24, 0x84, 0x79, 0x08, 0x5b, 0xd2, 0x26, 0x61, 0x97, 0x2d,
	0xd8, 0x7b, 0xd4, 0x74, 0xee, 0x86, 0x4f, 0x46, 0x2d, 0x83, 0xf7, 0xd8, 0x7b, 0x14, 0x35, 0x60,
	0x5b, 0x85, 0xa0, 0xee, 0x1e, 0x3b, 0x35, 0xd7, 0x24, 0x94, 0xcf, 0xd1, 0xd7, 0xd2, 0x61, 0xe2,
	0xf0, 0x35, 0xd8, 0xa1, 0xa3, 0x88, 0x71, 0x22, 0x57, 0x5b, 0xbb, 0xef, 0x85, 0xce, 0xc5, 0x1c,
	0x35, 0xb7, 0x53, 0x6d, 0x53, 0x2a, 0x75, 0x48, 0x6f, 0xc1, 0xba, 0x9c, 0xe3, 0x76, 0x78, 0x45,
	0x84, 0xaf, 0x06, 0xab, 0xa2, 0x2a, 0x2e, 0x49, 0xf4, 0x99, 0x04, 0xe5, 0xec, 0x7d, 0x0c, 0xf7,
	0x23, 0xca, 0xd3, 0xcd, 0x67, 0x9a, 0x95, 0x74, 0x54, 0xef, 0x44, 0x94, 0x4f, 0x73, 0x6f, 0x32,
	0x23, 0x5d, 0xdf, 0x06, 0x24, 0x88, 0x1f, 0x79, 0x2c, 0x18, 0xd8, 0x31, 0xbf, 0x36, 0x21, 0xe9,
	0xe9, 0xbd, 0x99, 0x68, 0x4e, 0xf8, 0xb5, 0x0e, 0xe7, 0x1b, 0x50, 0x36, 0xbd, 0xcd, 0xe9, 0x15,
	0xe1, 0xae, 0x1d, 0x51, 0xee, 0xd0, 0x20, 0x26, 0x03, 0x4d, 0xe3, 0x1c, 0x36, 0x5f, 0x67, 0x58,
	0xa9, 0xbb, 0x53, 0x2d, 0x7a, 0x02, 0xf7, 0x59, 0xa0, 0xe9, 0x65, 0x47, 0x34, 0x20, 0x5e, 0x7c,
	0x6d, 0xbb, 0x43, 0x7d, 0x5e, 0xb3, 0x63, 0xdc, 0x4b, 0x0c, 0xba, 0x5a, 0x7f, 0x68, 0xd4, 0xa8,
	0x0d, 0x77, 0xe4, 0x66, 0x93, 0x1c, 0x8a, 0x06, 0xa4, 0xef, 0x51, 0x57, 0x31, 0x3b, 0xdf, 0xbc,
	0x3b, 0x19, 0xd7, 0xb6, 0x3a, 0xcd, 0x96, 0x39, 0x53, 0x5b, 0x2b, 0xf1, 0x16, 0xeb, 0x3b, 0xf3,
	0xd0, 0x93, 0xfc, 0xcf, 0x3f, 0xaa, 0x2d, 0x29, 0x36, 0xbe, 0x03, 0xa8, 0x4b, 0x03, 0x97, 0x05,
	0x03, 0xd3, 0x44, 0xc7, 0x4c, 0xa8, 0x2b, 0x2e, 0x1d, 0x09, 0x92, 0x94, 0x59, 0x79, 0x83, 0x4d,
	0xfb, 0x3e, 0xb9, 0xc1, 0xbe, 0x07, 0x33, 0x9b, 0x13, 0xba, 0x07, 0xab, 0x8a, 0x01, 0xc9, 0x58,
	0xc4, 0x2b, 0x52, 0xec, 0xb8, 0xe8, 0x73, 0x00, 0x66, 0x15, 0x4b, 0x06, 0x60, 0x01, 0x17, 0x0c,
	0x32, 0xbd, 0xa6, 0x7e, 0x00, 0x77, 0x35, 0x4d, 0x28, 0x17, 0xdd, 0x99, 0x02, 0xbd, 0x62, 0xa5,
	0xda, 0x83, 0x02, 0x4f, 0xdc, 0xca, 0x19, 0xb9, 0xe4, 0x9a, 0x75, 0x28, 0x85, 0xf7, 0xfe, 0x68,
	0xc1, 0x1d, 0x93, 0x80, 0xe7, 0x94, 0xb3, 0x33, 0xe6, 0xe8, 0x64, 0x7e, 0x11, 0xf2, 0xce, 0x39,
	0x61, 0x41, 0x3a, 0xc8, 0x8b, 0x93, 0x71, 0x6d, 0xb5, 0x25, 0xb1, 0xce, 0x21, 0x5e, 0x55, 0xca,
	0x8e, 0x3b, 0x1f, 0x41, 0x66, 0x31, 0x82, 0xf9, 0xf1, 0x99, 0xbd, 0xcd, 0xf8, 0x5c, 0xf8, 0x34,
	0xc9, 0xdd, 0xee, 0xd3, 0xc4, 0x24, 0xee, 0xd7, 0x16, 0x14, 0xbb, 0x9c, 0x39, 0xd4, 0xdc, 0x4f,
	0xf2, 0x73, 0xf0, 0xda, 0xef, 0x87, 0x5e, 0x52, 0x05, 0x2d, 0xa1, 0x2a, 0x80, 0x3f, 0xf4, 0x62,
	0x16, 0x79, 0x6c, 0x3a, 0x4f, 0x66, 0x10, 0xb4, 0x0e, 0x99, 0x68, 0x64, 0x7a, 0x3c, 0x13, 0x8d,
	0x16, 0xce, 0x95, 0xbb, 0xcd, 0xb9, 0x5e, 0x7d, 0x5f, 0x3f, 0xfc, 0xbb, 0x05, 0x6b, 0x73, 0x63,
	0x1b, 0x7d, 0x13, 0x6a, 0xb8, 0xdd, 0x7b, 0x76, 0xfc, 0xbc, 0x6d, 0xf7, 0x4e, 0x0e, 0x4e, 0x4e,
	0x7b, 0xf6, 0xb3, 0x6e, 0xfb, 0xa9, 0x7d, 0xfa, 0xb4, 0xd7, 0x6d, 0xb7, 0x3a, 0x47, 0x9d, 0xf6,
	0xe1, 0xe6, 0x52, 0xe5, 0xde, 0xfb, 0x1f, 0xee, 0xde, 0xb9, 0xc1, 0x0c, 0x7d, 0x1d, 0x76, 0x16,
	0xe0, 0xde, 0x69, 0xab, 0xd5, 0xee, 0xf5, 0x36, 0xad, 0x4a, 0xe5, 0xfd, 0x0f, 0x77, 0x3f, 0x45,
	0x7b, 0x83, 0xdf, 0xd1, 0x41, 0xe7, 0xf8, 0x14, 0xb7, 0x37, 0x33, 0x37, 0xfa, 0x19, 0xed, 0x0d,
	0x7e, 0xed, 0xef, 0x77, 0x3b, 0xb8, 0x7d, 0xb8, 0x99, 0xbd, 0xd1, 0xcf, 0x68, 0x2b, 0xb9, 0x9f,
	0xfe, 0xb2, 0xba, 0xd4, 0x3c, 0x7a, 0x31, 0xa9, 0x5a, 0x1f, 0x4f, 0xaa, 0xd6, 0x9f, 0x26, 0x55,
	0xeb, 0x83, 0x97, 0xd5, 0xa5, 0x8f, 0x5f, 0x56, 0x97, 0xfe, 0xf0, 0xb2, 0xba, 0xf4, 0xc3, 0xb7,
	0x67, 0x16, 0xd2, 0x3e, 0x09, 0x5c, 0x75, 0xaf, 0x38, 0xa1, 0xd7, 0x50, 0x64, 0x6c, 0x8c, 0xcc,
	0x3f, 0x7a, 0xf4, 0x6a, 0xda, 0x5f, 0x51, 0xea, 0xaf, 0xfe, 0x23, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0x4d, 0x97, 0x94, 0x09, 0x12, 0x00, 0x00,
}

func (this *OracleRequestPacketData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OracleRequestPacketData)
	if !ok {
		that2, ok := that.(OracleRequestPacketData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClientID != that1.ClientID {
		return false
	}
	if this.OracleScriptID != that1.OracleScriptID {
		return false
	}
	if !bytes.Equal(this.Calldata, that1.Calldata) {
		return false
	}
	if this.AskCount != that1.AskCount {
		return false
	}
	if this.MinCount != that1.MinCount {
		return false
	}
	if len(this.FeeLimit) != len(that1.FeeLimit) {
		return false
	}
	for i := range this.FeeLimit {
		if !this.FeeLimit[i].Equal(&that1.FeeLimit[i]) {
			return false
		}
	}
	if this.PrepareGas != that1.PrepareGas {
		return false
	}
	if this.ExecuteGas != that1.ExecuteGas {
		return false
	}
	return true
}
func (this *OracleRequestPacketAcknowledgement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OracleRequestPacketAcknowledgement)
	if !ok {
		that2, ok := that.(OracleRequestPacketAcknowledgement)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RequestID != that1.RequestID {
		return false
	}
	return true
}
func (this *OracleResponsePacketData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OracleResponsePacketData)
	if !ok {
		that2, ok := that.(OracleResponsePacketData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClientID != that1.ClientID {
		return false
	}
	if this.RequestID != that1.RequestID {
		return false
	}
	if this.AnsCount != that1.AnsCount {
		return false
	}
	if this.RequestTime != that1.RequestTime {
		return false
	}
	if this.ResolveTime != that1.ResolveTime {
		return false
	}
	if this.ResolveStatus != that1.ResolveStatus {
		return false
	}
	if !bytes.Equal(this.Result, that1.Result) {
		return false
	}
	return true
}

func (m *OracleRequestPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleRequestPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleRequestPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExecuteGas != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ExecuteGas))
		i--
		dAtA[i] = 0x40
	}
	if m.PrepareGas != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.PrepareGas))
		i--
		dAtA[i] = 0x38
	}
	if len(m.FeeLimit) > 0 {
		for iNdEx := len(m.FeeLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MinCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.MinCount))
		i--
		dAtA[i] = 0x28
	}
	if m.AskCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.AskCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Calldata) > 0 {
		i -= len(m.Calldata)
		copy(dAtA[i:], m.Calldata)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Calldata)))
		i--
		dAtA[i] = 0x1a
	}
	if m.OracleScriptID != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.OracleScriptID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OracleRequestPacketAcknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleRequestPacketAcknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleRequestPacketAcknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestID != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OracleResponsePacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleResponsePacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleResponsePacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ResolveStatus != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ResolveStatus))
		i--
		dAtA[i] = 0x30
	}
	if m.ResolveTime != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ResolveTime))
		i--
		dAtA[i] = 0x28
	}
	if m.RequestTime != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.RequestTime))
		i--
		dAtA[i] = 0x20
	}
	if m.AnsCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.AnsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.RequestID != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *OracleRequestPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.OracleScriptID != 0 {
		n += 1 + sovOracle(uint64(m.OracleScriptID))
	}
	l = len(m.Calldata)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.AskCount != 0 {
		n += 1 + sovOracle(uint64(m.AskCount))
	}
	if m.MinCount != 0 {
		n += 1 + sovOracle(uint64(m.MinCount))
	}
	if len(m.FeeLimit) > 0 {
		for _, e := range m.FeeLimit {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	if m.PrepareGas != 0 {
		n += 1 + sovOracle(uint64(m.PrepareGas))
	}
	if m.ExecuteGas != 0 {
		n += 1 + sovOracle(uint64(m.ExecuteGas))
	}
	return n
}

func (m *OracleRequestPacketAcknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestID != 0 {
		n += 1 + sovOracle(uint64(m.RequestID))
	}
	return n
}

func (m *OracleResponsePacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.RequestID != 0 {
		n += 1 + sovOracle(uint64(m.RequestID))
	}
	if m.AnsCount != 0 {
		n += 1 + sovOracle(uint64(m.AnsCount))
	}
	if m.RequestTime != 0 {
		n += 1 + sovOracle(uint64(m.RequestTime))
	}
	if m.ResolveTime != 0 {
		n += 1 + sovOracle(uint64(m.ResolveTime))
	}
	if m.ResolveStatus != 0 {
		n += 1 + sovOracle(uint64(m.ResolveStatus))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *OracleRequestPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OracleRequestPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleRequestPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScriptID", wireType)
			}
			m.OracleScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Calldata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Calldata = append(m.Calldata[:0], dAtA[iNdEx:postIndex]...)
			if m.Calldata == nil {
				m.Calldata = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
			}
			m.AskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
			}
			m.MinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeLimit = append(m.FeeLimit, types.Coin{})
			if err := m.FeeLimit[len(m.FeeLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGas", wireType)
			}
			m.PrepareGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGas", wireType)
			}
			m.ExecuteGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OracleRequestPacketAcknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OracleRequestPacketAcknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleRequestPacketAcknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OracleResponsePacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OracleResponsePacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleResponsePacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnsCount", wireType)
			}
			m.AnsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTime", wireType)
			}
			m.RequestTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveTime", wireType)
			}
			m.ResolveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveStatus", wireType)
			}
			m.ResolveStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveStatus |= ResolveStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
