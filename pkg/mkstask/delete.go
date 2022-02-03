package mkstask

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var mksTaskDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes mkstask",
	Long:  "mkstask delete is used to delete resource",
	Run: func(cmd *cobra.Command, args []string) {
		var name string = ""
		if len(args) == 0 {
			klog.Fatalf("Name argument is required to identify your resource")
		} else {
			name = args[0]
		}
		deleteResource(cmd, name)
	},
}

func deleteResource(cmd *cobra.Command, name string) {
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
	er := mksClient.MkscontrollerV1alpha1().MksTasks(namespace).Delete(context.TODO(), name, v1.DeleteOptions{})
	if er != nil {
		fmt.Printf("Error!!! Coldn't delete the resource with name %s from the namespace %s\n", name, namespace)
		fmt.Errorf("Couldn't delete mksTsk", err.Error())
	} else {
		fmt.Println("Successively deleted")
	}
}

func MksTaskDelete() *cobra.Command {
	return mksTaskDeleteCmd
}

func init() {
	mksTaskDeleteCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")
}
