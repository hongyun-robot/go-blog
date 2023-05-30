package list

import (
	"context"
	"fmt"
	"front-gin/rpc/article/article_client"
	"front-gin/rpc/server/svc"
	"front-gin/src/global"
	"front-gin/src/types"
	"front-gin/src/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type GetArticleByPagingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByPagingLogic {
	return &GetArticleByPagingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByPagingLogic) GetArticleByPaging(req *types.GetArticleByPagingReq) (resp *types.ArticleRespose, err error) {
	fmt.Println("draft", req.PageSize)
	articleList, err := l.svcCtx.Article.GetArticleByPaging(l.ctx, &article_client.GetArticleByPagingRequest{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Draft:    req.Draft,
	})
	fmt.Println("qqqqqqqqq", articleList)
	if err != nil {
		return &types.ArticleRespose{
			Code:    http.StatusOK,
			Data:    []*types.ArticleData{},
			Message: global.ERROR,
		}, nil
	}

	articleData, err := utils.ArticleRpcToApi(articleList.GetArticleData())
	if err != nil {
		return &types.ArticleRespose{
			Code:    http.StatusOK,
			Data:    []*types.ArticleData{},
			Message: global.ERROR,
		}, nil
	}
	return &types.ArticleRespose{
		Code:    http.StatusOK,
		Data:    articleData,
		Message: global.SUCCESS,
	}, nil
}