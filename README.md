
```shell

go test github.com/chyroc/gotestexample/a -coverprofile=coverage.out -covermode=count -coverpkg=./...

go test github.com/chyroc/gotestexample/b -coverprofile=coverage.out.2 -covermode=count -coverpkg=./...

tail -n +2 coverage.out.2 >> coverage.out

go tool cover -html coverage.out -o coverage.html

```

![](./assets/fc8ff74c-1b7d-4ed5-926b-fa3f7c3bb0f0.png)

![](./assets/8cdf2a1d-f4f0-4d69-9aef-b022aa2f5e66.png)