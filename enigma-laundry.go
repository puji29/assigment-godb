package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"challenge-godb/entity"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s  password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
mainMenuLoop:
	for {
		fmt.Println("===============ENIGMA LAUNDRY============")
		fmt.Println("1. View All Data")
		fmt.Println("2. Add New Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Delete Data")
		fmt.Println("5. Exit")
		fmt.Println("=========================================")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("enter you choice :")
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:

			for {
				fmt.Println("Select Choice To View Data")
				fmt.Println("1. View Data Customer")
				fmt.Println("2. View Data Pelayanan")
				fmt.Println("3. View Data transaksi Bill")
				fmt.Println("4. View Data transaksi Bill Detail")
				fmt.Println("5. Back To Menu Utama")

				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("enter you choice :")
				scanner.Scan()
				SubChoice, _ := strconv.Atoi(scanner.Text())

				switch SubChoice {
				case 1:
					customers := getAllCustomer()
					fmt.Printf("%-5s %-12s %-12s\n", "ID", "Nama Customer", "No. Hp")
					for _, customer := range customers {
						fmt.Printf("%-5d %-12s %-12s\n", customer.Id, customer.Name_customer, customer.No_hp)
					}
				case 2:
					pelayanans := getAllPelayanan()
					fmt.Printf("%-5s %-12s %-12s %-12s\n", "ID", "Jenis Pelayanan ", "Satuan", "Harga")
					for _, pelayanan := range pelayanans {
						fmt.Printf("%-5d %-12s %-12s %-12d\n", pelayanan.Id, pelayanan.Jenis_pelayanan, pelayanan.Satuan, pelayanan.Harga)
					}
				case 3:
					trxbills := getAlltrxBill()
					fmt.Printf("%-5s %-12s %-12s %-12s\n", "ID", "No ", "Tanggal Masuk", "Diterima oleh")

					for _, trtrxbill := range trxbills {

						fmt.Printf("%-5d %-12d %-12s %-12s\n", trtrxbill.Id, trtrxbill.No, trtrxbill.Tanggal_masuk.Format("2006-01-02"), trtrxbill.Diterima_oleh)
					}
				case 4:
					totalBayarResults := totalBayar()

					trxBillDetails := getAlltrxBillDetail()

					fmt.Printf("%-5s %-12s %-12s %-12s %-10s %-15s %-12s\n", "ID", "Customer ID", "Pelayanan ID", "Trx Bill ID", "Jumlah", "Tanggal Keluar", "Total Bayar")

					for i, trxBillDetail := range trxBillDetails {
						totalBayar := totalBayarResults[i]

						fmt.Printf("%-5d %-12d %-12d %-12d %-10d %-15s %-12.f\n",
							trxBillDetail.Id, trxBillDetail.Customer_id, trxBillDetail.Pelayanan_id,
							trxBillDetail.Trx_bill_id, trxBillDetail.Jumlah, trxBillDetail.Tanggal_keluar.Format("2006-01-02"), totalBayar)
					}
					totalByrKes := totalBayarKeseluruhan()
					fmt.Printf("Totdal Keselurahan: %.f", totalByrKes)

				case 5:

					continue mainMenuLoop
				default:
					fmt.Println("invalid sub-Choice")
				}
			}
		case 2:
		mainMenuLoop2:
			for {
				fmt.Println("Select Choice To ADD Data")
				fmt.Println("1. Add Data Detail Customer ")
				fmt.Println("2. Add Data Transaksi detail ")
				fmt.Println("3. Back To Menu Utama")

				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("enter you choice :")
				scanner.Scan()
				SubChoice2, _ := strconv.Atoi(scanner.Text())

				switch SubChoice2 {
				case 1:
					scanner := bufio.NewScanner(os.Stdin)

					fmt.Println("Masukan Data Customer Detail")
					fmt.Print("Id :")
					scanner.Scan()
					newId, _ := strconv.Atoi(scanner.Text())

					if isCustomerExists(newId) {
						fmt.Println("ID pelanggan sudah ada dalam database. Tidak dapat melakukan insert.")
						continue mainMenuLoop2
					}

					fmt.Print("Nama :")
					scanner.Scan()
					newName := scanner.Text()

					fmt.Print("NO.Hp :")
					scanner.Scan()
					newHp := scanner.Text()
					//validasi panjang inputan
					if len(newHp) > 12 {
						fmt.Println("Panjang input melebih 12. Silakan coba lagi")
						continue mainMenuLoop2
					}

					customer := entity.Customer{Id: newId, Name_customer: newName, No_hp: newHp}
					addCustomer(customer)

					fmt.Println("Masukan Data Detail Pelayanan")
					fmt.Print("Id :")
					scanner.Scan()
					newIdP, _ := strconv.Atoi(scanner.Text())

					fmt.Print("Jenis Pelayanan :")
					scanner.Scan()
					newJenis := scanner.Text()

					fmt.Print("Satuan (KG/Buah) :")
					scanner.Scan()
					newSatuan := scanner.Text()

					newSatuan = strings.ToLower(newSatuan)

					if newSatuan != "Buah" && newSatuan != "KG" {
						fmt.Println("Harap Memasukan Hanya KG/Buah.")
						continue mainMenuLoop2
					}
					fmt.Print("Harga :")
					scanner.Scan()
					newHarga, _ := strconv.Atoi(scanner.Text())

					pelayanan := entity.Pelayanan{Id: newIdP, Jenis_pelayanan: newJenis, Satuan: newSatuan, Harga: newHarga}
					addPelayanan(pelayanan)

					fmt.Println("Masukan Data Transaksi Bill")
					fmt.Print("Id :")
					scanner.Scan()
					newIdT, _ := strconv.Atoi(scanner.Text())

					fmt.Print("No: ")
					scanner.Scan()
					newNo, _ := strconv.Atoi(scanner.Text())

					waktu := time.Now()

					fmt.Print("Diterima Oleh: ")
					scanner.Scan()
					newDo := scanner.Text()

					trxBills := entity.TrxBill{Id: newIdT, No: newNo, Tanggal_masuk: waktu, Diterima_oleh: newDo}
					addTxBill(trxBills)
				case 2:
					fmt.Println("Masukan Transaksi Detail")
					fmt.Print("Id: ")
					scanner.Scan()
					newTbd, err := strconv.Atoi(scanner.Text())

					if err != nil {
						fmt.Println("Input ID Tidak boleh Kosong.")
						continue mainMenuLoop2
					}

					fmt.Print("Masukan Customer id: ")
					scanner.Scan()
					newCid, _ := strconv.Atoi(scanner.Text())

					if !isCustomerExists(newCid) {
						fmt.Println("Data id customer tidak ada di tTabel Customer.")
						continue mainMenuLoop2
					}

					fmt.Print("Masukan Pelayanan Id: ")
					scanner.Scan()
					newPid, _ := strconv.Atoi(scanner.Text())

					if !getPelyananId(newPid) {
						fmt.Println("Data Id Pelayanan Tidak ada di Tabel Pelayanan.")
						continue mainMenuLoop2
					}

					fmt.Print("Masukan Transaksi Bill Id: ")
					scanner.Scan()
					newTBI, _ := strconv.Atoi(scanner.Text())

					fmt.Print("Jumlah: ")
					scanner.Scan()
					newJml, err := strconv.Atoi(scanner.Text())

					if err != nil || newJml == 0 {
						fmt.Println("Jumlah Harus Angka. Harap masukkan jumlah yang benar.")
						continue mainMenuLoop2
					}

					waktu := time.Now()

					trxBillDetails := entity.TrxBillDetail{Id: newTbd, Customer_id: newCid, Pelayanan_id: newPid, Trx_bill_id: newTBI, Jumlah: newJml, Tanggal_keluar: waktu}
					addTxBillDetail(trxBillDetails)
				case 3:
					continue mainMenuLoop
				default:
					fmt.Println("invalid sub-Choice")
				}
			}
		case 3:
		mainMenuLoop3:
			for {
				fmt.Println("Select Choice To Update Data")
				fmt.Println("1. Update Data Customer")
				fmt.Println("2. Update Data Pelayanan")
				fmt.Println("3. Back To Menu Utama")

				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("enter you choice :")
				scanner.Scan()
				SubChoice1, _ := strconv.Atoi(scanner.Text())

				switch SubChoice1 {
				case 1:

					scanner := bufio.NewScanner(os.Stdin)

					fmt.Print("Enter Customer Id to update: ")
					scanner.Scan()
					custId, _ := strconv.Atoi(scanner.Text())

					if !isCustomerExists(custId) {
						fmt.Println("Id Customer tidak ditemukan.")
						continue mainMenuLoop3
					}

					fmt.Print("Masukan Nama:")
					scanner.Scan()
					updateNama := scanner.Text()

					fmt.Print("Masukan NO.Hp :")
					scanner.Scan()
					updateNo_hp := scanner.Text()
					//validasi panjang inputan
					if len(updateNo_hp) > 12 {
						fmt.Println("Panjang input melebih 12. Silakan coba lagi")
						continue mainMenuLoop3
					}
					customer := entity.Customer{Id: custId, Name_customer: updateNama, No_hp: updateNo_hp}
					updateCustomer(customer)

				case 2:

					fmt.Print("Enter Pelyanan Id to update: ")
					scanner.Scan()
					layId, _ := strconv.Atoi(scanner.Text())

					if !getPelyananId(layId) {
						fmt.Println("Id Pelayanan tidak ditemukan.")
						continue mainMenuLoop3
					}

					fmt.Print("Masukan Jenis Pelayanan :")
					scanner.Scan()
					updateJp := scanner.Text()

					fmt.Print("Masukan Satuan (KG/Buah):")
					scanner.Scan()
					updateSt := scanner.Text()

					if updateSt != "Buah" && updateSt != "KG" {
						fmt.Println("Harap Memasukan Hanya KG/Buah.")
						continue mainMenuLoop3
					}

					fmt.Print("Masukan Harga:")
					scanner.Scan()
					updateHrg, _ := strconv.Atoi(scanner.Text())

					pelayanan := entity.Pelayanan{Id: layId, Jenis_pelayanan: updateJp, Satuan: updateSt, Harga: updateHrg}
					updatePelayanan(pelayanan)

				case 3:
					continue mainMenuLoop
				default:
					fmt.Println("invalid sub-Choice")
				}

			}
		case 4:
			for {
				fmt.Println("Select Choice To Delete Data")
				fmt.Println("1. Delete Data Customer")
				fmt.Println("2. Delete Data Pelayanan")
				fmt.Println("3. Back To Menu Utama")

				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("enter you choice :")
				scanner.Scan()
				SubChoice2, _ := strconv.Atoi(scanner.Text())

				switch SubChoice2 {
				case 1:
					fmt.Print("Masukan Customer Id: ")
					scanner.Scan()
					DelCid, _ := strconv.Atoi(scanner.Text())

					deleteCustomer(DelCid)

				case 2:
					fmt.Print("Masukan Pelayanan Id: ")
					scanner.Scan()
					delPid, _ := strconv.Atoi(scanner.Text())

					deletePelayanan(delPid)

				case 4:
					continue mainMenuLoop
				default:
					fmt.Println("invalid sub-Choice")
				}
			}
		case 5:
			os.Exit(0)
		default:
			fmt.Println("invalid choice")

		}

	}
}

