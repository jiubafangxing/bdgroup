package router

import (
	"github.com/bdgroup/config"
	"github.com/bdgroup/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func users() func(c *gin.Context) {
	return func(c *gin.Context) {
		usernames, err := service.QueryUsernames()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"userNames": usernames,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}

// 设置Cookies，如果对应cookies是新用户则新增
func setCookies() func(c *gin.Context) {
	return func(c *gin.Context) {
		var json struct {
			BDCookies string `json:"bdCookies" binding:"required"`
		}
		// 绑定 JSON 数据并校验
		if err := c.ShouldBindJSON(&json); err == nil {
			bdnInfo, err := config.Instance.SetBdnInfoByCookies(json.BDCookies)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			} else {

				curtime := time.Now()
				existUser, err := service.GetUserByUK(bdnInfo.UK)
				if nil != err {
					return
				}
				if nil == existUser {
					_, err := service.InsertUser(bdnInfo.Username, bdnInfo.UK, curtime, curtime, bdnInfo.Cookies)
					if err != nil {
						return
					}
				} else {
					err := service.UpdateUser(bdnInfo.UK, bdnInfo.Cookies, curtime)
					if err != nil {
						return
					}
				}
				log.Printf("百度网盘登录验证登录成功, 昵称:%s", bdnInfo.Username)
				c.JSON(http.StatusOK, gin.H{
					"username": bdnInfo.Username,
				})
			}
			return
		} else {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}

func swatchUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var json struct {
			UserName string `json:"userName" binding:"required"`
		}
		if err := c.ShouldBindJSON(&json); err == nil {
			user, err := service.GetByUserName(json.UserName)
			if err == nil {
				bdnInfo, err := config.Instance.SetBdnInfoByCookies(user.Cookies)
				if err != nil {
					c.JSON(http.StatusBadGateway, gin.H{
						"error": err.Error(),
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"username": bdnInfo.Username,
					})
					log.Printf("百度网盘登录验证登录成功, 昵称:%s", bdnInfo.Username)
				}
			}
		} else {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}
	}
}
