package withJson

import "encoding/json"

type HelloService struct {
}

func (*HelloService) Hello(req string, reply *string) error {
	*reply = "Hello: " + req
	return nil
}

// 客户端
type clientRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params"`
	Id     uint64         `json:"id"`
}

// 服务端
type serverRequest struct {
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
	Id     *json.RawMessage `json:"id"`
}
