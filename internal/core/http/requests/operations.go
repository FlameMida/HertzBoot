package requests

import (
	"HertzBoot/internal/core/entities"
	"HertzBoot/pkg/request"
)

type OperationsSearch struct {
	entities.Operations
	request.PageInfo
}
