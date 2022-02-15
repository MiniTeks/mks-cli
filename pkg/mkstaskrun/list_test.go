package mkstaskrun

import (
	"fmt"
	"testing"

	"github.com/MiniTeks/mks-cli/pkg/test"
)

func TestList(t *testing.T) {
	fc := &test.FakeMksParam{}
	fc.SetNamespace("default")
	fc.SetTestObjects(GetTestData()...)
	cs, _ := fc.Client(nil)

	tr := Command(cs)
	out, err := test.ExecuteCommand(tr, "list")
	fmt.Println(out)
	if err != nil {
		t.Fatalf("Cannot execute command: %v", err)
	} else if out != "1 testmtr1\n2 testmtr2\n3 testmtr3\n4 testmtr4\n" {
		t.Fatal("Cant find taskrun")
	}
}
