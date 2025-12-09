package main

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"
)


func main() {
	// This is a placeholder for the main function.
	fmt.Println("Hello, World!")

	mcpServer := server.NewMCPServer(
		"Kubernetes MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
	)

	streamableHTTPServer := server.NewStreamableHTTPServer(mcpServer)
	log.Printf("Starting MCP server on HTTP at port :8080...")
	log.Printf("MCP endpoint will be available at: http://localhost:8080/mcp")
	if err:=streamableHTTPServer.Start(":8080"); err != nil {
		log.Fatalf("Failed to start MCP server: %v", err)
	}
	
}