// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v4alpha/server_info.proto

package envoy_admin_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ServerInfo_State int32

const (
	ServerInfo_LIVE             ServerInfo_State = 0
	ServerInfo_DRAINING         ServerInfo_State = 1
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	ServerInfo_INITIALIZING     ServerInfo_State = 3
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
	2: "PRE_INITIALIZING",
	3: "INITIALIZING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":             0,
	"DRAINING":         1,
	"PRE_INITIALIZING": 2,
	"INITIALIZING":     3,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e5b029e0fc1706, []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

var CommandLineOptions_IpVersion_name = map[int32]string{
	0: "v4",
	1: "v6",
}

var CommandLineOptions_IpVersion_value = map[string]int32{
	"v4": 0,
	"v6": 1,
}

func (x CommandLineOptions_IpVersion) String() string {
	return proto.EnumName(CommandLineOptions_IpVersion_name, int32(x))
}

func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e5b029e0fc1706, []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	CommandLineOptions_Serve    CommandLineOptions_Mode = 0
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

var CommandLineOptions_Mode_name = map[int32]string{
	0: "Serve",
	1: "Validate",
	2: "InitOnly",
}

var CommandLineOptions_Mode_value = map[string]int32{
	"Serve":    0,
	"Validate": 1,
	"InitOnly": 2,
}

func (x CommandLineOptions_Mode) String() string {
	return proto.EnumName(CommandLineOptions_Mode_name, int32(x))
}

func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e5b029e0fc1706, []int{1, 1}
}

