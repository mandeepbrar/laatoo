package websocket

type rpcRequest struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	Id      string                 `json:"id"`
}

type rpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Id      string      `json:"id"`
}

type rpcNotification struct {
	Jsonrpc string      `json:"jsonrpc"`
	Data    interface{} `json:"data"`
	Id      string      `json:"id"`
}
