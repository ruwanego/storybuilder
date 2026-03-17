package repositories

import (
	"context"
	"fmt"

	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	"github.com/storybuilder/storybuilder/domain/boundary/repositories"
	"github.com/storybuilder/storybuilder/domain/entities"
)

// SampleSQLRepository is a generic SQL repository that implements database functionality
// agnostic of the underlying dialect (MySQL, Postgres, etc.) as the adapter handles parameter translation natively.
type SampleSQLRepository struct {
	db        adapters.DBAdapterInterface
	tableName string
}

// NewSampleSQLRepository creates a new instance of the repository.
func NewSampleSQLRepository(dbAdapter adapters.DBAdapterInterface, tableName string) repositories.SampleRepositoryInterface {
	if tableName == "" {
		tableName = "sample"
	}
	return &SampleSQLRepository{
		db:        dbAdapter,
		tableName: tableName,
	}
}

// Get retrieves a collection of Samples.
func (repo *SampleSQLRepository) Get(ctx context.Context) ([]entities.Sample, error) {
	query := fmt.Sprintf(`SELECT id, name, password FROM %s`, repo.tableName)
	parameters := map[string]any{}
	result, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}
	return repo.mapResult(result)
}

// GetByID retrieves a single Sample.
func (repo *SampleSQLRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {
	// NOTE: DBAdapters support named parameters natively depending on exact driver syntax,
	// and handles the internal dialect translations natively.
	query := fmt.Sprintf(`SELECT id, name, password FROM %s WHERE id=?id`, repo.tableName)
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
func (repo *SampleSQLRepository) Add(ctx context.Context, sample entities.Sample) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, password) VALUES(?name, ?password)`, repo.tableName)
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
func (repo *SampleSQLRepository) Edit(ctx context.Context, sample entities.Sample) error {
	query := fmt.Sprintf(`UPDATE %s SET name=?name, password=?password WHERE id=?id`, repo.tableName)
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
func (repo *SampleSQLRepository) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=?id`, repo.tableName)
	parameters := map[string]any{
		"id": id,
	}
	_, err := repo.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}
	return nil
}

// mapResult maps the raw database map result to domain entities.
func (repo *SampleSQLRepository) mapResult(result []map[string]any) (samples []entities.Sample, err error) {
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
