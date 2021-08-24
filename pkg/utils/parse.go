package utils

//
//func ParseUrl(host string, port string) string {
//	if port == "80" {
//		return "http://" + host
//	} else if port == "443" {
//		return  "https://" + host
//	} else if len(regexp.MustCompile("443").FindAllStringIndex(port, -1)) == 1 {
//		return "https://" + host + ":" + port
//	} else {
//		return "https://" + host + ":" + port
//	}
//}

func ParseUrl(host string, port string) string {
	if port == "443" {
		return "https://" + host
	} else {
		return "http://" + host + ":" + port
	}
}