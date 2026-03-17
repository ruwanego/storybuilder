package repositories

import (
	"context"

	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	"github.com/storybuilder/storybuilder/domain/boundary/repositories"
	"github.com/storybuilder/storybuilder/domain/entities"
)

// SampleMySQLRepository is an example repository that implements test database functionality.
type SampleMySQLRepository struct {
	db adapters.DBAdapterInterface
}

// NewSampleMySQLRepository creates a new instance of the repository.
func NewSampleMySQLRepository(dbAdapter adapters.DBAdapterInterface) repositories.SampleRepositoryInterface {
	return &SampleMySQLRepository{db: dbAdapter}
}

// Get retrieves a collection of Samples.
func (repo *SampleMySQLRepository) Get(ctx context.Context) ([]entities.Sample, error) {
	query := `SELECT id, name, password
				FROM sample`
	parameters := map[string]any{}
	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}
	return repo.mapResult(result)
}

// GetByID retrieves a single Sample.
func (repo *SampleMySQLRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {
	// NOTE: DBAdapters support named parameters, and you don't have to
	// 		 worry about the order in which those parameters are declared in the
	//		 query and in the parameters map. The DBAdapter will take care of that.
	query := `SELECT id, name, password
				FROM sample
				WHERE id=?id`
	parameters := map[string]any{
		"id": id,
	}
	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return entities.Sample{}, err
	}
	mapped, err := repo.mapResult(result)
	if err != nil {
		return entities.Sample{}, err
	}
	if len(mapped) == 0 {
		return entities.Sample{}, nil
	}
	return mapped[0], nil
}

// Add adds a new sample.
func (repo *SampleMySQLRepository) Add(ctx context.Context, sample entities.Sample) error {
	query := `INSERT INTO sample
				(name, password)
				VALUES(?name, ?password)
				`
	parameters := map[string]any{
		"name":     sample.Name,
		"password": sample.Password,
	}
	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}
	return nil
}

// Edit updates an existing sample identified by the id.
func (repo *SampleMySQLRepository) Edit(ctx context.Context, sample entities.Sample) error {
	query := `UPDATE sample
				SET name=?name, password=?password
				WHERE id=?id
				`
	parameters := map[string]any{
		"id":       sample.ID,
		"name":     sample.Name,
		"password": sample.Password,
	}
	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an existing sample identified by id.
func (repo *SampleMySQLRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM sample
				WHERE id=?id
				`
	parameters := map[string]any{
		"id": id,
	}
	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}
	return nil
}

// mapResult maps the result to entities.
func (repo *SampleMySQLRepository) mapResult(result []map[string]any) (samples []entities.Sample, err error) {
	// Applying type assertion in this manner will result in a panic when the db data structure changes.
	// This defer-recover pattern is used to recover from the panic and to return an error instead.
	// Notice the use of `named returned values` for this function (without which the recover pattern will not work).
	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
		}
	}()
	for _, row := range result {
		samples = append(samples, entities.Sample{
			ID:   int(row["id"].(int64)),
			Name: string(row["name"].([]byte)),
		})
	}
	return samples, err
}
