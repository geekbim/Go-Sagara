package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/internal/repository/psql/mapper"
	"sagara/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *regionRepository) getListRegionQuery() string {
	stmt := fmt.Sprintf(`SELECT r.id, r.name
		FROM %s r`, models.Region{}.TableName())

	return stmt
}

func (repository *regionRepository) GetRegion(ctx context.Context, options *request.Option) ([]*entity.Region, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.RegionList{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getListRegionQuery()

	if options.Key != "" {
		stmt += fmt.Sprintf(` WHERE name ILIKE $%d`, len(args)+1)
		args = append(args, fmt.Sprintf("%%%s%%", options.Key))
	}

	stmt += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(args)+1, len(args)+2)
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		regions := mapper.ToDomainListRegion(results.([]*models.RegionList))
		return regions, nil
	} else {
		return nil, nil
	}
}

func (repository *regionRepository) CountRegion(ctx context.Context, options *request.Option) (int32, error) {
	var (
		args  []interface{}
		count sql.NullInt32
		err   error
	)

	stmt := repository.getListRegionQuery()

	stmt = fmt.Sprintf(`SELECT COUNT(id) from (%s) as q`, stmt)

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}
