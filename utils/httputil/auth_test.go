package httputil

func getMockAccount() map[string]string {
	accounts := make(map[string]string)
	accounts["admin"] = "admin"
	accounts["user"] = "user"
	return accounts
}
