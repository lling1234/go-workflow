package types

import "errors"

type CommonResponse struct {
	Code    uint64      `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

//获取CommonResponse
func GetCommonResponse(err error, T any) (*CommonResponse, error) {
	if err != nil {
		return &CommonResponse{
			Code:    400,
			Msg:     err.Error(),
			Success: false,
		}, err
	}
	return &CommonResponse{
		Code:    200,
		Data:    T,
		Success: true,
	}, nil
}

func GetErrorCommonResponse(msg string) (*CommonResponse, error) {
	return &CommonResponse{
		Code:    400,
		Msg:     msg,
		Success: false,
	}, errors.New(msg)
}

func GetSuccessCommonResponse(T any) (*CommonResponse, error) {
	return &CommonResponse{
		Code:    200,
		Data:    T,
		Success: true,
	}, nil
}
func GetSuccessCommonPageResponse(T any) (*CommonResponse, error) {
	return &CommonResponse{
		Code:    200,
		Data:    T,
		Success: true,
	}, nil
}
