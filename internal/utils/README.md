# A fake `utils` module

This will work for *internal use only* (this has nothing to do with it being in the `/internal` directory)

Attempting to import this module from other repositories will result in the following error:
```bash
go mod tidy
# ... snip ...
  github.com/keilin-anz/go-work-mod-tidy-workaround/utils/math: github.com/keilin-anz/go-work-mod-tidy-workaround/utils@v0.0.0-20221005091720-9ec42131e4a7: parsing go.mod:
  module declares its path as: foo/utils
          but was required as: github.com/keilin-anz/go-work-mod-tidy-workaround/utils
```

This is fine if you just want to have internal modules with their own independant dependencies

(But entirely unusable otherwise)
