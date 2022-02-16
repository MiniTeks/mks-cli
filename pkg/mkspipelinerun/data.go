package mkspipelinerun

import (
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetTestData(prdata ...*v1alpha1.MksPipelineRun) []runtime.Object {
	var ro []runtime.Object
	for _, obj := range prdata {
		ro = append(ro, runtime.Object(obj))
	}
	return ro
}

var Prlist = []*v1alpha1.MksPipelineRun{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmpr1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "mprref1",
			},
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmpr2",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "mprref2",
			},
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmpr3",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "mprref3",
			},
		},
	},
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "testmpr4",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "mprref4",
			},
		},
	},
}

var Prdel = []*v1alpha1.MksPipelineRun{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "delmpr1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "deletedpipelineref",
			},
		},
	},
}

var Prget = []*v1alpha1.MksPipelineRun{
	{
		ObjectMeta: v1.ObjectMeta{
			Name:      "getmpr1",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "getpipelineref",
			},
		},
	},
}
