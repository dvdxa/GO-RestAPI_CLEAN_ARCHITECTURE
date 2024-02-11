package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/config"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/controller"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/db/postgres"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/logger"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/server"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/storage"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/usecase"
)

func Run() {
	cfg := config.InitConfigs()

	log := logger.InitLogger(&cfg.Logger)

	db := postgres.InitDB(&cfg.Database)

	storage := storage.NewStorage(db, log)

	service := usecase.NewService(storage, log)

	handler := controller.NewHandler(cfg, log, service)

	srv := server.NewServer(cfg, handler)

	var wg sync.WaitGroup

	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	go startServer(srv, ctx, &wg)
	waitForInterrupt(cancel, &wg)

	wg.Wait()
}

func startServer(srv *server.Server, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		err := srv.Run()
		if err != nil {
			log.Fatalf("failed to run server: %v\n", err)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("shutting down server gracefully")

		//5 sec to gracefully shutdown server
		shutDownCtx, shutDowncancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutDowncancel()

		err := srv.Shutdown(shutDownCtx)
		if err != nil {
			fmt.Printf("error shutting 	down server %v\n", err.Error())
		}
	}
}

func waitForInterrupt(cancel context.CancelFunc, wg *sync.WaitGroup) {

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer wg.Done()
		s := <-ch
		fmt.Printf("received signal %v\n", s)
		cancel()
	}()
}
