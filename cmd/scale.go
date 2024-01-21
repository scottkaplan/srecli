package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Manage the number of pods (replicas) running in a cluster",
	Long: `Manage the number of pods (replicas) running in a cluster
For example:

  sre scale --replicas 3 --pod healthy-pod: Run 3 healthy-pod pods

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scale called")
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
