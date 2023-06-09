package utils

import (
	"fmt"
	"front-gin/rpc/classify/classify_client"
	"front-gin/src/types"
	"github.com/duke-git/lancet/v2/convertor"
)

func ClassifyRpcToApi(data []*classify_client.ClassifyData) (classifyData []*types.ClassifyData, err error) {
	classifyData = make([]*types.ClassifyData, 0, len(data))
	for _, datum := range data {
		classify := types.ClassifyData{}
		err = convertor.CopyProperties(&classify, datum)
		if err != nil {
			fmt.Println("ClassifyRpcToApi", err)
			return nil, err
		}
		classifyData = append(classifyData, &classify)
	}

	return classifyData, nil
}