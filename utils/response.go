package utils

func SuccessResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}

func FailResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}
