package jobs

import (
	"context"
	"log"
	"time"

	"github.com/samverrall/background-tasks/background"
	"github.com/samverrall/background-tasks/database"
)

type Jobs struct {
	db database.BookingsRepo
}

func New(db database.BookingsRepo) *Jobs {
	return &Jobs{
		db: db,
	}
}

func (j *Jobs) UnlockSlots(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(time.Second * 5)

		for range t.C { // iterates each time a tick is delivered in the channnel
			log.Printf("Attempting to unlock stale time slots... \n")

			if err := j.db.UnlockBookings(ctx); err != nil {
				log.Printf("failed to unlock bookings: %s \n", err.Error())
			}
		}

	}
}
