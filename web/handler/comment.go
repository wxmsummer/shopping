package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	proto "github.com/wxmsummer/shopping/comment/proto/comment"
	"net/http"
)

func GetComments(c *gin.Context) {

	//server := grpc.NewService()
	//server.Init()
	//
	//// call the backend service
	//webClient := proto.NewCommentService("go.micro.service.comment", server.Client())
	//
	//productID, _ := strconv.Atoi(c.Param("productID"))
	//
	//rsp, err := webClient.GetComments(context.TODO(), &proto.GetCommentsReq{
	//	ProductID: int32(productID),
	//})
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, rsp)
	//	return
	//}
	//
	//c.JSON(http.StatusOK, rsp)
}

// 获取发表评价界面
func GetAddComment(c *gin.Context) {
	c.HTML(http.StatusOK, "add_comment.html", nil)
}

// 提交表单，发表评价
func PostAddComment(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	// 从网页获取表单数据，实例化评论
	comment := &proto.Comment{
		OrderID:    0,
		UserID:     0,
		ProductID:  0,
		Content:    c.PostForm("content"),
		CreateTime: "",
	}

	// call the backend service
	webClient := proto.NewCommentService("go.micro.service.comment", server.Client())
	rsp, err := webClient.AddComment(context.TODO(), &proto.AddCommentReq{
		Comment: comment,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)
}
