// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flare_rpc_meta.proto

package protocol

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RpcMeta struct {
	Request            *RpcRequestMeta  `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Response           *RpcResponseMeta `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	CompressType       int32            `protobuf:"varint,3,opt,name=compress_type,json=compressType" json:"compress_type"`
	CorrelationId      int64            `protobuf:"varint,4,opt,name=correlation_id,json=correlationId" json:"correlation_id"`
	AttachmentSize     int32            `protobuf:"varint,5,opt,name=attachment_size,json=attachmentSize" json:"attachment_size"`
	ChunkInfo          *ChunkInfo       `protobuf:"bytes,6,opt,name=chunk_info,json=chunkInfo" json:"chunk_info,omitempty"`
	AuthenticationData []byte           `protobuf:"bytes,7,opt,name=authentication_data,json=authenticationData" json:"authentication_data"`
	StreamSettings     *StreamSettings  `protobuf:"bytes,8,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
}

func (m *RpcMeta) Reset()         { *m = RpcMeta{} }
func (m *RpcMeta) String() string { return proto.CompactTextString(m) }
func (*RpcMeta) ProtoMessage()    {}
func (*RpcMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d45d6fa31ca5a, []int{0}
}
func (m *RpcMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMeta.Merge(m, src)
}
func (m *RpcMeta) XXX_Size() int {
	return m.Size()
}
func (m *RpcMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMeta proto.InternalMessageInfo

func (m *RpcMeta) GetRequest() *RpcRequestMeta {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *RpcMeta) GetResponse() *RpcResponseMeta {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RpcMeta) GetCompressType() int32 {
	if m != nil {
		return m.CompressType
	}
	return 0
}

func (m *RpcMeta) GetCorrelationId() int64 {
	if m != nil {
		return m.CorrelationId
	}
	return 0
}

func (m *RpcMeta) GetAttachmentSize() int32 {
	if m != nil {
		return m.AttachmentSize
	}
	return 0
}

func (m *RpcMeta) GetChunkInfo() *ChunkInfo {
	if m != nil {
		return m.ChunkInfo
	}
	return nil
}

func (m *RpcMeta) GetAuthenticationData() []byte {
	if m != nil {
		return m.AuthenticationData
	}
	return nil
}

func (m *RpcMeta) GetStreamSettings() *StreamSettings {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

type RpcRequestMeta struct {
	ServiceName  string `protobuf:"bytes,1,req,name=service_name,json=serviceName" json:"service_name"`
	MethodName   string `protobuf:"bytes,2,req,name=method_name,json=methodName" json:"method_name"`
	LogId        int64  `protobuf:"varint,3,opt,name=log_id,json=logId" json:"log_id"`
	TraceId      int64  `protobuf:"varint,4,opt,name=trace_id,json=traceId" json:"trace_id"`
	SpanId       int64  `protobuf:"varint,5,opt,name=span_id,json=spanId" json:"span_id"`
	ParentSpanId int64  `protobuf:"varint,6,opt,name=parent_span_id,json=parentSpanId" json:"parent_span_id"`
	RequestId    string `protobuf:"bytes,7,opt,name=request_id,json=requestId" json:"request_id"`
}

func (m *RpcRequestMeta) Reset()         { *m = RpcRequestMeta{} }
func (m *RpcRequestMeta) String() string { return proto.CompactTextString(m) }
func (*RpcRequestMeta) ProtoMessage()    {}
func (*RpcRequestMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d45d6fa31ca5a, []int{1}
}
func (m *RpcRequestMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcRequestMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcRequestMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcRequestMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcRequestMeta.Merge(m, src)
}
func (m *RpcRequestMeta) XXX_Size() int {
	return m.Size()
}
func (m *RpcRequestMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcRequestMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RpcRequestMeta proto.InternalMessageInfo

func (m *RpcRequestMeta) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RpcRequestMeta) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *RpcRequestMeta) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *RpcRequestMeta) GetTraceId() int64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *RpcRequestMeta) GetSpanId() int64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *RpcRequestMeta) GetParentSpanId() int64 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *RpcRequestMeta) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type RpcResponseMeta struct {
	ErrorCode int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode" json:"error_code"`
	ErrorText string `protobuf:"bytes,2,opt,name=error_text,json=errorText" json:"error_text"`
}

