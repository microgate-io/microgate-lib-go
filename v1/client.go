package v1

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

// DialMicrogate returns a new connection.
// It block until is does and keeps trying if it fails.
func DialMicrogate() *grpc.ClientConn {
	addr := "localhost:9191"
	healthy := false
	for !healthy {
		ping, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			// cannot use microgate logging here, no connection to it!
			log.Println("waiting to create connection to ", addr, err)
			time.Sleep(5 * time.Second)
		} else {
			// we know it is there, close the check
			ping.Close()
			healthy = true
		}
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		// cannot use microgate logging here, no connection to it!
		log.Println("unable to create connection to ", addr, err)
	}
	log.Println("connected to microgate on", addr)
	return conn
}
