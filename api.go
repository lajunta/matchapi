package main

import (
	"fmt"
	"net"
	"strings"
)

func showApi() string {
	ip := apiurl()
	str := fmt.Sprintf(`
	Links:
  ---------------------------------------------
	news:		%s/assets/news/1.html
	images:		%s/assets/images/1.jpg
	newsAPI:	%s/wp-json/cyberbuf/news
	bannerAPI:	%s/wp-json/cyberbuf/banner
	securityAPI:	%s/wp-json/cyberbuf/security
	solutionsAPI:	%s/wp-json/cyberbuf/solutions
	`, ip, ip, ip, ip, ip, ip)
	return str
}

func apiurl() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localconn := conn.LocalAddr().(*net.UDPAddr).String()
	ip := strings.Split(localconn, ":")[0]
	return "http://" + ip + ":8080"
}
