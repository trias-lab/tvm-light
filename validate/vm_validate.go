package validate

import (
	"tvm-light/proto/tm"
)

func RequestValidate(request *tm.ExecuteContractRequest) (isCorect bool, err error) {
	return true, nil;
}
