package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/franciscocunha55/kubernetes-mcp-server/k8s"

)

func ListNamespacesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientSet, err := GetClientSet()
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	namespacesList := k8s.ListNamespaces(clientSet)
	return mcp.NewToolResultStructured(namespacesList, "Namespaces listed successfully"), nil
}

var ToolListNamespaces = mcp.NewTool("list_namespaces",
	mcp.WithDescription("List all namespaces in the Kubernetes cluster"),
)
