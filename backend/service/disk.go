package service

import (
	"fmt"
	"time"

	"github.com/selahaddinislamoglu/moni/backend/model"
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskService struct {
	data model.Disk
}

func NewDiskService() Disk {
	disk := &DiskService{}
	disk.startMonitoring()
	return disk
}

func (d *DiskService) startMonitoring() {
	go func() {
		var prevReadBytes, prevWriteBytes uint64
		var prevReadTime int64
		for {
			start := time.Now()
			usages, err := disk.IOCounters()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if len(usages) == 0 {
				fmt.Println("No disk usage data found")
				continue
			}
			readTime := time.Now().UnixMilli()

			var readBytes, writeBytes uint64
			var readBytesDiff, writeBytesDiff uint64
			var timeDiff float64

			for _, usage := range usages {
				readBytes += usage.ReadBytes
				writeBytes += usage.WriteBytes
			}
			if prevReadBytes == 0 && prevWriteBytes == 0 {
				goto end
			}

			readBytesDiff = readBytes - prevReadBytes
			writeBytesDiff = writeBytes - prevWriteBytes
			timeDiff = float64(readTime-prevReadTime) / 1000
			d.data.Usage = (float64(readBytesDiff+writeBytesDiff) / timeDiff) / (1024 * 1024)
			d.data.Time = time.Now().Unix()
		end:
			prevReadBytes = readBytes
			prevWriteBytes = writeBytes
			prevReadTime = readTime
			time.Sleep(5*time.Second - time.Since(start))
		}
	}()
}

func (d *DiskService) GetUsageLastFiveSeconds() (model.Disk, error) {
	return d.data, nil
}