func (m *RpcResponseMeta) Reset()         { *m = RpcResponseMeta{} }
func (m *RpcResponseMeta) String() string { return proto.CompactTextString(m) }
func (*RpcResponseMeta) ProtoMessage()    {}
func (*RpcResponseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85d45d6fa31ca5a, []int{2}
}
func (m *RpcResponseMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcResponseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcResponseMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcResponseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcResponseMeta.Merge(m, src)
}
func (m *RpcResponseMeta) XXX_Size() int {
	return m.Size()
}
func (m *RpcResponseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcResponseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RpcResponseMeta proto.InternalMessageInfo

func (m *RpcResponseMeta) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *RpcResponseMeta) GetErrorText() string {
	if m != nil {
		return m.ErrorText
	}
	return ""
}

func init() {
	proto.RegisterType((*RpcMeta)(nil), "flare.rpc.policy.RpcMeta")
	proto.RegisterType((*RpcRequestMeta)(nil), "flare.rpc.policy.RpcRequestMeta")
	proto.RegisterType((*RpcResponseMeta)(nil), "flare.rpc.policy.RpcResponseMeta")
}

func init() { proto.RegisterFile("flare_rpc_meta.proto", fileDescriptor_a85d45d6fa31ca5a) }

var fileDescriptor_a85d45d6fa31ca5a = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0xa4, 0x49, 0x9a, 0x69, 0x92, 0xa2, 0xa5, 0x07, 0x53, 0x84, 0x6b, 0x8a, 0x2a,
	0x02, 0xa8, 0xae, 0x04, 0xe2, 0x82, 0xc4, 0x25, 0x45, 0x08, 0x1f, 0x40, 0xc8, 0xe9, 0x09, 0x0e,
	0xd6, 0xb2, 0x9e, 0x38, 0x16, 0xb6, 0x77, 0xd9, 0xdd, 0xa0, 0xb6, 0x0f, 0x81, 0x78, 0x1f, 0x5e,
	0xa0, 0xc7, 0x1e, 0x39, 0x21, 0x94, 0xbc, 0x08, 0xf2, 0xda, 0xa1, 0x4e, 0x10, 0x37, 0x7b, 0xbe,
	0xdf, 0xb7, 0x7f, 0xe6, 0x9b, 0x85, 0xbd, 0x69, 0x4a, 0x25, 0x86, 0x52, 0xb0, 0x30, 0x43, 0x4d,
	0x3d, 0x21, 0xb9, 0xe6, 0xe4, 0x96, 0xa9, 0x7a, 0x52, 0x30, 0x4f, 0xf0, 0x34, 0x61, 0x17, 0xfb,
	0x03, 0x2e, 0x74, 0xc2, 0x73, 0x55, 0x02, 0xfb, 0xb6, 0xd2, 0x12, 0x69, 0x96, 0xe4, 0xf1, 0x86,
	0xf5, 0xf0, 0x47, 0x0b, 0xba, 0x81, 0x60, 0x6f, 0x51, 0x53, 0xf2, 0x02, 0xba, 0x12, 0xbf, 0xcc,
	0x51, 0x69, 0xdb, 0x72, 0xad, 0xd1, 0xce, 0x53, 0xd7, 0xdb, 0x5c, 0xd8, 0x0b, 0x04, 0x0b, 0x4a,
	0xa6, 0xb0, 0x04, 0x2b, 0x03, 0x79, 0x09, 0xdb, 0x12, 0x95, 0xe0, 0xb9, 0x42, 0xbb, 0x69, 0xcc,
	0xf7, 0xff, 0x63, 0x2e, 0x21, 0xe3, 0xfe, 0x6b, 0x21, 0x8f, 0x60, 0xc0, 0x78, 0x26, 0x24, 0x2a,
	0x15, 0xea, 0x0b, 0x81, 0x76, 0xcb, 0xb5, 0x46, 0xed, 0xf1, 0xd6, 0xd5, 0xaf, 0x83, 0x46, 0xd0,
	0x5f, 0x49, 0x67, 0x17, 0x02, 0xc9, 0x13, 0x18, 0x32, 0x2e, 0x25, 0xa6, 0xb4, 0xb8, 0x61, 0x98,
	0x44, 0xf6, 0x96, 0x6b, 0x8d, 0x5a, 0x15, 0x3b, 0xa8, 0x69, 0x7e, 0x44, 0x8e, 0x61, 0x97, 0x6a,
	0x4d, 0xd9, 0x2c, 0xc3, 0x5c, 0x87, 0x2a, 0xb9, 0x44, 0xbb, 0x5d, 0x5b, 0x79, 0x78, 0x23, 0x4e,
	0x92, 0x4b, 0x24, 0xcf, 0x00, 0xd8, 0x6c, 0x9e, 0x7f, 0x0e, 0x93, 0x7c, 0xca, 0xed, 0x8e, 0xb9,
	0xc7, 0x5e, 0xed, 0x1e, 0xa7, 0x85, 0xe8, 0xe7, 0x53, 0x1e, 0xf4, 0xd8, 0xea, 0x93, 0x3c, 0x87,
	0xdb, 0x74, 0xae, 0x67, 0x98, 0xeb, 0x84, 0x95, 0x67, 0x8a, 0xa8, 0xa6, 0x76, 0xd7, 0xb5, 0x46,
	0xfd, 0x6a, 0x1f, 0xb2, 0x0e, 0xbc, 0xa2, 0x9a, 0x92, 0x31, 0xec, 0x96, 0xa9, 0x84, 0x0a, 0xb5,
	0x4e, 0xf2, 0x58, 0xd9, 0xdb, 0x66, 0xc3, 0x3b, 0xb5, 0x0d, 0x27, 0x86, 0x98, 0x54, 0x40, 0x30,
	0x54, 0x6b, 0xff, 0x87, 0xdf, 0x9a, 0x30, 0x5c, 0x4f, 0x84, 0x3c, 0x84, 0xbe, 0x42, 0xf9, 0x35,
	0x61, 0x18, 0xe6, 0x34, 0x43, 0xdb, 0x72, 0x9b, 0xa3, 0x5e, 0x75, 0x8c, 0x9d, 0x4a, 0x79, 0x47,
	0x33, 0x24, 0x47, 0xb0, 0x93, 0xa1, 0x9e, 0xf1, 0xa8, 0xe4, 0x9a, 0x35, 0x0e, 0x4a, 0xc1, 0x60,
	0x77, 0xa1, 0x93, 0xf2, 0xb8, 0x68, 0x73, 0xab, 0xd6, 0xe6, 0x76, 0xca, 0x63, 0x3f, 0x22, 0x07,
	0xb0, 0xad, 0x25, 0x65, 0xb8, 0x99, 0x42, 0xd7, 0x54, 0xfd, 0x88, 0xdc, 0x83, 0xae, 0x12, 0xd4,
	0xa4, 0xd4, 0xae, 0xe9, 0x9d, 0xa2, 0xe8, 0x47, 0xe4, 0x31, 0x0c, 0x05, 0x95, 0x26, 0x9a, 0x8a,
	0xea, 0xd4, 0xa8, 0x7e, 0xa9, 0x4d, 0x4a, 0xf6, 0x01, 0x40, 0x35, 0x6c, 0x05, 0x57, 0x74, 0x77,
	0x75, 0xdc, 0x5e, 0x55, 0xf7, 0xa3, 0xc3, 0x8f, 0xb0, 0xbb, 0x31, 0x64, 0x85, 0x0f, 0xa5, 0xe4,
	0x32, 0x64, 0x3c, 0x42, 0x33, 0xd8, 0xab, 0xf4, 0x7b, 0xa6, 0x7e, 0xca, 0x23, 0xbc, 0x81, 0x34,
	0x9e, 0x6b, 0x33, 0xc0, 0xbd, 0x35, 0xe8, 0x0c, 0xcf, 0xf5, 0x78, 0x7a, 0xb5, 0x70, 0xac, 0xeb,
	0x85, 0x63, 0xfd, 0x5e, 0x38, 0xd6, 0xf7, 0xa5, 0xd3, 0xb8, 0x5e, 0x3a, 0x8d, 0x9f, 0x4b, 0xa7,
	0x01, 0x7b, 0x8c, 0x67, 0xff, 0x8c, 0xfb, 0x78, 0xf0, 0xba, 0xa8, 0x04, 0x82, 0xbd, 0x2f, 0x9e,
	0xda, 0x1b, 0xeb, 0xc3, 0x51, 0x9c, 0xe8, 0xd9, 0xfc, 0x93, 0xc7, 0x78, 0x76, 0x62, 0xe8, 0x63,
	0x29, 0x58, 0xf9, 0x15, 0xf3, 0x13, 0xf3, 0x20, 0x19, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0x46, 0xfa, 0x62, 0xe5, 0x03, 0x00, 0x00,
}

