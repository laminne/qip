package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/approversrsrs/qip/cmd/bun/migrations"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	db := bun.NewDB(
		sql.OpenDB(
			pgdriver.NewConnector(
				pgdriver.WithDSN("postgres://postgres:qip@localhost:5432/qip?sslmode=disable"),
			),
		),
		pgdialect.New(),
	)

	app := &cli.App{
		Name: "bun",

		Commands: []*cli.Command{
			newDBCommand(migrate.NewMigrator(db, migrations.Migrations)),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newDBCommand(mi *migrate.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "qip Database migration tools",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create tables",
				Action: func(context *cli.Context) error {
					return mi.Init(context.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(context *cli.Context) error {
					if err := mi.Lock(context.Context); err != nil {
						return err
					}
					defer mi.Unlock(context.Context)

					g, err := mi.Migrate(context.Context)
					if err != nil {
						return err
					}

					if g.IsZero() {
						fmt.Println("Nothing to do (database is already migrated")
						return nil
					}

					fmt.Printf("migrated: %s\n", g)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(context *cli.Context) error {
					if err := mi.Lock(context.Context); err != nil {
						return err
					}
					defer mi.Unlock(context.Context)

					g, err := mi.Rollback(context.Context)
					if err != nil {
						return err
					}

					if g.IsZero() {
						fmt.Println("Nothing to do")
						return nil
					}

					fmt.Printf("rolled back %s\n", g)
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "create migration file",
				Action: func(context *cli.Context) error {
					name := strings.Join(context.Args().Slice(), "_")
					m, err := mi.CreateGoMigration(context.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration file %s (%s)\n", m.Name, m.Path)
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migration status",
				Action: func(context *cli.Context) error {
					m, err := mi.MigrationsWithStatus(context.Context)
					if err != nil {
						return err
					}

					fmt.Printf("migrations: %s\n", m)
					fmt.Printf("Unapplied Migrations: %s\n", m.Unapplied())
					fmt.Printf("Last migraton group: %s\n", m.LastGroup())
					return nil
				},
			},
		},
	}
}
