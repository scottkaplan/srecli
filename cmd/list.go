package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
		namespace, _ := cmd.Flags().GetString("namespace")

		kubeconfig := "/home/scott/.kube/config"
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		for _, pod := range pods.Items {
			fmt.Fprintln(w, pod.Name+"\t"+pod.Namespace+"\t")
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
