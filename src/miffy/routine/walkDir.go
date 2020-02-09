package routine

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//RunWarlDir func
func RunWarlDir() {
	flag.Parse()
	filesize := make(chan int64)
	var wg sync.WaitGroup
	tick := time.Tick(100 * time.Millisecond)
	// go func() {
	//此处必须保证wg先有队列进去后再监听Wait事件
	for _, root := range flag.Args() {
		wg.Add(1)
		go walk(root, filesize, &wg)
	}
	// }()

	go func() {
		wg.Wait()
		close(filesize)
	}()

	var nfiles, nbytes int64
	//打印并发程序的进度, 利用time.Tick
loop:
	for {
		select {
		case size, ok := <-filesize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func walk(dir string, ch chan<- int64, n *sync.WaitGroup) {
	//每次递归完成后将sync.waitgroup的数量减一
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walk(subdir, ch, n)
		} else {
			ch <- entry.Size()
		}
	}

}

func dirents(dir string) []os.FileInfo {
	//TODO: 可以利用缓冲通道来限速
	entires, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return entires
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.2f GB\n", nfiles, float64(nbytes)/1e9)
}
