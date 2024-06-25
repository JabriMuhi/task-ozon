package in_memory

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

func NewUser(id int, username string, password string, email string) *User {
	return &User{
		Id:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}

type Comment struct {
	Id      int
	Post_id int
	User_id int
	Text    string
}

type CommentParentChild struct {
	Parent_id   int
	Children_id int
	Level       int
}

func NewComment(id int, post_id int, user_id int, text string) *Comment {
	return &Comment{
		Id:      id,
		Post_id: post_id,
		User_id: user_id,
		Text:    text,
	}
}

func NewCommentParentChild(parent_id int, children_id int, level int) *CommentParentChild {
	return &CommentParentChild{
		Parent_id:   parent_id,
		Children_id: children_id,
		Level:       level,
	}
}

type Post struct {
	Id               int
	Title            string
	Content          string
	Author_id        int
	Comments_allowed bool
}

func NewPost(id int, title string, content string, author_id int, comments_allowed bool) *Post {
	return &Post{
		Id:               id,
		Title:            title,
		Content:          content,
		Author_id:        author_id,
		Comments_allowed: comments_allowed,
	}
}

type InMemory struct {
	Users               map[int]User
	Posts               map[int]Post
	Comments            map[int]Comment
	CommentsParentChild []CommentParentChild
}

func InitInMemory(users map[int]User, posts map[int]Post, comments map[int]Comment, commentsParentChild []CommentParentChild) *InMemory {
	return &InMemory{
		Users:               users,
		Posts:               posts,
		Comments:            comments,
		CommentsParentChild: commentsParentChild,
	}
}
