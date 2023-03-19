package main

import (
	"os"
	"path/filepath"
)

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func main() {
	//ReadDir
	path, err := os.Getwd()
	if err == nil {
		dirEntries, err := os.ReadDir(path)
		if err == nil {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v",
					dentry.Name(), dentry.IsDir())
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}

	//Stat
	targetFiles := []string{"no_such_file.txt", "config.json"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File does not exist: %v", name)
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(),
				info.Size())
		}
	}

	//Glob
	path2, err := os.Getwd()
	if err == nil {
		matches, err := filepath.Glob(filepath.Join(path2, "*.json"))
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}

	//WalkDir
	path3, err := os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path3, callback)
	} else {
		Printfln("Error %v", err.Error())
	}

}
