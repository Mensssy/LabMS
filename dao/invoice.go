package dao

import (
	"com.mensssy.LabMS/dao/db"
	"com.mensssy.LabMS/model"
)

var (
	pageSize = 8
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

func FindInvoices(userId string, pageNum int) ([]model.Invoice, int64, error) {
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
