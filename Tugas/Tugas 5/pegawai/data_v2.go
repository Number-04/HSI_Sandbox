package pegawai

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dustin/go-humanize"
)

type ListPegawai struct {
	List []Pegawai
}

type InformasiPegawaiV2 interface {
	TampilkanInformasiV2()
}

func (l ListPegawai) TampilkanInformasiV2() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "|\t ---- \t ------ \t ------------ \t ------------ \t|")
	fmt.Fprintln(w, "|\t Nama \t Posisi \t Gaji Bulanan \t Gaji Tahunan \t|")
	fmt.Fprintln(w, "|\t ---- \t ------ \t ------------ \t ------------ \t|")

	for _, p := range l.List {
		fmt.Fprintf(
			w,
			"|\t %s \t %s \t %s \t %s \t|\n",
			p.Nama,
			p.Posisi,
			FormatRupiah(p.GajiBulanan),
			FormatRupiah(p.HitungGajiTahunan()),
		)
	}

	w.Flush()
}

func FormatRupiah(n float64) string {
	return "Rp " + humanize.Comma(int64(n))
}
