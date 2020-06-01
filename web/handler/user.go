package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2/service/grpc"
	"net/http"
	"time"

	proto "github.com/wxmsummer/shopping/user/proto/user"
)

func Register(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())
	user := proto.User{
		Name: request["name"].(string),
		Phone: request["phone"].(string),
		Password: request["password"].(string),
	}
	rsp, err := webClient.Register(context.TODO(), &proto.RegisterReq{
		User: &user,
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

func Login(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	phone := "12345"
	password := "12345"

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())
	rsp, err := webClient.Login(context.TODO(), &proto.LoginReq{
		Phone: phone,
		Password: password,
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

func Logout(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id := 1

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())
	rsp, err := webClient.Logout(context.TODO(), &proto.LogoutReq{
		Id: int32(id),
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

func GetLevel(w http.ResponseWriter, r *http.Request) {

	server := grpc.NewService()
	server.Init()

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id := 1

	// call the backend service
	webClient := proto.NewUserService("go.micro.service.user", server.Client())
	rsp, err := webClient.GetLevel(context.TODO(), &proto.GetLevelReq{
		Id: int32(id),
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