package packagePrint

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	timeBefore := time.Now()
	Print()
	timeAfter := time.Now()
	fmt.Println("Time taken", timeAfter.Second()-timeBefore.Second(), "seconds")
}
