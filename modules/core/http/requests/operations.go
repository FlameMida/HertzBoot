package requests

import (
	"HertzBoot/app/request"
	"HertzBoot/modules/core/entities"
)

type OperationsSearch struct {
	entities.Operations
	request.PageInfo
}
