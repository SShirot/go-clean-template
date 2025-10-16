// Package translation provides translation repository implementations
package translation

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/di"
	domain "github.com/evrone/go-clean-template/internal/domain/translation"
	pkgpg "github.com/evrone/go-clean-template/pkg/postgres"
)

// postgresRepo implements translation repository for PostgreSQL
type postgresRepo struct {
	db di.PostgresInterface
}

// NewPostgresRepo creates a new PostgreSQL translation repository
func NewPostgresRepo(db di.PostgresInterface) di.TranslationRepoInterface {
	return &postgresRepo{
		db: db,
	}
}

// Store stores a new translation
func (r *postgresRepo) Store(ctx context.Context, translation interface{}) error {
	pg, ok := r.db.(*pkgpg.Postgres)
	if !ok || pg == nil || pg.Pool == nil {
		return fmt.Errorf("invalid postgres connection")
	}

	// Expect domain translation; accept entity by best-effort if needed
	tDomain, ok := translation.(*domain.Translation)
	if !ok {
		if tv, ok2 := translation.(domain.Translation); ok2 {
			tDomain = &tv
		} else {
			return fmt.Errorf("invalid translation type")
		}
	}

	sql, args, err := pg.Builder.
		Insert("history").
		Columns("source, destination, original, translation").
		Values(tDomain.Source, tDomain.Destination, tDomain.Original, tDomain.Translation).
		ToSql()
	if err != nil {
		return fmt.Errorf("repo - Store - build: %w", err)
	}

	_, err = pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo - Store - exec: %w", err)
	}

	return nil
}

// GetHistory retrieves translation history
func (r *postgresRepo) GetHistory(ctx context.Context, limit, offset int) ([]interface{}, error) {
	pg, ok := r.db.(*pkgpg.Postgres)
	if !ok || pg == nil || pg.Pool == nil {
		return nil, fmt.Errorf("invalid postgres connection")
	}

	sql, _, err := pg.Builder.
		Select("source, destination, original, translation").
		From("history").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo - GetHistory - build: %w", err)
	}

	rows, err := pg.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("repo - GetHistory - query: %w", err)
	}
	defer rows.Close()

	var out []interface{}
	for rows.Next() {
		e := domain.Translation{}
		if err := rows.Scan(&e.Source, &e.Destination, &e.Original, &e.Translation); err != nil {
			return nil, fmt.Errorf("repo - GetHistory - scan: %w", err)
		}
		out = append(out, e)
	}

	return out, nil
}

// GetByID retrieves a translation by ID
func (r *postgresRepo) GetByID(ctx context.Context, id string) (interface{}, error) {
	// Not used in demo; implement if needed
	return nil, fmt.Errorf("not implemented")
}

// Delete removes a translation
func (r *postgresRepo) Delete(ctx context.Context, id string) error {
	// Not used in demo; implement if needed
	return fmt.Errorf("not implemented")
}
