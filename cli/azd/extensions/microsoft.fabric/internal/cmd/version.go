// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Populated at build time
	Version   = "dev" // Default value for development builds
	Commit    = "none"
	BuildDate = "unknown"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version of the application",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Version: %s\nCommit: %s\nBuild Date: %s\n", Version, Commit, BuildDate)

			return nil
		},
	}
}
