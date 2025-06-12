// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package workspace

import (
	"errors"

	"github.com/azure/azure-dev/cli/azd/pkg/output"
	"github.com/spf13/cobra"
	"microsoft.fabric/internal/pkg/opts"
)

func newListCmd(o *opts.RootOpts) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists workspaces",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := getWorkspaceClient()
			if err != nil {
				return err
			}

			respList, err := client.ListWorkspaces(cmd.Context(), nil)
			if err != nil {
				return err
			}

			if o.Output == opts.FormatTable {
				f, err := output.NewFormatter("table")
				if err != nil {
					return err
				}

				tOpts := output.TableFormatterOptions{
					Columns: []output.Column{
						{
							Heading:       "ID",
							ValueTemplate: "{{.ID}}",
						},
						{
							Heading:       "DISPLAY NAME",
							ValueTemplate: "{{.DisplayName}}",
						},
						{
							Heading:       "CAPACITY ID",
							ValueTemplate: "{{.CapacityID}}",
						},
						{
							Heading:       "TYPE",
							ValueTemplate: "{{.Type}}",
						},
						{
							Heading:       "DESCRIPTION",
							ValueTemplate: "{{.Description}}",
						},
					},
				}

				if err := f.Format(respList, cmd.OutOrStdout(), tOpts); err != nil {
					return err
				}
			} else if o.Output == opts.FormatJSON {
				f, err := output.NewFormatter("json")
				if err != nil {
					return err
				}

				if err := f.Format(respList, cmd.OutOrStdout(), nil); err != nil {
					return err
				}
			} else if o.Output == opts.FormatEnvVars {
				return errors.New("dotenv output format is not supported for this command")
			}

			return nil
		},
	}
}
