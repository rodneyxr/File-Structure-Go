package filestructure

import (
	"errors"
	"fmt"
)

type FileStruct struct {
	Name  string
	parent *FileStruct
	files map[string]*FileStruct
}

// returns a new FileStruct type
func NewFileStruct(name string) (*FileStruct, error) {
	if name == "" || name == "." || name == ".." {
		return nil, errors.New("FileStruct: Invalid file name.")
	}

	fs := new(FileStruct)
	fs.Name = name
	fs.parent = fs
	fs.files = map[string]*FileStruct{
		".": fs,
		"..": fs.parent,
	}

	return fs, nil
}

func (fs *FileStruct) InsertFilePath(fp FilePath) (*FileStruct, error) {
	return nil, errors.New("FileStruct: not yet implemented.")
}

// creates all necessary directories to the specified filepath
func (fs *FileStruct) mkdir(fp FilePath) (*FileStruct, error) {
	return nil, errors.New("FileStruct: not yet implemented.")
}

func (fs *FileStruct) Print() {
	// iterate over all files in the file struct
	for _, file := range fs.files {
		// ignore the directory if itself or parent
		if file == fs || file.parent == fs.parent { continue }
		file.Print()
	}
	// finally print this file's name
	fmt.Println(fs.Name)
}

func (fs *FileStruct) IsNil() bool {
	return fs.Name == ""
}

func (fs FileStruct) String() string {
	return fs.Name
}