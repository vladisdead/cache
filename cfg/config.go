package cfg

func New() (*CFG, error) {
	cfg := CFG{
		Storage: Storage{
			Connstring: "user=test password=qwe host= port=5432 database=cache sslmode=disable",
		},
	}
	return &cfg, nil
}
