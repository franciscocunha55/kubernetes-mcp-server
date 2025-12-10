package main

import (
	"fmt"
	"log"
	"os"

	"github.com/franciscocunha55/kubernetes-mcp-server/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {

	mode := "stdio"
	if len(os.Args) > 1 && os.Args[1] == "http" {
		mode = "http"
	}

	mcpServer := server.NewMCPServer(
		"Kubernetes MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
	)

	mcpServer.AddTool(tools.ToolHelloWorld, tools.HelloHandler)
	mcpServer.AddTool(tools.ToolListNamespaces, tools.ListNamespacesHandler)

	if mode == "http" {
		streamableHTTPServer := server.NewStreamableHTTPServer(mcpServer, server.WithStateLess(true))
		log.Printf("Starting MCP server on HTTP at port :8080...")
		log.Printf("MCP endpoint will be available at: http://localhost:8080/mcp")
		if err := streamableHTTPServer.Start(":8080"); err != nil {
			log.Fatalf("Failed to start MCP server: %v", err)
		}
		
	} else {
		fmt.Fprintln(os.Stderr, "Starting MCP server on stdio...")
		if err := server.ServeStdio(mcpServer); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		}
	}
}


