package example

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func testable() error {
	rand.Seed(time.Now().UnixNano())
	if num := rand.Float32(); num > 0.5 {
		return errors.New(fmt.Sprintf("%f > 0.5", num))
	}
	return nil
}
