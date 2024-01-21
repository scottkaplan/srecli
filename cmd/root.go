package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sre",
	Short: "A CLI for doing k8s control plane stuff.",
	Long: `A CLI for doing k8s control plane stuff. Similar to kubectl.
         Used to scale resources in the cluster. For example:

  sre list: List pods in the cluster
  sre list --namespace scott: List pods in namespace scott
  sre info --namespace scott: Print info about pods in namespace scott
  sre info --pod healthy-pod: Print info about healthy-pod
  sre scale --replicas 3 --pod healthy-pod: Run 3 healthy-pod pods
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sre.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


