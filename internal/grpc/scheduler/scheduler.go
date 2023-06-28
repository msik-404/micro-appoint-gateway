package scheduler

import (
	"fmt"
	"os"
)

var (
	ConnString = fmt.Sprintf("%s:50051", os.Getenv("SCHEDULER_HOSTNAME"))
)
