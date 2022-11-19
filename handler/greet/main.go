package main

import (
	"context"

	"github.com/caospinac/aws-sls-go/domain"
	"github.com/caospinac/aws-sls-go/util"
)

func greet(ctx context.Context, event domain.EventRequest) (util.Response, util.ApiError) {
	res := util.NewResponse().WithMessage("Hello, world!")

	return res, nil
}

func main() {
	util.Start(greet)
}
