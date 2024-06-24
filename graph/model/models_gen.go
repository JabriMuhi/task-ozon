// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID       int    `json:"id"`
	PostID   int    `json:"postId"`
	ParentID *int   `json:"parentId,omitempty"`
	Content  string `json:"content"`
	Author   *User  `json:"author"`
	Level    int    `json:"level"`
}

type Link struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Mutation struct {
}

type Post struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	Author          *User  `json:"author"`
	CommentsAllowed bool   `json:"commentsAllowed"`
}

type Query struct {
}

type Subscription struct {
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
