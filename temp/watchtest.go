package main

import (
	"github.com/andreaskoch/go-fswatch"
	"fmt"
)

func main(){
	go func() {
		fileWatcher := fswatch.NewFileWatcher("test.txt").Start()
		for fileWatcher.IsRunning() {
			select {
			case <-fileWatcher.Modified:
				go func() {
					// file changed. do something.
					fmt.Println("changed")
				}()
			}
		}
	}()
}
