package service

import (
	"fmt"
	"strconv"

	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/dao"
	"com.mensssy.LabMS/model"
	"github.com/gin-gonic/gin"
)

var (
	PageSize = 8
)

func SubmitInvoice(c *gin.Context) {
	var invoice model.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "error parameters",
		})
		return
	}

	//在上下文中获取到userId
	userId, _ := (c.Get("userId"))
	invoice.UserId = userId.(string)

	invoiceId, err := dao.SaveInvoice(invoice)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(response.Internal_Server_Error, response.Body{
			Msg: "database error, invoice save failed",
		})
		return
	}

	c.JSON(response.OK, response.Body{
		Msg: "invoice save succeeded",
		Data: map[string]int{
			"invoiceId": invoiceId,
		},
	})
}

func GetInvoices(c *gin.Context) {
	userId, _ := c.Get("userId")
	pageNum, err := strconv.Atoi(c.Param("pageNum"))

	if err != nil {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "wrong pageNum format",
		})
		return
	}

	res, totalPageNum, err := dao.FindInvoices(userId.(string), pageNum)
	if err != nil {
		c.AbortWithStatusJSON(response.Internal_Server_Error, response.Body{
			Msg: "database error",
		})
		return
	}

	c.JSON(response.OK, response.Body{
		Msg: "get invoices succeeded",
		Data: map[string]interface{}{
			"invoices":     res,
			"totalPageNum": totalPageNum,
		},
	})

}

func UploadInvoiceDoc(c *gin.Context) {
	Upload(c, "invoiceDoc")
}

func DownloadInvoiceDoc(c *gin.Context) {
	Download(c, "invoiceDoc")
}
