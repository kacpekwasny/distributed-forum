package forum

import (
	"math"
	"math/rand"
	"net/http"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	"github.com/kacpekwasny/distributed-forum/pkg/enums"
)

var tpl *template.Template

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("was not able to aquire the filename of current running file")
	}
	dirname := filepath.Dir(filename)
	templatesGlobSelector := filepath.Join(dirname, "templates", "*.gohtml")

	tpl = template.Must(
		template.Must(
			template.ParseGlob(templatesGlobSelector),
		).ParseGlob(templatesGlobSelector),
	)
}

type ContentManagerInterface interface {
	AddPost(*User, *Post) error
	AddComment(*User, *Comment) error
	LikePost(postId Id, r Reaction)
	LikeComment(postId Id, r Reaction)
}

type Id uint64

func NewRandId() Id {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return Id(math.MaxUint64 * r.Float64())
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
	cmv.posts[p.Id] = p

	return nil
}

func NewContentManagerVolatile() *ConentMangerVolatile {
	return &ConentMangerVolatile{
		posts:    make(map[Id]*Post),
		comments: make(map[Id]*Comment),
	}
}

func RenderPost(w http.ResponseWriter, p *Post) {
	// var tmplFile1 = "/home/kacper/go/src/github.com/kacpekwasny/distributed-forum/pkg/forum/post.go.html"
	// var tmplFile2 = "/home/kacper/go/src/github.com/kacpekwasny/distributed-forum/pkg/forum/all-posts.go.html"
	err := tpl.Execute(w, []*Post{p, p})
	if err != nil {
		panic(err)
	}
}

type TimeStampable struct {
	Timestamp uint64
}

type Postable struct {
	Id       Id
	UserId   Id
	Contents string

	TimeStampable
}

type Reaction struct {
	UserId    Id
	ReactType enums.ReactType

	TimeStampable
}

type Reactionable struct {
	Reactions []Reaction
}

type Post struct {
	Postable
	Reactionable
}

type Comment struct {
	PostId Id

	Postable
	Reactionable
}

type User struct {
	Id       Id
	Username string
}
