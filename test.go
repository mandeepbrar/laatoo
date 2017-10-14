package main

/*
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk("/home/mandeep/goprogs/src/laatoo/services", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".json") {
			newpath := fmt.Sprint(strings.TrimSuffix(path, ".json"), ".yml")
			err := os.Rename(path, newpath)
			fmt.Printf("json found %s %s %s", path, newpath, err)
		}
		return nil
	})
	fmt.Println("search ended", err)
}
*/
