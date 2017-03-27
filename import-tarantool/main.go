package main

import (
	"encoding/csv"
	"flag"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/tarantool/go-tarantool"
)

var n = flag.Int("n", 64, "n")

func main() {
	flag.Parse()
	opts := tarantool.Opts{User: "guest"}
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		panic(err)
	}

	fp, err := os.Open("tours.csv")
	if err != nil {
		panic(err)
	}

	rd := csv.NewReader(fp)
	rd.Read()

	recCh := make(chan []int, 16)

	wg := sync.WaitGroup{}
	for i := 0; i < *n; i++ {
		wg.Add(1)
		go func() {
			for entry := range recCh {
				_, err := conn.Insert(conn.Schema.Spaces["tours"], entry)
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}()
	}

	for {
		record, err := rd.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		ints := make([]int, 0, len(record))
		for _, r := range record {
			i, _ := strconv.Atoi(r)
			ints = append(ints, i)
		}

		recCh <- ints
	}

	close(recCh)
	wg.Done()
}
