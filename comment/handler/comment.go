package handler

import (
	"context"
	"shopping/comment/repository"

	proto "shopping/comment/proto/comment"
)

type Comment struct{ Repo *repository.Comment }

func (e *Comment) GetComments(ctx context.Context, req *proto.GetCommentsReq, rsp *proto.GetCommentsResp) error {

	rsp.Code = 200
	rsp.Msg = "GetComments成功！"
	return nil
}

func (e *Comment) AddComment(ctx context.Context, req *proto.AddCommentReq, rsp *proto.Resp) error {

	rsp.Code = 200
	rsp.Msg = "AddComment成功！"
	return nil
}
