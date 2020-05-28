package handler
//
//import (
//	"context"
//	"encoding/json"
//	"net/http"
//	"time"
//
//	"github.com/micro/go-micro/v2/client"
//	proto "github.com/wxmsummer/shopping/payment/proto/payment"
//)
//
//func WebCall(w http.ResponseWriter, r *http.Request) {
//	// decode the incoming request as json
//	var request map[string]interface{}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// call the backend service
//	webClient := proto.NewWebService("go.micro.service.proto", client.DefaultClient)
//	rsp, err := webClient.Call(context.TODO(), &proto.Request{
//		Name: request["name"].(string),
//	})
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// we want to augment the response
//	response := map[string]interface{}{
//		"msg": rsp.Msg,
//		"ref": time.Now().UnixNano(),
//	}
//
//	// encode and write the response as json
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//}
