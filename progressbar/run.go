package progressbar

import (
	"github.com/cheggaaa/pb/v3"
)

const runProgressbarTemplate = `{{percent . | green}} {{bar . " " "━" "━" "─" " " | green}} {{counters . | green}} {{speed . "(%s/s)" | red}} {{rtime . "%s left" | blue}}`

func NewRun() *RunProgressBar {
	return &RunProgressBar{
		progressbar: pb.New64(0).
			SetWidth(128).
			SetTemplateString(runProgressbarTemplate),
	}
}

type RunProgressBar struct {
	totalObjects     int64
	completedObjects int64
	progressbar      *pb.ProgressBar
}

func (rp *RunProgressBar) Start() {
	if rp == nil {
		return
	}
	rp.progressbar.Start()
}

func (rp *RunProgressBar) Finish() {
	if rp == nil {
		return
	}
	rp.progressbar.Finish()
}

func (rp *RunProgressBar) AddCompletedCommands(commands int64) {
	if rp == nil {
		return
	}
	rp.progressbar.Add64(commands)
}

func (rp *RunProgressBar) AddTotalCommands(commands int64) {
	if rp == nil {
		return
	}
	rp.progressbar.AddTotal(commands)
}
