package pkg

import (
	"github.com/ms56bc/rate-limter/pkg/algorithm"
	"testing"
)

func TestFixedWindow_Evict_Based_On_time(t *testing.T) {
	fixedWindow := algorithm.NewFixedWindow()
	fixedWindow.HasLimitReached(NewIpAddressEntityId(""))
}
