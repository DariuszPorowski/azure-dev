// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Creates a new workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("CREATE...")

			return nil
		},
	}
}
