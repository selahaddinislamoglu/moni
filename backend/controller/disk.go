package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/moni/backend/service"
)

type DiskController struct {
	diskService service.Disk
}

func NewDiskController() *DiskController {
	return &DiskController{}
}

func (d *DiskController) SetupDiskService(diskService service.Disk) {
	d.diskService = diskService
}

func (d *DiskController) GetUsageLastFiveSeconds(ctx *gin.Context) {
	data, err := d.diskService.GetUsageLastFiveSeconds()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get disk usage"})
		return
	}
	ctx.JSON(200, data)
}
