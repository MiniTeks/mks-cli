package mkspipelinerun

import (
	"context"
	"fmt"

	vr "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var mksPrDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete PipelineRuns",
	Annotations: map[string]string{
		"commandType": "main",
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, err := cmd.Flags().GetString("config")
		if err != nil {
			klog.Fatalf("Error in getting kubeconfig path")
		} else {
			delete(cfgFile)
		}
	},
}

func deletecommand() *cobra.Command {
	return mksPrDelete
}

func delete(cfgFile string) {

	cfg, err := clientcmd.BuildConfigFromFlags("", cfgFile)

	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}
	mksClient, err := vr.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error generating new clientset: %v", err)
	}

	delErr := mksClient.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Delete(context.TODO(), resourceName, v1.DeleteOptions{})
	if delErr != nil {
		klog.Fatalf("Delete MksPipelineRun failed!", err.Error())
	}
	fmt.Println("Mks PipelineRun ", resourceName, " deleted")

}

func init() {

	mksPrDelete.Flags().StringVar(&resourceName, "rn", "", "Name of MksPipelineRun Resource to be deleted")
	mksPrDelete.MarkFlagRequired("rn")
	mksPrDelete.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")

}
