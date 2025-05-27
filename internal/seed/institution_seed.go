package seed

import (
	"log"
	"tindak_ai/internal/domain"
	"tindak_ai/pkg/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedInstitution(db *gorm.DB) {
	institution := []domain.Institution{

		{
			Name:         "SEKRETARIAT DAERAH",
			Address:      "JI. Kyai Singkil No 7",
			PostalCode:   "59511",
			ContactPhone: "(0291) 685322",
			ContactEmail: "setda@demakkab.go.id",
			Website:      helper.PtrString("http://setda.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291) 685625"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude:     -7.035695,
			Longitude:    110.422796,
		},
		{
			Name:         "SEKRETARIAT DPRD",
			Address:      "Jl. Sultan Trenggono No",
			PostalCode:   "59516",
			ContactPhone: "(0291) 681177",
			ContactEmail: "dprd@demakkab.go.id",
			Website:      helper.PtrString("http://dprd.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9102185,
			Longitude: 110.057836,
		},
		{
			Name:         "INSPEKTORAT",
			Address:      "JI.Kyai Mugni Nomor 1016",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 685972",
			ContactEmail: "inspektorat@demakkab.go.id",
			Website:      helper.PtrString("http://inspektorat.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9065837,
			Longitude: 110.0578123,
		},
		{
			Name:         "Bappelitbangda (BADAN PERENCANAAN PEMBANGUNAN, PENELITIAN DAN PENGEMBANGAN DAERAH)",
			Address:      "JL.Kyai Jebat No.30 BINTORO DEMAK Gedung hijau Lt.3",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 685663",
			ContactEmail: "bappedalitbang@demakkab.go.id",
			Website:      helper.PtrString("http://bappelitbangda.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291) 685632"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8912276,
			Longitude: 110.6352443,
		},
		{
			Name:         "BKPP (Badan Kepegawaian, Pendidikan Dan Pelatihan)",
			Address:      "JL.Raya Buyaran Demak No.65A Karangtengah Demak",
			PostalCode:   "", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 681643,690155,6904466",
			ContactEmail: "bkpp@demakkab.go.id",
			Website:      helper.PtrString("https://bkpp.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291)681643,690155"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9189216,
			Longitude: 110.592623,
		},
		{
			Name:         "BPKPAD (Badan Pengelolaan Keuangan, Pendapatan dan Aset Daerah)",
			Address:      "JI. Kyai Jebat No 881A",
			PostalCode:   "59511",
			ContactPhone: "(0291) 685560",
			ContactEmail: "bpkpad@demakkab.go.id",
			Website:      helper.PtrString("http://bpkpad.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8916951,
			Longitude: 110.6349157,
		},
		{
			Name:         "BPBD (Badan Penanggulanagan Bencana Daerah)",
			Address:      "JI. Bayangkara baru No",
			PostalCode:   "59515", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 682 200",
			ContactEmail: "bpbd@demakkab.go.id",
			Website:      helper.PtrString("http://bpbd.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.894811,
			Longitude: 110.6317553,
		},
		{
			Name:         "DINPEPUSAR(Dinas Perpustakaan Dan Kearsipan)",
			Address:      "Jl. Sultan Fatah No.67",
			PostalCode:   "59515", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 681075",
			ContactEmail: "dpk.demak@gmail.com",
			Website:      helper.PtrString("http://dinperpusar.demakkab.go.id"),
			Fax:          helper.PtrString("(0291) 681075"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8976443,
			Longitude: 110.6308102,
		},
		{
			Name:         "DINAS PERTANIAN DAN PANGAN",
			Address:      "JL.Sultan Hadiwijaya No.8",
			PostalCode:   "59515", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291)685013",
			ContactEmail: "dinpertanpangan@demakkab.go.id",
			Website:      helper.PtrString("http://dinpertanpangan.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291)685013"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.2473322,
			Longitude: 109.6953146,
		},
		{
			Name:         "DINPERKIM (Dinas Perumahan Dan Pemukiman)",
			Address:      "JL. Kyai Jebat No.35 LT.3",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291)685715",
			ContactEmail: "dinperkim.kabdemak@gmail.com",
			Website:      helper.PtrString("http://dinperkim.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291)685395"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8909744,
			Longitude: 105.759083,
		},
		{
			Name:         "DINPUTARU (Dinas PU dan Tata ruang)",
			Address:      "JL. Kyai Jebat No.35",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "0291 685123",
			ContactEmail: "dinputaru@demakkab.go.id",
			Website:      helper.PtrString("http://dinputaru.demakkab.go.id"),
			Fax:          helper.PtrString("0291 6905623"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8909744,
			Longitude: 105.759083,
		},
		{
			Name:         "DINKES (Dinas Kesehatan)",
			Address:      "JL. SULTAN HADIWIJAYA 44 DEMAK",
			PostalCode:   "59515",
			ContactPhone: "(0291) 685934", // Kode pos 59515 dari PDF ada di baris telepon
			ContactEmail: "dinkes.demakkab@gmail.com",
			Website:      helper.PtrString("http://dinkes.demakkab.go.id"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8906903,
			Longitude: 110.6304935,
		},
		{
			Name:         "DINAS PARIWISATA",
			Address:      "JL. SULTAN FATAH, NO. 53, DEMAK",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "", // Tidak ada nomor telepon di data asli
			ContactEmail: "Dinparta@demakkab.g.id",
			Website:      helper.PtrString("http://pariwisata.demakkab.go.id"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8960352,
			Longitude: 110.6376115,
		},
		{
			Name:         "DINAS LINGKUNGAN HIDUP",
			Address:      "JI. Bayangkara baru No 1",
			PostalCode:   "59515",
			ContactPhone: "(0291) 685677",
			ContactEmail: "dinlh@demakkab.go.id",
			Website:      helper.PtrString("http://dinlh.demakkab.go.id"),
			Fax:          helper.PtrString("(0291) 681911"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8968476,
			Longitude: 110.634439,
		},
		{
			Name:         "DINDUKCAPIL (Dinas Kependudukan Dan Pencatatan Sipil)",
			Address:      "JI.Kyai Mugni Nomor 1016",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 685972",
			ContactEmail: "dukcapil@demakkab.go.id",
			Website:      helper.PtrString("http://dukcapil.demakkab.go.id"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8906367,
			Longitude: 108.1993749,
		},
		{
			Name:         "DINAS PERHUBUNGAN",
			Address:      "Jl. Sultan Trenggono No. 16",
			PostalCode:   "59515",
			ContactPhone: "0291-685863",
			ContactEmail: "perhubungan@demakkab.go.id",
			Website:      helper.PtrString("http://dinhub.demakkab.go.id"), // 'g' diganti 'go' asumsi typo
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8906367,
			Longitude: 108.1993749,
		},
		{
			Name:         "DINAS PERDAGANGAN DAN KOPERASI", // Nama diperbaiki dari "DINAS PERDAGANGAN DAN KOPERASI" ke "DINAS PERDAGANGAN, KOPERASI, USAHA KECIL DAN MENENGAH" jika ini lebih akurat, namun saya akan tetap sesuai PDF "DINAS PERDAGANGAN DAN KOPERASI"
			Address:      "JL. Kyai Mugni 1016, Demak",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291681604)",
			ContactEmail: "info@dindagkop.demakkab.go.id",
			Website:      helper.PtrString("http://dindagkopukm.demakkab.go.id/"), // Website merujuk dindagkopukm
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8906367,
			Longitude: 108.1993749,
		},
		{
			Name:         "DPMPTSP (Dinas Penanaman Modal, Perizinan Terpadu Satu Pintu)",
			Address:      "Jl. Sultan Hadi Wijaya No 8",
			PostalCode:   "59515",
			ContactPhone: "(0291) 681011",
			ContactEmail: "dinpmptsp@demakkab.go.id",
			Website:      helper.PtrString("http://perizinan.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8906367,
			Longitude: 108.1993749,
		},
		{
			Name:         "DINKOMINFO(Dinas Komunikasi dan Informatika)",
			Address:      "JL.Sultan Hadiwijaya No.4 DEMAK",
			PostalCode:   "59515",
			ContactPhone: "(0291)685790",
			ContactEmail: "dinkominfo@demakab.go.id", // Asumsi 'demakab' seharusnya 'demakkab'
			Website:      helper.PtrString("https://dinkominfo.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9009401,
			Longitude: 110.6287881,
		},
		{
			Name:         "DINAS PENDIDIKAN DAN KEBUDAYAAN",
			Address:      "JL.Sitan Trenggono No.89 DEMAK", // Sitan -> Sultan? Sesuai PDF 'Sitan'
			PostalCode:   "59511",                               // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291)685242",
			ContactEmail: "dindikbud@demakkab.go.id",
			Website:      helper.PtrString("http://dindikbud.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291) 685364"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9029179,
			Longitude: 110.6275289,
		},
		{
			Name:         "DINAS SOSIAL PEMBERDAYYAN PEREMPUAN DAN PERLINDUNGAN ANAK",
			Address:      "JL.Kyai Singkil No.42 DEMAK",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291)665745",
			ContactEmail: "dinsosp2pa@gmail.com",
			Website:      helper.PtrString("http://dinsosp2pa.demakkab.go.id"),
			Fax:          helper.PtrString("(0291)665745"), // Telp/Fax sama
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,
		},
		{
			Name:         "DINAS KEPEMUDAAN DAN OLAHRAGA",
			Address:      "JL.Hadi Wijaya No.45 DEMAK",
			PostalCode:   "59515", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 6905626",
			ContactEmail: "dinpora@demakkab.go.id",
			Website:      helper.PtrString("http://dinpora.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,
		},
		{
			Name:         "DINPERMADESP2KB (Dinas Pemberdayaan Masyarakat Dan Desa, Pengendalian Penduduk dan KB)",
			Address:      "L.Kyai Jebat No.30 BINTORO DEMAK Gedung Hijau Lt.2", // L.Kyai -> Jl.Kyai?
			PostalCode:   "59511",
			ContactPhone: "(0291) 685792,685376",
			ContactEmail: "dinpermadesp2kb@demakkab.go.id",
			Website:      helper.PtrString("http://dinpermadesp2kb.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,
		},
		{
			Name:         "DINAS TENAGA KERJA DAN PERINDUSTRIAN",
			Address:      "JI. Kyai Mugni Nomor 1016",
			PostalCode:   "59515", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291) 681142",
			ContactEmail: "dinnakerind@demakkab.go.id",
			Website:      helper.PtrString("https://dinnakerind.demakkab.go.id/"),
			Fax:          helper.PtrString("(0291) 685262"),
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,
		},
		{
			Name:         "DINAS KELAUTAN DAN PERIKANAN",
			Address:      "Jl.sultan Hadi wijaya No53",
			PostalCode:   "59515",
			ContactPhone: "(0291) 685368",
			ContactEmail: "dinlutkan@demakkab.go.id",
			Website:      helper.PtrString("dinlutkan.demakkab.go.id"), // Tanpa http://
			Fax:          helper.PtrString("(0291) 685368"),            // Telp/Fax sama
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,
		},
		{
			Name:         "BAKESBANGPOL (BADAN KESATUAN BANGSA DAN POLITIK)",
			Address:      "JL.Kyai Jebat No.30 BINTORO DEMAK Gedung Hijau Lt.1",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "", // Tidak ada nomor telepon di data asli
			ContactEmail: "kesbangpolinmas@demakkab.go.id",
			Website:      helper.PtrString("http://kesbangpolinmas.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,

		},
		{
			Name:         "SATPOL PP",
			Address:      "JL.Kyai Jebat No.30 BINTORO DEMAK Gedung Hijau Lt.1",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291)685495",
			ContactEmail: "satpolppkabdemak@yahoo.com",
			Website:      helper.PtrString("http://satpolpp.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8904054,
			Longitude: 108.1941451,
		},
		{
			Name:         "SEKRETARIAT KOMISI PEMILIHAN UMUM",
			Address:      "JL.Kyai Turmudzi No.1 DEMAK",
			PostalCode:   "59511", // Data asli tidak mencantumkan kode pos eksplisit untuk entri ini
			ContactPhone: "(0291)681753",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://kpu.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8960239,
			Longitude: 110.6392812,
		},
		{
			Name:         "KECAMATAN DEMAK",
			Address:      "Jl. Sultan Fatah No.16",
			PostalCode:   "59511",
			ContactPhone: "0291 685001",
			ContactEmail: "kecdemakkota@gmail.com",
			Website:      helper.PtrString("http://kecdemak.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9008192,
			Longitude: 110.632233,
		},
		{
			Name:         "KECAMATAN KARANGTENGAH",
			Address:      "Jl. Raya Buyaran No.31",
			PostalCode:   "59561",
			ContactPhone: "0291 686385",
			ContactEmail: "kecamatankarangtengah06@gamail.com", // gamail.com -> gmail.com?
			Website:      helper.PtrString("http://keckarangtengah.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9200289,
			Longitude: 110.5959228,
		},
		{
			Name:         "KECAMATAN SAYUNG",
			Address:      "Jl. Raya saying Km. 10", // saying -> Sayung?
			PostalCode:   "59563",
			ContactPhone: "(024) 76450373",
			ContactEmail: "kec-sayung@gmail..com", // gmail..com -> gmail.com?
			Website:      helper.PtrString("http://kecksayung.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9420685,
			Longitude: 110.5082238,
		},
		{
			Name:         "KECAMATAN BONANG",
			Address:      "Jl. Raya Demak Moro Km 12",
			PostalCode:   "59552",
			ContactPhone: "0291 6908020",
			ContactEmail: "kec_Bonang@yahoo.com",
			Website:      helper.PtrString("http://kecbonang.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8400435,
			Longitude: 109.9598378,
		},
		{
			Name:         "KECAMATAN WEDUNG",
			Address:      "Jl. Raya Ngawen No 44A",
			PostalCode:   "59554",
			ContactPhone: "0291 6906099",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://kecwedung.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8400435,
			Longitude: 109.9598378,
		},
		{
			Name:         "KECAMATAN MIJEN",
			Address:      "Jl. Raya Mijen No 59",
			PostalCode:   "59583",
			ContactPhone: "0291 4256438",
			ContactEmail: "kecmijen59@gmail.com",
			Website:      helper.PtrString("http://kecmijen.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.076518,
			Longitude: 107.8700943,
		},
		{
			Name:         "KECAMATAN WONOSALAM",
			Address:      "Jl. Raya Demak-Purwodadi Km 5",
			PostalCode:   "59571",
			ContactPhone: "0291 685720",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://kecwonosalam.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.076518,
			Longitude: 107.8700943,
		},
		{
			Name:         "KECAMATAN DEMPET",
			Address:      "Jl. Raya Dempet No 25",
			PostalCode:   "59573",
			ContactPhone: "0291 685716",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://kecdempet.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.953412,
			Longitude: 110.3902627,
		},
		{
			Name:         "KECAMATAN GAJAH",
			Address:      "Jl. Raya Gajah No 45",
			PostalCode:   "59581",
			ContactPhone: "0291 685250",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://kecgajah.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.8709812,
			Longitude: 110.7315475,
		},
		{
			Name:         "KECAMATAN KARANGANYAR",
			Address:      "Jl. Raya Demak- Kudus Km 17",
			PostalCode:   "59552", // PDF menampilkan 59552, di baris lain 59582 tapi itu bagian dari alamat lain. Entri ini 59552
			ContactPhone: "0291 4101162",
			ContactEmail: "Ikeckaranganyardemak@gmail.com", // 'I' di awal email?
			Website:      helper.PtrString("http://keckaranganyar.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.2275241,
			Longitude: 110.8713911,
		},
		{
			Name:         "KECAMATAN KEBONAGUNG",
			Address:      "Jl. Raya Semarang-Purwodadi km.36",
			PostalCode:   "59563", // PDF menampilkan 59563, di baris lain 59572 tapi itu bagian dari alamat lain. Entri ini 59563
			ContactPhone: "0291 5135754",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://keckebonagung.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.0203635,
			Longitude: 110.7008597,
		},
		{
			Name:         "KECAMATAN GUNTUR",
			Address:      "Jl. Guntur Raya NO.",
			PostalCode:   "59565",
			ContactPhone: "0291 685445",
			ContactEmail: "kec", // Email tidak lengkap
			Website:      helper.PtrString("http://kecguntur.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -6.9787602,
			Longitude: 110.6149087,
		},
		{
			Name:         "KECAMATAN KARANGAWEN",
			Address:      "Jl. Raya Kr.Awen-Mranggen No.115",
			PostalCode:   "59566",
			ContactPhone: "024-70792667",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://keckarangawen.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.0442099,
			Longitude: 110.5753056,
		},
		{
			Name:         "KECAMATAN MRANGGEN",
			Address:      "Jl. Raya Mranggen No.172",
			PostalCode:   "59567",
			ContactPhone: "024-6723114",
			ContactEmail: "", // Tidak ada email di data asli
			Website:      helper.PtrString("http://kecmranggen.demakkab.go.id/"),
			Fax:          nil,
			LogoUrl: helper.PtrString("https://upload.wikimedia.org/wikipedia/commons/0/06/Lambang_Kabupaten_Demak.png"),
			Latitude: -7.0259765,
			Longitude: 110.4409107,
		},
	}

	for _, ins := range institution {
		if err := db.FirstOrCreate(&ins, domain.Institution{
			ID: uuid.New(),
			Name:         ins.Name,
			Address:      ins.Address,
			PostalCode:   ins.PostalCode,
			ContactPhone: ins.ContactPhone,
			ContactEmail: ins.ContactEmail,
			Website:      ins.Website,
			Fax:          ins.Fax,
			LogoUrl: ins.LogoUrl,
			Latitude:     ins.Latitude,
			Longitude:    ins.Longitude,
		}).Error; err != nil {
			log.Printf("failed to seed institution %s : %v", ins.Name, err)
		}
	}
}
