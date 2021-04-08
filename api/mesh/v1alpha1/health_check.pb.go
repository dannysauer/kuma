// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: mesh/v1alpha1/health_check.proto

package v1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// HealthCheck defines configuration for health checking.
type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of selectors to match dataplanes that should be configured to do
	// health checks.
	Sources []*Selector `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	// List of selectors to match services that need to be health checked.
	Destinations []*Selector `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// Configuration for various types of health checking.
	Conf *HealthCheck_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_health_check_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheck) GetSources() []*Selector {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *HealthCheck) GetDestinations() []*Selector {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *HealthCheck) GetConf() *HealthCheck_Conf {
	if x != nil {
		return x.Conf
	}
	return nil
}

// Conf defines configuration for various types of health checking.
type HealthCheck_Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interval between consecutive health checks.
	Interval *duration.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// Maximum time to wait for a health check response.
	Timeout *duration.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of consecutive unhealthy checks before considering a host
	// unhealthy.
	UnhealthyThreshold uint32 `protobuf:"varint,3,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	// Number of consecutive healthy checks before considering a host healthy.
	HealthyThreshold uint32 `protobuf:"varint,4,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	// If specified, Envoy will start health checking after for a random time in
	// ms between 0 and initial_jitter. This only applies to the first health
	// check.
	InitialJitter *duration.Duration `protobuf:"bytes,7,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	// If specified, during every interval Envoy will add interval_jitter to the
	// wait time.
	IntervalJitter *duration.Duration `protobuf:"bytes,8,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	// If specified, during every interval Envoy will add interval_ms *
	// interval_jitter_percent / 100 to the wait time. If interval_jitter_ms and
	// interval_jitter_percent are both set, both of them will be used to
	// increase the wait time.
	IntervalJitterPercent uint32 `protobuf:"varint,9,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	// Allows to configure panic threshold for Envoy cluster. If not specified,
	// the default is 50%. To disable panic mode, set to 0%.
	HealthyPanicThreshold *wrappers.FloatValue `protobuf:"bytes,10,opt,name=healthy_panic_threshold,json=healthyPanicThreshold,proto3" json:"healthy_panic_threshold,omitempty"`
	// If set to true, Envoy will not consider any hosts when the cluster is in
	// 'panic mode'. Instead, the cluster will fail all requests as if all hosts
	// are unhealthy. This can help avoid potentially overwhelming a failing
	// service.
	FailTrafficOnPanic *wrappers.BoolValue `protobuf:"bytes,11,opt,name=fail_traffic_on_panic,json=failTrafficOnPanic,proto3" json:"fail_traffic_on_panic,omitempty"`
	// Specifies the path to the file where Envoy can log health check events.
	// If empty, no event log will be written.
	EventLogPath string `protobuf:"bytes,12,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	// If set to true, health check failure events will always be logged. If set
	// to false, only the initial health check failure event will be logged. The
	// default value is false.
	AlwaysLogHealthCheckFailures *wrappers.BoolValue `protobuf:"bytes,13,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	// The "no traffic interval" is a special health check interval that is used
	// when a cluster has never had traffic routed to it. This lower interval
	// allows cluster information to be kept up to date, without sending a
	// potentially large amount of active health checking traffic for no reason.
	// Once a cluster has been used for traffic routing, Envoy will shift back
	// to using the standard health check interval that is defined. Note that
	// this interval takes precedence over any other. The default value for "no
	// traffic interval" is 60 seconds.
	NoTrafficInterval *duration.Duration `protobuf:"bytes,14,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	// Types that are assignable to Protocol:
	//	*HealthCheck_Conf_Tcp_
	//	*HealthCheck_Conf_Http_
	Protocol isHealthCheck_Conf_Protocol `protobuf_oneof:"protocol"`
}

func (x *HealthCheck_Conf) Reset() {
	*x = HealthCheck_Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck_Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck_Conf) ProtoMessage() {}

func (x *HealthCheck_Conf) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck_Conf.ProtoReflect.Descriptor instead.
func (*HealthCheck_Conf) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_health_check_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HealthCheck_Conf) GetInterval() *duration.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *HealthCheck_Conf) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *HealthCheck_Conf) GetUnhealthyThreshold() uint32 {
	if x != nil {
		return x.UnhealthyThreshold
	}
	return 0
}

func (x *HealthCheck_Conf) GetHealthyThreshold() uint32 {
	if x != nil {
		return x.HealthyThreshold
	}
	return 0
}

func (x *HealthCheck_Conf) GetInitialJitter() *duration.Duration {
	if x != nil {
		return x.InitialJitter
	}
	return nil
}

func (x *HealthCheck_Conf) GetIntervalJitter() *duration.Duration {
	if x != nil {
		return x.IntervalJitter
	}
	return nil
}

