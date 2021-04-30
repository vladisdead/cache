package cfg

type CFG struct {
	Storage Storage
}

type Storage struct {
	Connstring     string // Информация для подключения к БД.
	MigrationsPath string // Путь до файлов миграции БД.
}
