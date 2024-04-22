// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: service/rpc/task/task.proto

package task

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

// 空消息
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{0}
}

// 任务信息结构体
type TaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Type        int64  `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=Text,proto3" json:"Text,omitempty"`
	QuestionUrl string `protobuf:"bytes,6,opt,name=Question_url,json=QuestionUrl,proto3" json:"Question_url,omitempty"`
	AnswerUrl   string `protobuf:"bytes,7,opt,name=Answer_url,json=AnswerUrl,proto3" json:"Answer_url,omitempty"`
	Bonus       int64  `protobuf:"varint,8,opt,name=Bonus,proto3" json:"Bonus,omitempty"`
	State       int64  `protobuf:"varint,9,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *TaskInfo) Reset() {
	*x = TaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfo) ProtoMessage() {}

func (x *TaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfo.ProtoReflect.Descriptor instead.
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TaskInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskInfo) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TaskInfo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TaskInfo) GetQuestionUrl() string {
	if x != nil {
		return x.QuestionUrl
	}
	return ""
}

func (x *TaskInfo) GetAnswerUrl() string {
	if x != nil {
		return x.AnswerUrl
	}
	return ""
}

func (x *TaskInfo) GetBonus() int64 {
	if x != nil {
		return x.Bonus
	}
	return 0
}

func (x *TaskInfo) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

// Rank 结构体
type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank      int64  `protobuf:"varint,1,opt,name=Rank,proto3" json:"Rank,omitempty"`
	Id        int64  `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	AvatarUrl string `protobuf:"bytes,4,opt,name=Avatar_url,json=AvatarUrl,proto3" json:"Avatar_url,omitempty"`
	Point     int64  `protobuf:"varint,5,opt,name=Point,proto3" json:"Point,omitempty"`
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *Rank) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Rank) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rank) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Rank) GetPoint() int64 {
	if x != nil {
		return x.Point
	}
	return 0
}

type GetTaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User_Id int64 `protobuf:"varint,1,opt,name=User_Id,json=UserId,proto3" json:"User_Id,omitempty"`
}

func (x *GetTaskListRequest) Reset() {
	*x = GetTaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskListRequest) ProtoMessage() {}

func (x *GetTaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskListRequest.ProtoReflect.Descriptor instead.
func (*GetTaskListRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskListRequest) GetUser_Id() int64 {
	if x != nil {
		return x.User_Id
	}
	return 0
}

type TaskInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User_Id int64 `protobuf:"varint,1,opt,name=User_Id,json=UserId,proto3" json:"User_Id,omitempty"`
}

func (x *TaskInitRequest) Reset() {
	*x = TaskInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInitRequest) ProtoMessage() {}

func (x *TaskInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInitRequest.ProtoReflect.Descriptor instead.
func (*TaskInitRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskInitRequest) GetUser_Id() int64 {
	if x != nil {
		return x.User_Id
	}
	return 0
}

type TaskListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskList []*TaskInfo `protobuf:"bytes,1,rep,name=TaskList,proto3" json:"TaskList,omitempty"`
}

func (x *TaskListReply) Reset() {
	*x = TaskListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListReply) ProtoMessage() {}

func (x *TaskListReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListReply.ProtoReflect.Descriptor instead.
func (*TaskListReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskListReply) GetTaskList() []*TaskInfo {
	if x != nil {
		return x.TaskList
	}
	return nil
}

type DoTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task_Id int64 `protobuf:"varint,1,opt,name=Task_Id,json=TaskId,proto3" json:"Task_Id,omitempty"`
}

func (x *DoTaskRequest) Reset() {
	*x = DoTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTaskRequest) ProtoMessage() {}

func (x *DoTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTaskRequest.ProtoReflect.Descriptor instead.
func (*DoTaskRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{6}
}

func (x *DoTaskRequest) GetTask_Id() int64 {
	if x != nil {
		return x.Task_Id
	}
	return 0
}

type DoTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int64 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *DoTaskReply) Reset() {
	*x = DoTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTaskReply) ProtoMessage() {}

func (x *DoTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTaskReply.ProtoReflect.Descriptor instead.
func (*DoTaskReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{7}
}

func (x *DoTaskReply) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type TaskInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task_Id int64 `protobuf:"varint,1,opt,name=Task_Id,json=TaskId,proto3" json:"Task_Id,omitempty"`
}

func (x *TaskInfoRequest) Reset() {
	*x = TaskInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfoRequest) ProtoMessage() {}

