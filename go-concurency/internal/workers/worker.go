package workers

import (
	"fmt"
	"time"
)

func Work(name string, delay time.Duration) {
	fmt.Printf("%s started\n", name)
	time.Sleep(delay)
	fmt.Printf("%s finished\n", name)
}
