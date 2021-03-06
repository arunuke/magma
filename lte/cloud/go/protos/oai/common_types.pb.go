// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/oai/common_types.proto

package oai

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Ambr_BitrateUnitsAMBR int32

const (
	Ambr_BPS  Ambr_BitrateUnitsAMBR = 0
	Ambr_KBPS Ambr_BitrateUnitsAMBR = 1
)

var Ambr_BitrateUnitsAMBR_name = map[int32]string{
	0: "BPS",
	1: "KBPS",
}

var Ambr_BitrateUnitsAMBR_value = map[string]int32{
	"BPS":  0,
	"KBPS": 1,
}

func (x Ambr_BitrateUnitsAMBR) String() string {
	return proto.EnumName(Ambr_BitrateUnitsAMBR_name, int32(x))
}

func (Ambr_BitrateUnitsAMBR) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{3, 0}
}

// guti_t
type Guti struct {
	Plmn                 []byte   `protobuf:"bytes,1,opt,name=plmn,proto3" json:"plmn,omitempty"`
	MmeGid               uint32   `protobuf:"varint,2,opt,name=mme_gid,json=mmeGid,proto3" json:"mme_gid,omitempty"`
	MmeCode              uint32   `protobuf:"varint,3,opt,name=mme_code,json=mmeCode,proto3" json:"mme_code,omitempty"`
	MTmsi                uint32   `protobuf:"varint,4,opt,name=m_tmsi,json=mTmsi,proto3" json:"m_tmsi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Guti) Reset()         { *m = Guti{} }
func (m *Guti) String() string { return proto.CompactTextString(m) }
func (*Guti) ProtoMessage()    {}
func (*Guti) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{0}
}

func (m *Guti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Guti.Unmarshal(m, b)
}
func (m *Guti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Guti.Marshal(b, m, deterministic)
}
func (m *Guti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Guti.Merge(m, src)
}
func (m *Guti) XXX_Size() int {
	return xxx_messageInfo_Guti.Size(m)
}
func (m *Guti) XXX_DiscardUnknown() {
	xxx_messageInfo_Guti.DiscardUnknown(m)
}

var xxx_messageInfo_Guti proto.InternalMessageInfo

func (m *Guti) GetPlmn() []byte {
	if m != nil {
		return m.Plmn
	}
	return nil
}

func (m *Guti) GetMmeGid() uint32 {
	if m != nil {
		return m.MmeGid
	}
	return 0
}

func (m *Guti) GetMmeCode() uint32 {
	if m != nil {
		return m.MmeCode
	}
	return 0
}

func (m *Guti) GetMTmsi() uint32 {
	if m != nil {
		return m.MTmsi
	}
	return 0
}

// ecgi_t
type Ecgi struct {
	Plmn []byte `protobuf:"bytes,1,opt,name=plmn,proto3" json:"plmn,omitempty"`
	// serializing struct eci_t here without creating a new message
	EnbId                uint32   `protobuf:"varint,2,opt,name=enb_id,json=enbId,proto3" json:"enb_id,omitempty"`
	CellId               uint32   `protobuf:"varint,3,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Empty                uint32   `protobuf:"varint,4,opt,name=empty,proto3" json:"empty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ecgi) Reset()         { *m = Ecgi{} }
func (m *Ecgi) String() string { return proto.CompactTextString(m) }
func (*Ecgi) ProtoMessage()    {}
func (*Ecgi) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{1}
}

func (m *Ecgi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ecgi.Unmarshal(m, b)
}
func (m *Ecgi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ecgi.Marshal(b, m, deterministic)
}
func (m *Ecgi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ecgi.Merge(m, src)
}
func (m *Ecgi) XXX_Size() int {
	return xxx_messageInfo_Ecgi.Size(m)
}
func (m *Ecgi) XXX_DiscardUnknown() {
	xxx_messageInfo_Ecgi.DiscardUnknown(m)
}

var xxx_messageInfo_Ecgi proto.InternalMessageInfo

func (m *Ecgi) GetPlmn() []byte {
	if m != nil {
		return m.Plmn
	}
	return nil
}

func (m *Ecgi) GetEnbId() uint32 {
	if m != nil {
		return m.EnbId
	}
	return 0
}

func (m *Ecgi) GetCellId() uint32 {
	if m != nil {
		return m.CellId
	}
	return 0
}

func (m *Ecgi) GetEmpty() uint32 {
	if m != nil {
		return m.Empty
	}
	return 0
}

