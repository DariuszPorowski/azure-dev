// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Gets a workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("SHOW...")

			return nil
		},
	}
}
