package busniess

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/model"
)

func AddPostToArchive(year int, month int, post *model.Post) {
	selector := bson.M{"year": year, "month": month}
	update := bson.M{"$push": bson.M{"posts": post}}
	_, err := model.UpsertArchive(selector, update)
	if err != nil {
		panic(err)
	}
}

type ArchiveOut struct {
	ID    string `json:"id"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
}

func newArchiveOut(archive *model.Archive) *ArchiveOut {
	out := &ArchiveOut{
		ID:    archive.ID.Hex(),
		Year:  archive.Year,
		Month: archive.Month,
	}
	return out
}

type ListArchivesOut struct {
	*ArchiveOut
}

func ListArchives() []*ListArchivesOut {
	archives, err := model.ListArchives(nil)
	if err != nil {
		panic(err)
	}
	outs := make([]*ListArchivesOut, 0, len(archives))
	for _, archive := range archives {
		out := &ListArchivesOut{ArchiveOut: newArchiveOut(archive)}
		outs = append(outs, out)
	}
	return outs
}
