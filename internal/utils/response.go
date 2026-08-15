package utils

func Bad(statusCode int, msg string) any {
	return map[string]any{"status": statusCode, "message": msg}
}
