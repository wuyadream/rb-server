package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var rubbishMap = map[string]string {
	"电池": "有害垃圾",
	"玻璃": "玻金塑纸",
	"剩菜": "厨余垃圾",
	"沙发": "废旧家具",
	"盆栽": "年花年桔",
	"袜子": "废旧织物",
	"塑料袋": "其他垃圾",
}

func GetType(c *gin.Context) {

	rubbish := c.DefaultQuery("rubbish", "")

	rubbishType, ok  := rubbishMap[rubbish]

	if ok == false {
		rubbishType = "未知垃圾"
	}

	c.JSON(http.StatusOK, gin.H{"rubbish": rubbish, "type": rubbishType})

}