type ServerInfo struct {
	Version              string              `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	State                ServerInfo_State    `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v4alpha.ServerInfo_State" json:"state,omitempty"`
	UptimeCurrentEpoch   *duration.Duration  `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	UptimeAllEpochs      *duration.Duration  `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	HotRestartVersion    string              `protobuf:"bytes,5,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	CommandLineOptions   *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e5b029e0fc1706, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetHotRestartVersion() string {
	if m != nil {
		return m.HotRestartVersion
	}
	return ""
}

func (m *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if m != nil {
		return m.CommandLineOptions
	}
	return nil
}

type CommandLineOptions struct {
	BaseId                     uint64                       `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Concurrency                uint32                       `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	ConfigPath                 string                       `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	ConfigYaml                 string                       `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	AllowUnknownStaticFields   bool                         `protobuf:"varint,5,opt,name=allow_unknown_static_fields,json=allowUnknownStaticFields,proto3" json:"allow_unknown_static_fields,omitempty"`
	RejectUnknownDynamicFields bool                         `protobuf:"varint,26,opt,name=reject_unknown_dynamic_fields,json=rejectUnknownDynamicFields,proto3" json:"reject_unknown_dynamic_fields,omitempty"`
	AdminAddressPath           string                       `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	LocalAddressIpVersion      CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v4alpha.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	LogLevel                   string                       `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	ComponentLogLevel          string                       `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	LogFormat                  string                       `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	LogFormatEscaped           bool                         `protobuf:"varint,27,opt,name=log_format_escaped,json=logFormatEscaped,proto3" json:"log_format_escaped,omitempty"`
	LogPath                    string                       `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	ServiceCluster             string                       `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	ServiceNode                string                       `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	ServiceZone                string                       `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	FileFlushInterval          *duration.Duration           `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	DrainTime                  *duration.Duration           `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	ParentShutdownTime         *duration.Duration           `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	Mode                       CommandLineOptions_Mode      `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v4alpha.CommandLineOptions_Mode" json:"mode,omitempty"`
	DisableHotRestart          bool                         `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	EnableMutexTracing         bool                         `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	RestartEpoch               uint32                       `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	CpusetThreads              bool                         `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	DisabledExtensions         []string                     `protobuf:"bytes,28,rep,name=disabled_extensions,json=disabledExtensions,proto3" json:"disabled_extensions,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                     `json:"-"`
	XXX_unrecognized           []byte                       `json:"-"`
	XXX_sizecache              int32                        `json:"-"`
}

func (m *CommandLineOptions) Reset()         { *m = CommandLineOptions{} }
func (m *CommandLineOptions) String() string { return proto.CompactTextString(m) }
func (*CommandLineOptions) ProtoMessage()    {}
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e5b029e0fc1706, []int{1}
}

func (m *CommandLineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineOptions.Unmarshal(m, b)
}
func (m *CommandLineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineOptions.Marshal(b, m, deterministic)
}
func (m *CommandLineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineOptions.Merge(m, src)
}
func (m *CommandLineOptions) XXX_Size() int {
	return xxx_messageInfo_CommandLineOptions.Size(m)
}
func (m *CommandLineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineOptions proto.InternalMessageInfo

func (m *CommandLineOptions) GetBaseId() uint64 {
	if m != nil {
		return m.BaseId
	}
	return 0
}

func (m *CommandLineOptions) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *CommandLineOptions) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *CommandLineOptions) GetConfigYaml() string {
	if m != nil {
		return m.ConfigYaml
	}
	return ""
}

func (m *CommandLineOptions) GetAllowUnknownStaticFields() bool {
	if m != nil {
		return m.AllowUnknownStaticFields
	}
	return false
}

func (m *CommandLineOptions) GetRejectUnknownDynamicFields() bool {
	if m != nil {
		return m.RejectUnknownDynamicFields
	}
	return false
}

func (m *CommandLineOptions) GetAdminAddressPath() string {
	if m != nil {
		return m.AdminAddressPath
	}
	return ""
}

func (m *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if m != nil {
		return m.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (m *CommandLineOptions) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetComponentLogLevel() string {
	if m != nil {
		return m.ComponentLogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormatEscaped() bool {
	if m != nil {
		return m.LogFormatEscaped
	}
	return false
}

func (m *CommandLineOptions) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *CommandLineOptions) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *CommandLineOptions) GetServiceNode() string {
	if m != nil {
		return m.ServiceNode
	}
	return ""
}

func (m *CommandLineOptions) GetServiceZone() string {
	if m != nil {
		return m.ServiceZone
	}
	return ""
}

func (m *CommandLineOptions) GetFileFlushInterval() *duration.Duration {
	if m != nil {
		return m.FileFlushInterval
	}
	return nil
}

func (m *CommandLineOptions) GetDrainTime() *duration.Duration {
	if m != nil {
		return m.DrainTime
	}
	return nil
}

func (m *CommandLineOptions) GetParentShutdownTime() *duration.Duration {
	if m != nil {
		return m.ParentShutdownTime
	}
	return nil
}

func (m *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return CommandLineOptions_Serve
}

func (m *CommandLineOptions) GetDisableHotRestart() bool {
	if m != nil {
		return m.DisableHotRestart
	}
	return false
}

func (m *CommandLineOptions) GetEnableMutexTracing() bool {
	if m != nil {
		return m.EnableMutexTracing
	}
	return false
}

func (m *CommandLineOptions) GetRestartEpoch() uint32 {
	if m != nil {
		return m.RestartEpoch
	}
	return 0
}

func (m *CommandLineOptions) GetCpusetThreads() bool {
	if m != nil {
		return m.CpusetThreads
	}
	return false
}

func (m *CommandLineOptions) GetDisabledExtensions() []string {
	if m != nil {
		return m.DisabledExtensions
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.admin.v4alpha.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterEnum("envoy.admin.v4alpha.CommandLineOptions_IpVersion", CommandLineOptions_IpVersion_name, CommandLineOptions_IpVersion_value)
	proto.RegisterEnum("envoy.admin.v4alpha.CommandLineOptions_Mode", CommandLineOptions_Mode_name, CommandLineOptions_Mode_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v4alpha.ServerInfo")
	proto.RegisterType((*CommandLineOptions)(nil), "envoy.admin.v4alpha.CommandLineOptions")
}

func init() {
	proto.RegisterFile("envoy/admin/v4alpha/server_info.proto", fileDescriptor_b9e5b029e0fc1706)
}

var fileDescriptor_b9e5b029e0fc1706 = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0xc7, 0x63, 0x30, 0xc6, 0x3e, 0x7c, 0x2d, 0x03, 0x24, 0x0b, 0x94, 0xd4, 0x80, 0x50, 0xb8,
	0x48, 0xed, 0x36, 0x89, 0xaa, 0x2a, 0x55, 0xab, 0x12, 0x30, 0xe9, 0x52, 0x42, 0xd0, 0x42, 0x91,
	0x92, 0x9b, 0xd1, 0xb0, 0x3b, 0xb6, 0x97, 0xce, 0xce, 0xac, 0x76, 0x67, 0x1d, 0xdc, 0xab, 0x5e,
	0xf6, 0x19, 0xfa, 0x28, 0xbd, 0xaf, 0xd4, 0xdb, 0x3e, 0x46, 0x2f, 0xfa, 0x0e, 0xd5, 0x9c, 0x59,
	0x7f, 0x10, 0x90, 0xc8, 0x15, 0xcc, 0x39, 0xbf, 0xf3, 0xf7, 0xcc, 0x39, 0x67, 0xff, 0xb0, 0xc3,
	0x65, 0x4f, 0xf5, 0x9b, 0x2c, 0x8c, 0x23, 0xd9, 0xec, 0xbd, 0x60, 0x22, 0xe9, 0xb2, 0x66, 0xc6,
	0xd3, 0x1e, 0x4f, 0x69, 0x24, 0xdb, 0xaa, 0x91, 0xa4, 0x4a, 0x2b, 0xb2, 0x84, 0x58, 0x03, 0xb1,
	0x46, 0x81, 0xad, 0x3d, 0xee, 0x28, 0xd5, 0x11, 0xbc, 0x89, 0xc8, 0x65, 0xde, 0x6e, 0x86, 0x79,
	0xca, 0x74, 0xa4, 0xa4, 0x2d, 0x5a, 0xdb, 0x2e, 0xb4, 0xa5, 0x54, 0x1a, 0xe3, 0x59, 0x33, 0xe4,
	0x49, 0xca, 0x83, 0x71, 0x68, 0x23, 0x0f, 0x13, 0x76, 0x83, 0xc9, 0x34, 0xd3, 0x79, 0x56, 0xa4,
	0x37, 0x6f, 0xa5, 0x7b, 0x3c, 0xcd, 0x22, 0x25, 0x23, 0xd9, 0xb1, 0xc8, 0xd6, 0x7f, 0x93, 0x00,
	0x67, 0x78, 0x63, 0x4f, 0xb6, 0x15, 0x71, 0x61, 0xba, 0x40, 0xdc, 0x52, 0xbd, 0xb4, 0x5b, 0xf3,
	0x07, 0x47, 0xf2, 0x2d, 0x4c, 0x19, 0x6d, 0xee, 0x4e, 0xd4, 0x4b, 0xbb, 0xf3, 0xcf, 0x76, 0x1a,
	0x77, 0x3c, 0xaa, 0x31, 0x52, 0x6a, 0x9c, 0x19, 0xd8, 0xb7, 0x35, 0xe4, 0x27, 0x58, 0xce, 0x13,
	0x1d, 0xc5, 0x9c, 0x06, 0x79, 0x9a, 0x72, 0xa9, 0x29, 0x4f, 0x54, 0xd0, 0x75, 0x27, 0xeb, 0xa5,
	0xdd, 0x99, 0x67, 0xab, 0x0d, 0xdb, 0x8b, 0xc6, 0xa0, 0x17, 0x8d, 0x83, 0xa2, 0x17, 0x3e, 0xb1,
	0x65, 0xfb, 0xb6, 0xaa, 0x65, 0x8a, 0x48, 0x0b, 0x16, 0x0b, 0x31, 0x26, 0x84, 0x15, 0xca, 0xdc,
	0xf2, 0x7d, 0x4a, 0x0b, 0xb6, 0x66, 0x4f, 0x08, 0x54, 0xc9, 0x48, 0x03, 0x96, 0xba, 0x4a, 0xd3,
	0x94, 0x67, 0x9a, 0xa5, 0x9a, 0x0e, 0x9e, 0x3d, 0x85, 0xcf, 0x5e, 0xec, 0x2a, 0xed, 0xdb, 0xcc,
	0x45, 0xd1, 0x80, 0x77, 0xb0, 0x1c, 0xa8, 0x38, 0x66, 0x32, 0xa4, 0x22, 0x92, 0x9c, 0xaa, 0x04,
	0x5b, 0xea, 0x56, 0xf0, 0x97, 0x9f, 0xdc, 0xd9, 0x8f, 0x7d, 0x5b, 0x70, 0x1c, 0x49, 0xfe, 0xd6,
	0xe2, 0x3e, 0x09, 0x6e, 0xc5, 0xb6, 0x5e, 0xc3, 0x14, 0xb6, 0x8b, 0x54, 0xa1, 0x7c, 0xec, 0x5d,
	0xb4, 0x9c, 0x07, 0x64, 0x16, 0xaa, 0x07, 0xfe, 0x9e, 0x77, 0xe2, 0x9d, 0xbc, 0x76, 0x4a, 0x64,
	0x19, 0x9c, 0x53, 0xbf, 0x45, 0xbd, 0x13, 0xef, 0xdc, 0xdb, 0x3b, 0xf6, 0xde, 0x9b, 0xe8, 0x04,
	0x71, 0x60, 0xf6, 0x46, 0x64, 0xf2, 0x65, 0xfd, 0x8f, 0xbf, 0x7e, 0x7f, 0xbc, 0x0e, 0xab, 0x37,
	0xee, 0xf2, 0x7c, 0x6c, 0x2c, 0x5b, 0xff, 0x02, 0x90, 0xdb, 0xb7, 0x22, 0x8f, 0x60, 0xfa, 0x92,
	0x65, 0x9c, 0x46, 0x21, 0xce, 0xbd, 0xec, 0x57, 0xcc, 0xd1, 0x0b, 0x49, 0x1d, 0x66, 0x02, 0x25,
	0xed, 0xd4, 0x82, 0x3e, 0x0e, 0x7f, 0xce, 0x1f, 0x0f, 0x91, 0xcf, 0x91, 0x68, 0x47, 0x1d, 0x9a,
	0x30, 0x6d, 0x47, 0x5a, 0xf3, 0xc1, 0x86, 0x4e, 0x99, 0xee, 0x8e, 0x01, 0x7d, 0x16, 0x0b, 0x9c,
	0xd4, 0x10, 0x78, 0xc7, 0x62, 0x41, 0xbe, 0x83, 0x75, 0x26, 0x84, 0xfa, 0x40, 0x73, 0xf9, 0x8b,
	0x54, 0x1f, 0x24, 0x35, 0x4b, 0x13, 0x05, 0xb4, 0x1d, 0x71, 0x11, 0x66, 0x38, 0x91, 0xaa, 0xef,
	0x22, 0xf2, 0xb3, 0x25, 0xce, 0x10, 0x38, 0xc4, 0x3c, 0xd9, 0x83, 0x8d, 0x94, 0x5f, 0xf1, 0x40,
	0x0f, 0xeb, 0xc3, 0xbe, 0x64, 0xf1, 0x48, 0x60, 0x0d, 0x05, 0xd6, 0x2c, 0x54, 0x28, 0x1c, 0x58,
	0xa4, 0x90, 0x78, 0x0a, 0x04, 0x9b, 0x45, 0x59, 0x18, 0xa6, 0x3c, 0xcb, 0xec, 0x53, 0x2a, 0x78,
	0x53, 0x07, 0x33, 0x7b, 0x36, 0x81, 0x0f, 0xba, 0x02, 0x57, 0xa8, 0x80, 0x89, 0x21, 0x1d, 0x25,
	0xc3, 0xf5, 0x99, 0xc6, 0xaf, 0xe3, 0xab, 0x4f, 0xdc, 0x86, 0x86, 0x97, 0x14, 0xeb, 0xe5, 0xaf,
	0xa0, 0x64, 0xf1, 0x33, 0xc3, 0x30, 0x59, 0x87, 0x9a, 0x50, 0x1d, 0x2a, 0x78, 0x8f, 0x0b, 0xb7,
	0x8a, 0x17, 0xaa, 0x0a, 0xd5, 0x39, 0x36, 0x67, 0xb3, 0xc2, 0x81, 0x8a, 0x13, 0x25, 0xcd, 0x17,
	0x35, 0xc2, 0x6a, 0x76, 0x85, 0x87, 0xa9, 0xe3, 0x01, 0xbf, 0x01, 0x60, 0xa8, 0xb6, 0x4a, 0x63,
	0xa6, 0x5d, 0x40, 0xcc, 0xc8, 0x1f, 0x62, 0xc0, 0x74, 0x61, 0x94, 0xa6, 0x3c, 0x0b, 0x58, 0xc2,
	0x43, 0x77, 0x1d, 0xbb, 0xe7, 0x0c, 0xb1, 0x96, 0x8d, 0x93, 0x55, 0x30, 0x17, 0xb1, 0x9d, 0x9a,
	0xb1, 0x5e, 0x21, 0x94, 0x9d, 0xf8, 0x13, 0x58, 0x30, 0x2e, 0x18, 0x05, 0x9c, 0x06, 0x22, 0xcf,
	0x34, 0x4f, 0xdd, 0x39, 0x24, 0xe6, 0x8b, 0xf0, 0xbe, 0x8d, 0x92, 0x4d, 0x98, 0x1d, 0x80, 0x52,
	0x85, 0xdc, 0x9d, 0x47, 0x6a, 0xa6, 0x88, 0x9d, 0xa8, 0x90, 0x8f, 0x23, 0xbf, 0x2a, 0xc9, 0xdd,
	0x85, 0x1b, 0xc8, 0x7b, 0x25, 0x39, 0xf1, 0x60, 0xa9, 0x1d, 0x09, 0x4e, 0xdb, 0x22, 0xcf, 0xba,
	0x34, 0x92, 0x9a, 0xa7, 0x3d, 0x26, 0x5c, 0xe7, 0x3e, 0x4b, 0x58, 0x34, 0x55, 0x87, 0xa6, 0xc8,
	0x2b, 0x6a, 0xc8, 0x37, 0x00, 0x61, 0xca, 0x22, 0x49, 0x8d, 0x57, 0xb8, 0x8b, 0xf7, 0x29, 0xd4,
	0x10, 0x3e, 0x8f, 0x62, 0xb4, 0xb8, 0x84, 0xa1, 0xb5, 0x65, 0xdd, 0x5c, 0x87, 0x66, 0x0d, 0x51,
	0x83, 0xdc, 0x6b, 0x71, 0xb6, 0xec, 0xac, 0xa8, 0x42, 0xb1, 0x1f, 0xa0, 0x1c, 0x9b, 0x7e, 0x2c,
	0xe1, 0x36, 0x3d, 0xfd, 0xd4, 0x6d, 0x7a, 0xa3, 0x42, 0xee, 0x63, 0xa5, 0x59, 0x8d, 0x30, 0xca,
	0xd8, 0xa5, 0xe0, 0x74, 0xcc, 0xe5, 0xdc, 0x87, 0x38, 0xcc, 0xc5, 0x22, 0xf5, 0xe3, 0xd0, 0xe4,
	0xc8, 0x97, 0xb0, 0xcc, 0x25, 0xe2, 0x71, 0xae, 0xf9, 0x35, 0xd5, 0x29, 0x0b, 0x22, 0xd9, 0x71,
	0x1f, 0x61, 0x01, 0xb1, 0xb9, 0x37, 0x26, 0x75, 0x6e, 0x33, 0x64, 0x1b, 0xe6, 0x06, 0xde, 0x69,
	0xcd, 0xdc, 0x45, 0x6f, 0x98, 0x2d, 0x82, 0xd6, 0xab, 0x77, 0x60, 0x3e, 0x48, 0xf2, 0x8c, 0x6b,
	0xaa, 0xbb, 0x29, 0x67, 0x61, 0xe6, 0xae, 0xa2, 0xe0, 0x9c, 0x8d, 0x9e, 0xdb, 0x20, 0x69, 0x0e,
	0x6f, 0x1b, 0x52, 0x7e, 0xad, 0xb9, 0xcc, 0xd0, 0x5a, 0x3f, 0xab, 0x4f, 0xee, 0xd6, 0x7c, 0x32,
	0x48, 0xb5, 0x86, 0x99, 0xad, 0x75, 0xa8, 0x8d, 0xbe, 0x91, 0x0a, 0x4c, 0xf4, 0x5e, 0x38, 0x0f,
	0xf0, 0xef, 0xd7, 0x4e, 0x69, 0xeb, 0x0b, 0x28, 0x9b, 0x4e, 0x90, 0x1a, 0x4c, 0xa1, 0xf3, 0x59,
	0x3b, 0xbd, 0x60, 0x22, 0x0a, 0x99, 0xe6, 0x4e, 0xc9, 0x9c, 0x3c, 0x19, 0xe9, 0xb7, 0x52, 0xf4,
	0x9d, 0x89, 0x97, 0xbb, 0xc6, 0x34, 0xb7, 0x61, 0xf3, 0x23, 0xd3, 0xbc, 0xdd, 0xdf, 0xa3, 0x72,
	0x75, 0xd6, 0x99, 0x3b, 0x2a, 0x57, 0x97, 0x9d, 0x95, 0xa3, 0x72, 0x75, 0xc5, 0x79, 0xe8, 0xd7,
	0x62, 0x76, 0x8d, 0x76, 0x95, 0xf9, 0x8e, 0xf9, 0x57, 0x5d, 0x5e, 0x51, 0xc9, 0x62, 0x4e, 0x05,
	0x97, 0xaf, 0xbe, 0xff, 0xf3, 0xb7, 0xbf, 0xff, 0xa9, 0x4c, 0x38, 0x93, 0xb0, 0x19, 0x29, 0x3b,
	0xc3, 0x24, 0x55, 0xd7, 0xfd, 0xbb, 0xc6, 0xf9, 0x6a, 0x61, 0x64, 0xd2, 0xa7, 0x66, 0x47, 0x4e,
	0x4b, 0x97, 0x15, 0x5c, 0x96, 0xe7, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x09, 0x90, 0x05,
	0x67, 0x08, 0x00, 0x00,
}