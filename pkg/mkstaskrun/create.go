package mkstaskrun

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createMksTaskRun(mksclient *versioned.Clientset) *cobra.Command {
	cc := &cobra.Command{
		Use:   "create",
		Short: "Create a MksTaskRun in default namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("MksTaskRun called")
			fs, _ := cmd.Flags().GetString("taskRef")
			if fs == "" {
				fmt.Errorf("taskRef not defined")
			}
			cmtr := &v1alpha1.MksTaskRun{
				TypeMeta:   metav1.TypeMeta{Kind: "MksTaskRun"},
				ObjectMeta: metav1.ObjectMeta{Name: "mkstaskrun" + fs},
				Spec: v1alpha1.MksTaskRunSpec{
					TaskRef: v1alpha1.MksTaskRef{
						Name: fs,
					},
				},
			}
			obj, err := mksclient.MkscontrollerV1alpha1().MksTaskRuns("default").Create(context.TODO(), cmtr, metav1.CreateOptions{})
			if err != nil {
				return nil
			}
			fmt.Println(obj)
			return nil
		},
	}
	cc.Flags().String("taskRef", "", "Task name to be executed")
	return cc
}
