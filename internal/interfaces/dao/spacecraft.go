package dao

import (
	"context"
	"github.com/mariojuzar/fleet-inventory/internal/domain/model"
	"github.com/mariojuzar/fleet-inventory/internal/domain/repository"
	"github.com/mariojuzar/fleet-inventory/internal/infrastructures/mysql"
	"github.com/rs/zerolog/log"
	"time"
)

type SpaceCraftRepository struct {
	DB *mysql.DB
}

type SpaceCraftRepositoryOpts struct {
	DB *mysql.DB
}

func NewSpaceCraft(opts *SpaceCraftRepositoryOpts) repository.SpaceCraftRepository {
	return &SpaceCraftRepository{DB: opts.DB}
}

func (repo *SpaceCraftRepository) Insert(ctx context.Context, model *model.SpaceCraft) error {
	tx, err := repo.DB.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var args []any
	query := `INSERT INTO space_crafts (name, class, crew, image, value, status) VALUES (?, ?, ?, ?, ?, ?); `

	args = append(args, model.Name)
	args = append(args, model.Class)
	args = append(args, model.Crew)
	args = append(args, model.Image)
	args = append(args, model.Value)
	args = append(args, model.Status)

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		log.Err(err).Msg("Failed to insert space craft")
		return err
	}
	lastId, _ := res.LastInsertId()
	model.Id = int(lastId)

	armQuery, armArgs := repo.generateArmamentsInsertQuery(model)
	_, err = tx.ExecContext(ctx, armQuery, armArgs...)
	if err != nil {
		log.Err(err).Msg("Failed to insert space craft armaments")
		return err
	}

	return tx.Commit()
}

func (repo *SpaceCraftRepository) generateArmamentsInsertQuery(model *model.SpaceCraft) (string, []any) {
	query := ""
	var args []any

	for _, armament := range model.Armaments {
		query += `INSERT INTO space_crafts_armaments(space_craft_id, armament_id, qty) VALUES (?, ?, ?);`
		args = append(args, model.Id)
		args = append(args, armament.ArmamentId)
		args = append(args, armament.Quantity)
	}

	return query, args
}

func (repo *SpaceCraftRepository) Update(ctx context.Context, id int, model *model.SpaceCraftUpdate) error {
	tx, err := repo.DB.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var args []any

	query := `UPDATE space_crafts SET updated_at = ? `
	args = append(args, time.Now())

	if model.Name != nil {
		query += `, name = ?`
		args = append(args, model.Name)
	}

	if model.Crew != nil {
		query += `, crew = ?`
		args = append(args, model.Crew)
	}

	if model.Class != nil {
		query += `, class = ?`
		args = append(args, model.Class)
	}

	if model.Status != nil {
		query += `, status = ?`
		args = append(args, model.Status)
	}

	if model.Value != nil {
		query += `, value = ?`
		args = append(args, model.Value)
	}

	if model.Image != nil {
		query += `, image = ?`
		args = append(args, model.Image)
	}

	query += " WHERE id = ? and is_deleted = false"
	args = append(args, id)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (repo *SpaceCraftRepository) Get(ctx context.Context, id int) (*model.SpaceCraft, error) {
	query := `SELECT sc.id as id, sc.name name, sc.class as class, sc.crew as crew, sc.image as image, sc.value as value, sc.status as status, sc.created_at as created_at, sc.updated_at as updated_at 
			  FROM space_crafts sc
			  WHERE sc.id = ? and sc.is_deleted = false`

	result := &model.SpaceCraft{}
	err := repo.DB.Conn.GetContext(ctx, result, query, id)
	if err != nil {
		return nil, err
	}

	var armResult []model.SpaceCraftArmamentEmbed
	armQuery := `SELECT a.name as title, sca.qty as qty
				 FROM space_crafts_armaments sca
				 LEFT JOIN space_crafts sc ON sc.id = sca.space_craft_id and sc.is_deleted = false
				 LEFT JOIN armaments a ON a.id = sca.armament_id and a.is_deleted = false
				 WHERE sca.space_craft_id = ? and sca.is_deleted = false`
	err = repo.DB.Conn.SelectContext(ctx, &armResult, armQuery, id)
	if err != nil {
		return nil, err
	}
	result.Armaments = armResult

	return result, err
}

func (repo *SpaceCraftRepository) Delete(ctx context.Context, id int) error {
	query := `UPDATE space_crafts SET is_deleted = true, updated_at = ? WHERE id = ? and is_deleted = false`

	_, err := repo.DB.Conn.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		log.Err(err).Msg("Failed to soft delete space craft")
		return err
	}

	return nil
}

func (repo *SpaceCraftRepository) Fetch(ctx context.Context, filter *model.SpaceCraftFetchFilter) ([]model.SpaceCraft, error) {
	var args []any
	query := `SELECT sc.id as id, sc.name name, sc.class as class, sc.crew as crew, sc.image as image, sc.value as value, sc.status as status, sc.created_at as created_at, sc.updated_at as updated_at 
			  FROM space_crafts sc
			  WHERE sc.is_deleted = false `

	if filter.Name != "" {
		query += ` AND sc.name LIKE ? `
		args = append(args, filter.Name)
	}
	if filter.Status != "" {
		query += ` AND sc.status = ? `
		args = append(args, filter.Status)
	}
	if filter.Class != "" {
		query += ` AND sc.class = ? `
		args = append(args, filter.Class)
	}

	var result []model.SpaceCraft
	err := repo.DB.Conn.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Err(err).Msg("Failed to fetch space craft")
		return nil, err
	}
	return result, err
}
