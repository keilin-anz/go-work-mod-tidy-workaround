module go-work-mod-tidy-workaround/other

go 1.19

require (
	github.com/keilin-anz/go-work-mod-tidy-workaround/utils v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/keilin-anz/go-work-mod-tidy-workaround/utils => ../utils
