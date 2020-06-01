package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2/service/grpc"
	"net/http"
	"time"

	proto "github.com/wxmsummer/shopping/comment/proto/comment"
)

func GetComments(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()
	
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	webClient := proto.NewCommentService("go.micro.service.comment", server.Client())
	rsp, err := webClient.GetComments(context.TODO(), &proto.GetCommentsReq{
		ProductID: request["productID"].(int32),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func AddComment(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()
	
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 实例化评论
	comment := &proto.Comment{
		Id:         0,
		UserID:     0,
		ProductID:  0,
		Content:    "",
		CreateTime: "",
	}

	// call the backend service
	webClient := proto.NewCommentService("go.micro.service.comment", server.Client())
	rsp, err := webClient.AddComment(context.TODO(), &proto.AddCommentReq{
		Comment: comment,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}