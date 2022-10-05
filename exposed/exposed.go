package exposed

import (
	"fmt"
	"go-work-mod-tidy-workaround/other"
	"go-work-mod-tidy-workaround/utils/math"
)

// Try exposing an internal module func through this public module
var Add = math.Add

func DoSomething() {
	fmt.Println(other.ToYaml(struct {
		Message string
		Result  int
	}{
		Message: "this came from the exposed module",
		Result:  Add(20, 20),
	}))
}