func (x *HealthCheck_Conf) GetIntervalJitterPercent() uint32 {
	if x != nil {
		return x.IntervalJitterPercent
	}
	return 0
}

func (x *HealthCheck_Conf) GetHealthyPanicThreshold() *wrappers.FloatValue {
	if x != nil {
		return x.HealthyPanicThreshold
	}
	return nil
}

func (x *HealthCheck_Conf) GetFailTrafficOnPanic() *wrappers.BoolValue {
	if x != nil {
		return x.FailTrafficOnPanic
	}
	return nil
}

func (x *HealthCheck_Conf) GetEventLogPath() string {
	if x != nil {
		return x.EventLogPath
	}
	return ""
}

func (x *HealthCheck_Conf) GetAlwaysLogHealthCheckFailures() *wrappers.BoolValue {
	if x != nil {
		return x.AlwaysLogHealthCheckFailures
	}
	return nil
}

func (x *HealthCheck_Conf) GetNoTrafficInterval() *duration.Duration {
	if x != nil {
		return x.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck_Conf) GetProtocol() isHealthCheck_Conf_Protocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (x *HealthCheck_Conf) GetTcp() *HealthCheck_Conf_Tcp {
	if x, ok := x.GetProtocol().(*HealthCheck_Conf_Tcp_); ok {
		return x.Tcp
	}
	return nil
}

func (x *HealthCheck_Conf) GetHttp() *HealthCheck_Conf_Http {
	if x, ok := x.GetProtocol().(*HealthCheck_Conf_Http_); ok {
		return x.Http
	}
	return nil
}

type isHealthCheck_Conf_Protocol interface {
	isHealthCheck_Conf_Protocol()
}

type HealthCheck_Conf_Tcp_ struct {
	Tcp *HealthCheck_Conf_Tcp `protobuf:"bytes,5,opt,name=tcp,proto3,oneof"`
}

type HealthCheck_Conf_Http_ struct {
	Http *HealthCheck_Conf_Http `protobuf:"bytes,6,opt,name=http,proto3,oneof"`
}

func (*HealthCheck_Conf_Tcp_) isHealthCheck_Conf_Protocol() {}

func (*HealthCheck_Conf_Http_) isHealthCheck_Conf_Protocol() {}

// Tcp defines optional configuration for specifying bytes to send and
// expected response during the health check
type HealthCheck_Conf_Tcp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bytes which will be send during the health check to the target
	Send *wrappers.BytesValue `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	// Bytes blocks expected as a response. When checking the response,
	// “fuzzy” matching is performed such that each block must be found, and
	// in the order specified, but not necessarily contiguous.
	Receive []*wrappers.BytesValue `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
}

func (x *HealthCheck_Conf_Tcp) Reset() {
	*x = HealthCheck_Conf_Tcp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck_Conf_Tcp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck_Conf_Tcp) ProtoMessage() {}

func (x *HealthCheck_Conf_Tcp) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck_Conf_Tcp.ProtoReflect.Descriptor instead.
func (*HealthCheck_Conf_Tcp) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_health_check_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *HealthCheck_Conf_Tcp) GetSend() *wrappers.BytesValue {
	if x != nil {
		return x.Send
	}
	return nil
}

func (x *HealthCheck_Conf_Tcp) GetReceive() []*wrappers.BytesValue {
	if x != nil {
		return x.Receive
	}
	return nil
}

// Http defines optional Http configuration which will instruct the service
// the health check will be made for is an http service. It's mutually
// exclusive with the Tcp block so when provided you can't provide
// the Tcp configuration
type HealthCheck_Conf_Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP path which will be requested during the health check
	// (ie. /health)
	//  +required
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` //
	// The list of HTTP headers which should be added to each health check
	// request
	//  +optional
	RequestHeadersToAdd []*HealthCheck_Conf_Http_HeaderValueOption `protobuf:"bytes,2,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// List of HTTP response statuses which are considered healthy
	//  +optional
	ExpectedStatuses []*wrappers.UInt32Value `protobuf:"bytes,3,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
}

func (x *HealthCheck_Conf_Http) Reset() {
	*x = HealthCheck_Conf_Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck_Conf_Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck_Conf_Http) ProtoMessage() {}

func (x *HealthCheck_Conf_Http) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck_Conf_Http.ProtoReflect.Descriptor instead.
func (*HealthCheck_Conf_Http) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_health_check_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *HealthCheck_Conf_Http) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HealthCheck_Conf_Http) GetRequestHeadersToAdd() []*HealthCheck_Conf_Http_HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *HealthCheck_Conf_Http) GetExpectedStatuses() []*wrappers.UInt32Value {
	if x != nil {
		return x.ExpectedStatuses
	}
	return nil
}

type HealthCheck_Conf_Http_HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name
	//  +required
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value
	//  +optional
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HealthCheck_Conf_Http_HeaderValue) Reset() {
	*x = HealthCheck_Conf_Http_HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck_Conf_Http_HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck_Conf_Http_HeaderValue) ProtoMessage() {}

func (x *HealthCheck_Conf_Http_HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck_Conf_Http_HeaderValue.ProtoReflect.Descriptor instead.
func (*HealthCheck_Conf_Http_HeaderValue) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_health_check_proto_rawDescGZIP(), []int{0, 0, 1, 0}
}

func (x *HealthCheck_Conf_Http_HeaderValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HealthCheck_Conf_Http_HeaderValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type HealthCheck_Conf_Http_HeaderValueOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key/Value representation of the HTTP header
	//  +required
	Header *HealthCheck_Conf_Http_HeaderValue `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The bool value which if true (default) will mean the header values
	// should be appended to already present ones
	//  +optional
	Append *wrappers.BoolValue `protobuf:"bytes,2,opt,name=append,proto3" json:"append,omitempty"`
}

