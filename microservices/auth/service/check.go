package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	//common "github.com/Yux77Yux/platform_backend/generated/common"
)

// 原本用于自定义实现OAuth的Check，即envoy默认调用的，没看到文档，不做了
var ()

func (s *Server) Check(ctx context.Context, req *generated.CheckRequest) (*generated.CheckResponse, error) {
	return nil, nil
}
