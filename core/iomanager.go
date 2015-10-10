package core

import(
	"path/filepath"
	"io/ioutil"
	"log"
	"io"
	"os"
	"fmt"
)

//const COPY_PATH = "C:\\Users\\Bruno\\Dropbox\\.gync"

type IOManager struct {
	Path string
}

func NewIOManager(path string) *IOManager {
	return &IOManager{path}
}

func (io IOManager) GetPath() string {
	dir, _ := filepath.Abs(io.Path)
	
	return dir
}

func (ioMan *IOManager) SaveObj(obj []byte) {
	
	err := ioutil.WriteFile(ioMan.Path, obj, 0600)

	if err != nil {
		log.Fatal(err)
	}
}


func (ioMan *IOManager) LoadFile() []byte {
	b, err := ioutil.ReadFile(ioMan.GetPath())
	
	if err != nil {
		log.Fatal("Unable to open file: ", err)
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

func (ioMan *IOManager) putEmptyJson(path string) {
	err := ioutil.WriteFile(path, []byte("{}"), 0600)

	if err != nil {
		log.Fatal(err)
	}
}

func InitDirectory(name string, dir string) {
	if _,err := os.Stat(filepath.Join(dir, name)); os.IsNotExist(err) {

		err := os.MkdirAll(filepath.Join(dir, name), 0777)
	
		if err != nil {
			log.Fatal("failed to initialize directory: ", err)
		}
	}
}

func CopyDirContents(src string, dst string) error {
	if _,err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	err := filepath.Walk(src, func (path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(src, path)

		if err != nil {
			return err
		}

		destination := filepath.Join(dst, relativePath)

		if info.IsDir() {
			err := os.MkdirAll(destination, 0777)

			if err != nil {
				DeleteDir(destination)
				log.Fatal(err)
			}
		} else {
			err := CopyFile(path, destination)

			if err != nil {
				DeleteDir(destination)
				log.Fatal(err)
			}
		}
		
		return err
	})

	return err
}

func DeleteDir(path string) {
	info, err := os.Stat(path)

	if err != nil {
		log.Fatal("Error deleting: ",  err)
	}

	if info.IsDir() {
		err := os.RemoveAll(path)

		if err != nil {
			log.Fatal("Error deleting: ",  err)
		}		
	} else {
		err := os.Remove(path)
		
		if err != nil {
			log.Fatal("Error deleting: ",  err)
		}		
	}
}