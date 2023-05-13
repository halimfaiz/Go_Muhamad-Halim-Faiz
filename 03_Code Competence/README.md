Aplikasi ini adalah Sebuah code Backend untuk toko barang elektronik.

Memiliki fungsi utama untuk memudahkan pengguna dalan mengelola data barang elektronik

Feature :
1. Register dan Login untuk mengakses fitur lainnya.
3. Create Category
4. Create Item
5. Get Item By ID
6. Get All Items
7. Get All Items By ID Category
8. Get All Items By Item Name
9. Update Item
10. Delete Item

Guide :
1. jangan lupa setelah meng clone ubah configurasi database di Config sesuai dengan database yang digunakan di local.
2. untuk mendapatkan akses dalam menggunakan fitur-fitur contohnya seperti Get All Items maka perlu register dan login untuk mendapatkan token JWT, setelah itu pakai `localhost:8080/items` dengan method GET serta token yang dimasukan pada auth di postman untuk pengetesan ubah localhost sesuai dengan DB Host di Config dan ubah 8080 sesuai port di main.go