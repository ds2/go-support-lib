package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ds_2/go-support-lib/docker"
)

func init() {
	fmt.Println("Wir sind da")
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(false)
	logrus.SetLevel(logrus.DebugLevel)
}

func TestGetCurrentContainers(t *testing.T) {
	var containers = docker.GetCurrentContainers()
	assert.NotNil(t, containers, "the given containers array is nil!!")
}

func TestDurationSince(t *testing.T) {
	var now = time.Now()
	var seconds = int64(34)
	var minutes = int64(54)
	var hours = int64(1)
	var durationSeconds = seconds + minutes*60 + hours*60*60
	var nowInPast = now.Unix() - durationSeconds
	var lostMinutes = docker.GetExecutingSeconds(nowInPast, now)
	assert.True(t, lostMinutes > 0, "No lost seconds??")
	assert.Equal(t, uint64(durationSeconds), lostMinutes, "The time is not as it should be!")
}
