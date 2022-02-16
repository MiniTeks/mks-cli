package mkstask

import (
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetTestData(tdata ...*v1alpha1.MksTask) []runtime.Object {
	var ro []runtime.Object
	for _, obj := range tdata {
		ro = append(ro, runtime.Object(obj))
	}
	return ro
}

var Tlist = []*v1alpha1.MksTask{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testt1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask1",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testt2",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask2",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testt3",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask3",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testt4",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask4",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	},
}

var Tdel = []*v1alpha1.MksTask{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "delmt1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "deletedtaskstep",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	},
}

var Tget = []*v1alpha1.MksTask{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "getmt1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "gettaskstep",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	},
}
