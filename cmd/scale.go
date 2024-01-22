package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Manage the number of pods (replicas) running in a deployment",
	Long: `Manage the number of pods (replicas) running in a deployment
For example:

  sre scale --replicas 3 --deployment prod-deploy: Run 3 pods in the prod-deploy deployment

`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		podName, _ := cmd.Flags().GetString("pod")
		if len(podName) == 0 {
			fmt.Println("--pod is required")
			os.Exit(1)
		}

		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		if len(kubeconfig) == 0 {
			kubeconfig = homedir.HomeDir()+"/.kube/config"
		}
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("scale called")
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
	infoCmd.Flags().String("deployment", "", "name of deployment")
}
