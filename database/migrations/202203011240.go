package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kachit/golang-api-skeleton/models"
	"gorm.io/gorm"
	"strconv"
)

func init() {
	var cnt = 1
	name := func() string {
		cnt++
		return "202203011240" + strconv.Itoa(cnt)
	}

	Migrations = append(Migrations, &gormigrate.Migration{
		ID: name() + "_create_users_id_seq",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`CREATE SEQUENCE "users__id_seq" 
									INCREMENT 1
									MINVALUE  1
									MAXVALUE 2147483647
									START 1
									CACHE 1;
								`).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec(`DROP SEQUENCE IF EXISTS "users_id_seq";`).Error
		},
	})

	Migrations = append(Migrations, &gormigrate.Migration{
		ID: name() + "_create_users_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`CREATE TABLE "users" (
								  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
								  "name" varchar(100) NOT NULL,
								  "email" varchar(100) NOT NULL,
								  "password" varchar(100) NOT NULL,
								  "created_at" timestamptz(6) NOT NULL,
								  "modified_at" timestamptz(6),
								  "deleted_at" timestamptz(6),
								  PRIMARY KEY ("id"),
								  CONSTRAINT ux_users_email UNIQUE (email)
								)
								;`).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(models.TableUsers)
		},
	})
}
