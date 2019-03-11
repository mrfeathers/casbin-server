// Code generated by protoc-gen-go. DO NOT EDIT.
// source: casbin.proto

package proto

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

type NewEnforcerRequest struct {
	ModelText            string   `protobuf:"bytes,1,opt,name=modelText,proto3" json:"modelText,omitempty"`
	AdapterHandle        int32    `protobuf:"varint,2,opt,name=adapterHandle,proto3" json:"adapterHandle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewEnforcerRequest) Reset()         { *m = NewEnforcerRequest{} }
func (m *NewEnforcerRequest) String() string { return proto.CompactTextString(m) }
func (*NewEnforcerRequest) ProtoMessage()    {}
func (*NewEnforcerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{0}
}

func (m *NewEnforcerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewEnforcerRequest.Unmarshal(m, b)
}
func (m *NewEnforcerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewEnforcerRequest.Marshal(b, m, deterministic)
}
func (m *NewEnforcerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewEnforcerRequest.Merge(m, src)
}
func (m *NewEnforcerRequest) XXX_Size() int {
	return xxx_messageInfo_NewEnforcerRequest.Size(m)
}
func (m *NewEnforcerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewEnforcerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewEnforcerRequest proto.InternalMessageInfo

func (m *NewEnforcerRequest) GetModelText() string {
	if m != nil {
		return m.ModelText
	}
	return ""
}

func (m *NewEnforcerRequest) GetAdapterHandle() int32 {
	if m != nil {
		return m.AdapterHandle
	}
	return 0
}

type NewEnforcerReply struct {
	Handler              int32    `protobuf:"varint,1,opt,name=handler,proto3" json:"handler,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewEnforcerReply) Reset()         { *m = NewEnforcerReply{} }
func (m *NewEnforcerReply) String() string { return proto.CompactTextString(m) }
func (*NewEnforcerReply) ProtoMessage()    {}
func (*NewEnforcerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{1}
}

func (m *NewEnforcerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewEnforcerReply.Unmarshal(m, b)
}
func (m *NewEnforcerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewEnforcerReply.Marshal(b, m, deterministic)
}
func (m *NewEnforcerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewEnforcerReply.Merge(m, src)
}
func (m *NewEnforcerReply) XXX_Size() int {
	return xxx_messageInfo_NewEnforcerReply.Size(m)
}
func (m *NewEnforcerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NewEnforcerReply.DiscardUnknown(m)
}

var xxx_messageInfo_NewEnforcerReply proto.InternalMessageInfo

func (m *NewEnforcerReply) GetHandler() int32 {
	if m != nil {
		return m.Handler
	}
	return 0
}

type NewAdapterRequest struct {
	AdapterName          string   `protobuf:"bytes,1,opt,name=adapterName,proto3" json:"adapterName,omitempty"`
	DriverName           string   `protobuf:"bytes,2,opt,name=driverName,proto3" json:"driverName,omitempty"`
	ConnectString        string   `protobuf:"bytes,3,opt,name=connectString,proto3" json:"connectString,omitempty"`
	DbSpecified          bool     `protobuf:"varint,4,opt,name=dbSpecified,proto3" json:"dbSpecified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAdapterRequest) Reset()         { *m = NewAdapterRequest{} }
func (m *NewAdapterRequest) String() string { return proto.CompactTextString(m) }
func (*NewAdapterRequest) ProtoMessage()    {}
func (*NewAdapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{2}
}

func (m *NewAdapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAdapterRequest.Unmarshal(m, b)
}
func (m *NewAdapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAdapterRequest.Marshal(b, m, deterministic)
}
func (m *NewAdapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAdapterRequest.Merge(m, src)
}
func (m *NewAdapterRequest) XXX_Size() int {
	return xxx_messageInfo_NewAdapterRequest.Size(m)
}
func (m *NewAdapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAdapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewAdapterRequest proto.InternalMessageInfo

func (m *NewAdapterRequest) GetAdapterName() string {
	if m != nil {
		return m.AdapterName
	}
	return ""
}

func (m *NewAdapterRequest) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *NewAdapterRequest) GetConnectString() string {
	if m != nil {
		return m.ConnectString
	}
	return ""
}

func (m *NewAdapterRequest) GetDbSpecified() bool {
	if m != nil {
		return m.DbSpecified
	}
	return false
}

type NewAdapterReply struct {
	Handler              int32    `protobuf:"varint,1,opt,name=handler,proto3" json:"handler,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAdapterReply) Reset()         { *m = NewAdapterReply{} }
func (m *NewAdapterReply) String() string { return proto.CompactTextString(m) }
func (*NewAdapterReply) ProtoMessage()    {}
func (*NewAdapterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{3}
}

func (m *NewAdapterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAdapterReply.Unmarshal(m, b)
}
func (m *NewAdapterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAdapterReply.Marshal(b, m, deterministic)
}
func (m *NewAdapterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAdapterReply.Merge(m, src)
}
func (m *NewAdapterReply) XXX_Size() int {
	return xxx_messageInfo_NewAdapterReply.Size(m)
}
func (m *NewAdapterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAdapterReply.DiscardUnknown(m)
}

var xxx_messageInfo_NewAdapterReply proto.InternalMessageInfo

func (m *NewAdapterReply) GetHandler() int32 {
	if m != nil {
		return m.Handler
	}
	return 0
}

type EnforceRequest struct {
	EnforcerHandler      int32    `protobuf:"varint,1,opt,name=enforcerHandler,proto3" json:"enforcerHandler,omitempty"`
	Params               []string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnforceRequest) Reset()         { *m = EnforceRequest{} }
func (m *EnforceRequest) String() string { return proto.CompactTextString(m) }
func (*EnforceRequest) ProtoMessage()    {}
func (*EnforceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{4}
}

func (m *EnforceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnforceRequest.Unmarshal(m, b)
}
func (m *EnforceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnforceRequest.Marshal(b, m, deterministic)
}
func (m *EnforceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnforceRequest.Merge(m, src)
}
func (m *EnforceRequest) XXX_Size() int {
	return xxx_messageInfo_EnforceRequest.Size(m)
}
func (m *EnforceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnforceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnforceRequest proto.InternalMessageInfo

func (m *EnforceRequest) GetEnforcerHandler() int32 {
	if m != nil {
		return m.EnforcerHandler
	}
	return 0
}

func (m *EnforceRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type BoolReply struct {
	Res                  bool     `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolReply) Reset()         { *m = BoolReply{} }
func (m *BoolReply) String() string { return proto.CompactTextString(m) }
func (*BoolReply) ProtoMessage()    {}
func (*BoolReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{5}
}

func (m *BoolReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolReply.Unmarshal(m, b)
}
func (m *BoolReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolReply.Marshal(b, m, deterministic)
}
func (m *BoolReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolReply.Merge(m, src)
}
func (m *BoolReply) XXX_Size() int {
	return xxx_messageInfo_BoolReply.Size(m)
}
func (m *BoolReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolReply.DiscardUnknown(m)
}

var xxx_messageInfo_BoolReply proto.InternalMessageInfo

func (m *BoolReply) GetRes() bool {
	if m != nil {
		return m.Res
	}
	return false
}

type EmptyRequest struct {
	Handler              int32    `protobuf:"varint,1,opt,name=handler,proto3" json:"handler,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{6}
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

func (m *EmptyRequest) GetHandler() int32 {
	if m != nil {
		return m.Handler
	}
	return 0
}

type EmptyReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyReply) Reset()         { *m = EmptyReply{} }
func (m *EmptyReply) String() string { return proto.CompactTextString(m) }
func (*EmptyReply) ProtoMessage()    {}
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{7}
}

func (m *EmptyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyReply.Unmarshal(m, b)
}
func (m *EmptyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyReply.Marshal(b, m, deterministic)
}
func (m *EmptyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyReply.Merge(m, src)
}
func (m *EmptyReply) XXX_Size() int {
	return xxx_messageInfo_EmptyReply.Size(m)
}
func (m *EmptyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyReply.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyReply proto.InternalMessageInfo

type PolicyRequest struct {
	EnforcerHandler      int32    `protobuf:"varint,1,opt,name=enforcerHandler,proto3" json:"enforcerHandler,omitempty"`
	PType                string   `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	Params               []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyRequest) Reset()         { *m = PolicyRequest{} }
func (m *PolicyRequest) String() string { return proto.CompactTextString(m) }
func (*PolicyRequest) ProtoMessage()    {}
func (*PolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{8}
}

func (m *PolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyRequest.Unmarshal(m, b)
}
func (m *PolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyRequest.Marshal(b, m, deterministic)
}
func (m *PolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyRequest.Merge(m, src)
}
func (m *PolicyRequest) XXX_Size() int {
	return xxx_messageInfo_PolicyRequest.Size(m)
}
func (m *PolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyRequest proto.InternalMessageInfo

func (m *PolicyRequest) GetEnforcerHandler() int32 {
	if m != nil {
		return m.EnforcerHandler
	}
	return 0
}

func (m *PolicyRequest) GetPType() string {
	if m != nil {
		return m.PType
	}
	return ""
}

func (m *PolicyRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type SimpleGetRequest struct {
	EnforcerHandler      int32    `protobuf:"varint,1,opt,name=enforcerHandler,proto3" json:"enforcerHandler,omitempty"`
	PType                string   `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleGetRequest) Reset()         { *m = SimpleGetRequest{} }
func (m *SimpleGetRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleGetRequest) ProtoMessage()    {}
func (*SimpleGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{9}
}

func (m *SimpleGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleGetRequest.Unmarshal(m, b)
}
func (m *SimpleGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleGetRequest.Marshal(b, m, deterministic)
}
func (m *SimpleGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleGetRequest.Merge(m, src)
}
func (m *SimpleGetRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleGetRequest.Size(m)
}
func (m *SimpleGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleGetRequest proto.InternalMessageInfo

func (m *SimpleGetRequest) GetEnforcerHandler() int32 {
	if m != nil {
		return m.EnforcerHandler
	}
	return 0
}

func (m *SimpleGetRequest) GetPType() string {
	if m != nil {
		return m.PType
	}
	return ""
}

type ArrayReply struct {
	Array                []string `protobuf:"bytes,1,rep,name=array,proto3" json:"array,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayReply) Reset()         { *m = ArrayReply{} }
func (m *ArrayReply) String() string { return proto.CompactTextString(m) }
func (*ArrayReply) ProtoMessage()    {}
func (*ArrayReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{10}
}

func (m *ArrayReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayReply.Unmarshal(m, b)
}
func (m *ArrayReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayReply.Marshal(b, m, deterministic)
}
func (m *ArrayReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayReply.Merge(m, src)
}
func (m *ArrayReply) XXX_Size() int {
	return xxx_messageInfo_ArrayReply.Size(m)
}
func (m *ArrayReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayReply.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayReply proto.InternalMessageInfo

func (m *ArrayReply) GetArray() []string {
	if m != nil {
		return m.Array
	}
	return nil
}

type FilteredPolicyRequest struct {
	EnforcerHandler      int32    `protobuf:"varint,1,opt,name=enforcerHandler,proto3" json:"enforcerHandler,omitempty"`
	PType                string   `protobuf:"bytes,2,opt,name=pType,proto3" json:"pType,omitempty"`
	FieldIndex           int32    `protobuf:"varint,3,opt,name=fieldIndex,proto3" json:"fieldIndex,omitempty"`
	FieldValues          []string `protobuf:"bytes,4,rep,name=fieldValues,proto3" json:"fieldValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilteredPolicyRequest) Reset()         { *m = FilteredPolicyRequest{} }
func (m *FilteredPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*FilteredPolicyRequest) ProtoMessage()    {}
func (*FilteredPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{11}
}

func (m *FilteredPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilteredPolicyRequest.Unmarshal(m, b)
}
func (m *FilteredPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilteredPolicyRequest.Marshal(b, m, deterministic)
}
func (m *FilteredPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilteredPolicyRequest.Merge(m, src)
}
func (m *FilteredPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_FilteredPolicyRequest.Size(m)
}
func (m *FilteredPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilteredPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilteredPolicyRequest proto.InternalMessageInfo

func (m *FilteredPolicyRequest) GetEnforcerHandler() int32 {
	if m != nil {
		return m.EnforcerHandler
	}
	return 0
}

func (m *FilteredPolicyRequest) GetPType() string {
	if m != nil {
		return m.PType
	}
	return ""
}

func (m *FilteredPolicyRequest) GetFieldIndex() int32 {
	if m != nil {
		return m.FieldIndex
	}
	return 0
}

func (m *FilteredPolicyRequest) GetFieldValues() []string {
	if m != nil {
		return m.FieldValues
	}
	return nil
}

type UserRoleRequest struct {
	EnforcerHandler      int32    `protobuf:"varint,1,opt,name=enforcerHandler,proto3" json:"enforcerHandler,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRoleRequest) Reset()         { *m = UserRoleRequest{} }
func (m *UserRoleRequest) String() string { return proto.CompactTextString(m) }
func (*UserRoleRequest) ProtoMessage()    {}
func (*UserRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{12}
}

func (m *UserRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRoleRequest.Unmarshal(m, b)
}
func (m *UserRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRoleRequest.Marshal(b, m, deterministic)
}
func (m *UserRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRoleRequest.Merge(m, src)
}
func (m *UserRoleRequest) XXX_Size() int {
	return xxx_messageInfo_UserRoleRequest.Size(m)
}
func (m *UserRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRoleRequest proto.InternalMessageInfo

func (m *UserRoleRequest) GetEnforcerHandler() int32 {
	if m != nil {
		return m.EnforcerHandler
	}
	return 0
}

func (m *UserRoleRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserRoleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type PermissionRequest struct {
	EnforcerHandler      int32    `protobuf:"varint,1,opt,name=enforcerHandler,proto3" json:"enforcerHandler,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Permissions          []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermissionRequest) Reset()         { *m = PermissionRequest{} }
func (m *PermissionRequest) String() string { return proto.CompactTextString(m) }
func (*PermissionRequest) ProtoMessage()    {}
func (*PermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{13}
}

func (m *PermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionRequest.Unmarshal(m, b)
}
func (m *PermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionRequest.Marshal(b, m, deterministic)
}
func (m *PermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionRequest.Merge(m, src)
}
func (m *PermissionRequest) XXX_Size() int {
	return xxx_messageInfo_PermissionRequest.Size(m)
}
func (m *PermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionRequest proto.InternalMessageInfo

func (m *PermissionRequest) GetEnforcerHandler() int32 {
	if m != nil {
		return m.EnforcerHandler
	}
	return 0
}

func (m *PermissionRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *PermissionRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Array2DReply struct {
	D2                   []*Array2DReplyD `protobuf:"bytes,1,rep,name=d2,proto3" json:"d2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Array2DReply) Reset()         { *m = Array2DReply{} }
func (m *Array2DReply) String() string { return proto.CompactTextString(m) }
func (*Array2DReply) ProtoMessage()    {}
func (*Array2DReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{14}
}

func (m *Array2DReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Array2DReply.Unmarshal(m, b)
}
func (m *Array2DReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Array2DReply.Marshal(b, m, deterministic)
}
func (m *Array2DReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Array2DReply.Merge(m, src)
}
func (m *Array2DReply) XXX_Size() int {
	return xxx_messageInfo_Array2DReply.Size(m)
}
func (m *Array2DReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Array2DReply.DiscardUnknown(m)
}

var xxx_messageInfo_Array2DReply proto.InternalMessageInfo

func (m *Array2DReply) GetD2() []*Array2DReplyD {
	if m != nil {
		return m.D2
	}
	return nil
}

type Array2DReplyD struct {
	D1                   []string `protobuf:"bytes,1,rep,name=d1,proto3" json:"d1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Array2DReplyD) Reset()         { *m = Array2DReplyD{} }
func (m *Array2DReplyD) String() string { return proto.CompactTextString(m) }
func (*Array2DReplyD) ProtoMessage()    {}
func (*Array2DReplyD) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0f85ba9164bca5, []int{14, 0}
}

func (m *Array2DReplyD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Array2DReplyD.Unmarshal(m, b)
}
func (m *Array2DReplyD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Array2DReplyD.Marshal(b, m, deterministic)
}
func (m *Array2DReplyD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Array2DReplyD.Merge(m, src)
}
func (m *Array2DReplyD) XXX_Size() int {
	return xxx_messageInfo_Array2DReplyD.Size(m)
}
func (m *Array2DReplyD) XXX_DiscardUnknown() {
	xxx_messageInfo_Array2DReplyD.DiscardUnknown(m)
}

var xxx_messageInfo_Array2DReplyD proto.InternalMessageInfo

func (m *Array2DReplyD) GetD1() []string {
	if m != nil {
		return m.D1
	}
	return nil
}

func init() {
	proto.RegisterType((*NewEnforcerRequest)(nil), "proto.NewEnforcerRequest")
	proto.RegisterType((*NewEnforcerReply)(nil), "proto.NewEnforcerReply")
	proto.RegisterType((*NewAdapterRequest)(nil), "proto.NewAdapterRequest")
	proto.RegisterType((*NewAdapterReply)(nil), "proto.NewAdapterReply")
	proto.RegisterType((*EnforceRequest)(nil), "proto.EnforceRequest")
	proto.RegisterType((*BoolReply)(nil), "proto.BoolReply")
	proto.RegisterType((*EmptyRequest)(nil), "proto.EmptyRequest")
	proto.RegisterType((*EmptyReply)(nil), "proto.EmptyReply")
	proto.RegisterType((*PolicyRequest)(nil), "proto.PolicyRequest")
	proto.RegisterType((*SimpleGetRequest)(nil), "proto.SimpleGetRequest")
	proto.RegisterType((*ArrayReply)(nil), "proto.ArrayReply")
	proto.RegisterType((*FilteredPolicyRequest)(nil), "proto.FilteredPolicyRequest")
	proto.RegisterType((*UserRoleRequest)(nil), "proto.UserRoleRequest")
	proto.RegisterType((*PermissionRequest)(nil), "proto.PermissionRequest")
	proto.RegisterType((*Array2DReply)(nil), "proto.Array2DReply")
	proto.RegisterType((*Array2DReplyD)(nil), "proto.Array2DReply.d")
}

func init() { proto.RegisterFile("casbin.proto", fileDescriptor_df0f85ba9164bca5) }

var fileDescriptor_df0f85ba9164bca5 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0xae, 0xd3, 0xa6, 0x77, 0xf9, 0x92, 0xb6, 0xc9, 0xa6, 0x17, 0xdc, 0xa8, 0x54, 0xc1, 0x02,
	0x11, 0x09, 0x14, 0xe9, 0x02, 0xc7, 0xc1, 0x49, 0x27, 0x9a, 0x94, 0x92, 0x80, 0xb8, 0x5e, 0xe4,
	0x5c, 0x11, 0xaf, 0xae, 0x77, 0x5a, 0x8c, 0x1c, 0xdb, 0xac, 0x9d, 0xbb, 0xe6, 0xa7, 0xf0, 0xc0,
	0x03, 0xff, 0x14, 0x79, 0x6d, 0xa7, 0x76, 0x48, 0x4a, 0x9d, 0xe6, 0xc9, 0xde, 0xf1, 0xcc, 0xf7,
	0xcd, 0xce, 0x7e, 0x33, 0x5e, 0x54, 0x4c, 0xc3, 0xbf, 0xb2, 0x9c, 0x8e, 0x27, 0xdc, 0xc0, 0x65,
	0x45, 0xf9, 0xd0, 0x7e, 0x03, 0xbb, 0xa0, 0x0f, 0xe7, 0xce, 0xb5, 0x2b, 0x4c, 0x12, 0x3a, 0xfd,
	0x39, 0x25, 0x3f, 0x60, 0xc7, 0x28, 0x4d, 0x5c, 0x4e, 0xf6, 0x3b, 0xba, 0x0d, 0x54, 0xa5, 0xa5,
	0xb4, 0x4b, 0xfa, 0x9d, 0x81, 0x7d, 0x8a, 0x3d, 0x83, 0x1b, 0x5e, 0x40, 0x62, 0x68, 0x38, 0xdc,
	0x26, 0xb5, 0xd0, 0x52, 0xda, 0x45, 0x3d, 0x6b, 0xd4, 0xbe, 0x44, 0x35, 0x83, 0xec, 0xd9, 0x33,
	0xa6, 0xe2, 0xc9, 0xef, 0xf2, 0xab, 0x90, 0xa8, 0x45, 0x3d, 0x59, 0x6a, 0x7f, 0x2b, 0xa8, 0x5d,
	0xd0, 0x87, 0x5e, 0x04, 0x91, 0xe4, 0xd1, 0x42, 0x39, 0x06, 0xbd, 0x30, 0x26, 0x14, 0x67, 0x92,
	0x36, 0xb1, 0x13, 0x80, 0x0b, 0xeb, 0x7d, 0xec, 0x50, 0x90, 0x0e, 0x29, 0x4b, 0x98, 0xab, 0xe9,
	0x3a, 0x0e, 0x99, 0xc1, 0x38, 0x10, 0x96, 0x73, 0xa3, 0x6e, 0x4b, 0x97, 0xac, 0x31, 0xe4, 0xe1,
	0x57, 0x63, 0x8f, 0x4c, 0xeb, 0xda, 0x22, 0xae, 0xee, 0xb4, 0x94, 0xf6, 0x53, 0x3d, 0x6d, 0xd2,
	0xbe, 0xc0, 0x41, 0x3a, 0xbd, 0xfb, 0x37, 0xa3, 0x63, 0x3f, 0xde, 0x77, 0xb2, 0x91, 0x36, 0x0e,
	0x28, 0xae, 0xc4, 0x30, 0x13, 0xb3, 0x68, 0x66, 0x0d, 0xec, 0x7a, 0x86, 0x30, 0x26, 0xbe, 0x5a,
	0x68, 0x6d, 0xb7, 0x4b, 0x7a, 0xbc, 0xd2, 0x3e, 0x46, 0xa9, 0xef, 0xba, 0x76, 0x44, 0x5d, 0xc5,
	0xb6, 0x20, 0x5f, 0x42, 0x3c, 0xd5, 0xc3, 0x57, 0xad, 0x8d, 0xca, 0xf9, 0xc4, 0x0b, 0x66, 0x09,
	0xe1, 0xea, 0xe4, 0x2a, 0x40, 0xec, 0xe9, 0xd9, 0x33, 0xed, 0x06, 0x7b, 0x23, 0xd7, 0xb6, 0xcc,
	0x59, 0xfe, 0x4c, 0x0f, 0x51, 0xf4, 0xde, 0xcd, 0xbc, 0xa4, 0xea, 0xd1, 0x22, 0x95, 0xff, 0x76,
	0x26, 0x7f, 0x1d, 0xd5, 0xb1, 0x35, 0xf1, 0x6c, 0x1a, 0x50, 0xb0, 0x21, 0x2e, 0x4d, 0x03, 0x7a,
	0x42, 0x18, 0xd1, 0x56, 0x42, 0x1f, 0x23, 0x5c, 0xa9, 0x8a, 0x24, 0x8e, 0x16, 0xda, 0x5f, 0x0a,
	0x9e, 0xfd, 0x68, 0xd9, 0x01, 0x09, 0xe2, 0x9b, 0xdd, 0xe9, 0x09, 0x70, 0x6d, 0x91, 0xcd, 0x7f,
	0x72, 0x38, 0xdd, 0x4a, 0x5d, 0x15, 0xf5, 0x94, 0x25, 0x14, 0x95, 0x5c, 0xfd, 0x6a, 0xd8, 0x53,
	0xf2, 0xd5, 0x1d, 0x99, 0x55, 0xda, 0xa4, 0x99, 0x38, 0xb8, 0xf4, 0x49, 0xe8, 0xae, 0xbd, 0x86,
	0x50, 0x18, 0x76, 0xa6, 0x3e, 0x89, 0x38, 0x27, 0xf9, 0x1e, 0xda, 0x84, 0x6b, 0x53, 0x2c, 0x72,
	0xf9, 0xae, 0xf9, 0xa8, 0x8d, 0x48, 0x4c, 0x2c, 0xdf, 0xb7, 0x5c, 0x67, 0x33, 0x34, 0x2d, 0x94,
	0xbd, 0x39, 0x64, 0x72, 0xd0, 0x69, 0x93, 0xf6, 0x33, 0x2a, 0xf2, 0x64, 0xba, 0x3f, 0x44, 0x67,
	0xf3, 0x19, 0x0a, 0xbc, 0x2b, 0x0f, 0xa6, 0xdc, 0x7d, 0x16, 0x4d, 0xa0, 0x4e, 0xda, 0xa1, 0xc3,
	0xf5, 0x02, 0xef, 0x36, 0xeb, 0x50, 0x38, 0xdb, 0x47, 0x81, 0x3f, 0x8f, 0x0f, 0xb1, 0xc0, 0x9f,
	0x77, 0xff, 0xa9, 0x63, 0xf7, 0x4c, 0x8e, 0x2e, 0x76, 0x86, 0x72, 0x6a, 0xa6, 0xb0, 0xa3, 0x18,
	0xe9, 0xbf, 0x13, 0xac, 0xf9, 0xd1, 0xb2, 0x4f, 0xa1, 0xe0, 0xb7, 0xd8, 0x29, 0x70, 0xd7, 0xca,
	0x4c, 0xbd, 0x73, 0xcc, 0x0e, 0x9f, 0x66, 0x63, 0xc9, 0x97, 0x08, 0xe1, 0x6b, 0x3c, 0x89, 0x41,
	0x59, 0xb2, 0x99, 0x6c, 0xbf, 0x37, 0xab, 0xb1, 0x79, 0xde, 0xb2, 0xda, 0x16, 0xfb, 0x06, 0xf8,
	0xc5, 0x35, 0x62, 0x11, 0xb2, 0x7a, 0x12, 0x98, 0xea, 0xda, 0x66, 0x2d, 0x6b, 0x9c, 0xc7, 0x8d,
	0x8d, 0xf7, 0x94, 0x3b, 0xee, 0x05, 0x4a, 0x3d, 0x9e, 0xd0, 0x1d, 0xc6, 0x1e, 0x99, 0x16, 0x58,
	0x9a, 0xe6, 0x2b, 0xec, 0xf7, 0x38, 0x0f, 0x87, 0x67, 0xfe, 0xd8, 0x6f, 0x51, 0xd1, 0x69, 0xe2,
	0xce, 0x93, 0x7d, 0x78, 0xe4, 0x6b, 0xd4, 0xa2, 0xc8, 0xf5, 0x88, 0x87, 0x38, 0x8c, 0xc2, 0xb3,
	0xad, 0xce, 0x8e, 0x63, 0xdf, 0xa5, 0x13, 0x60, 0x29, 0xd2, 0x1b, 0x1c, 0x65, 0x91, 0xd2, 0x09,
	0xe5, 0x87, 0x7b, 0x89, 0xd2, 0x80, 0x82, 0xfb, 0xce, 0xae, 0xbe, 0xa4, 0x1d, 0x64, 0x41, 0xf6,
	0x07, 0x14, 0xfc, 0x7f, 0x35, 0x56, 0x84, 0x0f, 0x51, 0x1b, 0x50, 0x90, 0xab, 0x1a, 0x2b, 0x90,
	0xde, 0xa0, 0x91, 0x42, 0x7a, 0x78, 0x35, 0x56, 0xee, 0xab, 0xd6, 0xe3, 0x7c, 0x20, 0xdc, 0xa9,
	0x67, 0x39, 0x37, 0xb9, 0x0f, 0xba, 0x8f, 0x46, 0xa2, 0xce, 0xb5, 0x31, 0x4e, 0x13, 0xb1, 0xac,
	0x8d, 0x70, 0x9e, 0x88, 0xe4, 0x71, 0x89, 0x8c, 0x70, 0x9c, 0xd5, 0xda, 0x02, 0x52, 0x7e, 0xb9,
	0x5d, 0xe2, 0x93, 0x25, 0xea, 0x7d, 0x34, 0xec, 0xf7, 0x52, 0x4d, 0x0b, 0x30, 0x79, 0xd4, 0x7c,
	0x2e, 0x45, 0xf4, 0xf0, 0x6a, 0xad, 0x80, 0x19, 0xe1, 0x28, 0xa5, 0xc5, 0x5c, 0xdb, 0x5a, 0x81,
	0x78, 0x89, 0x93, 0x45, 0x75, 0x6f, 0x02, 0xf6, 0x95, 0xec, 0xde, 0x9e, 0x6d, 0x8f, 0xa7, 0x57,
	0x7f, 0x90, 0x19, 0xf8, 0xf7, 0xcf, 0xed, 0xbb, 0x5b, 0x8c, 0xb6, 0xc5, 0xce, 0x50, 0x8f, 0x62,
	0x65, 0x36, 0x73, 0x80, 0xe4, 0x8f, 0xb6, 0x78, 0x8b, 0x5a, 0x0e, 0xf2, 0x1d, 0xf6, 0x22, 0x90,
	0xb7, 0xb9, 0xf9, 0xfb, 0x60, 0x29, 0xfe, 0xb7, 0x8f, 0xa3, 0xef, 0x99, 0x41, 0x78, 0x21, 0x58,
	0x9b, 0x3e, 0x89, 0xcf, 0x47, 0xff, 0x12, 0xe5, 0x08, 0x23, 0xbc, 0x5a, 0xe5, 0x21, 0x3f, 0x45,
	0x35, 0x45, 0x1e, 0x45, 0xe7, 0xa3, 0x7e, 0x81, 0xd2, 0xd0, 0xf0, 0xd7, 0xf9, 0xeb, 0x0e, 0x0d,
	0x7f, 0xbd, 0x9f, 0xdf, 0x6b, 0xd4, 0x86, 0x86, 0xff, 0x98, 0x91, 0x9a, 0x50, 0xaf, 0x8b, 0xd1,
	0xff, 0x1c, 0x0d, 0xcb, 0xed, 0xdc, 0x08, 0xcf, 0xec, 0xd0, 0xad, 0x11, 0xd6, 0xc9, 0x8f, 0xbc,
	0xfa, 0xe5, 0xe8, 0xea, 0x36, 0x0a, 0x17, 0x23, 0xe5, 0x6a, 0x57, 0x5a, 0xbf, 0xfa, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x19, 0xe7, 0x0a, 0xa4, 0x8d, 0x0e, 0x00, 0x00,
}
