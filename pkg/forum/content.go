package forum

import (
	"google.golang.org/genproto/googleapis/type/datetime"
)

type ContentManagerInterface interface {
	AddPost(*User, *Post) error
	AddComment(*User, *Comment) error
}

type Id uint32

func GetId() Id {

}

type PostInterface interface{}
type CommentInterface interface{}
type UserInterface interface{}

type ConentMangerVolatile struct {
	posts    map[Id]*Post
	comments map[Id]*Comment
}

func (cmv *ConentMangerVolatile) AddPost(u *User, p *Post) error {
	// todo: validate post
	cmv.posts[p.id] = p

	return nil
}

func NewContentManagerVolatile() *ConentMangerVolatile {
	return &ConentMangerVolatile{
		posts:    make(map[Id]*Post),
		comments: make(map[Id]*Comment),
	}
}

type Postable struct {
	id        Id
	author    *User
	contents  string
	timestamp datetime.DateTime
}

type Likeable struct {
	likes    []Id
	dislikes []Id
}

type Post struct {
	Postable
	Likeable
}

type Comment struct {
	Postable
	Likeable

	postId Id
}

type User struct {
	id       Id
	username string
}
