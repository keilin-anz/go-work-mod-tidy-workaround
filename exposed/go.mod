module github.com/keilin-anz/go-work-mod-tidy-workaround/exposed

go 1.19

replace github.com/keilin-anz/go-work-mod-tidy-workaround/other => ../internal/other

replace github.com/keilin-anz/go-work-mod-tidy-workaround/utils => ../internal/utils

require (
	github.com/keilin-anz/go-work-mod-tidy-workaround/other v0.0.0-00010101000000-000000000000
	github.com/keilin-anz/go-work-mod-tidy-workaround/utils v0.0.0-00010101000000-000000000000
)

require gopkg.in/yaml.v3 v3.0.1 // indirect
