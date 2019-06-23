// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

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

type ServiceType_Enum int32

const (
	ServiceType_NOT_SET  ServiceType_Enum = 0
	ServiceType_SS_LIBEV ServiceType_Enum = 1
	ServiceType_SSR      ServiceType_Enum = 2
)

var ServiceType_Enum_name = map[int32]string{
	0: "NOT_SET",
	1: "SS_LIBEV",
	2: "SSR",
}

var ServiceType_Enum_value = map[string]int32{
	"NOT_SET":  0,
	"SS_LIBEV": 1,
	"SSR":      2,
}

func (x ServiceType_Enum) String() string {
	return proto.EnumName(ServiceType_Enum_name, int32(x))
}

func (ServiceType_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0, 0}
}

type ServiceType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceType) Reset()         { *m = ServiceType{} }
func (m *ServiceType) String() string { return proto.CompactTextString(m) }
func (*ServiceType) ProtoMessage()    {}
func (*ServiceType) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *ServiceType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceType.Unmarshal(m, b)
}
func (m *ServiceType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceType.Marshal(b, m, deterministic)
}
func (m *ServiceType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceType.Merge(m, src)
}
func (m *ServiceType) XXX_Size() int {
	return xxx_messageInfo_ServiceType.Size(m)
}
func (m *ServiceType) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceType.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceType proto.InternalMessageInfo

type GeneralRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralRequest) Reset()         { *m = GeneralRequest{} }
func (m *GeneralRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralRequest) ProtoMessage()    {}
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *GeneralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralRequest.Unmarshal(m, b)
}
func (m *GeneralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralRequest.Marshal(b, m, deterministic)
}
func (m *GeneralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralRequest.Merge(m, src)
}
func (m *GeneralRequest) XXX_Size() int {
	return xxx_messageInfo_GeneralRequest.Size(m)
}
func (m *GeneralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralRequest proto.InternalMessageInfo

func (m *GeneralRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GeneralResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (m *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(m, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

type HeartbeatStreamServer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatStreamServer) Reset()         { *m = HeartbeatStreamServer{} }
func (m *HeartbeatStreamServer) String() string { return proto.CompactTextString(m) }
func (*HeartbeatStreamServer) ProtoMessage()    {}
func (*HeartbeatStreamServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{3}
}

func (m *HeartbeatStreamServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatStreamServer.Unmarshal(m, b)
}
func (m *HeartbeatStreamServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatStreamServer.Marshal(b, m, deterministic)
}
func (m *HeartbeatStreamServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatStreamServer.Merge(m, src)
}
func (m *HeartbeatStreamServer) XXX_Size() int {
	return xxx_messageInfo_HeartbeatStreamServer.Size(m)
}
func (m *HeartbeatStreamServer) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatStreamServer.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatStreamServer proto.InternalMessageInfo

type SyncStreamServer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncStreamServer) Reset()         { *m = SyncStreamServer{} }
func (m *SyncStreamServer) String() string { return proto.CompactTextString(m) }
func (*SyncStreamServer) ProtoMessage()    {}
func (*SyncStreamServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{4}
}

func (m *SyncStreamServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncStreamServer.Unmarshal(m, b)
}
func (m *SyncStreamServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncStreamServer.Marshal(b, m, deterministic)
}
func (m *SyncStreamServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncStreamServer.Merge(m, src)
}
func (m *SyncStreamServer) XXX_Size() int {
	return xxx_messageInfo_SyncStreamServer.Size(m)
}
func (m *SyncStreamServer) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncStreamServer.DiscardUnknown(m)
}

var xxx_messageInfo_SyncStreamServer proto.InternalMessageInfo

type GetAvailablePortRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UsedPorts            []int32  `protobuf:"varint,2,rep,packed,name=used_ports,json=usedPorts,proto3" json:"used_ports,omitempty"`
	PortFrom             int32    `protobuf:"varint,3,opt,name=port_from,json=portFrom,proto3" json:"port_from,omitempty"`
	PortTo               int32    `protobuf:"varint,4,opt,name=port_to,json=portTo,proto3" json:"port_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAvailablePortRequest) Reset()         { *m = GetAvailablePortRequest{} }
func (m *GetAvailablePortRequest) String() string { return proto.CompactTextString(m) }
func (*GetAvailablePortRequest) ProtoMessage()    {}
func (*GetAvailablePortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{5}
}

func (m *GetAvailablePortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAvailablePortRequest.Unmarshal(m, b)
}
func (m *GetAvailablePortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAvailablePortRequest.Marshal(b, m, deterministic)
}
func (m *GetAvailablePortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAvailablePortRequest.Merge(m, src)
}
func (m *GetAvailablePortRequest) XXX_Size() int {
	return xxx_messageInfo_GetAvailablePortRequest.Size(m)
}
func (m *GetAvailablePortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAvailablePortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAvailablePortRequest proto.InternalMessageInfo

func (m *GetAvailablePortRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetAvailablePortRequest) GetUsedPorts() []int32 {
	if m != nil {
		return m.UsedPorts
	}
	return nil
}

func (m *GetAvailablePortRequest) GetPortFrom() int32 {
	if m != nil {
		return m.PortFrom
	}
	return 0
}

func (m *GetAvailablePortRequest) GetPortTo() int32 {
	if m != nil {
		return m.PortTo
	}
	return 0
}

type GetAvailablePortResponse struct {
	Port                 int32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAvailablePortResponse) Reset()         { *m = GetAvailablePortResponse{} }
func (m *GetAvailablePortResponse) String() string { return proto.CompactTextString(m) }
func (*GetAvailablePortResponse) ProtoMessage()    {}
func (*GetAvailablePortResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{6}
}

func (m *GetAvailablePortResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAvailablePortResponse.Unmarshal(m, b)
}
func (m *GetAvailablePortResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAvailablePortResponse.Marshal(b, m, deterministic)
}
func (m *GetAvailablePortResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAvailablePortResponse.Merge(m, src)
}
func (m *GetAvailablePortResponse) XXX_Size() int {
	return xxx_messageInfo_GetAvailablePortResponse.Size(m)
}
func (m *GetAvailablePortResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAvailablePortResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAvailablePortResponse proto.InternalMessageInfo

func (m *GetAvailablePortResponse) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type CreateServiceRequest struct {
	Token                string           `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PortFrom             int32            `protobuf:"varint,2,opt,name=port_from,json=portFrom,proto3" json:"port_from,omitempty"`
	PortTo               int32            `protobuf:"varint,3,opt,name=port_to,json=portTo,proto3" json:"port_to,omitempty"`
	Type                 ServiceType_Enum `protobuf:"varint,4,opt,name=type,proto3,enum=ServiceType_Enum" json:"type,omitempty"`
	Method               string           `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Password             string           `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string           `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{7}
}

func (m *CreateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceRequest.Unmarshal(m, b)
}
func (m *CreateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceRequest.Merge(m, src)
}
func (m *CreateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceRequest.Size(m)
}
func (m *CreateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceRequest proto.InternalMessageInfo

func (m *CreateServiceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateServiceRequest) GetPortFrom() int32 {
	if m != nil {
		return m.PortFrom
	}
	return 0
}

func (m *CreateServiceRequest) GetPortTo() int32 {
	if m != nil {
		return m.PortTo
	}
	return 0
}

func (m *CreateServiceRequest) GetType() ServiceType_Enum {
	if m != nil {
		return m.Type
	}
	return ServiceType_NOT_SET
}

func (m *CreateServiceRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CreateServiceRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateServiceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateServiceResponse struct {
	ServiceId            string   `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceResponse) Reset()         { *m = CreateServiceResponse{} }
func (m *CreateServiceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceResponse) ProtoMessage()    {}
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{8}
}

func (m *CreateServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceResponse.Unmarshal(m, b)
}
func (m *CreateServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceResponse.Merge(m, src)
}
func (m *CreateServiceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceResponse.Size(m)
}
func (m *CreateServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceResponse proto.InternalMessageInfo

func (m *CreateServiceResponse) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *CreateServiceResponse) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *CreateServiceResponse) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type StopServiceRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ServiceId            string   `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopServiceRequest) Reset()         { *m = StopServiceRequest{} }
func (m *StopServiceRequest) String() string { return proto.CompactTextString(m) }
func (*StopServiceRequest) ProtoMessage()    {}
func (*StopServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{9}
}

func (m *StopServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopServiceRequest.Unmarshal(m, b)
}
func (m *StopServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopServiceRequest.Marshal(b, m, deterministic)
}
func (m *StopServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopServiceRequest.Merge(m, src)
}
func (m *StopServiceRequest) XXX_Size() int {
	return xxx_messageInfo_StopServiceRequest.Size(m)
}
func (m *StopServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopServiceRequest proto.InternalMessageInfo

func (m *StopServiceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *StopServiceRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type RemoveServiceRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ServiceId            string   `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServiceRequest) Reset()         { *m = RemoveServiceRequest{} }
func (m *RemoveServiceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveServiceRequest) ProtoMessage()    {}
func (*RemoveServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{10}
}

