package dao

import (
	"context"
	"github.com/mariojuzar/fleet-inventory/internal/domain/model"
	"github.com/mariojuzar/fleet-inventory/internal/domain/repository"
	"github.com/mariojuzar/fleet-inventory/internal/infrastructures/mysql"
	"github.com/rs/zerolog/log"
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

func (repo *SpaceCraftRepository) Update(ctx context.Context, model *model.SpaceCraft) error {
	//TODO implement me
	panic("implement me")
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
	query := `UPDATE space_crafts SET is_deleted = true WHERE id = ?`

	_, err := repo.DB.Conn.ExecContext(ctx, query, id)
	if err != nil {
		log.Err(err).Msg("Failed to soft delete space craft")
		return err
	}

	return nil
}

func (repo *SpaceCraftRepository) Fetch(ctx context.Context, filter *model.SpaceCraftFetchFilter) ([]model.SpaceCraft, error) {
	//TODO implement me
	panic("implement me")
}
