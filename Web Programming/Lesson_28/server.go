/*
	Temel bir ORM aracını nasıl kullanabileceğimizi öğrenmeye çalışıyoruz.
*/
package main

import (
	"fmt"

	"entity/southwind"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := gorm.Open("sqlite3", "db\\southwind.sdb")
	db.LogMode(true)
	defer db.Close()
	if err == nil {
		//db.SingularTable(true)
		db.AutoMigrate(&Southwind.Employee{}, &Southwind.Email{})
		db.Model(&Southwind.Employee{}).Related(&Southwind.Email{})

		burakMails := []Southwind.Email{
			Southwind.Email{Mail: "selim@buraksenyurt.com", IsActive: true},
			Southwind.Email{Mail: "burak.senyurt@southwind.com", IsActive: false},
			Southwind.Email{Mail: "burakselimsenyurt@gmail.com", IsActive: true},
		}

		burak := Southwind.Employee{FirstName: "burak", LastName: "senyurt", Emails: burakMails}
		db.Create(&burak)

		loraMails := []Southwind.Email{
			Southwind.Email{Mail: "lora@kimbilll.moon", IsActive: true},
			Southwind.Email{Mail: "kimbill.the.black.lora@southwind.com", IsActive: true},
		}
		lora := Southwind.Employee{FirstName: "Lora", LastName: "Kimbılll", Emails: loraMails}
		db.Create(&lora)

		WriteToScreen(burak)
		WriteToScreen(lora)

		var burki Southwind.Employee
		db.Find(&burki, "ID=?", 1) //Önce
		db.Model(&burki).Update("LastName", "Selim Senyurt")
		WriteToScreen(burki)

		var buffon Southwind.Employee

		db.Model(&buffon).Where("ID=?", 2).Updates(map[string]interface{}{"FirstName": "Cianluici", "LastName": "Buffon"})
		db.First(&buffon, 2) //Direkt primary key üstünden(varsayılan olarak ID) arama yapar
		WriteToScreen(buffon)
	} else {
		fmt.Println(err.Error())
	}
}

func WriteToScreen(e Southwind.Employee) {
	fmt.Printf("%d\t%s,%s,%s\n", e.ID, e.FirstName, e.LastName, e.CreatedAt)
	for _, email := range e.Emails {
		fmt.Printf("\t%d:%s\n", email.ID, email.Mail)
	}
}