func (x *TaskInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfoRequest.ProtoReflect.Descriptor instead.
func (*TaskInfoRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{8}
}

func (x *TaskInfoRequest) GetTask_Id() int64 {
	if x != nil {
		return x.Task_Id
	}
	return 0
}

type TaskInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TaskInfo `protobuf:"bytes,1,opt,name=Task,proto3" json:"Task,omitempty"`
}

func (x *TaskInfoReply) Reset() {
	*x = TaskInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfoReply) ProtoMessage() {}

func (x *TaskInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfoReply.ProtoReflect.Descriptor instead.
func (*TaskInfoReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{9}
}

func (x *TaskInfoReply) GetTask() *TaskInfo {
	if x != nil {
		return x.Task
	}
	return nil
}

type RankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank []*Rank `protobuf:"bytes,1,rep,name=Rank,proto3" json:"Rank,omitempty"`
}

func (x *RankRequest) Reset() {
	*x = RankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRequest) ProtoMessage() {}

func (x *RankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRequest.ProtoReflect.Descriptor instead.
func (*RankRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{10}
}

func (x *RankRequest) GetRank() []*Rank {
	if x != nil {
		return x.Rank
	}
	return nil
}

type RankReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank []*Rank `protobuf:"bytes,1,rep,name=Rank,proto3" json:"Rank,omitempty"`
}

func (x *RankReply) Reset() {
	*x = RankReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankReply) ProtoMessage() {}

func (x *RankReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankReply.ProtoReflect.Descriptor instead.
func (*RankReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{11}
}

func (x *RankReply) GetRank() []*Rank {
	if x != nil {
		return x.Rank
	}
	return nil
}

type FailTaskListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailTaskList []*TaskInfo `protobuf:"bytes,1,rep,name=FailTaskList,proto3" json:"FailTaskList,omitempty"`
}

func (x *FailTaskListReply) Reset() {
	*x = FailTaskListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailTaskListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailTaskListReply) ProtoMessage() {}

func (x *FailTaskListReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailTaskListReply.ProtoReflect.Descriptor instead.
func (*FailTaskListReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{12}
}

func (x *FailTaskListReply) GetFailTaskList() []*TaskInfo {
	if x != nil {
		return x.FailTaskList
	}
	return nil
}

type AdminCheckTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task_Id    int64 `protobuf:"varint,1,opt,name=Task_Id,json=TaskId,proto3" json:"Task_Id,omitempty"`
	ActionType int64 `protobuf:"varint,2,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (x *AdminCheckTaskRequest) Reset() {
	*x = AdminCheckTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCheckTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCheckTaskRequest) ProtoMessage() {}

func (x *AdminCheckTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCheckTaskRequest.ProtoReflect.Descriptor instead.
func (*AdminCheckTaskRequest) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{13}
}

func (x *AdminCheckTaskRequest) GetTask_Id() int64 {
	if x != nil {
		return x.Task_Id
	}
	return 0
}

func (x *AdminCheckTaskRequest) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type TaskVisualReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageData []byte `protobuf:"bytes,1,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
}

func (x *TaskVisualReply) Reset() {
	*x = TaskVisualReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_task_task_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskVisualReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskVisualReply) ProtoMessage() {}

func (x *TaskVisualReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_task_task_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskVisualReply.ProtoReflect.Descriptor instead.
func (*TaskVisualReply) Descriptor() ([]byte, []int) {
	return file_service_rpc_task_task_proto_rawDescGZIP(), []int{14}
}

func (x *TaskVisualReply) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

var File_service_rpc_task_task_proto protoreflect.FileDescriptor

var file_service_rpc_task_task_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xde, 0x01, 0x0a,
	0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x73, 0x0a,
	0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x5f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x0d, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a,
	0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0d, 0x44, 0x6f,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x54,
	0x61, 0x73, 0x6b, 0x5f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x54, 0x61, 0x73, 0x6b, 0x5f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x2d, 0x0a, 0x0b, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x52, 0x61, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x22, 0x2b, 0x0a, 0x09, 0x52, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x52, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x22, 0x47, 0x0a, 0x11, 0x46, 0x61, 0x69, 0x6c, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x0c, 0x46,
	0x61, 0x69, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0c, 0x46, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x50, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b,
	0x5f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x30, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x32, 0xb6, 0x03, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3c, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x13, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x04,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x0c, 0x46, 0x61, 0x69, 0x6c,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x46, 0x61, 0x69,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a,
	0x0a, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x0a, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x12, 0x0b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_rpc_task_task_proto_rawDescOnce sync.Once
	file_service_rpc_task_task_proto_rawDescData = file_service_rpc_task_task_proto_rawDesc
)

