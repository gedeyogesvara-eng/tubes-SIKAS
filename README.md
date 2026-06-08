# 💰 SIKAS - Aplikasi Kas Mahasiswa
SIKAS adalah aplikasi berbasis CLI Command Line Interface yang ditulis dengan bahasa pemrograman Golang. Aplikasi ini dirancang untuk mempermudah pengelolaan uang kas di dalam lingkup kelas atau organisasi mahasiswa, mulai dari pencatatan pembayaran, manajemen pengeluaran, hingga tracking status tunggakan.

# ✨ Fitur Utama

Aplikasi ini dilengkapi dengan fitur-fitur kas yang interaktif:
1. Tampilkan Daftar Kas & Status: Menampilkan tabel seluruh mahasiswa, nominal bayar, tanggal bayar terbaru, status lunas, sisa tunggakan, total saldo kelas, dan ringkasan kuota kelulusan kas.
2. Input Pembayaran Kas: Melakukan setoran kas mahasiswa dengan nominal tertentu. Sistem otomatis mendeteksi status kelunasan berdasarkan nominal batas atas.
3. Input Pengeluaran Kelas: Mencatat nominal pengeluaran kelas disertai dengan keterangan kegunaannya. Sistem memiliki validasi agar pengeluaran tidak melebihi sisa saldo kas saat ini.
4. Cari Data Mahasiswa: Mempermudah pencarian profil tunggakan mahasiswa menggunakan kata kunci nama atau NIM.
5. Urutkan Penunggak Kas: Mengurutkan daftar mahasiswa berdasarkan kontribusi pembayaran terkecil ke terbesar (*ascending*) menggunakan algoritma *Bubble Sort*.
6. Tambah Data Baru: Menambahkan entitas mahasiswa baru ke dalam database internal.
7. Edit Data: Mengubah data NIM, Nama, atau total akumulasi pembayaran kas mahasiswa apabila terjadi kesalahan input.
8. Menghapus Data: Menghapus data mahasiswa dari sistem dilengkapi dengan dialog konfirmasi (`yes/no`).

# 🛠️ Spesifikasi & Teknologi
  1. Bahasa Pemrograman: Go (Golang) versi 25.1 .
  2. Paradigma: Structured Programming dengan manipulasi objek array dinamis Slice.
  3. Konstanta Batas Kas: Rp 100.000 per orang (kasperorang = 100000).

# 🚀 Cara Menjalankan Aplikasi
Pastikan Anda sudah menginstal Go compiler di komputer Anda sebelum mengikuti langkah-langkah di bawah ini.
# 1. Clone Repositori
git clone [https://github.com/username-kamu/sikas-aplikasi-kas.git](https://github.com/username-kamu/sikas-aplikasi-kas.git)
cd sikas-aplikasi-kas
# 2. Jalankan Program (Tanpa Compile)
Anda bisa langsung mengeksekusi file kode utama:
go run main.go

# 💻 Panduan Penggunaan Menu (Struktur Menu CLI)
Saat aplikasi dijalankan, Anda akan dihadapkan pada menu interaktif berikut:
___________________________________________
       SIKAS APLIKASI KAS MAHASISWA       
___________________________________________
1. Tampilkan Daftar Kas & Status
2. Input Pembayaran Kas
3. Input Pengeluaran Kelas
4. Cari Data Mahasiswa
5. Urutkan Penunggak Kas
6. Tambah Data Baru
7. Edit Data
8. Menghapus Data
9. Keluar
___________________________________________
Pilih menu (1-9): 
Ketik angka 1-9 lalu tekan Enter untuk mengeksekusi menu pilihan Anda.

# 👥 Tim Kontributor
Proyek ini dibangun secara kolaboratif oleh anggota tim berikut (berdasarkan modul fungsi yang dikembangkan):

* Gede Yogi Yogesvara Dita Diasta (109082500037)
  1. Pengembangan Menu Utama (`main`)
  2. Fungsi Penayangan Daftar Kas (`tampilkandaftarkas`)
  3. Fungsi Pencarian Profil (`caridata`)


* Rizky Iffat Venardi (2311102301)
  1. Fungsi Setoran & Logika Pembayaran (`pembayarankas`)
  2. Fungsi Pembaruan Data (`editData`)
  3. Fungsi Penghapusan Entitas (`HapusData`)


* Ignatius Steven M. (109082500089)
  1. Fungsi Manajemen Finansial Keluar (`pengeluaranKelas`)
  2. Fungsi Algoritma Pengurutan Penunggak (`urutkanPenunggakKas`)
  3. Fungsi Pendaftaran Anggota Baru (`TambahMahasiswa`)
