# A useful MCP Tool


## Usage

### Install
```shell
go install github.com/chslink/mcp-tools/cmd/mcpgen@latest
```


### WriteYourCode 
Create Go functions with Tool prefix and standardized comments:
```go
// ToolExample Example tool description
// @param1 Parameter 1 description
// @param2 Parameter 2 description
func ToolExample(param1 string, param2 int) (*mcp_golang.ToolResponse, error) {
// Implementation logic...
}
```

### GenerateCode
Execute generation command:
```shell
mcpgen -root=./sometools
```

### Import sometools package

```go
package main
import (
	"log"
	
	_ "github.com/yourrepository/sometools"
	"github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/chslink/mcp-tools/gen"
)

func main() {
	done := make(chan struct{})
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	var err error
	// Register tools
	for _, tool := range gen.GetTools() {
		err = server.RegisterTool(tool.Name, tool.Description, tool.Func)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}
```
