ORM (Object-Relational-Mapping)

sebuah cara untuk dapat mengkonversi data menggunakan bahasa pemrograman yang berorientasi objek. dalam golang menggunakan Struct.

GORM

merupakan library golang yang membantu dalam pembuatan ORM

Kelebihan :

- mengurangi query yang repetitive
- otomatis fetch data
- cara simple untuk melakukan screening data sebelum dimasukkan ke database

Kekurangan :

- perlu penulisan code yang cukup banyak dalam prosesnya.
- beberapa fungsi di SQL tidak dapat digunakan

Database Migration

- Sebuah cara untuk memperbaharui versi database untuk mengikuti perubahan versi aplikasi. perubahan akan selalu sama.
- Memudahkan untuk update versi atau rollback ke versi sebelumnya.
- Dapat melacak perubahan pada struktur database
Riwayat struktur database yang ditulis pada kode
- selalu compatible dengan perubahan versi yang terjadi

LIBRARY : https://gorm.io/docs/

Install Gorm:  go get -u gorm.io/gorm

Install MySQL Driver: go get -u gorm.io/driver/mysql


connect MySQL menggunakan GORM :

     <username>:<password>@/<db_name>?charset=utf8&parseTime=True&loc=Local
