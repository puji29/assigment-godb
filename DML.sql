--DML
insert into mst_customer (id, nama_customer, no_hp) values (1, 'Jessica', '08598798712');
insert into mst_customer (id, nama_customer, no_hp) values (2, 'Puji', '08598798713');
insert into mst_pelayanan (id, jenis_pelayanan, satuan, harga) values (1, 'Cuci + Strika', 'KG','7000');
insert into trx_bill (id, no, tanggal_masuk, diterima_oleh) values (1, 123,'20023-11-07','mirna');
insert  into  trx_bill_detail (id, customer_id, pelayanan_id, trx_bill_id, jumlah, tanggal_keluar) values (1,1,1,1, 5,'2023-11-07');

update mst_customer set nama_customer = 'adit' where id= 2;

delete from mst_customer where id= 2;