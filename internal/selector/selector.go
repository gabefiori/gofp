package selector

// Displays a series of options for user selection.
type Selector interface {
	Run(inputChan chan string) (string, error)
}

