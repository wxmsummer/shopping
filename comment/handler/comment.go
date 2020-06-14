package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"shopping/comment/model"
	"shopping/comment/repository"

	proto "shopping/comment/proto/comment"
)

type Comment struct{ Repo *repository.Comment }

// 对已完成的该笔订单可进行评价
func (e *Comment) AddComment(ctx context.Context, req *proto.AddCommentReq, rsp *proto.Resp) (err error) {

	comment := model.Comment{
		UserID:     req.Comment.Id,
		ProductID:  req.Comment.ProductID,
		OrderID:    req.Comment.OrderId,
		Star:       req.Comment.Star,
		Content:    req.Comment.Content,
		CreateTime: req.Comment.CreateTime,
	}
	// 注册账号到数据库
	err = e.Repo.Create(&comment)
	if err != nil {
		log.Error("e.Repo.Create(&comment) err")
		return err
	}
	rsp.Code = 200
	rsp.Msg = "AddComment成功！"
	return nil
}

// 根据订单号获取该订单的评价
func (e *Comment) GetCommentByOrderId(ctx context.Context, req *proto.GetCommentsReq, rsp *proto.GetCommentsResp) error {

	rsp.Code = 200
	rsp.Msg = "GetComments成功！"
	return nil
}

// 根据商品id获取该商品的所有评价
func (e *Comment) GetCommentsByProductId(ctx context.Context, req *proto.GetCommentsReq, rsp *proto.GetCommentsResp) error {

	rsp.Code = 200
	rsp.Msg = "GetComments成功！"
	return nil
}


