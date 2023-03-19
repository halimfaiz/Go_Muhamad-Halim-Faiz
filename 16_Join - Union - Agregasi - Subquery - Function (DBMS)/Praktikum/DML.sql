/* Insert 5 operators pada table operators */
INSERT INTO operators (name) values ('TOKO A');
INSERT INTO operators (name) values ('TOKO B');
INSERT INTO operators (name) values ('TOKO C');
INSERT INTO operators (name) values ('TOKO D');
INSERT INTO operators (name) values ('TOKO E');

/* Insert 3 product type */
INSERT INTO product_types (name) values ('Maknanan');
INSERT INTO product_types (name) values ('Baju');
INSERT INTO product_types (name) values ('Minuman');

/* Insert 2 product dengan product type id = 1, dan operators id = 3 */
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (1,3,'0101', 'Taro',1);
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (1,3,'0102', 'Taro 1',1);

/* Insert 3 product dengan product type id = 2, dan operators id = 1 */
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (2,1,'0103', 'kaos',1);
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (2,1,'0104', 'kaos 1',1);
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (2,1,'0105', 'kaos 2',1);

/* Insert 3 product dengan product type id = 3, dan operators id = 4 */
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (3,4,'0106', 'cola',1);
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (3,4,'0107', 'cola 1',1);
INSERT INTO products (product_type_id, operator_id, code, name, status)
VALUES (3,4,'0108', 'cola 2',1);

/* Insert product description pada setiap product */
INSERT INTO product_descriptions (product_id, description)
VALUES (1,'Makanan Ringan');
INSERT INTO product_descriptions (product_id, description)
VALUES (2,'Makanan Ringan');
INSERT INTO product_descriptions (product_id, description)
VALUES (3,'Kaos Polos Cotton Combed 30s');
INSERT INTO product_descriptions (product_id, description)
VALUES (4,'Kaos Polos Cotton Combed 30s');
INSERT INTO product_descriptions (product_id, description)
VALUES (5,'Kaos Polos Cotton Combed 30s');
INSERT INTO product_descriptions (product_id, description)
VALUES (6,'Minuman Ringan Berkarbonasi');
INSERT INTO product_descriptions (product_id, description)
VALUES (7,'Minuman Ringan Berkarbonasi');
INSERT INTO product_descriptions (product_id, description)
VALUES (7,'Minuman Ringan Berkarbonasi');

/* Insert 3 payment methods */
INSERT INTO payment_methods (name, status) 
VALUES ('DANA',1);
INSERT INTO payment_methods (name, status) 
VALUES ('GOPAY',2);
INSERT INTO payment_methods (name, status) 
VALUES ('OVO',3);

/* Insert 5 user pada tabel user */
INSERT INTO users (name,gender, dob)
VALUES ('Halim','L', '2002-08-29');
INSERT INTO users (name,gender, dob)
VALUES ('Nun','P', '2001-11-11');
INSERT INTO users (name,gender, dob)
VALUES ('Bim','L', '2002-12-12');
INSERT INTO users (name,gender, dob)
VALUES ('Rin','P', '2005-05-15');
INSERT INTO users (name,gender, dob)
VALUES ('Ray','L', '2006-06-16');

/* Insert 3 transaksi di masing-masing user */
/* USER 1 */
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (1,1,'BERHASIL',1, 2000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (1,2,'BERHASIL',2, 20000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (1,3,'BERHASIL',3, 7500);

/* USER 2 */
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (2,1,'BERHASIL',1, 2000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (2,2,'BERHASIL',2, 20000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (2,3,'BERHASIL',3, 7500);

/* USER 3 */
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (3,1,'BERHASIL',1, 2000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (3,2,'BERHASIL',2, 20000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (3,3,'BERHASIL',3, 7500);

/* USER 4 */
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (4,1,'BERHASIL',1, 2000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (4,2,'BERHASIL',2, 20000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (4,3,'BERHASIL',3, 7500);

/* USER 5 */
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (5,1,'BERHASIL',1, 2000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (5,2,'BERHASIL',2, 20000);
INSERT INTO transactions (user_id,payment_method_id,status,total_qty,total_price)
VALUES (5,3,'BERHASIL',3, 7500);


/* Insert 3 product di masing-masing transaksi */
/* USER 1 */
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (1,1,'BERHASIL',1,2000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (1,2,'BERHASIL',1,10000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (1,3,'BERHASIL',1,2500);

/* USER 2 */
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (2,1,'BERHASIL',2,2000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (2,2,'BERHASIL',2,10000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (2,3,'BERHASIL',2,2500);

/* USER 3 */
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (3,1,'BERHASIL',3,2000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (3,2,'BERHASIL',3,10000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (3,3,'BERHASIL',3,2500);

/* USER 4 */
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (4,1,'BERHASIL',3,2000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (4,2,'BERHASIL',3,10000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (4,3,'BERHASIL',3,2500);

/* USER 5 */
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (5,1,'BERHASIL',1,2000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (5,2,'BERHASIL',1,10000);
INSERT INTO transaction_details (transaction_id,product_id,status,qty,price)
VALUES (5,3,'BERHASIL',1,2500);


/* Menampilkan nama user / pelanggan dengan gender Laki-laki */
select * from users where gender='L';

/* Menampilkan product dengan id = 3 */
SELECT * FROM products WHERE id = 3;

/*Menampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’. */
SELECT * FROM users WHERE name LIKE '%a%'
AND datediff(NOW(), created_at) <= 7
AND DATE(created_at) >= DATE(NOW() - INTERVAL 7 DAY);

/*Menghitung jumlah user / pelanggan dengan status gender Perempuan */
SELECT count(id) FROM users WHERE gender = "p";

/* Menampilkan data pelanggan dengan urutan sesuai nama abjad */
SELECT * FROM users ORDER BY name ASC;

/* Menampilkan 5 data pada data product */
SELECT * FROM products LIMIT 5;

/* Mengubah data product id 1 dengan nama ‘product dummy’ */
UPDATE products SET name = 'product dummy' WHERE id = 1;
SELECT * FROM products WHERE id = 1;

/* Mengupdate qty = 3 pada transaction detail dengan product id = 1*/
UPDATE transaction_details SET qty = 3 WHERE product_id =1;
SELECT * FROM transaction_details WHERE  product_id = 1;

/* Delete data pada tabel product dengan id = 1 */
ALTER TABLE products
ADD COLUMN deleted_at DATETIME NULL;
UPDATE products SET deleted_at = NOW() WHERE id = 1;
SELECT * FROM products;

/* Delete pada pada tabel product dengan product type id 1 */
UPDATE products SET deleted_at = NOW() WHERE product_type_id = 1;
SELECT * FROM products;
