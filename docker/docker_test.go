// Copyright 2023 Dirk Strauss
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
	if os.Getenv("CI") != "" {
		t.Skip("Skipped due to CI environment")
	}
	GetCurrentContainers()
	//assert.NotNil(t, containers, "the given containers array is nil!!")
}

func TestDurationSince(t *testing.T) {
	var now = time.Now()
	var seconds = int64(34)
	var minutes = int64(54)
	var hours = int64(1)
	var durationSeconds = seconds + minutes*60 + hours*60*60
	var nowInPast = now.Unix() - durationSeconds
	var lostMinutes = GetExecutingSeconds(nowInPast, now)
	assert.True(t, lostMinutes > 0, "No lost seconds??")
	assert.Equal(t, uint64(durationSeconds), lostMinutes, "The time is not as it should be!")
}
