package request

import (
	"github.com/test-instructor/cheetah/server/model/common/request"
	"github.com/test-instructor/cheetah/server/model/interfacecase"
)

type ApiCaseSearch struct {
	interfacecase.ApiCase
	request.PageInfo
	FrontCase bool `json:"front_case" form:"front_case"`
}