func getAllCustomer() []entity.Customer {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_customer;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomer(rows)

	return customers
}

func scanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name_customer, &customer.No_hp)
		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return customers
}

func addCustomer(customer entity.Customer) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "INSERT INTO mst_customer (id, nama_customer, no_hp) VALUES ($1, $2, $3);"

	_, err := db.Exec(sqlStatement, customer.Id, customer.Name_customer, customer.No_hp)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data Customer")
	}

}

func isCustomerExists(customerID int) bool {
	db := connectDb()
	defer db.Close()

	query := "SELECT COUNT(*) FROM mst_customer WHERE id = $1"
	var count int

	err := db.QueryRow(query, customerID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func updateCustomer(customer entity.Customer) {

	db := connectDb()
	defer db.Close()

	sqlStatement := "UPDATE mst_customer SET nama_customer = $2, no_hp = $3 WHERE id=$1; "

	_, err := db.Exec(sqlStatement, customer.Id, customer.Name_customer, customer.No_hp)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Updated DAta Customer")
	}
}

func deleteCustomer(id int) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "DELETE FROM mst_pelayanan WHERE id = $1;"

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data Customer")
	}

}
func getAllPelayanan() []entity.Pelayanan {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_pelayanan;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	pelayanans := scanPelayanan(rows)

	return pelayanans
}

