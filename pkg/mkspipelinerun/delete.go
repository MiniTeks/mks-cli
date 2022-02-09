package mkspipelinerun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func deletecommand(mksclient *versioned.Clientset) *cobra.Command {
	mksPrDelete := &cobra.Command{
		Use:   "delete",
		Short: "Delete PipelineRuns",
		Annotations: map[string]string{
			"commandType": "main",
		},
		Run: func(cmd *cobra.Command, args []string) {
			delErr := mksclient.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Delete(context.TODO(), resourceName, v1.DeleteOptions{})
			if delErr != nil {
				klog.Fatalf("Delete MksPipelineRun failed!", delErr.Error())
			}
			fmt.Println("Mks PipelineRun ", resourceName, " deleted")
		},
	}
	mksPrDelete.Flags().StringVar(&resourceName, "rn", "", "Name of MksPipelineRun Resource to be deleted")
	mksPrDelete.MarkFlagRequired("rn")
	mksPrDelete.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")

	return mksPrDelete
}