// eps_subscribed_qos_profile_t
type EpsSubscribedQosProfile struct {
	Qci uint32 `protobuf:"varint,1,opt,name=qci,proto3" json:"qci,omitempty"`
	// serializing allocation_retention_priority_t here instead of adding a
	// new message type
	PriorityLevel           uint32   `protobuf:"varint,2,opt,name=priority_level,json=priorityLevel,proto3" json:"priority_level,omitempty"`
	PreEmptionVulnerability uint32   `protobuf:"varint,3,opt,name=pre_emption_vulnerability,json=preEmptionVulnerability,proto3" json:"pre_emption_vulnerability,omitempty"`
	PreEmptionCapability    uint32   `protobuf:"varint,4,opt,name=pre_emption_capability,json=preEmptionCapability,proto3" json:"pre_emption_capability,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *EpsSubscribedQosProfile) Reset()         { *m = EpsSubscribedQosProfile{} }
func (m *EpsSubscribedQosProfile) String() string { return proto.CompactTextString(m) }
func (*EpsSubscribedQosProfile) ProtoMessage()    {}
func (*EpsSubscribedQosProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{2}
}

func (m *EpsSubscribedQosProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EpsSubscribedQosProfile.Unmarshal(m, b)
}
func (m *EpsSubscribedQosProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EpsSubscribedQosProfile.Marshal(b, m, deterministic)
}
func (m *EpsSubscribedQosProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpsSubscribedQosProfile.Merge(m, src)
}
func (m *EpsSubscribedQosProfile) XXX_Size() int {
	return xxx_messageInfo_EpsSubscribedQosProfile.Size(m)
}
func (m *EpsSubscribedQosProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_EpsSubscribedQosProfile.DiscardUnknown(m)
}

var xxx_messageInfo_EpsSubscribedQosProfile proto.InternalMessageInfo

func (m *EpsSubscribedQosProfile) GetQci() uint32 {
	if m != nil {
		return m.Qci
	}
	return 0
}

func (m *EpsSubscribedQosProfile) GetPriorityLevel() uint32 {
	if m != nil {
		return m.PriorityLevel
	}
	return 0
}

func (m *EpsSubscribedQosProfile) GetPreEmptionVulnerability() uint32 {
	if m != nil {
		return m.PreEmptionVulnerability
	}
	return 0
}

func (m *EpsSubscribedQosProfile) GetPreEmptionCapability() uint32 {
	if m != nil {
		return m.PreEmptionCapability
	}
	return 0
}

// ambr_t
type Ambr struct {
	BrUl uint64 `protobuf:"varint,1,opt,name=br_ul,json=brUl,proto3" json:"br_ul,omitempty"`
	BrDl uint64 `protobuf:"varint,2,opt,name=br_dl,json=brDl,proto3" json:"br_dl,omitempty"`
	// Unit (either bps or Kbps)
	BrUnit               Ambr_BitrateUnitsAMBR `protobuf:"varint,3,opt,name=br_unit,json=brUnit,proto3,enum=magma.lte.oai.Ambr_BitrateUnitsAMBR" json:"br_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Ambr) Reset()         { *m = Ambr{} }
func (m *Ambr) String() string { return proto.CompactTextString(m) }
func (*Ambr) ProtoMessage()    {}
func (*Ambr) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{3}
}

func (m *Ambr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ambr.Unmarshal(m, b)
}
func (m *Ambr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ambr.Marshal(b, m, deterministic)
}
func (m *Ambr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ambr.Merge(m, src)
}
func (m *Ambr) XXX_Size() int {
	return xxx_messageInfo_Ambr.Size(m)
}
func (m *Ambr) XXX_DiscardUnknown() {
	xxx_messageInfo_Ambr.DiscardUnknown(m)
}

var xxx_messageInfo_Ambr proto.InternalMessageInfo

func (m *Ambr) GetBrUl() uint64 {
	if m != nil {
		return m.BrUl
	}
	return 0
}

func (m *Ambr) GetBrDl() uint64 {
	if m != nil {
		return m.BrDl
	}
	return 0
}

func (m *Ambr) GetBrUnit() Ambr_BitrateUnitsAMBR {
	if m != nil {
		return m.BrUnit
	}
	return Ambr_BPS
}

