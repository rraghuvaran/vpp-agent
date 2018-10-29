// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l3.proto

package l3

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

type StaticRoute_RouteType int32

const (
	StaticRoute_INTRA_VRF StaticRoute_RouteType = 0
	StaticRoute_INTER_VRF StaticRoute_RouteType = 1
	StaticRoute_DROP      StaticRoute_RouteType = 2
)

var StaticRoute_RouteType_name = map[int32]string{
	0: "INTRA_VRF",
	1: "INTER_VRF",
	2: "DROP",
}
var StaticRoute_RouteType_value = map[string]int32{
	"INTRA_VRF": 0,
	"INTER_VRF": 1,
	"DROP":      2,
}

func (x StaticRoute_RouteType) String() string {
	return proto.EnumName(StaticRoute_RouteType_name, int32(x))
}
func (StaticRoute_RouteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{0, 0}
}

type IPScanNeighbor_Mode int32

const (
	IPScanNeighbor_DISABLED IPScanNeighbor_Mode = 0
	IPScanNeighbor_IPv4     IPScanNeighbor_Mode = 1
	IPScanNeighbor_IPv6     IPScanNeighbor_Mode = 2
	IPScanNeighbor_BOTH     IPScanNeighbor_Mode = 3
)

var IPScanNeighbor_Mode_name = map[int32]string{
	0: "DISABLED",
	1: "IPv4",
	2: "IPv6",
	3: "BOTH",
}
var IPScanNeighbor_Mode_value = map[string]int32{
	"DISABLED": 0,
	"IPv4":     1,
	"IPv6":     2,
	"BOTH":     3,
}

func (x IPScanNeighbor_Mode) String() string {
	return proto.EnumName(IPScanNeighbor_Mode_name, int32(x))
}
func (IPScanNeighbor_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{3, 0}
}

// Static Routes
type StaticRoute struct {
	Type              StaticRoute_RouteType `protobuf:"varint,10,opt,name=type,proto3,enum=l3.StaticRoute_RouteType" json:"type,omitempty"`
	VrfId             uint32                `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	DstNetwork        string                `protobuf:"bytes,3,opt,name=dst_network,json=dstNetwork,proto3" json:"dst_network,omitempty"`
	NextHopAddr       string                `protobuf:"bytes,4,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	OutgoingInterface string                `protobuf:"bytes,5,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	Weight            uint32                `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Preference        uint32                `protobuf:"varint,7,opt,name=preference,proto3" json:"preference,omitempty"`
	// (a poor man's primary and backup)
	ViaVrfId             uint32   `protobuf:"varint,8,opt,name=via_vrf_id,json=viaVrfId,proto3" json:"via_vrf_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticRoute) Reset()         { *m = StaticRoute{} }
func (m *StaticRoute) String() string { return proto.CompactTextString(m) }
func (*StaticRoute) ProtoMessage()    {}
func (*StaticRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{0}
}
func (m *StaticRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticRoute.Unmarshal(m, b)
}
func (m *StaticRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticRoute.Marshal(b, m, deterministic)
}
func (dst *StaticRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticRoute.Merge(dst, src)
}
func (m *StaticRoute) XXX_Size() int {
	return xxx_messageInfo_StaticRoute.Size(m)
}
func (m *StaticRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticRoute.DiscardUnknown(m)
}

var xxx_messageInfo_StaticRoute proto.InternalMessageInfo

func (m *StaticRoute) GetType() StaticRoute_RouteType {
	if m != nil {
		return m.Type
	}
	return StaticRoute_INTRA_VRF
}

func (m *StaticRoute) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *StaticRoute) GetDstNetwork() string {
	if m != nil {
		return m.DstNetwork
	}
	return ""
}

func (m *StaticRoute) GetNextHopAddr() string {
	if m != nil {
		return m.NextHopAddr
	}
	return ""
}

func (m *StaticRoute) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *StaticRoute) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *StaticRoute) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *StaticRoute) GetViaVrfId() uint32 {
	if m != nil {
		return m.ViaVrfId
	}
	return 0
}

