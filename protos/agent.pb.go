// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/agent.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
	return fileDescriptor_be14160ed34e84b5, []int{0, 0}
}

type ServiceEncryptionMethod_Enum int32

const (
	ServiceEncryptionMethod_NOT_SET                ServiceEncryptionMethod_Enum = 0
	ServiceEncryptionMethod_AES_256_CFB            ServiceEncryptionMethod_Enum = 1
	ServiceEncryptionMethod_AES_128_GCM            ServiceEncryptionMethod_Enum = 2
	ServiceEncryptionMethod_AES_192_GCM            ServiceEncryptionMethod_Enum = 3
	ServiceEncryptionMethod_AES_256_GCM            ServiceEncryptionMethod_Enum = 4
	ServiceEncryptionMethod_CHACHA20_IETF_POLY1305 ServiceEncryptionMethod_Enum = 5
	ServiceEncryptionMethod_AES_256_CTR            ServiceEncryptionMethod_Enum = 6
	ServiceEncryptionMethod_CHACHA20               ServiceEncryptionMethod_Enum = 7
	ServiceEncryptionMethod_CHACHA20_IETF          ServiceEncryptionMethod_Enum = 8
)

var ServiceEncryptionMethod_Enum_name = map[int32]string{
	0: "NOT_SET",
	1: "AES_256_CFB",
	2: "AES_128_GCM",
	3: "AES_192_GCM",
	4: "AES_256_GCM",
	5: "CHACHA20_IETF_POLY1305",
	6: "AES_256_CTR",
	7: "CHACHA20",
	8: "CHACHA20_IETF",
}

var ServiceEncryptionMethod_Enum_value = map[string]int32{
	"NOT_SET":                0,
	"AES_256_CFB":            1,
	"AES_128_GCM":            2,
	"AES_192_GCM":            3,
	"AES_256_GCM":            4,
	"CHACHA20_IETF_POLY1305": 5,
	"AES_256_CTR":            6,
	"CHACHA20":               7,
	"CHACHA20_IETF":          8,
}

func (x ServiceEncryptionMethod_Enum) String() string {
	return proto.EnumName(ServiceEncryptionMethod_Enum_name, int32(x))
}

func (ServiceEncryptionMethod_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{1, 0}
}

type ServiceStatus_Enum int32

const (
	ServiceStatus_NOT_SET ServiceStatus_Enum = 0
	ServiceStatus_RUNNING ServiceStatus_Enum = 1
	ServiceStatus_STOPPED ServiceStatus_Enum = 2
)

var ServiceStatus_Enum_name = map[int32]string{
	0: "NOT_SET",
	1: "RUNNING",
	2: "STOPPED",
}

var ServiceStatus_Enum_value = map[string]int32{
	"NOT_SET": 0,
	"RUNNING": 1,
	"STOPPED": 2,
}

func (x ServiceStatus_Enum) String() string {
	return proto.EnumName(ServiceStatus_Enum_name, int32(x))
}

func (ServiceStatus_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{2, 0}
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
	return fileDescriptor_be14160ed34e84b5, []int{0}
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

type ServiceEncryptionMethod struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceEncryptionMethod) Reset()         { *m = ServiceEncryptionMethod{} }
func (m *ServiceEncryptionMethod) String() string { return proto.CompactTextString(m) }
func (*ServiceEncryptionMethod) ProtoMessage()    {}
func (*ServiceEncryptionMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{1}
}

func (m *ServiceEncryptionMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEncryptionMethod.Unmarshal(m, b)
}
func (m *ServiceEncryptionMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEncryptionMethod.Marshal(b, m, deterministic)
}
func (m *ServiceEncryptionMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEncryptionMethod.Merge(m, src)
}
func (m *ServiceEncryptionMethod) XXX_Size() int {
	return xxx_messageInfo_ServiceEncryptionMethod.Size(m)
}
func (m *ServiceEncryptionMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEncryptionMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEncryptionMethod proto.InternalMessageInfo

type ServiceStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceStatus) Reset()         { *m = ServiceStatus{} }
func (m *ServiceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceStatus) ProtoMessage()    {}
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{2}
}

