package infra

import (
	"database/sql"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func (repo *Repository) sanitizeString(name string) error {
	if !regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`).MatchString(strings.ToUpper(name)) {
		return fmt.Errorf("invalid identifier format: %s", name)
	}

	reserved := []string{"SELECT", "DROP", "INSERT", "UPDATE", "DELETE", "ALTER", "CREATE", "GRANT", "REVOKE", "UNION", "WHERE", "ORDER", "GROUP"}

	if slices.Contains(reserved, strings.ToUpper(name)) {
		return fmt.Errorf("reserved word not allowed: %s", name)
	}

	return nil
}

func (repo *Repository) makeData(data map[string]any) ([]string, []string, []any, error) {
	if len(data) == 0 {
		return nil, nil, nil, fmt.Errorf("no data provided")
	}

	columns := make([]string, 0, len(data))
	placeholders := make([]string, 0, len(data))
	values := make([]any, 0, len(data))

	i := 1

	for col, val := range data {
		err := repo.sanitizeString(fmt.Sprintf("%s", col))
		if err != nil {
			return nil, nil, nil, err
		}
		columns = append(columns, col)
		values = append(values, val)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		i++
	}
	return columns, placeholders, values, nil
}

func (repo *Repository) Insert(table string, data map[string]any) error {
	columns, placeholders, values, err := repo.makeData(data)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
	_, err = repo.db.Exec(query, values...)
	return err
}

func (repo *Repository) UnsafeRaw(query string)
