package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
Нужно получить токен пользователя, провалидировать его и записать в контекст
*/

const (
	authorizationHandler = "Authorization"
	userCtx              = "userCode"
)

// userIdentity - проверка авторизации
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHandler)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[1] == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	// чтобы каждый раз не ходить в БД за id пользователя по его коду -
	// в userCtx будем держать его uuid код
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

// getUserId - получение UserId из контекста
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
