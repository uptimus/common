package application

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/uptimus/common/logging"
)

type Shutdown = func(handler io.Closer)
type Launcher = func(ctx context.Context, addShutdownHook Shutdown) error

type Application struct {
	closable []io.Closer
	launcher Launcher
	done     chan bool
	quit     chan os.Signal
}

func (a *Application) Launch(ctx context.Context) {
	logging.Info("Application starting")

	go a.gracefulShutdown(ctx)

	err := a.launcher(ctx, a.registerDeferredShutdown)
	if err != nil {
		logging.Fatal("Application couldn't start", zap.Error(err))
		a.quit <- syscall.SIGTERM
	}
	<-a.done

	logging.Info("Application shutdown")
}

func (a *Application) gracefulShutdown(ctx context.Context) {
	<-a.quit

	for _, closer := range a.closable {
		_ = closer.Close()
	}

	close(a.done)
}

func (a *Application) registerDeferredShutdown(handler io.Closer) {
	a.closable = append(a.closable, handler)
}

func NewApplication(launcher Launcher) *Application {
	closable := make([]io.Closer, 0)
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	return &Application{
		closable: closable,
		launcher: launcher,
		done:     done,
		quit:     quit,
	}
}
