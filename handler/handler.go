package handler

import (
	helper "mockup_server/helper"
	repo "mockup_server/repo"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GenerateReadErrorResponse(c, err)
		return
	}

	resp, errReq := repo.SendLoginReq(request.Username, request.Password)
	if errReq != nil {
		helper.GenerateReadErrorResponse(c, errReq)
		return
	}

	c.JSON(helper.FormatResponse(helper.READSUCCESS_200, nil, struct {
		Token string `json:"token"`
	}{
		Token: repo.GenerateJWT(*resp),
	}))
}

func Logout(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.GenerateReadErrorResponse(c, err)
		return
	}

	resp, errReq := repo.SendLogoutReq(request.Username)
	if errReq != nil {
		helper.GenerateReadErrorResponse(c, errReq)
		return
	}

	c.JSON(helper.FormatResponse(helper.READSUCCESS_200, nil, resp))
}
