package middleware

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/constraints"
)

func IsProperString(value *string, maxLength int) (bool, error) {
	if value != nil {
		if len(*value) > int(maxLength) {
			return false, errors.New(fmt.Sprintf(
				"value should be shorter than %d",
				maxLength,
			))
		}
	}
	return true, nil
}

func IsProperInteger[T constraints.Integer](value *T, low T, high T) (bool, error) {
	if value != nil {
		if *value > high || *value <= low {
			return false, errors.New(fmt.Sprintf(
				"value should be smaller than %d and greater than %d",
				high,
				low,
			))
		}
	}
	return true, nil
}

func IsProperObjectIDHex(hex string) (bool, error) {
	if len(hex) != 24 {
		return false, errors.New("This is not proper hex value for objectID")
	}
	return true, nil
}

func GetNPerPageValue(c *gin.Context) (int64, error) {
	query := c.DefaultQuery("nPerPage", "100")
	nPerPage, err := strconv.Atoi(query)
	if err != nil {
		return int64(nPerPage), err
	}
	if nPerPage < 0 {
		return int64(nPerPage), errors.New("nPerPage should be positive number")
	}
	return int64(nPerPage), nil
}
