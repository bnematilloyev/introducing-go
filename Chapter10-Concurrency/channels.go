package main

/*
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func main() {
	// Channels - provide a way
	// for two goroutines to communicate

	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}
*/

// Channel direction

/*
func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		for {
			ch1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			ch2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-ch1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-ch2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("Timeout")
				//default:
				//	fmt.Println("Nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
*/
/*
	| Type               | Sender behavior                      | Receiver behavior                         |
	| ------------------ | ------------------------------------ | ----------------------------------------- |
	| Unbuffered channel | bloklanadi, receiver tayyor bo‘lmasa | bloklanadi, sender ma’lumot yubormaguncha |
	| Buffered channel   | bloklanmaydi, buffer bo‘sh bo‘lsa    | bloklanmaydi, bufferda element bo‘lsa     |
*/
