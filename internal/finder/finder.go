package finder

import (
	"log"
	"strings"
	"sync"
)

type FinderOpts struct {
	Sources    []Source
	OutputChan chan string
	HomeDir    string
}

func Run(opts *FinderOpts) {
	var wg sync.WaitGroup

	for _, source := range opts.Sources {
		wg.Add(1)

		go func(s Source) {
			defer wg.Done()

			paths, err := s.Find()

			if err != nil {
				log.Fatal(err)
			}

			for _, p := range paths {
				opts.OutputChan <- "~" + strings.TrimPrefix(p, opts.HomeDir)
			}
		}(source)
	}

	wg.Wait()
	close(opts.OutputChan)
}
