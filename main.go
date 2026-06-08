package main

import (
	"fmt"
	"time"
)

type Mahasiswa struct {
	Nama         string
	NIM          string
	Totalbayar   int
	StatusLunas  bool
	TanggalBayar string
}

var datamahasiswa []Mahasiswa
var totalpengeluaran int

const kasperorang = 100000

// dibuat Gede Yogi Yogesvara Dita Diasta (109082500037)
func tampilkandaftarkas() {
	var m Mahasiswa
	var totalsaldo int

	fmt.Println("Daftar Kas : ")
	fmt.Printf("NIM\t\tNama\t\tTotal Bayar\t\tTanggal\t\tStatus\t\tSisa Tunggakan\n")
	for i := 0; i < len(datamahasiswa); i++ {
		m = datamahasiswa[i]
		totalsaldo = totalsaldo + m.Totalbayar
		fmt.Printf("%s\t\t%s\t\t%d\t\t\t%s\t%s\t%d\n", m.NIM, m.Nama, m.Totalbayar, m.TanggalBayar, statuslunas(m.StatusLunas), kasperorang-m.Totalbayar)
	}
	fmt.Printf("Total Saldo Kas: Rp %d\n", totalsaldo-totalpengeluaran)
	fmt.Printf("Jumlah Mahasiswa Lunas: %d dari %d Mahasiswa\n", hitunglunas(), len(datamahasiswa))
}

func hitunglunas() int {
	var count int = 0

	for i := 0; i < len(datamahasiswa); i++ {
		if datamahasiswa[i].StatusLunas {
			count++
		}
	}
	return count
}

func statuslunas(lunas bool) string {
	if lunas {
		return "Lunas"
	} else {
		return "Belum Lunas"
	}
}

// dibuat Rizky Iffat Venardi  (2311102301)
func pembayarankas() {
	var input string
	fmt.Print("Masukkan NIM atau Nama Mahasiswa: ")
	fmt.Scanln(&input)
	idx := -1
	i := 0
	for i < len(datamahasiswa) {
		if datamahasiswa[i].NIM == input || datamahasiswa[i].Nama == input {
			idx = i
		}
		i++
	}
	if idx != -1 {
		fmt.Println("\nData ditemukan!")
		fmt.Println("NIM  :", datamahasiswa[idx].NIM)
		fmt.Println("Nama :", datamahasiswa[idx].Nama)
		var nominal int
		fmt.Print("Masukkan nominal pembayaran: ")
		fmt.Scanln(&nominal)
		tgl := time.Now().Format("02-01-2006")
		datamahasiswa[idx].TanggalBayar = tgl
		datamahasiswa[idx].Totalbayar += nominal
		if datamahasiswa[idx].Totalbayar >= kasperorang {
			datamahasiswa[idx].StatusLunas = true
		}
		fmt.Println("\nPEMBAYARAN BERHASIL")
		fmt.Println("Tanggal Bayar :", datamahasiswa[idx].TanggalBayar)
		fmt.Println("Total Bayar   :", datamahasiswa[idx].Totalbayar)
		if datamahasiswa[idx].StatusLunas {
			fmt.Println("Status : LUNAS")
		} else {
			sisa := kasperorang - datamahasiswa[idx].Totalbayar
			fmt.Println("Status : BELUM LUNAS")
			fmt.Println("Sisa Tunggakan :", sisa)
		}
	} else {
		fmt.Println("\nData mahasiswa tidak ditemukan!")
	}
}

// dibuat Ignatius Steven M. (109082500089)
func pengeluaranKelas() {
	var jumlah int
	var keterangan string

	fmt.Println("=== Pengeluaran Kelas ===")

	fmt.Print("Masukkan jumlah pengeluaran: ")
	fmt.Scanln(&jumlah)

	fmt.Print("Masukkan keterangan: ")
	fmt.Scanln(&keterangan)

	totalKas := 0
	for i := 0; i < len(datamahasiswa); i++ {
		totalKas += datamahasiswa[i].Totalbayar
	}

	saldo := totalKas - totalpengeluaran

	if jumlah > saldo {
		fmt.Println("Pengeluaran gagal, saldo kas tidak mencukupi")
		return
	}

	totalpengeluaran += jumlah

	fmt.Println()
	fmt.Println("PENGELUARAN BERHASIL")
	fmt.Println("Jumlah Pengeluaran :", jumlah)
	fmt.Println("Keterangan         :", keterangan)
	fmt.Println("Sisa Saldo Kas     :", totalKas-totalpengeluaran)
}

