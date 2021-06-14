// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/subscriberauth.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// ErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by EPC. Diameter Base Protocol errors are reflected in gRPC status code
type M5GErrorCode int32

const (
	M5GErrorCode_UNDEFINED M5GErrorCode = 0
	// Default success code
	M5GErrorCode_MULTI_ROUND_AUTH        M5GErrorCode = 1001
	M5GErrorCode_SUCCESS                 M5GErrorCode = 2001
	M5GErrorCode_LIMITED_SUCCESS         M5GErrorCode = 2002
	M5GErrorCode_COMMAND_UNSUPORTED      M5GErrorCode = 3001
	M5GErrorCode_UNABLE_TO_DELIVER       M5GErrorCode = 3002
	M5GErrorCode_REALM_NOT_SERVED        M5GErrorCode = 3003
	M5GErrorCode_TOO_BUSY                M5GErrorCode = 3004
	M5GErrorCode_LOOP_DETECTED           M5GErrorCode = 3005
	M5GErrorCode_REDIRECT_INDICATION     M5GErrorCode = 3006
	M5GErrorCode_APPLICATION_UNSUPPORTED M5GErrorCode = 3007
	M5GErrorCode_INVALIDH_DR_BITS        M5GErrorCode = 3008
	M5GErrorCode_INVALID_AVP_BITS        M5GErrorCode = 3009
	M5GErrorCode_UNKNOWN_PEER            M5GErrorCode = 3010
	M5GErrorCode_AUTHENTICATION_REJECTED M5GErrorCode = 4001
	M5GErrorCode_OUT_OF_SPACE            M5GErrorCode = 4002
	M5GErrorCode_ELECTION_LOST           M5GErrorCode = 4003
	M5GErrorCode_AUTHORIZATION_REJECTED  M5GErrorCode = 5003
	// Permanent Failures 7.4.3
	M5GErrorCode_USER_UNKNOWN             M5GErrorCode = 5001
	M5GErrorCode_UNKNOWN_SESSION_ID       M5GErrorCode = 5002
	M5GErrorCode_UNKNOWN_EPS_SUBSCRIPTION M5GErrorCode = 5420
	M5GErrorCode_RAT_NOT_ALLOWED          M5GErrorCode = 5421
	M5GErrorCode_ROAMING_NOT_ALLOWED      M5GErrorCode = 5004
	M5GErrorCode_EQUIPMENT_UNKNOWN        M5GErrorCode = 5422
	M5GErrorCode_UNKOWN_SERVING_NODE      M5GErrorCode = 5423
	// Transient Failures 7.4.4
	M5GErrorCode_AUTHENTICATION_DATA_UNAVAILABLE M5GErrorCode = 4181
)

var M5GErrorCode_name = map[int32]string{
	0:    "UNDEFINED",
	1001: "MULTI_ROUND_AUTH",
	2001: "SUCCESS",
	2002: "LIMITED_SUCCESS",
	3001: "COMMAND_UNSUPORTED",
	3002: "UNABLE_TO_DELIVER",
	3003: "REALM_NOT_SERVED",
	3004: "TOO_BUSY",
	3005: "LOOP_DETECTED",
	3006: "REDIRECT_INDICATION",
	3007: "APPLICATION_UNSUPPORTED",
	3008: "INVALIDH_DR_BITS",
	3009: "INVALID_AVP_BITS",
	3010: "UNKNOWN_PEER",
	4001: "AUTHENTICATION_REJECTED",
	4002: "OUT_OF_SPACE",
	4003: "ELECTION_LOST",
	5003: "AUTHORIZATION_REJECTED",
	5001: "USER_UNKNOWN",
	5002: "UNKNOWN_SESSION_ID",
	5420: "UNKNOWN_EPS_SUBSCRIPTION",
	5421: "RAT_NOT_ALLOWED",
	5004: "ROAMING_NOT_ALLOWED",
	5422: "EQUIPMENT_UNKNOWN",
	5423: "UNKOWN_SERVING_NODE",
	4181: "AUTHENTICATION_DATA_UNAVAILABLE",
}

