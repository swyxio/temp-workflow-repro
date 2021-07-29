## temporal go sdk wrong error repro

this repo is meant to show how the go sdk offers the wrong error message when you run it.


![CleanShot 2021-07-28 at 13 44 54@2x](https://user-images.githubusercontent.com/6764957/127393422-05d0020f-59e4-447d-ae6d-f12c46de2188.png)

## instructions

1. have temporal server running (https://docs.temporal.io/docs/server/quick-install)
2. ` go run worker/worker.go`
3. `go run starter/starter.go`
