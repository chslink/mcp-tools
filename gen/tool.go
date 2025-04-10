package gen

import (
	"encoding/json"
	"sync"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type Tool struct {
	Name        string
	Description string
	Args        any
	Func        any
}

var (
	mux   sync.Mutex
	tools []Tool
)

func RegisterTool(tool Tool) {
	mux.Lock()
	defer mux.Unlock()
	tools = append(tools, tool)
}

func GetTools() []Tool {
	mux.Lock()
	defer mux.Unlock()
	return tools
}

func ConvertMcpResp(val any, err error) (*mcp_golang.ToolResponse, error) {
	if err != nil {
		return nil, err
	}
	if ret, ok := val.(string); ok {
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(ret)), nil
	}
	if ret, ok := val.(*string); ok {
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(*ret)), nil
	}
	buff, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(buff))), nil
}
