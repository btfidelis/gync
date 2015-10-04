package core

import(
	"path/filepath"
	"io/ioutil"
	"log"
)

const storagePath string = "../storage/"

type IOManager struct {
	Path string
}

func NewIOManager(path string) *IOManager {
	return &IOManager{path}
}

func (io IOManager) GetPath() string {
	dir, _ := filepath.Abs(storagePath + io.Path)
	
	return dir
}

func (ioMan *IOManager) SaveObj(obj []byte) {
	
	err := ioutil.WriteFile(storagePath + ioMan.Path, obj, 0600)

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

func CopyFile(srcPath, srcDest string) err error {
	srcFile, err := os.Stat(srcFile)

	if err != nil {
		return err
	}

	if !srcFile.Mode().IsRegular() {
		return fmt.Errorf("Unregular source file: %s (%s)", srcFile.Name(), srcFile.Mode().String());
	}

	destFile, err := os.Stat(srcDest)

	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if !(srcDest.Mode().IsRegular()) {
			return fmt.Errorf("Unregular source file: %s (%s)", srcFile.Name(), srcFile.Mode().String());
		}
		if os.SameFile(srcFile, destFile) {
			return nil
		}
	}

	if err = os.Link(srcFile, srcDest) {
		return nil
	}

	err = copyFileContents(srcPath, srcDest)

	return err
}

func copyFileContents(srcPath, destPath string) error {
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

	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}

	return destFile.Sync()
}