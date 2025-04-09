// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: customer.proto

package customer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Branch struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Name        *string                `protobuf:"bytes,2,opt,name=name"`
	xxx_hidden_CustomerId  *string                `protobuf:"bytes,3,opt,name=customer_id,json=customerId"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Branch) Reset() {
	*x = Branch{}
	mi := &file_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Branch) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *Branch) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *Branch) GetCustomerId() string {
	if x != nil {
		if x.xxx_hidden_CustomerId != nil {
			return *x.xxx_hidden_CustomerId
		}
		return ""
	}
	return ""
}

func (x *Branch) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *Branch) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *Branch) SetCustomerId(v string) {
	x.xxx_hidden_CustomerId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *Branch) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Branch) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *Branch) HasCustomerId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *Branch) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *Branch) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Name = nil
}

func (x *Branch) ClearCustomerId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_CustomerId = nil
}

type Branch_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id         *string
	Name       *string
	CustomerId *string
}

func (b0 Branch_builder) Build() *Branch {
	m0 := &Branch{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_Id = b.Id
	}
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Name = b.Name
	}
	if b.CustomerId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_CustomerId = b.CustomerId
	}
	return m0
}

type BranchSummary struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Name        *string                `protobuf:"bytes,2,opt,name=name"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *BranchSummary) Reset() {
	*x = BranchSummary{}
	mi := &file_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BranchSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchSummary) ProtoMessage() {}

func (x *BranchSummary) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BranchSummary) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *BranchSummary) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *BranchSummary) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *BranchSummary) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *BranchSummary) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *BranchSummary) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *BranchSummary) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *BranchSummary) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Name = nil
}

type BranchSummary_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id   *string
	Name *string
}

func (b0 BranchSummary_builder) Build() *BranchSummary {
	m0 := &BranchSummary{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_Id = b.Id
	}
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_Name = b.Name
	}
	return m0
}

type GetBranchesForContactRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ContactId   *string                `protobuf:"bytes,1,opt,name=contact_id,json=contactId"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetBranchesForContactRequest) Reset() {
	*x = GetBranchesForContactRequest{}
	mi := &file_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchesForContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchesForContactRequest) ProtoMessage() {}

func (x *GetBranchesForContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetBranchesForContactRequest) GetContactId() string {
	if x != nil {
		if x.xxx_hidden_ContactId != nil {
			return *x.xxx_hidden_ContactId
		}
		return ""
	}
	return ""
}

func (x *GetBranchesForContactRequest) SetContactId(v string) {
	x.xxx_hidden_ContactId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *GetBranchesForContactRequest) HasContactId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetBranchesForContactRequest) ClearContactId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_ContactId = nil
}

type GetBranchesForContactRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	ContactId *string
}

func (b0 GetBranchesForContactRequest_builder) Build() *GetBranchesForContactRequest {
	m0 := &GetBranchesForContactRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.ContactId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_ContactId = b.ContactId
	}
	return m0
}

type GetBranchesForContactResponse struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Branches *[]*BranchSummary      `protobuf:"bytes,1,rep,name=branches"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *GetBranchesForContactResponse) Reset() {
	*x = GetBranchesForContactResponse{}
	mi := &file_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchesForContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchesForContactResponse) ProtoMessage() {}

func (x *GetBranchesForContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetBranchesForContactResponse) GetBranches() []*BranchSummary {
	if x != nil {
		if x.xxx_hidden_Branches != nil {
			return *x.xxx_hidden_Branches
		}
	}
	return nil
}

func (x *GetBranchesForContactResponse) SetBranches(v []*BranchSummary) {
	x.xxx_hidden_Branches = &v
}

type GetBranchesForContactResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Branches []*BranchSummary
}

func (b0 GetBranchesForContactResponse_builder) Build() *GetBranchesForContactResponse {
	m0 := &GetBranchesForContactResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Branches = &b.Branches
	return m0
}

type GetBranchRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetBranchRequest) Reset() {
	*x = GetBranchRequest{}
	mi := &file_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchRequest) ProtoMessage() {}

func (x *GetBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetBranchRequest) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *GetBranchRequest) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *GetBranchRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetBranchRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

type GetBranchRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id *string
}

func (b0 GetBranchRequest_builder) Build() *GetBranchRequest {
	m0 := &GetBranchRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Id = b.Id
	}
	return m0
}

