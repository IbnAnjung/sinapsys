package http

const RequestIdContextKey = "RequestID"
const UserIdContextKey = "UserID"
const UserIsPremiumContextKey = "UserIsPremium"

func GetStandartSuccessResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}