// apn_configuration_t
type ApnConfig struct {
	ContextIdentifier    uint32                   `protobuf:"varint,1,opt,name=context_identifier,json=contextIdentifier,proto3" json:"context_identifier,omitempty"`
	IpAddress            []string                 `protobuf:"bytes,2,rep,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PdnType              uint32                   `protobuf:"varint,3,opt,name=pdn_type,json=pdnType,proto3" json:"pdn_type,omitempty"`
	ServiceSelection     []byte                   `protobuf:"bytes,4,opt,name=service_selection,json=serviceSelection,proto3" json:"service_selection,omitempty"`
	SubscribedQos        *EpsSubscribedQosProfile `protobuf:"bytes,5,opt,name=subscribed_qos,json=subscribedQos,proto3" json:"subscribed_qos,omitempty"`
	Ambr                 *Ambr                    `protobuf:"bytes,6,opt,name=ambr,proto3" json:"ambr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ApnConfig) Reset()         { *m = ApnConfig{} }
func (m *ApnConfig) String() string { return proto.CompactTextString(m) }
func (*ApnConfig) ProtoMessage()    {}
func (*ApnConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{4}
}

func (m *ApnConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApnConfig.Unmarshal(m, b)
}
func (m *ApnConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApnConfig.Marshal(b, m, deterministic)
}
func (m *ApnConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApnConfig.Merge(m, src)
}
func (m *ApnConfig) XXX_Size() int {
	return xxx_messageInfo_ApnConfig.Size(m)
}
func (m *ApnConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ApnConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ApnConfig proto.InternalMessageInfo

func (m *ApnConfig) GetContextIdentifier() uint32 {
	if m != nil {
		return m.ContextIdentifier
	}
	return 0
}

func (m *ApnConfig) GetIpAddress() []string {
	if m != nil {
		return m.IpAddress
	}
	return nil
}

func (m *ApnConfig) GetPdnType() uint32 {
	if m != nil {
		return m.PdnType
	}
	return 0
}

func (m *ApnConfig) GetServiceSelection() []byte {
	if m != nil {
		return m.ServiceSelection
	}
	return nil
}

func (m *ApnConfig) GetSubscribedQos() *EpsSubscribedQosProfile {
	if m != nil {
		return m.SubscribedQos
	}
	return nil
}

func (m *ApnConfig) GetAmbr() *Ambr {
	if m != nil {
		return m.Ambr
	}
	return nil
}

// apn_config_profile
type ApnConfigProfile struct {
	ContextIdentifier    uint32       `protobuf:"varint,1,opt,name=context_identifier,json=contextIdentifier,proto3" json:"context_identifier,omitempty"`
	AllApnConfInd        uint32       `protobuf:"varint,2,opt,name=all_apn_conf_ind,json=allApnConfInd,proto3" json:"all_apn_conf_ind,omitempty"`
	ApnConfigs           []*ApnConfig `protobuf:"bytes,3,rep,name=apn_configs,json=apnConfigs,proto3" json:"apn_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ApnConfigProfile) Reset()         { *m = ApnConfigProfile{} }
func (m *ApnConfigProfile) String() string { return proto.CompactTextString(m) }
func (*ApnConfigProfile) ProtoMessage()    {}
func (*ApnConfigProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ff32391755a7c2e, []int{5}
}

func (m *ApnConfigProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApnConfigProfile.Unmarshal(m, b)
}
func (m *ApnConfigProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApnConfigProfile.Marshal(b, m, deterministic)
}
func (m *ApnConfigProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApnConfigProfile.Merge(m, src)
}
func (m *ApnConfigProfile) XXX_Size() int {
	return xxx_messageInfo_ApnConfigProfile.Size(m)
}
func (m *ApnConfigProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ApnConfigProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ApnConfigProfile proto.InternalMessageInfo

func (m *ApnConfigProfile) GetContextIdentifier() uint32 {
	if m != nil {
		return m.ContextIdentifier
	}
	return 0
}

func (m *ApnConfigProfile) GetAllApnConfInd() uint32 {
	if m != nil {
		return m.AllApnConfInd
	}
	return 0
}

func (m *ApnConfigProfile) GetApnConfigs() []*ApnConfig {
	if m != nil {
		return m.ApnConfigs
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.lte.oai.Ambr_BitrateUnitsAMBR", Ambr_BitrateUnitsAMBR_name, Ambr_BitrateUnitsAMBR_value)
	proto.RegisterType((*Guti)(nil), "magma.lte.oai.Guti")
	proto.RegisterType((*Ecgi)(nil), "magma.lte.oai.Ecgi")
	proto.RegisterType((*EpsSubscribedQosProfile)(nil), "magma.lte.oai.EpsSubscribedQosProfile")
	proto.RegisterType((*Ambr)(nil), "magma.lte.oai.Ambr")
	proto.RegisterType((*ApnConfig)(nil), "magma.lte.oai.ApnConfig")
	proto.RegisterType((*ApnConfigProfile)(nil), "magma.lte.oai.ApnConfigProfile")
}

func init() { proto.RegisterFile("lte/protos/oai/common_types.proto", fileDescriptor_9ff32391755a7c2e) }

var fileDescriptor_9ff32391755a7c2e = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0x1b, 0x49,
	0x10, 0xdd, 0xc1, 0x63, 0x03, 0x05, 0x46, 0xa6, 0x81, 0x65, 0x38, 0xa0, 0xf5, 0x5a, 0xcb, 0x62,
	0x69, 0xb5, 0xb6, 0x44, 0x72, 0x49, 0xa4, 0x1c, 0x6c, 0x82, 0x90, 0x95, 0x20, 0x91, 0x01, 0x72,
	0xc8, 0xa5, 0x33, 0x33, 0x5d, 0xb6, 0x5a, 0xea, 0x2f, 0xba, 0xdb, 0x28, 0xfe, 0x23, 0x91, 0xf2,
	0x03, 0xf2, 0x4f, 0xf2, 0xc3, 0xa2, 0x1e, 0x7b, 0x1c, 0x07, 0x91, 0x43, 0x6e, 0xd5, 0xef, 0x55,
	0xd7, 0x7b, 0x7a, 0x53, 0xd3, 0xf0, 0xb7, 0xf0, 0xd8, 0x37, 0x56, 0x7b, 0xed, 0xfa, 0x3a, 0xe3,
	0xfd, 0x42, 0x4b, 0xa9, 0x15, 0xf5, 0x33, 0x83, 0xae, 0x57, 0xe2, 0xa4, 0x29, 0xb3, 0x89, 0xcc,
	0x7a, 0xc2, 0x63, 0x4f, 0x67, 0xbc, 0x83, 0x10, 0x5f, 0x4e, 0x3d, 0x27, 0x04, 0x62, 0x23, 0xa4,
	0x4a, 0xa2, 0x76, 0xd4, 0xdd, 0x4e, 0xcb, 0x9a, 0x1c, 0xc2, 0xba, 0x94, 0x48, 0x27, 0x9c, 0x25,
	0x6b, 0xed, 0xa8, 0xdb, 0x4c, 0x1b, 0x52, 0xe2, 0x25, 0x67, 0xe4, 0x08, 0x36, 0x02, 0x51, 0x68,
	0x86, 0x49, 0xad, 0x64, 0x42, 0xe3, 0xb9, 0x66, 0x48, 0x0e, 0xa0, 0x21, 0xa9, 0x97, 0x8e, 0x27,
	0x71, 0x49, 0xd4, 0xe5, 0xad, 0x74, 0xbc, 0xf3, 0x11, 0xe2, 0x8b, 0x62, 0xf2, 0xb4, 0xcc, 0x01,
	0x34, 0x50, 0xe5, 0x74, 0xa9, 0x52, 0x47, 0x95, 0x8f, 0x58, 0x50, 0x2f, 0x50, 0x88, 0x80, 0xcf,
	0x35, 0x1a, 0xe1, 0x38, 0x62, 0x64, 0x1f, 0xea, 0x28, 0x8d, 0x9f, 0x55, 0x0a, 0xe5, 0xa1, 0xf3,
	0x2d, 0x82, 0xc3, 0x0b, 0xe3, 0x6e, 0xa6, 0xb9, 0x2b, 0x2c, 0xcf, 0x91, 0xbd, 0xd3, 0xee, 0xda,
	0xea, 0x31, 0x17, 0x48, 0x5a, 0x50, 0xbb, 0x2f, 0x78, 0x29, 0xda, 0x4c, 0x43, 0x49, 0x4e, 0x60,
	0xc7, 0x58, 0xae, 0x2d, 0xf7, 0x33, 0x2a, 0xf0, 0x01, 0xc5, 0x42, 0xbb, 0x59, 0xa1, 0x6f, 0x03,
	0x48, 0x5e, 0xc2, 0x91, 0xb1, 0x48, 0x83, 0x02, 0xd7, 0x8a, 0x3e, 0x4c, 0x85, 0x42, 0x9b, 0xe5,
	0x5c, 0x70, 0x3f, 0x5b, 0xb8, 0x3a, 0x34, 0x16, 0x2f, 0xe6, 0xfc, 0xfb, 0x55, 0x9a, 0x3c, 0x87,
	0x3f, 0x57, 0xef, 0x16, 0x99, 0xa9, 0x2e, 0xce, 0x7d, 0xef, 0xff, 0xb8, 0x78, 0xbe, 0xe4, 0x3a,
	0x9f, 0x23, 0x88, 0x07, 0x32, 0xb7, 0x64, 0x0f, 0xea, 0xb9, 0xa5, 0x53, 0x51, 0xba, 0x8e, 0xd3,
	0x38, 0xb7, 0x77, 0x62, 0x01, 0xb2, 0xb9, 0xdb, 0x12, 0x7c, 0x2d, 0xc8, 0x2b, 0x58, 0x0f, 0x9d,
	0x8a, 0xfb, 0xd2, 0xd2, 0xce, 0xd9, 0x3f, 0xbd, 0x9f, 0xbe, 0x71, 0x2f, 0xcc, 0xeb, 0x0d, 0xb9,
	0xb7, 0x99, 0xc7, 0x3b, 0xc5, 0xbd, 0x1b, 0x5c, 0x0d, 0xd3, 0xb4, 0x91, 0xdb, 0x70, 0xe8, 0x9c,
	0x40, 0xeb, 0x31, 0x47, 0xd6, 0xa1, 0x36, 0xbc, 0xbe, 0x69, 0xfd, 0x41, 0x36, 0x20, 0x7e, 0x13,
	0xaa, 0xa8, 0xf3, 0x65, 0x0d, 0x36, 0x07, 0x46, 0x9d, 0x6b, 0x35, 0xe6, 0x13, 0xf2, 0x3f, 0x90,
	0x42, 0x2b, 0x8f, 0x9f, 0x3c, 0xe5, 0x0c, 0x95, 0xe7, 0x63, 0x8e, 0x76, 0x11, 0xf0, 0xee, 0x82,
	0x19, 0x2d, 0x09, 0x72, 0x0c, 0xc0, 0x0d, 0xcd, 0x18, 0xb3, 0xe8, 0x5c, 0xb2, 0xd6, 0xae, 0x75,
	0x37, 0xd3, 0x4d, 0x6e, 0x06, 0x73, 0x20, 0xec, 0x93, 0x61, 0xf3, 0x35, 0xad, 0xf6, 0xc9, 0x30,
	0x75, 0x3b, 0x33, 0x48, 0xfe, 0x83, 0x5d, 0x87, 0xf6, 0x81, 0x17, 0x48, 0x1d, 0x0a, 0x2c, 0x42,
	0x5c, 0x65, 0x80, 0xdb, 0x69, 0x6b, 0x41, 0xdc, 0x54, 0x38, 0xb9, 0x82, 0x1d, 0xb7, 0xfc, 0xfe,
	0xf4, 0x5e, 0xbb, 0xa4, 0xde, 0x8e, 0xba, 0x5b, 0x67, 0xff, 0x3e, 0x0a, 0xe4, 0x17, 0x7b, 0x92,
	0x36, 0xdd, 0x2a, 0x4a, 0x4e, 0x21, 0xce, 0x64, 0x6e, 0x93, 0x46, 0x39, 0x64, 0xef, 0x89, 0x54,
	0xd3, 0xb2, 0xa1, 0xf3, 0x35, 0x82, 0xd6, 0x32, 0x9b, 0x6a, 0xe9, 0x7e, 0x33, 0xa2, 0x53, 0x68,
	0x65, 0x42, 0xd0, 0xcc, 0x28, 0x5a, 0x68, 0x35, 0xa6, 0x5c, 0x55, 0xff, 0x43, 0x33, 0x13, 0x62,
	0x31, 0x7d, 0xa4, 0x18, 0x79, 0x01, 0x5b, 0x55, 0x13, 0x9f, 0xb8, 0xa4, 0xd6, 0xae, 0x75, 0xb7,
	0xce, 0x92, 0xc7, 0xe6, 0x2a, 0x37, 0x29, 0x64, 0x55, 0xe9, 0x86, 0x7f, 0x7d, 0x38, 0x2e, 0xdb,
	0xfa, 0xe1, 0x99, 0x28, 0x84, 0x9e, 0xb2, 0xfe, 0x44, 0xaf, 0xbc, 0x17, 0x79, 0xa3, 0xac, 0x9f,
	0x7d, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xce, 0x7f, 0x5b, 0xba, 0x48, 0x04, 0x00, 0x00,
}
