1. Recursive

    Merupakan suatu keadaan dimana sebuah function dapat menyelesaikan masalah dengan cara memanggil dirinya sendiri dengan scope masalah yang lebih kecil.

          func factorial(number int) int {
              if number == 1 {
                return number
              } else {
              return number * factorial(number - 1)
             }
          }

    Beberapa masalah dapat diselesaikan dengan mudah menggunakan recursive dan juga penulisan code menjadi lebih pendek. akan tetapi membutuhkan pemahaman terhadap pemecahan masalahnya sebelum melakukan perulangan recursive.

2. Number Theory

    Salah satu percabangan dari ilmu matematika mengenai bilangan. beberapa contohnya seperti :
    - Faktorial
    - Bilangan Prima
    - FPB (faktor persekutuan teresar)
    - KPK (kelipatan persekutuan terkecil)

3. Sort

    Digunakan untuk mengurutkan data. salah satu jenis pengurutannya yaitu :
    - Selection Sort - O (n^2)
    - Counting Sort - O (n + k)
    - Merge Sort - O (log n)

    Package sort dapat digunakan untuk memudahkan dalam mengurutkan tipe data.
