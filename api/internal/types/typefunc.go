package types

import "errors"

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
