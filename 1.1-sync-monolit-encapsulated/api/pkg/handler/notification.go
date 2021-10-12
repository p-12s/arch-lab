package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/utils"
)

func (h *Handler) sendEmail(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input common.NotificationSendReq
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = utils.SendEmail(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Notification.SaveNotification(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "OK",
	})
}
