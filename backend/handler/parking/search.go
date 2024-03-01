package parking

import (
	"backend/model/parking"
	"backend/model/requests"
	"backend/model/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SearchParkingLotsByLocation godoc
// @Summary      Search Parking Lots By latitude and longitude
// @Description  通过经纬度搜索附近的停车场
// @Tags         parking
// @Accept       json
// @Produce      json
// @Param        latitude  query  float64  true  "Latitude 经纬度"
// @Param        longitude  query  float64  true  "Longitude 经纬度"
// @Success      200  {object}  response.Response
// @Router       /parking/search [post]
func SearchParkingLotsByLocation(c *gin.Context) {
	var request requests.ParkingSearchByLatitudeAndLongitude
	if err := c.ShouldBind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	parkingLots, err := parking.FindNearbyParkingLots(request.Latitude, request.Longitude)
	if err != nil {
		response.Error(c, err, "搜索停车场失败")
		return
	}

	response.JSON(c, parkingLots)
}

// SearchHistoryByUserID godoc
// @Summary      Search Parking History By UserID
// @Description  通过用户ID查找用户的停车历史
// @Tags         parking
// @Accept       json
// @Produce      json
// @Param        latitude  query  float64  true  "Latitude 经纬度"
// @Param        longitude  query  float64  true  "Longitude 经纬度"
// @Success      200  {object}  response.Response
// @Router       /parking/history [get]
func SearchHistoryByUserID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	historyRecords, err := parking.SearchParkingHistoryByUserID(userID)
	if err != nil {
		response.Error(c, err, "搜索停车历史失败")
		return
	}

	response.JSON(c, historyRecords) // 使用你的响应工具返回数据
}

func SearchReserveByUserID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userId")) // 从查询参数中获取用户ID
	pr := parking.ParkingReservationModel{
		UserID: userID,
	}
	reservations, err := pr.GetByUserID()
	if err != nil {
		response.Error(c, err, "搜索预定的停车历史失败")
		return
	}
	response.JSON(c, reservations)
}

func CancelReserveByID(c *gin.Context) {
	reserveID, _ := strconv.Atoi(c.Query("reserveId")) // 从查询参数中获取用户ID
	err := parking.CancelReserveByReserveId(reserveID)
	if err != nil {
		response.Error(c, err, "取消预定失败")
		return
	}
	response.JSON(c, "取消预定成功") // 使用你的响应工具返回数据

}
