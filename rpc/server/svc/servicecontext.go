package svc

import (
	"front-gin/config/etc"
	"front-gin/rpc/article/article"
	"front-gin/rpc/classify/classify"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   etc.Config
	Article  article.Article
	Classify classify.Classify
	//Base     base.Base
}

func NewServiceContext(c etc.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Article:  article.NewArticle(zrpc.MustNewClient(c.ArticleRpc)),
		Classify: classify.NewClassify(zrpc.MustNewClient(c.ClassifyRpc)),
		//Base:     base.NewBase(zrpc.MustNewClient(c.BaseRpc)),
	}
}