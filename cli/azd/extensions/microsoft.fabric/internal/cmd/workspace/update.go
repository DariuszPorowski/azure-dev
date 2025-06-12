// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("UPDATE...")

			return nil
		},
	}
}
