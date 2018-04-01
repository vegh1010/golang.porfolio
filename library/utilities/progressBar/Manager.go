package progressBar

import (
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"

	"sync"
)

type Manager struct {
	Progress *mpb.Progress
	WG       sync.WaitGroup
	Bars     []*mpb.Bar
}

func (self *Manager) AddBar(name string, total int) *mpb.Bar {
	// Add a bar
	// You're not limited to just a single bar, add as many as you need
	bar := self.Progress.AddBar(int64(total),
		// Prepending decorators
		mpb.PrependDecorators(
			// StaticName decorator with minWidth and no width sync options
			// If you need to change name while rendering, use DynamicName
			decor.StaticName(name, len(name), 0),
			// ETA decorator with minWidth and width sync options
			// DSyncSpace is shortcut for DwidthSync|DextraSpace
			decor.ETA(4, decor.DSyncSpace),
		),
		// Appending decorators
		mpb.AppendDecorators(
			// Percentage decorator with minWidth and no width sync options
			decor.Percentage(5, 0),
		),
	)
	self.Bars = append(self.Bars, bar)

	return bar
}

func (self *Manager) Stop() {
	self.Progress.Stop() // Stop mpb's rendering goroutine
}

func (self *Manager) Done() {
	self.WG.Done()
}

func (self *Manager) Wait() {
	self.WG.Wait() // Wait for goroutines to finish
}

func (self *Manager) Complete() {
	for _, bar := range self.Bars {
		bar.Complete()
	}
}

//-------------------IMPLEMENTATION EXAMPLE----------------//
//func SendData() {
//	BManager := New(1)
//	total := 100
//	bar := BManager.AddBar("SampleBar", total)
//
//	go func() {
//		defer BManager.Done()
//		for i := 0; i < total; i++ {
//			//logic here
//
//			bar.Incr(1) // increment progress bar
//		}
//	}()
//
//	BManager.Wait()
//	BManager.Stop()
//}
