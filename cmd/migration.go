package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

func newMigrationCmd() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate tables",
	}

	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "apply new migrate ",
		RunE:  migrateUp,
	}

	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "rollback to previous db version",
		RunE:  migrateDown,
	}
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	return migrateCmd

}

func migrateUp(cmd *cobra.Command, _ []string) error {
	cfgPath, err := cmd.Flags().GetString("config")
	if err != nil {
		fmt.Println("invlaid flag for migrate command")
		return err
	}
	cfg := getConfig(cfgPath)
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	m, err := migrate.New("file://store/migrations", databaseUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) || err == nil {
		return nil
	}
	return err

}

func migrateDown(cmd *cobra.Command, _ []string) error {
	cfgPath, err := cmd.Flags().GetString("config")
	if err != nil {
		fmt.Println("invlaid flag for migrate command")
		return err
	}
	cfg := getConfig(cfgPath)
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	m, err := migrate.New("file://store/migrations", databaseUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = m.Down()
	if errors.Is(err, migrate.ErrNoChange) || err == nil {
		return nil
	}
	return err

}
