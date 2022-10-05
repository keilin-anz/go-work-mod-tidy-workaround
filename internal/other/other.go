package other

import (
	"go-work-mod-tidy-workaround/utils/math"

	// YAML library Included purely to check if `go mod tidy` works as expected
	"gopkg.in/yaml.v3"
)

func ToYaml(i any) string {
	result, err := yaml.Marshal(i)
	if err != nil {
		return "ceci ñ'est pas une erreur"
	}
	return string(result)
}

func AddTwenty(i int) int {
	return math.Add(i, 20)
}
