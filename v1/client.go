package v1

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

// DialMicrogate returns a new connection.
// It block until is does and keeps trying if it fails.
func DialMicrogate() *grpc.ClientConn {
	addr := "localhost:9191"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	for err != nil {
		// cannot use microgate logging here, no connection to it!
		log.Println("failed to connect to ", addr, err)
		time.Sleep(1 * time.Second)
	}
	log.Println("connected to microgate on :9191")
	return conn
}
