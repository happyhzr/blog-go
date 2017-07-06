package busniess

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/model"
)

type FileIn struct {
	Name      string
	Path      string
	Size      int64
	CreatedBy string
}

type FileOut struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func newFileOut(file *model.File) *FileOut {
	out := &FileOut{
		ID:   file.ID.Hex(),
		Name: file.Name,
		Path: file.Path,
		Size: file.Size,
	}
	return out
}

type CreateFileIn struct {
	FileIn
}

type CreateFileOut struct {
	FileOut
}

func CreateFile(in *CreateFileIn) *CreateFileOut {
	file := &model.File{
		Name:      in.Name,
		Path:      in.Path,
		Size:      in.Size,
		CreatedBy: bson.ObjectIdHex(in.CreatedBy),
		CreatedAt: time.Now().Unix(),
	}
	err := file.Save()
	if err != nil {
		panic(err)
	}
	out := &CreateFileOut{FileOut: *newFileOut(file)}
	return out
}

type ListFilesOut struct {
	FileOut
}

func ListFiles() []*ListFilesOut {
	files, err := model.ListFiles(nil)
	if err != nil {
		panic(err)
	}
	outs := make([]*ListFilesOut, 0, len(files))
	for _, file := range files {
		out := &ListFilesOut{FileOut: *newFileOut(file)}
		outs = append(outs, out)
	}
	return outs
}
