package handler

import (
	"net/http"
	"strconv"
	"survorest/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
func (h *userHandler) Create(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}
	user, err := h.userService.LoginUserForm(input)
	if err != nil || user.IsAdmin != "admin" {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}
	session := sessions.Default(c)
	session.Set("userID", user.Id)
	session.Set("username", user.FullName)
	session.Save()

	c.Redirect(http.StatusFound, "/users")

}

func (h *userHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}
func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUser()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userHandler) Delete(c *gin.Context) {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	err := h.userService.DeleteUser(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.Redirect(http.StatusFound, "/users")
}