func scanPelayanan(rows *sql.Rows) []entity.Pelayanan {
	pelayanans := []entity.Pelayanan{}
	var err error

	for rows.Next() {
		pelayanan := entity.Pelayanan{}
		err := rows.Scan(&pelayanan.Id, &pelayanan.Jenis_pelayanan, &pelayanan.Satuan, &pelayanan.Harga)
		if err != nil {
			panic(err)
		}

		pelayanans = append(pelayanans, pelayanan)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return pelayanans
}

func addPelayanan(pelayanan entity.Pelayanan) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "INSERT INTO mst_pelayanan (id, jenis_pelayanan, satuan, harga) VALUES ($1, $2, $3, $4);"

	_, err := db.Exec(sqlStatement, pelayanan.Id, pelayanan.Jenis_pelayanan, pelayanan.Satuan, pelayanan.Harga)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data Pelayanan")
	}

}

func getPelyananId(pelayananID int) bool {
	db := connectDb()
	defer db.Close()

	query := "SELECT COUNT(*) FROM mst_pelayanan WHERE id = $1"
	var count int

	err := db.QueryRow(query, pelayananID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func updatePelayanan(pelayanan entity.Pelayanan) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "UPDATE mst_pelayanan SET jenis_pelayanan = $2, satuan = $3, harga = $4 WHERE id=$1; "

	_, err := db.Exec(sqlStatement, pelayanan.Id, pelayanan.Jenis_pelayanan, pelayanan.Satuan, pelayanan.Harga)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Updated Data Pelayanan")
	}
}

