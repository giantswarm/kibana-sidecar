package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/giantswarm/kibana-sidecar/config"
	"github.com/giantswarm/kibana-sidecar/service/kibana"
)

// DaemonCmd is the command cobra executes when no sub-command is called
var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Runs a kibana-sidecar as a service",
	Long: `Running as a service/daemon, kibana-sidecar will constantly
make sure that the Kibana configuration is what we configured. This
will override config changes made interactively.`,
	Run: runDaemon,
}

func init() {
	RootCmd.AddCommand(DaemonCmd)
}

func runDaemon(cmd *cobra.Command, args []string) {

	for {
		kibana.WriteConfig()
		time.Sleep(config.DaemonWaitInterval)
	}

}