func (m *RemoveServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveServiceRequest.Unmarshal(m, b)
}
func (m *RemoveServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveServiceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveServiceRequest.Merge(m, src)
}
func (m *RemoveServiceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveServiceRequest.Size(m)
}
func (m *RemoveServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveServiceRequest proto.InternalMessageInfo

func (m *RemoveServiceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RemoveServiceRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func init() {
	proto.RegisterEnum("ServiceType_Enum", ServiceType_Enum_name, ServiceType_Enum_value)
	proto.RegisterType((*ServiceType)(nil), "ServiceType")
	proto.RegisterType((*GeneralRequest)(nil), "GeneralRequest")
	proto.RegisterType((*GeneralResponse)(nil), "GeneralResponse")
	proto.RegisterType((*HeartbeatStreamServer)(nil), "HeartbeatStreamServer")
	proto.RegisterType((*SyncStreamServer)(nil), "SyncStreamServer")
	proto.RegisterType((*GetAvailablePortRequest)(nil), "GetAvailablePortRequest")
	proto.RegisterType((*GetAvailablePortResponse)(nil), "GetAvailablePortResponse")
	proto.RegisterType((*CreateServiceRequest)(nil), "CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "CreateServiceResponse")
	proto.RegisterType((*StopServiceRequest)(nil), "StopServiceRequest")
	proto.RegisterType((*RemoveServiceRequest)(nil), "RemoveServiceRequest")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0x93, 0x34, 0x93, 0xfe, 0xb8, 0x43, 0xd3, 0x9a, 0xa0, 0x4a, 0xc5, 0x12, 0x28,
	0x54, 0x74, 0x5b, 0x15, 0x0e, 0xc0, 0x01, 0xa9, 0x45, 0xa5, 0x44, 0xa0, 0x82, 0xec, 0x88, 0x03,
	0x97, 0xc8, 0x89, 0x87, 0xd4, 0xa2, 0xf6, 0x9a, 0xf5, 0xa6, 0x28, 0x77, 0x1e, 0x83, 0x97, 0xe2,
	0x8d, 0xd0, 0x6e, 0xac, 0x28, 0x4e, 0x9c, 0x8a, 0x03, 0xa7, 0xdd, 0xf9, 0xd9, 0xdd, 0x6f, 0xe6,
	0xfb, 0x66, 0xa1, 0x19, 0x8c, 0x28, 0x91, 0x2c, 0x15, 0x5c, 0x72, 0xf7, 0x15, 0x34, 0x7d, 0x12,
	0xb7, 0xd1, 0x90, 0x7a, 0x93, 0x94, 0xdc, 0x43, 0xb0, 0x2e, 0x92, 0x71, 0x8c, 0x4d, 0xa8, 0x5f,
	0x7d, 0xea, 0xf5, 0xfd, 0x8b, 0x9e, 0x7d, 0x0f, 0xd7, 0x61, 0xcd, 0xf7, 0xfb, 0x1f, 0xbb, 0xe7,
	0x17, 0x5f, 0x6c, 0x03, 0xeb, 0x60, 0xfa, 0xbe, 0x67, 0x57, 0xdc, 0x27, 0xb0, 0x79, 0x49, 0x09,
	0x89, 0xe0, 0xc6, 0xa3, 0x1f, 0x63, 0xca, 0x24, 0xee, 0x40, 0x55, 0xf2, 0xef, 0x94, 0x38, 0xc6,
	0x81, 0xd1, 0x69, 0x78, 0x53, 0xc3, 0xdd, 0x86, 0xad, 0x59, 0x5e, 0x96, 0xf2, 0x24, 0x23, 0x77,
	0x0f, 0x5a, 0xef, 0x29, 0x10, 0x72, 0x40, 0x81, 0xf4, 0xa5, 0xa0, 0x20, 0x56, 0x20, 0x48, 0xb8,
	0x08, 0xb6, 0x3f, 0x49, 0x86, 0x05, 0xdf, 0x2f, 0x03, 0xf6, 0x2e, 0x49, 0x9e, 0xdd, 0x06, 0xd1,
	0x4d, 0x30, 0xb8, 0xa1, 0xcf, 0x5c, 0xc8, 0x3b, 0x5f, 0xc4, 0x7d, 0x80, 0x71, 0x46, 0x61, 0x3f,
	0xe5, 0x42, 0x66, 0x4e, 0xe5, 0xc0, 0xec, 0x54, 0xbd, 0x86, 0xf2, 0xa8, 0xa3, 0x19, 0x3e, 0x84,
	0x86, 0x8a, 0xf4, 0xbf, 0x09, 0x1e, 0x3b, 0xe6, 0x81, 0xd1, 0xa9, 0x7a, 0x6b, 0xca, 0xf1, 0x4e,
	0xf0, 0x18, 0xf7, 0xa0, 0xae, 0x83, 0x92, 0x3b, 0x96, 0x0e, 0xd5, 0x94, 0xd9, 0xe3, 0x2e, 0x03,
	0x67, 0x19, 0xc5, 0xb4, 0x1e, 0x44, 0xb0, 0x54, 0x96, 0x46, 0x51, 0xf5, 0xf4, 0xde, 0xfd, 0x63,
	0xc0, 0xce, 0x5b, 0x41, 0x81, 0xa4, 0xbc, 0xc1, 0x77, 0x63, 0x2e, 0x80, 0xaa, 0xac, 0x06, 0x65,
	0xce, 0x83, 0xc2, 0xc7, 0x60, 0xc9, 0x49, 0x4a, 0x1a, 0xea, 0xe6, 0xe9, 0x36, 0x9b, 0xe3, 0x92,
	0x29, 0x22, 0x3d, 0x1d, 0xc6, 0x5d, 0xa8, 0xc5, 0x24, 0xaf, 0x79, 0xe8, 0x54, 0xf5, 0x9b, 0xb9,
	0x85, 0x6d, 0x58, 0x4b, 0x83, 0x2c, 0xfb, 0xc9, 0x45, 0xe8, 0xd4, 0x74, 0x64, 0x66, 0xab, 0x9a,
	0x92, 0x20, 0x26, 0xa7, 0xae, 0xfd, 0x7a, 0xef, 0xc6, 0xd0, 0x5a, 0x28, 0x29, 0x6f, 0xc0, 0x3e,
	0x40, 0x36, 0x75, 0xf5, 0xa3, 0x30, 0x2f, 0xac, 0x91, 0x7b, 0xba, 0x21, 0x3e, 0x82, 0xf5, 0x21,
	0x4f, 0x64, 0x10, 0x25, 0x24, 0x54, 0x42, 0x45, 0x27, 0x34, 0x67, 0xbe, 0x6e, 0x38, 0x6b, 0xa1,
	0x39, 0xd7, 0xc2, 0x2e, 0xa0, 0x2f, 0x79, 0xfa, 0x4f, 0xfd, 0x2b, 0x22, 0xa8, 0x2c, 0x20, 0x70,
	0x3f, 0xc0, 0x8e, 0x47, 0x31, 0xbf, 0xa5, 0xff, 0x70, 0xd9, 0xe9, 0x6f, 0x13, 0xd6, 0xcf, 0xd4,
	0x10, 0xe5, 0x97, 0xe1, 0x6b, 0xd8, 0xb8, 0xe2, 0x21, 0xcd, 0x34, 0x8d, 0x5b, 0xac, 0x38, 0x1a,
	0xed, 0x5d, 0x56, 0x2a, 0xf8, 0x13, 0x03, 0x9f, 0x81, 0xa5, 0x24, 0xbf, 0x7c, 0x64, 0x9b, 0x2d,
	0x8e, 0xc2, 0x89, 0x81, 0x4f, 0xc1, 0xea, 0x26, 0x51, 0xc9, 0x03, 0x36, 0x5b, 0x18, 0x32, 0xbc,
	0x04, 0x7b, 0x51, 0xb0, 0xe8, 0xb0, 0x15, 0x93, 0xd4, 0x7e, 0xc0, 0x56, 0xaa, 0xfb, 0x0d, 0x6c,
	0x14, 0x58, 0xc7, 0x16, 0x2b, 0x13, 0x76, 0x7b, 0x97, 0x95, 0x8b, 0xe3, 0x05, 0x34, 0xe7, 0x68,
	0xc4, 0xfb, 0x6c, 0x99, 0xd4, 0x12, 0xf8, 0x2f, 0x61, 0xa3, 0xc0, 0x18, 0xb6, 0x58, 0x19, 0x83,
	0xcb, 0x27, 0xcf, 0x0f, 0xbf, 0x76, 0x46, 0x91, 0xbc, 0x1e, 0x0f, 0xd8, 0x90, 0xc7, 0xc7, 0x23,
	0x7e, 0x14, 0x8d, 0x92, 0x48, 0xd2, 0xf1, 0x74, 0x39, 0xd2, 0xdf, 0xdf, 0xb1, 0xfe, 0xfe, 0xb2,
	0x41, 0x4d, 0xaf, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x67, 0x41, 0x26, 0x15, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	NodeHeartbeat(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (AgentService_NodeHeartbeatClient, error)
	Sync(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (AgentService_SyncClient, error)
	Init(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	GetAvailablePort(ctx context.Context, in *GetAvailablePortRequest, opts ...grpc.CallOption) (*GetAvailablePortResponse, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) NodeHeartbeat(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (AgentService_NodeHeartbeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentService_serviceDesc.Streams[0], "/AgentService/NodeHeartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceNodeHeartbeatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_NodeHeartbeatClient interface {
	Recv() (*HeartbeatStreamServer, error)
	grpc.ClientStream
}

type agentServiceNodeHeartbeatClient struct {
	grpc.ClientStream
}

func (x *agentServiceNodeHeartbeatClient) Recv() (*HeartbeatStreamServer, error) {
	m := new(HeartbeatStreamServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) Sync(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (AgentService_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentService_serviceDesc.Streams[1], "/AgentService/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_SyncClient interface {
	Recv() (*SyncStreamServer, error)
	grpc.ClientStream
}

type agentServiceSyncClient struct {
	grpc.ClientStream
}

func (x *agentServiceSyncClient) Recv() (*SyncStreamServer, error) {
	m := new(SyncStreamServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) Init(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/AgentService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) GetAvailablePort(ctx context.Context, in *GetAvailablePortRequest, opts ...grpc.CallOption) (*GetAvailablePortResponse, error) {
	out := new(GetAvailablePortResponse)
	err := c.cc.Invoke(ctx, "/AgentService/GetAvailablePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/AgentService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/AgentService/StopService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/AgentService/RemoveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	NodeHeartbeat(*GeneralRequest, AgentService_NodeHeartbeatServer) error
	Sync(*GeneralRequest, AgentService_SyncServer) error
	Init(context.Context, *GeneralRequest) (*GeneralResponse, error)
	GetAvailablePort(context.Context, *GetAvailablePortRequest) (*GetAvailablePortResponse, error)
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	StopService(context.Context, *StopServiceRequest) (*GeneralResponse, error)
	RemoveService(context.Context, *RemoveServiceRequest) (*GeneralResponse, error)
}

// UnimplementedAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) NodeHeartbeat(req *GeneralRequest, srv AgentService_NodeHeartbeatServer) error {
	return status.Errorf(codes.Unimplemented, "method NodeHeartbeat not implemented")
}
func (*UnimplementedAgentServiceServer) Sync(req *GeneralRequest, srv AgentService_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedAgentServiceServer) Init(ctx context.Context, req *GeneralRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedAgentServiceServer) GetAvailablePort(ctx context.Context, req *GetAvailablePortRequest) (*GetAvailablePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePort not implemented")
}
func (*UnimplementedAgentServiceServer) CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (*UnimplementedAgentServiceServer) StopService(ctx context.Context, req *StopServiceRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopService not implemented")
}
func (*UnimplementedAgentServiceServer) RemoveService(ctx context.Context, req *RemoveServiceRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_NodeHeartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GeneralRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).NodeHeartbeat(m, &agentServiceNodeHeartbeatServer{stream})
}

type AgentService_NodeHeartbeatServer interface {
	Send(*HeartbeatStreamServer) error
	grpc.ServerStream
}

type agentServiceNodeHeartbeatServer struct {
	grpc.ServerStream
}

func (x *agentServiceNodeHeartbeatServer) Send(m *HeartbeatStreamServer) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GeneralRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).Sync(m, &agentServiceSyncServer{stream})
}

type AgentService_SyncServer interface {
	Send(*SyncStreamServer) error
	grpc.ServerStream
}

type agentServiceSyncServer struct {
	grpc.ServerStream
}

func (x *agentServiceSyncServer) Send(m *SyncStreamServer) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Init(ctx, req.(*GeneralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_GetAvailablePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailablePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetAvailablePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentService/GetAvailablePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetAvailablePort(ctx, req.(*GetAvailablePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentService/StopService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgentService/RemoveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _AgentService_Init_Handler,
		},
		{
			MethodName: "GetAvailablePort",
			Handler:    _AgentService_GetAvailablePort_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _AgentService_CreateService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _AgentService_StopService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _AgentService_RemoveService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NodeHeartbeat",
			Handler:       _AgentService_NodeHeartbeat_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Sync",
			Handler:       _AgentService_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agent.proto",
}
