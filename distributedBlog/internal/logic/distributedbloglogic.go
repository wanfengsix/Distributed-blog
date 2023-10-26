package logic

import (
	"context"

	"distributedBlog/internal/svc"
	"distributedBlog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DistributedBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDistributedBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DistributedBlogLogic {
	return &DistributedBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DistributedBlogLogic) DistributedBlog(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
