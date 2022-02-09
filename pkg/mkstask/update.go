package mkstask

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

var myFlags2 fl = fl{}

func MksTaskUpdate(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "updates mkstask resources ",
		Long:  "mkstask update is used to update",
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string = ""
			if len(args) == 0 {
				klog.Fatalf("Name argument is required for your resource")
			} else {
				name = args[0]
			}
			mt := &v1alpha1.MksTask{
				TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
				ObjectMeta: v1.ObjectMeta{Name: name},
				Spec:       v1alpha1.MksTaskSpec{Name: myFlags2.stepname, Image: myFlags2.image, Command: myFlags2.command, Args: myFlags2.args},
			}
			crt, err := mksclient.MkscontrollerV1alpha1().MksTasks(myFlags2.namespace).Update(context.TODO(), mt, v1.UpdateOptions{})
			if err != nil {
				fmt.Errorf("Couldn't update mksTsk", err.Error())
				return err
			}
			fmt.Println(crt)
			fmt.Println(args, "mkstask update called")
			return nil
		},
	}
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.namespace, "namespace", "n", "default", "namespace of the mksTaskResource")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.stepname, "stepname", "s", "", "provide step name")
	mksTaskUpdateCmd.MarkFlagRequired("sn")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.image, "image", "i", "", "name of the image resource")
	mksTaskUpdateCmd.MarkFlagRequired("i")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.command, "command", "c", "", "commands")
	mksTaskUpdateCmd.MarkFlagRequired("command")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.args, "args", "a", "", "arguments to the commmand")
	mksTaskUpdateCmd.MarkFlagRequired("args")
	return mksTaskUpdateCmd
}
