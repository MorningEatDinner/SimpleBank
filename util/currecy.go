package util

// 加入的所有支持的货币种类
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrecy(currency string) bool {
	//返回是否是支持的货币类型
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