var M5GErrorCode_value = map[string]int32{
	"UNDEFINED":                       0,
	"MULTI_ROUND_AUTH":                1001,
	"SUCCESS":                         2001,
	"LIMITED_SUCCESS":                 2002,
	"COMMAND_UNSUPORTED":              3001,
	"UNABLE_TO_DELIVER":               3002,
	"REALM_NOT_SERVED":                3003,
	"TOO_BUSY":                        3004,
	"LOOP_DETECTED":                   3005,
	"REDIRECT_INDICATION":             3006,
	"APPLICATION_UNSUPPORTED":         3007,
	"INVALIDH_DR_BITS":                3008,
	"INVALID_AVP_BITS":                3009,
	"UNKNOWN_PEER":                    3010,
	"AUTHENTICATION_REJECTED":         4001,
	"OUT_OF_SPACE":                    4002,
	"ELECTION_LOST":                   4003,
	"AUTHORIZATION_REJECTED":          5003,
	"USER_UNKNOWN":                    5001,
	"UNKNOWN_SESSION_ID":              5002,
	"UNKNOWN_EPS_SUBSCRIPTION":        5420,
	"RAT_NOT_ALLOWED":                 5421,
	"ROAMING_NOT_ALLOWED":             5004,
	"EQUIPMENT_UNKNOWN":               5422,
	"UNKOWN_SERVING_NODE":             5423,
	"AUTHENTICATION_DATA_UNAVAILABLE": 4181,
}

func (x M5GErrorCode) String() string {
	return proto.EnumName(M5GErrorCode_name, int32(x))
}

func (M5GErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a742c81ad67f98d, []int{0}
}

