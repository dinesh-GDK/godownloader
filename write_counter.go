package godownloader

import "github.com/dinesh-GDK/multibar"

type WriteCounter struct {
	currTotal   float64
	Total       float64
	ProgressBar multibar.ProgressFunc
	Freq        uint64
	UpdateFreq  uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.currTotal += float64(n)
	wc.Freq += 1

	if wc.Freq%wc.UpdateFreq == 0 {
		wc.ProgressBar(int((wc.currTotal / wc.Total) * 100))
		wc.Freq = 0
	}
	return n, nil
}
