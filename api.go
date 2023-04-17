package main

import (
	"fmt"
	"net"
	"strings"
)

func showApi() string {
	ip := apiurl()
	str := fmt.Sprintf(`
API 地址与说明:
-------------------------------
news:		%s/assets/news/1.html
images:		%s/assets/images/1.jpg
newsAPI:	%s/wp-json/nautica/news
newsAPI:	%s/wp-json/nautica/videos
faqAPI:		%s/wp-json/nautica/faq
bannerAPI:	%s/wp-json/nautica/banner
securityAPI:	%s/wp-json/nautica/security	
solutionsAPI:	%s/wp-json/nautica/solutions
`, ip, ip, ip, ip, ip, ip, ip, ip)
	return str
}

func apiurl() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localconn := conn.LocalAddr().(*net.UDPAddr).String()
	ip := strings.Split(localconn, ":")[0]
	return "http://" + ip + ":8080"
}
