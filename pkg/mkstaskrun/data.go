package mkstaskrun

import (
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetTestData(trdata ...*v1alpha1.MksTaskRun) []runtime.Object {
	var ro []runtime.Object
	for _, obj := range trdata {
		ro = append(ro, runtime.Object(obj))
	}
	return ro
}

var Trlist = []*v1alpha1.MksTaskRun{
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

var Trdel = []*v1alpha1.MksTaskRun{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "delmtr1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "deletedtaskref",
			},
		},
	},
}

var Trget = []*v1alpha1.MksTaskRun{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "getmtr1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "gettaskref",
			},
		},
	},
}
