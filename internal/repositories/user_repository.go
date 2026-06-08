package repositories

type UserRepository struct {
	client *MongoRepositoryContext
}

func NewUserRepository(uri, dbName, collectionName string) (*UserRepository, error) {
	mongoRepo, err := NewMongoRepositoryContext(uri, dbName, collectionName)
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		client: mongoRepo,
	}, nil
}
