package sqlc

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testStore Store

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"
	migrationURL := "file://../../../db/migrations"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	DBsource, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to Docker: %s", err)
	}
	RunDBMigration(migrationURL, DBsource)
	connPool, err := pgxpool.New(context.Background(), DBsource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
