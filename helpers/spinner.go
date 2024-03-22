package helpers

//this is how a very simple spinner works
//func Spinner(stop chan bool) {
//	const delay = 100 * time.Millisecond
//	const loops = 10000
//	for i := 0; i < loops; i++ {
//		stopSpinner := <-stop
//		if !stopSpinner {
//			for _, r := range `-\|/` {
//				fmt.Printf("\r%c", r)
//				time.Sleep(delay)
//			}
//		}
//	}
//}

//here some spinner sample with simulated work:

//func main() {
//	pauseItForAMoment := false
//	// create a spinner
//	spinner, err := createSpinner()
//	if err != nil {
//		fmt.Printf("failed to make spinner from config struct: %v\n", err)
//		os.Exit(1)
//	}
//
//	// start the spinner
//	if err := spinner.Start(); err != nil {
//		panic(err)
//	}
//
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	var (
//		total   int
//		current int
//	)
//
//	spinnerCh := make(chan int, 0)
//	data := make(chan string)
//
//	// only want one go routine at a time but this is not important
//	max := make(chan struct{}, 1)
//
//	go func(s *yacspin.Spinner) {
//		// tried moving this into the worker goroutine also, but same effect
//		for range spinnerCh {
//			current += 1
//			if !pauseItForAMoment {
//				if err := s.Pause(); err != nil {
//					panic(err)
//				}
//				s.Message(fmt.Sprintf("%d/%d", current, total))
//				if err := s.Unpause(); err != nil {
//					panic(err)
//				}
//			}
//		}
//	}(spinner)
//
//	go func() {
//		defer wg.Done()
//		for d := range data {
//			wg.Add(1)
//			go func(wg *sync.WaitGroup, d string) {
//				max <- struct{}{}
//				defer func() {
//					<-max
//				}()
//
//				// function is doing work and printing the result once done.
//				pauseItForAMoment = true
//				spinner.Prefix(d + " ")
//				spinner.Stop()
//				spinner.Start()
//				pauseItForAMoment = false
//
//				// sends a value to the spinner go routine so that it can show
//				// the updated count
//				time.Sleep(4 * time.Second)
//				spinnerCh <- 1
//				wg.Done()
//			}(wg, d)
//		}
//	}()
//
//	// simulate queing some work
//	ss := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
//	for _, s := range ss {
//		data <- s
//	}
//	total = len(ss)
//
//	close(data)
//	wg.Wait()
//	close(spinnerCh)
//}
//
//func createSpinner() (*yacspin.Spinner, error) {
//	// build the configuration, each field is documented
//	cfg := yacspin.Config{
//		Frequency: 100 * time.Millisecond,
//		CharSet:   yacspin.CharSets[40],
//		Suffix:    " ", // puts a least one space between the animating spinner and the Message
//		// Message:           "collecting files",
//		SuffixAutoColon:   true,
//		ColorAll:          true,
//		Colors:            []string{"fgYellow"},
//		StopCharacter:     "✓",
//		StopColors:        []string{"fgGreen"},
//		StopMessage:       "done",
//		StopFailCharacter: "✗",
//		StopFailColors:    []string{"fgRed"},
//		StopFailMessage:   "failed",
//	}
//
//	s, err := yacspin.New(cfg)
//	if err != nil {
//		return nil, fmt.Errorf("failed to make spinner from struct: %w", err)
//	}
//
//	return s, nil
//}
