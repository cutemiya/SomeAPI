package titlerepo

import (
	"api/model"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"reflect"
)

type Repository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewRepository(logger *zap.SugaredLogger, db *sqlx.DB) Repository {
	return Repository{
		logger: logger,
		db:     db,
	}
}

func (r Repository) InsertData(ctx context.Context,
	userModel model.User,
	locationModel model.Location,
	loginModel model.Login,
	pictureModel model.Picture) error {
	var (
		psql       = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		locationId int
		loginId    int
		pictureId  int
	)

	query := psql.Insert("Locations").
		Columns("NumberLocation", "NameLocation", "City", "StateLocation", "Country", "Postcode", "Latitude", "Longitude", "OffsetLocation", "Description").
		Values(locationModel.Number,
			locationModel.Name,
			locationModel.City,
			locationModel.State,
			locationModel.Country,
			fmt.Sprintf("%s", locationModel.Postcode),
			locationModel.Latitude,
			locationModel.Longitude,
			locationModel.Offset,
			locationModel.Description).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&locationId)
	r.logger.Debugf("%s", reflect.TypeOf(locationId))
	if err != nil {
		return err
	}

	query = psql.Insert("Logins").
		Columns("Uuid", "Username", "Password", "Salt", "Md5", "Sha1", "Sha256").
		Values(loginModel.Uuid,
			loginModel.Username,
			loginModel.Password,
			loginModel.Salt,
			loginModel.MD5,
			loginModel.SHA1,
			loginModel.SHA256).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err = query.QueryRowContext(ctx).Scan(&loginId)
	if err != nil {
		return err
	}

	query = psql.Insert("Pictures").
		Columns("Large", "Medium", "Thumbnail").
		Values(pictureModel.Large, pictureModel.Medium, pictureModel.Thumbnail).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err = query.QueryRowContext(ctx).Scan(&pictureId)
	if err != nil {
		return err
	}

	sqlQuery, args, err := psql.Insert("Users").
		Columns("Gender", "Title", "First", "Last", "Email", "Date", "Age",
			"RegisteredDate", "RegisteredAge", "Phone", "Cell", "IdName",
			"IdValue", "Nat", "LocationId", "LoginId", "PictureId").
		Values(userModel.Gender,
			userModel.Title,
			userModel.First,
			userModel.Last,
			userModel.Email,
			userModel.Date,
			userModel.Age,
			userModel.RegisteredDate,
			userModel.RegisteredAge,
			userModel.Phone,
			userModel.Cell,
			userModel.IdName,
			userModel.IdValue,
			userModel.Nat,
			locationId,
			loginId,
			pictureId).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) SelectData(ctx context.Context) (string, error) {
	var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sqlQuery, args, err := psql.Select("Gender", "Title", "First", "Last").From("Users").ToSql()
	if err != nil {
		return "", err
	}

	row, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return "", err
	}

	var (
		result string
		gender string
		title  string
		first  string
		last   string
	)

	for row.Next() {
		err := row.Scan(&gender, &title, &first, &last)
		if err != nil {
			return "", err
		}

		result += fmt.Sprintf("%s %s %s %s \n", gender, title, first, last)

	}

	return result, nil
}
