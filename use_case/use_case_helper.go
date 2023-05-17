package use_case

import (
	"capstone/constant"
	"capstone/delivery/dto_response"
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
