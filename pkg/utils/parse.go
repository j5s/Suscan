package utils

func ParseUrl(host string, port string) string {
	if port == "443" {
		return "https://" + host
	} else {
		return "http://" + host + ":" + port
	}
}