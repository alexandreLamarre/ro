# ro

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-%23007d9c)
![Release](https://img.shields.io/github/v/release/alexandreLamarre/ro)
[![GoDoc](https://godoc.org/github.com/alexandreLamarre/ro?status.svg)](https://pkg.go.dev/github.com/alexandreLamarre/ro)
[![Lint](https://github.com/alexandreLamarre/ro/actions/workflows/lint.yaml/badge.svg)](https://github.com/alexandreLamarre/ro/actions/workflows/lint.yaml)
[![Test](https://github.com/alexandreLamarre/ro/actions/workflows/test.yaml/badge.svg)](https://github.com/alexandreLamarre/ro/actions/workflows/test.yaml)
[![License](https://img.shields.io/github/license/alexandreLamarre/ro)](./LICENSE)


`alexandreLamarre/ro` (short for "range-over") is a generic iterator Go library based on Go 1.22+ [rangefunc](https://go.dev/wiki/RangefuncExperiment) experiment.

The library is built for composability and readibility while providing similar functionality to other language's interator libraries such as [itertools](https://docs.python.org/3/library/itertools.html)

This library is intended to be paired with [`samber/lo`](https://github.com/samber/lo), where applicable.

## 🚀 Install

```sh
go get github.com/alexandreLamarre/ro@v0.1.0
```

## ✔️ Requirements

This library requires you to enable the go experimental feature [rangefunc](https://go.dev/wiki/RangefuncExperiment) :

```sh
export GOEXPERIMENT=rangefunc
```

## 💡 Usage

You can import `ro` using:

```go
import (
    "github.com/alexandreLamarre/ro"
)
```


## 👀 Example

```go
it := ro.Drop(
    ro.Apply(
        ro.Limit(
            ro.Permutations(
                ro.ToSlice(
                    ro.Range(0, 6, 1),
                ),
                5,
            ),
            5,
        ),
        func(perm []int) int {
            return digitsToInt(perm)
        },
    ),
    func(i int) bool {
        return i > 10000
    },
)
for v := range it {
    fmt.Println(v)
}
```

yields the following output:
```sh
1234
2134
```

The above iterator yields all numbers < 10000 that are formed from the digits of the first 5 generated permutations of {0,1,2,3,4,5} of size 5