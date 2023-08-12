package config

//func InitDb() (*gorm.DB, error) {
//	var err error
//	dsn := os.Getenv("DB_DSN")
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Silent),
//	})
//	if err != nil {
//		return nil, err
//	}
//	db.AutoMigrate(&model.User{}, &model.Post{})
//	return db, nil
//}
