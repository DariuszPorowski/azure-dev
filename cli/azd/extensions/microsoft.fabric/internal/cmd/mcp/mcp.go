// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package mcp

import (
	"github.com/spf13/cobra"
	"microsoft.fabric/internal/pkg/opts"
)

func NewMcpCmd(o *opts.RootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Manage MCP Server",
	}

	// CRUDL
	cmd.AddCommand(newStdioCmd())

	return cmd
}
