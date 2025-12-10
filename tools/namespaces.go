package tools

import (
	"context"
	"fmt"

	"github.com/franciscocunha55/kubernetes-mcp-server/k8s"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListNamespacesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientSet, err := GetClientSet()
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	namespacesList, err := k8s.ListNamespaces(clientSet)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var namespaceNames []string
	for _, ns := range namespacesList.Items {
		namespaceNames = append(namespaceNames, ns.Name)
	}

	textOutput := fmt.Sprintf("Found %d namespaces:\n", len(namespaceNames))
	for _, name := range namespaceNames {
		textOutput += fmt.Sprintf("- %s\n", name)
	}

	result := map[string]interface{}{
		"namespaces": namespaceNames,
		"count":      len(namespaceNames),
	}

	return mcp.NewToolResultStructured(result, textOutput), nil
}

var ToolListNamespaces = mcp.NewTool("list_namespaces",
	mcp.WithDescription("List all namespaces in the Kubernetes cluster"),
)
