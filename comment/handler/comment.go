package handler

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"shopping/comment/model"
	"shopping/comment/repository"

	proto "shopping/comment/proto/comment"
)

type Comment struct{ Repo *repository.Comment }

// 如何解决评价服务的评价和商品服务的评价不同步的问题？
// 后端服务调用接口
// 可以周期性更新评价

// 对已完成的该笔订单可进行评价
func (e *Comment) AddComment(ctx context.Context, req *proto.AddCommentReq, rsp *proto.Resp) (err error) {

	comment := model.Comment{
		UserID:     req.Comment.UserID,
		ProductID:  req.Comment.ProductID,
		OrderID:    req.Comment.OrderID,
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
func (e *Comment) GetCommentByOrderId(ctx context.Context, req *proto.GetCommentByOrderIdReq, rsp *proto.GetCommentByOrderIdResp) error {

	var comment proto.Comment
	err := e.Repo.Db.Where("order_id = ?" , req.OrderID).Find(&comment).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = "GetCommentByOrderId success"
	rsp.Comment = &comment

	return nil
}

// 根据商品id获取该商品的所有评价
func (e *Comment) GetCommentsByProductId(ctx context.Context, req *proto.GetCommentsByProductIdReq, rsp *proto.GetCommentsByProductIdResp) error {

	var comments []*proto.Comment

	err := e.Repo.Db.Where("product_id = ?" , req.ProductID).Find(&comments).Error
	if err != nil {
		return err
	}

	rsp.Code = 200
	rsp.Msg = "GetCommentsByProductId success"
	rsp.Comments = comments

	return nil
}


