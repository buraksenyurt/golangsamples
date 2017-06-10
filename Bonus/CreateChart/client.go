package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
)

const (
	Width  = 640
	Height = 480
)

func main() {
	var axis []string
	axis = append(axis, "OEM", "core i7 Cpu", "keyboard", "mouse", "1080P Monitor")
	var values []float64
	values = append(values, 40, 7, 12, 12, 45)
	// PieChart üzerinde sunacağımız veriyi tutan yapımız
	// Chart başlığını, x ve y değerlerini taşıyor
	data := ChartData{Title: "Stock Values", Axis: axis, Values: values}
	c := CreatePieChart("Our Stock Report", data)
	SavePieToFile("stocks.png", c)
}

// PieChart'ı dosyaya kaydeden fonksiyon
func SavePieToFile(fileName string, c chart.PieChart) {
	imgFile, _ := os.Create(fileName) //Sistemde bir dosya oluştur
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(img, img.Bounds(), bg, image.ZP, draw.Src)
	igr := imgg.AddTo(img, 0, 0, Width, Height, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)              //char ile imajı burada eşleştirdik
	png.Encode(imgFile, img) // img zaten igr ilişkilendirilmişti. Burada fiziki dosyaya kaydedilmiş oluyor
	imgFile.Close()
}

//Pie Chart oluşturan fonksiyonumuz
func CreatePieChart(title string, data ChartData) chart.PieChart {
	pie := chart.PieChart{Title: title}                 //Başlığını verdik
	pie.AddDataPair(data.Title, data.Axis, data.Values) //Verileri aktardık
	pie.Inner = 0
	pie.Key.Border = -1
	pie.FmtVal = chart.AbsoluteValue
	return pie
}

type ChartData struct {
	Title  string
	Axis   []string
	Values []float64
}
