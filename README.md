# Discussion 31116 minimal reproduction

Reproduction for [Renovate discussion 31116](https://github.com/renovatebot/renovate/discussions/31116)

## Current behavior

The repository has 2 direct dependencies which are both left intentionally on an older version
* `github.com/julienschmidt/httprouter v1.2.0` (current version v1.3.0)
* `gocloud.dev v0.37.0` (current version v0.39.0)


:x: Renovate fails to detect the direct dependency `gocloud.dev` and therefore doesn't create a Pull Request to update it.

:white_check_mark: It successfully detects the direct dependency to `github.com/julienschmidt/httprouter v1.2.0` and also creates a pull request to update it.

## Expected behavior

Renovate should detect the direct dependency to `gocloud.dev v0.37.0` and create a pull request to update it to the current minor version `v0.39.0`.
