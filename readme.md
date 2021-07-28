## temporal go sdk wrong error repro

this repo is meant to show how the go sdk offers the wrong error message when you run it.


![CleanShot 2021-07-28 at 13 44 54@2x](https://user-images.githubusercontent.com/6764957/127393422-05d0020f-59e4-447d-ae6d-f12c46de2188.png)

## instructions

1. have temporal server running
2. ` go run worker/worker.go`
3. `go run starter/starter.go`

when running you should see this error: `PanicError reflect: Call with too few input arguments PanicStack activity for my-queue [panic]`

This happens because `MyActivity` uses `workflow.Context` as its first argument, instead of Go's native `context.Context`. Once you fix this, the code runs as expected.

The problem here is that the error `Call with too few input arguments` is incorrect/misleading.
