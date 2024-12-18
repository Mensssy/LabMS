package dao

import (
	"fmt"
	"time"

	"com.mensssy.LabMS/dao/db"
	"com.mensssy.LabMS/model"
	"gorm.io/gorm"
)

func SaveInvoice(invoice model.Invoice) (int, error) {
	db := db.SqlDB
	tx := db.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return 0, tx.Error
	}

	// 创建发票
	res := tx.Create(&invoice)

	if res.Error != nil {
		tx.Rollback()
		return 0, res.Error
	}

	return invoice.InvoiceId, tx.Commit().Error
}

// 根据userId查询发票
func FindInvoices4User(userId string, pageSize int, pageNum int) ([]model.Invoice, int64, error) {
	db := db.SqlDB

	var totalInvoiceNum int64
	var invoices []model.Invoice

	//获取发票总数
	res := db.Model(&model.Invoice{}).Where("user_Id = ?", userId).Count(&totalInvoiceNum)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	//获取总页数=发票总数/页大小
	totalPageNum := (totalInvoiceNum / int64(pageSize)) + 1

	//获取该页发票信息
	res = db.Where("user_id = ?", userId).Order("invoice_id DESC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&invoices)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	return invoices, totalPageNum, nil
}

// 根据发票的状态查询发票，为管理员功能
func FindInvoices4Stat(stat int, pageSize int, pageNum int) ([]model.Invoice, int64, error) {
	db := db.SqlDB

	var totalInvoiceNum int64
	var invoices []model.Invoice

	//获取发票总数
	res := db.Model(&model.Invoice{}).Where("state = ?", stat).Count(&totalInvoiceNum)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	//获取总页数=发票总数/页大小
	totalPageNum := (totalInvoiceNum / int64(pageSize)) + 1

	//获取该页发票信息
	res = db.Where("state = ?", stat).Order("invoice_id DESC").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&invoices)

	if res.Error != nil {
		return nil, 0, res.Error
	}

	return invoices, totalPageNum, nil
}

func SetInvoiceStat(invoiceIds []int, stat int) error {
	db := db.SqlDB
	tx := db.Begin()

	var res *gorm.DB
	// 设置为已送报时，放入一个批次，批次以送报日期命名
	if stat == 4 {
		res = tx.Model(&model.Invoice{}).Where("invoice_id IN ?", invoiceIds).Updates(map[string]interface{}{"state": 4, "delivery_date": time.Now().Format("2006-01-02")})
	} else {

		res = tx.Model(&model.Invoice{}).Where("invoice_id IN ?", invoiceIds).Update("state", stat)
	}

	if res.Error != nil {
		fmt.Println(res.Error.Error())
		tx.Rollback()
		return res.Error
	}

	return tx.Commit().Error
}

func UpdateInvoice(invoice model.Invoice) error {
	db := db.SqlDB
	tx := db.Begin()

	res := tx.Model(&model.Invoice{}).Where("invoice_id = ?", invoice.InvoiceId).Updates(&invoice)
	if res.Error != nil {
		return res.Error
	}

	return tx.Commit().Error
}

func GetBatches() ([]string, error) {
	db := db.SqlDB

	var batches []time.Time

	res := db.Model(&model.Invoice{}).Where("state >= 4").Distinct().Pluck("delivery_date", &batches)
	dates := make([]string, len(batches))

	//获取批次名列表 即送报日期
	for v := range batches {
		dates[v] = batches[v].Format("2006-01-02")
		if dates[v] == "2004-02-18" {
			dates = dates[0:v]
		}
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return dates, nil
}

func GetBatch(delivery_date string, groupType string, pageNum int, pageSize int) ([]model.Invoice, int, error) {
	db := db.SqlDB

	var totalInvoiceNum int64
	//获取发票总数
	res := db.Model(&model.Invoice{}).Where("delivery_date = ? AND state >= 4", delivery_date).Count(&totalInvoiceNum)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	//获取总页数=发票总数/页大小
	totalPageNum := (totalInvoiceNum / int64(pageSize)) + 1

	var invoices []model.Invoice

	res = db.Model(&model.Invoice{}).Where("delivery_date = ? AND state >= 4", delivery_date).Order(groupType).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&invoices)

	if res.Error != nil {
		return nil, 0, res.Error
	}
	return invoices, int(totalPageNum), nil
}
