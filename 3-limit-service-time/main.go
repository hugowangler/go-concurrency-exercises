//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	start := time.Now()
	if u.IsPremium {
		process()
		u.TimeUsed += int64(time.Since(start))
	} else {
		c := make(chan int)
		go func() {
			process()
			c <- 1
		}()

		select {
		case <-c:
			u.TimeUsed += int64(time.Since(start))
			return true
		case <-time.After((time.Second * 10) - time.Duration(u.TimeUsed)):
			return false
		}
	}
	return true
}

func main() {
	RunMockServer()
}
