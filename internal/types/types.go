package types

type HelloReq struct {
	Msg string `form:"msg"`
}

type HelloResp struct {
	Msg string `json:"msg"`
}
