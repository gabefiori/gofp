package ui

import (
	fzf "github.com/junegunn/fzf/src"
)

func Run(inputChan chan string) (string, error) {
	outputChan := make(chan string)
	resultChan := make(chan string)

	go func() {
		for out := range outputChan {
			resultChan <- out
		}

		close(resultChan)
	}()

	options, err := fzf.ParseOptions(true, nil)

	if err != nil {
		return "", err
	}

	options.Input = inputChan
	options.Output = outputChan

	_, err = fzf.Run(options)

	close(outputChan)

	if err != nil {
		return "", err
	}

	return <-resultChan, nil
}
