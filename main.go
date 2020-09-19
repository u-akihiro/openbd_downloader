package main

import (
	"context"
	"downloader/downloader"
	p "downloader/parser"
	"downloader/util"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"
	"golang.org/x/sync/semaphore"
)

func main() {
	start := time.Now()

	b, _ := downloader.DownloadCoverage()
	coverage, _ := p.ParserCoverageJSON(b)

	sliced := util.Slice(coverage, util.ChunkSize)

	sem := semaphore.NewWeighted(4)

	wg := &sync.WaitGroup{}
	for i := 0; i < len(sliced); i++ {
		wg.Add(1)
		sem.Acquire(context.Background(), 1)

		go func(slice []string) {
			defer wg.Done()
			defer sem.Release(1)
			json, _ := downloader.DownloadBookInfo(slice)

			t := time.Now()
			entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
			id := ulid.MustNew(ulid.Timestamp(t), entropy)

			ioutil.WriteFile("./json/"+id.String()+".json", json, 0644)
			fmt.Println("Download done.")
		}(sliced[i])
	}

	wg.Wait()

	end := time.Now()
	duration := end.Sub(start)

	fmt.Printf("%v\n", duration)
	// for i := 0; i < len(res.([]interface{})); i++ {
	// 	fmt.Println(res.([]interface{})[i].(map[string]interface{})["summary"].(map[string]interface{})["title"])
	// 	fmt.Println(res.([]interface{})[i].(map[string]interface{})["summary"].(map[string]interface{})["cover"])
	// }
}
