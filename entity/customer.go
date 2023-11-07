package entity

import (
	"time"
)

type Customer struct {
	Id            int
	Name_customer string
	No_hp         string
}

type Pelayanan struct {
	Id              int
	Jenis_pelayanan string
	Satuan          string
	Harga           int
}

type TrxBill struct {
	Id            int
	No            int
	Tanggal_masuk time.Time
	Diterima_oleh string
}

type TrxBillDetail struct {
	Id             int
	Customer_id    int
	Pelayanan_id   int
	Trx_bill_id    int
	Jumlah         int
	Tanggal_keluar time.Time
	total_Bayar    int
}
