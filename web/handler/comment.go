package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	proto "github.com/wxmsummer/shopping/comment/proto/comment"
	"net/http"
	"strconv"
	"time"
)

func GetComments(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	webClient := proto.NewCommentService("go.micro.service.comment", server.Client())
	rsp, err := webClient.GetCommentsByProductId(context.TODO(), &proto.GetCommentsByProductIdReq{
		ProductID: c.Param("product_id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)
}

// 获取发表评价界面
func GetAddComment(c *gin.Context) {
	c.HTML(http.StatusOK, "add_comment.html", nil)
}

// 提交表单，发表评价
func PostAddComment(c *gin.Context) {

	server := grpc.NewService()
	server.Init()

	userId, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(302, "/user/login")
	}

	star, _ := strconv.Atoi(c.PostForm("star"))

	// 从网页获取表单数据，实例化评论
	comment := &proto.Comment{
		OrderID:    c.PostForm("order_id"),
		UserID:     userId,
		ProductID:  c.PostForm("product_id"),
		Star:       int32(star),
		Content:    c.PostForm("content"),
		CreateTime: time.Now().Unix(),
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
