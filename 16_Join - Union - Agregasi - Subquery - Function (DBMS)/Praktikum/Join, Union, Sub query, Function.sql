/*gabungan data transaksi user 1 dan 2  */
SELECT * FROM transactions WHERE user_id IN(1,2);

/* Jumlah transaksi user id 1*/
select user_id, Sum(total_price) as total_price FROM transactions WHERE user_id = 1;

/*  Total transaksi dengan product type 2 */
SELECT count(*) FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON p.id = td.product_id
WHERE p.product_type_id = 2;

/*Menampilkan semua field table product dan field name table product type yang saling berhubungan. */
SELECT p.*, pt.name FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

/* Menampilkan semua field table transaction, field name table product dan field name table user*/
SELECT t.*, p.name AS products, u.name AS users FROM transaction_details td
JOIN transactions t ON t.id = td.transaction_id
JOIN users u ON u.id = t.user_id
JOIN products p ON p.id = td.product_id
WHERE p.deleted_at IS NULL

/* function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud */
 delimiter $$
 CREATE TRIGGER delete_transaction_details
 BEFORE DELETE ON transactions FOR EACH ROW
 BEGIN
 DECLARE trans_id INT;
 SET trans_id = OLD.id;
 DELETE FROM transaction_details WHERE transaction_id = trans_id;
 END $$
 DELETE FROM transactions WHERE id = 5;
 SELECT * FROM transaction_details;
 
 /* function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.*/
 delimiter $$
 CREATE TRIGGER update_transaction_id
 AFTER DELETE ON transaction_details FOR EACH ROW
 BEGIN
 DECLARE trans_id INT;
 SET trans_id = OLD.transaction_id;
 UPDATE transactions SET total_qty = (SELECT sum(qty) FROM transaction_details WHERE transaction_id = trans_id)
 WHERE id = trans_id;
 END $$
 
DELETE FROM transaction_details WHERE transaction_id = 2 AND product_id = 1;

SELECT * FROM transaction_details WHERE transaction_id = 2;
 

 /* Menampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. */
 /* Menggunakan NOT IN */
 SELECT * FROM products WHERE id NOT IN(
 SELECT product_id FROM transaction_details
 );
 
 /* Menggunakan IN */
SELECT * FROM products WHERE id IN (
SELECT p.id FROM products p
LEFT OUTER JOIN transaction_details td on p.id = td.product_id
WHERE td.transaction_id IS NULL
);
 