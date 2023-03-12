package FileWordCount

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func CountTotalWordInDir(dirPath string) int {
	file, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	list, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	totalNumberOfFilesInDir := len(list)
	var total_words_count uint64
	var wg sync.WaitGroup
	for i := 0; i < totalNumberOfFilesInDir; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			wordsInFile := CountWordFile(filepath.Join(dirPath, filepath.Base(list[i].Name())))
			for c := 0; c < wordsInFile; c++ {
				atomic.AddUint64(&total_words_count, 1)
			}
		}(i)
	}
	wg.Wait()

	return int(total_words_count)
}

func CountWordFile(path string) int {
	fileHandle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)
	count := 0
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			break
		}
		count++
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return count
}
