// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("DELETE...")

			return nil
		},
	}
}
