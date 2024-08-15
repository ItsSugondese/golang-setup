package database

import "gorm.io/gorm"

func GetAllSchemas(db *gorm.DB) ([]string, error) {
	var schemas []string
	query := `
        SELECT nspname
        FROM pg_namespace
        WHERE nspname NOT IN ('information_schema', 'pg_catalog')
        AND nspname NOT LIKE 'pg_%';
    `
	err := db.Raw(query).Scan(&schemas).Error
	if err != nil {
		return nil, err
	}
	return schemas, nil
}
