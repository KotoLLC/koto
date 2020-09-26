// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: user.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
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

type UserFriendsFriendOfFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	InviteStatus string `protobuf:"bytes,2,opt,name=invite_status,json=inviteStatus,proto3" json:"invite_status,omitempty"`
}

func (x *UserFriendsFriendOfFriend) Reset() {
	*x = UserFriendsFriendOfFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsFriendOfFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsFriendOfFriend) ProtoMessage() {}

func (x *UserFriendsFriendOfFriend) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsFriendOfFriend.ProtoReflect.Descriptor instead.
func (*UserFriendsFriendOfFriend) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserFriendsFriendOfFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserFriendsFriendOfFriend) GetInviteStatus() string {
	if x != nil {
		return x.InviteStatus
	}
	return ""
}

type UserFriendsFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User                        `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Friends []*UserFriendsFriendOfFriend `protobuf:"bytes,2,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsFriend) Reset() {
	*x = UserFriendsFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsFriend) ProtoMessage() {}

func (x *UserFriendsFriend) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsFriend.ProtoReflect.Descriptor instead.
func (*UserFriendsFriend) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserFriendsFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserFriendsFriend) GetFriends() []*UserFriendsFriendOfFriend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserFriendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*UserFriendsFriend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsResponse) Reset() {
	*x = UserFriendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsResponse) ProtoMessage() {}

func (x *UserFriendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsResponse.ProtoReflect.Descriptor instead.
func (*UserFriendsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserFriendsResponse) GetFriends() []*UserFriendsFriend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserFriendsOfFriendsResponseFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	InviteStatus string  `protobuf:"bytes,2,opt,name=invite_status,json=inviteStatus,proto3" json:"invite_status,omitempty"`
	Friends      []*User `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsOfFriendsResponseFriend) Reset() {
	*x = UserFriendsOfFriendsResponseFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsOfFriendsResponseFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsOfFriendsResponseFriend) ProtoMessage() {}

func (x *UserFriendsOfFriendsResponseFriend) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsOfFriendsResponseFriend.ProtoReflect.Descriptor instead.
func (*UserFriendsOfFriendsResponseFriend) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserFriendsOfFriendsResponseFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserFriendsOfFriendsResponseFriend) GetInviteStatus() string {
	if x != nil {
		return x.InviteStatus
	}
	return ""
}

func (x *UserFriendsOfFriendsResponseFriend) GetFriends() []*User {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserFriendsOfFriendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*UserFriendsOfFriendsResponseFriend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsOfFriendsResponse) Reset() {
	*x = UserFriendsOfFriendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsOfFriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsOfFriendsResponse) ProtoMessage() {}

func (x *UserFriendsOfFriendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsOfFriendsResponse.ProtoReflect.Descriptor instead.
func (*UserFriendsOfFriendsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserFriendsOfFriendsResponse) GetFriends() []*UserFriendsOfFriendsResponseFriend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	IsAdmin   bool     `protobuf:"varint,2,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	OwnedHubs []string `protobuf:"bytes,3,rep,name=owned_hubs,json=ownedHubs,proto3" json:"owned_hubs,omitempty"`
}

func (x *UserMeResponse) Reset() {
	*x = UserMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeResponse) ProtoMessage() {}

func (x *UserMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeResponse.ProtoReflect.Descriptor instead.
func (*UserMeResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserMeResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserMeResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *UserMeResponse) GetOwnedHubs() []string {
	if x != nil {
		return x.OwnedHubs
	}
	return nil
}

type UserEditProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailChanged    bool   `protobuf:"varint,1,opt,name=email_changed,json=emailChanged,proto3" json:"email_changed,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	AvatarChanged   bool   `protobuf:"varint,3,opt,name=avatar_changed,json=avatarChanged,proto3" json:"avatar_changed,omitempty"`
	AvatarId        string `protobuf:"bytes,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	PasswordChanged bool   `protobuf:"varint,5,opt,name=password_changed,json=passwordChanged,proto3" json:"password_changed,omitempty"`
	CurrentPassword string `protobuf:"bytes,6,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	NewPassword     string `protobuf:"bytes,7,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *UserEditProfileRequest) Reset() {
	*x = UserEditProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEditProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEditProfileRequest) ProtoMessage() {}

func (x *UserEditProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEditProfileRequest.ProtoReflect.Descriptor instead.
func (*UserEditProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserEditProfileRequest) GetEmailChanged() bool {
	if x != nil {
		return x.EmailChanged
	}
	return false
}

func (x *UserEditProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserEditProfileRequest) GetAvatarChanged() bool {
	if x != nil {
		return x.AvatarChanged
	}
	return false
}

func (x *UserEditProfileRequest) GetAvatarId() string {
	if x != nil {
		return x.AvatarId
	}
	return ""
}

func (x *UserEditProfileRequest) GetPasswordChanged() bool {
	if x != nil {
		return x.PasswordChanged
	}
	return false
}

func (x *UserEditProfileRequest) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *UserEditProfileRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type UserUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *UserUsersRequest) Reset() {
	*x = UserUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUsersRequest) ProtoMessage() {}