// dibuat Gede Yogi Yogesvara Dita Diasta (109082500037)
func caridata() {
	var input string

	fmt.Print("Masukkan NIM atau Nama : ")
	fmt.Scanln(&input)
	idx := -1
	ketemu := false

	for i := 0; i < len(datamahasiswa) && !ketemu; i++ {
		if datamahasiswa[i].NIM == input || datamahasiswa[i].Nama == input {
			idx = i
			ketemu = true
		}
	}

	if idx != -1 {
		fmt.Println("\nData Mahasiswa!")
		fmt.Println("NIM  : ", datamahasiswa[idx].NIM)
		fmt.Println("Nama : ", datamahasiswa[idx].Nama)
		fmt.Println("Total Bayar : Rp", datamahasiswa[idx].Totalbayar)
		fmt.Println("Status      : ", statuslunas(datamahasiswa[idx].StatusLunas))
		fmt.Println("Sisa Tunggakan : Rp ", kasperorang-datamahasiswa[idx].Totalbayar)
	} else {
		fmt.Println("\nData mahasiswa tidak ditemukan!")
	}
}

// dibuat Ignatius Steven M. (109082500089)
func urutkanPenunggakKas() {
	var temp Mahasiswa

	for i := 0; i < len(datamahasiswa)-1; i++ {
		for j := 0; j < len(datamahasiswa)-i-1; j++ {
			if datamahasiswa[j].Totalbayar > datamahasiswa[j+1].Totalbayar {
				temp = datamahasiswa[j]
				datamahasiswa[j] = datamahasiswa[j+1]
				datamahasiswa[j+1] = temp
			}
		}
	}

	fmt.Println("Daftar Mahasiswa (Diurutkan dari pembayaran terkecil)")
	fmt.Println("_____________________________________________________")

	for i := 0; i < len(datamahasiswa); i++ {
		fmt.Printf("%d. %s (%s) - Total Bayar: Rp %d\n",
			i+1,
			datamahasiswa[i].Nama,
			datamahasiswa[i].NIM,
			datamahasiswa[i].Totalbayar)
	}
}

