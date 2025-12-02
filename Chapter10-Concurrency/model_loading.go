package main

import (
	"fmt"
	"time"
)

/*
	func main() {
		logs := make(chan string)

		// Worker
		go func() {
			for i := 1; i <= 10; i++ {
				logs <- fmt.Sprintf("Loading module %d/10...", i)
				time.Sleep(400 * time.Millisecond)
			}
			close(logs)
		}()

		// UI
		for msg := range logs {
			fmt.Println(msg)
		}
		fmt.Println("Done!")
	}
*/

/*
func main() {
	done := make(chan bool)

	// Spinner UI
	go func() {
		chars := `-\|/`
		for {
			for _, c := range chars {
				select {
				case <-done:
					return
				default:
					fmt.Printf("\rLoading... %c", c)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	// Backend work
	time.Sleep(3 * time.Second)
	done <- true

	fmt.Println("\nCompleted!")
}
*/

type JobStatus struct {
	Name  string
	Step  int
	Total int
}

func main() {
	statusChan := make(chan JobStatus)

	// 3 ta parallel job
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for step := 1; step <= 10; step++ {
				statusChan <- JobStatus{
					Name:  fmt.Sprintf("Job %d", id),
					Step:  step,
					Total: 10,
				}
				time.Sleep(time.Duration(200+id*50) * time.Millisecond)
			}
		}(i)
	}

	// UI renderer
	go func() {
		progress := map[string]JobStatus{}

		for st := range statusChan {
			progress[st.Name] = st

			fmt.Print("\033[H\033[2J") // clear screen

			for _, p := range progress {
				percent := p.Step * 100 / p.Total
				fmt.Printf("%s [%d/%d] %d%%\n", p.Name, p.Step, p.Total, percent)
			}
		}
	}()

	time.Sleep(4 * time.Second)
}
