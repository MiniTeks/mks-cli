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

var mksTaskGetCmd = &cobra.Command{
	Use:   "get",
	Short: "gets mkstasks",
	Long:  "mkstask get is used to get",
	Run: func(cmd *cobra.Command, args []string) {
		var name string = ""
		if len(args) == 0 {
			klog.Fatalf("A Name argument is required to get your resource")
		} else {
			name = args[0]
		}
		getResorce(cmd, name)
		// fmt.Println(args, "mkstask get called")
	},
}

func getResorce(cmd *cobra.Command, name string) {
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
	crt, err := mksClient.MkscontrollerV1alpha1().MksTasks(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		fmt.Printf("Error!!! Coldn't get any resource with name %s inside %s\n", name, namespace)
		klog.Fatal(err.Error())
	}
	fmt.Println("Here is your requested resource")
	displayCrt(crt)
}

func displayCrt(crt *v1alpha1.MksTask) {
	fmt.Println("name : ", crt.Name)
	fmt.Println("namespace : ", crt.Namespace)
	fmt.Println("spec : ", crt.Spec)
}

func MksTaskGet() *cobra.Command {
	return mksTaskGetCmd
}

func init() {
	mksTaskGetCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")
}
