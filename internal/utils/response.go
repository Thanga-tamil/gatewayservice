package utils

func Bad(statusCode int, msg string) any {
	return map[string]any{"status": statusCode, "message": msg}
}

func Success(statusCode int, data any, msg string) any {
	return map[string]any{"status": statusCode, "data": data,  "message": msg}
}