func file_service_rpc_task_task_proto_rawDescGZIP() []byte {
	file_service_rpc_task_task_proto_rawDescOnce.Do(func() {
		file_service_rpc_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rpc_task_task_proto_rawDescData)
	})
	return file_service_rpc_task_task_proto_rawDescData
}

var file_service_rpc_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_service_rpc_task_task_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: task.Empty
	(*TaskInfo)(nil),              // 1: task.TaskInfo
	(*Rank)(nil),                  // 2: task.Rank
	(*GetTaskListRequest)(nil),    // 3: task.GetTaskListRequest
	(*TaskInitRequest)(nil),       // 4: task.TaskInitRequest
	(*TaskListReply)(nil),         // 5: task.TaskListReply
	(*DoTaskRequest)(nil),         // 6: task.DoTaskRequest
	(*DoTaskReply)(nil),           // 7: task.DoTaskReply
	(*TaskInfoRequest)(nil),       // 8: task.TaskInfoRequest
	(*TaskInfoReply)(nil),         // 9: task.TaskInfoReply
	(*RankRequest)(nil),           // 10: task.RankRequest
	(*RankReply)(nil),             // 11: task.RankReply
	(*FailTaskListReply)(nil),     // 12: task.FailTaskListReply
	(*AdminCheckTaskRequest)(nil), // 13: task.AdminCheckTaskRequest
	(*TaskVisualReply)(nil),       // 14: task.TaskVisualReply
}
var file_service_rpc_task_task_proto_depIdxs = []int32{
	1,  // 0: task.TaskListReply.TaskList:type_name -> task.TaskInfo
	1,  // 1: task.TaskInfoReply.Task:type_name -> task.TaskInfo
	2,  // 2: task.RankRequest.Rank:type_name -> task.Rank
	2,  // 3: task.RankReply.Rank:type_name -> task.Rank
	1,  // 4: task.FailTaskListReply.FailTaskList:type_name -> task.TaskInfo
	3,  // 5: task.Task.GetTaskList:input_type -> task.GetTaskListRequest
	4,  // 6: task.Task.TaskInit:input_type -> task.TaskInitRequest
	6,  // 7: task.Task.DoTask:input_type -> task.DoTaskRequest
	8,  // 8: task.Task.TaskInfo:input_type -> task.TaskInfoRequest
	10, // 9: task.Task.Rank:input_type -> task.RankRequest
	0,  // 10: task.Task.FailTaskList:input_type -> task.Empty
	13, // 11: task.Task.AdminCheckTask:input_type -> task.AdminCheckTaskRequest
	0,  // 12: task.Task.TaskVisual:input_type -> task.Empty
	5,  // 13: task.Task.GetTaskList:output_type -> task.TaskListReply
	5,  // 14: task.Task.TaskInit:output_type -> task.TaskListReply
	7,  // 15: task.Task.DoTask:output_type -> task.DoTaskReply
	9,  // 16: task.Task.TaskInfo:output_type -> task.TaskInfoReply
	11, // 17: task.Task.Rank:output_type -> task.RankReply
	12, // 18: task.Task.FailTaskList:output_type -> task.FailTaskListReply
	0,  // 19: task.Task.AdminCheckTask:output_type -> task.Empty
	14, // 20: task.Task.TaskVisual:output_type -> task.TaskVisualReply
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_service_rpc_task_task_proto_init() }
func file_service_rpc_task_task_proto_init() {
	if File_service_rpc_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_rpc_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_service_rpc_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInfo); i {
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
		file_service_rpc_task_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rank); i {
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
		file_service_rpc_task_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskListRequest); i {
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
		file_service_rpc_task_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInitRequest); i {
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
		file_service_rpc_task_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListReply); i {
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
		file_service_rpc_task_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTaskRequest); i {
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
		file_service_rpc_task_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTaskReply); i {
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
		file_service_rpc_task_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInfoRequest); i {
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
		file_service_rpc_task_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInfoReply); i {
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
		file_service_rpc_task_task_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRequest); i {
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
		file_service_rpc_task_task_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankReply); i {
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
		file_service_rpc_task_task_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailTaskListReply); i {
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
		file_service_rpc_task_task_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCheckTaskRequest); i {
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
		file_service_rpc_task_task_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskVisualReply); i {
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
			RawDescriptor: file_service_rpc_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_rpc_task_task_proto_goTypes,
		DependencyIndexes: file_service_rpc_task_task_proto_depIdxs,
		MessageInfos:      file_service_rpc_task_task_proto_msgTypes,
	}.Build()
	File_service_rpc_task_task_proto = out.File
	file_service_rpc_task_task_proto_rawDesc = nil
	file_service_rpc_task_task_proto_goTypes = nil
	file_service_rpc_task_task_proto_depIdxs = nil
}
