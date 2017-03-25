package main

import (
	"encoding/csv"
	"flag"
	"math/rand"
	"os"
	"strconv"
)

var n = flag.Int("n", 10000000, "How many entries to generate")

func main() {
	wr := csv.NewWriter(os.Stdout)
	defer wr.Flush()

	wr.Write([]string{
		"tour_id",
		"package_id",
		"date_start",
		"stars",
		"pansion",
		"nights",
		"region_id",
		"hotel_id",
		"airport_id",
		"price",
	})

	for i := 0; i < *n; i++ {
		wr.Write([]string{
			strconv.FormatInt(int64(i), 10),
			strconv.FormatInt(rand.Int63(), 10),
			strconv.FormatInt(rand.Int63n(90), 10),
			strconv.FormatInt(rand.Int63n(10), 10),
			strconv.FormatInt(rand.Int63n(10), 10),
			strconv.FormatInt(rand.Int63n(30), 10),
			strconv.FormatInt(rand.Int63n(1000), 10),
			strconv.FormatInt(rand.Int63n(1000000), 10),
			strconv.FormatInt(rand.Int63n(10), 10),
			strconv.FormatInt(rand.Int63n(10000), 10),
		})
	}
}
