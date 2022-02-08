package mkstaskrun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func deleteMksTaskRun() *cobra.Command {
	cc := &cobra.Command{
		Use:   "delete",
		Short: "Delete a MkstaskRun in default namespace",
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
			delerr := mksclient.MkscontrollerV1alpha1().MksTaskRuns("default").Delete(context.TODO(), fs, metav1.DeleteOptions{})
			if delerr != nil {
				return nil
			}
			return nil
		},
	}
	cc.Flags().String("name", "", "Name of the taskrun to be deleted")
	return cc
}
