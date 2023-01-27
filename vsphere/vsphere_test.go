// Copyright 2023 Dirk Strauss
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vsphere

import (
	"context"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/vmware/govmomi/vim25"
	"gitlab.com/ds_2/go-support-lib/common"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(false)
	logrus.SetLevel(logrus.DebugLevel)
}

func TestCreateClient(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipped due to CI environment")
	}
	type args struct {
		myContext     context.Context
		serverUrlStr  string
		username      string
		pw            string
		allowInsecure bool
	}
	tests := []struct {
		name    string
		args    args
		want    *vim25.Client
		wantErr bool
	}{
		{
			name: "testGenUrl1",
			args: args{
				myContext:     context.Background(),
				serverUrlStr:  common.GetEnvVar("VSPHERE_URL", "https://my.vsphere.local:12345/"),
				username:      common.GetEnvVar("VSPHERE_USERNAME", "nouser"),
				pw:            common.GetEnvVar("VSPHERE_PW", "nopw"),
				allowInsecure: true,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateClient(tt.args.myContext, tt.args.serverUrlStr, tt.args.username, tt.args.pw, tt.args.allowInsecure, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// we have a client!
			if got == nil {
				t.Errorf("CreateClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
