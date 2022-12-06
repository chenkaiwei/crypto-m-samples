package logic

import (
	"context"

	"standardDemo/internal/svc"
	"standardDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiDemo2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMultiDemo2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiDemo2Logic {
	return &MultiDemo2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MultiDemo2Logic) MultiDemo2(req *types.SimpleMsg) (resp *types.SimpleMsg, err error) {
	logx.Info("MultiDemo2 called")

	resp = &types.SimpleMsg{
		Message: "your message is revieved --" + req.Message,
	}
	return
}
