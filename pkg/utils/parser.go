package utils

import "strconv"

func ParseUint(value interface{}) uint {
	switch v := value.(type) {
	case string:
		val, _ := strconv.Atoi(v)
		if val >= 0 {
			return uint(val)
		}
		return 0
	case float64:
		if v >= 0 {
			return uint(v)
		}
		return 0
	case int:
		if v >= 0 {
			return uint(v)
		}
		return 0
	case uint:
		return v
	default:
		return 0
	}
}