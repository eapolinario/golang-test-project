package main

import "fmt"

import "github.com/flyteorg/flyte/flytestdlib/errors"

const ErrRefreshingToken errors.ErrorCode = "TOKEN_REFRESH_FAILURE!"

func main() {
  fmt.Println(ErrRefreshingToken)
}
