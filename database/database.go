package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type HashName struct {
	Id    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
	Items []Item
}

type Item struct {
	ShadowPayId uint `gorm:"uniqueIndex"`
	HashNameId  uint
	HashName    HashName `gorm:"foreignKey:HashNameId;references:Id;constraint:OnDelete:CASCADE"`
	Price       float64
}

func Main() {
	dbHost := "localhost" // IP-адрес вашего удаленного сервера
	dbName := "test"      // Имя вашей базы данных
	dbUsername := "root"  // Имя пользователя базы данных
	dbPass := "root"      // Пароль пользователя базы данных
	dbPort := "5432"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUsername, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&HashName{}, &Item{})
	if err != nil {
		log.Fatal(err)
	}

	itemsSlice := []Item{
		{ShadowPayId: 1, HashNameId: uint(7), Price: 11.1},
		{ShadowPayId: 2, HashNameId: uint(7), Price: 22.2},
		{ShadowPayId: 3, HashNameId: uint(7), Price: 33.3},
	}
	err = Upsert(itemsSlice, 500,
		[]string{"shadow_pay_id"},
		[]string{"shadow_pay_id", "hash_name_id", "price"},
		db,
	)
	if err != nil {
		log.Fatal(err)
	}
	//DB::table()->updateOrInsert()
	//DB::table('apartments')->upsert($data, 'id');
	// Create
	//hashNameOne := HashName{Name: "One"}
	//hashNameTwo := HashName{Name: "Two"}
	//db.Create(&hashNameOne)
	//db.Create(&hashNameTwo)
	//db.Create(&Item{ShadowPayId: 1111, HashNameId: 7, Name: "one_1"})
	//db.Create(&Item{ShadowPayId: 222, HashName: hashNameOne, Name: "one_2"})
	//db.Create(&Item{ShadowPayId: 333, HashName: hashNameOne, Name: "one_3"})
	//db.Create(&Item{ShadowPayId: 444, HashName: hashNameTwo, Name: "two_1"})
	//db.Create(&Item{ShadowPayId: 555, HashName: hashNameTwo, Name: "two_2"})

	// GET
	// all
	//var hashNames []HashName
	//result := db.Find(&hashNames)
	//fmt.Println(hashNames)
	//fmt.Println(result.RowsAffected)
	// first
	// where
	//var hashName HashName
	//result := db.Where("name = ?", "Two").First(&hashName)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	fmt.Println("not found")
	//} else {
	//	fmt.Println(hashName)
	//db.Where("hash_name_id = ?", hashName.Id).Delete(&Item{})
	//db.Delete(&hashName)
	//}
	// UPSERT
	//var items []Item
	//var hash HashName
	//db.Where("name = ?", "Two").First(&hash)
	//for i := 0; i < 10; i++ {
	//	newHash := hash
	//	if i == 8 {
	//		items = append(items, Item{ShadowPayId: uint(5), HashName: newHash, Name: "new_3"})
	//		continue
	//	}
	//	items = append(items, Item{ShadowPayId: uint(i), HashName: newHash, Name: "new_3"})
	//
	//}

	//for _, item := range items {
	//	db.Create(&item)
	//}
	// Transaction
	// auto
	//db.Transaction(func(tx *gorm.DB) error {
	//	for _, item := range items {
	//		if err := tx.Create(&item).Error; err != nil {
	//			return err
	//		}
	//	}
	//
	//	return nil
	//})

	//result := db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "shadow_pay_id"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"hash_name_id", "name"}),
	//}).Create(&items)
	//if result.Error != nil {
	//	fmt.Println("Какая-то ошибка")
	//	fmt.Println(err)
	//}
}

func Upsert(value interface{}, sizeChunk int, conflictFields, insertFields []string, db *gorm.DB) error {
	var conflictColumns []clause.Column
	for _, conflictName := range conflictFields {
		conflictColumns = append(conflictColumns, clause.Column{Name: conflictName})
	}

	result := db.Clauses(clause.OnConflict{
		Columns:   conflictColumns,
		DoUpdates: clause.AssignmentColumns(insertFields),
	}).CreateInBatches(value, sizeChunk)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
