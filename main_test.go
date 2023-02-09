package packagePrint

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	timeBefore := time.Now()
	Print()
	timeAfter := time.Now()
	timeTaken := timeAfter.Second() - timeBefore.Second()
	if timeTaken != 2 {
		t.Fail()
	}
}
