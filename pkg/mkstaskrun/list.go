package mkstaskrun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func listMksTaskRun() *cobra.Command {
	cc := &cobra.Command{
		Use:   "list",
		Short: "List all MkstaskRuns in default namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			CfgFile, _ := cmd.Flags().GetString("config")

			cfg, err := clientcmd.BuildConfigFromFlags("", CfgFile)
			if err != nil {
				return err
			}
			mksclient, err := versioned.NewForConfig(cfg)
			if err != nil {
				return nil
			}
			objlist, err := mksclient.MkscontrollerV1alpha1().MksTaskRuns("default").List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				return nil
			}
			for i, obj := range objlist.Items {
				fmt.Println(i+1, obj.GetName())
			}
			return nil
		},
	}
	return cc
}