func (x *HealthCheck_Conf_Http_HeaderValueOption) Reset() {
	*x = HealthCheck_Conf_Http_HeaderValueOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck_Conf_Http_HeaderValueOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck_Conf_Http_HeaderValueOption) ProtoMessage() {}

func (x *HealthCheck_Conf_Http_HeaderValueOption) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_health_check_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck_Conf_Http_HeaderValueOption.ProtoReflect.Descriptor instead.
func (*HealthCheck_Conf_Http_HeaderValueOption) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_health_check_proto_rawDescGZIP(), []int{0, 0, 1, 1}
}

func (x *HealthCheck_Conf_Http_HeaderValueOption) GetHeader() *HealthCheck_Conf_Http_HeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HealthCheck_Conf_Http_HeaderValueOption) GetAppend() *wrappers.BoolValue {
	if x != nil {
		return x.Append
	}
	return nil
}

var File_mesh_v1alpha1_health_check_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_health_check_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x0d,
	0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x40, 0x0a,
	0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x04, 0x63,
	0x6f, 0x6e, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x52,
	0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0xe1, 0x0b, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x41,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x3f, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x38, 0x0a, 0x13, 0x75, 0x6e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x12, 0x75, 0x6e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x34, 0x0a, 0x11,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x10, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6a, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4a, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x4a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x4a, 0x69, 0x74, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x12, 0x53, 0x0a, 0x17, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f, 0x70, 0x61, 0x6e, 0x69,
	0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x4d, 0x0a, 0x15, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4f, 0x6e, 0x50,
	0x61, 0x6e, 0x69, 0x63, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x62, 0x0a, 0x20, 0x61, 0x6c,
	0x77, 0x61, 0x79, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1c, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x4c, 0x6f, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x49,
	0x0a, 0x13, 0x6e, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6e, 0x6f, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3c, 0x0a, 0x03, 0x74, 0x63, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x63, 0x70,
	0x48, 0x00, 0x52, 0x03, 0x74, 0x63, 0x70, 0x12, 0x3f, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x48, 0x00, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x6d, 0x0a, 0x03, 0x54, 0x63, 0x70, 0x12,
	0x2f, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64,
	0x12, 0x35, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x1a, 0xb3, 0x03, 0x0a, 0x04, 0x48, 0x74, 0x74, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x70, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x55, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x2a, 0x05, 0x10, 0xd8, 0x04, 0x28, 0x64, 0x52, 0x10, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x1a, 0x35, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x96, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b,
	0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_v1alpha1_health_check_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_health_check_proto_rawDescData = file_mesh_v1alpha1_health_check_proto_rawDesc
)

func file_mesh_v1alpha1_health_check_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_health_check_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_health_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_health_check_proto_rawDescData)
	})
	return file_mesh_v1alpha1_health_check_proto_rawDescData
}

