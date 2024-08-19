package database

import "gorm.io/gorm"

func GetAllSchemasRepo(db *gorm.DB) ([]string, error) {
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

func SchemaExistsRepo(db *gorm.DB, schemaName string) (bool, error) {
	var exists bool
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM pg_namespace
            WHERE nspname = ?
        );
    `
	err := db.Raw(query, schemaName).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