// dibuat Ignatius Steven M. (109082500089)
func TambahMahasiswa() {
	var nim, nama string

	fmt.Println("=== Tambah Mahasiswa ===")

	fmt.Println("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Println("Masukkan Nama: ")
	fmt.Scanln(&nama)

	mhsBaru := Mahasiswa{
		NIM:         nim,
		Nama:        nama,
		Totalbayar:  0,
		StatusLunas: false,
	}

	datamahasiswa = append(datamahasiswa, mhsBaru)

	fmt.Println("Data mahasiswa berhasil ditambahkan")
}

// dibuat Rizky Iffat Venardi  (2311102301)
func editData() {
	if len(datamahasiswa) == 0 {
		fmt.Println("\n✗ Belum ada Data Mahasiswa!")
		return
	}
	var inputNIM string
	fmt.Print("\nMasukkan NIM Mahasiswa yang akan diedit: ")
	fmt.Scanln(&inputNIM)
	idx := -1
	for i := 0; i < len(datamahasiswa); i++ {
		if datamahasiswa[i].NIM == inputNIM {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("\n✗ Mahasiswa dengan NIM tidak ditemukan!")
		return
	}
	fmt.Println("\nDATA MAHASISWA")
	fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
	fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
	fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
	fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
	if !datamahasiswa[idx].StatusLunas {
		fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
	}
	fmt.Println("\nPILIH DATA YANG AKAN DIEDIT")
	fmt.Println("1. Edit NIM")
	fmt.Println("2. Edit Nama")
	fmt.Println("3. Edit Total Bayar")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scanln(&pilihan)
	if pilihan == 1 {
		var NIMbaru string
		fmt.Print("\nMasukkan NIM baru: ")
		fmt.Scanln(&NIMbaru)
		datamahasiswa[idx].NIM = NIMbaru
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ NIM berhasil diubah!")
	} else if pilihan == 2 {
		var namabaru string
		fmt.Print("\nMasukkan Nama Baru: ")
		fmt.Scanln(&namabaru)
		datamahasiswa[idx].Nama = namabaru
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ Nama berhasil diubah!")
	} else if pilihan == 3 {
		fmt.Printf("\nTotal Bayar Saat Ini: Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Sisa Tunggakan Saat Ini: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		var totalbaru int
		fmt.Print("Masukkan Total Bayar baru: Rp ")
		fmt.Scanln(&totalbaru)
		if totalbaru < 0 {
			fmt.Println("\n✗ Total bayar tidak boleh negatif!")
			return
		}
		datamahasiswa[idx].Totalbayar = totalbaru
		if datamahasiswa[idx].Totalbayar >= kasperorang {
			datamahasiswa[idx].StatusLunas = true
		} else {
			datamahasiswa[idx].StatusLunas = false
		}
		fmt.Println("\nDATA SETELAH DIEDIT")
		fmt.Printf("NIM         : %s\n", datamahasiswa[idx].NIM)
		fmt.Printf("Nama        : %s\n", datamahasiswa[idx].Nama)
		fmt.Printf("Total Bayar : Rp %d\n", datamahasiswa[idx].Totalbayar)
		fmt.Printf("Status      : %s\n", statuslunas(datamahasiswa[idx].StatusLunas))
		if !datamahasiswa[idx].StatusLunas {
			fmt.Printf("Sisa Tunggakan: Rp %d\n", kasperorang-datamahasiswa[idx].Totalbayar)
		}
		fmt.Println("\n✓ Total bayar berhasil diubah!")
	} else {
		fmt.Println("\n✗ Pilihan tidak valid! Silakan pilih 1, 2, atau 3.")
		return
	}
}

// dibuat Rizky Iffat Venardi  (2311102301)
func HapusData() {
	if len(datamahasiswa) == 0 {
		fmt.Println("\n✗ Belum ada data Mahasiswa!")
		return
	}
	var nim string
	fmt.Print("\nMasukkan NIM Mahasiswa yang akan dihapus: ")
	fmt.Scanln(&nim)
	idx := -1
	ditemukan := false
	for i := 0; i < len(datamahasiswa) && !ditemukan; i++ {
		if datamahasiswa[i].NIM == nim {
			idx = i
			ditemukan = true
		}
	}
	if idx == -1 {
		fmt.Println("✗ Mahasiswa dengan NIM tersebut tidak ditemukan!")
		return
	}
	fmt.Println("\n=== Detail Data yang Akan Dihapus ===")
	fmt.Printf("NIM        : %s\n", datamahasiswa[idx].NIM)
	fmt.Printf("Nama       : %s\n", datamahasiswa[idx].Nama)
	fmt.Printf("Total Bayar: Rp %d\n", datamahasiswa[idx].Totalbayar)
	fmt.Print("\nYakin ingin menghapus data ini? (yes/no): ")
	var konfirmasi string
	fmt.Scanln(&konfirmasi)
	if konfirmasi == "yes" || konfirmasi == "YES" || konfirmasi == "y" || konfirmasi == "Y" {
		for i := idx; i < len(datamahasiswa)-1; i++ {
			datamahasiswa[i] = datamahasiswa[i+1]
		}
		datamahasiswa = datamahasiswa[:len(datamahasiswa)-1]
		fmt.Println("\n✓ Data berhasil dihapus!")
	} else {
		fmt.Println("\n✗ Penghapusan data dibatalkan.")
	}
}

// dibuat Gede Yogi Yogesvara Dita Diasta (109082500037)
func main() {
	var pilih int

	for {
		fmt.Println("___________________________________________")
		fmt.Println("       SIKAS APLIKASI KAS MAHASISWA       ")
		fmt.Println("___________________________________________")
		fmt.Println("1. Tampilkan Daftar Kas & Status")
		fmt.Println("2. Input Pembayaran Kas")
		fmt.Println("3. Input Pengeluaran Kelas")
		fmt.Println("4. Cari Data Mahasiswa")
		fmt.Println("5. Urutkan Penunggak Kas")
		fmt.Println("6. Tambah Data Baru")
		fmt.Println("7. Edit Data")
		fmt.Println("8. Menghapus Data")
		fmt.Println("9. Keluar")
		fmt.Println("___________________________________________")
		fmt.Print("Pilih menu (1-9): ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			fmt.Println("Anda memilih 1 : Tampilkan Daftar Kas & Status")
			fmt.Println()
			tampilkandaftarkas()
		case 2:
			fmt.Println("Anda memilih 2 : Input Pembayaran Kas")
			fmt.Println()
			pembayarankas()
		case 3:
			fmt.Println("Anda memilih 3 : Input Pengeluaran Kelas")
			fmt.Println()
			pengeluaranKelas()
		case 4:
			fmt.Println("Anda memilih 4 : Cari Data Mahasiswa")
			fmt.Println()
			caridata()
		case 5:
			fmt.Println("Anda memilih 5 : Urutkan Penunggak Kas")
			fmt.Println()
			urutkanPenunggakKas()
		case 6:
			fmt.Println("Anda memilih 6 : Tambah Data Baru")
			fmt.Println()
			TambahMahasiswa()
		case 7:
			fmt.Println("Anda memilih 7 : Edit Data")
			fmt.Println()
			editData()
		case 8:
			fmt.Println("Anda memilih 8 : Menghapus Data")
			fmt.Println()
			HapusData()
		case 9:
			fmt.Println("Anda memilih 9 : Keluar")
			fmt.Println("Terima kasih telah menggunakan aplikasi SIKAS")
			return
		}
	}
}