type GetBranchResponse struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Branch *Branch                `protobuf:"bytes,1,opt,name=branch"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetBranchResponse) Reset() {
	*x = GetBranchResponse{}
	mi := &file_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchResponse) ProtoMessage() {}

func (x *GetBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetBranchResponse) GetBranch() *Branch {
	if x != nil {
		return x.xxx_hidden_Branch
	}
	return nil
}

func (x *GetBranchResponse) SetBranch(v *Branch) {
	x.xxx_hidden_Branch = v
}

func (x *GetBranchResponse) HasBranch() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Branch != nil
}

func (x *GetBranchResponse) ClearBranch() {
	x.xxx_hidden_Branch = nil
}

type GetBranchResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Branch *Branch
}

func (b0 GetBranchResponse_builder) Build() *GetBranchResponse {
	m0 := &GetBranchResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Branch = b.Branch
	return m0
}

type GetRecentPurchasesByBranchRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Page        int32                  `protobuf:"varint,2,opt,name=page"`
	xxx_hidden_PageSize    int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetRecentPurchasesByBranchRequest) Reset() {
	*x = GetRecentPurchasesByBranchRequest{}
	mi := &file_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecentPurchasesByBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentPurchasesByBranchRequest) ProtoMessage() {}

func (x *GetRecentPurchasesByBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetRecentPurchasesByBranchRequest) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *GetRecentPurchasesByBranchRequest) GetPage() int32 {
	if x != nil {
		return x.xxx_hidden_Page
	}
	return 0
}

func (x *GetRecentPurchasesByBranchRequest) GetPageSize() int32 {
	if x != nil {
		return x.xxx_hidden_PageSize
	}
	return 0
}

func (x *GetRecentPurchasesByBranchRequest) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *GetRecentPurchasesByBranchRequest) SetPage(v int32) {
	x.xxx_hidden_Page = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *GetRecentPurchasesByBranchRequest) SetPageSize(v int32) {
	x.xxx_hidden_PageSize = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *GetRecentPurchasesByBranchRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetRecentPurchasesByBranchRequest) HasPage() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *GetRecentPurchasesByBranchRequest) HasPageSize() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *GetRecentPurchasesByBranchRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *GetRecentPurchasesByBranchRequest) ClearPage() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Page = 0
}

func (x *GetRecentPurchasesByBranchRequest) ClearPageSize() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_PageSize = 0
}

type GetRecentPurchasesByBranchRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id       *string
	Page     *int32
	PageSize *int32
}

func (b0 GetRecentPurchasesByBranchRequest_builder) Build() *GetRecentPurchasesByBranchRequest {
	m0 := &GetRecentPurchasesByBranchRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_Id = b.Id
	}
	if b.Page != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Page = *b.Page
	}
	if b.PageSize != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_PageSize = *b.PageSize
	}
	return m0
}

type GetRecentPurchasesByBranchResponse struct {
	state                   protoimpl.MessageState                      `protogen:"opaque.v1"`
	xxx_hidden_Items        *[]*GetRecentPurchasesByBranchResponse_Item `protobuf:"bytes,1,rep,name=items"`
	xxx_hidden_TotalRecords int32                                       `protobuf:"varint,2,opt,name=total_records,json=totalRecords"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *GetRecentPurchasesByBranchResponse) Reset() {
	*x = GetRecentPurchasesByBranchResponse{}
	mi := &file_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecentPurchasesByBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentPurchasesByBranchResponse) ProtoMessage() {}

func (x *GetRecentPurchasesByBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetRecentPurchasesByBranchResponse) GetItems() []*GetRecentPurchasesByBranchResponse_Item {
	if x != nil {
		if x.xxx_hidden_Items != nil {
			return *x.xxx_hidden_Items
		}
	}
	return nil
}

func (x *GetRecentPurchasesByBranchResponse) GetTotalRecords() int32 {
	if x != nil {
		return x.xxx_hidden_TotalRecords
	}
	return 0
}

func (x *GetRecentPurchasesByBranchResponse) SetItems(v []*GetRecentPurchasesByBranchResponse_Item) {
	x.xxx_hidden_Items = &v
}

func (x *GetRecentPurchasesByBranchResponse) SetTotalRecords(v int32) {
	x.xxx_hidden_TotalRecords = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *GetRecentPurchasesByBranchResponse) HasTotalRecords() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *GetRecentPurchasesByBranchResponse) ClearTotalRecords() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_TotalRecords = 0
}

type GetRecentPurchasesByBranchResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Items        []*GetRecentPurchasesByBranchResponse_Item
	TotalRecords *int32
}

func (b0 GetRecentPurchasesByBranchResponse_builder) Build() *GetRecentPurchasesByBranchResponse {
	m0 := &GetRecentPurchasesByBranchResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Items = &b.Items
	if b.TotalRecords != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_TotalRecords = *b.TotalRecords
	}
	return m0
}

