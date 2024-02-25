package user

import (
	"backend/model/requests"
	"fmt"
	"strconv"

	"backend/model/response"

	"backend/model/user"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Add New Bank Card
// @Description  添加新的银行卡
// @Tags         sign
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.SignupUsingPhoneRequest  true  "Phone--手机号||Password-- 密码|| Name--昵称"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /user/card  [post]
//
//	AddNewCard  添加新的银行卡
func AddNewCard(c *gin.Context) {
	// 1. 验证表单
	request := requests.AddNewCardRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Abort500(c, "绑定参数失败")
		return
	}
	userID, _ := strconv.Atoi(request.UserID) // 从查询参数中获取用户ID

	// 2. 验证，创建数据
	cardModel := user.BankCardsBound{
		UserId:       userID,
		CardNumber:   request.CardNumber,
		CardPassword: request.CardPassword,
		BankName:     request.BankName,
		ExpiryDate:   request.ExpiryDate,
	}
	if !cardModel.Valid() {
		response.BadRequest(c, fmt.Errorf("请求的参数不合法"))
		return
	}
	err := cardModel.Create()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	if cardModel.ID > 0 {
		response.JSON(c, "绑定银行卡成功")
	} else {
		response.Abort500(c, "绑定银行卡失败，请稍后尝试~")
		return
	}

}

func GetBoundedCard(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	cardModel := user.BankCardsBound{
		UserId: userID,
	}
	boundedCards, err := cardModel.Get()
	if err != nil {
		response.Error(c, err, "获取绑定的银行卡失败")
		return
	}

	response.JSON(c, boundedCards) // 使用你的响应工具返回数据
}
