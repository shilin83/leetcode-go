package helpers

var (
	MaxInt = 1<<31 - 1
	MinInt = -1 << 31
)

func IF[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}

	return falseValue
}
