package mkstask

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func MksTaskDelete(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "deletes mkstask",
		Long:  "mkstask delete is used to delete resource",
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string = ""
			if len(args) == 0 {
				klog.Fatalf("Name argument is required to identify your resource")
			} else {
				name = args[0]
			}
			namespace, _ := cmd.Flags().GetString("namespace")
			er := mksclient.MkscontrollerV1alpha1().MksTasks(namespace).Delete(context.TODO(), name, v1.DeleteOptions{})
			if er != nil {
				fmt.Printf("Error!!! Coldn't delete the resource with name %s from the namespace %s\n", name, namespace)
				fmt.Errorf("Couldn't delete mksTsk", er.Error())
				return er
			} else {
				fmt.Println("Successively deleted")
			}
			return nil
		},
	}
	mksTaskDeleteCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")

	return mksTaskDeleteCmd
}
