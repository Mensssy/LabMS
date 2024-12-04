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
	pageSize4User  = 8
	pageSize4Admin = 11
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
	invoice.State = 1

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

func UserGetInvoices(c *gin.Context) {
	userId, _ := c.Get("userId")
	pageNum, err := strconv.Atoi(c.Param("pageNum"))

	if err != nil || pageNum < 1 {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "wrong pageNum format",
		})
		return
	}

	res, totalPageNum, err := dao.FindInvoices4User(userId.(string), pageSize4User, pageNum)
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

func AdminGetInvoices(c *gin.Context) {
	invoiceStat := c.Param("invoiceState")
	switch invoiceStat {
	case "submitted":
		getInvoicesByStat(1, c)
	case "checkpassed":
		getInvoicesByStat(2, c)
	case "delivered":
		getInvoicesByStat(4, c)
	default:
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "illegal invoiceState",
		})
	}
}

// 1:待审核 2:审核通过 3:驳回 4:已送报 5:报销成功 6:报销失败
func getInvoicesByStat(stat int, c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Param("pageNum"))

	if err != nil || pageNum < 1 {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "wrong pageNum",
		})
		return
	}

	res, totalPageNum, err := dao.FindInvoices4Stat(stat, pageSize4Admin, pageNum)
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

type invoiceIds struct {
	Ids []int `json:"ids"`
}

func SetInvoiceStat(c *gin.Context) {
	state, err := strconv.Atoi(c.Param("state"))
	if err != nil || (state < 1 || state > 6) {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "wrong or illegal state",
		})
		return
	}
	var ids invoiceIds
	err = c.ShouldBindJSON(&ids)
	if err != nil {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "wrong or illegal ids",
		})
		return
	}

	err = dao.SetInvoiceStat(ids.Ids, state)
	if err != nil {
		c.AbortWithStatusJSON(response.Internal_Server_Error, response.Body{
			Msg: "database error",
		})
		return
	}

	c.JSON(response.OK, response.Body{
		Msg: "set invoice state succeeded",
	})
}
