// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: mgmt/v1alpha1/auth.proto

package mgmtv1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type LoginCliRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The oauth code
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// The oauth redirect uri that the client uses during the oauth request
	RedirectUri string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *LoginCliRequest) Reset() {
	*x = LoginCliRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCliRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCliRequest) ProtoMessage() {}

func (x *LoginCliRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCliRequest.ProtoReflect.Descriptor instead.
func (*LoginCliRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginCliRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LoginCliRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type LoginCliResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The access token that is returned on successful login
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *LoginCliResponse) Reset() {
	*x = LoginCliResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCliResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCliResponse) ProtoMessage() {}

func (x *LoginCliResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCliResponse.ProtoReflect.Descriptor instead.
func (*LoginCliResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginCliResponse) GetAccessToken() *AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type GetAuthStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthStatusRequest) Reset() {
	*x = GetAuthStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthStatusRequest) ProtoMessage() {}

func (x *GetAuthStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAuthStatusRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{2}
}

type GetAuthStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not the server has authentication enabled.
	// This tells the client if it is expected to send access tokens.
	IsEnabled bool `protobuf:"varint,1,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
}

func (x *GetAuthStatusResponse) Reset() {
	*x = GetAuthStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthStatusResponse) ProtoMessage() {}

func (x *GetAuthStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthStatusResponse.ProtoReflect.Descriptor instead.
func (*GetAuthStatusResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthStatusResponse) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

// A decoded representation of an Access token from the backing auth server
type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The access token that will be provided in subsequent requests to provide authenticated access to the Api
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Token that can be used to retrieve a refreshed access token.
	// Will not be provided if the offline_access scope is not provided in the initial login flow.
	RefreshToken *string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3,oneof" json:"refresh_token,omitempty"`
	// Relative time in seconds that the access token will expire. Combine with the current time to get the expires_at time.
	ExpiresIn int64 `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	// The scopes that the access token have
	Scope string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	// The identity token of the authenticated user
	IdToken *string `protobuf:"bytes,5,opt,name=id_token,json=idToken,proto3,oneof" json:"id_token,omitempty"`
	// The token type. For JWTs, this will be `Bearer`
	TokenType string `protobuf:"bytes,6,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *AccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AccessToken) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

func (x *AccessToken) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *AccessToken) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *AccessToken) GetIdToken() string {
	if x != nil && x.IdToken != nil {
		return *x.IdToken
	}
	return ""
}

func (x *AccessToken) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

type GetAuthorizeUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state that's generated by the client that is passed along to prevent tampering
	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// The redirect uri that the client will be redirected back to during the auth request
	RedirectUri string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// The scopes the client is requesting as a part of the oauth login request
	Scope string `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *GetAuthorizeUrlRequest) Reset() {
	*x = GetAuthorizeUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorizeUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorizeUrlRequest) ProtoMessage() {}

func (x *GetAuthorizeUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorizeUrlRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorizeUrlRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *GetAuthorizeUrlRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetAuthorizeUrlRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *GetAuthorizeUrlRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type GetAuthorizeUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The generated url that is the client will be redirected to during the Oauth flow
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetAuthorizeUrlResponse) Reset() {
	*x = GetAuthorizeUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorizeUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorizeUrlResponse) ProtoMessage() {}

func (x *GetAuthorizeUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorizeUrlResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorizeUrlResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *GetAuthorizeUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetCliIssuerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCliIssuerRequest) Reset() {
	*x = GetCliIssuerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCliIssuerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCliIssuerRequest) ProtoMessage() {}

func (x *GetCliIssuerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCliIssuerRequest.ProtoReflect.Descriptor instead.
func (*GetCliIssuerRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{7}
}

type GetCliIssuerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The backing authentication issuer url
	IssuerUrl string `protobuf:"bytes,1,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`
	// The audience that will be used in the access token. This corresponds to the "aud" claim
	Audience string `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
}

func (x *GetCliIssuerResponse) Reset() {
	*x = GetCliIssuerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCliIssuerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCliIssuerResponse) ProtoMessage() {}

func (x *GetCliIssuerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCliIssuerResponse.ProtoReflect.Descriptor instead.
func (*GetCliIssuerResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *GetCliIssuerResponse) GetIssuerUrl() string {
	if x != nil {
		return x.IssuerUrl
	}
	return ""
}

func (x *GetCliIssuerResponse) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

type RefreshCliRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token used to retrieve a new access token.
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshCliRequest) Reset() {
	*x = RefreshCliRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshCliRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshCliRequest) ProtoMessage() {}

func (x *RefreshCliRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshCliRequest.ProtoReflect.Descriptor instead.
func (*RefreshCliRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *RefreshCliRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshCliResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The access token that is returned on successful refresh
	AccessToken *AccessToken `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *RefreshCliResponse) Reset() {
	*x = RefreshCliResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshCliResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshCliResponse) ProtoMessage() {}

