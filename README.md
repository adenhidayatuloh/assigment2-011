# assigment2-011

Nama	: Aden Hidayatuloh
Kode Peserta : GLNG-KS08-011

Table Database : 

create table Orders (

	Order_id serial primary key,
	Customer_name varchar (225) not null,
	Ordered_at timestamptz default now(),
	Created_at timestamptz default now(),
	Updated_at timestamptz default now()

	
);

create table Items (

	Item_id serial primary key,
	Item_code varchar (225) not null,
	Quantity int not null,
	description text not null,
	Order_id int not null,
	Created_at timestamptz default now(),
	Updated_at timestamptz default now(),
	constraint Items_Order_id_fk
	foreign key (Order_id) references Orders(Order_id) on delete cascade

	
);
