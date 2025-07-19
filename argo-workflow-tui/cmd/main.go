package main

import (
	"fmt"
	"log"
	
	config "argo-workflow-tui/pkg" // Import your config package
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	
	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	fmt.Printf("✅ Configuration loaded successfully!\n")
	fmt.Printf("🌐 Argo Server URL: %s\n", cfg.ArgoServerURL)
	fmt.Printf("🔑 Authentication: %s\n", func() string {
		if cfg.ArgoToken != "" {
			return "ARGO_TOKEN (set)"
		}
		if cfg.ArgoIDToken != "" {
			return "ARGO_ID_TOKEN (set)" 
		}
		return "none"
	}())

	fmt.Printf("\n🚀 Ready to connect to Argo Workflows server!\n")
	fmt.Printf("📝 Next steps:\n")
	fmt.Printf("   - Implement workflow listing\n")
	fmt.Printf("   - Add workflow creation\n") 
	fmt.Printf("   - Build TUI interface\n")
}


// func listWorkflows(cfg *config.Config) {
// 	// This function lists workflow templates and their available arguments.
// 	// It assumes the Argo server is accessible and authentication is handled via the config.

// 	// Import statements would be:
// 	// "context"
// 	// "github.com/argoproj/argo-workflows/v3/pkg/apiclient"
// 	// "github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflowtemplate"
// 	// "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/scheme"
// 	// "google.golang.org/grpc/metadata"
// 	// "os"
// 	// "fmt"

// 	ctx := context.Background()

// 	// Prepare authentication metadata
// 	md := metadata.New(nil)
// 	if cfg.ArgoToken != "" {
// 		md.Append("authorization", "Bearer "+cfg.ArgoToken)
// 	} else if cfg.ArgoIDToken != "" {
// 		md.Append("authorization", "Bearer "+cfg.ArgoIDToken)
// 	}
// 	ctx = metadata.NewOutgoingContext(ctx, md)

// 	// Connect to Argo server
// 	argoClient, err := apiclient.NewClient(&apiclient.ClientOptions{
// 		ServerAddr: cfg.ArgoServerURL,
// 		// InsecureSkipVerify: true, // Set if needed for self-signed certs
// 	})
// 	if err != nil {
// 		fmt.Printf("❌ Failed to create Argo client: %v\n", err)
// 		return
// 	}
// 	defer argoClient.Close()

// 	wftmplServiceClient, err := argoClient.NewWorkflowTemplateServiceClient()
// 	if err != nil {
// 		fmt.Printf("❌ Failed to create WorkflowTemplate service client: %v\n", err)
// 		return
// 	}

// 	// List all workflow templates in all namespaces (or default)
// 	namespace := "default"
// 	listReq := &workflowtemplate.WorkflowTemplateListRequest{Namespace: namespace}
// 	listResp, err := wftmplServiceClient.ListWorkflowTemplates(ctx, listReq)
// 	if err != nil {
// 		fmt.Printf("❌ Failed to list workflow templates: %v\n", err)
// 		return
// 	}

// 	if len(listResp.Items) == 0 {
// 		fmt.Println("No workflow templates found.")
// 		return
// 	}

// 	fmt.Printf("📋 Workflow Templates in namespace '%s':\n", namespace)
// 	for _, tmpl := range listResp.Items {
// 		fmt.Printf("\n- Name: %s\n", tmpl.GetName())
// 		if tmpl.Spec.Arguments != nil && len(tmpl.Spec.Arguments.Parameters) > 0 {
// 			fmt.Println("  Arguments:")
// 			for _, param := range tmpl.Spec.Arguments.Parameters {
// 				fmt.Printf("    - %s", param.Name)
// 				if param.Value != nil {
// 					fmt.Printf(" (default: %s)", param.Value.String())
// 				}
// 				if param.Description != "" {
// 					fmt.Printf(" - %s", param.Description)
// 				}
// 				fmt.Println()
// 			}
// 		} else {
// 			fmt.Println("  Arguments: (none)")
// 		}
// 	}
// }