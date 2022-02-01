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

var myFlags fl = fl{}

var mksTaskCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "ceates mkstask",
	Long:  "mkstask create is used to create",
	Run: func(cmd *cobra.Command, args []string) {
		var name string = ""
		if len(args) == 0 {
			klog.Fatalf("Name argument is required for your resource")
		} else {
			name = args[0]
		}
		createResource(cmd, name)
		fmt.Println(args, "mkstask create called")
	},
}

func createResource(cmd *cobra.Command, name string) {
	CfgFile, _ := cmd.Flags().GetString("config")
	cfg, err := clientcmd.BuildConfigFromFlags("", CfgFile)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}
	mksClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building mks clientset: %v", err)
	}
	mt := &v1alpha1.MksTask{
		TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
		ObjectMeta: v1.ObjectMeta{Name: name},
		Spec:       v1alpha1.MksTaskSpec{Name: myFlags.stepname, Image: myFlags.image, Command: myFlags.command, Args: myFlags.args},
	}
	namespace, _ := cmd.Flags().GetString("namespace")
	crt, err := mksClient.MkscontrollerV1alpha1().MksTasks(myFlags.namespace).Create(context.TODO(), mt, v1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error!!! Coldn't create the resource with name %s in the namespace %s\n", name, namespace)
		fmt.Errorf("Couldn't create mksTsk", err.Error())
	} else {
		fmt.Println(string(crt.UID))
	}
	fmt.Println(crt)
}

func MksTaskCreate() *cobra.Command {
	return mksTaskCreateCmd
}

func init() {
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.namespace, "namespace", "n", "default", "namespace of the mksTaskResource")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.stepname, "stepname", "s", "", "provide step name")
	mksTaskCreateCmd.MarkFlagRequired("sn")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.image, "image", "i", "", "name of the image resource")
	mksTaskCreateCmd.MarkFlagRequired("i")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.command, "command", "c", "", "commands")
	mksTaskCreateCmd.MarkFlagRequired("command")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.args, "args", "a", "", "arguments to the commmand")
	mksTaskCreateCmd.MarkFlagRequired("args")
}
