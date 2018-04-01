package progressBar

import (
	"github.com/vbauerster/mpb"

	"os"
	"sync"
	"time"
)

func New(numBars int) *Manager {
	p := mpb.New(
		// override default (80) width
		mpb.WithWidth(100),
		// override default "[=>-]" format
		mpb.WithFormat("╢▌▌░╟"),
		// override default 100ms refresh rate
		mpb.WithRefreshRate(120*time.Millisecond),

		mpb.Output(os.Stderr),
	)

	var wg sync.WaitGroup
	wg.Add(numBars)

	return &Manager{
		Progress: p,
		WG:       wg,
	}
}
