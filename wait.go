package golib

import (
	"net"
	"time"
)

func WaitTCP(addr string, d time.Duration) bool {
	log.Printf("Waiting for TCP to be available at %s", addr)
	if d < 10*time.Second {
		d = 1 * time.Minute
	}
	// Try once a second to connect
	for startTime := time.Now(); time.Since(startTime) < d; time.Sleep(3 * time.Second) {
		conn, err := net.DialTimeout("tcp", addr, time.Second)

		if err == nil {
			// Connection successful
			log.Printf("TCP came up on %s", addr)
			closeErr := conn.Close()
			if closeErr != nil {
				log.Printf("Error closing TCP connection in waitTCP: %s", closeErr)
			}

			return true
		}

		log.Printf("Tried to connect to %s, got error: %s. Will retry in 1 second.", addr, err)
	}
	// Timed out
	return false
	// panic(fmt.Sprintf("Timeout out waiting for service to start on %s", addr))
}
