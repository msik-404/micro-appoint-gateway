package users 

import (
	"fmt"
	"os"
)

var (
	ConnString = fmt.Sprintf("%s:50051", os.Getenv("USERS_HOSTNAME"))
)
