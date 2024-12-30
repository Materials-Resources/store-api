// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: search.proto

package search

import (
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

// Common product fields extracted into a base message
type ProductBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`                                               // Unique identifier
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                             // Product name
	Description    string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`                               // Product description
	SerialNumber   string `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`         // Renamed from 'sn' for clarity
	ProductGroupId string `protobuf:"bytes,5,opt,name=product_group_id,json=productGroupId,proto3" json:"product_group_id,omitempty"` // ID of the associated product group
}

func (x *ProductBase) Reset() {
	*x = ProductBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBase) ProtoMessage() {}

func (x *ProductBase) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBase.ProtoReflect.Descriptor instead.
func (*ProductBase) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *ProductBase) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ProductBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductBase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductBase) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *ProductBase) GetProductGroupId() string {
	if x != nil {
		return x.ProductGroupId
	}
	return ""
}

// Pagination and Facet Metadata
type PageMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalRecords int32 `protobuf:"varint,1,opt,name=total_records,json=totalRecords,proto3" json:"total_records,omitempty"` // Total results across all pages
	CurrentPage  int32 `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`    // Current page being viewed
	TotalPages   int32 `protobuf:"varint,3,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`       // Total number of pages available
}

func (x *PageMetadata) Reset() {
	*x = PageMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageMetadata) ProtoMessage() {}

func (x *PageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageMetadata.ProtoReflect.Descriptor instead.
func (*PageMetadata) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *PageMetadata) GetTotalRecords() int32 {
	if x != nil {
		return x.TotalRecords
	}
	return 0
}

func (x *PageMetadata) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *PageMetadata) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

// Represents a single aggregation result in a bucket
type BucketAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`    // The name or value of the aggregated item
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // The count or frequency of this item
}

func (x *BucketAggregation) Reset() {
	*x = BucketAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketAggregation) ProtoMessage() {}

func (x *BucketAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketAggregation.ProtoReflect.Descriptor instead.
func (*BucketAggregation) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *BucketAggregation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BucketAggregation) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Contains overall aggregation results for all fields using a map
type Aggregations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldBuckets map[string]*Bucket `protobuf:"bytes,1,rep,name=field_buckets,json=fieldBuckets,proto3" json:"field_buckets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Maps field names to their respective buckets
}

func (x *Aggregations) Reset() {
	*x = Aggregations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aggregations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aggregations) ProtoMessage() {}

func (x *Aggregations) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aggregations.ProtoReflect.Descriptor instead.
func (*Aggregations) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *Aggregations) GetFieldBuckets() map[string]*Bucket {
	if x != nil {
		return x.FieldBuckets
	}
	return nil
}

// Represents a bucket tied to a field, containing multiple aggregations
type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aggregations []*BucketAggregation `protobuf:"bytes,1,rep,name=aggregations,proto3" json:"aggregations,omitempty"` // List of aggregation results for this field
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{4}
}

func (x *Bucket) GetAggregations() []*BucketAggregation {
	if x != nil {
		return x.Aggregations
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{5}
}

func (x *Filter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Product-related definitions
// Refactored to a single message using ProductBase
type ProductDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ProductBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *ProductDocument) Reset() {
	*x = ProductDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDocument) ProtoMessage() {}

func (x *ProductDocument) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDocument.ProtoReflect.Descriptor instead.
func (*ProductDocument) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{6}
}

func (x *ProductDocument) GetBase() *ProductBase {
	if x != nil {
		return x.Base
	}
	return nil
}

// Search result product, based on ProductBase
type ProductResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *ProductBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *ProductResult) Reset() {
	*x = ProductResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResult) ProtoMessage() {}

func (x *ProductResult) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResult.ProtoReflect.Descriptor instead.
func (*ProductResult) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{7}
}

func (x *ProductResult) GetBase() *ProductBase {
	if x != nil {
		return x.Base
	}
	return nil
}

// Request and Response Definitions for RPCs
type SearchProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`                                                                                             // Search query
	Page    int32              `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                                                                                              // Page number (1-indexed)
	SortBy  string             `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`                                                                             // Optional sort parameter
	Filters map[string]*Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Optional filters
}

func (x *SearchProductsRequest) Reset() {
	*x = SearchProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsRequest) ProtoMessage() {}

func (x *SearchProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsRequest.ProtoReflect.Descriptor instead.
func (*SearchProductsRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{8}
}

func (x *SearchProductsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchProductsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchProductsRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *SearchProductsRequest) GetFilters() map[string]*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type SearchProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results      []*ProductResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PageMetadata *PageMetadata    `protobuf:"bytes,2,opt,name=page_metadata,json=pageMetadata,proto3" json:"page_metadata,omitempty"`
	Aggregations *Aggregations    `protobuf:"bytes,3,opt,name=aggregations,proto3" json:"aggregations,omitempty"`
}

