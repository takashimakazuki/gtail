package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)


func initQueue() ([] string, int) {
	queue := []string{}
	cursor := 0
	return queue, cursor
}

func enqueue(queue []string, value string) []string{
	queue = append(queue, value)
	return queue
}

func dequeue(queue []string) []string {
	queue = queue[1:]
	return queue
}

func show(queues []string) {
	for _, queue := range queues {
		fmt.Printf("\x1b[32m%s\x1b[0m\n", queue)
	}
}

func tail(stream *os.File, n int) []string {
	queue, cursor := initQueue()
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		if n < 1 {
			n = int(math.Abs(float64(n)))
			if n == 0 {
				n = 10
			}
		}
		queue = enqueue(queue, scanner.Text())
		if n-1 < cursor {
			queue = dequeue(queue)
		}
		cursor++
	}
	return queue
}

func main() {
	const USAGE string = "Usage: gtail [-n #] [file]"
	intOpt := flag.Int("n", 10, USAGE)
	flag.Usage = func() {
		fmt.Println(USAGE)
	}
	flag.Parse()
	n := int(math.Abs(float64(*intOpt)))
	if flag.NArg() > 0 {

		for i:= 0; i < flag.NArg(); i++ {
			if i > 0 {
				fmt.Print("\n")
			}
			if flag.NArg() != 1 {
				fmt.Println("==> " + flag.Arg(i) + " <==")
			}
			fp, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Printf("\u001B[31m%s\u001B[0m\n", "Error: No such file or directory")
				os.Exit(1)
			}
			defer fp.Close()
			show(tail(fp, n))
		}
	}

}