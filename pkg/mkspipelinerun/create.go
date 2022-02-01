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

var mksPrCreate = &cobra.Command{
	Use:   "create",
	Short: "Create PipelineRuns",
	Annotations: map[string]string{
		"commandType": "main",
	},
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func createcommand() *cobra.Command {
	return mksPrCreate
}

func create() {

	cfg, err := clientcmd.BuildConfigFromFlags("", cfgFile)

	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}
	mksClient, err := vr.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error generating new clientset: %v", err)
	}

	deployment := &v1alpha1.MksPipelineRun{
		TypeMeta:   v1.TypeMeta{Kind: "MksPipelineRun"},
		ObjectMeta: v1.ObjectMeta{Name: resourceName},
		Spec:       v1alpha1.MksPipelineRunSpec{PipelineRef: v1alpha1.MksPipelineRunRef{Name: pipelineRunRef}},
	}

	crt, err := mksClient.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Create(context.TODO(), deployment, v1.CreateOptions{})
	if err != nil {
		klog.Fatalf("Create MksPipelineRun failed!", err.Error())
	}
	fmt.Println("Mks Pipeline created with UID: ", crt.UID)

}

func init() {

	mksPrCreate.Flags().StringVar(&resourceName, "n", "", "Name of MksPipelineRun Resource")
	mksPrCreate.MarkFlagRequired("n")
	mksPrCreate.Flags().StringVar(&pipelineRunRef, "pr", "", "Name for PipelineRun Ref")
	mksPrCreate.MarkFlagRequired("pr")
	mksPrCreate.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")
}
