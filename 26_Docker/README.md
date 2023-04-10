Docker

Sebuah aplikasi open-source yang memungkinkan untuk dapat membuat, menguji, dan menerapkan aplikasi secara cepat.

Container

Sebuah wadah untuk mengemas dan menjalankan suatu aplikasi. container ini mencakup kode, runtime, system tools, dan pengaturan.

Strategi Deployment
1. Big-Bang Deployment Strategy

    kelebihan :
    1. mudah diimplementasikan. tinggal replace.
    2. perubahan ke system langsung 100% secara instan.

   Kekurangan :
   1. Terlalu beresiko, rata - rata downtime cukup lama.
 2. Rollout Deployment Strategy

    Kelebihan :
    1. Lebih aman dan less downtime dari versi sebelumnya.

    Kekurangan :
    1. Akan ada 2 versi aplikasi yang berjalan berbarengan sampai server terdeploy, dan bisa membuat bingung.
    2. Karena sifatnya satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan sampai semua server terken efeknya.
    3. Tidak ada kontrol request. Server yang baru ter-deploy dengan aplikasi versi baru, langsung mendapat request yang sama banyaknya dengan server lainnya.
3. Blue / Green Deployment Strategy

    Kelebihan :
    1. Perubahan sangat cepat, sekali switch service langsung berubah 100%
    2. Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment.

   Kekurangan :
   1. Resource yang dibutuhkan lebih banyak. karena setiap Deployment harus menyediakan service yang serupa environmentnya dengan yang sedang berjalan di production.
   2. Testing harus benar-benar diprioritaskan sebelum di switch, aplikasi harus dipastikan aman dari request yang tiba-tiba banyak.
4. Canary Deployment Strategy

   Kelebihan :
   1. Cukup aman
   2. Mudah untuk rollback jika terjadi error / bug, tanpa berimbas kesemua server

    Kekurangan :
    1. Untuk mencapai 100% butuh waktu lama dibandingkan dengan Blue / Green Deployment. Blue / Green Deployment, aplikasi dapat langsung 100% terdeploy keseluruh user.