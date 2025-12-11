package handler

import (
    "encoding/json"
    "net/http"
)

// Blog represents a simple blog structure
type Blog struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

// GetBlogsHandler godoc
// @Summary Get all blogs
// @Tags Blog
// @Success 200 {array} Blog
// @Router /blogs [get]
func GetBlogsHandler(w http.ResponseWriter, r *http.Request) {
    blogs := []Blog{
        {ID: 1, Title: "First Post", Content: "Welcome to Osmity Blog!"},
    }
    json.NewEncoder(w).Encode(blogs)
}

// CreateBlogHandler godoc
// @Summary Create a new blog
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog body Blog true "Blog data"
// @Success 200 {object} Blog
// @Router /blogs [post]
func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
    var b Blog
    json.NewDecoder(r.Body).Decode(&b)

    // 仮の ID 付与
    b.ID = 999

    json.NewEncoder(w).Encode(b)
}
