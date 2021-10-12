package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

func (h *Handler) createCard(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cardCode := uuid.New().String() // пусть будет uuid для простоты, в примре был формат "1235-6892"
	err = h.services.Loyalty.CreateUserCard(common.LoyaltyCard{
		UserId: userId,
		Code:   cardCode,
		Status: common.LOYALTY_CARD_CREATED,
	})
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":             userId,
		"loyalty_card_number": cardCode,
	})
}
