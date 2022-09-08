package page

import (
	"act/common/models"
	"encoding/json"
	"errors"
	"math"

	"google.golang.org/protobuf/runtime/protoimpl"
)

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}
type CommonRpcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//返回数据JSON
	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

// 查询分页
func PageResult(page *PageRequest, total int64, data interface{}, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	return JsonResult(&models.PageResponse{
		Records: data,
		Current: page.Offset/page.Limit + 1,
		Size:    page.Limit,
		Total:   total,
		Pages:   int64(math.Ceil(float64(total / page.Limit))),
	}, err)
}
func JsonResult(data interface{}, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	if data == nil {
		return &CommonRpcRes{}, errors.New("not found")
	}
	resJson, err := json.Marshal(data)
	return Result(string(resJson), err)
}
func Result(data string, err error) (*CommonRpcRes, error) {
	if err != nil {
		return &CommonRpcRes{}, err
	}
	return &CommonRpcRes{
		Json: data,
	}, nil
}
