package service

import (
	"fmt"
	"time"

	"github.com/selahaddinislamoglu/moni/internal/model"
	"github.com/shirou/gopsutil/v4/net"
)

type NetworkService struct {
	data      model.Network
	publisher Publisher
}

func NewNetworkService() Network {
	network := &NetworkService{}
	network.startMonitoring()
	return network
}

func (d *NetworkService) Setup(Publisher Publisher) {
	d.publisher = Publisher
}

func (d *NetworkService) startMonitoring() {
	go func() {
		var prevBytesRecv, prevBytesSent uint64
		var prevReadTime int64
		for {
			start := time.Now()
			usages, err := net.IOCounters(false)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if len(usages) == 0 {
				fmt.Println("No disk usage data found")
				continue
			}

			readTime := time.Now().UnixMilli()
			usage := usages[0]
			recvBytesDiff := usage.BytesRecv - prevBytesRecv
			sendBytesDiff := usage.BytesSent - prevBytesSent
			timeDiff := float64(readTime-prevReadTime) / 1000

			if prevBytesRecv == 0 && prevBytesSent == 0 {
				goto end
			}

			d.data.Usage = (float64(recvBytesDiff+sendBytesDiff) / timeDiff) / (1024 * 1024)
			d.data.Time = time.Now().Unix()
			d.publisher.Publish(NetworkTopic, d.data)
		end:
			prevBytesRecv = usage.BytesRecv
			prevBytesSent = usage.BytesSent
			prevReadTime = readTime
			time.Sleep(5*time.Second - time.Since(start))
		}
	}()
}

func (d *NetworkService) GetUsageLastFiveSeconds() (model.Network, error) {
	return d.data, nil
}
