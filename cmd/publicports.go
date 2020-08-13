package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var showType bool

// publicportsCmd represents the publicports command
var publicportsCmd = &cobra.Command{
	Use:   "publicports",
	Short: "Prints all public ports used by docker containers",
	Long: `Prints all published ports used by docker containers.
These can be used to add the necessary rules to your firewall.
Format: 80\n443\n`,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cli, err := client.NewEnvClient()
		if err != nil {
			return err
		}

		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			return err
		}

		for _, container := range containers {
			for _, port := range container.Ports {
				if port.PublicPort != 0 {
					if showType {
						fmt.Printf("%d/%s\n", port.PublicPort, port.Type)
					} else {
						fmt.Printf("%d\n", port.PublicPort)
					}
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(publicportsCmd)
	publicportsCmd.Flags().BoolVarP(&showType, "type", "t", false, "add types, format: 80/tcp\\n1234/udp\\n")
}
