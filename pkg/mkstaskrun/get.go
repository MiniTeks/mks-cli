package mkstaskrun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func getMksTaskRun() *cobra.Command {
	cc := &cobra.Command{
		Use:   "get",
		Short: "Get a MkstaskRun in default namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fs, _ := cmd.Flags().GetString("name")
			if fs == "" {
				fmt.Errorf("TaskName not defined")
			}
			CfgFile, _ := cmd.Flags().GetString("config")

			cfg, err := clientcmd.BuildConfigFromFlags("", CfgFile)
			if err != nil {
				return err
			}
			mksclient, err := versioned.NewForConfig(cfg)
			if err != nil {
				return nil
			}
			obj, err := mksclient.MkscontrollerV1alpha1().MksTaskRuns("default").Get(context.TODO(), fs, metav1.GetOptions{})
			if err != nil {
				return nil
			}
			fmt.Println(obj)
			return nil
		},
	}
	cc.Flags().String("name", "", "Name of the taskrun to be deleted")
	return cc
}
