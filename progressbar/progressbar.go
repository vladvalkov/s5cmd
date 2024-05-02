package progressbar

import (
	"fmt"
	"sync/atomic"

	"github.com/cheggaaa/pb/v3"
)

type CopyProgressBar struct {
	totalObjects     int64
	completedObjects int64
	progressbar      *pb.ProgressBar
}

const copyProgressbarTemplate = `{{percent . | green}} {{bar . " " "━" "━" "─" " " | green}} {{counters . | green}} {{speed . "(%s/s)" | red}} {{rtime . "%s left" | blue}} {{ string . "objects" | yellow}}`

func NewCopy() *CopyProgressBar {
	return &CopyProgressBar{
		progressbar: pb.New64(0).
			Set(pb.Bytes, true).
			Set(pb.SIBytesPrefix, true).
			SetWidth(128).
			Set("objects", fmt.Sprintf("(%d/%d)", 0, 0)).
			SetTemplateString(copyProgressbarTemplate),
	}
}

func (cp *CopyProgressBar) Start() {
	if cp == nil {
		return
	}

	cp.progressbar.Start()
}

func (cp *CopyProgressBar) Finish() {
	if cp == nil {
		return
	}

	cp.progressbar.Finish()
}

func (cp *CopyProgressBar) IncrementCompletedObjects() {
	if cp == nil {
		return
	}

	atomic.AddInt64(&cp.completedObjects, 1)
	cp.progressbar.Set("objects", fmt.Sprintf("(%d/%d)", cp.completedObjects, cp.totalObjects))
}

func (cp *CopyProgressBar) IncrementTotalObjects() {
	if cp == nil {
		return
	}

	atomic.AddInt64(&cp.totalObjects, 1)
	cp.progressbar.Set("objects", fmt.Sprintf("(%d/%d)", cp.completedObjects, cp.totalObjects))
}

func (cp *CopyProgressBar) AddCompletedBytes(bytes int64) {
	if cp == nil {
		return
	}

	cp.progressbar.Add64(bytes)
}

func (cp *CopyProgressBar) AddTotalBytes(bytes int64) {
	if cp == nil {
		return
	}

	cp.progressbar.AddTotal(bytes)
}
