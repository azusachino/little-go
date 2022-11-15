package leak

import (
	"go.uber.org/goleak"
	"testing"
)

func TestGetData(t *testing.T) {
	GetData()
}

func TestGetDataWithGoleak(t *testing.T) {
	defer goleak.VerifyNone(t)
	GetData()
}
