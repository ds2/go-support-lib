package docker

import (
	"context"
	"fmt"
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
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		var thisInfo = ContainerInfo{
			Id:      container.ID,
			Image:   container.Image,
			Started: container.Created,
		}
		fmt.Println("Info about this container is: ", thisInfo)
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