func (m *RpcMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StreamSettings != nil {
		{
			size, err := m.StreamSettings.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlareRpcMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.AuthenticationData != nil {
		i -= len(m.AuthenticationData)
		copy(dAtA[i:], m.AuthenticationData)
		i = encodeVarintFlareRpcMeta(dAtA, i, uint64(len(m.AuthenticationData)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ChunkInfo != nil {
		{
			size, err := m.ChunkInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlareRpcMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.AttachmentSize))
	i--
	dAtA[i] = 0x28
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.CorrelationId))
	i--
	dAtA[i] = 0x20
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.CompressType))
	i--
	dAtA[i] = 0x18
	if m.Response != nil {
		{
			size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlareRpcMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlareRpcMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RpcRequestMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcRequestMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcRequestMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.RequestId)
	copy(dAtA[i:], m.RequestId)
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(len(m.RequestId)))
	i--
	dAtA[i] = 0x3a
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.ParentSpanId))
	i--
	dAtA[i] = 0x30
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.SpanId))
	i--
	dAtA[i] = 0x28
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.TraceId))
	i--
	dAtA[i] = 0x20
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.LogId))
	i--
	dAtA[i] = 0x18
	i -= len(m.MethodName)
	copy(dAtA[i:], m.MethodName)
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(len(m.MethodName)))
	i--
	dAtA[i] = 0x12
	i -= len(m.ServiceName)
	copy(dAtA[i:], m.ServiceName)
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(len(m.ServiceName)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RpcResponseMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcResponseMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcResponseMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.ErrorText)
	copy(dAtA[i:], m.ErrorText)
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(len(m.ErrorText)))
	i--
	dAtA[i] = 0x12
	i = encodeVarintFlareRpcMeta(dAtA, i, uint64(m.ErrorCode))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintFlareRpcMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovFlareRpcMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RpcMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovFlareRpcMeta(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovFlareRpcMeta(uint64(l))
	}
	n += 1 + sovFlareRpcMeta(uint64(m.CompressType))
	n += 1 + sovFlareRpcMeta(uint64(m.CorrelationId))
	n += 1 + sovFlareRpcMeta(uint64(m.AttachmentSize))
	if m.ChunkInfo != nil {
		l = m.ChunkInfo.Size()
		n += 1 + l + sovFlareRpcMeta(uint64(l))
	}
	if m.AuthenticationData != nil {
		l = len(m.AuthenticationData)
		n += 1 + l + sovFlareRpcMeta(uint64(l))
	}
	if m.StreamSettings != nil {
		l = m.StreamSettings.Size()
		n += 1 + l + sovFlareRpcMeta(uint64(l))
	}
	return n
}

