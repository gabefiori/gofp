package finder

import (
	"log"
	"strings"
	"sync"

	"github.com/mitchellh/go-homedir"
)

func Run(sources []Source, output chan string) {
	home, err := homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, source := range sources {
		wg.Add(1)

		go func(s Source) {
			defer wg.Done()

			paths, err := s.Find()

			if err != nil {
				log.Fatal(err)
			}

			for _, p := range paths {
				output <- "~" + strings.TrimPrefix(p, home)
			}
		}(source)
	}

	wg.Wait()
	close(output)
}
