package spinner

import (
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func New(Prefix, Suffix string) *Manager {
	loading := []string{"|", "/", "-", "\\"}
	s := spinner.New(loading, 100*time.Millisecond)
	s.Prefix = Prefix // Prefix text before the spinner
	s.Suffix = Suffix // Append text after the spinner
	s.Writer = os.Stderr

	return &Manager{
		Spinner: s,
	}
}

type Manager struct {
	Spinner *spinner.Spinner
	GoStop  bool
}

func (self *Manager) Start() {
	self.Spinner.Start()
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			if self.GoStop {
				return
			}
		}
	}()
}

func (self *Manager) Stop() {
	if !self.GoStop {
		self.GoStop = true
		self.Spinner.Stop()
	}
}

func (self *Manager) Restart() {
	self.Spinner.Restart()
}