type GetRecentPurchasesByBranchResponse_Item struct {
	state                         protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ProductId          *string                `protobuf:"bytes,1,opt,name=product_id,json=productId"`
	xxx_hidden_ProductSn          *string                `protobuf:"bytes,2,opt,name=product_sn,json=productSn"`
	xxx_hidden_ProductName        *string                `protobuf:"bytes,3,opt,name=product_name,json=productName"`
	xxx_hidden_ProductDescription *string                `protobuf:"bytes,4,opt,name=product_description,json=productDescription"`
	xxx_hidden_OrderedQuantity    float64                `protobuf:"fixed64,5,opt,name=ordered_quantity,json=orderedQuantity"`
	xxx_hidden_UnitType           *string                `protobuf:"bytes,6,opt,name=unit_type,json=unitType"`
	XXX_raceDetectHookData        protoimpl.RaceDetectHookData
	XXX_presence                  [1]uint32
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *GetRecentPurchasesByBranchResponse_Item) Reset() {
	*x = GetRecentPurchasesByBranchResponse_Item{}
	mi := &file_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecentPurchasesByBranchResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentPurchasesByBranchResponse_Item) ProtoMessage() {}

func (x *GetRecentPurchasesByBranchResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetRecentPurchasesByBranchResponse_Item) GetProductId() string {
	if x != nil {
		if x.xxx_hidden_ProductId != nil {
			return *x.xxx_hidden_ProductId
		}
		return ""
	}
	return ""
}

func (x *GetRecentPurchasesByBranchResponse_Item) GetProductSn() string {
	if x != nil {
		if x.xxx_hidden_ProductSn != nil {
			return *x.xxx_hidden_ProductSn
		}
		return ""
	}
	return ""
}

func (x *GetRecentPurchasesByBranchResponse_Item) GetProductName() string {
	if x != nil {
		if x.xxx_hidden_ProductName != nil {
			return *x.xxx_hidden_ProductName
		}
		return ""
	}
	return ""
}

func (x *GetRecentPurchasesByBranchResponse_Item) GetProductDescription() string {
	if x != nil {
		if x.xxx_hidden_ProductDescription != nil {
			return *x.xxx_hidden_ProductDescription
		}
		return ""
	}
	return ""
}

func (x *GetRecentPurchasesByBranchResponse_Item) GetOrderedQuantity() float64 {
	if x != nil {
		return x.xxx_hidden_OrderedQuantity
	}
	return 0
}

func (x *GetRecentPurchasesByBranchResponse_Item) GetUnitType() string {
	if x != nil {
		if x.xxx_hidden_UnitType != nil {
			return *x.xxx_hidden_UnitType
		}
		return ""
	}
	return ""
}

func (x *GetRecentPurchasesByBranchResponse_Item) SetProductId(v string) {
	x.xxx_hidden_ProductId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 6)
}

func (x *GetRecentPurchasesByBranchResponse_Item) SetProductSn(v string) {
	x.xxx_hidden_ProductSn = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 6)
}

func (x *GetRecentPurchasesByBranchResponse_Item) SetProductName(v string) {
	x.xxx_hidden_ProductName = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 6)
}

func (x *GetRecentPurchasesByBranchResponse_Item) SetProductDescription(v string) {
	x.xxx_hidden_ProductDescription = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 6)
}

func (x *GetRecentPurchasesByBranchResponse_Item) SetOrderedQuantity(v float64) {
	x.xxx_hidden_OrderedQuantity = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 6)
}

func (x *GetRecentPurchasesByBranchResponse_Item) SetUnitType(v string) {
	x.xxx_hidden_UnitType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 6)
}

func (x *GetRecentPurchasesByBranchResponse_Item) HasProductId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetRecentPurchasesByBranchResponse_Item) HasProductSn() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *GetRecentPurchasesByBranchResponse_Item) HasProductName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *GetRecentPurchasesByBranchResponse_Item) HasProductDescription() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *GetRecentPurchasesByBranchResponse_Item) HasOrderedQuantity() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *GetRecentPurchasesByBranchResponse_Item) HasUnitType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *GetRecentPurchasesByBranchResponse_Item) ClearProductId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_ProductId = nil
}

func (x *GetRecentPurchasesByBranchResponse_Item) ClearProductSn() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_ProductSn = nil
}

func (x *GetRecentPurchasesByBranchResponse_Item) ClearProductName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_ProductName = nil
}