func (m *RpcRequestMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	n += 1 + l + sovFlareRpcMeta(uint64(l))
	l = len(m.MethodName)
	n += 1 + l + sovFlareRpcMeta(uint64(l))
	n += 1 + sovFlareRpcMeta(uint64(m.LogId))
	n += 1 + sovFlareRpcMeta(uint64(m.TraceId))
	n += 1 + sovFlareRpcMeta(uint64(m.SpanId))
	n += 1 + sovFlareRpcMeta(uint64(m.ParentSpanId))
	l = len(m.RequestId)
	n += 1 + l + sovFlareRpcMeta(uint64(l))
	return n
}

func (m *RpcResponseMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFlareRpcMeta(uint64(m.ErrorCode))
	l = len(m.ErrorText)
	n += 1 + l + sovFlareRpcMeta(uint64(l))
	return n
}

func sovFlareRpcMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFlareRpcMeta(x uint64) (n int) {
	return sovFlareRpcMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RpcMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlareRpcMeta
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
			return fmt.Errorf("proto: RpcMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &RpcRequestMeta{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &RpcResponseMeta{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressType", wireType)
			}
			m.CompressType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationId", wireType)
			}
			m.CorrelationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CorrelationId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentSize", wireType)
			}
			m.AttachmentSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AttachmentSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChunkInfo == nil {
				m.ChunkInfo = &ChunkInfo{}
			}
			if err := m.ChunkInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticationData = append(m.AuthenticationData[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthenticationData == nil {
				m.AuthenticationData = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StreamSettings == nil {
				m.StreamSettings = &StreamSettings{}
			}
			if err := m.StreamSettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlareRpcMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlareRpcMeta
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
func (m *RpcRequestMeta) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlareRpcMeta
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
			return fmt.Errorf("proto: RpcRequestMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcRequestMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogId", wireType)
			}
			m.LogId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			m.TraceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TraceId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpanId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentSpanId", wireType)
			}
			m.ParentSpanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParentSpanId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlareRpcMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("service_name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("method_name")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RpcResponseMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlareRpcMeta
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
			return fmt.Errorf("proto: RpcResponseMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcResponseMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlareRpcMeta
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
				return ErrInvalidLengthFlareRpcMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlareRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlareRpcMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlareRpcMeta
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
func skipFlareRpcMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlareRpcMeta
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
					return 0, ErrIntOverflowFlareRpcMeta
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
					return 0, ErrIntOverflowFlareRpcMeta
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
				return 0, ErrInvalidLengthFlareRpcMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFlareRpcMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFlareRpcMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFlareRpcMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlareRpcMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFlareRpcMeta = fmt.Errorf("proto: unexpected end of group")
)