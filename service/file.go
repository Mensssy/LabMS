package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/util"
	"github.com/gin-gonic/gin"
)

var (
	invoiceDocPath string = "/usr/local/labMS/invoiceDocs/"
)

func Upload(c *gin.Context, fileType string) {
	//获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(response.Bad_Request, err)
	}

	switch fileType {
	//发票凭证
	case "invoiceDoc":
		userId, _ := c.Get("userId")
		invoiceId := c.PostForm("invoiceId")

		//文件名称为invoiceId的加密
		fileSuffix := filepath.Ext(file.Filename)
		file.Filename = util.Encrypt(invoiceId) + fileSuffix

		//保存在对应加密usrId文件夹下
		finalName := invoiceDocPath + userId.(string) + "/" + file.Filename

		//保证目录存在
		err = os.MkdirAll(filepath.Dir(finalName), 0755)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(response.Internal_Server_Error, response.Body{
				Msg: "make dir failed",
			})
			return
		}

		//保存文件
		err = c.SaveUploadedFile(file, finalName)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(response.Internal_Server_Error, response.Body{
				Msg: "file save failed",
			})
			return
		}
		c.JSON(response.OK, response.Body{
			Msg: "upload invoice document succeed",
		})

		return
	}

}

func Download(c *gin.Context, fileType string) {
	switch fileType {
	case "invoiceDoc":
		userId, _ := c.Get("userId")
		invoiceId := c.Param("invoiceId")

		fileDir := invoiceDocPath + userId.(string) + "/"
		fileNameWithoutSuffix := util.Encrypt(invoiceId)

		files, err := os.ReadDir(fileDir)
		if err != nil {
			c.AbortWithStatusJSON(response.Internal_Server_Error, response.Body{
				Msg: "user not exists",
			})
		}

		for _, file := range files {
			//系统中的文件名
			fileName := file.Name()
			//文件名去掉后缀
			if fileNameWithoutSuffix == strings.TrimSuffix(fileName, filepath.Ext(fileName)) {
				tarFilePath := fileDir + fileName

				//设置请求头
				c.Header("Content-Type", "application/octet-stream")
				c.Header("Content-Disposition", "attachment;filename="+"发票编号"+invoiceId+"凭证"+filepath.Ext(fileName))

				//读取文件
				c.File(tarFilePath)

				return
			}
		}
	}
}