// IP ARP Entries
type ARPEntry struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PhysAddress          string   `protobuf:"bytes,3,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Static               bool     `protobuf:"varint,4,opt,name=static,proto3" json:"static,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ARPEntry) Reset()         { *m = ARPEntry{} }
func (m *ARPEntry) String() string { return proto.CompactTextString(m) }
func (*ARPEntry) ProtoMessage()    {}
func (*ARPEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{1}
}
func (m *ARPEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ARPEntry.Unmarshal(m, b)
}
func (m *ARPEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ARPEntry.Marshal(b, m, deterministic)
}
func (dst *ARPEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ARPEntry.Merge(dst, src)
}
func (m *ARPEntry) XXX_Size() int {
	return xxx_messageInfo_ARPEntry.Size(m)
}
func (m *ARPEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ARPEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ARPEntry proto.InternalMessageInfo

func (m *ARPEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ARPEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ARPEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *ARPEntry) GetStatic() bool {
	if m != nil {
		return m.Static
	}
	return false
}

type ProxyARP struct {
	Interfaces           []*ProxyARP_Interface `protobuf:"bytes,1,rep,name=interfaces" json:"interfaces,omitempty"`
	Ranges               []*ProxyARP_Range     `protobuf:"bytes,2,rep,name=ranges" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProxyARP) Reset()         { *m = ProxyARP{} }
func (m *ProxyARP) String() string { return proto.CompactTextString(m) }
func (*ProxyARP) ProtoMessage()    {}
func (*ProxyARP) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{2}
}
func (m *ProxyARP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP.Unmarshal(m, b)
}
func (m *ProxyARP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP.Marshal(b, m, deterministic)
}
func (dst *ProxyARP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP.Merge(dst, src)
}
func (m *ProxyARP) XXX_Size() int {
	return xxx_messageInfo_ProxyARP.Size(m)
}
func (m *ProxyARP) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP proto.InternalMessageInfo

func (m *ProxyARP) GetInterfaces() []*ProxyARP_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ProxyARP) GetRanges() []*ProxyARP_Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type ProxyARP_Interface struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyARP_Interface) Reset()         { *m = ProxyARP_Interface{} }
func (m *ProxyARP_Interface) String() string { return proto.CompactTextString(m) }
func (*ProxyARP_Interface) ProtoMessage()    {}
func (*ProxyARP_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{2, 0}
}
func (m *ProxyARP_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP_Interface.Unmarshal(m, b)
}
func (m *ProxyARP_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP_Interface.Marshal(b, m, deterministic)
}
func (dst *ProxyARP_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP_Interface.Merge(dst, src)
}
func (m *ProxyARP_Interface) XXX_Size() int {
	return xxx_messageInfo_ProxyARP_Interface.Size(m)
}
func (m *ProxyARP_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP_Interface proto.InternalMessageInfo

func (m *ProxyARP_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProxyARP_Range struct {
	FirstIpAddr          string   `protobuf:"bytes,1,opt,name=first_ip_addr,json=firstIpAddr,proto3" json:"first_ip_addr,omitempty"`
	LastIpAddr           string   `protobuf:"bytes,2,opt,name=last_ip_addr,json=lastIpAddr,proto3" json:"last_ip_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyARP_Range) Reset()         { *m = ProxyARP_Range{} }
func (m *ProxyARP_Range) String() string { return proto.CompactTextString(m) }
func (*ProxyARP_Range) ProtoMessage()    {}
func (*ProxyARP_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{2, 1}
}
func (m *ProxyARP_Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP_Range.Unmarshal(m, b)
}
func (m *ProxyARP_Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP_Range.Marshal(b, m, deterministic)
}
func (dst *ProxyARP_Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP_Range.Merge(dst, src)
}
func (m *ProxyARP_Range) XXX_Size() int {
	return xxx_messageInfo_ProxyARP_Range.Size(m)
}
func (m *ProxyARP_Range) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP_Range.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP_Range proto.InternalMessageInfo

func (m *ProxyARP_Range) GetFirstIpAddr() string {
	if m != nil {
		return m.FirstIpAddr
	}
	return ""
}

func (m *ProxyARP_Range) GetLastIpAddr() string {
	if m != nil {
		return m.LastIpAddr
	}
	return ""
}

// Enables/disables IP neighbor scanning
type IPScanNeighbor struct {
	Mode                 IPScanNeighbor_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=l3.IPScanNeighbor_Mode" json:"mode,omitempty"`
	ScanInterval         uint32              `protobuf:"varint,2,opt,name=scan_interval,json=scanInterval,proto3" json:"scan_interval,omitempty"`
	MaxProcTime          uint32              `protobuf:"varint,3,opt,name=max_proc_time,json=maxProcTime,proto3" json:"max_proc_time,omitempty"`
	MaxUpdate            uint32              `protobuf:"varint,4,opt,name=max_update,json=maxUpdate,proto3" json:"max_update,omitempty"`
	ScanIntDelay         uint32              `protobuf:"varint,5,opt,name=scan_int_delay,json=scanIntDelay,proto3" json:"scan_int_delay,omitempty"`
	StaleThreshold       uint32              `protobuf:"varint,6,opt,name=stale_threshold,json=staleThreshold,proto3" json:"stale_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IPScanNeighbor) Reset()         { *m = IPScanNeighbor{} }
func (m *IPScanNeighbor) String() string { return proto.CompactTextString(m) }
func (*IPScanNeighbor) ProtoMessage()    {}
func (*IPScanNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_l3_0527e5292e3a3685, []int{3}
}
func (m *IPScanNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPScanNeighbor.Unmarshal(m, b)
}
func (m *IPScanNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPScanNeighbor.Marshal(b, m, deterministic)
}
func (dst *IPScanNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPScanNeighbor.Merge(dst, src)
}
func (m *IPScanNeighbor) XXX_Size() int {
	return xxx_messageInfo_IPScanNeighbor.Size(m)
}
func (m *IPScanNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IPScanNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IPScanNeighbor proto.InternalMessageInfo

func (m *IPScanNeighbor) GetMode() IPScanNeighbor_Mode {
	if m != nil {
		return m.Mode
	}
	return IPScanNeighbor_DISABLED
}

func (m *IPScanNeighbor) GetScanInterval() uint32 {
	if m != nil {
		return m.ScanInterval
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxProcTime() uint32 {
	if m != nil {
		return m.MaxProcTime
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxUpdate() uint32 {
	if m != nil {
		return m.MaxUpdate
	}
	return 0
}

func (m *IPScanNeighbor) GetScanIntDelay() uint32 {
	if m != nil {
		return m.ScanIntDelay
	}
	return 0
}

func (m *IPScanNeighbor) GetStaleThreshold() uint32 {
	if m != nil {
		return m.StaleThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*StaticRoute)(nil), "l3.StaticRoute")
	proto.RegisterType((*ARPEntry)(nil), "l3.ARPEntry")
	proto.RegisterType((*ProxyARP)(nil), "l3.ProxyARP")
	proto.RegisterType((*ProxyARP_Interface)(nil), "l3.ProxyARP.Interface")
	proto.RegisterType((*ProxyARP_Range)(nil), "l3.ProxyARP.Range")
	proto.RegisterType((*IPScanNeighbor)(nil), "l3.IPScanNeighbor")
	proto.RegisterEnum("l3.StaticRoute_RouteType", StaticRoute_RouteType_name, StaticRoute_RouteType_value)
	proto.RegisterEnum("l3.IPScanNeighbor_Mode", IPScanNeighbor_Mode_name, IPScanNeighbor_Mode_value)
}

func init() { proto.RegisterFile("l3.proto", fileDescriptor_l3_0527e5292e3a3685) }

var fileDescriptor_l3_0527e5292e3a3685 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdf, 0x6e, 0xda, 0x30,
	0x14, 0xc6, 0x1b, 0x4a, 0x59, 0x38, 0x21, 0x8c, 0x59, 0x5a, 0x97, 0x55, 0xdd, 0xca, 0xb2, 0x49,
	0x43, 0x9b, 0xca, 0x05, 0x4c, 0xbd, 0xa7, 0xa2, 0x53, 0x23, 0xad, 0x6d, 0xe4, 0xb2, 0xde, 0x46,
	0x6e, 0x62, 0x20, 0x5a, 0x12, 0x47, 0xb6, 0x49, 0xe1, 0x76, 0x4f, 0xb3, 0xe7, 0xd9, 0x03, 0xec,
	0x59, 0x26, 0x3b, 0x09, 0xd0, 0x1b, 0x94, 0xf3, 0x9d, 0x9f, 0x75, 0xfe, 0x7c, 0x07, 0x30, 0x93,
	0xf1, 0x30, 0xe7, 0x4c, 0x32, 0xd4, 0x48, 0xc6, 0xee, 0xdf, 0x06, 0x58, 0xf7, 0x92, 0xc8, 0x38,
	0xc4, 0x6c, 0x25, 0x29, 0x3a, 0x87, 0xa6, 0xdc, 0xe4, 0xd4, 0x81, 0xbe, 0x31, 0xe8, 0x8e, 0xde,
	0x0e, 0x93, 0xf1, 0x70, 0x2f, 0x3d, 0xd4, 0xbf, 0xb3, 0x4d, 0x4e, 0xb1, 0xc6, 0xd0, 0x6b, 0x68,
	0x15, 0x7c, 0x1e, 0xc4, 0x91, 0x63, 0xf4, 0x8d, 0x81, 0x8d, 0x8f, 0x0a, 0x3e, 0xf7, 0x22, 0x74,
	0x06, 0x56, 0x24, 0x64, 0x90, 0x51, 0xf9, 0xc4, 0xf8, 0x2f, 0xe7, 0xb0, 0x6f, 0x0c, 0xda, 0x18,
	0x22, 0x21, 0x6f, 0x4b, 0x05, 0xb9, 0x60, 0x67, 0x74, 0x2d, 0x83, 0x25, 0xcb, 0x03, 0x12, 0x45,
	0xdc, 0x69, 0x6a, 0xc4, 0x52, 0xe2, 0x35, 0xcb, 0x27, 0x51, 0xc4, 0xd1, 0x39, 0x20, 0xb6, 0x92,
	0x0b, 0x16, 0x67, 0x8b, 0x20, 0xce, 0x24, 0xe5, 0x73, 0x12, 0x52, 0xe7, 0x48, 0x83, 0xaf, 0xea,
	0x8c, 0x57, 0x27, 0xd0, 0x31, 0xb4, 0x9e, 0x68, 0xbc, 0x58, 0x4a, 0xa7, 0xa5, 0x5b, 0xa9, 0x22,
	0xf4, 0x1e, 0x20, 0xe7, 0x74, 0x4e, 0x39, 0xcd, 0x42, 0xea, 0xbc, 0xd0, 0xb9, 0x3d, 0x05, 0x9d,
	0x02, 0x14, 0x31, 0x09, 0xaa, 0x31, 0x4c, 0x9d, 0x37, 0x8b, 0x98, 0x3c, 0xa8, 0x49, 0xdc, 0x31,
	0xb4, 0xb7, 0x33, 0x23, 0x1b, 0xda, 0xde, 0xed, 0x0c, 0x4f, 0x82, 0x07, 0xfc, 0xbd, 0x77, 0x50,
	0x85, 0x57, 0x58, 0x87, 0x06, 0x32, 0xa1, 0x39, 0xc5, 0x77, 0x7e, 0xaf, 0xe1, 0xfe, 0x36, 0xc0,
	0x9c, 0x60, 0xff, 0x2a, 0x93, 0x7c, 0x83, 0x4e, 0xa1, 0xbd, 0xeb, 0xde, 0xd0, 0xdd, 0xef, 0x04,
	0xf4, 0x0e, 0x20, 0x2e, 0x57, 0x40, 0x85, 0x70, 0x1a, 0x55, 0x5a, 0x2f, 0x80, 0x0a, 0x81, 0x3e,
	0x40, 0x27, 0x5f, 0x6e, 0xc4, 0x16, 0x28, 0x37, 0x69, 0x29, 0xad, 0x46, 0x8e, 0xa1, 0x25, 0xb4,
	0x43, 0x7a, 0x87, 0x26, 0xae, 0x22, 0xf7, 0x9f, 0x01, 0xa6, 0xcf, 0xd9, 0x7a, 0x33, 0xc1, 0x3e,
	0xba, 0x00, 0xd8, 0xd6, 0x14, 0x8e, 0xd1, 0x3f, 0x1c, 0x58, 0xa3, 0x63, 0x65, 0x6e, 0x4d, 0x0c,
	0xb7, 0x8b, 0xc4, 0x7b, 0x24, 0xfa, 0x02, 0x2d, 0x4e, 0xb2, 0x05, 0x55, 0xad, 0xa9, 0x37, 0xe8,
	0xd9, 0x1b, 0xac, 0x52, 0xb8, 0x22, 0x4e, 0xce, 0xa0, 0xbd, 0x73, 0x03, 0x41, 0x33, 0x23, 0x69,
	0x3d, 0xb0, 0xfe, 0x3e, 0xb9, 0x81, 0x23, 0xfd, 0x42, 0xb9, 0x3f, 0x8f, 0xb9, 0x90, 0x41, 0x35,
	0x7a, 0x45, 0x59, 0x5a, 0xf4, 0x4a, 0xf7, 0xfb, 0xd0, 0x49, 0xc8, 0x1e, 0x52, 0xae, 0x06, 0x94,
	0x56, 0x12, 0xee, 0x9f, 0x06, 0x74, 0x3d, 0xff, 0x3e, 0x24, 0xd9, 0xad, 0x72, 0xfa, 0x91, 0x71,
	0xf4, 0x15, 0x9a, 0x29, 0x8b, 0xca, 0xaa, 0xdd, 0xd1, 0x1b, 0xd5, 0xec, 0x73, 0x62, 0x78, 0xc3,
	0x22, 0x8a, 0x35, 0x84, 0x3e, 0x82, 0x2d, 0x42, 0x92, 0x95, 0xb7, 0x55, 0x90, 0x44, 0x97, 0xb0,
	0x71, 0x47, 0x89, 0x5e, 0xa5, 0xa9, 0x56, 0x53, 0xb2, 0x0e, 0x72, 0xce, 0xc2, 0x40, 0xc6, 0x29,
	0xd5, 0x0e, 0xd8, 0xd8, 0x4a, 0xc9, 0xda, 0xe7, 0x2c, 0x9c, 0xc5, 0xa9, 0xf6, 0x50, 0x31, 0xab,
	0x3c, 0x22, 0x92, 0x6a, 0x17, 0x6c, 0xdc, 0x4e, 0xc9, 0xfa, 0xa7, 0x16, 0xd0, 0x27, 0xe8, 0xd6,
	0x75, 0x82, 0x88, 0x26, 0x64, 0xa3, 0x6f, 0x78, 0x57, 0x68, 0xaa, 0x34, 0xf4, 0x19, 0x5e, 0x0a,
	0x49, 0x12, 0x1a, 0xc8, 0x25, 0xa7, 0x62, 0xc9, 0x92, 0xa8, 0xba, 0xe3, 0xae, 0x96, 0x67, 0xb5,
	0xea, 0x8e, 0xa0, 0xa9, 0x86, 0x40, 0x1d, 0x30, 0xa7, 0xde, 0xfd, 0xe4, 0xf2, 0xc7, 0xd5, 0xb4,
	0x77, 0xa0, 0x8e, 0xcf, 0xf3, 0x8b, 0x6f, 0xe5, 0x19, 0x7a, 0x7e, 0x71, 0xd1, 0x6b, 0xa8, 0xaf,
	0xcb, 0xbb, 0xd9, 0x75, 0xef, 0xf0, 0xb1, 0xa5, 0xff, 0xf0, 0xe3, 0xff, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xdc, 0x5b, 0x7e, 0x43, 0xfc, 0x03, 0x00, 0x00,
}
