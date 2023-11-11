package forum

import (
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kacpekwasny/distributed-forum/pkg/enums"
	"github.com/kacpekwasny/distributed-forum/pkg/utils"
)

var tpl *template.Template

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("was not able to aquire the filename of current running file")
	}
	dirname := filepath.Dir(filename)
	templatesGlobSelector := filepath.Join(dirname, "templates", "*.go.html")

	tpl = template.Must(
		template.
			New("forum").
			Funcs(utils.FuncMapCommon).
			ParseGlob(templatesGlobSelector))

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
	err := tpl.Execute(w, []*Post{p, p})
	utils.Pife(err)
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
