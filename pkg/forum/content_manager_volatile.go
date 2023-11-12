package forum

type ConentMangerVolatile struct {
	posts    map[Id]*Story
	comments map[Id]*Comment
}

func (cmv *ConentMangerVolatile) AddStory(u *User, p *Story) error {
	// todo: validate post
	cmv.posts[p.Id] = p

	return nil
}

func NewContentManagerVolatile() *ConentMangerVolatile {
	return &ConentMangerVolatile{
		posts:    make(map[Id]*Story),
		comments: make(map[Id]*Comment),
	}
}
