package article

import (
	"context"
	"front-gin/rpc/article/article_client"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddArticleRequest         = article_client.AddArticleRequest
	ArticleData               = article_client.ArticleData
	ArticleList               = article_client.ArticleList
	ClassifyData              = article_client.ClassifyData
	DeleteArticleRequest      = article_client.DeleteArticleRequest
	GetArticleByPagingRequest = article_client.GetArticleByPagingRequest
	GetArticleRequest         = article_client.GetArticleRequest
	UpdateArticleRequest      = article_client.UpdateArticleRequest

	Article interface {
		AddArticle(ctx context.Context, in *AddArticleRequest, opts ...grpc.CallOption) (*ArticleList, error)
		GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*ArticleList, error)
		GetArticleByPaging(ctx context.Context, in *GetArticleByPagingRequest, opts ...grpc.CallOption) (*ArticleList, error)
		UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*ArticleList, error)
		DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*ArticleList, error)
	}

	defaultArticle struct {
		cli zrpc.Client
	}
)

func NewArticle(cli zrpc.Client) Article {
	return &defaultArticle{
		cli: cli,
	}
}

func (m *defaultArticle) AddArticle(ctx context.Context, in *AddArticleRequest, opts ...grpc.CallOption) (*ArticleList, error) {
	client := article_client.NewArticleClient(m.cli.Conn())
	return client.AddArticle(ctx, in, opts...)
}

func (m *defaultArticle) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*ArticleList, error) {
	client := article_client.NewArticleClient(m.cli.Conn())
	return client.GetArticle(ctx, in, opts...)
}

func (m *defaultArticle) GetArticleByPaging(ctx context.Context, in *GetArticleByPagingRequest, opts ...grpc.CallOption) (*ArticleList, error) {
	client := article_client.NewArticleClient(m.cli.Conn())
	return client.GetArticleByPaging(ctx, in, opts...)
}

func (m *defaultArticle) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*ArticleList, error) {
	client := article_client.NewArticleClient(m.cli.Conn())
	return client.UpdateArticle(ctx, in, opts...)
}

func (m *defaultArticle) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*ArticleList, error) {
	client := article_client.NewArticleClient(m.cli.Conn())
	return client.DeleteArticle(ctx, in, opts...)
}