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

			err := s.Find(opts.OutputChan, func(p string) string {
				return "~" + strings.TrimPrefix(p, opts.HomeDir)
			})

			if err != nil {
				log.Fatal(err)
			}

		}(source)
	}

	wg.Wait()
	close(opts.OutputChan)
}
