package store

import (
	"log"

	"github.com/robfig/cron/v3"
)

func (pg PgStore) ScheduleCleanUp() {
	c := cron.New()
	c.AddFunc("CLEANUP every 1 min", func() {
		err := cleanupDeletedTodo(pg)
		if err != nil {
			log.Println("Failed to cleanup todos", err)
		}
	})
	c.Start()
	select {}
}

func cleanupDeletedTodo(pg PgStore) error {
	_, err := pg.db.Exec(`
        DELETE FROM todo
        WHERE status = 'deleted'
          AND created_at < NOW() - INTERVAL '1 minutes'
    `)
	return err
}
