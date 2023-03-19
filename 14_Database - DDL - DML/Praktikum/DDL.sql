/*2.a membuat table user*/

CREATE TABLE users ( 
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    gender char(1) NOT NULL,
    dob DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

/*2.b Membuat table payment_method*/

 CREATE TABLE payment_methods (
    id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

/*2.b Membuat table operators*/

CREATE TABLE operators (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

/*2.b Membuat table product_types*/

CREATE TABLE product_types (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

/*2.c Membuat table transaction*/

CREATE TABLE transactions (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id int(11),
    payment_method_id int(11),
    status varchar(10),
    total_qty int(11),
    total_price NUMERIC(25,2),
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);

/*2.b Membuat table products*/

CREATE TABLE products (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_type_id int(11),
    operator_id int(11),
    code varchar(50), 
    name varchar(255),
    status smallint,
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(product_type_id) REFERENCES product_types(id),
	FOREIGN KEY(operator_id) REFERENCES operators(id)
);

/*2.b Membuat table product_descriptions*/

CREATE TABLE product_descriptions (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	product_id int(11),
    description TEXT,
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(product_id) REFERENCES products(id)
);

/*2.c Membuat table transaction_details*/

CREATE TABLE transaction_details (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	transaction_id int(11),
    product_id int(11),
    status varchar(10),
	qty int(11),
    price NUMERIC(25,2),
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(transaction_id) REFERENCES transactions(id),
    FOREIGN KEY(product_id) REFERENCES products(id)
);

/*3. Membuat table kurirs*/

CREATE TABLE kurirs (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()    
);

/* 4. Menambah column di table kurirs*/
ALTER TABLE kurirs
ADD ongkos_dasar decimal;

/* 5. Rename table kurirs*/

ALTER TABLE kurirs
RENAME TO shipping;

/* 6. Hapus table kurirs*/

DROP TABLE shipping;

/* 7.a Membuat table payment_method_descriptions*/

CREATE TABLE payment_method_descriptions (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    payment_method_id int(11),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);

/* 7.b. Membuat table alamats*/

CREATE TABLE alamats (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id int(11),
    alamat varchar(30),
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

/* 7.c Membuat table user_payment_method_detail*/

CREATE TABLE user_payment_method_details (
	id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id int(11),
    payment_method_id int(11),
	created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);
