package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/laminne/notepod/pkg/models/entity"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		if _, err := db.NewCreateTable().Model(&entity.User{}).Exec(context.Background()); err != nil {
			log.Fatal(err)
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		if _, err := db.NewDropTable().Model(&entity.User{}).Exec(context.Background()); err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
