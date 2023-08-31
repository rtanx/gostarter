package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rtanx/gostarter/env"
	"github.com/rtanx/gostarter/internal/http/routes"
	"github.com/rtanx/gostarter/internal/infrastructure/logger"
	"github.com/rtanx/gostarter/internal/infrastructure/validation"
	"github.com/spf13/cobra"
	// "pandatech.io/dashboard/env"
	// "pandatech.io/dashboard/internal/db"
	// "pandatech.io/dashboard/internal/http/routes"
	// "pandatech.io/dashboard/internal/infrastructure/logger"
	// "pandatech.io/dashboard/internal/infrastructure/validation"
)

var (
	httpCmd = &cobra.Command{
		Use:   "start",
		Short: "Start HTTP REST API",
		Run:   initApp,
	}
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

func initApp(_ *cobra.Command, _ []string) {
	defer logger.Sync()
	logger.Info("starting HTTP Server...")
	gin.SetMode(env.GinMode())
	binding.Validator = validation.NewValidator()
	r := gin.Default()
	r.RedirectFixedPath = true
	routes.RoutesHandler(r)

	srvr := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.AppPort()),
		Handler: r,
	}

	go func() {
		err := srvr.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("failed serving server.", logger.String("error", err.Error()))
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)

	// kill with no param (default) send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL (but we can't catch it, so we don't need to add it)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutdown server...")

	// 5 seconds timeout to waiting signal.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srvr.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown with error.", logger.String("error", err.Error()))
	}

	<-ctx.Done()

	logger.Info("server exiting")
}
