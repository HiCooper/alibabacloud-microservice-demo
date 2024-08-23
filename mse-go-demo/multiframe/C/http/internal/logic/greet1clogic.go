package logic

import (
	"context"
	"fmt"
	"os"

	"github.com/aliyun/alibabacloud-microservice-demo/mse-go-demo/multiframe/C/http/internal/svc"
	"github.com/aliyun/alibabacloud-microservice-demo/mse-go-demo/multiframe/C/http/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type Greet1cLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreet1cLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Greet1cLogic {
	return &Greet1cLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Greet1cLogic) Greet1c(req *types.GreetReq) (resp *types.GreetResp, err error) {
	fmt.Printf("[Greet1cLogic.Greet1c] receive request: %v\n", req)

	resp = &types.GreetResp{
		Message: fmt.Sprintf("hello %s", req.Name),
	}

	tag := os.Getenv("MSE_ALICLOUD_SERVICE_TAG")
	if tag == "" {
		tag = "base"
	}
	ip := os.Getenv("KUBERNETES_POD_IP")
	resp.CallChain = fmt.Sprintf("C:%s:%s", tag, ip)
	fmt.Printf("[Greet1cLogic.Greet1c] response: %v\n", resp)
	return
}