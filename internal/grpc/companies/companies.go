package companies

import (
	"fmt"
	"os"
)

var (
	ConnString = fmt.Sprintf("%s:50051", os.Getenv("COMPANIES_HOSTNAME"))
)
