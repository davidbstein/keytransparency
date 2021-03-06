// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type/type.proto

/*
Package type_go_proto is a generated protocol buffer package.

It is generated from these files:
	type/type.proto
	type/keymaster.proto
	type/authz.proto

It has these top-level messages:
	User
	Metadata
	SigningKey
	VerifyingKey
	KeySet
*/
package type_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"
import google_crypto_tink "github.com/google/tink/proto/tink_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// User represents plain account information that gets committed to and obfuscated in Entry.
type User struct {
	// domain_id specifies the domain.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	// app_id specifies that application.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	// user_id specifies the user.
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// public_key_data is the public key material for this account.
	PublicKeyData []byte `protobuf:"bytes,4,opt,name=public_key_data,json=publicKeyData,proto3" json:"public_key_data,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys *google_crypto_tink.Keyset `protobuf:"bytes,5,opt,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
	// status is set when account is part of a batch operation.
	Status *google_rpc.Status `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *User) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *User) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *User) GetPublicKeyData() []byte {
	if m != nil {
		return m.PublicKeyData
	}
	return nil
}

func (m *User) GetAuthorizedKeys() *google_crypto_tink.Keyset {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

func (m *User) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "google.keytransparency.type.User")
}

func init() { proto.RegisterFile("type/type.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xf7, 0x6d, 0x47, 0x1b, 0x3f, 0x0a, 0x01, 0xe9, 0xd0, 0x6e, 0x8a, 0x0b, 0x29,
	0x2e, 0x12, 0xd0, 0xb5, 0x08, 0xea, 0xa6, 0x74, 0x37, 0xe2, 0xc6, 0xcd, 0x90, 0x26, 0x61, 0x1a,
	0xda, 0x4e, 0x2e, 0xc9, 0x9d, 0x45, 0xfc, 0xcd, 0xfe, 0x08, 0x49, 0x32, 0x2a, 0xb8, 0xc9, 0xc7,
	0x73, 0xce, 0x81, 0x7b, 0x0f, 0x99, 0x62, 0x00, 0xcd, 0xe3, 0xc1, 0xc0, 0x59, 0xb4, 0x74, 0xd1,
	0x5a, 0xdb, 0x1e, 0x34, 0xdb, 0xeb, 0x80, 0x4e, 0x74, 0x1e, 0x84, 0xd3, 0x9d, 0x0c, 0x2c, 0x5a,
	0xe6, 0xb3, 0x2c, 0x72, 0x07, 0x92, 0x7b, 0x14, 0xd8, 0xfb, 0x9c, 0x9a, 0x13, 0x34, 0xdd, 0x3e,
	0xbf, 0xaf, 0x3f, 0x0b, 0x32, 0x7a, 0xf3, 0xda, 0xd1, 0x05, 0x99, 0x28, 0x7b, 0x14, 0xa6, 0x6b,
	0x8c, 0xaa, 0x8a, 0x65, 0xb1, 0x9a, 0xd4, 0xa7, 0x19, 0xac, 0x15, 0xbd, 0x22, 0xa5, 0x00, 0x88,
	0xca, 0xbf, 0xa4, 0x8c, 0x05, 0xc0, 0x5a, 0xd1, 0x19, 0x39, 0xe9, 0xbd, 0x76, 0x91, 0xff, 0x4f,
	0xbc, 0x8c, 0xdf, 0xb5, 0xa2, 0x37, 0x64, 0x0a, 0xfd, 0xf6, 0x60, 0x64, 0xb3, 0xd7, 0xa1, 0x51,
	0x02, 0x45, 0x35, 0x5a, 0x16, 0xab, 0xf3, 0xfa, 0x22, 0xe3, 0x8d, 0x0e, 0x2f, 0x02, 0x05, 0x7d,
	0x26, 0x53, 0xd1, 0xe3, 0xce, 0x3a, 0xf3, 0xa1, 0x55, 0xf4, 0xfa, 0x6a, 0xbc, 0x2c, 0x56, 0x67,
	0x77, 0x73, 0x36, 0x6c, 0x26, 0x5d, 0x00, 0xb4, 0x2c, 0x4d, 0xbc, 0xd1, 0xc1, 0x6b, 0xac, 0x2f,
	0x7f, 0x23, 0x91, 0xd0, 0x5b, 0x52, 0xe6, 0xf5, 0xaa, 0x32, 0x65, 0xe9, 0x77, 0xd6, 0x81, 0x64,
	0xaf, 0x49, 0xa9, 0x07, 0xc7, 0xd3, 0xe3, 0xfb, 0x43, 0x6b, 0x70, 0xd7, 0x6f, 0x99, 0xb4, 0x47,
	0x3e, 0x14, 0xf4, 0xa7, 0x3d, 0x2e, 0xad, 0xd3, 0x5c, 0x80, 0xe1, 0x3f, 0x75, 0x37, 0xad, 0x6d,
	0x52, 0x5f, 0xdb, 0x32, 0x5d, 0xf7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xe0, 0x10, 0x2d,
	0x8b, 0x01, 0x00, 0x00,
}
