package main

import (
	"log"

	"github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"

	"github.com/chslink/mcp-tools/gen"
	_ "github.com/chslink/mcp-tools/tools"
)

func main() {
	done := make(chan struct{})

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	var err error
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