var file_mesh_v1alpha1_health_check_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mesh_v1alpha1_health_check_proto_goTypes = []interface{}{
	(*HealthCheck)(nil),                             // 0: kuma.mesh.v1alpha1.HealthCheck
	(*HealthCheck_Conf)(nil),                        // 1: kuma.mesh.v1alpha1.HealthCheck.Conf
	(*HealthCheck_Conf_Tcp)(nil),                    // 2: kuma.mesh.v1alpha1.HealthCheck.Conf.Tcp
	(*HealthCheck_Conf_Http)(nil),                   // 3: kuma.mesh.v1alpha1.HealthCheck.Conf.Http
	(*HealthCheck_Conf_Http_HeaderValue)(nil),       // 4: kuma.mesh.v1alpha1.HealthCheck.Conf.Http.HeaderValue
	(*HealthCheck_Conf_Http_HeaderValueOption)(nil), // 5: kuma.mesh.v1alpha1.HealthCheck.Conf.Http.HeaderValueOption
	(*Selector)(nil),                                // 6: kuma.mesh.v1alpha1.Selector
	(*duration.Duration)(nil),                       // 7: google.protobuf.Duration
	(*wrappers.FloatValue)(nil),                     // 8: google.protobuf.FloatValue
	(*wrappers.BoolValue)(nil),                      // 9: google.protobuf.BoolValue
	(*wrappers.BytesValue)(nil),                     // 10: google.protobuf.BytesValue
	(*wrappers.UInt32Value)(nil),                    // 11: google.protobuf.UInt32Value
}
var file_mesh_v1alpha1_health_check_proto_depIdxs = []int32{
	6,  // 0: kuma.mesh.v1alpha1.HealthCheck.sources:type_name -> kuma.mesh.v1alpha1.Selector
	6,  // 1: kuma.mesh.v1alpha1.HealthCheck.destinations:type_name -> kuma.mesh.v1alpha1.Selector
	1,  // 2: kuma.mesh.v1alpha1.HealthCheck.conf:type_name -> kuma.mesh.v1alpha1.HealthCheck.Conf
	7,  // 3: kuma.mesh.v1alpha1.HealthCheck.Conf.interval:type_name -> google.protobuf.Duration
	7,  // 4: kuma.mesh.v1alpha1.HealthCheck.Conf.timeout:type_name -> google.protobuf.Duration
	7,  // 5: kuma.mesh.v1alpha1.HealthCheck.Conf.initial_jitter:type_name -> google.protobuf.Duration
	7,  // 6: kuma.mesh.v1alpha1.HealthCheck.Conf.interval_jitter:type_name -> google.protobuf.Duration
	8,  // 7: kuma.mesh.v1alpha1.HealthCheck.Conf.healthy_panic_threshold:type_name -> google.protobuf.FloatValue
	9,  // 8: kuma.mesh.v1alpha1.HealthCheck.Conf.fail_traffic_on_panic:type_name -> google.protobuf.BoolValue
	9,  // 9: kuma.mesh.v1alpha1.HealthCheck.Conf.always_log_health_check_failures:type_name -> google.protobuf.BoolValue
	7,  // 10: kuma.mesh.v1alpha1.HealthCheck.Conf.no_traffic_interval:type_name -> google.protobuf.Duration
	2,  // 11: kuma.mesh.v1alpha1.HealthCheck.Conf.tcp:type_name -> kuma.mesh.v1alpha1.HealthCheck.Conf.Tcp
	3,  // 12: kuma.mesh.v1alpha1.HealthCheck.Conf.http:type_name -> kuma.mesh.v1alpha1.HealthCheck.Conf.Http
	10, // 13: kuma.mesh.v1alpha1.HealthCheck.Conf.Tcp.send:type_name -> google.protobuf.BytesValue
	10, // 14: kuma.mesh.v1alpha1.HealthCheck.Conf.Tcp.receive:type_name -> google.protobuf.BytesValue
	5,  // 15: kuma.mesh.v1alpha1.HealthCheck.Conf.Http.request_headers_to_add:type_name -> kuma.mesh.v1alpha1.HealthCheck.Conf.Http.HeaderValueOption
	11, // 16: kuma.mesh.v1alpha1.HealthCheck.Conf.Http.expected_statuses:type_name -> google.protobuf.UInt32Value
	4,  // 17: kuma.mesh.v1alpha1.HealthCheck.Conf.Http.HeaderValueOption.header:type_name -> kuma.mesh.v1alpha1.HealthCheck.Conf.Http.HeaderValue
	9,  // 18: kuma.mesh.v1alpha1.HealthCheck.Conf.Http.HeaderValueOption.append:type_name -> google.protobuf.BoolValue
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_health_check_proto_init() }
func file_mesh_v1alpha1_health_check_proto_init() {
	if File_mesh_v1alpha1_health_check_proto != nil {
		return
	}
	file_mesh_v1alpha1_selector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mesh_v1alpha1_health_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
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
		file_mesh_v1alpha1_health_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck_Conf); i {
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
		file_mesh_v1alpha1_health_check_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck_Conf_Tcp); i {
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
		file_mesh_v1alpha1_health_check_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck_Conf_Http); i {
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
		file_mesh_v1alpha1_health_check_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck_Conf_Http_HeaderValue); i {
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
		file_mesh_v1alpha1_health_check_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck_Conf_Http_HeaderValueOption); i {
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
	file_mesh_v1alpha1_health_check_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HealthCheck_Conf_Tcp_)(nil),
		(*HealthCheck_Conf_Http_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mesh_v1alpha1_health_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mesh_v1alpha1_health_check_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_health_check_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_health_check_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_health_check_proto = out.File
	file_mesh_v1alpha1_health_check_proto_rawDesc = nil
	file_mesh_v1alpha1_health_check_proto_goTypes = nil
	file_mesh_v1alpha1_health_check_proto_depIdxs = nil
}
