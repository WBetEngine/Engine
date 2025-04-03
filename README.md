# Roadmap Pengembangan Frontend dan Backend untuk Situs Judi Online

## 1. Struktur Direktori Proyek
- **Struktur Direktori**:
  - **/WBE/**: Direktori utama proyek.
    - **frontend/**: Direktori untuk semua file frontend.
      - **admin/**: Direktori untuk bagian admin.
        - **backend/**: Backend untuk admin.
          - **database/**: Berisi file dan konfigurasi terkait database.
            - `db.go`: Mendefinisikan koneksi ke Supabase, termasuk inisialisasi dan pengelolaan pool koneksi.
            - `queries.go`: Berisi fungsi-fungsi untuk operasi database (select, insert, update, delete).
            - `migrations.go`: Mengelola migrasi database jika diperlukan.
          - **handlers/**: Berisi fungsi handler untuk memproses permintaan.
            - `auth_handler.go`: Menangani proses otentikasi admin (login/logout).
            - `user_handler.go`: Menangani operasi CRUD untuk data pengguna.
            - `dashboard_handler.go`: Menangani permintaan untuk halaman dashboard.
          - **router/**: Berisi logika routing untuk backend admin.
            - `router.go`: Mendefinisikan semua rute API dan endpoint untuk admin panel.
            - `middleware_group.go`: Mengelompokkan middleware berdasarkan kebutuhan rute.
          - **middleware/**: Berisi fungsi middleware untuk pemrosesan permintaan.
            - `auth_middleware.go`: Memeriksa otentikasi untuk rute yang dilindungi.
            - `logger_middleware.go`: Mencatat log permintaan dan respon.
            - `rate_limiter.go`: Membatasi jumlah permintaan dari satu sumber.
          - **model/**: Berisi model data dan skema.
            - `user_model.go`: Mendefinisikan struktur data pengguna.
            - `transaction_model.go`: Mendefinisikan struktur data transaksi.
            - `game_model.go`: Mendefinisikan struktur data game.
          - **template/**: Berisi template Go untuk rendering HTML.
            - `admin_templates.go`: Fungsi untuk merender template admin.
        - **assets/**: Aset untuk frontend admin.
          - **js/**: File JavaScript untuk frontend admin.
            - `admin.js`: JavaScript utama untuk fungsionalitas admin panel.
            - `dashboard.js`: JavaScript khusus untuk dashboard admin.
            - `htmx.min.js`: Library HTMX untuk interaksi tanpa reload halaman.
            - `chart.js`: Library untuk menampilkan grafik di dashboard.
          - **css/**: File CSS untuk frontend admin.
            - `admin.css`: Style khusus untuk admin panel.
            - `tailwind.css`: File CSS Tailwind yang telah di-compile.
          - **image/**: Aset gambar untuk frontend admin.
            - `logo.png`: Logo situs.
            - `favicon.ico`: Favicon untuk browser.
            - `background.jpg`: Gambar latar belakang.
        - **template/**: Template HTML untuk admin.
          - `base.html`: Template dasar yang berisi struktur HTML umum, header, footer, sidebar, dan navbar.
          - `login_admin.html`: Halaman login untuk admin.
          - `register_superadmin.html`: Halaman pendaftaran untuk super admin.
          - `dashboard.html`: Halaman dashboard dengan statistik situs.
          - `user_management.html`: Halaman untuk mengelola data pengguna.
          - `transaction_history.html`: Halaman untuk melihat riwayat transaksi.
          - `game_management.html`: Halaman untuk mengelola game.
          - `settings.html`: Halaman pengaturan admin panel.
      - **user/**: Direktori untuk bagian pengguna.
        - **backend/**: Backend untuk pengguna.
          - **database/**: Berisi file dan konfigurasi terkait database.
            - `db.go`: Mendefinisikan koneksi ke Supabase, termasuk inisialisasi dan pengelolaan pool koneksi.
            - `queries.go`: Berisi fungsi-fungsi untuk operasi database (select, insert, update, delete).
            - `cache.go`: Mengelola caching data menggunakan Redis.
          - **handlers/**: Berisi fungsi handler untuk memproses permintaan.
            - `auth_handler.go`: Menangani proses otentikasi pengguna (login/register).
            - `game_handler.go`: Menangani permintaan terkait game.
            - `transaction_handler.go`: Menangani permintaan terkait transaksi.
            - `profile_handler.go`: Menangani permintaan terkait profil pengguna.
          - **router/**: Berisi logika routing untuk backend pengguna.
            - `router.go`: Mendefinisikan semua rute API dan endpoint untuk frontend pengguna.
            - `middleware_group.go`: Mengelompokkan middleware berdasarkan kebutuhan rute.
          - **middleware/**: Berisi fungsi middleware untuk pemrosesan permintaan.
            - `auth_middleware.go`: Memeriksa otentikasi untuk rute yang dilindungi.
            - `logger_middleware.go`: Mencatat log permintaan dan respon.
            - `rate_limiter.go`: Membatasi jumlah permintaan dari satu sumber.
          - **model/**: Berisi model data dan skema.
            - `user_model.go`: Mendefinisikan struktur data pengguna.
            - `transaction_model.go`: Mendefinisikan struktur data transaksi.
            - `game_model.go`: Mendefinisikan struktur data game.
          - **template/**: Berisi template Go untuk rendering HTML.
            - `user_templates.go`: Fungsi untuk merender template pengguna.
        - **assets/**: Aset untuk frontend pengguna.
          - **js/**: File JavaScript untuk frontend pengguna.
            - `main.js`: JavaScript utama untuk fungsionalitas frontend pengguna.
            - `game.js`: JavaScript khusus untuk halaman game.
            - `htmx.min.js`: Library HTMX untuk interaksi tanpa reload halaman.
            - `carousel.js`: JavaScript untuk slider pada halaman beranda.
          - **css/**: File CSS untuk frontend pengguna.
            - `style.css`: Style khusus untuk frontend pengguna.
            - `tailwind.css`: File CSS Tailwind yang telah di-compile.
            - `game.css`: Style khusus untuk halaman game.
          - **image/**: Aset gambar untuk frontend pengguna.
            - `logo.png`: Logo situs.
            - `favicon.ico`: Favicon untuk browser.
            - `banner/`: Folder untuk gambar banner.
            - `games/`: Folder untuk gambar game.
        - **template/**: Template HTML untuk pengguna.
          - `base.html`: Template dasar dengan header, footer, sidebar, dan navbar.
          - `login.html`: Halaman login untuk pengguna.
          - `register.html`: Halaman pendaftaran untuk pengguna.
          - `dashboard.html`: Halaman dashboard pengguna dengan informasi akun.
          - `home.html`: Halaman beranda dengan promosi dan game populer.
          - `history.html`: Halaman riwayat transaksi pengguna.
          - `slot.html`: Halaman untuk kategori game slot.
          - `live_casino.html`: Halaman untuk kategori live casino.
          - `togel.html`: Halaman untuk kategori togel.
          - `sports.html`: Halaman untuk kategori taruhan olahraga.
          - `fishing.html`: Halaman untuk kategori permainan memancing.
          - `crash.html`: Halaman untuk kategori game crash.
          - `poker.html`: Halaman untuk kategori poker.
          - `promo.html`: Halaman yang menampilkan promosi aktif.
          - `faq.html`: Halaman FAQ untuk bantuan pengguna.
          - `terms.html`: Halaman syarat dan ketentuan.
          - `privacy.html`: Halaman kebijakan privasi.
          - `guide.html`: Halaman panduan bermain.
          - `deposit.html`: Halaman untuk melakukan deposit.
          - `withdraw.html`: Halaman untuk melakukan penarikan dana.
          - `profile.html`: Halaman profil pengguna untuk mengedit informasi.

          <!-- Provider Game Slot -->
          - `slots_pragmatic.html`: Halaman slot dari provider Pragmatic Play.
          - `slots_pgsoft.html`: Halaman slot dari provider PG Soft.
          - `slots_joker.html`: Halaman slot dari provider Joker.
          - `slots_habanero.html`: Halaman slot dari provider Habanero.
          - `slots_cq9.html`: Halaman slot dari provider CQ9.
          - `slots_yggdrasil.html`: Halaman slot dari provider Yggdrasil.
          - `slots_microgaming.html`: Halaman slot dari provider Microgaming.
          - `slots_playtech.html`: Halaman slot dari provider Playtech.
          - `slots_spadegaming.html`: Halaman slot dari provider Spadegaming.
          - `slots_jili.html`: Halaman slot dari provider Jili.
          - `slots_hacksaw.html`: Halaman slot dari provider Hacksaw.
          - `slots_afb.html`: Halaman slot dari provider AFB.

          <!-- Provider Live Casino -->
          - `casino_pragmatic.html`: Halaman casino dari provider Pragmatic Play.
          - `casino_evolution.html`: Halaman casino dari provider Evolution.
          - `casino_joker.html`: Halaman casino dari provider Joker.
          - `casino_cq9.html`: Halaman casino dari provider CQ9.
          - `casino_microgaming.html`: Halaman casino dari provider Microgaming.
          - `casino_playtech.html`: Halaman casino dari provider Playtech.
          - `casino_spadegaming.html`: Halaman casino dari provider Spadegaming.
          - `casino_habanero.html`: Halaman casino dari provider Habanero.

          <!-- Provider Sports -->
          - `sports_afb88.html`: Halaman sports dari provider AFB88.
          - `sports_saba.html`: Halaman sports dari provider Saba.
          - `sports_568win.html`: Halaman sports dari provider 568Win.

- **Koneksi Antar File**:
  - **Frontend User Templates**:
    - `C:\laragon\www\WBE\frontend\user\template\base.html`:
      - Berisi struktur HTML dasar yang digunakan oleh semua halaman.
      - Menyertakan header, footer, sidebar, dan navbar.
      - Memuat `C:\laragon\www\WBE\frontend\user\assets\css\tailwind.css` dan `C:\laragon\www\WBE\frontend\user\assets\css\style.css`.
      - Memuat `C:\laragon\www\WBE\frontend\user\assets\js\htmx.min.js` untuk interaksi AJAX.
      - Memuat `C:\laragon\www\WBE\frontend\user\assets\js\main.js` untuk fungsionalitas umum.

    - `C:\laragon\www\WBE\frontend\user\template\dashboard.html`:
      - Memperluas `C:\laragon\www\WBE\frontend\user\template\base.html` untuk tata letak yang konsisten.
      - Menampilkan informasi pengguna dari `C:\laragon\www\WBE\frontend\user\backend\handlers\profile_handler.go`.
      - Menampilkan riwayat transaksi dari `C:\laragon\www\WBE\frontend\user\backend\handlers\transaction_handler.go`.
      - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\user\assets\js\dashboard.js` untuk interaksi dinamis.
      - Gaya diterapkan dari `C:\laragon\www\WBE\frontend\user\assets\css\dashboard.css`.

    - `C:\laragon\www\WBE\frontend\user\template\login.html`:
      - Memperluas `C:\laragon\www\WBE\frontend\user\template\base.html`.
      - Berisi form login yang mengirim data ke `C:\laragon\www\WBE\frontend\user\backend\handlers\auth_handler.go`.
      - Menggunakan HTMX untuk mengirim data login tanpa memuat ulang halaman.

    - `C:\laragon\www\WBE\frontend\user\template\register.html`:
      - Memperluas `C:\laragon\www\WBE\frontend\user\template\base.html`.
      - Berisi form registrasi yang mengirim data ke `C:\laragon\www\WBE\frontend\user\backend\handlers\auth_handler.go`.
      - Menggunakan HTMX untuk mengirim data registrasi tanpa memuat ulang halaman.

    - `C:\laragon\www\WBE\frontend\user\template\slot.html` dan halaman game lainnya:
      - Memperluas `C:\laragon\www\WBE\frontend\user\template\base.html`.
      - Mendapatkan data game dari `C:\laragon\www\WBE\frontend\user\backend\handlers\game_handler.go`.
      - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\user\assets\js\game.js`.
  
  - **Backend User**:
    - `C:\laragon\www\WBE\main.go`:
      - Titik masuk utama aplikasi Go.
      - Memuat konfigurasi dari `.env` menggunakan `C:\laragon\www\WBE\frontend\user\backend\config\config.go`.
      - Menginisialisasi koneksi database menggunakan `C:\laragon\www\WBE\frontend\user\backend\database\db.go`.
      - Memanggil `C:\laragon\www\WBE\frontend\user\backend\router\router.go` untuk menyiapkan rute.

    - `C:\laragon\www\WBE\frontend\user\backend\router\router.go`:
      - Mendefinisikan semua rute API dan halaman web.
      - Menggunakan middleware dari `C:\laragon\www\WBE\frontend\user\backend\middleware\`.
      - Memanggil handler yang sesuai dari `C:\laragon\www\WBE\frontend\user\backend\handlers\`.

    - `C:\laragon\www\WBE\frontend\user\backend\handlers\auth_handler.go`:
      - Menangani permintaan login dan registrasi.
      - Menggunakan `C:\laragon\www\WBE\frontend\user\backend\database\queries.go` untuk interaksi database.
      - Menggunakan `C:\laragon\www\WBE\frontend\user\backend\model\user_model.go` untuk struktur data pengguna.
      - Membuat token JWT untuk otentikasi.

    - `C:\laragon\www\WBE\frontend\user\backend\handlers\game_handler.go`:
      - Menangani permintaan terkait game.
      - Mengambil data game dari `C:\laragon\www\WBE\frontend\user\backend\database\queries.go`.
      - Menggunakan `C:\laragon\www\WBE\frontend\user\backend\model\game_model.go` untuk struktur data game.

    - `C:\laragon\www\WBE\frontend\user\backend\handlers\transaction_handler.go`:
      - Menangani transaksi deposit dan penarikan.
      - Menggunakan `C:\laragon\www\WBE\frontend\user\backend\database\queries.go` untuk operasi database.
      - Menggunakan `C:\laragon\www\WBE\frontend\user\backend\model\transaction_model.go`.

  - **Koneksi Database**:
    - `C:\laragon\www\WBE\frontend\user\backend\database\db.go`:
      - Menggunakan package database/sql dan driver PostgreSQL untuk koneksi ke Supabase.
      - Membaca konfigurasi koneksi dari `.env` (SUPABASE_DB_URL, dll).
      - Menyediakan fungsi untuk mendapatkan koneksi database.

    - `C:\laragon\www\WBE\frontend\user\backend\database\queries.go`:
      - Berisi fungsi SQL untuk berbagai operasi pada tabel users, transactions, dan games.
      - Memanggil `C:\laragon\www\WBE\frontend\user\backend\database\db.go` untuk mendapatkan koneksi.

    - `C:\laragon\www\WBE\frontend\user\backend\database\cache.go`:
      - Mengelola caching dengan Redis untuk mengurangi beban database.
      - Membaca konfigurasi Redis dari `.env` (REDIS_ADDR, REDIS_PASSWORD).

  - **Frontend Admin Templates**:
    - `C:\laragon\www\WBE\frontend\admin\template\base.html`:
      - Berisi struktur HTML dasar untuk admin panel.
      - Memuat `C:\laragon\www\WBE\frontend\admin\assets\css\tailwind.css` dan `C:\laragon\www\WBE\frontend\admin\assets\css\admin.css`.
      - Memuat `C:\laragon\www\WBE\frontend\admin\assets\js\htmx.min.js` dan `C:\laragon\www\WBE\frontend\admin\assets\js\admin.js`.

    - `C:\laragon\www\WBE\frontend\admin\template\dashboard.html`:
      - Memperluas `C:\laragon\www\WBE\frontend\admin\template\base.html`.
      - Menampilkan statistik situs dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\dashboard_handler.go`.
      - Menggunakan `C:\laragon\www\WBE\frontend\admin\assets\js\chart.js` untuk menampilkan grafik.

    - `C:\laragon\www\WBE\frontend\admin\template\user_management.html`:
      - Memperluas `C:\laragon\www\WBE\frontend\admin\template\base.html`.
      - Menampilkan dan mengelola data pengguna dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\user_handler.go`.
      - Menggunakan HTMX untuk operasi CRUD tanpa reload halaman.

## 2. Persiapan Proyek
- **Inisialisasi Proyek**: Buat proyek baru untuk frontend dan backend.
- **Konfigurasi Supabase**: Daftarkan akun di Supabase dan buat proyek baru. Dapatkan URL dan kunci API untuk integrasi.
- **Konfigurasi .env**:
  - Berikut adalah konfigurasi lengkap dari file `.env`:
    ```
    # Supabase Database
    SUPABASE_DB_URL=postgresql://postgres.lepbutgfaszhvcclbgcy:Qwewenz1*@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres
    SUPABASE_HOST=aws-0-ap-southeast-1.pooler.supabase.com
    SUPABASE_PORT=6543
    SUPABASE_USER=postgres.lepbutgfaszhvcclbgcy
    SUPABASE_PASSWORD=Qwewenz1*
    SUPABASE_DBNAME=postgres
    SUPABASE_SSLMODE=require

    # JWT Configuration
    JWT_SECRET=sh_admin_panel_secret_key_2024
    JWT_EXPIRATION=24h
    REFRESH_SECRET=sh_admin_panel_refresh_key_2024
    REFRESH_EXPIRATION=168h

    # Redis
    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=

    # Server
    PORT=8080
    ```

## 3. Pengembangan Frontend User
- **Struktur Proyek**: Buat struktur direktori untuk file HTML, CSS, dan JavaScript.
- **Buat `base.html`**: 
  - Buat template dasar yang akan digunakan oleh halaman lain.
  - Sertakan elemen umum seperti header, footer, sidebar, dan navbar.
  - Gunakan Tailwind CSS untuk styling elemen-elemen ini.
- **Buat `home.html`**:
  - Desain halaman beranda dengan informasi dasar dan navigasi.
- **Buat `register.html`**:
  - **Struktur Formulir Pendaftaran di `C:\laragon\www\WBE\frontend\user\template\register.html`**
  - **File JavaScript yang Diperlukan**:
    - `C:\laragon\www\WBE\frontend\user\assets\js\validation.js`: Berisi fungsi validasi umum untuk semua form
    - `C:\laragon\www\WBE\frontend\user\assets\js\registration.js`: Berisi logika khusus untuk halaman pendaftaran
    - `C:\laragon\www\WBE\frontend\user\assets\js\utils.js`: Berisi fungsi utilitas seperti debounce
    - `C:\laragon\www\WBE\frontend\user\assets\js\api.js`: Berisi fungsi untuk memanggil API
  - **Formulir Wajib**:
    - Username Login: Wajib, minimal 4 huruf atau angka
    - Email: Wajib, harus berformat @gmail.com
    - Nomor HP: Wajib, format +62
    - Nama Lengkap: Wajib, tidak boleh ada angka
    - Kata Sandi & Konfirmasi Kata Sandi: Minimal 8 karakter, dilengkapi tombol toggle untuk show/hidden password
    - Mata Uang: Otomatis IDR (tidak bisa diedit)
    - Kode Referral: Opsional
  - **Opsi Metode Deposit (Pilih Salah Satu)**:
    - Bank Transfer: Pilihan BCA, Mandiri, BNI, BRI, Mestika, dll.
    - E-Wallet: Pilihan DANA, OVO, GoPay, ShopeePay, LinkAja, Flip, dll.
    - Pulsa: Pilihan Telkomsel, XL
  - **Keamanan & Syarat**:
    - Checkbox Persetujuan Syarat & Ketentuan
    - Tombol Register
    - Tautan ke Halaman Login untuk yang sudah punya akun
  - **Validasi Frontend**:
    - Semua input pendaftaran menggunakan JavaScript untuk validasi real-time saat pengguna mengetik (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\registration.js`)
    - Mengimplementasikan debounce untuk mengurangi spam request API saat validasi (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\utils.js`)
    - Menampilkan pesan error langsung di bawah field yang tidak valid (styling di `C:\laragon\www\WBE\frontend\user\assets\css\forms.css`)
  - **Pengiriman Data ke Backend**:
    - Saat menekan tombol "Register", data dikirim ke backend melalui HTTP POST request ke endpoint API pendaftaran (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\api.js`)
    - Data dikirim dalam format JSON
    - Menampilkan loading spinner saat pendaftaran berlangsung (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\ui.js`)
  - **Koneksi Backend**:
    - Endpoint `/register` di `C:\laragon\www\WBE\frontend\user\backend\handlers\auth_handler.go` menerima data
    - Backend melakukan validasi tambahan di `C:\laragon\www\WBE\frontend\user\backend\middleware\validator.go`
    - Password di-hash menggunakan bcrypt di `C:\laragon\www\WBE\frontend\user\backend\utils\password.go`
    - Data disimpan di Supabase melalui fungsi di `C:\laragon\www\WBE\frontend\user\backend\database\user_queries.go`
    - Setelah berhasil, backend membuat JWT Token di `C:\laragon\www\WBE\frontend\user\backend\utils\jwt.go`
    - Frontend menerima respons dan menampilkan notifikasi sukses/error (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\ui.js`)
    - Jika sukses, user diarahkan ke dashboard.html (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\navigation.js`)

- **Buat `login_user.html`**:
  - **Struktur Formulir Login di `C:\laragon\www\WBE\frontend\user\template\login.html`**
  - **File JavaScript yang Diperlukan**:
    - `C:\laragon\www\WBE\frontend\user\assets\js\validation.js`: Berisi fungsi validasi umum
    - `C:\laragon\www\WBE\frontend\user\assets\js\login.js`: Berisi logika khusus untuk halaman login
    - `C:\laragon\www\WBE\frontend\user\assets\js\api.js`: Berisi fungsi untuk memanggil API
    - `C:\laragon\www\WBE\frontend\user\assets\js\ui.js`: Berisi fungsi UI seperti loading spinner dan notifikasi
  - **Formulir Login**:
    - Username/Email: Satu kolom untuk memasukkan username atau email
    - Kata Sandi: Input bertipe password dengan tombol toggle show/hidden
    - Checkbox "Remember Me" untuk menyimpan sesi login
    - Tombol Login untuk submit form
    - Tautan "Belum Punya Akun? Daftar Sekarang" ke halaman registrasi
    - Tautan "Lupa Kata Sandi" untuk reset password
  - **Validasi Frontend**:
    - Validasi JavaScript untuk memastikan input tidak kosong (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\login.js`)
    - Pesan error yang jelas jika validasi gagal (styling di `C:\laragon\www\WBE\frontend\user\assets\css\forms.css`)
  - **Pengiriman Data ke Backend**:
    - Data dikirim ke backend melalui HTTP POST request ke endpoint login (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\api.js`)
    - Menampilkan loading spinner saat proses login berlangsung (implementasi di `C:\laragon\www\WBE\frontend\user\assets\js\ui.js`)
  - **Koneksi Backend**:
    - Endpoint `/login` di `C:\laragon\www\WBE\frontend\user\backend\handlers\auth_handler.go` menerima data
    - Backend mencari user di Supabase melalui fungsi di `C:\laragon\www\WBE\frontend\user\backend\database\user_queries.go`
    - Memverifikasi password menggunakan fungsi di `C:\laragon\www\WBE\frontend\user\backend\utils\password.go`
    - Jika berhasil, backend membuat JWT Token di `C:\laragon\www\WBE\frontend\user\backend\utils\jwt.go`
    - Jika "Remember Me" dicentang, token disimpan dalam cookie HTTP-only di `C:\laragon\www\WBE\frontend\user\backend\utils\cookie.go`
    - Frontend menerima respons dan mengalihkan ke dashboard menggunakan `C:\laragon\www\WBE\frontend\user\assets\js\navigation.js`
    - Jika gagal, menampilkan pesan error yang sesuai menggunakan `C:\laragon\www\WBE\frontend\user\assets\js\ui.js`
    - Backend menerapkan pembatasan percobaan login gagal di `C:\laragon\www\WBE\frontend\user\backend\middleware\rate_limiter.go`

## 4. Pengembangan Backend
- **Inisialisasi Proyek Go**: 
  - Setup project Go di `C:\laragon\www\WBE\frontend\user\backend\`
  - File utama: `C:\laragon\www\WBE\main.go` - Entry point aplikasi
  - Konfigurasi aplikasi: `C:\laragon\www\WBE\frontend\user\backend\config\config.go` - Load konfigurasi dari .env
  
- **Struktur Database**:
  - `C:\laragon\www\WBE\frontend\user\backend\database\migrations\`
    - `001_create_users_table.go` - Membuat tabel users
    - `002_create_transactions_table.go` - Membuat tabel transactions
    - `003_create_games_table.go` - Membuat tabel games
    - `004_create_game_providers_table.go` - Membuat tabel game_providers
    - `005_create_bets_table.go` - Membuat tabel bets
    - `006_create_bank_accounts_table.go` - Membuat tabel bank_accounts
    - `007_create_e_wallets_table.go` - Membuat tabel e_wallets
    - `008_create_admins_table.go` - Membuat tabel admins untuk menyimpan data admin dan superadmin
      - Struktur tabel:
        ```sql
        CREATE TABLE admins (
          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
          username VARCHAR(50) NOT NULL UNIQUE,
          email VARCHAR(100) NOT NULL UNIQUE,
          password VARCHAR(255) NOT NULL,
          role VARCHAR(20) NOT NULL DEFAULT 'admin',
          is_active BOOLEAN DEFAULT true,
          last_login TIMESTAMP,
          created_at TIMESTAMP DEFAULT NOW(),
          updated_at TIMESTAMP DEFAULT NOW()
        );
        ```
  
- **Database Connection**:
  - `C:\laragon\www\WBE\frontend\user\backend\database\db.go`
    - Implementasi koneksi ke Supabase PostgreSQL
    - Menggunakan environment variables dari file .env
    - Menyediakan pool koneksi database
  
- **Query dan Repository**:
  - `C:\laragon\www\WBE\frontend\user\backend\database\user_queries.go`
    - `FindUserByUsername(username string)` - Mencari user berdasarkan username
    - `FindUserByEmail(email string)` - Mencari user berdasarkan email
    - `CreateUser(user User)` - Membuat user baru
    - `UpdateUser(user User)` - Memperbarui data user
    - `DeleteUser(userID int)` - Menghapus user
  
  - `C:\laragon\www\WBE\frontend\user\backend\database\transaction_queries.go`
    - `CreateTransaction(transaction Transaction)` - Membuat transaksi baru
    - `GetTransactionsByUserID(userID int)` - Mendapatkan transaksi berdasarkan user ID
    - `GetTransactionByID(transactionID int)` - Mendapatkan transaksi berdasarkan ID
    - `UpdateTransactionStatus(transactionID int, status string)` - Memperbarui status transaksi
  
  - `C:\laragon\www\WBE\frontend\user\backend\database\game_queries.go`
    - `GetAllGames()` - Mendapatkan semua game
    - `GetGamesByProvider(providerID int)` - Mendapatkan game berdasarkan provider
    - `GetGameByID(gameID int)` - Mendapatkan game berdasarkan ID
    - `GetPopularGames(limit int)` - Mendapatkan game populer

- **Integrasi Supabase**:
  - `C:\laragon\www\WBE\frontend\user\backend\database\supabase.go`
    - Konfigurasi Supabase client
    - Fungsi untuk query data dari Supabase
    - Fungsi untuk memanipulasi data di Supabase
    - Implementasi caching dengan Redis

- **Models**:
  - `C:\laragon\www\WBE\frontend\user\backend\model\user_model.go`
    - Struktur data User
    - Validasi data User
    - JSON marshaling/unmarshaling
  
  - `C:\laragon\www\WBE\frontend\user\backend\model\transaction_model.go`
    - Struktur data Transaction
    - Validasi data Transaction
    - JSON marshaling/unmarshaling
  
  - `C:\laragon\www\WBE\frontend\user\backend\model\game_model.go`
    - Struktur data Game, GameProvider
    - Validasi data Game
    - JSON marshaling/unmarshaling
  
  - `C:\laragon\www\WBE\frontend\user\backend\model\bank_model.go`
    - Struktur data untuk metode pembayaran (Bank, E-Wallet, Pulsa)
    - Validasi data pembayaran
    - JSON marshaling/unmarshaling

- **Endpoint API**:
  - `C:\laragon\www\WBE\frontend\user\backend\router\router.go`
    - Setup router menggunakan Gin/Echo
    - Definisi semua routes (auth, user, games, transactions)
    - Penerapan middleware pada routes
  
  - `C:\laragon\www\WBE\frontend\user\backend\handlers\auth_handler.go`
    - Endpoint `/register` - Menerima dan memproses data registrasi
    - Endpoint `/login` - Menerima dan memproses data login
    - Endpoint `/logout` - Proses logout
    - Endpoint `/refresh-token` - Memperbarui token JWT
  
  - `C:\laragon\www\WBE\frontend\user\backend\handlers\user_handler.go`
    - Endpoint `/user/profile` - Mendapatkan data profil user
    - Endpoint `/user/update-profile` - Memperbarui data profil
    - Endpoint `/user/change-password` - Memproses perubahan password
  
  - `C:\laragon\www\WBE\frontend\user\backend\handlers\game_handler.go`
    - Endpoint `/games` - Mendapatkan daftar semua game
    - Endpoint `/games/:id` - Mendapatkan detail game
    - Endpoint `/games/provider/:provider` - Mendapatkan game berdasarkan provider
    - Endpoint `/games/popular` - Mendapatkan game populer
  
  - `C:\laragon\www\WBE\frontend\user\backend\handlers\transaction_handler.go`
    - Endpoint `/transactions/deposit` - Memproses deposit
    - Endpoint `/transactions/withdraw` - Memproses penarikan
    - Endpoint `/transactions/history` - Mendapatkan riwayat transaksi
    - Endpoint `/transactions/:id` - Mendapatkan detail transaksi

- **Middleware**:
  - `C:\laragon\www\WBE\frontend\user\backend\middleware\auth_middleware.go`
    - Validasi token JWT
    - Verifikasi user terautentikasi
    - Implementasi role-based access control
  
  - `C:\laragon\www\WBE\frontend\user\backend\middleware\validator.go`
    - Validasi input request dari pengguna
    - Sanitasi data input
    - Error handling untuk input tidak valid
  
  - `C:\laragon\www\WBE\frontend\user\backend\middleware\rate_limiter.go`
    - Pembatasan jumlah request dari satu IP
    - Proteksi terhadap serangan brute force
    - Penggunaan Redis untuk menyimpan counter request
  
  - `C:\laragon\www\WBE\frontend\user\backend\middleware\logger_middleware.go`
    - Log semua request dan response
    - Log error dan warning
    - Integrasi dengan sistem monitoring

- **Utils**:
  - `C:\laragon\www\WBE\frontend\user\backend\utils\password.go`
    - Fungsi `HashPassword(password string)` - Hashing password dengan bcrypt
    - Fungsi `VerifyPassword(hash, password string)` - Verifikasi password
  
  - `C:\laragon\www\WBE\frontend\user\backend\utils\jwt.go`
    - Fungsi `GenerateToken(userID int)` - Membuat JWT token
    - Fungsi `VerifyToken(tokenString string)` - Validasi JWT token
    - Fungsi `GenerateRefreshToken(userID int)` - Membuat refresh token
  
  - `C:\laragon\www\WBE\frontend\user\backend\utils\cookie.go`
    - Fungsi untuk mengelola HTTP-only cookies
    - Fungsi untuk menyimpan token dalam cookie
    - Fungsi untuk menghapus cookie saat logout
  
  - `C:\laragon\www\WBE\frontend\user\backend\utils\response.go`
    - Fungsi untuk format response API yang konsisten
    - Handling error response
    - Format success response

- **Autentikasi**:
  - JWT token implementation
  - Refresh token mechanism
  - Cookie-based authentication
  - Secure password handling dengan bcrypt

## 5. Pengembangan Admin Panel
- **Struktur Admin Panel**:
  - Setup direktori di `C:\laragon\www\WBE\frontend\admin\`
  - Gunakan Go templates untuk rendering HTML dari backend

- **Backend Admin**:
  - `C:\laragon\www\WBE\main.go`
    - Entry point untuk aplikasi admin
    - Setup router dan middleware
  
  - `C:\laragon\www\WBE\frontend\admin\backend\router\admin_router.go`
    - Definisi routes untuk admin panel
    - Implementasi middleware khusus admin

- **Alur Kerja Sistem Pengecekan Superadmin**:
  - **Saat Server Dimulai**:
    - Server akan menghubungkan ke database Supabase menggunakan koneksi yang telah dikonfigurasi dalam file `.env`.
    - Implementasi di `C:\laragon\www\WBE\main.go` memanggil fungsi `checkSuperadmin()` saat inisialisasi server.
    
  - **Pengecekan Tabel `admins`**:
    - Server akan melakukan pengecekan ke tabel `admins` untuk melihat apakah sudah ada entri dengan `role = 'superadmin'`.
    - Implementasi di `C:\laragon\www\WBE\frontend\admin\backend\database\admin_queries.go` dengan fungsi `CountSuperadmins()` yang mengembalikan jumlah superadmin.
    - Jika belum ada Superadmin, maka sistem akan otomatis mengarahkan ke halaman `register_superadmin.html` untuk mendaftarkan Superadmin pertama kali.
    - Middleware di `C:\laragon\www\WBE\frontend\admin\backend\middleware\superadmin_checker.go` menangani logika redirect ini.
    
  - **Halaman Pendaftaran Superadmin (Register_superadmin.html)**:
    - Halaman ini BERDIRI SENDIRI dan merupakan halaman yang muncul pertama kali saat sistem belum memiliki superadmin.
    - Terletak di `C:\laragon\www\WBE\frontend\admin\template\register_superadmin.html`
    - Halaman ini tidak menggunakan template base.html dari admin panel dan memiliki desain minimal.
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\register_superadmin.js` untuk validasi client-side.
    - Superadmin baru akan diminta untuk mengisi formulir pendaftaran dengan data berikut:
      - **Username**: Harus unik dan minimal 4 karakter.
      - **Email**: Harus menggunakan format email yang valid (`@gmail.com`).
      - **Password**: Harus di-hash dan dilengkapi dengan ikon untuk **show/hide password**.
      - **Konfirmasi Password**: Memastikan bahwa password yang dimasukkan sama.
      - **Tombol Register**: Untuk mengirimkan data pendaftaran ke backend.
    
  - **Koneksi Backend untuk Pendaftaran Superadmin**:
    - Form pendaftaran mengirim data ke `C:\laragon\www\WBE\frontend\admin\backend\handlers\admin_auth_handler.go` dengan fungsi `RegisterSuperadminHandler()`.
    - Handler menjalankan validasi tambahan menggunakan `C:\laragon\www\WBE\frontend\admin\backend\middleware\validator.go`.
    - Password di-hash menggunakan bcrypt di `C:\laragon\www\WBE\frontend\admin\backend\utils\password.go`.
    
  - **Penyimpanan Data Superadmin**:
    - Setelah Superadmin mendaftar dengan benar, data akan disimpan di tabel `admins` dalam database Supabase dengan `role = 'superadmin'`.
    - Fungsi `CreateAdmin(admin Admin)` di `C:\laragon\www\WBE\frontend\admin\backend\database\admin_queries.go` menangani penyimpanan data.
    - Query SQL yang digunakan:
      ```sql
      INSERT INTO admins (username, email, password, role) VALUES ($1, $2, $3, 'superadmin') RETURNING id
      ```
    
  - **Setelah Pendaftaran Sukses**:
    - Jika pendaftaran berhasil, Superadmin diarahkan ke halaman login (`login_admin.html`).
    - Menggunakan HTTP redirect di `C:\laragon\www\WBE\frontend\admin\backend\handlers\admin_auth_handler.go`.
    
  - **Server Berjalan Normal**:
    - Setelah Superadmin berhasil mendaftar, server akan melanjutkan berjalan normal.
    - Middleware `superadmin_checker.go` hanya akan aktif jika tidak ada superadmin yang terdeteksi.
    
  - **Halaman Login Admin (Login_admin.html)**:
    - Halaman ini BERDIRI SENDIRI dan terpisah dari admin panel.
    - Terletak di `C:\laragon\www\WBE\frontend\admin\template\login_admin.html`
    - Digunakan oleh superadmin dan admin lain untuk login ke sistem.
    - Tidak menggunakan template base.html dari admin panel untuk menjaga keamanan.
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\login_admin.js` untuk validasi dan interaksi.
    - Authenticasi ditangani oleh `C:\laragon\www\WBE\frontend\admin\backend\handlers\admin_auth_handler.go` dengan fungsi `LoginAdminHandler()`.
    - Token JWT disimpan sebagai HTTP-only cookie untuk keamanan.
    - Middleware di `C:\laragon\www\WBE\frontend\admin\backend\middleware\admin_auth_middleware.go` menangani verifikasi token dan role checking.
    - Setelah login berhasil, admin diarahkan ke dashboard admin panel sesuai dengan role mereka.

  - **Model Data Admin**:
    - Struktur data admin didefinisikan di `C:\laragon\www\WBE\frontend\admin\backend\model\admin_model.go`:
      ```go
      type Admin struct {
          ID        string    `json:"id"`
          Username  string    `json:"username" validate:"required,min=4,max=50"`
          Email     string    `json:"email" validate:"required,email"`
          Password  string    `json:"password,omitempty" validate:"required,min=8"`
          Role      string    `json:"role" validate:"required,oneof=admin superadmin"`
          IsActive  bool      `json:"is_active"`
          LastLogin time.Time `json:"last_login,omitempty"`
          CreatedAt time.Time `json:"created_at"`
          UpdatedAt time.Time `json:"updated_at"`
      }
      ```

  - **Log Akses Admin**:
    - Semua aktivitas login dan tindakan admin dicatat di `C:\laragon\www\WBE\frontend\admin\backend\middleware\admin_logger_middleware.go`.
    - Log disimpan di database untuk audit trail dan keamanan.
    - Fungsi `LogAdminActivity(adminID, action, details string)` di `C:\laragon\www\WBE\frontend\admin\backend\utils\logger.go` menangani pencatatan aktivitas.

  - **Pendaftaran Admin Role Lain oleh Superadmin**:
    - Setelah superadmin berhasil login, mereka memiliki akses ke halaman administrator.html untuk mengelola admin lain.
    - Terletak di `C:\laragon\www\WBE\frontend\admin\template\administrator.html`
    - Memperluas `C:\laragon\www\WBE\frontend\admin\template\base.html` karena sudah berada di dalam admin panel.
    - Halaman ini hanya bisa diakses oleh user dengan role 'superadmin'.
    - Berisi form pendaftaran admin baru dengan field:
      - Username (wajib)
      - Email (wajib)
      - Password (wajib)
      - Role (admin, dengan level akses terbatas)
      - Status (aktif/nonaktif)
    - Superadmin dapat melihat daftar semua admin, mengedit, atau menonaktifkan akun admin.
    - Fitur ini ditangani oleh `C:\laragon\www\WBE\frontend\admin\backend\handlers\administrator_handler.go`
    - Menggunakan HTMX untuk operasi CRUD pada admin tanpa reload halaman penuh.
    - Data admin biasa yang didaftarkan oleh superadmin disimpan dalam tabel `admins` yang sama dengan superadmin, hanya berbeda pada nilai kolom `role` yang diatur sebagai 'admin'.
    - Query SQL untuk pendaftaran admin biasa:
      ```sql
      INSERT INTO admins (username, email, password, role, is_active) VALUES ($1, $2, $3, 'admin', $4) RETURNING id
      ```
    - Admin yang telah terdaftar dapat login melalui halaman login_admin.html seperti superadmin, tetapi memiliki akses yang lebih terbatas sesuai dengan rolenya.

- **Halaman Admin Panel**:
  - `C:\laragon\www\WBE\frontend\admin\template\dashboard.html`
    - Halaman dashboard admin
    - Menampilkan statistik situs (jumlah user, transaksi, pendapatan)
    - Grafik untuk visualisasi data
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\dashboard.js`
    - Mengambil data dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\dashboard_handler.go`
  
  - `C:\laragon\www\WBE\frontend\admin\template\user_management.html`
    - Menampilkan daftar semua pengguna
    - Fitur search dan filter pengguna
    - Fungsi CRUD untuk data pengguna
    - Menggunakan HTMX untuk operasi CRUD tanpa reload halaman
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\user_management.js`
    - Mengambil data dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\user_management_handler.go`
  
  - `C:\laragon\www\WBE\frontend\admin\template\transaction_management.html`
    - Menampilkan semua transaksi (deposit/withdraw)
    - Filter berdasarkan status, jenis, dan tanggal
    - Persetujuan/penolakan withdraw
    - Konfirmasi deposit manual
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\transaction_management.js`
    - Mengambil data dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\transaction_management_handler.go`
  
  - `C:\laragon\www\WBE\frontend\admin\template\game_management.html`
    - Manajemen game dan provider
    - Pengaturan status aktif/nonaktif game
    - Pengaturan RTP (Return to Player) untuk slot games
    - Pengaturan maintenance mode
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\game_management.js`
    - Mengambil data dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\game_management_handler.go`
  
  - `C:\laragon\www\WBE\frontend\admin\template\settings.html`
    - Pengaturan umum situs
    - Manajemen bank dan metode pembayaran
    - Pengaturan promo dan bonus
    - Backup dan restore database
    - Menggunakan JavaScript dari `C:\laragon\www\WBE\frontend\admin\assets\js\settings.js`
    - Mengambil data dari `C:\laragon\www\WBE\frontend\admin\backend\handlers\settings_handler.go`

- **Handler Admin**:
  - `C:\laragon\www\WBE\frontend\admin\backend\handlers\dashboard_handler.go`
    - Fungsi untuk mengumpulkan statistik untuk dashboard
    - Proses data untuk grafik dan visualisasi
  
  - `C:\laragon\www\WBE\frontend\admin\backend\handlers\user_management_handler.go`
    - Fungsi CRUD untuk manajemen pengguna
    - Pencarian dan filter pengguna
    - Fungsi untuk reset password, ban/unban
  
  - `C:\laragon\www\WBE\frontend\admin\backend\handlers\transaction_management_handler.go`
    - Fungsi untuk melihat dan mengelola transaksi
    - Approval/rejection untuk transaksi withdraw
    - Konfirmasi deposit manual
  
  - `C:\laragon\www\WBE\frontend\admin\backend\handlers\game_management_handler.go`
    - Fungsi untuk manajemen game dan provider
    - Toggle status aktif/nonaktif
    - Pengaturan RTP
  
  - `C:\laragon\www\WBE\frontend\admin\backend\handlers\settings_handler.go`
    - Fungsi untuk pengaturan umum situs
    - Manajemen metode pembayaran
    - Manajemen promo dan bonus

- **Assets Admin**:
  - `C:\laragon\www\WBE\frontend\admin\assets\js\admin.js`
    - JavaScript umum untuk semua halaman admin
    - Fungsi utilitas dan helpers
  
  - `C:\laragon\www\WBE\frontend\admin\assets\js\dashboard.js`
    - Implementasi grafik dan chart menggunakan Chart.js
    - Real-time data update menggunakan WebSockets
  
  - `C:\laragon\www\WBE\frontend\admin\assets\js\user_management.js`
    - Implementasi CRUD operations dengan HTMX
    - Filter dan pencarian user
  
  - `C:\laragon\www\WBE\frontend\admin\assets\js\transaction_management.js`
    - Filter transaksi
    - Konfirmasi approval/reject
  
  - `C:\laragon\www\WBE\frontend\admin\assets\js\game_management.js`
    - Toggle game status
    - Set RTP values
    - Management provider games
  
  - `C:\laragon\www\WBE\frontend\admin\assets\css\admin.css`
    - Styling khusus untuk admin panel
    - Dark/light theme
    - Responsive layout

## 6. Pengujian dan Deployment
- **Pengujian**: Lakukan pengujian unit dan integrasi untuk memastikan semua fitur berfungsi dengan baik.
- **Deployment**: Deploy aplikasi ke server atau platform cloud, dan pastikan semua konfigurasi sudah benar.

## 7. Pemeliharaan dan Pembaruan
- **Pembaruan Reguler**: Jaga sistem tetap diperbarui dengan fitur baru dan patch keamanan.
- **Umpan Balik Pengguna**: Terus kumpulkan umpan balik pengguna untuk meningkatkan platform.
