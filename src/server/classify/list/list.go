package list

import (
	"context"
	"fmt"
	"front-gin/rpc/classify/classify_client"
	"front-gin/rpc/server/svc"
	"front-gin/src/global"
	"front-gin/src/types"
	"front-gin/src/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifyLogic {
	return &GetClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassifyLogic) GetClassify(req *types.GetClassifyReq) (resp *types.ClassifyRespose, err error) {
	var classify *classify_client.ClassifyList

	if req.ID == -1 {
		classify, err = l.svcCtx.Classify.GetAllClassify(l.ctx, &classify_client.GetAllClassifyRequest{})
	} else {
		classify, err = l.svcCtx.Classify.GetClassify(l.ctx, &classify_client.GetClassifyRequest{
			Id: []uint64{uint64(req.ID)},
		})
	}

	if err != nil {
		fmt.Println("暂无数据")
		l.Errorf("id=%v,暂无数据\n", req.ID)
		l.Error(err)
		return &types.ClassifyRespose{
			Code:    200,
			Data:    []*types.ClassifyData{},
			Message: global.NULL_DATA,
		}, nil
	}

	if len(classify.GetClassify()) <= 0 {
		return &types.ClassifyRespose{
			Code:    200,
			Data:    []*types.ClassifyData{},
			Message: global.NULL_DATA,
		}, nil
	}

	classifyData, err := utils.ClassifyRpcToApi(classify.GetClassify())
	if err != nil {
		return &types.ClassifyRespose{
			Code:    200,
			Data:    []*types.ClassifyData{},
			Message: global.ERROR,
		}, nil
	}

	return &types.ClassifyRespose{
		Code:    200,
		Data:    classifyData,
		Message: global.SUCCESS,
	}, nil
}