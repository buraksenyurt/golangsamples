/*
	Lesson_20
	Temel Dosya işlemlerine bir bakalım
	Tüm fonksiyonlar için https://golang.org/pkg/os/ adresine gidebiliriz
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	writeFileStats("Cover.jpg")
	writeFileStats("NoFile.jpg")
	writeToFile("golang.txt", "go inanılmaz keyifli bir dilmiş. Sanırım C kökenlilerin kolayca öğrenebileceği bir dil.")
	readFile("golang.txt")
	//readFile("cover.jpg")
	readFileWithIO("golang.txt")
	//readFileWithIO("cover.jpg")

	// Aşağıdaki kod parçasında Player yapısından bir slice içeriğini
	// satır satır bir dosyaya aktarma işlemi yapılmaktadır
	players := []Player{
		Player{1, "baltazar", 80},
		Player{2, "orvel", 23},
		Player{3, "nadya", 48},
		Player{4, "obi van", 91},
		Player{5, "şumi", 77},
	}
	f, err := os.Create("players.dat")
	if err == nil {
		var content string = "Players\n"
		for _, player := range players {
			content += player.ToString()
		}
		f.WriteString(content)
	} else {
		fmt.Printf("Dosya oluşturulurken hata oluştu\n\t%s\n", err.Error())
	}
}

type Player struct {
	Id    int
	Title string
	Level int
}

func (p Player) ToString() string {
	// int değerleri string'e dönüştürmek için strconv paketindeki Itoa fonksiyonunu kullandık
	return strconv.Itoa(p.Id) + "|" + p.Title + "|" + strconv.Itoa(p.Level) + "\n"
}

//dosya içeriğini okuma örneği - 1
//parametre olarak gelen dosya içeriğini byte byte okuyup ekrana basacak
func readFile(fullpath string) {
	f, err := os.Open(fullpath)
	if err == nil {
		defer f.Close()
		fileInfo, _ := f.Stat()                    // dosya boyutunu bulmak için
		fileBytes := make([]byte, fileInfo.Size()) //dosya boyutu kadarlık bir slice oluşturulur
		read, _ := f.Read(fileBytes)               //dosya içeriği fileBytes kesitine yazılır. read'e okunan byte sayısı döner
		fmt.Printf("\n%s (%d byte okundu)\n\t%s\n", fullpath, read, fileBytes)
	} else {
		fmt.Printf("Dosyayı açamadım dostum\n\t%s\n", err.Error())
	}
}

// Dosya içeriğini okuma örneği - 2
func readFileWithIO(fullpath string) {
	content, err := ioutil.ReadFile(fullpath) //dosya içeriğini byte array olarak döndürür
	if err == nil {
		fmt.Println("'", string(content), "'") //content ASCII değerlerini barındırır. Bu nedenle string dönüşümü yapılmıştır. txt dışında jpg içeriğini okuduğumuzda ekrana nasıl bastığına bakın.
	} else {
		fmt.Printf("Dosya okunmaya çalışılırken hata oluştu\n\t%s\n", err.Error())
	}
}

//dosya oluşturma ve içerisine string içerik yazma örneği
func writeToFile(fullpath string, content string) {
	f, err := os.Create(fullpath)
	if err == nil { //dosya yaratılırken hata oluşmadıysa
		defer f.Close()                //fonksiyondan çıkarken çalışacak
		r, _ := f.WriteString(content) //içeriği dosyaya yazıyor
		fmt.Println("işlem sonucu", r)
	} else {
		fmt.Printf("Dosya oluşturulurken hata oluştu\n\t%s\n", err.Error())
	}
}

//bir dosya hakkında bilgiler edinme örneği
func writeFileStats(file string) {
	f, err := os.Open(file) //dosyayı açtık
	if err == nil {
		defer f.Close()      //writeFileStats fonksiyonundan çıkarken çalışacak
		fInfo, _ := f.Stat() // Stat fonksiyonu da Open gibi iki değer döndürür. _ error değeridir
		fmt.Println(fInfo.Size(), "bytes")
		fmt.Println("File name is ", fInfo.Name())
		fmt.Println("File modes are ", fInfo.Mode().String())
		fmt.Println(fInfo.ModTime(), "Last Changed")
	} else {
		fmt.Printf("Dosyayı açamadım dostum\n\t%s\n", err.Error())
	}
}
