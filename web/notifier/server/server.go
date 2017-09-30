package server

import (

	//"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Email(ctx context.Context, in *notifier.EmailRequest) (*notifier.EmailResponse, error) {

	err := shared.SendEmail(in.Email, "Welcome To GoLand","htttps://github.com/wizelineacademy/GoWorkshop")

  if err == nil {
		return &notifier.EmailResponse {
			Code: 200,
			Message: "OK",
		},nil
	}

	return &notifier.EmailResponse {
		Code: 500,
		Message: err.Error(),
	}, err

}
