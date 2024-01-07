package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"
	"sagara/domain/entity"
	"sagara/internal/delivery/request"
	"sagara/internal/repository/psql/mapper"
	"sagara/internal/repository/psql/models"
	"sagara/pkg/common"

	"github.com/rocketlaunchr/dbq/v2"
)

func (repository *officeRepository) getListOfficeQuery() string {
	stmt := fmt.Sprintf(
		`SELECT
			o.id
			, o.name
			, d.path AS image_path
			, o.room_type
			, o.description
			, o.additional_data->>'capacity' AS additional_data_capacity
			, o.additional_data->>'table' AS additional_data_table
			, o.additional_data->>'chair' AS additional_data_chair
			, o.additional_data->'special_offer'->>'title' AS additional_data_special_offer_title
			, o.additional_data->'special_offer'->>'description' AS additional_data_special_offer_description
			, o.additional_data->'special_offer'->>'image_path' AS additional_data_special_offer_image_path
			, o.additional_data->'special_price'->>'image_path' AS additional_data_special_price_image_path
		FROM %s o
		JOIN %s d ON o.image_reference = d.id`,
		models.Office{}.TableName(),
		models.DataObject{}.TableName())

	return stmt
}

func (repository *officeRepository) GetOffice(ctx context.Context, options *request.Option) ([]*entity.Office, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: false, ConcreteStruct: models.OfficeList{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getListOfficeQuery()

	if options.RegionId > 0 {
		stmt += fmt.Sprintf(` WHERE o.region_id = $%d`, len(args)+1)
		args = append(args, options.RegionId)
	}

	stmt += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(args)+1, len(args)+2)
	args = append(args, options.Pagination.Limit)
	args = append(args, options.Pagination.Page)

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		offices := mapper.ToDomainListOffice(results.([]*models.OfficeList))
		return offices, nil
	} else {
		return nil, nil
	}
}

func (repository *officeRepository) CountOffice(ctx context.Context, options *request.Option) (int32, error) {
	var (
		args  []interface{}
		count sql.NullInt32
		err   error
	)

	stmt := repository.getListOfficeQuery()

	stmt = fmt.Sprintf(`SELECT COUNT(id) from (%s) as q`, stmt)

	err = repository.db.QueryRowContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int32, nil
}

func (repository *officeRepository) FindOfficeById(ctx context.Context, id common.ID) (*entity.Office, error) {
	var args []interface{}

	opts := &dbq.Options{SingleResult: true, ConcreteStruct: models.OfficeList{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	stmt := repository.getListOfficeQuery()

	stmt += ` WHERE o.id = $1`
	args = append(args, id.String())

	results := dbq.MustQ(ctx, repository.db, stmt, opts, args...)
	if results != nil {
		order := mapper.ToDomainOffice(results.(*models.OfficeList))
		return order, nil
	} else {
		return nil, nil
	}
}
