package mkspipelinerun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func getcommand(mksclient *versioned.Clientset) *cobra.Command {
	mksPrGet := &cobra.Command{
		Use:   "get",
		Short: "Get PipelineRuns",
		Annotations: map[string]string{
			"commandType": "main",
		},
		Run: func(cmd *cobra.Command, args []string) {
			get, err := mksclient.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Get(context.TODO(), resourceName, v1.GetOptions{})
			if err != nil {
				klog.Fatalf("Get MksPipelineRun failed!", err.Error())
			} else {
				printPipelineRun(get)
			}
		},
	}
	mksPrGet.Flags().StringVar(&resourceName, "rn", "", "Name of MksPipelineRun Resource to be fetched")
	mksPrGet.MarkFlagRequired("rn")
	mksPrGet.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")
	return mksPrGet
}

func printPipelineRun(get *v1alpha1.MksPipelineRun) {
	fmt.Println("MKS PIPELINERUN: ")
	fmt.Printf("\n")
	fmt.Println("UID: ", get.UID)
	fmt.Println("Name: ", get.Name)
	fmt.Println("Namespace: ", get.Namespace)
	fmt.Println("PipelineRef: ", get.Spec.PipelineRef)
	fmt.Println("Created At: ", get.CreationTimestamp)
}