func (x *RefreshCliResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshCliResponse.ProtoReflect.Descriptor instead.
func (*RefreshCliResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{10}
}

func (x *RefreshCliResponse) GetAccessToken() *AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type CheckTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckTokenRequest) Reset() {
	*x = CheckTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenRequest) ProtoMessage() {}

func (x *CheckTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenRequest.ProtoReflect.Descriptor instead.
func (*CheckTokenRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{11}
}

type CheckTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckTokenResponse) Reset() {
	*x = CheckTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenResponse) ProtoMessage() {}

func (x *CheckTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_auth_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenResponse.ProtoReflect.Descriptor instead.
func (*CheckTokenResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_auth_proto_rawDescGZIP(), []int{12}
}

var File_mgmt_v1alpha1_auth_proto protoreflect.FileDescriptor

var file_mgmt_v1alpha1_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43,
	0x6c, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55,
	0x72, 0x69, 0x22, 0x51, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6c, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xed, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x69, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x1d, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0x41, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x53, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43,
	0x6c, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa3, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6c, 0x69,
	0x12, 0x1e, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6c,
	0x69, 0x12, 0x20, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x69, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x2e, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc5, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x63, 0x6c, 0x65, 0x75,
	0x73, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x6f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x67, 0x6d, 0x74, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0e, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mgmt_v1alpha1_auth_proto_rawDescOnce sync.Once
	file_mgmt_v1alpha1_auth_proto_rawDescData = file_mgmt_v1alpha1_auth_proto_rawDesc
)

func file_mgmt_v1alpha1_auth_proto_rawDescGZIP() []byte {
	file_mgmt_v1alpha1_auth_proto_rawDescOnce.Do(func() {
		file_mgmt_v1alpha1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_mgmt_v1alpha1_auth_proto_rawDescData)
	})
	return file_mgmt_v1alpha1_auth_proto_rawDescData
}

