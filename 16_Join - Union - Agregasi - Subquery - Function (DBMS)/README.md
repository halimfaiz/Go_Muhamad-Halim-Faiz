1. Join
   untuk mengombinasikan record dari 2 tabel atau lebih.
   - Inner Join (Irisan)
   - Left Join (semua data yang berada di kiri Termasuk data di kanan yang merupakan bagian dari data kiri)
   - Right Join(semua data yang berada di kanan Termasuk data di kiri yang merupakan bagian dari data kanan)

2. Aggregate
   - MIN
   - MAX
   - SUM
   - AVG
   - COUNT

3. Subquery
   Dapat digunakan dengan :
   - SELECT
   - INSERT
   - UPDATE
   - DELETE

    Bersama dengan operator seperti :
    - =         (sama dengan)
    - <,>       (lebih kecil dari, lebih besar dari)
    - '>=,<=    (lebih besar dari sama dengan, lebih kecil dari sama dengan)

4. Function
   - function untuk mengembalikan jumlah data dari tweet per user

            DELIMITER $$
            CREATE FUNCTION sf_count_tweer_peruser
            (user_id_p_ INT)
            RETURNS IN DETERMINISTIC
            BEGIN
            DECLARE toal INT;
            SELECT COUNT (*) INTO total FROM tweets
            WHERE user_id = user_id_p AND type = 'tweets';
            RETURN total;
            END $$
            DELIMITER;

        1. Delimiter

            Memberitahu MYSQL Delimiter yang digunakan. pada contoh diatas Delimiter yang digunakan adalah $$
        2. CREATE FUNCTION =

            Header untuk membuat function

        3. RETURN

            Untuk menentukan tipe data yang akan di Return oleh function
        4. DETERMINISTIC / NOT DETERMINISTIC

            Untuk menentukan yang bisa menggunakan function ini adalah pembuatnya saja (Deterministic) atau user siapa saja (Not Deterministic)
        5. BEGIN END

            Merupakan Body dari function jadi semua SQL nya di tulis disini
   - function untuk Delete data yang berhubungan dengan TABLE users yang sedang di Delete datanya.

            Delimiter $$
            CREATE TRIGGER delete_all_data_user
            BEFORE DELETE ON  users FOR EACH ROW
            BEGIN
            --Declare variables--
            DECLARE v_user_id INT;
            SET v_user_id = OLD.id;
            --Trigger Code--
            DELETE FROM tweets WHERE user_id = v_user_id;
            DELETE FROM user_followers WHERE user_id = v_user_id;
            END $$

        6. DELIMITER

            Memberitahu MYSQL Delimiter yang digunakan. pada contoh diatas Delimiter yang digunakan adalah $$
        7. CREATE TRIGGER

            Header untuk membuat Trigger Function
        8. DECLARE

            Merupakan Syntax untuk mendeclare kan suatu hal / variabel
        9. OLD

            Merupakan Variabel yang menyimpan nilai dari dalam data yang Sedang berinteraksi / dipanggil
        10. NEW

            Merupakan Variabel yang menyimpan nilai dari dalam data yang Baru masuk / sedang di input
        11. BEGIN END

            Merupakan Body dari Function jadi semua SQL nya di tulis disini