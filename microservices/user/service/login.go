package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	//model "github.com/Yux77Yux/platform_backend/microservices/user/model"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	var (
		res *generated.LoginResponse
		err error
	)

	// user_credentials := &model.UserCredentials{
	// 	Username: req.GetUserCredential().Username,
	// 	Password: req.GetUserCredential().Password,
	// }

	return res, err
}
