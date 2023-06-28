// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: schedulerpb.proto

package schedulerpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId *string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3,oneof" json:"customer_id,omitempty"`
	CompanyId  *string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3,oneof" json:"company_id,omitempty"`
	ServiceId  *string `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3,oneof" json:"service_id,omitempty"`
	EmployeeId *string `protobuf:"bytes,4,opt,name=employee_id,json=employeeId,proto3,oneof" json:"employee_id,omitempty"`
	StartTime  *int64  `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	EndTime    *int64  `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
}

func (x *AddOrderRequest) Reset() {
	*x = AddOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRequest) ProtoMessage() {}

func (x *AddOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRequest.ProtoReflect.Descriptor instead.
func (*AddOrderRequest) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrderRequest) GetCustomerId() string {
	if x != nil && x.CustomerId != nil {
		return *x.CustomerId
	}
	return ""
}

func (x *AddOrderRequest) GetCompanyId() string {
	if x != nil && x.CompanyId != nil {
		return *x.CompanyId
	}
	return ""
}

func (x *AddOrderRequest) GetServiceId() string {
	if x != nil && x.ServiceId != nil {
		return *x.ServiceId
	}
	return ""
}

func (x *AddOrderRequest) GetEmployeeId() string {
	if x != nil && x.EmployeeId != nil {
		return *x.EmployeeId
	}
	return ""
}

func (x *AddOrderRequest) GetStartTime() int64 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *AddOrderRequest) GetEndTime() int64 {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return 0
}

type OrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId *string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3,oneof" json:"customer_id,omitempty"`
	CompanyId  *string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3,oneof" json:"company_id,omitempty"`
	IsCanceled *bool   `protobuf:"varint,3,opt,name=is_canceled,json=isCanceled,proto3,oneof" json:"is_canceled,omitempty"`
	NPerPage   *int64  `protobuf:"varint,4,opt,name=n_per_page,json=nPerPage,proto3,oneof" json:"n_per_page,omitempty"`
	StartValue *string `protobuf:"bytes,5,opt,name=start_value,json=startValue,proto3,oneof" json:"start_value,omitempty"`
	StartDate  *int64  `protobuf:"varint,6,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
}

func (x *OrdersRequest) Reset() {
	*x = OrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersRequest) ProtoMessage() {}

func (x *OrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersRequest.ProtoReflect.Descriptor instead.
func (*OrdersRequest) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{1}
}

func (x *OrdersRequest) GetCustomerId() string {
	if x != nil && x.CustomerId != nil {
		return *x.CustomerId
	}
	return ""
}

func (x *OrdersRequest) GetCompanyId() string {
	if x != nil && x.CompanyId != nil {
		return *x.CompanyId
	}
	return ""
}

func (x *OrdersRequest) GetIsCanceled() bool {
	if x != nil && x.IsCanceled != nil {
		return *x.IsCanceled
	}
	return false
}

func (x *OrdersRequest) GetNPerPage() int64 {
	if x != nil && x.NPerPage != nil {
		return *x.NPerPage
	}
	return 0
}

func (x *OrdersRequest) GetStartValue() string {
	if x != nil && x.StartValue != nil {
		return *x.StartValue
	}
	return ""
}

func (x *OrdersRequest) GetStartDate() int64 {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CustomerId *string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3,oneof" json:"customer_id,omitempty"`
	CompanyId  *string `protobuf:"bytes,3,opt,name=company_id,json=companyId,proto3,oneof" json:"company_id,omitempty"`
	ServiceId  *string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3,oneof" json:"service_id,omitempty"`
	EmployeeId *string `protobuf:"bytes,5,opt,name=employee_id,json=employeeId,proto3,oneof" json:"employee_id,omitempty"`
	OrderTime  *int64  `protobuf:"varint,6,opt,name=order_time,json=orderTime,proto3,oneof" json:"order_time,omitempty"`
	IsCanceled *bool   `protobuf:"varint,7,opt,name=is_canceled,json=isCanceled,proto3,oneof" json:"is_canceled,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Order) GetCustomerId() string {
	if x != nil && x.CustomerId != nil {
		return *x.CustomerId
	}
	return ""
}

func (x *Order) GetCompanyId() string {
	if x != nil && x.CompanyId != nil {
		return *x.CompanyId
	}
	return ""
}

func (x *Order) GetServiceId() string {
	if x != nil && x.ServiceId != nil {
		return *x.ServiceId
	}
	return ""
}

func (x *Order) GetEmployeeId() string {
	if x != nil && x.EmployeeId != nil {
		return *x.EmployeeId
	}
	return ""
}

func (x *Order) GetOrderTime() int64 {
	if x != nil && x.OrderTime != nil {
		return *x.OrderTime
	}
	return 0
}

func (x *Order) GetIsCanceled() bool {
	if x != nil && x.IsCanceled != nil {
		return *x.IsCanceled
	}
	return false
}

type OrdersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrdersReply) Reset() {
	*x = OrdersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersReply) ProtoMessage() {}

func (x *OrdersReply) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersReply.ProtoReflect.Descriptor instead.
func (*OrdersReply) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{3}
}

func (x *OrdersReply) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CustomerId *string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3,oneof" json:"customer_id,omitempty"`
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{4}
}

