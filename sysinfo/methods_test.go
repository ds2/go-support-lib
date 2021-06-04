package sysinfo

import (
	"testing"
)

func TestGetNodeDetails(t *testing.T) {
	t.Run("simpleTest", func(t *testing.T) {
		var _, err = GetNodeDetails()
		if err != nil {
			t.Errorf("GetNodeDetails Error = %v", err)
		}
	})
}

func TestGetDiskSizeInfo(t *testing.T) {
	t.Run("simpleTest", func(t *testing.T) {
		var _, err = GetDiskSizeInfo("/")
		if err != nil {
			t.Errorf("GetNodeDetails Error = %v", err)
		}
	})
}