var file_mgmt_v1alpha1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_mgmt_v1alpha1_auth_proto_goTypes = []any{
	(*LoginCliRequest)(nil),         // 0: mgmt.v1alpha1.LoginCliRequest
	(*LoginCliResponse)(nil),        // 1: mgmt.v1alpha1.LoginCliResponse
	(*GetAuthStatusRequest)(nil),    // 2: mgmt.v1alpha1.GetAuthStatusRequest
	(*GetAuthStatusResponse)(nil),   // 3: mgmt.v1alpha1.GetAuthStatusResponse
	(*AccessToken)(nil),             // 4: mgmt.v1alpha1.AccessToken
	(*GetAuthorizeUrlRequest)(nil),  // 5: mgmt.v1alpha1.GetAuthorizeUrlRequest
	(*GetAuthorizeUrlResponse)(nil), // 6: mgmt.v1alpha1.GetAuthorizeUrlResponse
	(*GetCliIssuerRequest)(nil),     // 7: mgmt.v1alpha1.GetCliIssuerRequest
	(*GetCliIssuerResponse)(nil),    // 8: mgmt.v1alpha1.GetCliIssuerResponse
	(*RefreshCliRequest)(nil),       // 9: mgmt.v1alpha1.RefreshCliRequest
	(*RefreshCliResponse)(nil),      // 10: mgmt.v1alpha1.RefreshCliResponse
	(*CheckTokenRequest)(nil),       // 11: mgmt.v1alpha1.CheckTokenRequest
	(*CheckTokenResponse)(nil),      // 12: mgmt.v1alpha1.CheckTokenResponse
}
var file_mgmt_v1alpha1_auth_proto_depIdxs = []int32{
	4,  // 0: mgmt.v1alpha1.LoginCliResponse.access_token:type_name -> mgmt.v1alpha1.AccessToken
	4,  // 1: mgmt.v1alpha1.RefreshCliResponse.access_token:type_name -> mgmt.v1alpha1.AccessToken
	0,  // 2: mgmt.v1alpha1.AuthService.LoginCli:input_type -> mgmt.v1alpha1.LoginCliRequest
	9,  // 3: mgmt.v1alpha1.AuthService.RefreshCli:input_type -> mgmt.v1alpha1.RefreshCliRequest
	11, // 4: mgmt.v1alpha1.AuthService.CheckToken:input_type -> mgmt.v1alpha1.CheckTokenRequest
	7,  // 5: mgmt.v1alpha1.AuthService.GetCliIssuer:input_type -> mgmt.v1alpha1.GetCliIssuerRequest
	5,  // 6: mgmt.v1alpha1.AuthService.GetAuthorizeUrl:input_type -> mgmt.v1alpha1.GetAuthorizeUrlRequest
	2,  // 7: mgmt.v1alpha1.AuthService.GetAuthStatus:input_type -> mgmt.v1alpha1.GetAuthStatusRequest
	1,  // 8: mgmt.v1alpha1.AuthService.LoginCli:output_type -> mgmt.v1alpha1.LoginCliResponse
	10, // 9: mgmt.v1alpha1.AuthService.RefreshCli:output_type -> mgmt.v1alpha1.RefreshCliResponse
	12, // 10: mgmt.v1alpha1.AuthService.CheckToken:output_type -> mgmt.v1alpha1.CheckTokenResponse
	8,  // 11: mgmt.v1alpha1.AuthService.GetCliIssuer:output_type -> mgmt.v1alpha1.GetCliIssuerResponse
	6,  // 12: mgmt.v1alpha1.AuthService.GetAuthorizeUrl:output_type -> mgmt.v1alpha1.GetAuthorizeUrlResponse
	3,  // 13: mgmt.v1alpha1.AuthService.GetAuthStatus:output_type -> mgmt.v1alpha1.GetAuthStatusResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_mgmt_v1alpha1_auth_proto_init() }
func file_mgmt_v1alpha1_auth_proto_init() {
	if File_mgmt_v1alpha1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mgmt_v1alpha1_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoginCliRequest); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LoginCliResponse); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuthStatusRequest); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuthStatusResponse); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AccessToken); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuthorizeUrlRequest); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAuthorizeUrlResponse); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetCliIssuerRequest); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetCliIssuerResponse); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshCliRequest); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshCliResponse); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CheckTokenRequest); i {
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
		file_mgmt_v1alpha1_auth_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*CheckTokenResponse); i {
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
	file_mgmt_v1alpha1_auth_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mgmt_v1alpha1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mgmt_v1alpha1_auth_proto_goTypes,
		DependencyIndexes: file_mgmt_v1alpha1_auth_proto_depIdxs,
		MessageInfos:      file_mgmt_v1alpha1_auth_proto_msgTypes,
	}.Build()
	File_mgmt_v1alpha1_auth_proto = out.File
	file_mgmt_v1alpha1_auth_proto_rawDesc = nil
	file_mgmt_v1alpha1_auth_proto_goTypes = nil
	file_mgmt_v1alpha1_auth_proto_depIdxs = nil
}
