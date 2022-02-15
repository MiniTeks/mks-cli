package mkstaskrun

import (
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetTestData() []runtime.Object {
	var ro []runtime.Object
	for _, obj := range trdata {
		ro = append(ro, runtime.Object(obj))
	}
	return ro
}

var trdata = []*v1alpha1.MksTaskRun{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmtr1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "mtrref1",
			},
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmtr2",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "mtrref2",
			},
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmtr3",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "mtrref3",
			},
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmtr4",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "mtrref4",
			},
		},
	},
}
