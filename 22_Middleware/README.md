Middleware

Sebuah entitas yang terhubung ke pemrosesan request dan response server. middleware ini dapat digunakan untuk mengimplementasikan berbagai fungsi di antara request HTTP yang masuk ke server dan sebelum response.

Third party middleware antaralain :

1. Negroni
2. Echo (https://echo.labstack.com/middleware/)
3. Interpose
4. Alice

dalam middleware terdapat :
1. log middleware.

   memberikan log informasi terkait HTTP request, dan sebagai datasource untuk analisis
2. auth middleware

    sebuah authentifikasi untuk mengidentifikasi User dan mengamankan data dan informasi.
    - basic auth
    - JWT auth