func (x *UserUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUsersRequest.ProtoReflect.Descriptor instead.
func (*UserUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type UserUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserUsersResponse) Reset() {
	*x = UserUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUsersResponse) ProtoMessage() {}

func (x *UserUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUsersResponse.ProtoReflect.Descriptor instead.
func (*UserUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *UserUserRequest) Reset() {
	*x = UserUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUserRequest) ProtoMessage() {}

func (x *UserUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUserRequest.ProtoReflect.Descriptor instead.
func (*UserUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type UserUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserUserResponse) Reset() {
	*x = UserUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUserResponse) ProtoMessage() {}

func (x *UserUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUserResponse.ProtoReflect.Descriptor instead.
func (*UserUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *UserUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserRegisterFCMTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Os       string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *UserRegisterFCMTokenRequest) Reset() {
	*x = UserRegisterFCMTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterFCMTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterFCMTokenRequest) ProtoMessage() {}

func (x *UserRegisterFCMTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterFCMTokenRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterFCMTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *UserRegisterFCMTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserRegisterFCMTokenRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UserRegisterFCMTokenRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70,
	0x63, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f,
	0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x6c, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x47, 0x0a,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x22, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x23, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x69, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x5f, 0x68,
	0x75, 0x62, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x64,
	0x48, 0x75, 0x62, 0x73, 0x22, 0x90, 0x02, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2d, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x43, 0x4d, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x32, 0x8f, 0x03, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12,
	0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x02, 0x4d, 0x65, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x43, 0x4d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x46, 0x43, 0x4d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_user_proto_goTypes = []interface{}{
	(*UserFriendsFriendOfFriend)(nil),          // 0: rpc.UserFriendsFriendOfFriend
	(*UserFriendsFriend)(nil),                  // 1: rpc.UserFriendsFriend
	(*UserFriendsResponse)(nil),                // 2: rpc.UserFriendsResponse
	(*UserFriendsOfFriendsResponseFriend)(nil), // 3: rpc.UserFriendsOfFriendsResponseFriend
	(*UserFriendsOfFriendsResponse)(nil),       // 4: rpc.UserFriendsOfFriendsResponse
	(*UserMeResponse)(nil),                     // 5: rpc.UserMeResponse
	(*UserEditProfileRequest)(nil),             // 6: rpc.UserEditProfileRequest
	(*UserUsersRequest)(nil),                   // 7: rpc.UserUsersRequest
	(*UserUsersResponse)(nil),                  // 8: rpc.UserUsersResponse
	(*UserUserRequest)(nil),                    // 9: rpc.UserUserRequest
	(*UserUserResponse)(nil),                   // 10: rpc.UserUserResponse
	(*UserRegisterFCMTokenRequest)(nil),        // 11: rpc.UserRegisterFCMTokenRequest
	(*User)(nil),                               // 12: rpc.User
	(*Empty)(nil),                              // 13: rpc.Empty
}
var file_user_proto_depIdxs = []int32{
	12, // 0: rpc.UserFriendsFriendOfFriend.user:type_name -> rpc.User
	12, // 1: rpc.UserFriendsFriend.user:type_name -> rpc.User
	0,  // 2: rpc.UserFriendsFriend.friends:type_name -> rpc.UserFriendsFriendOfFriend
	1,  // 3: rpc.UserFriendsResponse.friends:type_name -> rpc.UserFriendsFriend
	12, // 4: rpc.UserFriendsOfFriendsResponseFriend.user:type_name -> rpc.User
	12, // 5: rpc.UserFriendsOfFriendsResponseFriend.friends:type_name -> rpc.User
	3,  // 6: rpc.UserFriendsOfFriendsResponse.friends:type_name -> rpc.UserFriendsOfFriendsResponseFriend
	12, // 7: rpc.UserMeResponse.user:type_name -> rpc.User
	12, // 8: rpc.UserUsersResponse.users:type_name -> rpc.User
	12, // 9: rpc.UserUserResponse.user:type_name -> rpc.User
	13, // 10: rpc.UserService.Friends:input_type -> rpc.Empty
	13, // 11: rpc.UserService.FriendsOfFriends:input_type -> rpc.Empty
	13, // 12: rpc.UserService.Me:input_type -> rpc.Empty
	6,  // 13: rpc.UserService.EditProfile:input_type -> rpc.UserEditProfileRequest
	7,  // 14: rpc.UserService.Users:input_type -> rpc.UserUsersRequest
	9,  // 15: rpc.UserService.User:input_type -> rpc.UserUserRequest
	11, // 16: rpc.UserService.RegisterFCMToken:input_type -> rpc.UserRegisterFCMTokenRequest
	2,  // 17: rpc.UserService.Friends:output_type -> rpc.UserFriendsResponse
	4,  // 18: rpc.UserService.FriendsOfFriends:output_type -> rpc.UserFriendsOfFriendsResponse
	5,  // 19: rpc.UserService.Me:output_type -> rpc.UserMeResponse
	13, // 20: rpc.UserService.EditProfile:output_type -> rpc.Empty
	8,  // 21: rpc.UserService.Users:output_type -> rpc.UserUsersResponse
	10, // 22: rpc.UserService.User:output_type -> rpc.UserUserResponse
	13, // 23: rpc.UserService.RegisterFCMToken:output_type -> rpc.Empty
	17, // [17:24] is the sub-list for method output_type
	10, // [10:17] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsFriendOfFriend); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsFriend); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsOfFriendsResponseFriend); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsOfFriendsResponse); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeResponse); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEditProfileRequest); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUsersRequest); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUsersResponse); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUserRequest); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUserResponse); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterFCMTokenRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
