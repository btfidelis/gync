package core

import(
	"path/filepath"
	"io/ioutil"
	"log"
	"io"
	"os"
	"fmt"
)

const STORAGE string = "../storage/"
const COPY_PATH = "C:\\Users\\Bruno\\Dropbox\\.gync"

type IOManager struct {
	Path string
}

func NewIOManager(path string) *IOManager {
	return &IOManager{path}
}

func (io IOManager) GetPath() string {
	dir, _ := filepath.Abs(STORAGE + io.Path)
	
	return dir
}

func (ioMan *IOManager) SaveObj(obj []byte) {
	
	err := ioutil.WriteFile(STORAGE + ioMan.Path, obj, 0600)

	if err != nil {
		log.Fatal(err)
	}
}


func (ioMan *IOManager) LoadFile() []byte {
	b, err := ioutil.ReadFile(ioMan.GetPath())
	
	if err != nil {
		log.Fatal("Unable to open file")
		return nil
	}

	return b
}

func CopyFile(srcPath string, destPath string) error {
	srcFile, err := os.Stat(srcPath)

	if err != nil {
		return err
	}

	if !srcFile.Mode().IsRegular() {
		return fmt.Errorf("Unregular source file: %s (%s)", srcFile.Name(), srcFile.Mode().String());
	}

	destFile, err := os.Stat(destPath)

	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if !(destFile.Mode().IsRegular()) {
			return fmt.Errorf("Unregular source file: %s (%s)", srcFile.Name(), srcFile.Mode().String());
		}
		if os.SameFile(srcFile, destFile) {
			return nil
		}
	}

	if err := os.Link(srcPath, destPath); err == nil {
		return err
	}

	return copyFileContents(srcPath, destPath)

}

func copyFileContents(srcPath string, destPath string) error {
	srcFile, err := os.Open(srcPath)

	if err != nil {
		return err
	}

	defer srcFile.Close()

	destFile, err := os.Create(destPath)

	if err != nil {
		return err
	}

	defer func() {
		cerr := destFile.Close()
		if err == nil {
			err = cerr
		}
	}()
	fmt.Println("coping file")
	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}

	return destFile.Sync()
}