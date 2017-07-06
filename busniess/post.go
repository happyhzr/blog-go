package busniess

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/insisthzr/blog-back/model"
)

type PostIn struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedBy string `json:"-"`
}

type PostOut struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	CreatedAt int64    `json:"created_at"`
	CreatedBy *UserOut `json:"created_by"`
}

func newPostOut(post *model.Post) *PostOut {
	out := &PostOut{
		ID:        post.ID.Hex(),
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
	}
	if post.CreatedBy != "" {
		user, err := model.GetUserByID(post.CreatedBy)
		if err != nil {
			if err != mgo.ErrNotFound {
				panic(err)
			}
		} else {
			out.CreatedBy = newUserOut(user)
		}
	}
	return out
}

type CreatePostIn struct {
	PostIn
}

type CreatePostOut struct {
	PostOut
}

func CreatePost(in *CreatePostIn) *CreatePostOut {
	post := &model.Post{
		Title:     in.Title,
		Body:      in.Body,
		CreatedBy: bson.ObjectIdHex(in.CreatedBy),
		CreatedAt: time.Now().Unix(),
	}
	err := post.Save()
	if err != nil {
		panic(err)
	}
	t := time.Unix(post.CreatedAt, 0)
	AddPostToArchive(t.Year(), int(t.Month()), post.ID)
	out := &CreatePostOut{PostOut: *newPostOut(post)}
	return out
}

type ListPostsIn struct {
	Range SkipLimit
}

type ListPostsOut struct {
	PostOut
}

func ListPosts(in *ListPostsIn) []*ListPostsOut {
	posts, err := model.ListPostsWithRange(nil, in.Range.Skip, in.Range.Limit)
	if err != nil {
		panic(err)
	}
	outs := make([]*ListPostsOut, 0, len(posts))
	for _, post := range posts {
		out := &ListPostsOut{PostOut: *newPostOut(post)}
		outs = append(outs, out)
	}
	return outs
}

func CountPosts() int {
	count, err := model.CountPosts(nil)
	if err != nil {
		panic(err)
	}
	return count
}
