# Testing options to workaround `go mod tidy` issue in workspaces

See the [original issue here](https://github.com/golang/go/issues/50750)

This repo consitutes a minimal example of the *closest I've gotten* to making `go mod tidy` not attempt to download my locally-defined modules.

The hope was to find an adequate workaround pattern, but I have not.  However this approach *does work* for some limited use cases.

> :warning: Caveat: my use case is very specific, so this will likely not fit your needs


- [Testing options to workaround `go mod tidy` issue in workspaces](#testing-options-to-workaround-go-mod-tidy-issue-in-workspaces)
- [TLDR](#tldr)
  - [My requirements](#my-requirements)
  - [Approach](#approach)
    - [Use cases included here](#use-cases-included-here)
  - [Conclusion](#conclusion)
    - [What else?](#what-else)
      - [Alternative approaches](#alternative-approaches)
  - [Other branches (alternatives](#other-branches-alternatives)
    - [Opinion](#opinion)
# TLDR

Go Workspaces are not an adequate replacement for `go mod edit -replace`, and they don't even seem to play well together.

`go mod tidy` will not behave properly for modules which have dependencies on other modules within the same workspace (it will attempt to download the latest published copy instead)

You can work around `go mod tidy`'s promsicuous downloading behaviour by using invalid URLs as module names, but that brings caveats (see rest of this document)

See more in [Conclusion](#conclusion)

## My requirements

- monorepo with multiple modules (obviously)
- want modules to depend upon each other (segregation of duty and requirements)
  - and this is basically **all** I want to use Go Workspaces for
- :warning: **IMPORTANT** releases of the modules are in lock-step, ie. I will release new versions of _all modules_ at once
  - This means I don't care about the scenario mentioned [here](https://github.com/golang/go/issues/50750#issuecomment-1194322256)
  - I want locally defined modules to Just Work™ (and never try to download a remote version)
- :warning: **IMPORTANT** this isn't really a requirement, but a limitation to this workaround
  - Any shared modules cannot be imported from other repositories see [utils/README](./utils/README.md) for more details


## Approach

1. If you specify your workspace module names as a URL (like you're supposed to), then `go mod tidy` gets promiscuous and attempts to download them
2. Solution: just use non-URL module names for any modules which need local import, ie.
   - instead of `github.com/keilin-anz/go-work-mod-tidy-workaround/utils`
   - use `go-work-mod-tidy-workaround/utils`

### Use cases included here

| Scenario                                                                              | link                         | `go mod tidy`                          | `go work sync` | import from another repo                                                                                                                              |
| ------------------------------------------------------------------------------------- | ---------------------------- | -------------------------------------- | -------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| A `cmd` with no `go.mod` which depends on internal modules                            | See [here](./cmd)            | **doesn't work** (no `go.mod`)         | **works**      | **doesn't work** (as expected)                                                                                                                        |
| An **internal** module which has no dependencies                                      | See [here](./internal/math)  | **works** - obviously, no dependencies | **works**      | **doesn't work** <br><br>(module is downloaded, but go detects a mismatch between internal module name and the public one)                            |
| An **internal** module which has external and internal dependencies                   | See [here](./internal/other) | **works!**                             | **works**      | Same as above                                                                                                                                         |
| An **exposed** (importable externally) module with external and internal dependencies | See [here](./exposed)        | **works**                              | **works**      | **ALMOST works** <br><br>Module downloads, module is correctly included (as expected), but then module fails to build *despite building fine locally* |

## Conclusion

It's *almost* a workaround, but ultimately only useful for monorepo's where any shared functionality isn't expected to be downloaded by other repositories

You **CAN:**

- Happily segregate local modules and their external dependencies
- Freely import said local modules and use them
  - in "module-less" go files like in [cmd](./cmd)
  - other local modules like in [internal/other](./internal/other)
  - modules which have fully-formed URL names like in [exposed](./exposed)
- ***almost*** use `go work` as a replacement for `go mod edit -replace` *BUT NOT QUITE*

You **CAN'T:**

- Expose any modules which have a dependency on the internal modules
- Trust that the "magic" behaviour of happily skipping download on non-URL imports in `go mod tidy` won't change and break this
- Have your cake and eat it to

### What else?

You may be able to get a more sane use case working, eg. with:

- modules that can actually be imported elsewhere
- modules which can internally depend upon local changes rather than using external

**BUT** it appears to require a fair amount of manual effort, something like:

- add `replace` clauses to all of your modules which depend on each other (and *don't forget to update them if you change any imports! :fire:*)
- cross your fingers that doesn't break `go work sync`
- `EDIT:` Nope - that doesn't seem to work
  - Fine locally
  - Breaks on external import (just in a different way):
    ```
    go get github.com/keilin-anz/go-work-mod-tidy-workaround/exposed@with-replace-and-url-names
    go: downloading github.com/keilin-anz/go-work-mod-tidy-workaround/exposed v0.0.0-20221005103718-e629069e8429
    go: downloading github.com/keilin-anz/go-work-mod-tidy-workaround v0.0.0-20221005103718-e629069e8429
    go: downloading github.com/keilin-anz/go-work-mod-tidy-workaround/utils v0.0.0-00010101000000-000000000000
    go: downloading github.com/keilin-anz/go-work-mod-tidy-workaround/other v0.0.0-00010101000000-000000000000
    github.com/keilin-anz/go-work-mod-tidy-workaround/exposed imports
    	github.com/keilin-anz/go-work-mod-tidy-workaround/other: github.com/keilin-anz/go-work-mod-tidy-workaround/other@v0.0.0-00010101000000-000000000000: invalid version: unknown revision 000000000000
    github.com/keilin-anz/go-work-mod-tidy-workaround/exposed imports
    	github.com/keilin-anz/go-work-mod-tidy-workaround/utils/math: github.com/keilin-anz/go-work-mod-tidy-workaround/utils@v0.0.0-00010101000000-000000000000: invalid version: unknown revision 000000000000

#### Alternative approaches


## Other branches (alternatives

- [with-replace-and-url-names](https://github.com/keilin-anz/go-work-mod-tidy-workaround/tree/with-replace-and-url-names) - attempt at using legit URL names and the `replace` directive inside individual `go.mod` files
  - This is not an optimal solution since it requires a lot of manual updates, but figured it worth trying out
  - it **did not work**
    ```bash
    go get github.com/keilin-anz/go-work-mod-tidy-workaround/exposed@with-replace-and-url-names
    ```
    and you get
    ```
    github.com/keilin-anz/go-work-mod-tidy-workaround/exposed imports
  	github.com/keilin-anz/go-work-mod-tidy-workaround/other: github.com/keilin-anz/go-work-mod-tidy-workaround/other@v0.0.0-00010101000000-000000000000: invalid version: unknown revision 000000000000
  github.com/keilin-anz/go-work-mod-tidy-workaround/exposed imports
  	github.com/keilin-anz/go-work-mod-tidy-workaround/utils/math: github.com/keilin-anz/go-work-mod-tidy-workaround/utils@v0.0.0-00010101000000-000000000000: invalid version: unknown revision 000000000000
    ```
- [with-replace-in-go-work](https://github.com/keilin-anz/go-work-mod-tidy-workaround/tree/with-replace-in-go-work) - attempt at achieving similar to the above but by using `replace` directives in the `go.work` file
  - Requires you to specify a version
  - That version is then not recognised when importing from outside, eg. try:
    ```bash
    $ go get github.com/keilin-anz/go-work-mod-tidy-workaround/exposed@with-replace-in-go-work

    go: downloading github.com/keilin-anz/go-work-mod-tidy-workaround v0.0.0-20221005105612-75af353f75f5
    go: downloading github.com/keilin-anz/go-work-mod-tidy-workaround/exposed v0.0.0-20221005105612-75af353f75f5
    github.com/keilin-anz/go-work-mod-tidy-workaround/exposed imports
    	github.com/keilin-anz/go-work-mod-tidy-workaround/other: cannot find module providing package github.com/keilin-anz/go-work-mod-tidy-workaround/other
    github.com/keilin-anz/go-work-mod-tidy-workaround/exposed imports
    	github.com/keilin-anz/go-work-mod-tidy-workaround/utils/math: cannot find module providing package github.com/keilin-anz/go-work-mod-tidy-workaround/utils/math
    ```

### Opinion

- `go.work` should be a proper replacement for using `go mod edit -replace`, it's too close to good DX not to
- If the decision is made not to solve [issues like this](https://github.com/golang/go/issues/50750) then:
  - an appropriate pattern should be documented
  - documentation should stop saying "go.work can be used instead of adding replace directives to work across multiple modules." as noted [here](https://github.com/golang/go/issues/50750#issuecomment-1236104735)
