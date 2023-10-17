package database

import "context"

type BookingsRepo interface {
	UnlockBookings(ctx context.Context) error
}

type Bookings struct{}

func NewBookingsRepo() *Bookings {
	return &Bookings{}
}

func (b *Bookings) UnlockBookings(ctx context.Context) error {
	// db.Exec("DELETE FROM tmp_bookings....")
	return nil
}
