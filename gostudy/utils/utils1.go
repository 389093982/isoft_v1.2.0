package utils

func GetOpratorResult(n1 float64, n2 float64, operator byte) (string, float64) {
	switch operator {
	case '+':
		return "ok", n1 + n2
	case '-':
		return "ok", n1 - n2
	case '*':
		return "ok", n1 * n2
	case '/':
		return "ok", n1 / n2
	case '%':
		return "ok", float64(int64(n1) % int64(n2))
	default:
		return "error", 0
	}
}