func (x *CancelRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *CancelRequest) GetCustomerId() string {
	if x != nil && x.CustomerId != nil {
		return *x.CustomerId
	}
	return ""
}

type AvaliableTimeSlotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId       *string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3,oneof" json:"company_id,omitempty"`
	ServiceId       *string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3,oneof" json:"service_id,omitempty"`
	ServiceDuration *int32  `protobuf:"varint,3,opt,name=service_duration,json=serviceDuration,proto3,oneof" json:"service_duration,omitempty"`
	Date            *int64  `protobuf:"varint,4,opt,name=date,proto3,oneof" json:"date,omitempty"`
	StartValue      *string `protobuf:"bytes,5,opt,name=start_value,json=startValue,proto3,oneof" json:"start_value,omitempty"`
	NPerPage        *int64  `protobuf:"varint,6,opt,name=n_per_page,json=nPerPage,proto3,oneof" json:"n_per_page,omitempty"`
}

func (x *AvaliableTimeSlotsRequest) Reset() {
	*x = AvaliableTimeSlotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvaliableTimeSlotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvaliableTimeSlotsRequest) ProtoMessage() {}

func (x *AvaliableTimeSlotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvaliableTimeSlotsRequest.ProtoReflect.Descriptor instead.
func (*AvaliableTimeSlotsRequest) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{5}
}

func (x *AvaliableTimeSlotsRequest) GetCompanyId() string {
	if x != nil && x.CompanyId != nil {
		return *x.CompanyId
	}
	return ""
}

func (x *AvaliableTimeSlotsRequest) GetServiceId() string {
	if x != nil && x.ServiceId != nil {
		return *x.ServiceId
	}
	return ""
}

func (x *AvaliableTimeSlotsRequest) GetServiceDuration() int32 {
	if x != nil && x.ServiceDuration != nil {
		return *x.ServiceDuration
	}
	return 0
}

func (x *AvaliableTimeSlotsRequest) GetDate() int64 {
	if x != nil && x.Date != nil {
		return *x.Date
	}
	return 0
}

func (x *AvaliableTimeSlotsRequest) GetStartValue() string {
	if x != nil && x.StartValue != nil {
		return *x.StartValue
	}
	return ""
}

func (x *AvaliableTimeSlotsRequest) GetNPerPage() int64 {
	if x != nil && x.NPerPage != nil {
		return *x.NPerPage
	}
	return 0
}

type TimeSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *int64 `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	EndTime   *int64 `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
}

func (x *TimeSlot) Reset() {
	*x = TimeSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSlot) ProtoMessage() {}

func (x *TimeSlot) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSlot.ProtoReflect.Descriptor instead.
func (*TimeSlot) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{6}
}

func (x *TimeSlot) GetStartTime() int64 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *TimeSlot) GetEndTime() int64 {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return 0
}

type EmployeeTimeSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *string     `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name      *string     `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Surname   *string     `protobuf:"bytes,3,opt,name=surname,proto3,oneof" json:"surname,omitempty"`
	TimeSlots []*TimeSlot `protobuf:"bytes,4,rep,name=time_slots,json=timeSlots,proto3" json:"time_slots,omitempty"`
}

func (x *EmployeeTimeSlot) Reset() {
	*x = EmployeeTimeSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeTimeSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeTimeSlot) ProtoMessage() {}

func (x *EmployeeTimeSlot) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeTimeSlot.ProtoReflect.Descriptor instead.
func (*EmployeeTimeSlot) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{7}
}

func (x *EmployeeTimeSlot) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *EmployeeTimeSlot) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *EmployeeTimeSlot) GetSurname() string {
	if x != nil && x.Surname != nil {
		return *x.Surname
	}
	return ""
}

func (x *EmployeeTimeSlot) GetTimeSlots() []*TimeSlot {
	if x != nil {
		return x.TimeSlots
	}
	return nil
}

type AvaliableTimeSlotsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeTimeSlots []*EmployeeTimeSlot `protobuf:"bytes,1,rep,name=employee_time_slots,json=employeeTimeSlots,proto3" json:"employee_time_slots,omitempty"`
}

func (x *AvaliableTimeSlotsReply) Reset() {
	*x = AvaliableTimeSlotsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedulerpb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvaliableTimeSlotsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvaliableTimeSlotsReply) ProtoMessage() {}

