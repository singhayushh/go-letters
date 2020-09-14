package backend

import (
	"github.com/gin-gonic/gin"
)

// RenderHome ...
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Goletters",
	})
}

// Mail ...
func Mail(c *gin.Context) {
	senderName := c.PostForm("sender-name")
	senderEmail := c.PostForm("sender-email")
	receiverName := c.PostForm("receiver-name")
	receiverEmail := c.PostForm("receiver-email")
	password := c.PostForm("password")
	title := c.PostForm("title")
	body := c.PostForm("body")
	err := SendMail(body, title, senderName, senderEmail, password, receiverName, receiverEmail)
	if err != nil {
		c.HTML(200, "index.html", gin.H{
			"title": "Failed!",
		})
	} else {
		c.HTML(200, "index.html", gin.H{
			"title": "Success!",
		})
	}
}
