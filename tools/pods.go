package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func HelloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

var ToolHelloWorld = mcp.NewTool("hello_world",
	mcp.WithDescription("Say hello to someone"),
	mcp.WithString("name",
		mcp.Required(),
		mcp.Description("Name of the person to greet"),
	),
)
