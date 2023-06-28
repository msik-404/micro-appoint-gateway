package employees 

import (
	"fmt"
	"os"
)

var (
	ConnString = fmt.Sprintf("%s:50051", os.Getenv("EMPLOYEES_HOSTNAME"))
)
