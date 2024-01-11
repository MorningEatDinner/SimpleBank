package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/xiaorui/simplebank/util"
)

var validCurrecy validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currecy, ok := fieldLevel.Field().Interface().(string); ok {
		//检查这个currecy是否是支持的类型
		return util.IsSupportedCurrecy(currecy)
	}

	return false
}
