package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func GetCurrentContainers() []ContainerInfo {
	var rc []ContainerInfo
	{
	}
	cli, err := client.NewEnvClient()
	if err != nil {
		logrus.Panic("Error when setting up the new client: ", err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		logrus.Debugln(container.ID[:10], " and image=",container.Image, "with labels=",container.Labels)
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
		logrus.Infoln("Info about this container is:", thisInfo)
		rc = append(rc, thisInfo)
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

func StopContainer(containerId string, maxWaitTime time.Duration) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	errContainerStop := cli.ContainerStop(ctx, containerId, &maxWaitTime)
	if errContainerStop != nil {
		logrus.Warn("Could not stop container ", containerId, " due to: ", errContainerStop)
		return false
	}
	return true
}

func TerminateContainer(containerId string) bool {
	cli, err := client.NewEnvClient()
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
