//go:build e2e

package e2e

import "github.com/jackc/pgx/v5/pgxpool"

var TestDbInstance *pgxpool.Pool
