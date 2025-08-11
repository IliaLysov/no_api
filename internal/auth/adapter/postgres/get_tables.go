package postgres

import (
	"context"
	"fmt"
	"log"
)

func (p *Postgres) GetTables(ctx context.Context) (err error) {
	rows, err := p.pool.Query(ctx, `
		SELECT tablename 
		FROM pg_catalog.pg_tables
		WHERE schemaname NOT IN ('pg_catalog', 'information_schema');
	`)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
	return nil
}
