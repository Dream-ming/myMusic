package user

func UserInfoHandler(c *gin.Context) {
	userID := 1
	user, err := service.GetUserInfo(userID)
    if err != nil {
        c.JSON(500, gin.H{"error": "获取用户信息失败"})
        return
    }
    c.JSON(200, user)
}