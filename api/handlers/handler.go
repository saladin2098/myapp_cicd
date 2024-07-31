package handlers


import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Message struct {
    Message string `json:"message"`
}

// Hello godoc
// @Summary Returns a hello message
// @Description Get a hello message
// @Tags example
// @Produce json
// @Success 200 {object} Message
// @Router /api/hello [get]
func Hello(c *gin.Context) {
    c.JSON(http.StatusOK, Message{Message: "Hello, World! megajin!!!!!"})
}

// Goodbye godoc
// @Summary Returns a goodbye message
// @Description Post a goodbye message
// @Tags example
// @Produce json
// @Param message body Message true "Message"
// @Success 200 {object} Message
// @Router /api/goodbye [post]
func Goodbye(c *gin.Context) {
    var msg Message
    if err := c.ShouldBindJSON(&msg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, Message{Message: "Goodbye, ninmadir" + msg.Message})
}
