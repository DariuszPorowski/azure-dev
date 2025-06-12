// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
	"microsoft.fabric/internal/cmd/mcp"
	"microsoft.fabric/internal/cmd/workspace"
	"microsoft.fabric/internal/pkg/opts"
)

var pf = &opts.RootOpts{}

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "fabric <command> [options]",
		Short:         "AZD Extension for Microsoft Fabric",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// --debug
	rootCmd.PersistentFlags().BoolVarP(&pf.Debug, "debug", "d", false, "Enable debug output")

	// --output
	rootCmd.PersistentFlags().VarP(
		enumflag.New(&pf.Output, "output", opts.Formats, enumflag.EnumCaseInsensitive),
		"output", "o",
		fmt.Sprintf("The output format (the supported formats are: %s).", strings.Join(opts.GetFormats(), ", ")))
	rootCmd.PersistentFlags().Lookup("output").NoOptDefVal = "json"

	// --query
	rootCmd.PersistentFlags().StringVarP(&pf.Query, "query", "q", "", "Query to filter results (JSONPath syntax)")

	rootCmd.AddCommand(workspace.NewWorkspaceCmd(pf))
	rootCmd.AddCommand(mcp.NewMcpCmd(pf))
	// rootCmd.AddCommand(newContextCommand())
	// rootCmd.AddCommand(newPromptCommand())
	rootCmd.AddCommand(newVersionCmd())

	return rootCmd
}
