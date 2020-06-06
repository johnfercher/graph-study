package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/johnfercher/graph-study/russiandoll/models"
)

type RussianDollRepository struct {
	db *sql.DB
}

func NewRussianDollRepository(db *sql.DB) *RussianDollRepository {
	return &RussianDollRepository{
		db: db,
	}
}

func (self *RussianDollRepository) GetById(ctx context.Context, id string) (*models.RussianDoll, error) {
	results, err := self.db.QueryContext(ctx, fmt.Sprintf(`WITH RECURSIVE get_dolls AS (
    SELECT
        c.id as doll_id
         , cp.parent_id
    FROM russian_doll as c
             LEFT JOIN russian_doll_parent cp on cp.russian_doll_id = c.id
    WHERE
            c.id = '%s'
    UNION ALL
    SELECT
        cp2.russian_doll_id
         , cp2.parent_id
    FROM russian_doll_parent cp2
             INNER JOIN get_dolls o on o.doll_id = cp2.parent_id
)
SELECT c.id, c.name, parent_id
FROM get_dolls gc
         JOIN russian_doll c on c.id = gc.doll_id`, id))
	if err != nil {
		return nil, err
	}

	russianDolls := make(map[string]models.RussianDoll)
	var idRead sql.NullString
	var name sql.NullString
	var parentId sql.NullString

	for results.Next() {
		err = results.Scan(&idRead, &name, &parentId)
		if err != nil {
			return nil, err
		}

		russianDoll := models.RussianDoll{
			Id:       idRead.String,
			Name:     name.String,
			ParentId: parentId.String,
		}

		russianDolls[russianDoll.Id] = russianDoll
	}

	russianDoll := self.BuildTree(id, russianDolls)

	return &russianDoll, nil
}

func (self *RussianDollRepository) GetAll(ctx context.Context) ([]*models.RussianDoll, error) {
	results, err := self.db.QueryContext(ctx, `SELECT id, name, parent_id FROM russian_doll LEFT JOIN russian_doll_parent as p ON p.russian_doll_id = id`)
	if err != nil {
		return nil, err
	}

	russianDolls := []*models.RussianDoll{}

	for results.Next() {
		var idRead sql.NullString
		var name sql.NullString
		var parentId sql.NullString

		err = results.Scan(&idRead, &name, &parentId)
		if err != nil {
			return nil, err
		}

		russianDoll := &models.RussianDoll{
			Id:       idRead.String,
			Name:     name.String,
			ParentId: parentId.String,
		}

		russianDolls = append(russianDolls, russianDoll)
	}

	return russianDolls, nil
}

func (self *RussianDollRepository) Create(ctx context.Context, russianDoll *models.RussianDoll) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO russian_doll (id, name) values ('%s', '%s')`, russianDoll.Id, russianDoll.Name))
	if err != nil {
		return err
	}

	return nil
}

func (self *RussianDollRepository) Update(ctx context.Context, russianDoll *models.RussianDoll) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`UPDATE russian_doll SET name = '%s' WHERE id = '%s'`, russianDoll.Name, russianDoll.Id))
	if err != nil {
		return err
	}

	return nil
}

func (self *RussianDollRepository) Delete(ctx context.Context, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM russian_doll WHERE id = '%s'`, id))
	if err != nil {
		return err
	}

	return nil
}

func (self *RussianDollRepository) RemoveFromRussianDoll(ctx context.Context, parentId string, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`DELETE FROM russian_doll_parent WHERE parent_id = '%s' AND russian_doll_id='%s'`, parentId, id))
	if err != nil {
		return err
	}

	return nil
}

func (self *RussianDollRepository) AddToRussianDoll(ctx context.Context, parentId string, id string) error {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO russian_doll_parent (russian_doll_id, parent_id) VALUES ('%s', '%s')`, id, parentId))
	if err != nil {
		return err
	}

	return nil
}

func (self *RussianDollRepository) BuildTree(parentId string, russianDollMap map[string]models.RussianDoll) models.RussianDoll {
	parent := russianDollMap[parentId]
	delete(russianDollMap, parentId)

	russianDoll, _ := self.buildTree(parent, russianDollMap)
	return russianDoll
}

func (self *RussianDollRepository) buildTree(parent models.RussianDoll, russianDollMap map[string]models.RussianDoll) (models.RussianDoll, map[string]models.RussianDoll) {
	var dollsFromParent []models.RussianDoll

	for _, doll := range russianDollMap {
		if parent.Id == doll.ParentId {
			dollsFromParent = append(dollsFromParent, doll)
			delete(russianDollMap, doll.Id)
		}
	}

	for _, parentDoll := range dollsFromParent {
		builtParentDoll, remain := self.buildTree(parentDoll, russianDollMap)

		russianDollMap = remain
		parent.RussianDolls = append(parent.RussianDolls, builtParentDoll)
	}

	return parent, russianDollMap
}