func deletePelayanan(id int) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "DELETE FROM mst_pelayanan WHERE id = $1;"

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data Pelayanan")
	}

}
func getAlltrxBill() []entity.TrxBill {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM trx_bill;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	trxBills := scanTrxBill(rows)

	return trxBills
}

func scanTrxBill(rows *sql.Rows) []entity.TrxBill {
	trxBills := []entity.TrxBill{}
	var err error

	for rows.Next() {
		trxBill := entity.TrxBill{}
		err := rows.Scan(&trxBill.Id, &trxBill.No, &trxBill.Tanggal_masuk, &trxBill.Diterima_oleh)
		if err != nil {
			panic(err)
		}

		trxBills = append(trxBills, trxBill)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return trxBills
}

func addTxBill(trxBill entity.TrxBill) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "INSERT INTO trx_bill (id, no, tanggal_masuk, diterima_oleh) VALUES ($1, $2, $3, $4);"

	_, err := db.Exec(sqlStatement, trxBill.Id, trxBill.No, trxBill.Tanggal_masuk, trxBill.Diterima_oleh)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("succesfully Insert Data Transaction Bill ")
	}

}

func getAlltrxBillDetail() []entity.TrxBillDetail {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM trx_bill_detail;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	trxBillDetails := scanTrxBillDetail(rows)

	return trxBillDetails
}

func scanTrxBillDetail(rows *sql.Rows) []entity.TrxBillDetail {
	trxBillDeatails := []entity.TrxBillDetail{}
	var err error

	for rows.Next() {
		trxBillDetail := entity.TrxBillDetail{}
		err := rows.Scan(&trxBillDetail.Id, &trxBillDetail.Customer_id, &trxBillDetail.Pelayanan_id, &trxBillDetail.Trx_bill_id, &trxBillDetail.Jumlah, &trxBillDetail.Tanggal_keluar)
		if err != nil {
			panic(err)
		}

		trxBillDeatails = append(trxBillDeatails, trxBillDetail)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return trxBillDeatails
}

func addTxBillDetail(trxBillDetail entity.TrxBillDetail) {
	db := connectDb()
	defer db.Close()

	sqlStatement := "INSERT INTO trx_bill_detail (id, customer_id,pelayanan_id, trx_bill_id, jumlah, tanggal_keluar) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := db.Exec(sqlStatement, trxBillDetail.Id, trxBillDetail.Customer_id, trxBillDetail.Pelayanan_id, trxBillDetail.Trx_bill_id, trxBillDetail.Jumlah, trxBillDetail.Tanggal_keluar)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data Transaction Bill Detail")
	}

}
func totalBayar() []float64 {
	db := connectDb()
	defer db.Close()

	query := "SELECT tbd.jumlah * mp.harga AS total from trx_bill_detail tbd JOIN mst_pelayanan mp ON tbd.customer_id = mp.id;"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var results []float64

	for rows.Next() {
		var total float64
		err = rows.Scan(&total)
		if err != nil {
			panic(err)
		}
		results = append(results, total)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return results
}

func totalBayarKeseluruhan() []float64 {
	db := connectDb()
	defer db.Close()

	query := "SELECT SUM(tbd.jumlah * mp.harga) AS total from trx_bill_detail tbd JOIN mst_pelayanan mp ON tbd.customer_id = mp.id;"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var results []float64

	for rows.Next() {
		var total float64
		err = rows.Scan(&total)
		if err != nil {
			panic(err)
		}
		results = append(results, total)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return results
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Connected")
	}
	return db
}
