package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samverrall/background-tasks/background"
	"github.com/samverrall/background-tasks/database"
	"github.com/samverrall/background-tasks/jobs"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bookingsRepo := database.NewBookingsRepo()

	serviceJobs := jobs.New(bookingsRepo)

	background.Go(serviceJobs.UnlockSlots(ctx))

	// Spin up HTTP server, register endpoints etc etc

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