func (x *AvaliableTimeSlotsReply) ProtoReflect() protoreflect.Message {
	mi := &file_schedulerpb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvaliableTimeSlotsReply.ProtoReflect.Descriptor instead.
func (*AvaliableTimeSlotsReply) Descriptor() ([]byte, []int) {
	return file_schedulerpb_proto_rawDescGZIP(), []int{8}
}

func (x *AvaliableTimeSlotsReply) GetEmployeeTimeSlots() []*EmployeeTimeSlot {
	if x != nil {
		return x.EmployeeTimeSlots
	}
	return nil
}

var File_schedulerpb_proto protoreflect.FileDescriptor

var file_schedulerpb_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0xc9, 0x02, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0a, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x08, 0x6e, 0x50, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22,
	0xde, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x0a, 0x69,
	0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x22, 0x39, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2a, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x61, 0x0a, 0x0d, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0xd0,
	0x02, 0x0a, 0x19, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0a, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x08, 0x6e, 0x50, 0x65, 0x72, 0x50,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x6a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x22, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xb1, 0x01,
	0x0a, 0x10, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c,
	0x6f, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x34, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x53, 0x6c, 0x6f, 0x74, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x68, 0x0a, 0x17, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x13,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x11, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x32, 0xc6, 0x02, 0x0a, 0x03,
	0x41, 0x70, 0x69, 0x12, 0x42, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x6e, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61,
	0x6e, 0x79, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x6c, 0x6f, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x61, 0x6c, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x73, 0x69, 0x6b, 0x2d, 0x34, 0x30, 0x34, 0x2f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedulerpb_proto_rawDescOnce sync.Once
	file_schedulerpb_proto_rawDescData = file_schedulerpb_proto_rawDesc
)

func file_schedulerpb_proto_rawDescGZIP() []byte {
	file_schedulerpb_proto_rawDescOnce.Do(func() {
		file_schedulerpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedulerpb_proto_rawDescData)
	})
	return file_schedulerpb_proto_rawDescData
}

var file_schedulerpb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schedulerpb_proto_goTypes = []interface{}{
	(*AddOrderRequest)(nil),           // 0: schedulerpb.AddOrderRequest
	(*OrdersRequest)(nil),             // 1: schedulerpb.OrdersRequest
	(*Order)(nil),                     // 2: schedulerpb.Order
	(*OrdersReply)(nil),               // 3: schedulerpb.OrdersReply
	(*CancelRequest)(nil),             // 4: schedulerpb.CancelRequest
	(*AvaliableTimeSlotsRequest)(nil), // 5: schedulerpb.AvaliableTimeSlotsRequest
	(*TimeSlot)(nil),                  // 6: schedulerpb.TimeSlot
	(*EmployeeTimeSlot)(nil),          // 7: schedulerpb.EmployeeTimeSlot
	(*AvaliableTimeSlotsReply)(nil),   // 8: schedulerpb.AvaliableTimeSlotsReply
	(*emptypb.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_schedulerpb_proto_depIdxs = []int32{
	2, // 0: schedulerpb.OrdersReply.orders:type_name -> schedulerpb.Order
	6, // 1: schedulerpb.EmployeeTimeSlot.time_slots:type_name -> schedulerpb.TimeSlot
	7, // 2: schedulerpb.AvaliableTimeSlotsReply.employee_time_slots:type_name -> schedulerpb.EmployeeTimeSlot
	0, // 3: schedulerpb.Api.AddOrder:input_type -> schedulerpb.AddOrderRequest
	1, // 4: schedulerpb.Api.FindManyOrders:input_type -> schedulerpb.OrdersRequest
	4, // 5: schedulerpb.Api.CancelOrder:input_type -> schedulerpb.CancelRequest
	5, // 6: schedulerpb.Api.FindManyAvaliableTimeSlots:input_type -> schedulerpb.AvaliableTimeSlotsRequest
	9, // 7: schedulerpb.Api.AddOrder:output_type -> google.protobuf.Empty
	3, // 8: schedulerpb.Api.FindManyOrders:output_type -> schedulerpb.OrdersReply
	9, // 9: schedulerpb.Api.CancelOrder:output_type -> google.protobuf.Empty
	8, // 10: schedulerpb.Api.FindManyAvaliableTimeSlots:output_type -> schedulerpb.AvaliableTimeSlotsReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_schedulerpb_proto_init() }
func file_schedulerpb_proto_init() {
	if File_schedulerpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schedulerpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderRequest); i {
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
		file_schedulerpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersRequest); i {
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
		file_schedulerpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_schedulerpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersReply); i {
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
		file_schedulerpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRequest); i {
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
		file_schedulerpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvaliableTimeSlotsRequest); i {
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
		file_schedulerpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSlot); i {
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
		file_schedulerpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeTimeSlot); i {
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
		file_schedulerpb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvaliableTimeSlotsReply); i {
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
	file_schedulerpb_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_schedulerpb_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_schedulerpb_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_schedulerpb_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_schedulerpb_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_schedulerpb_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_schedulerpb_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schedulerpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schedulerpb_proto_goTypes,
		DependencyIndexes: file_schedulerpb_proto_depIdxs,
		MessageInfos:      file_schedulerpb_proto_msgTypes,
	}.Build()
	File_schedulerpb_proto = out.File
	file_schedulerpb_proto_rawDesc = nil
	file_schedulerpb_proto_goTypes = nil
	file_schedulerpb_proto_depIdxs = nil
}