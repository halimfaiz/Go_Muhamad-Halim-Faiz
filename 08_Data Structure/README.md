Data Structure
- Array
  dapat menampung 1 tipe data dengan jumlah fix yang ditentukan ketika deklarasi
  ex : [5]string{}
       [...]string{}
       "[...]" dapat digunakan ketika tidak tahu butuh berapa panjang array
    looping array :
    primes := [5]int{1,2,3,4,5}
    for _, value := range primes {
                        fmt.Println(value)
                        }
- Slice
  mirip seperti array tetapi tidak memerlukan deklarasi jumlah capacity
  ex : numbers :=[]int{1,2,3,4,5} || var numbers = make([]int,5,10) || var part_primes []int = primes[1:4]
  dengan mendeklarasi 5 element, ketika append element ke ex pertama panjang slice akan menjadi 6, dan capacitynya meningkat 2x lipat menjadi 10
  dan ex kedua menjelaskan bahwa dengan menggunakan 'make' dapat mengatur panjang slice yakni 5 dan capacity slice yakni 10
  serta ex ke-3 menunjukan pembuatan slice dari array diatas. [1:4] menunjukkan index milik array yang akan diambil ke slice dengan hasil [2,3,4]

- append(): menambahkan element ke slice
  copy(): mengcopy slice ke slice lain

- Map
  map sedikit berbeda karena dia memiliki Key dan Value yang mana Key nya selalu unique tidak bisa sama
  1. ex :salary_b := map[string]int{}
   salary_b ["surabaya"] = 1 //menambah key n value di map
      delete(salary_b["surabaya"]) //menghapus value menggunakan key
      looping map :
      for key, value := range salary_b {
        fmt.println (key, value)
      }

- function declaration
  ex : func sum(number int)int{
       return number + number
  }
  dengan adanya func sum diatas dengan parameter 'number' tipe data integer dan return integer penulisan di func main menjadi seperti berikut :
        func main (){
            fmt.println(sum(2))
        }
        maka akan menjadi 2+2 = 4

