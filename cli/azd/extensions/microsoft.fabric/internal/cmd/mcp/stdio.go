// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	stdlog "log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	fabcore "github.com/microsoft/fabric-sdk-go/fabric/core"
	"github.com/spf13/cobra"
	"microsoft.fabric/internal/pkg/azdctx"
)

func newStdioCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stdio",
		Short: "Starts the MCP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			log := slog.New(slog.NewTextHandler(os.Stderr, nil))
			log.Info("Logger initialized")

			ctx := cmd.Context()

			ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			defer stop() // Ensure stop is called to release resources
			log.Info("Signal handling initialized")

			opts := []server.ServerOption{
				server.WithToolCapabilities(true),
				server.WithResourceCapabilities(true, true),
				server.WithLogging(),
			}

			// Create a new MCP server
			mcpServer := server.NewMCPServer(
				"fabric-mcp-server",
				"0.0.1",
				opts...,
			)

			// Add tool
			tool := mcp.NewTool("hello_world",
				mcp.WithDescription("Say hello to someone"),
				mcp.WithString("name",
					mcp.Required(),
					mcp.Description("Name of the person to greet"),
				),
			) // Add tool handler
			mcpServer.AddTool(tool, helloHandler)

			workspaceTool := mcp.NewTool("get_workspaces",
				mcp.WithDescription("Get Fabric workspaces - returns all workspaces as JSON or a specific workspace by display name"),
				mcp.WithString("displayName",
					mcp.Description("Optional: specific workspace display name to retrieve"),
				),
			)

			mcpServer.AddTool(workspaceTool, workspaceHandler)

			stdioServer := server.NewStdioServer(mcpServer)

			stdioLogger := stdlog.New(cmd.OutOrStdout(), "[StdioServer] ", 0) // Use logger's writer
			stdioServer.SetErrorLogger(stdioLogger)
			log.Info("Stdio server transport created")

			// Start listening for messages
			errC := make(chan error, 1)
			go func() {
				log.Info("Starting to listen on stdio...")
				in, out := io.Reader(cmd.InOrStdin()), io.Writer(cmd.OutOrStdout())
				errC <- stdioServer.Listen(ctx, in, out)
			}()

			_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "Fabric MCP Server running on stdio\n")
			log.Info("Server running, waiting for requests or signals...")

			// Wait for shutdown signal
			select {
			case <-ctx.Done():
				log.Info("Shutdown signal received, context cancelled.")
			case err := <-errC:
				if err != nil && err != context.Canceled {
					log.Error("Server encountered an error %v", err)
					// We might want os.Exit(1) here depending on desired behavior
				} else {
					log.Info("Server listener stopped gracefully.")
				}
			}

			log.Info("Server shutting down.")

			// Start the stdio server
			// if err := server.ServeStdio(s); err != nil {
			// 	fmt.Printf("Server error: %v\n", err)
			// }

			return nil
		},
	}
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Missing required parameter", err), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

func workspaceHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	displayName := request.GetString("displayName", "")

	client, err := getWorkspaceClient()
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting workspace client", err), nil
	}

	respList, err := client.ListWorkspaces(ctx, nil)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error listing workspaces", err), nil
	}

	// If display name is provided, find specific workspace
	if displayName != "" {
		for _, entity := range respList {
			if *entity.DisplayName == displayName {
				entityJSON, err := entity.MarshalJSON()
				if err != nil {
					return mcp.NewToolResultErrorFromErr("Error marshaling workspace", err), nil
				}

				// Return the raw JSON for the specific workspace
				return mcp.NewToolResultText(string(entityJSON)), nil
			}
		}

		// Return an error if workspace with display name not found
		return mcp.NewToolResultErrorf("Workspace with display name '%s' not found", displayName), nil
	}

	// If no display name provided, return all workspaces as JSON array
	listJSON, err := json.MarshalIndent(respList, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error marshaling workspace list", err), nil
	}

	// Return the raw JSON array of all workspaces
	return mcp.NewToolResultText(string(listJSON)), nil
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