// Authentication Information Request (Section 7.2.5)
type M5GAuthenticationInformationRequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// Visted site identifier
	VisitedPlmn []byte `protobuf:"bytes,2,opt,name=visited_plmn,json=visitedPlmn,proto3" json:"visited_plmn,omitempty"`
	// Number of vectors to request in response
	NumRequestedEutranVectors uint32 `protobuf:"varint,3,opt,name=num_requested_eutran_vectors,json=numRequestedEutranVectors,proto3" json:"num_requested_eutran_vectors,omitempty"`
	// Indicates to the HSS the values are requested for immediate attach
	ImmediateResponsePreferred bool `protobuf:"varint,4,opt,name=immediate_response_preferred,json=immediateResponsePreferred,proto3" json:"immediate_response_preferred,omitempty"`
	// Concatenation of RAND and AUTS in the case of a resync attach case
	ResyncInfo []byte `protobuf:"bytes,5,opt,name=resync_info,json=resyncInfo,proto3" json:"resync_info,omitempty"`
	// Number of UTRAN/GERAN vectors to request in response (optional)
	NumRequestedUtranGeranVectors uint32 `protobuf:"varint,6,opt,name=num_requested_utran_geran_vectors,json=numRequestedUtranGeranVectors,proto3" json:"num_requested_utran_geran_vectors,omitempty"`
	// UTRAN/GERAN Resync Info (optional)
	UtranGeranResyncInfo []byte   `protobuf:"bytes,7,opt,name=utran_geran_resync_info,json=utranGeranResyncInfo,proto3" json:"utran_geran_resync_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M5GAuthenticationInformationRequest) Reset()         { *m = M5GAuthenticationInformationRequest{} }
func (m *M5GAuthenticationInformationRequest) String() string { return proto.CompactTextString(m) }
func (*M5GAuthenticationInformationRequest) ProtoMessage()    {}
func (*M5GAuthenticationInformationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a742c81ad67f98d, []int{0}
}

func (m *M5GAuthenticationInformationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M5GAuthenticationInformationRequest.Unmarshal(m, b)
}
func (m *M5GAuthenticationInformationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M5GAuthenticationInformationRequest.Marshal(b, m, deterministic)
}
func (m *M5GAuthenticationInformationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M5GAuthenticationInformationRequest.Merge(m, src)
}
func (m *M5GAuthenticationInformationRequest) XXX_Size() int {
	return xxx_messageInfo_M5GAuthenticationInformationRequest.Size(m)
}
func (m *M5GAuthenticationInformationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_M5GAuthenticationInformationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_M5GAuthenticationInformationRequest proto.InternalMessageInfo

func (m *M5GAuthenticationInformationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *M5GAuthenticationInformationRequest) GetVisitedPlmn() []byte {
	if m != nil {
		return m.VisitedPlmn
	}
	return nil
}

func (m *M5GAuthenticationInformationRequest) GetNumRequestedEutranVectors() uint32 {
	if m != nil {
		return m.NumRequestedEutranVectors
	}
	return 0
}

func (m *M5GAuthenticationInformationRequest) GetImmediateResponsePreferred() bool {
	if m != nil {
		return m.ImmediateResponsePreferred
	}
	return false
}

func (m *M5GAuthenticationInformationRequest) GetResyncInfo() []byte {
	if m != nil {
		return m.ResyncInfo
	}
	return nil
}

func (m *M5GAuthenticationInformationRequest) GetNumRequestedUtranGeranVectors() uint32 {
	if m != nil {
		return m.NumRequestedUtranGeranVectors
	}
	return 0
}

func (m *M5GAuthenticationInformationRequest) GetUtranGeranResyncInfo() []byte {
	if m != nil {
		return m.UtranGeranResyncInfo
	}
	return nil
}

// Authentication Information Answer (Section 7.2.6)
type M5GAuthenticationInformationAnswer struct {
	// EPC error code on failure
	ErrorCode M5GErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=magma.lte.M5GErrorCode" json:"error_code,omitempty"`
	// Authentication vectors matching the requested number
	EutranVectors        []*M5GAuthenticationInformationAnswer_EUTRANVector `protobuf:"bytes,2,rep,name=eutran_vectors,json=eutranVectors,proto3" json:"eutran_vectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *M5GAuthenticationInformationAnswer) Reset()         { *m = M5GAuthenticationInformationAnswer{} }
func (m *M5GAuthenticationInformationAnswer) String() string { return proto.CompactTextString(m) }
func (*M5GAuthenticationInformationAnswer) ProtoMessage()    {}
func (*M5GAuthenticationInformationAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a742c81ad67f98d, []int{1}
}

func (m *M5GAuthenticationInformationAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M5GAuthenticationInformationAnswer.Unmarshal(m, b)
}
func (m *M5GAuthenticationInformationAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M5GAuthenticationInformationAnswer.Marshal(b, m, deterministic)
}
func (m *M5GAuthenticationInformationAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M5GAuthenticationInformationAnswer.Merge(m, src)
}
func (m *M5GAuthenticationInformationAnswer) XXX_Size() int {
	return xxx_messageInfo_M5GAuthenticationInformationAnswer.Size(m)
}
func (m *M5GAuthenticationInformationAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_M5GAuthenticationInformationAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_M5GAuthenticationInformationAnswer proto.InternalMessageInfo

func (m *M5GAuthenticationInformationAnswer) GetErrorCode() M5GErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return M5GErrorCode_UNDEFINED
}

func (m *M5GAuthenticationInformationAnswer) GetEutranVectors() []*M5GAuthenticationInformationAnswer_EUTRANVector {
	if m != nil {
		return m.EutranVectors
	}
	return nil
}

// 3GPP TS 29.272, 7.3.18 E-UTRAN-Vector
// For details about fields read 3GPP 33.401
type M5GAuthenticationInformationAnswer_EUTRANVector struct {
	Rand                 []byte   `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Xres                 []byte   `protobuf:"bytes,2,opt,name=xres,proto3" json:"xres,omitempty"`
	Autn                 []byte   `protobuf:"bytes,3,opt,name=autn,proto3" json:"autn,omitempty"`
	Kasme                []byte   `protobuf:"bytes,4,opt,name=kasme,proto3" json:"kasme,omitempty"`
	Ck                   []byte   `protobuf:"bytes,5,opt,name=ck,proto3" json:"ck,omitempty"`
	Ik                   []byte   `protobuf:"bytes,6,opt,name=ik,proto3" json:"ik,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) Reset() {
	*m = M5GAuthenticationInformationAnswer_EUTRANVector{}
}
func (m *M5GAuthenticationInformationAnswer_EUTRANVector) String() string {
	return proto.CompactTextString(m)
}
func (*M5GAuthenticationInformationAnswer_EUTRANVector) ProtoMessage() {}
func (*M5GAuthenticationInformationAnswer_EUTRANVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a742c81ad67f98d, []int{1, 0}
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M5GAuthenticationInformationAnswer_EUTRANVector.Unmarshal(m, b)
}
func (m *M5GAuthenticationInformationAnswer_EUTRANVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M5GAuthenticationInformationAnswer_EUTRANVector.Marshal(b, m, deterministic)
}
func (m *M5GAuthenticationInformationAnswer_EUTRANVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M5GAuthenticationInformationAnswer_EUTRANVector.Merge(m, src)
}
func (m *M5GAuthenticationInformationAnswer_EUTRANVector) XXX_Size() int {
	return xxx_messageInfo_M5GAuthenticationInformationAnswer_EUTRANVector.Size(m)
}
func (m *M5GAuthenticationInformationAnswer_EUTRANVector) XXX_DiscardUnknown() {
	xxx_messageInfo_M5GAuthenticationInformationAnswer_EUTRANVector.DiscardUnknown(m)
}

var xxx_messageInfo_M5GAuthenticationInformationAnswer_EUTRANVector proto.InternalMessageInfo

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) GetRand() []byte {
	if m != nil {
		return m.Rand
	}
	return nil
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) GetXres() []byte {
	if m != nil {
		return m.Xres
	}
	return nil
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) GetAutn() []byte {
	if m != nil {
		return m.Autn
	}
	return nil
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) GetKasme() []byte {
	if m != nil {
		return m.Kasme
	}
	return nil
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) GetCk() []byte {
	if m != nil {
		return m.Ck
	}
	return nil
}

func (m *M5GAuthenticationInformationAnswer_EUTRANVector) GetIk() []byte {
	if m != nil {
		return m.Ik
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.lte.M5GErrorCode", M5GErrorCode_name, M5GErrorCode_value)
	proto.RegisterType((*M5GAuthenticationInformationRequest)(nil), "magma.lte.M5GAuthenticationInformationRequest")
	proto.RegisterType((*M5GAuthenticationInformationAnswer)(nil), "magma.lte.M5GAuthenticationInformationAnswer")
	proto.RegisterType((*M5GAuthenticationInformationAnswer_EUTRANVector)(nil), "magma.lte.M5GAuthenticationInformationAnswer.EUTRANVector")
}

func init() { proto.RegisterFile("lte/protos/subscriberauth.proto", fileDescriptor_6a742c81ad67f98d) }

var fileDescriptor_6a742c81ad67f98d = []byte{
	// 923 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4f, 0x73, 0xdb, 0x44,
	0x14, 0xaf, 0xdd, 0xb4, 0x89, 0x37, 0x72, 0xa2, 0x2e, 0xa1, 0x72, 0x93, 0x74, 0xec, 0x06, 0x66,
	0xe2, 0xc9, 0x0c, 0xce, 0x4c, 0x98, 0x70, 0xe0, 0x02, 0xb2, 0xb4, 0x75, 0x44, 0xa5, 0x95, 0x58,
	0xad, 0xdc, 0xa1, 0x97, 0x1d, 0xc5, 0xde, 0xa4, 0x1a, 0x5b, 0x52, 0x58, 0x49, 0x05, 0x4e, 0x70,
	0xa5, 0xf0, 0x19, 0x98, 0x01, 0xae, 0xfc, 0x39, 0xf3, 0x1f, 0xfa, 0x0d, 0x60, 0x86, 0x0f, 0xc1,
	0x91, 0x6f, 0xc0, 0x68, 0x65, 0x7b, 0x9c, 0x1c, 0x0a, 0x9c, 0xf4, 0xf4, 0xdb, 0xf7, 0xde, 0xef,
	0xa7, 0xf7, 0xd3, 0x93, 0x40, 0x7b, 0x9a, 0xf3, 0xc3, 0x0b, 0x91, 0xe6, 0x69, 0x76, 0x98, 0x15,
	0xa7, 0xd9, 0x48, 0x44, 0xa7, 0x5c, 0x84, 0x45, 0xfe, 0xb8, 0x27, 0x51, 0xd8, 0x88, 0xc3, 0xf3,
	0x38, 0xec, 0x4d, 0x73, 0xbe, 0xf7, 0xf4, 0x3a, 0x78, 0xc9, 0x39, 0x1e, 0xe8, 0x45, 0xfe, 0x98,
	0x27, 0x79, 0x34, 0x0a, 0xf3, 0x28, 0x4d, 0xac, 0xe4, 0x2c, 0x15, 0xb1, 0x0c, 0x09, 0x7f, 0xb7,
	0xe0, 0x59, 0x0e, 0x77, 0x40, 0xa3, 0xc8, 0xb8, 0x60, 0x49, 0x18, 0xf3, 0x56, 0xad, 0x53, 0xeb,
	0x36, 0xc8, 0x5a, 0x09, 0xe0, 0x30, 0xe6, 0xf0, 0x1e, 0x50, 0x9e, 0x44, 0x59, 0x94, 0xf3, 0x31,
	0xbb, 0x98, 0xc6, 0x49, 0xab, 0xde, 0xa9, 0x75, 0x15, 0xb2, 0x3e, 0xc3, 0xbc, 0x69, 0x9c, 0xc0,
	0x37, 0xc0, 0x6e, 0x52, 0xc4, 0x4c, 0x54, 0xed, 0xf8, 0x98, 0xf1, 0x22, 0x17, 0x61, 0xc2, 0x9e,
	0xf0, 0x51, 0x9e, 0x8a, 0xac, 0x75, 0xbd, 0x53, 0xeb, 0x36, 0xc9, 0x9d, 0xa4, 0x88, 0xc9, 0x3c,
	0x05, 0xc9, 0x8c, 0x61, 0x95, 0x00, 0xdf, 0x04, 0xbb, 0x51, 0x1c, 0xf3, 0x71, 0x14, 0xe6, 0x9c,
	0x09, 0x9e, 0x5d, 0xa4, 0x49, 0xc6, 0xd9, 0x85, 0xe0, 0x67, 0x5c, 0x08, 0x3e, 0x6e, 0xad, 0x74,
	0x6a, 0xdd, 0x35, 0xb2, 0xbd, 0xc8, 0x21, 0xb3, 0x14, 0x6f, 0x9e, 0x01, 0xdb, 0x60, 0x5d, 0xf0,
	0xec, 0x83, 0x64, 0xc4, 0xa2, 0xe4, 0x2c, 0x6d, 0xdd, 0x90, 0x22, 0x41, 0x05, 0x95, 0x4f, 0x0c,
	0x4f, 0xc0, 0xbd, 0xcb, 0x1a, 0x2b, 0x89, 0xe7, 0x7c, 0x59, 0xe8, 0x4d, 0x29, 0xf4, 0xee, 0xb2,
	0xd0, 0xa0, 0x4c, 0x1b, 0xf0, 0x25, 0xb1, 0xc7, 0x40, 0x5b, 0xae, 0x5d, 0xa6, 0x5d, 0x95, 0xb4,
	0x5b, 0xc5, 0xa2, 0x86, 0x2c, 0x04, 0xec, 0x3d, 0xab, 0x83, 0xbd, 0xe7, 0x99, 0xa1, 0x27, 0xd9,
	0x7b, 0x5c, 0xc0, 0xd7, 0x00, 0xe0, 0x42, 0xa4, 0x82, 0x8d, 0xd2, 0x71, 0x65, 0xc6, 0xc6, 0x91,
	0xd6, 0x5b, 0x78, 0xda, 0x73, 0x8e, 0x07, 0xa8, 0x3c, 0x37, 0xd2, 0x31, 0x27, 0x0d, 0x3e, 0x0f,
	0x61, 0x08, 0x36, 0xae, 0x4c, 0xbd, 0xde, 0xb9, 0xde, 0x5d, 0x3f, 0x7a, 0xfd, 0x72, 0xed, 0xbf,
	0xd0, 0xf7, 0x50, 0x40, 0x89, 0x8e, 0xab, 0x47, 0x25, 0x4d, 0xbe, 0xec, 0xd2, 0xf6, 0x47, 0x35,
	0xa0, 0x2c, 0x9f, 0x43, 0x08, 0x56, 0x44, 0x98, 0x8c, 0xa5, 0x4a, 0x85, 0xc8, 0xb8, 0xc4, 0xde,
	0x17, 0x3c, 0x9b, 0xbd, 0x26, 0x32, 0x2e, 0xb1, 0xb0, 0xc8, 0x13, 0xf9, 0x1e, 0x28, 0x44, 0xc6,
	0x70, 0x0b, 0xdc, 0x98, 0x84, 0x59, 0xcc, 0xa5, 0xb7, 0x0a, 0xa9, 0x6e, 0xe0, 0x06, 0xa8, 0x8f,
	0x26, 0x33, 0xf7, 0xea, 0xa3, 0x49, 0x79, 0x1f, 0x4d, 0xa4, 0x2d, 0x0a, 0xa9, 0x47, 0x93, 0x83,
	0xbf, 0x57, 0x80, 0xb2, 0x3c, 0x01, 0xd8, 0x04, 0x8d, 0x00, 0x9b, 0xe8, 0xbe, 0x85, 0x91, 0xa9,
	0x5e, 0x83, 0x2f, 0x02, 0xd5, 0x09, 0x6c, 0x6a, 0x31, 0xe2, 0x06, 0xd8, 0x64, 0x7a, 0x40, 0x4f,
	0xd4, 0xbf, 0x56, 0xa1, 0x02, 0x56, 0xfd, 0xc0, 0x30, 0x90, 0xef, 0xab, 0xbf, 0x6f, 0xc2, 0x2d,
	0xb0, 0x69, 0x5b, 0x8e, 0x45, 0x91, 0xc9, 0xe6, 0xe8, 0x1f, 0x9b, 0x50, 0x03, 0xd0, 0x70, 0x1d,
	0x47, 0xc7, 0x26, 0x0b, 0xb0, 0x1f, 0x78, 0x2e, 0xa1, 0xc8, 0x54, 0xbf, 0xd3, 0xe0, 0x6d, 0x70,
	0x2b, 0xc0, 0x7a, 0xdf, 0x46, 0x8c, 0xba, 0xcc, 0x44, 0xb6, 0x35, 0x44, 0x44, 0xfd, 0x5e, 0x2b,
	0xb9, 0x08, 0xd2, 0x6d, 0x87, 0x61, 0x97, 0x32, 0x1f, 0x91, 0x21, 0x32, 0xd5, 0x1f, 0x34, 0xd8,
	0x04, 0x6b, 0xd4, 0x75, 0x59, 0x3f, 0xf0, 0xdf, 0x51, 0x7f, 0xd4, 0x20, 0x04, 0x4d, 0xdb, 0x75,
	0x3d, 0x66, 0x22, 0x8a, 0x8c, 0xb2, 0xe3, 0x4f, 0x1a, 0x6c, 0x81, 0x17, 0x08, 0x32, 0x2d, 0x82,
	0x0c, 0xca, 0x2c, 0x6c, 0x5a, 0x86, 0x4e, 0x2d, 0x17, 0xab, 0x3f, 0x6b, 0x70, 0x17, 0x68, 0xba,
	0xe7, 0xd9, 0x33, 0xa4, 0x12, 0x32, 0x53, 0xf2, 0x8b, 0x64, 0xb4, 0xf0, 0x50, 0xb7, 0x2d, 0xf3,
	0x84, 0x99, 0x84, 0xf5, 0x2d, 0xea, 0xab, 0xbf, 0x2e, 0xc3, 0x4c, 0x1f, 0x7a, 0x15, 0xfc, 0x9b,
	0x06, 0x6f, 0x01, 0x25, 0xc0, 0x0f, 0xb0, 0xfb, 0x10, 0x33, 0x0f, 0x21, 0xa2, 0x3e, 0xab, 0xda,
	0x07, 0xf4, 0x04, 0x61, 0x3a, 0x67, 0x20, 0xe8, 0xad, 0x4a, 0xd6, 0xe7, 0xed, 0xb2, 0xc0, 0x0d,
	0x28, 0x73, 0xef, 0x33, 0xdf, 0xd3, 0x0d, 0xa4, 0x7e, 0xd1, 0x2e, 0xd5, 0x23, 0x1b, 0x19, 0x32,
	0xd5, 0x76, 0x7d, 0xaa, 0x7e, 0xd9, 0x86, 0x3b, 0xe0, 0x76, 0xd9, 0xc4, 0x25, 0xd6, 0xa3, 0x2b,
	0x3d, 0x3e, 0xd9, 0x97, 0xa4, 0x3e, 0x22, 0x6c, 0xc6, 0xac, 0x7e, 0xbc, 0x5f, 0x0e, 0x76, 0xae,
	0xc3, 0x47, 0xbe, 0x5f, 0x56, 0x58, 0xa6, 0xfa, 0x74, 0x1f, 0xde, 0x05, 0xad, 0xf9, 0x01, 0xf2,
	0x7c, 0xe6, 0x07, 0x7d, 0xdf, 0x20, 0x96, 0x27, 0x67, 0xf1, 0xd5, 0x41, 0x69, 0x13, 0xd1, 0xa9,
	0x9c, 0xae, 0x6e, 0xdb, 0xee, 0x43, 0x64, 0xaa, 0x5f, 0x1f, 0xc8, 0xd9, 0xb9, 0xba, 0x63, 0xe1,
	0xc1, 0xa5, 0x93, 0x4f, 0xf7, 0x4b, 0x9f, 0xd0, 0xdb, 0x81, 0xe5, 0x39, 0x08, 0xd3, 0x05, 0xff,
	0x37, 0xb2, 0x22, 0xc0, 0x0f, 0x2a, 0x7a, 0x32, 0xac, 0x0a, 0x4d, 0xa4, 0x7e, 0x7b, 0x00, 0x5f,
	0x06, 0xed, 0x2b, 0xe3, 0x30, 0x75, 0xaa, 0xb3, 0x00, 0xeb, 0x43, 0xdd, 0xb2, 0x4b, 0xcb, 0xd5,
	0x3f, 0x3b, 0x47, 0x9f, 0xd5, 0xc0, 0x8e, 0x73, 0x3c, 0xf0, 0x17, 0x1f, 0xdb, 0xcb, 0x3b, 0x04,
	0x3f, 0x04, 0xbb, 0xcf, 0x5b, 0x2c, 0xd8, 0xfb, 0x8f, 0x1b, 0x38, 0xfb, 0xe4, 0x6c, 0xbf, 0xf2,
	0xbf, 0x36, 0x76, 0xef, 0x5a, 0x7f, 0xe7, 0xd1, 0x1d, 0x59, 0x71, 0x58, 0xfe, 0x1a, 0x46, 0xd3,
	0xb4, 0x18, 0x1f, 0x9e, 0xa7, 0xb3, 0x7f, 0xc4, 0xe9, 0x4d, 0x79, 0x7d, 0xf5, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xd6, 0x27, 0xb2, 0x38, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// M5GSubscriberAuthenticationClient is the client API for M5GSubscriberAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type M5GSubscriberAuthenticationClient interface {
	// Authentication-Information (Code 318)
	M5GAuthenticationInformation(ctx context.Context, in *M5GAuthenticationInformationRequest, opts ...grpc.CallOption) (*M5GAuthenticationInformationAnswer, error)
}

type m5GSubscriberAuthenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewM5GSubscriberAuthenticationClient(cc grpc.ClientConnInterface) M5GSubscriberAuthenticationClient {
	return &m5GSubscriberAuthenticationClient{cc}
}

func (c *m5GSubscriberAuthenticationClient) M5GAuthenticationInformation(ctx context.Context, in *M5GAuthenticationInformationRequest, opts ...grpc.CallOption) (*M5GAuthenticationInformationAnswer, error) {
	out := new(M5GAuthenticationInformationAnswer)
	err := c.cc.Invoke(ctx, "/magma.lte.M5GSubscriberAuthentication/M5GAuthenticationInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// M5GSubscriberAuthenticationServer is the server API for M5GSubscriberAuthentication service.
type M5GSubscriberAuthenticationServer interface {
	// Authentication-Information (Code 318)
	M5GAuthenticationInformation(context.Context, *M5GAuthenticationInformationRequest) (*M5GAuthenticationInformationAnswer, error)
}

// UnimplementedM5GSubscriberAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedM5GSubscriberAuthenticationServer struct {
}

func (*UnimplementedM5GSubscriberAuthenticationServer) M5GAuthenticationInformation(ctx context.Context, req *M5GAuthenticationInformationRequest) (*M5GAuthenticationInformationAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method M5GAuthenticationInformation not implemented")
}

func RegisterM5GSubscriberAuthenticationServer(s *grpc.Server, srv M5GSubscriberAuthenticationServer) {
	s.RegisterService(&_M5GSubscriberAuthentication_serviceDesc, srv)
}

func _M5GSubscriberAuthentication_M5GAuthenticationInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(M5GAuthenticationInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(M5GSubscriberAuthenticationServer).M5GAuthenticationInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.M5GSubscriberAuthentication/M5GAuthenticationInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(M5GSubscriberAuthenticationServer).M5GAuthenticationInformation(ctx, req.(*M5GAuthenticationInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _M5GSubscriberAuthentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.M5GSubscriberAuthentication",
	HandlerType: (*M5GSubscriberAuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "M5GAuthenticationInformation",
			Handler:    _M5GSubscriberAuthentication_M5GAuthenticationInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/subscriberauth.proto",
}
