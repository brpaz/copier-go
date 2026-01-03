package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	Version   = "dev"
	BuildDate = "unknown"
	Commit    = "unknown"
)

func main() {
	cmd := &cli.Command{
		Name:                  "{{ project_slug }}",
		Version:               Version,
		Usage:                 "{{ project_description }}",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "hello",
				Aliases: []string{"h"},
				Usage:   "Say hello",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "name to greet",
						Value:   "World",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					name := cmd.String("name")
					fmt.Printf("Hello, %s!\n", name)
					return nil
				},
				ShellComplete: func(ctx context.Context, cmd *cli.Command) {
					// Custom completion suggestions
					if cmd.NArg() == 0 {
						fmt.Println("--name")
						fmt.Println("-n")
					}
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show version information",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Printf("{{ project_slug }} version %s\n", Version)
					fmt.Printf("Build Date: %s\n", BuildDate)
					fmt.Printf("Commit: %s\n", Commit)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
