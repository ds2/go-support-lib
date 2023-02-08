// Copyright 2022 Dirk Strauss
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

package docker

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func GetCurrentContainers() []*ContainerInfo {
	var rc []*ContainerInfo
	{
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Panic("Error when setting up the new client: ", err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		logrus.Debugln(container.ID[:10], " and image=", container.Image, "with labels=", container.Labels)
		var thisInfo = ContainerInfo{
			Id:              container.ID,
			Image:           container.Image,
			Started:         container.Created,
			Names:           container.Names,
			Imageid:         container.ImageID,
			InternalCommand: container.Command,
			State:           container.State,
			Status:          container.Status,
		}
		logrus.Infoln("Info about this container is:", &thisInfo)
		rc = append(rc, &thisInfo)
	}
	return rc
}

func getDurationBetween(time1 time.Time, time2 time.Time) time.Duration {
	return time1.Sub(time2)
}

func GetExecutingSeconds(startSecondsUnix int64, nowTime time.Time) uint64 {
	logrus.Debug("Start time to calculate from is ", startSecondsUnix)
	var containerTimeLocal = time.Unix(startSecondsUnix, 0)
	var duration = getDurationBetween(nowTime, containerTimeLocal)
	logrus.Debug("The duration between these two times is ", duration)
	var durationInSeconds = uint64(duration.Seconds())
	logrus.Debug("Duration in seconds is ", durationInSeconds)
	return durationInSeconds
}

func StopContainer(containerId string, maxWaitTime *time.Duration) bool {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	var maxWaitSeconds int = int(maxWaitTime.Seconds())
	stopOptions := container.StopOptions{
		Signal:  "",
		Timeout: &maxWaitSeconds,
	}
	errContainerStop := cli.ContainerStop(ctx, containerId, stopOptions)
	if errContainerStop != nil {
		logrus.Warn("Could not stop container ", containerId, " due to: ", errContainerStop)
		return false
	}
	return true
}

func TerminateContainer(containerId string) bool {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	errContainerStop := cli.ContainerKill(ctx, containerId, "KILL")
	if errContainerStop != nil {
		logrus.Warn("Could not kill container ", containerId, " due to: ", errContainerStop)
		return false
	}
	return true
}

// StopAndTerminateContainer tries to stop a given container first. If stop is not successful after the given waitTime, the container gets terminated.
func StopAndTerminateContainer(containerId string, maxWaitTime *time.Duration) bool {
	if !StopContainer(containerId, maxWaitTime) {
		logrus.Warn("Could not stop the container, will now try to terminate it.")
		return TerminateContainer(containerId)
	}
	return true
}
