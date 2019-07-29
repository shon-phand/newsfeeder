package handler

import (
	"fmt"
	"net/http"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := newsfeedPostRequest{}
		fmt.Println("Before Binding", requestBody)
		c.Bind(&requestBody)
		fmt.Println("After binding", requestBody)
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		fmt.Println("item value before adding", item)
		feed.Add(item)

		c.Status(http.StatusOK)
	}
}
