package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/blog-api-v2/models"
	"github.com/gorilla/mux"
)

var posts = []models.Post{
	{
		ID:      "p1",
		Title:   "First Post",
		Content: "This is my first blog post",
		Author:  "Eddy",
	},
	{
		ID:      "p2",
		Title:   "Second Post",
		Content: "Welcome to my blog!",
		Author:  "Mouity",
	},
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, post := range posts {
		if post.ID == id {
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post
	_ = json.NewDecoder(r.Body).Decode(&newPost)

	// Fix ID generation
	newPost.ID = "p" + strconv.Itoa(len(posts)+1)

	posts = append(posts, newPost)
	json.NewEncoder(w).Encode(newPost)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i, post := range posts {
		if post.ID == id {
			_ = json.NewDecoder(r.Body).Decode(&posts[i])
			posts[i].ID = id
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.Write([]byte("Post deleted"))
			return
		}
	}
	http.Error(w, "Post not found", http.StatusBadRequest)
}
