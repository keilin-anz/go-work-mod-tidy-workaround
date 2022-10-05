package main

import (
	"fmt"

	"go-work-mod-tidy-workaround/other"
	"go-work-mod-tidy-workaround/utils/math"

	"github.com/keilin-anz/go-work-mod-tidy-workaround/exposed"
)

func main() {
	result := other.ToYaml([]struct {
		Result int
	}{
		// Use our local-only module
		{Result: math.Add(1, 1)},
		// Use a local-only module which uses another local-only module
		{Result: other.AddTwenty(1)},
		// Use a module exposed for external import
		{Result: exposed.Add(2, 2)},
	})

	fmt.Printf("from cmd:\n%s\n", result)
	exposed.DoSomething()
}
