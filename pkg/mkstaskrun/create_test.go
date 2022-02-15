package mkstaskrun

import (
	"testing"

	"github.com/MiniTeks/mks-cli/pkg/test"
)

func TestCreate(t *testing.T) {
	fc := &test.FakeMksParam{}
	fc.SetNamespace("default")
	fc.SetTestObjects(GetTestData()...)
	cs, _ := fc.Client(nil)

	tr := Command(cs)
	_, err := test.ExecuteCommand(tr, "create", "--taskRef=testcreate")
	out, _ := test.ExecuteCommand(tr, "get", "--name=mkstaskruntestcreate")
	if err != nil {
		t.Fatalf("Cannot execute command: %v", err)
	} else if out == "" {
		t.Fatal("Cant create taskrun")
	}
}
