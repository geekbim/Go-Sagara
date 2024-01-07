package postgres_repository

import (
	"context"
	"sagara/domain/entity"
	"sagara/internal/repository/psql/helper"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *userRepository) StoreUser(ctx context.Context, user *entity.User) error {
	var err error

	err = dbq.Tx(ctx, repository.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		err = helper.StoreUser(ctx, E, user)
		if err != nil {
			return
		}

		txCommit()
	})

	return err
}
