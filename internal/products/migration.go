package products

import "gorm.io/gorm"

func MigrateDb(db *gorm.DB) error {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Product{})
	return err
}
