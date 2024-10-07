//go:build integration

package integration_tests

import "github.com/jackc/pgx/v5/pgxpool"

var TestDbInstance *pgxpool.Pool
