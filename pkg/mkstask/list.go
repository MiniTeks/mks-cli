package mkstask

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var mksTaskListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists mkstasks",
	Long:  "mkstask list is used to list",
	Run: func(cmd *cobra.Command, args []string) {
		ListResource(cmd)
	},
}

func ListResource(cmd *cobra.Command) {
	CfgFile, _ := cmd.Flags().GetString("config")
	cfg, err := clientcmd.BuildConfigFromFlags("", CfgFile)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}
	mksClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building mks clientset: %v", err)
	}
	namespace, _ := cmd.Flags().GetString("namespace")
	crt, err := mksClient.MkscontrollerV1alpha1().MksTasks(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Error!!! Coldn't get the resource(s) from the namespace %s\n", namespace)
		fmt.Errorf("Couldn't create mksTsk", err.Error())
	}
	fmt.Printf("Here are resources in the namespace %s\n", namespace)
	displayCrtList(crt)
}

func displayCrtList(crt *v1alpha1.MksTaskList) {
	for _, item := range crt.Items {
		fmt.Print("\n")
		displayCrt(&item)
	}
}

func MksTaskList() *cobra.Command {
	return mksTaskListCmd
}

func init() {
	mksTaskListCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")
}
