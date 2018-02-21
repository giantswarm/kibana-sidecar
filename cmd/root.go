package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/giantswarm/kibana-sidecar/config"
)

// RootCmd is the command cobra executes when no sub-command is called
var RootCmd = &cobra.Command{
	Use:   "kibana-sidecar",
	Short: "A little confiiguration service for Kibana",
	Long:  "This service writes Kibana confguration to an Elasticsearch endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(
		&config.ElasticsearchEndpoint,
		"elasticsearch-endpoint",
		"",
		config.ElasticsearchEndpointDefault,
		"Elasticsearch endpoint URL, e. g. 'http://elasticsearch:9200'")
}

// Execute runs the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
