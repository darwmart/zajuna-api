package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"zajunaApi/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Obtener usuarios con paginación o filtros
func (r *UserRepository) GetUsers(filters map[string]string, page, pageSize int) ([]models.User, int, error) {
	var whereClauses []string
	var args []interface{}
	argIndex := 1

	allowedFilters := map[string]string{
		"id":        "id",
		"username":  "username",
		"firstname": "firstname",
		"lastname":  "lastname",
		"email":     "email",
		"idnumber":  "idnumber",
		"auth":      "auth",
	}

	for key, column := range allowedFilters {
		if val, ok := filters[key]; ok && val != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("%s ILIKE $%d", column, argIndex))
			args = append(args, "%"+val+"%")
			argIndex++
		}
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// total registros
	var total int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM mdl_user %s", whereSQL)
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// consulta principal
	var query string
	if len(whereClauses) > 0 {
		query = fmt.Sprintf(`
			SELECT id, username, firstname, lastname, email, auth
			FROM mdl_user
			%s
			ORDER BY id`, whereSQL)
	} else {
		offset := (page - 1) * pageSize
		query = fmt.Sprintf(`
			SELECT id, username, firstname, lastname, email, auth
			FROM mdl_user
			%s
			ORDER BY id
			LIMIT %d OFFSET %d`, whereSQL, pageSize, offset)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Auth); err != nil {
			return nil, 0, err
		}
		u.FullName = u.FirstName + " " + u.LastName
		users = append(users, u)
	}

	return users, total, nil
}
