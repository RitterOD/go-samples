package randomstr

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomStr(n int) string {
	fmt.Printf("random string with %d", n)
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:n]
}