func (x *SearchProductsResponse) Reset() {
	*x = SearchProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsResponse) ProtoMessage() {}

func (x *SearchProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsResponse.ProtoReflect.Descriptor instead.
func (*SearchProductsResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{9}
}

func (x *SearchProductsResponse) GetResults() []*ProductResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchProductsResponse) GetPageMetadata() *PageMetadata {
	if x != nil {
		return x.PageMetadata
	}
	return nil
}

func (x *SearchProductsResponse) GetAggregations() *Aggregations {
	if x != nil {
		return x.Aggregations
	}
	return nil
}

type AddProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductDocument `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *AddProductRequest) Reset() {
	*x = AddProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductRequest) ProtoMessage() {}

func (x *AddProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductRequest.ProtoReflect.Descriptor instead.
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{10}
}

func (x *AddProductRequest) GetProducts() []*ProductDocument {
	if x != nil {
		return x.Products
	}
	return nil
}

type AddProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int32 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddProductResponse) Reset() {
	*x = AddProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductResponse) ProtoMessage() {}

func (x *AddProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductResponse.ProtoReflect.Descriptor instead.
func (*AddProductResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{11}
}

func (x *AddProductResponse) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x22, 0xa4, 0x01, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x22, 0x77, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x52, 0x0a, 0x11, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a,
	0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x12, 0x47, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x4d, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc7, 0x01, 0x0a,
	0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x70, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0c, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xb1, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_search_proto_goTypes = []any{
	(*ProductBase)(nil),            // 0: search.v1.ProductBase
	(*PageMetadata)(nil),           // 1: search.v1.PageMetadata
	(*BucketAggregation)(nil),      // 2: search.v1.BucketAggregation
	(*Aggregations)(nil),           // 3: search.v1.Aggregations
	(*Bucket)(nil),                 // 4: search.v1.Bucket
	(*Filter)(nil),                 // 5: search.v1.Filter
	(*ProductDocument)(nil),        // 6: search.v1.ProductDocument
	(*ProductResult)(nil),          // 7: search.v1.ProductResult
	(*SearchProductsRequest)(nil),  // 8: search.v1.SearchProductsRequest
	(*SearchProductsResponse)(nil), // 9: search.v1.SearchProductsResponse
	(*AddProductRequest)(nil),      // 10: search.v1.AddProductRequest
	(*AddProductResponse)(nil),     // 11: search.v1.AddProductResponse
	nil,                            // 12: search.v1.Aggregations.FieldBucketsEntry
	nil,                            // 13: search.v1.SearchProductsRequest.FiltersEntry
}
var file_search_proto_depIdxs = []int32{
	12, // 0: search.v1.Aggregations.field_buckets:type_name -> search.v1.Aggregations.FieldBucketsEntry
	2,  // 1: search.v1.Bucket.aggregations:type_name -> search.v1.BucketAggregation
	0,  // 2: search.v1.ProductDocument.base:type_name -> search.v1.ProductBase
	0,  // 3: search.v1.ProductResult.base:type_name -> search.v1.ProductBase
	13, // 4: search.v1.SearchProductsRequest.filters:type_name -> search.v1.SearchProductsRequest.FiltersEntry
	7,  // 5: search.v1.SearchProductsResponse.results:type_name -> search.v1.ProductResult
	1,  // 6: search.v1.SearchProductsResponse.page_metadata:type_name -> search.v1.PageMetadata
	3,  // 7: search.v1.SearchProductsResponse.aggregations:type_name -> search.v1.Aggregations
	6,  // 8: search.v1.AddProductRequest.products:type_name -> search.v1.ProductDocument
	4,  // 9: search.v1.Aggregations.FieldBucketsEntry.value:type_name -> search.v1.Bucket
	5,  // 10: search.v1.SearchProductsRequest.FiltersEntry.value:type_name -> search.v1.Filter
	8,  // 11: search.v1.SearchService.SearchProducts:input_type -> search.v1.SearchProductsRequest
	10, // 12: search.v1.SearchService.AddProduct:input_type -> search.v1.AddProductRequest
	9,  // 13: search.v1.SearchService.SearchProducts:output_type -> search.v1.SearchProductsResponse
	11, // 14: search.v1.SearchService.AddProduct:output_type -> search.v1.AddProductResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProductBase); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PageMetadata); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BucketAggregation); i {
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
		file_search_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Aggregations); i {
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
		file_search_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Bucket); i {
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
		file_search_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Filter); i {
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
		file_search_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ProductDocument); i {
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
		file_search_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ProductResult); i {
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
		file_search_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SearchProductsRequest); i {
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
		file_search_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SearchProductsResponse); i {
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
		file_search_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*AddProductRequest); i {
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
		file_search_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*AddProductResponse); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
