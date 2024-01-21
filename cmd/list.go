package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists pods in a cluster",
	Long: `Lists pods in a cluster
--namespace filters by namespace. For example:

  sre list: List pods in the cluster
  sre list --namespace scott: List pods in namespace scott
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		namespace, _ := cmd.Flags().GetString("namespace")
		fmt.Println("namespace="+namespace)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