func (m *ServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatus.Unmarshal(m, b)
}
func (m *ServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatus.Marshal(b, m, deterministic)
}
func (m *ServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatus.Merge(m, src)
}
func (m *ServiceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceStatus.Size(m)
}
func (m *ServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatus proto.InternalMessageInfo

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
	return fileDescriptor_be14160ed34e84b5, []int{3}
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
	return fileDescriptor_be14160ed34e84b5, []int{4}
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
	return fileDescriptor_be14160ed34e84b5, []int{5}
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

type ServiceInfo struct {
	ServiceId            string             `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	UserId               string             `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContainerId          string             `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Port                 int32              `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Status               ServiceStatus_Enum `protobuf:"varint,5,opt,name=status,proto3,enum=ServiceStatus_Enum" json:"status,omitempty"`
	StatsResult          int64              `protobuf:"varint,6,opt,name=stats_result,json=statsResult,proto3" json:"stats_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServiceInfo) Reset()         { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()    {}
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{6}
}

func (m *ServiceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfo.Unmarshal(m, b)
}
func (m *ServiceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfo.Marshal(b, m, deterministic)
}
func (m *ServiceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfo.Merge(m, src)
}
func (m *ServiceInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceInfo.Size(m)
}
func (m *ServiceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfo proto.InternalMessageInfo

func (m *ServiceInfo) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ServiceInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ServiceInfo) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *ServiceInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServiceInfo) GetStatus() ServiceStatus_Enum {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_NOT_SET
}

func (m *ServiceInfo) GetStatsResult() int64 {
	if m != nil {
		return m.StatsResult
	}
	return 0
}

type SyncStreamServer struct {
	Services             []*ServiceInfo `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SyncStreamServer) Reset()         { *m = SyncStreamServer{} }
func (m *SyncStreamServer) String() string { return proto.CompactTextString(m) }
func (*SyncStreamServer) ProtoMessage()    {}
func (*SyncStreamServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{7}
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

func (m *SyncStreamServer) GetServices() []*ServiceInfo {
	if m != nil {
		return m.Services
	}
	return nil
}

type SyncRequest struct {
	SyncInterval         int64    `protobuf:"varint,1,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{8}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetSyncInterval() int64 {
	if m != nil {
		return m.SyncInterval
	}
	return 0
}

type CreateServiceRequest struct {
	Token                string                       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PortFrom             int32                        `protobuf:"varint,2,opt,name=port_from,json=portFrom,proto3" json:"port_from,omitempty"`
	PortTo               int32                        `protobuf:"varint,3,opt,name=port_to,json=portTo,proto3" json:"port_to,omitempty"`
	Type                 ServiceType_Enum             `protobuf:"varint,4,opt,name=type,proto3,enum=ServiceType_Enum" json:"type,omitempty"`
	EncryptionMethod     ServiceEncryptionMethod_Enum `protobuf:"varint,5,opt,name=encryption_method,json=encryptionMethod,proto3,enum=ServiceEncryptionMethod_Enum" json:"encryption_method,omitempty"`
	Password             string                       `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string                       `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{9}
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

func (m *CreateServiceRequest) GetEncryptionMethod() ServiceEncryptionMethod_Enum {
	if m != nil {
		return m.EncryptionMethod
	}
	return ServiceEncryptionMethod_NOT_SET
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
	return fileDescriptor_be14160ed34e84b5, []int{10}
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
	return fileDescriptor_be14160ed34e84b5, []int{11}
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
	return fileDescriptor_be14160ed34e84b5, []int{12}
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
	proto.RegisterEnum("ServiceEncryptionMethod_Enum", ServiceEncryptionMethod_Enum_name, ServiceEncryptionMethod_Enum_value)
	proto.RegisterEnum("ServiceStatus_Enum", ServiceStatus_Enum_name, ServiceStatus_Enum_value)
	proto.RegisterType((*ServiceType)(nil), "ServiceType")
	proto.RegisterType((*ServiceEncryptionMethod)(nil), "ServiceEncryptionMethod")
	proto.RegisterType((*ServiceStatus)(nil), "ServiceStatus")
	proto.RegisterType((*GeneralRequest)(nil), "GeneralRequest")
	proto.RegisterType((*GeneralResponse)(nil), "GeneralResponse")
	proto.RegisterType((*HeartbeatStreamServer)(nil), "HeartbeatStreamServer")
	proto.RegisterType((*ServiceInfo)(nil), "ServiceInfo")
	proto.RegisterType((*SyncStreamServer)(nil), "SyncStreamServer")
	proto.RegisterType((*SyncRequest)(nil), "SyncRequest")
	proto.RegisterType((*CreateServiceRequest)(nil), "CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "CreateServiceResponse")
	proto.RegisterType((*StopServiceRequest)(nil), "StopServiceRequest")
	proto.RegisterType((*RemoveServiceRequest)(nil), "RemoveServiceRequest")
}

func init() { proto.RegisterFile("protos/agent.proto", fileDescriptor_be14160ed34e84b5) }

var fileDescriptor_be14160ed34e84b5 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0xc7, 0xce, 0xff, 0x71, 0xd2, 0x3a, 0x7b, 0x4d, 0x1b, 0x05, 0x9d, 0x14, 0x16, 0x81, 0xc2,
	0xa1, 0x6e, 0x7a, 0x3e, 0x0e, 0xdd, 0x21, 0x74, 0x52, 0x1a, 0xd2, 0xd6, 0x70, 0x97, 0x56, 0x76,
	0x40, 0x82, 0x17, 0xcb, 0x4d, 0xf6, 0x72, 0x16, 0xf5, 0xae, 0xb1, 0x37, 0x45, 0xf9, 0x00, 0x7c,
	0x14, 0x3e, 0x08, 0xef, 0x3c, 0xf0, 0x91, 0xd0, 0xae, 0x5d, 0x93, 0xa4, 0x2e, 0xf0, 0x70, 0x4f,
	0xde, 0xf9, 0xcd, 0xce, 0x78, 0xe6, 0x37, 0xbf, 0xb1, 0x01, 0x45, 0x31, 0x17, 0x3c, 0x19, 0xfa,
	0x4b, 0xca, 0x04, 0x51, 0x06, 0x7e, 0x09, 0x86, 0x4b, 0xe3, 0xdb, 0x60, 0x4e, 0x67, 0xeb, 0x88,
	0xe2, 0x27, 0x50, 0x9e, 0xb0, 0x55, 0x88, 0x0c, 0xa8, 0x4d, 0x2f, 0x67, 0x9e, 0x3b, 0x99, 0x99,
	0x1f, 0xa0, 0x26, 0xd4, 0x5d, 0xd7, 0x7b, 0x6d, 0x9f, 0x4e, 0x7e, 0x30, 0x35, 0x54, 0x83, 0x92,
	0xeb, 0x3a, 0xa6, 0x8e, 0xff, 0xd0, 0xe0, 0x28, 0x8b, 0x9d, 0xb0, 0x79, 0xbc, 0x8e, 0x44, 0xc0,
	0xd9, 0x1b, 0x2a, 0xde, 0xf1, 0x05, 0xfe, 0x5d, 0x2b, 0x4a, 0xb4, 0x0f, 0xc6, 0x68, 0xe2, 0x7a,
	0xd6, 0xf3, 0x2f, 0xbd, 0xf1, 0xd9, 0xa9, 0xa9, 0xdd, 0x01, 0x4f, 0xad, 0x17, 0xde, 0xf9, 0xf8,
	0x8d, 0xa9, 0xe7, 0xc0, 0x4b, 0x4b, 0x01, 0xa5, 0xcd, 0x10, 0x09, 0x94, 0x51, 0x0f, 0x0e, 0xc7,
	0x17, 0xa3, 0xf1, 0xc5, 0xc8, 0x3a, 0xf1, 0xec, 0xc9, 0xec, 0xcc, 0xbb, 0xba, 0x7c, 0xfd, 0xe3,
	0xd3, 0x67, 0x27, 0xcf, 0xcd, 0xca, 0x56, 0xfe, 0x99, 0x63, 0x56, 0x65, 0xe5, 0x77, 0x97, 0xcd,
	0x1a, 0x6a, 0x43, 0x6b, 0x2b, 0xd4, 0xac, 0xe3, 0x57, 0xd0, 0xca, 0x5a, 0x70, 0x85, 0x2f, 0x56,
	0x09, 0x3e, 0x2e, 0xaa, 0xdb, 0x80, 0x9a, 0xf3, 0xfd, 0x74, 0x6a, 0x4f, 0xcf, 0x4d, 0x4d, 0x1a,
	0xee, 0xec, 0xf2, 0xea, 0x6a, 0xf2, 0x8d, 0xa9, 0xe3, 0x4f, 0x61, 0xef, 0x9c, 0x32, 0x1a, 0xfb,
	0x37, 0x0e, 0xfd, 0x65, 0x45, 0x13, 0x81, 0x0e, 0xa0, 0x22, 0xf8, 0xcf, 0x94, 0x75, 0xb5, 0xbe,
	0x36, 0x68, 0x38, 0xa9, 0x81, 0xdb, 0xb0, 0x9f, 0xdf, 0x4b, 0x22, 0xce, 0x12, 0x8a, 0x8f, 0xa0,
	0x73, 0x41, 0xfd, 0x58, 0x5c, 0x53, 0x5f, 0xb8, 0x22, 0xa6, 0x7e, 0x28, 0x2b, 0xa1, 0x31, 0xfe,
	0x53, 0xcb, 0x67, 0x62, 0xb3, 0xb7, 0x1c, 0x3d, 0x06, 0x48, 0x52, 0xd3, 0x0b, 0x16, 0x59, 0xda,
	0x46, 0x86, 0xd8, 0x0b, 0x74, 0x04, 0xb5, 0x55, 0x42, 0x63, 0xe9, 0xd3, 0x95, 0xaf, 0x2a, 0x4d,
	0x7b, 0x81, 0x3e, 0x82, 0xe6, 0x9c, 0x33, 0xe1, 0x07, 0x2c, 0xf5, 0x96, 0x94, 0xd7, 0xc8, 0x31,
	0x7b, 0x81, 0x10, 0x94, 0x23, 0x1e, 0x8b, 0x6e, 0xb9, 0xaf, 0x0d, 0x2a, 0x8e, 0x3a, 0xa3, 0xcf,
	0xa1, 0x9a, 0x28, 0x2e, 0xba, 0x95, 0xbe, 0x36, 0xd8, 0xb3, 0x1e, 0x91, 0x2d, 0x86, 0x88, 0xa4,
	0xc7, 0xc9, 0xae, 0xc8, 0x77, 0xc8, 0x53, 0xe2, 0xc5, 0x34, 0x59, 0xdd, 0x88, 0x6e, 0xb5, 0xaf,
	0x0d, 0x4a, 0x8e, 0xa1, 0x30, 0x47, 0x41, 0xf8, 0x6b, 0x30, 0xdd, 0x35, 0x9b, 0x6f, 0xb6, 0x88,
	0x06, 0x50, 0xcf, 0x1a, 0x48, 0xba, 0x5a, 0xbf, 0x34, 0x30, 0xac, 0x26, 0xd9, 0x68, 0xd9, 0xc9,
	0xbd, 0xd8, 0x02, 0x43, 0x46, 0xdf, 0xb1, 0xfb, 0x31, 0xb4, 0x92, 0x35, 0x9b, 0x7b, 0x01, 0x13,
	0x34, 0xbe, 0xf5, 0x6f, 0x14, 0x1d, 0x25, 0xa7, 0x29, 0x41, 0x3b, 0xc3, 0xf0, 0x6f, 0x3a, 0x1c,
	0x8c, 0x63, 0xea, 0x0b, 0x9a, 0xe5, 0xfc, 0xd7, 0xd9, 0xa0, 0x0f, 0xa1, 0x21, 0x1b, 0xf7, 0xde,
	0xc6, 0x3c, 0x54, 0x14, 0x56, 0x9c, 0xba, 0x04, 0xce, 0x62, 0x1e, 0x4a, 0x76, 0x95, 0x53, 0x70,
	0xc5, 0x5f, 0xc5, 0xa9, 0x4a, 0x73, 0xc6, 0xd1, 0x27, 0x50, 0x16, 0xeb, 0x88, 0x2a, 0xea, 0xf6,
	0xac, 0x36, 0xd9, 0xd8, 0xa2, 0x94, 0x22, 0xe5, 0x46, 0xdf, 0x42, 0x9b, 0xe6, 0xcb, 0xe1, 0x85,
	0x6a, 0x3b, 0x32, 0x62, 0x1f, 0x93, 0x07, 0xb6, 0x27, 0x8d, 0x37, 0xe9, 0x0e, 0x8c, 0x7a, 0x50,
	0x8f, 0xfc, 0x24, 0xf9, 0x95, 0xc7, 0x0b, 0x45, 0x74, 0xc3, 0xc9, 0x6d, 0x39, 0x49, 0xe6, 0x87,
	0xb4, 0x5b, 0x53, 0xb8, 0x3a, 0xe3, 0x10, 0x3a, 0x3b, 0x34, 0xa4, 0xd2, 0xfb, 0x2f, 0x45, 0xed,
	0x0a, 0x47, 0x7f, 0x58, 0x38, 0xa5, 0x7f, 0x84, 0x83, 0x6d, 0x40, 0xae, 0xe0, 0xd1, 0xff, 0xe2,
	0x7c, 0xbb, 0x02, 0x7d, 0xa7, 0x02, 0xfc, 0x1d, 0x1c, 0x38, 0x34, 0xe4, 0xb7, 0xf4, 0x3d, 0x24,
	0xb3, 0xfe, 0xd2, 0xa1, 0x39, 0x92, 0x9f, 0xbc, 0x2c, 0x19, 0xfa, 0x0a, 0x5a, 0x53, 0xbe, 0xa0,
	0xf9, 0xf6, 0xa1, 0x7d, 0xb2, 0xbd, 0xc4, 0xbd, 0x43, 0x52, 0xb8, 0x9a, 0x27, 0x1a, 0xfa, 0x0c,
	0xca, 0x52, 0x8f, 0xa8, 0x49, 0x36, 0x64, 0xd9, 0x6b, 0x93, 0x5d, 0x89, 0xa7, 0x57, 0x6d, 0x16,
	0x14, 0x64, 0x37, 0xc9, 0xce, 0xb7, 0x00, 0xbd, 0x82, 0xd6, 0xd6, 0xa4, 0x50, 0x87, 0x14, 0x09,
	0xb8, 0x77, 0x48, 0x8a, 0x07, 0xfa, 0x05, 0x18, 0x1b, 0xd4, 0xa3, 0x47, 0xe4, 0xfe, 0x20, 0x0a,
	0xde, 0xfa, 0x02, 0x5a, 0x5b, 0x2c, 0xa3, 0x0e, 0x29, 0x62, 0xfd, 0x7e, 0xe4, 0xe9, 0x93, 0x9f,
	0x06, 0xcb, 0x40, 0xbc, 0x5b, 0x5d, 0x93, 0x39, 0x0f, 0x87, 0x4b, 0x7e, 0x1c, 0x2c, 0x59, 0x20,
	0xe8, 0x30, 0x7d, 0x1c, 0xab, 0x1f, 0xcc, 0x30, 0xfd, 0xdb, 0x5c, 0x57, 0xd5, 0xf3, 0xd9, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0xdb, 0x56, 0x0f, 0x7e, 0x06, 0x00, 0x00,
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
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (AgentService_SyncClient, error)
	Init(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
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

func (c *agentServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (AgentService_SyncClient, error) {
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
	Sync(*SyncRequest, AgentService_SyncServer) error
	Init(context.Context, *GeneralRequest) (*GeneralResponse, error)
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	StopService(context.Context, *StopServiceRequest) (*GeneralResponse, error)
	RemoveService(context.Context, *RemoveServiceRequest) (*GeneralResponse, error)
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
	m := new(SyncRequest)
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
	Metadata: "protos/agent.proto",
}
