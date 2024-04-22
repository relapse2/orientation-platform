// Code generated by goctl. DO NOT EDIT.
// Source: task.proto

package server

import (
	"context"

	"orientation-platform/service/rpc/task/internal/logic"
	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"
)

type TaskServer struct {
	svcCtx *svc.ServiceContext
	task.UnimplementedTaskServer
}

func NewTaskServer(svcCtx *svc.ServiceContext) *TaskServer {
	return &TaskServer{
		svcCtx: svcCtx,
	}
}

func (s *TaskServer) GetTaskList(ctx context.Context, in *task.GetTaskListRequest) (*task.TaskListReply, error) {
	l := logic.NewGetTaskListLogic(ctx, s.svcCtx)
	return l.GetTaskList(in)
}

func (s *TaskServer) TaskInit(ctx context.Context, in *task.TaskInitRequest) (*task.TaskListReply, error) {
	l := logic.NewTaskInitLogic(ctx, s.svcCtx)
	return l.TaskInit(in)
}

func (s *TaskServer) DoTask(ctx context.Context, in *task.DoTaskRequest) (*task.DoTaskReply, error) {
	l := logic.NewDoTaskLogic(ctx, s.svcCtx)
	return l.DoTask(in)
}

func (s *TaskServer) TaskInfo(ctx context.Context, in *task.TaskInfoRequest) (*task.TaskInfoReply, error) {
	l := logic.NewTaskInfoLogic(ctx, s.svcCtx)
	return l.TaskInfo(in)
}

func (s *TaskServer) Rank(ctx context.Context, in *task.RankRequest) (*task.RankReply, error) {
	l := logic.NewRankLogic(ctx, s.svcCtx)
	return l.Rank(in)
}

func (s *TaskServer) FailTaskList(ctx context.Context, in *task.Empty) (*task.FailTaskListReply, error) {
	l := logic.NewFailTaskListLogic(ctx, s.svcCtx)
	return l.FailTaskList(in)
}

func (s *TaskServer) AdminCheckTask(ctx context.Context, in *task.AdminCheckTaskRequest) (*task.Empty, error) {
	l := logic.NewAdminCheckTaskLogic(ctx, s.svcCtx)
	return l.AdminCheckTask(in)
}

func (s *TaskServer) TaskVisual(ctx context.Context, in *task.Empty) (*task.TaskVisualReply, error) {
	l := logic.NewTaskVisualLogic(ctx, s.svcCtx)
	return l.TaskVisual(in)
}
