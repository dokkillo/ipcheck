package main
import (
    "fmt"
	"os"
	"net"
)

func main() {

	_, subnet, _ := net.ParseCIDR(os.Args[1])
	ip := net.ParseIP(os.Args[2])
	if subnet.Contains(ip) {
    	fmt.Println(ip)
	}
}