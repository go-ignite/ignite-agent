// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/agent.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{3}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{4}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type HeartbeatRequest struct {
	Interval             *duration.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{5}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

type HeartbeatStreamServer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatStreamServer) Reset()         { *m = HeartbeatStreamServer{} }
func (m *HeartbeatStreamServer) String() string { return proto.CompactTextString(m) }
func (*HeartbeatStreamServer) ProtoMessage()    {}
func (*HeartbeatStreamServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{6}
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
	ContainerName        string             `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	ContainerId          string             `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Port                 int32              `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Status               ServiceStatus_Enum `protobuf:"varint,4,opt,name=status,proto3,enum=ServiceStatus_Enum" json:"status,omitempty"`
	StatsResult          int64              `protobuf:"varint,5,opt,name=stats_result,json=statsResult,proto3" json:"stats_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServiceInfo) Reset()         { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()    {}
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{7}
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

func (m *ServiceInfo) GetContainerName() string {
	if m != nil {
		return m.ContainerName
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
	return fileDescriptor_be14160ed34e84b5, []int{8}
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
	SyncInterval         *duration.Duration `protobuf:"bytes,1,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	NodeId               string             `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{9}
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

func (m *SyncRequest) GetSyncInterval() *duration.Duration {
	if m != nil {
		return m.SyncInterval
	}
	return nil
}

func (m *SyncRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type CreateServiceRequest struct {
	PortFrom             int32                        `protobuf:"varint,1,opt,name=port_from,json=portFrom,proto3" json:"port_from,omitempty"`
	PortTo               int32                        `protobuf:"varint,2,opt,name=port_to,json=portTo,proto3" json:"port_to,omitempty"`
	Type                 ServiceType_Enum             `protobuf:"varint,3,opt,name=type,proto3,enum=ServiceType_Enum" json:"type,omitempty"`
	EncryptionMethod     ServiceEncryptionMethod_Enum `protobuf:"varint,4,opt,name=encryption_method,json=encryptionMethod,proto3,enum=ServiceEncryptionMethod_Enum" json:"encryption_method,omitempty"`
	Password             string                       `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	UserId               string                       `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NodeId               string                       `protobuf:"bytes,7,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{10}
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

func (m *CreateServiceRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateServiceRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type CreateServiceResponse struct {
	ContainerName        string   `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
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
	return fileDescriptor_be14160ed34e84b5, []int{11}
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

func (m *CreateServiceResponse) GetContainerName() string {
	if m != nil {
		return m.ContainerName
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
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopServiceRequest) Reset()         { *m = StopServiceRequest{} }
func (m *StopServiceRequest) String() string { return proto.CompactTextString(m) }
func (*StopServiceRequest) ProtoMessage()    {}
func (*StopServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{12}
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

func (m *StopServiceRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type RemoveServiceRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServiceRequest) Reset()         { *m = RemoveServiceRequest{} }
func (m *RemoveServiceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveServiceRequest) ProtoMessage()    {}
func (*RemoveServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be14160ed34e84b5, []int{13}
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

func (m *RemoveServiceRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
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
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
	proto.RegisterType((*HeartbeatRequest)(nil), "HeartbeatRequest")
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
	// 803 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc7, 0x71, 0xbe, 0xf7, 0x38, 0x49, 0x9d, 0x69, 0xb3, 0x1b, 0x82, 0x40, 0x61, 0x24, 0xa4,
	0x50, 0xb4, 0x93, 0xd4, 0xd5, 0x96, 0xae, 0x84, 0x2a, 0x65, 0x53, 0x6f, 0xd7, 0xa8, 0xcd, 0xae,
	0xec, 0x80, 0x04, 0x37, 0x96, 0x13, 0x4f, 0x52, 0x4b, 0x6b, 0x8f, 0xb1, 0xc7, 0x8b, 0xf2, 0x10,
	0x3c, 0x03, 0x57, 0x3c, 0x05, 0x57, 0xbc, 0x19, 0x9a, 0xb1, 0xe3, 0x26, 0x69, 0x90, 0xe0, 0x82,
	0x2b, 0xfb, 0xfc, 0xe7, 0x9c, 0x33, 0x33, 0xbf, 0x73, 0x8e, 0x0d, 0x28, 0x8a, 0x19, 0x67, 0xc9,
	0xc8, 0x5d, 0xd3, 0x90, 0x13, 0x69, 0xf4, 0xbf, 0x58, 0x33, 0xb6, 0xbe, 0xa7, 0x23, 0x69, 0x2d,
	0xd2, 0xd5, 0xc8, 0x4b, 0x63, 0x97, 0xfb, 0x2c, 0xcc, 0xd6, 0xf1, 0x25, 0xa8, 0x36, 0x8d, 0x1f,
	0xfc, 0x25, 0x9d, 0x6f, 0x22, 0x8a, 0x9f, 0x42, 0xc5, 0x08, 0xd3, 0x00, 0xa9, 0x50, 0x9f, 0xdd,
	0xce, 0x1d, 0xdb, 0x98, 0x6b, 0x9f, 0xa0, 0x26, 0x34, 0x6c, 0xdb, 0x79, 0x6b, 0x5e, 0x19, 0x3f,
	0x6a, 0x0a, 0xaa, 0x43, 0xd9, 0xb6, 0x2d, 0xad, 0x84, 0xff, 0x52, 0xe0, 0x2c, 0x8f, 0x35, 0xc2,
	0x65, 0xbc, 0x89, 0x44, 0xda, 0x77, 0x94, 0xbf, 0x67, 0x1e, 0xfe, 0x43, 0x39, 0x96, 0xe8, 0x11,
	0xa8, 0x13, 0xc3, 0x76, 0xf4, 0x8b, 0x17, 0xce, 0xf4, 0xfa, 0x4a, 0x53, 0xb6, 0xc2, 0x33, 0xfd,
	0xa5, 0xf3, 0x66, 0xfa, 0x4e, 0x2b, 0x15, 0xc2, 0xa5, 0x2e, 0x85, 0xf2, 0x6e, 0x88, 0x10, 0x2a,
	0xa8, 0x0f, 0xa7, 0xd3, 0x9b, 0xc9, 0xf4, 0x66, 0xa2, 0x8f, 0x1d, 0xd3, 0x98, 0x5f, 0x3b, 0x77,
	0xb7, 0x6f, 0x7f, 0x7a, 0xf6, 0x7c, 0x7c, 0xa1, 0x55, 0xf7, 0xf2, 0xcf, 0x2d, 0xad, 0x26, 0x4e,
	0xbe, 0x75, 0xd6, 0xea, 0xa8, 0x03, 0xad, 0xbd, 0x50, 0xad, 0x81, 0x5f, 0x41, 0x2b, 0xbf, 0x82,
	0xcd, 0x5d, 0x9e, 0x26, 0xf8, 0xfc, 0xd8, 0xb9, 0x55, 0xa8, 0x5b, 0x3f, 0xcc, 0x66, 0xe6, 0xec,
	0x8d, 0xa6, 0x08, 0xc3, 0x9e, 0xdf, 0xde, 0xdd, 0x19, 0xaf, 0xb5, 0x12, 0x6e, 0x43, 0xd3, 0x08,
	0x22, 0xbe, 0xb1, 0xe8, 0x2f, 0x29, 0x4d, 0x38, 0x7e, 0x04, 0xad, 0xdc, 0x4e, 0x22, 0x16, 0x26,
	0x14, 0x9b, 0xa0, 0xdd, 0x50, 0x37, 0xe6, 0x0b, 0xea, 0xf2, 0xdc, 0x09, 0x5d, 0x40, 0xc3, 0x0f,
	0x39, 0x8d, 0x1f, 0xdc, 0xfb, 0x9e, 0x32, 0x50, 0x86, 0xaa, 0xfe, 0x29, 0xc9, 0xca, 0x44, 0xb6,
	0x65, 0x22, 0xaf, 0xf3, 0x32, 0x59, 0x85, 0x2b, 0x3e, 0x83, 0x6e, 0x91, 0xca, 0xe6, 0x31, 0x75,
	0x03, 0x71, 0x74, 0x1a, 0xe3, 0x3f, 0x95, 0xa2, 0x88, 0x66, 0xb8, 0x62, 0xe8, 0x2b, 0x68, 0x2f,
	0x59, 0xc8, 0x5d, 0x3f, 0xa4, 0xb1, 0x13, 0xba, 0x01, 0x95, 0xbb, 0x9c, 0x58, 0xad, 0x42, 0x9d,
	0xb9, 0x01, 0x45, 0x5f, 0x42, 0xf3, 0x83, 0x9b, 0xef, 0xf5, 0x4a, 0xd2, 0x49, 0x2d, 0x34, 0xd3,
	0x43, 0x08, 0x2a, 0x11, 0x8b, 0x79, 0xaf, 0x3c, 0x50, 0x86, 0x55, 0x4b, 0xbe, 0xa3, 0x6f, 0xa0,
	0x96, 0x48, 0x56, 0xbd, 0xca, 0x40, 0x19, 0xb6, 0xf5, 0xc7, 0x64, 0x8f, 0x20, 0x11, 0xf8, 0xac,
	0xdc, 0x45, 0xec, 0x21, 0xde, 0x12, 0x27, 0xa6, 0x49, 0x7a, 0xcf, 0x7b, 0xd5, 0x81, 0x32, 0x2c,
	0x5b, 0xaa, 0xd4, 0x2c, 0x29, 0xe1, 0xef, 0x40, 0xb3, 0x37, 0xe1, 0x72, 0xf7, 0x46, 0x68, 0x08,
	0x8d, 0x24, 0x4b, 0x9a, 0xf4, 0x94, 0x41, 0x79, 0xa8, 0xea, 0x4d, 0xb2, 0x73, 0x43, 0xab, 0x58,
	0xc5, 0x2b, 0x50, 0x45, 0xf4, 0x16, 0xed, 0x2b, 0x68, 0x25, 0x9b, 0x70, 0xe9, 0xfc, 0x7b, 0xbe,
	0x4d, 0xe1, 0x6f, 0xe6, 0xee, 0xe8, 0x0c, 0xea, 0x21, 0xf3, 0xe8, 0x07, 0x1c, 0x35, 0x61, 0x9a,
	0x1e, 0xfe, 0xad, 0x04, 0x4f, 0xa6, 0x31, 0x75, 0x39, 0xcd, 0xcf, 0xb1, 0xdd, 0xf1, 0x33, 0x38,
	0x11, 0x58, 0x9c, 0x55, 0xcc, 0x02, 0xb9, 0x5b, 0xd5, 0x6a, 0x08, 0xe1, 0x3a, 0x66, 0x81, 0x48,
	0x27, 0x17, 0x39, 0x93, 0xe9, 0xaa, 0x56, 0x4d, 0x98, 0x73, 0x51, 0xa2, 0x0a, 0xdf, 0x44, 0x54,
	0x82, 0x6d, 0xeb, 0x1d, 0xb2, 0x33, 0x83, 0x19, 0x40, 0xb9, 0x8c, 0xbe, 0x87, 0x0e, 0x2d, 0x46,
	0xcb, 0x09, 0xe4, 0x6c, 0xe5, 0xd8, 0x3f, 0x27, 0xff, 0x30, 0x7b, 0x59, 0xbc, 0x46, 0x0f, 0x64,
	0xd4, 0x87, 0x46, 0xe4, 0x26, 0xc9, 0xaf, 0x2c, 0xf6, 0x64, 0x19, 0x4e, 0xac, 0xc2, 0x16, 0xe7,
	0x4c, 0x93, 0xac, 0x0b, 0x6a, 0xd9, 0xb5, 0x85, 0x69, 0x7a, 0xbb, 0x3c, 0xea, 0x7b, 0x3c, 0x52,
	0xe8, 0x1e, 0xe0, 0xc8, 0x1a, 0xfe, 0xff, 0x6d, 0x3e, 0xfc, 0x2d, 0x20, 0x9b, 0xb3, 0xe8, 0xa0,
	0x06, 0x87, 0xc9, 0x94, 0x8f, 0x92, 0xe1, 0x4b, 0x78, 0x62, 0xd1, 0x80, 0x3d, 0xd0, 0xff, 0x1c,
	0xaa, 0xff, 0x5e, 0x82, 0xe6, 0x44, 0x7c, 0x52, 0xf3, 0x50, 0xf4, 0x12, 0x4e, 0x8a, 0x41, 0x44,
	0x1d, 0x72, 0x38, 0xdf, 0xfd, 0x53, 0x72, 0x74, 0x4e, 0xc7, 0x0a, 0xfa, 0x1a, 0x2a, 0xa2, 0x5b,
	0x51, 0x93, 0xec, 0x34, 0x6d, 0xbf, 0x43, 0x0e, 0x07, 0x60, 0xac, 0x88, 0x4e, 0xde, 0x03, 0x8c,
	0xba, 0xe4, 0x58, 0xff, 0xf5, 0x4f, 0xc9, 0xf1, 0x3a, 0xe8, 0xa0, 0xee, 0x90, 0x42, 0x8f, 0xc9,
	0xc7, 0xdc, 0xfa, 0x6d, 0xb2, 0xf7, 0xb1, 0x42, 0x2f, 0xa0, 0xb5, 0x07, 0x09, 0x75, 0xc9, 0x31,
	0x68, 0x87, 0x71, 0x57, 0x4f, 0x7f, 0x1e, 0xae, 0x7d, 0xfe, 0x3e, 0x5d, 0x90, 0x25, 0x0b, 0x46,
	0x6b, 0x76, 0xee, 0xaf, 0x43, 0x9f, 0xd3, 0x51, 0xf6, 0x38, 0x97, 0xff, 0xa3, 0xec, 0x0f, 0x94,
	0x2c, 0x6a, 0xf2, 0xf9, 0xfc, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x00, 0x5f, 0xb6, 0xad,
	0x06, 0x00, 0x00,
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
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (AgentService_HeartbeatClient, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (AgentService_SyncClient, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (AgentService_HeartbeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentService_serviceDesc.Streams[0], "/AgentService/Heartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceHeartbeatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_HeartbeatClient interface {
	Recv() (*HeartbeatStreamServer, error)
	grpc.ClientStream
}

type agentServiceHeartbeatClient struct {
	grpc.ClientStream
}

func (x *agentServiceHeartbeatClient) Recv() (*HeartbeatStreamServer, error) {
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

func (c *agentServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/AgentService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/AgentService/StopService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/AgentService/RemoveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	Heartbeat(*HeartbeatRequest, AgentService_HeartbeatServer) error
	Sync(*SyncRequest, AgentService_SyncServer) error
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	StopService(context.Context, *StopServiceRequest) (*EmptyResponse, error)
	RemoveService(context.Context, *RemoveServiceRequest) (*EmptyResponse, error)
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_Heartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HeartbeatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).Heartbeat(m, &agentServiceHeartbeatServer{stream})
}

type AgentService_HeartbeatServer interface {
	Send(*HeartbeatStreamServer) error
	grpc.ServerStream
}

type agentServiceHeartbeatServer struct {
	grpc.ServerStream
}

func (x *agentServiceHeartbeatServer) Send(m *HeartbeatStreamServer) error {
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
			StreamName:    "Heartbeat",
			Handler:       _AgentService_Heartbeat_Handler,
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
