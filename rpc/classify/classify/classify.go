// Code generated by goctl. DO NOT EDIT.
// Source: classify.proto

package classify

import (
	"context"
	"front-gin/rpc/classify/classify_client"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddClassifyRequest    = classify_client.AddClassifyRequest
	ArticleData           = classify_client.ArticleData
	ClassifyData          = classify_client.ClassifyData
	ClassifyList          = classify_client.ClassifyList
	DeleteClassifyRequest = classify_client.DeleteClassifyRequest
	GetAllClassifyRequest = classify_client.GetAllClassifyRequest
	GetClassifyRequest    = classify_client.GetClassifyRequest
	UpdateClassifyRequest = classify_client.UpdateClassifyRequest

	Classify interface {
		AddClassify(ctx context.Context, in *AddClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error)
		GetClassify(ctx context.Context, in *GetClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error)
		GetAllClassify(ctx context.Context, in *GetAllClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error)
		UpdateClassify(ctx context.Context, in *UpdateClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error)
		DeleteClassify(ctx context.Context, in *DeleteClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error)
	}

	defaultClassify struct {
		cli zrpc.Client
	}
)

func NewClassify(cli zrpc.Client) Classify {
	return &defaultClassify{
		cli: cli,
	}
}

func (m *defaultClassify) AddClassify(ctx context.Context, in *AddClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error) {
	client := classify_client.NewClassifyClient(m.cli.Conn())
	return client.AddClassify(ctx, in, opts...)
}

func (m *defaultClassify) GetClassify(ctx context.Context, in *GetClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error) {
	client := classify_client.NewClassifyClient(m.cli.Conn())
	return client.GetClassify(ctx, in, opts...)
}

func (m *defaultClassify) GetAllClassify(ctx context.Context, in *GetAllClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error) {
	client := classify_client.NewClassifyClient(m.cli.Conn())
	return client.GetAllClassify(ctx, in, opts...)
}

func (m *defaultClassify) UpdateClassify(ctx context.Context, in *UpdateClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error) {
	client := classify_client.NewClassifyClient(m.cli.Conn())
	return client.UpdateClassify(ctx, in, opts...)
}

func (m *defaultClassify) DeleteClassify(ctx context.Context, in *DeleteClassifyRequest, opts ...grpc.CallOption) (*ClassifyList, error) {
	client := classify_client.NewClassifyClient(m.cli.Conn())
	return client.DeleteClassify(ctx, in, opts...)
}