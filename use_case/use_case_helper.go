package use_case

import (
	"capstone/constant"
	"capstone/delivery/dto_response"

	"golang.org/x/sync/errgroup"
)

var (
	panicIsPath    = true
	panicIsNotPath = false
)

func panicIfErr(err error, excludedErrs ...error) {
	if err != nil {
		for _, excludedErr := range excludedErrs {
			if err == excludedErr {
				return
			}
		}
		panic(err)
	}
}

func panicIfRepositoryError(err error, message string, isPath bool) {
	if err != nil {
		if err == constant.ErrNoData {
			if isPath {
				panic(dto_response.NewNotFoundResponse(message))
			}
			panic(dto_response.NewBadRequestResponse(message))
		}
		panic(err)
	}
}

func await(fn func(group *errgroup.Group)) error {
	group := new(errgroup.Group)
	fn(group)
	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}
