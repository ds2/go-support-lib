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
