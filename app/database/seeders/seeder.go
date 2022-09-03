package seeders

import (
	"github.com/ilhamadikusuma31/ditoko/app/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{} //interface kosonng, Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun

}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {

	for _, j := range RegisterSeeders(db) {
		err := db.Debug().Create(j.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
