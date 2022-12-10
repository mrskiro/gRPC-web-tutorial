package main

import (
	"context"
	"log"
	"net/http"
	v1 "server/gen/proto/v1"
	"server/gen/proto/v1/v1connect"

	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"
)

type notificationHandler struct{}

func (h notificationHandler) SendNotification(
	ctx context.Context,
	req *connect.Request[v1.SendNotificationRequest],
) (*connect.Response[emptypb.Empty], error) {
	log.Println(req.Msg.GetUserId(), req.Msg.GetType(), req.Msg.Body)
	res := connect.NewResponse[emptypb.Empty](nil)
	return res, nil
}

func (h notificationHandler) NotificationList(
	ctx context.Context,
	req *connect.Request[v1.NotificationListRequest],
) (*connect.Response[v1.NotificationListResponse], error) {
	res := connect.NewResponse(&v1.NotificationListResponse{
		Notifications: []*v1.Notification{
			{
				Id:   1,
				Type: v1.NotificationType_NOTIFICATION_TYPE_EMAIL,
				Body: "Emailの通知です",
			},
		},
	})
	return res, nil
}

func main() {
	notificationHandler := notificationHandler{}
	path, handler := v1connect.NewNotificationServiceHandler(notificationHandler)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	handler = cors.Default().Handler(h2c.NewHandler(mux, &http2.Server{}))
	http.ListenAndServe(
		"localhost:8080",
		handler,
	)
}
