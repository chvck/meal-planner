// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service/service.proto

package service

import (
	context "context"
	fmt "fmt"
	model "github.com/chvck/meal-planner/proto/model"
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

type AllRecipesRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllRecipesRequest) Reset()         { *m = AllRecipesRequest{} }
func (m *AllRecipesRequest) String() string { return proto.CompactTextString(m) }
func (*AllRecipesRequest) ProtoMessage()    {}
func (*AllRecipesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{0}
}

func (m *AllRecipesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllRecipesRequest.Unmarshal(m, b)
}
func (m *AllRecipesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllRecipesRequest.Marshal(b, m, deterministic)
}
func (m *AllRecipesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllRecipesRequest.Merge(m, src)
}
func (m *AllRecipesRequest) XXX_Size() int {
	return xxx_messageInfo_AllRecipesRequest.Size(m)
}
func (m *AllRecipesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllRecipesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllRecipesRequest proto.InternalMessageInfo

func (m *AllRecipesRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *AllRecipesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type AllRecipesResponse struct {
	Recipes              []*model.Recipe `protobuf:"bytes,1,rep,name=recipes,proto3" json:"recipes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AllRecipesResponse) Reset()         { *m = AllRecipesResponse{} }
func (m *AllRecipesResponse) String() string { return proto.CompactTextString(m) }
func (*AllRecipesResponse) ProtoMessage()    {}
func (*AllRecipesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{1}
}

func (m *AllRecipesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllRecipesResponse.Unmarshal(m, b)
}
func (m *AllRecipesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllRecipesResponse.Marshal(b, m, deterministic)
}
func (m *AllRecipesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllRecipesResponse.Merge(m, src)
}
func (m *AllRecipesResponse) XXX_Size() int {
	return xxx_messageInfo_AllRecipesResponse.Size(m)
}
func (m *AllRecipesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllRecipesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllRecipesResponse proto.InternalMessageInfo

func (m *AllRecipesResponse) GetRecipes() []*model.Recipe {
	if m != nil {
		return m.Recipes
	}
	return nil
}

type RecipeByIDRequest struct {
	RecipeId             int64    `protobuf:"varint,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecipeByIDRequest) Reset()         { *m = RecipeByIDRequest{} }
func (m *RecipeByIDRequest) String() string { return proto.CompactTextString(m) }
func (*RecipeByIDRequest) ProtoMessage()    {}
func (*RecipeByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{2}
}

func (m *RecipeByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipeByIDRequest.Unmarshal(m, b)
}
func (m *RecipeByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipeByIDRequest.Marshal(b, m, deterministic)
}
func (m *RecipeByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipeByIDRequest.Merge(m, src)
}
func (m *RecipeByIDRequest) XXX_Size() int {
	return xxx_messageInfo_RecipeByIDRequest.Size(m)
}
func (m *RecipeByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipeByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecipeByIDRequest proto.InternalMessageInfo

func (m *RecipeByIDRequest) GetRecipeId() int64 {
	if m != nil {
		return m.RecipeId
	}
	return 0
}

type RecipeByIDResponse struct {
	Recipe               *model.Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RecipeByIDResponse) Reset()         { *m = RecipeByIDResponse{} }
func (m *RecipeByIDResponse) String() string { return proto.CompactTextString(m) }
func (*RecipeByIDResponse) ProtoMessage()    {}
func (*RecipeByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{3}
}

func (m *RecipeByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipeByIDResponse.Unmarshal(m, b)
}
func (m *RecipeByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipeByIDResponse.Marshal(b, m, deterministic)
}
func (m *RecipeByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipeByIDResponse.Merge(m, src)
}
func (m *RecipeByIDResponse) XXX_Size() int {
	return xxx_messageInfo_RecipeByIDResponse.Size(m)
}
func (m *RecipeByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipeByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecipeByIDResponse proto.InternalMessageInfo

func (m *RecipeByIDResponse) GetRecipe() *model.Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

type CreateRecipeRequest struct {
	Recipe               *model.Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateRecipeRequest) Reset()         { *m = CreateRecipeRequest{} }
func (m *CreateRecipeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRecipeRequest) ProtoMessage()    {}
func (*CreateRecipeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{4}
}

func (m *CreateRecipeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRecipeRequest.Unmarshal(m, b)
}
func (m *CreateRecipeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRecipeRequest.Marshal(b, m, deterministic)
}
func (m *CreateRecipeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRecipeRequest.Merge(m, src)
}
func (m *CreateRecipeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRecipeRequest.Size(m)
}
func (m *CreateRecipeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRecipeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRecipeRequest proto.InternalMessageInfo

func (m *CreateRecipeRequest) GetRecipe() *model.Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

type CreateRecipeResponse struct {
	Recipe               *model.Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateRecipeResponse) Reset()         { *m = CreateRecipeResponse{} }
func (m *CreateRecipeResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRecipeResponse) ProtoMessage()    {}
func (*CreateRecipeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{5}
}

func (m *CreateRecipeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRecipeResponse.Unmarshal(m, b)
}
func (m *CreateRecipeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRecipeResponse.Marshal(b, m, deterministic)
}
func (m *CreateRecipeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRecipeResponse.Merge(m, src)
}
func (m *CreateRecipeResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRecipeResponse.Size(m)
}
func (m *CreateRecipeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRecipeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRecipeResponse proto.InternalMessageInfo

func (m *CreateRecipeResponse) GetRecipe() *model.Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

type UpdateRecipeRequest struct {
	Recipe               *model.Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateRecipeRequest) Reset()         { *m = UpdateRecipeRequest{} }
func (m *UpdateRecipeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRecipeRequest) ProtoMessage()    {}
func (*UpdateRecipeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{6}
}

func (m *UpdateRecipeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRecipeRequest.Unmarshal(m, b)
}
func (m *UpdateRecipeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRecipeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRecipeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRecipeRequest.Merge(m, src)
}
func (m *UpdateRecipeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRecipeRequest.Size(m)
}
func (m *UpdateRecipeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRecipeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRecipeRequest proto.InternalMessageInfo

func (m *UpdateRecipeRequest) GetRecipe() *model.Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

type UpdateRecipeResponse struct {
	Recipe               *model.Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateRecipeResponse) Reset()         { *m = UpdateRecipeResponse{} }
func (m *UpdateRecipeResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateRecipeResponse) ProtoMessage()    {}
func (*UpdateRecipeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{7}
}

func (m *UpdateRecipeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRecipeResponse.Unmarshal(m, b)
}
func (m *UpdateRecipeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRecipeResponse.Marshal(b, m, deterministic)
}
func (m *UpdateRecipeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRecipeResponse.Merge(m, src)
}
func (m *UpdateRecipeResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateRecipeResponse.Size(m)
}
func (m *UpdateRecipeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRecipeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRecipeResponse proto.InternalMessageInfo

func (m *UpdateRecipeResponse) GetRecipe() *model.Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

type DeleteRecipeRequest struct {
	RecipeId             int64    `protobuf:"varint,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRecipeRequest) Reset()         { *m = DeleteRecipeRequest{} }
func (m *DeleteRecipeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRecipeRequest) ProtoMessage()    {}
func (*DeleteRecipeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{8}
}

func (m *DeleteRecipeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRecipeRequest.Unmarshal(m, b)
}
func (m *DeleteRecipeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRecipeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRecipeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRecipeRequest.Merge(m, src)
}
func (m *DeleteRecipeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRecipeRequest.Size(m)
}
func (m *DeleteRecipeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRecipeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRecipeRequest proto.InternalMessageInfo

func (m *DeleteRecipeRequest) GetRecipeId() int64 {
	if m != nil {
		return m.RecipeId
	}
	return 0
}

type DeleteRecipeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRecipeResponse) Reset()         { *m = DeleteRecipeResponse{} }
func (m *DeleteRecipeResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRecipeResponse) ProtoMessage()    {}
func (*DeleteRecipeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{9}
}

func (m *DeleteRecipeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRecipeResponse.Unmarshal(m, b)
}
func (m *DeleteRecipeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRecipeResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRecipeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRecipeResponse.Merge(m, src)
}
func (m *DeleteRecipeResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRecipeResponse.Size(m)
}
func (m *DeleteRecipeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRecipeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRecipeResponse proto.InternalMessageInfo

type LoginUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserRequest) Reset()         { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()    {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{10}
}

func (m *LoginUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserRequest.Unmarshal(m, b)
}
func (m *LoginUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserRequest.Marshal(b, m, deterministic)
}
func (m *LoginUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserRequest.Merge(m, src)
}
func (m *LoginUserRequest) XXX_Size() int {
	return xxx_messageInfo_LoginUserRequest.Size(m)
}
func (m *LoginUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserRequest proto.InternalMessageInfo

func (m *LoginUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginUserResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserResponse) Reset()         { *m = LoginUserResponse{} }
func (m *LoginUserResponse) String() string { return proto.CompactTextString(m) }
func (*LoginUserResponse) ProtoMessage()    {}
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{11}
}

func (m *LoginUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserResponse.Unmarshal(m, b)
}
func (m *LoginUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserResponse.Marshal(b, m, deterministic)
}
func (m *LoginUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserResponse.Merge(m, src)
}
func (m *LoginUserResponse) XXX_Size() int {
	return xxx_messageInfo_LoginUserResponse.Size(m)
}
func (m *LoginUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserResponse proto.InternalMessageInfo

func (m *LoginUserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateUserRequest struct {
	User                 *model.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string      `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{12}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *model.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	User                 *model.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e2f8c9a3669d2, []int{13}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *model.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*AllRecipesRequest)(nil), "chvck.mealplanner.service.AllRecipesRequest")
	proto.RegisterType((*AllRecipesResponse)(nil), "chvck.mealplanner.service.AllRecipesResponse")
	proto.RegisterType((*RecipeByIDRequest)(nil), "chvck.mealplanner.service.RecipeByIDRequest")
	proto.RegisterType((*RecipeByIDResponse)(nil), "chvck.mealplanner.service.RecipeByIDResponse")
	proto.RegisterType((*CreateRecipeRequest)(nil), "chvck.mealplanner.service.CreateRecipeRequest")
	proto.RegisterType((*CreateRecipeResponse)(nil), "chvck.mealplanner.service.CreateRecipeResponse")
	proto.RegisterType((*UpdateRecipeRequest)(nil), "chvck.mealplanner.service.UpdateRecipeRequest")
	proto.RegisterType((*UpdateRecipeResponse)(nil), "chvck.mealplanner.service.UpdateRecipeResponse")
	proto.RegisterType((*DeleteRecipeRequest)(nil), "chvck.mealplanner.service.DeleteRecipeRequest")
	proto.RegisterType((*DeleteRecipeResponse)(nil), "chvck.mealplanner.service.DeleteRecipeResponse")
	proto.RegisterType((*LoginUserRequest)(nil), "chvck.mealplanner.service.LoginUserRequest")
	proto.RegisterType((*LoginUserResponse)(nil), "chvck.mealplanner.service.LoginUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "chvck.mealplanner.service.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "chvck.mealplanner.service.CreateUserResponse")
}

func init() { proto.RegisterFile("proto/service/service.proto", fileDescriptor_a34e2f8c9a3669d2) }

var fileDescriptor_a34e2f8c9a3669d2 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x28, 0x49, 0xeb, 0xa1, 0x07, 0xb2, 0x8d, 0x4a, 0x70, 0x85, 0x88, 0xf6, 0x14, 0x44,
	0x71, 0x21, 0x1c, 0x10, 0xc7, 0x96, 0x4a, 0xa8, 0x88, 0x52, 0x64, 0xd4, 0x0b, 0x17, 0xe4, 0xc6,
	0x13, 0x58, 0x75, 0xe3, 0x35, 0xbb, 0x6e, 0x11, 0x7f, 0xc1, 0x27, 0xa3, 0x78, 0x36, 0xc9, 0xa6,
	0x76, 0x57, 0x86, 0xe6, 0x92, 0x68, 0x66, 0xdf, 0x9b, 0xf7, 0xe2, 0xd9, 0xe7, 0xc0, 0x5e, 0xae,
	0x55, 0xa1, 0x0e, 0x0c, 0xea, 0x6b, 0x31, 0xc6, 0xf9, 0x77, 0x54, 0x76, 0xd9, 0xe3, 0xf1, 0x8f,
	0xeb, 0xf1, 0x65, 0x34, 0xc5, 0x44, 0xe6, 0x32, 0xc9, 0x32, 0xd4, 0x91, 0x05, 0x84, 0x8f, 0x88,
	0x37, 0x55, 0x29, 0x4a, 0xfa, 0x24, 0x0e, 0x3f, 0x84, 0xee, 0xa1, 0x94, 0x31, 0x8e, 0x45, 0x8e,
	0x26, 0xc6, 0x9f, 0x57, 0x68, 0x0a, 0xb6, 0x0b, 0x1d, 0x35, 0x99, 0x18, 0x2c, 0xfa, 0xad, 0x41,
	0x6b, 0xb8, 0x11, 0xdb, 0x8a, 0xf5, 0xa0, 0x2d, 0xc5, 0x54, 0x14, 0xfd, 0x7b, 0x83, 0xd6, 0xb0,
	0x1d, 0x53, 0xc1, 0xcf, 0x80, 0xb9, 0x23, 0x4c, 0xae, 0x32, 0x83, 0xec, 0x2d, 0x6c, 0x6a, 0x6a,
	0xf5, 0x5b, 0x83, 0x8d, 0xe1, 0x83, 0xd1, 0xd3, 0xa8, 0x6a, 0x8f, 0x9c, 0x10, 0x35, 0x9e, 0xe3,
	0xf9, 0x4b, 0xe8, 0x52, 0xeb, 0xe8, 0xf7, 0xc9, 0xf1, 0xdc, 0xd3, 0x1e, 0x04, 0x74, 0xfe, 0x4d,
	0xa4, 0xd6, 0xd6, 0x16, 0x35, 0x4e, 0x52, 0x7e, 0x0a, 0xcc, 0x65, 0x58, 0x0b, 0x6f, 0xa0, 0x43,
	0x88, 0x12, 0xdf, 0xc0, 0x81, 0x85, 0xf3, 0x4f, 0xb0, 0xf3, 0x4e, 0x63, 0x52, 0xa0, 0xed, 0x5b,
	0x0b, 0xff, 0x3d, 0xef, 0x0c, 0x7a, 0xab, 0xf3, 0xd6, 0x60, 0xf0, 0x3c, 0x4f, 0xd7, 0x6a, 0x70,
	0x75, 0xde, 0x5d, 0x0d, 0x8e, 0x60, 0xe7, 0x18, 0x25, 0xde, 0x34, 0xe8, 0x5d, 0xe2, 0x2e, 0xf4,
	0x56, 0x39, 0x64, 0x82, 0x7f, 0x80, 0x87, 0x1f, 0xd5, 0x77, 0x91, 0x9d, 0x1b, 0xd4, 0xf3, 0x41,
	0x21, 0x6c, 0x5d, 0x19, 0xd4, 0x59, 0x32, 0x25, 0x6b, 0x41, 0xbc, 0xa8, 0x67, 0x67, 0x79, 0x62,
	0xcc, 0x2f, 0xa5, 0xd3, 0xf2, 0xa2, 0x06, 0xf1, 0xa2, 0xe6, 0xcf, 0xa0, 0xeb, 0xcc, 0xb2, 0xbf,
	0xb2, 0x07, 0xed, 0x42, 0x5d, 0x62, 0x66, 0x27, 0x51, 0xc1, 0x2f, 0xa0, 0x4b, 0x4b, 0x73, 0x75,
	0x5f, 0xc1, 0xfd, 0x99, 0x8e, 0x7d, 0x1c, 0x4f, 0x6e, 0x7d, 0x1c, 0x25, 0xa7, 0x84, 0x7a, 0xed,
	0xbc, 0x07, 0xe6, 0x6a, 0x58, 0x3f, 0xff, 0x2e, 0x32, 0xfa, 0xd3, 0x01, 0x76, 0x8a, 0x89, 0xfc,
	0x4c, 0x80, 0x2f, 0x14, 0x7b, 0x26, 0x00, 0x96, 0xd1, 0x64, 0xfb, 0xd1, 0xad, 0x2f, 0x88, 0xa8,
	0xf2, 0x12, 0x08, 0x5f, 0x34, 0x44, 0x5b, 0xd3, 0x02, 0x60, 0x19, 0x41, 0xaf, 0x54, 0x25, 0xdb,
	0x5e, 0xa9, 0x9a, 0x5c, 0x2b, 0xd8, 0x76, 0xe3, 0xc4, 0x22, 0x0f, 0xbd, 0x26, 0xc7, 0xe1, 0x41,
	0x63, 0xfc, 0x52, 0xd0, 0x8d, 0x87, 0x57, 0xb0, 0x26, 0x97, 0x5e, 0xc1, 0xda, 0xdc, 0x29, 0xd8,
	0x76, 0xa3, 0xe0, 0x15, 0xac, 0xc9, 0x99, 0x57, 0xb0, 0x2e, 0x63, 0xb3, 0xed, 0x2d, 0x2f, 0xa2,
	0x77, 0x7b, 0x95, 0x4c, 0x78, 0xb7, 0x57, 0x73, 0xbb, 0x27, 0x10, 0x2c, 0x22, 0xc8, 0x9e, 0x7b,
	0xb8, 0x37, 0x43, 0x1f, 0xee, 0x37, 0x03, 0x93, 0xce, 0x51, 0xf0, 0x75, 0xd3, 0x1e, 0x5e, 0x74,
	0xca, 0xff, 0xba, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xe2, 0xa6, 0x36, 0x3e, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MealPlannerServiceClient is the client API for MealPlannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MealPlannerServiceClient interface {
	AllRecipes(ctx context.Context, in *AllRecipesRequest, opts ...grpc.CallOption) (*AllRecipesResponse, error)
	RecipeByID(ctx context.Context, in *RecipeByIDRequest, opts ...grpc.CallOption) (*RecipeByIDResponse, error)
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error)
	DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type mealPlannerServiceClient struct {
	cc *grpc.ClientConn
}

func NewMealPlannerServiceClient(cc *grpc.ClientConn) MealPlannerServiceClient {
	return &mealPlannerServiceClient{cc}
}

func (c *mealPlannerServiceClient) AllRecipes(ctx context.Context, in *AllRecipesRequest, opts ...grpc.CallOption) (*AllRecipesResponse, error) {
	out := new(AllRecipesResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/AllRecipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mealPlannerServiceClient) RecipeByID(ctx context.Context, in *RecipeByIDRequest, opts ...grpc.CallOption) (*RecipeByIDResponse, error) {
	out := new(RecipeByIDResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/RecipeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mealPlannerServiceClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/CreateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mealPlannerServiceClient) UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error) {
	out := new(UpdateRecipeResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/UpdateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mealPlannerServiceClient) DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error) {
	out := new(DeleteRecipeResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/DeleteRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mealPlannerServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mealPlannerServiceClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/chvck.mealplanner.service.MealPlannerService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MealPlannerServiceServer is the server API for MealPlannerService service.
type MealPlannerServiceServer interface {
	AllRecipes(context.Context, *AllRecipesRequest) (*AllRecipesResponse, error)
	RecipeByID(context.Context, *RecipeByIDRequest) (*RecipeByIDResponse, error)
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error)
	DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
}

func RegisterMealPlannerServiceServer(s *grpc.Server, srv MealPlannerServiceServer) {
	s.RegisterService(&_MealPlannerService_serviceDesc, srv)
}

func _MealPlannerService_AllRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllRecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).AllRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/AllRecipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).AllRecipes(ctx, req.(*AllRecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MealPlannerService_RecipeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).RecipeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/RecipeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).RecipeByID(ctx, req.(*RecipeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MealPlannerService_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/CreateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MealPlannerService_UpdateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).UpdateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/UpdateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).UpdateRecipe(ctx, req.(*UpdateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MealPlannerService_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/DeleteRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).DeleteRecipe(ctx, req.(*DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MealPlannerService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MealPlannerService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MealPlannerServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chvck.mealplanner.service.MealPlannerService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MealPlannerServiceServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MealPlannerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chvck.mealplanner.service.MealPlannerService",
	HandlerType: (*MealPlannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllRecipes",
			Handler:    _MealPlannerService_AllRecipes_Handler,
		},
		{
			MethodName: "RecipeByID",
			Handler:    _MealPlannerService_RecipeByID_Handler,
		},
		{
			MethodName: "CreateRecipe",
			Handler:    _MealPlannerService_CreateRecipe_Handler,
		},
		{
			MethodName: "UpdateRecipe",
			Handler:    _MealPlannerService_UpdateRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _MealPlannerService_DeleteRecipe_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _MealPlannerService_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _MealPlannerService_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/service.proto",
}