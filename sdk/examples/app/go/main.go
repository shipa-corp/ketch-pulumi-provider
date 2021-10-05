package main

import (
	"context"
	"fmt"
	"os"

	"github.com/brunoa19/ketch-pulumi-provider/sdk/go/ketch"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	// define our program that creates our pulumi resources.
	// we refer to this style as "inline" pulumi programs where both program + automation can be compiled in the same
	// binary. no need for separate projects.
	deployFunc = func(ctx *pulumi.Context) error {
		framework, err := ketch.NewFramework(ctx, "go-app-framework-1", &ketch.FrameworkArgs{
			Framework: &ketch.FrameworkFrameworkArgs{
				Name: pulumi.String("go-app-framework-1"),
				IngressController: &ketch.FrameworkFrameworkIngressControllerArgs{
					ClassName:       pulumi.String("istio"),
					ServiceEndpoint: pulumi.String("1.2.3.4"),
					Type:            pulumi.String("istio"),
				},
			},
		})

		if err != nil {
			return err
		}

		app, err := ketch.NewApp(ctx, "go-app-1", &ketch.AppArgs{
			App: &ketch.AppAppArgs{
				Name:      pulumi.String("go-app-1"),
				Image:     pulumi.String("docker.io/shipasoftware/bulletinboard:1.0"),
				Framework: framework.Framework.Name(),
				Cnames:    pulumi.StringArray{pulumi.String("go-app-1.ketch.io"), pulumi.String("go-app-2.ketch.io")},
				Ports:     pulumi.IntArray{pulumi.Int(8008), pulumi.Int(8009)},
				Units:     pulumi.IntPtr(5),
				Processes: &ketch.AppAppProcessArray{
					&ketch.AppAppProcessArgs{
						Name: pulumi.String("web"),
						Cmds: pulumi.StringArray{
							pulumi.String("docker-entrypoint.sh"),
							pulumi.String("npm"),
							pulumi.String("start"),
						},
					},
				},
				RoutingSettings: &ketch.AppAppRoutingSettingsArgs{
					Weight: pulumi.IntPtr(100),
				},
				Version: pulumi.IntPtr(1),
			},
		})

		ctx.Export("app-name", app.App.Name().ToStringOutput())
		return nil

	}
)

func main() {
	// to destroy our program, we can run `go run main.go destroy`
	destroy := false
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		if argsWithoutProg[0] == "destroy" {
			destroy = true
		}
	}

	ctx := context.Background()

	projectName := "ketch-app-go"
	// we use a simple stack name here, but recommend using auto.FullyQualifiedStackName for maximum specificity.
	stackName := "dev"
	// stackName := auto.FullyQualifiedStackName("myOrgOrUser", projectName, stackName)

	// create or select a stack matching the specified name and project.
	// this will set up a workspace with everything necessary to run our inline program (deployFunc)
	s, err := auto.UpsertStackInlineSource(ctx, stackName, projectName, deployFunc)

	fmt.Printf("Created/Selected stack %q\n", stackName)

	fmt.Println("Successfully set config")
	fmt.Println("Starting refresh")

	_, err = s.Refresh(ctx)
	if err != nil {
		fmt.Printf("Failed to refresh stack: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Refresh succeeded!")

	if destroy {
		fmt.Println("Starting stack destroy")

		// wire up our destroy to stream progress to stdout
		stdoutStreamer := optdestroy.ProgressStreams(os.Stdout)

		// destroy our stack and exit early
		_, err := s.Destroy(ctx, stdoutStreamer)
		if err != nil {
			fmt.Printf("Failed to destroy stack: %v", err)
		}
		fmt.Println("Stack successfully destroyed")
		os.Exit(0)
	}

	fmt.Println("Starting update")

	// wire up our update to stream progress to stdout
	stdoutStreamer := optup.ProgressStreams(os.Stdout)

	// run the update to deploy our s3 website
	res, err := s.Up(ctx, stdoutStreamer)
	if err != nil {
		fmt.Printf("Failed to update stack: %v\n\n", err)
		os.Exit(1)
	}

	fmt.Println("Update succeeded!")

	fmt.Printf("Outputs: %+v\n", res.Outputs)
}
