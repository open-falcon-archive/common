package model

import (
	"fmt"
)

// code == 0 => success
// code == 1 => bad request
type SimpleRpcResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (this *SimpleRpcResponse) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s>", this.Code, this.Msg)
}

type NullRpcRequest struct {
}
