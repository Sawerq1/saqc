package repo

func (repo *SQLRepository) Init() {
	if _, err := repo.Commands.Init.ExecContext(repo.Context); err != nil {
		repo.Logger.Panic("Cannot exec init command")
	}
}

func (repo *SQLRepository) Seed() {
	if _, err := repo.Commands.Seed.ExecContext(repo.Context); err != nil {
		repo.Logger.Panic("Cannot exec seed command")
	}
}
