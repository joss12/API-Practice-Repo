package scanner

import (
	"net"
	"time"
)

func ScanPorts(host string) []int {
	openPorts := []int{}
	portsToCheck := []int{
		21, 22, 23, 25, 53, 80, 110, 143, 443, 3306, 6379, 8080, 8443,
	}

	// Scan ports in parallel
	results := make(chan int, len(portsToCheck))

	for _, port := range portsToCheck {
		go func(p int) {
			address := net.JoinHostPort(host, itoa(p))
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err == nil {
				conn.Close()
				results <- p
			} else {
				results <- -1
			}
		}(port)
	}

	// Collect open ports
	for i := 0; i < len(portsToCheck); i++ {
		p := <-results
		if p != -1 {
			openPorts = append(openPorts, p)
		}
	}

	return openPorts
}

// convert int to string without strconv
func itoa(i int) string {
	return string(rune('0'+i/10)) + string(rune('0'+i%10))
}
