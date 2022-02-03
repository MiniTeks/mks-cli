package mkspipelinerun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	vr "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var mksPrList = &cobra.Command{
	Use:   "list",
	Short: "List PipelineRuns",
	Annotations: map[string]string{
		"commandType": "main",
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, err := cmd.Flags().GetString("config")
		if err != nil {
			klog.Fatalf("Error in getting kubeconfig path")
		} else {
			list(cfgFile)
		}
	},
}

func listcommand() *cobra.Command {

	return mksPrList
}

func list(cfgFile string) {
	cfg, err := clientcmd.BuildConfigFromFlags("", cfgFile)

	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}
	mksClient, err := vr.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error generating new clientset: %v", err)
	}
	fet, err := mksClient.MkscontrollerV1alpha1().MksPipelineRuns(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Error!!! Coldn't get the resource(s) from the namespace %s\n", namespace)
		fmt.Errorf("Couldn't create mksPipelineRun", err.Error())
	} else {
		printList(fet)
	}

}

func printList(fet *v1alpha1.MksPipelineRunList) {
	fmt.Println("Here is List")
	fmt.Printf("\n")
	for i := range fet.Items {
		fmt.Println("MKS PIPELINERUN: ", i+1)
		fmt.Println("UID: ", fet.Items[i].UID)
		fmt.Println("Name: ", fet.Items[i].Name)
		fmt.Println("Namespace: ", fet.Items[i].Namespace)
		fmt.Println("PipelineRef: ", fet.Items[i].Spec.PipelineRef)
		fmt.Println("Created At: ", fet.Items[i].CreationTimestamp)

		fmt.Printf("\n")
	}

}

func init() {
	mksPrList.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")
}
