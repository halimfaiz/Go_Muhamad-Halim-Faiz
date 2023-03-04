1. Package strings

   a.) Len

        Digunakan untuk mengukur panjang character suatu string.

    b.) Compare

        untuk membandingkan tiap character antara dua string.

    c.) Contains

        Digunakan untuk membandingkan apakah satu string merupakan bagian dari string lain. atau apakah satu string dengan character panjang mengandung semua character pada string yang lebih pendek.

    d.) Replace

        Penggunaan replace untuk menggantikan character suatu string menjadi character baru atau pun menghapusnya menjadi string kosong.

        contoh :

        s := "apa kabar"
        t := strings.Replace(s,"a","i", -1)
        fmt.printf("%s\n", t)

        outputnya akan menjadi "ipi kibir"
        strings.replace untuk merubah / mereplace character string.
        - s merupakan variabel string yang akan diubah
        - "a" merupakan character string lama
        - "i" meruapakan character string baru
        - -1 digunakan untuk merubah semua "a" menjadi "i"

2. Struct

   Merupakan type yang berisikan property atau function(methods)

        Deklarasi struct Person dengan property name dan age :

        type Person struct {
        	name string
	        age  int
        }

        func main() {
	    var person1 = Person{
		name: "Halim",
		age:  20,
	        }
	    fmt.Println(person1)
        }

3. Method
   Merupakan function yang melekat pada suatu type (seperti Struct atau type data lainnya).

        Deklarasi Method :
        func (receiver StructType) functionName(input type) returnType {
        // block statement method
        }

        Deklarasi Function :
        func functionName(input type) returnType {
        // block statement function
        }

    Alasan menggunakan Method lebih baik dibanding function :

    1. Method membantu untuk melakukan Object-Oriented style code di Golang.

    2. Method membantu untuk menghindari permasalahan penamaan pada code. Meski dengan nama method sama, jika objectnya berbeda akan memanggil function yang berbeda.

    3. Method calls lebih mudah di baca dan dipahami daripada function calls


