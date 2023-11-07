--DDl
create table public.mst_customer (
	id int not null primary key,
	nama_customer varchar(100) not null,
	no_hp int not null
);

create table public.mst_pelayanan (
	id int not null primary key,
	jenis_pelayanan varchar(100) not null,
	satuan varchar(100) not null,
	harga int not null
);

create table public.trx_bill (
	id int not null primary key,
	no int not null,
	tanggal_masuk date not null,
	diterima_oleh varchar(100) not null
);

create table public.trx_bill_detail (
	id int not null primary key,
	customer_id int not null,
	pelayanan_id int not null,
	trx_bill_id int not null,
	jumlah int not null
);

alter table trx_bill_detail 
add constraint fk_customer
foreign key (customer_id)
references mst_customer(id)

ALTER TABLE public.mst_customer ALTER COLUMN no_hp TYPE varchar(12) USING no_hp::varchar(12);

select tbd.jumlah * mp.harga as total 
from trx_bill_detail tbd 
join mst_pelayanan mp on tbd.customer_id = mp.id;

select sum(tbd.jumlah * mp.harga)  as total_keseluruhan
from trx_bill_detail tbd 
join mst_pelayanan mp on tbd.customer_id = mp.id;