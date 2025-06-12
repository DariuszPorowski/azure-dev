// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package workspace

import (
	fabcore "github.com/microsoft/fabric-sdk-go/fabric/core"
	"github.com/spf13/cobra"
	"microsoft.fabric/internal/pkg/azdctx"
	"microsoft.fabric/internal/pkg/opts"
)

func NewWorkspaceCmd(o *opts.RootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workspace",
		Short: "Manage workspace",
	}

	// CRUDL
	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newShowCmd())
	cmd.AddCommand(newUpdateCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newListCmd(o))

	return cmd
}

func getWorkspaceClient() (*fabcore.WorkspacesClient, error) {
	cred, err := azdctx.GetAzdCredential()
	if err != nil {
		return nil, err
	}

	// fabricClient, err := fabric.NewClient(cred, nil, nil)
	// if err != nil {
	// 	return nil, err
	// }

	fabricCF, err := fabcore.NewClientFactory(cred, nil, nil)
	if err != nil {
		return nil, err
	}

	// fabricCF := fabcore.NewClientFactoryWithClient(*fabricClient)
	fabWS := fabricCF.NewWorkspacesClient()

	return fabWS, nil
}
