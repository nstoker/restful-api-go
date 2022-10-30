package service

import (
	"context"
	"time"
)

func (s Service) Delete(ctx context.Context, id int) error {
	todo, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	tx, err := s.repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	now := time.Now().UTC()
	todo.DeletedOn = &now
	err = s.repo.Update(ctx, todo)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