func (x *GetRecentPurchasesByBranchResponse_Item) ClearProductDescription() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_ProductDescription = nil
}

func (x *GetRecentPurchasesByBranchResponse_Item) ClearOrderedQuantity() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_OrderedQuantity = 0
}

func (x *GetRecentPurchasesByBranchResponse_Item) ClearUnitType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_UnitType = nil
}

type GetRecentPurchasesByBranchResponse_Item_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	ProductId          *string
	ProductSn          *string
	ProductName        *string
	ProductDescription *string
	OrderedQuantity    *float64
	UnitType           *string
}

func (b0 GetRecentPurchasesByBranchResponse_Item_builder) Build() *GetRecentPurchasesByBranchResponse_Item {
	m0 := &GetRecentPurchasesByBranchResponse_Item{}
	b, x := &b0, m0
	_, _ = b, x
	if b.ProductId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 6)
		x.xxx_hidden_ProductId = b.ProductId
	}
	if b.ProductSn != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 6)
		x.xxx_hidden_ProductSn = b.ProductSn
	}
	if b.ProductName != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 6)
		x.xxx_hidden_ProductName = b.ProductName
	}
	if b.ProductDescription != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 6)
		x.xxx_hidden_ProductDescription = b.ProductDescription
	}
	if b.OrderedQuantity != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 6)
		x.xxx_hidden_OrderedQuantity = *b.OrderedQuantity
	}
	if b.UnitType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 6)
		x.xxx_hidden_UnitType = b.UnitType
	}
	return m0
}

var File_customer_proto protoreflect.FileDescriptor

var file_customer_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x4d, 0x0a,
	0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0d,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x57, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22,
	0x64, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x73, 0x42, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xf8, 0x02, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x42, 0x79, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x42, 0x79, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0xe0, 0x01,
	0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x32, 0xcc, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x29, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x42, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x2e,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x03, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var file_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_customer_proto_goTypes = []any{
	(*Branch)(nil),                                  // 0: customer.v1.Branch
	(*BranchSummary)(nil),                           // 1: customer.v1.BranchSummary
	(*GetBranchesForContactRequest)(nil),            // 2: customer.v1.GetBranchesForContactRequest
	(*GetBranchesForContactResponse)(nil),           // 3: customer.v1.GetBranchesForContactResponse
	(*GetBranchRequest)(nil),                        // 4: customer.v1.GetBranchRequest
	(*GetBranchResponse)(nil),                       // 5: customer.v1.GetBranchResponse
	(*GetRecentPurchasesByBranchRequest)(nil),       // 6: customer.v1.GetRecentPurchasesByBranchRequest
	(*GetRecentPurchasesByBranchResponse)(nil),      // 7: customer.v1.GetRecentPurchasesByBranchResponse
	(*GetRecentPurchasesByBranchResponse_Item)(nil), // 8: customer.v1.GetRecentPurchasesByBranchResponse.Item
}
var file_customer_proto_depIdxs = []int32{
	1, // 0: customer.v1.GetBranchesForContactResponse.branches:type_name -> customer.v1.BranchSummary
	0, // 1: customer.v1.GetBranchResponse.branch:type_name -> customer.v1.Branch
	8, // 2: customer.v1.GetRecentPurchasesByBranchResponse.items:type_name -> customer.v1.GetRecentPurchasesByBranchResponse.Item
	2, // 3: customer.v1.CustomerService.GetBranchesForContact:input_type -> customer.v1.GetBranchesForContactRequest
	4, // 4: customer.v1.CustomerService.GetBranch:input_type -> customer.v1.GetBranchRequest
	6, // 5: customer.v1.CustomerService.GetRecentPurchasesByBranch:input_type -> customer.v1.GetRecentPurchasesByBranchRequest
	3, // 6: customer.v1.CustomerService.GetBranchesForContact:output_type -> customer.v1.GetBranchesForContactResponse
	5, // 7: customer.v1.CustomerService.GetBranch:output_type -> customer.v1.GetBranchResponse
	7, // 8: customer.v1.CustomerService.GetRecentPurchasesByBranch:output_type -> customer.v1.GetRecentPurchasesByBranchResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customer_proto_init() }
func file_customer_proto_init() {
	if File_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_customer_proto_rawDesc), len(file_customer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_proto_goTypes,
		DependencyIndexes: file_customer_proto_depIdxs,
		MessageInfos:      file_customer_proto_msgTypes,
	}.Build()
	File_customer_proto = out.File
	file_customer_proto_goTypes = nil
	file_customer_proto_depIdxs = nil
}
