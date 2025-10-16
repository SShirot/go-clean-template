// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/evrone/go-clean-template/config"
	grpcController "github.com/evrone/go-clean-template/internal/controller/grpc"
	"github.com/evrone/go-clean-template/internal/controller/http"
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/internal/di/wire"
	domainTranslation "github.com/evrone/go-clean-template/internal/domain/translation"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

// translationAdapter adapts di.TranslationServiceInterface to usecase.Translation
type translationAdapter struct {
	svc di.TranslationServiceInterface
}

func (a translationAdapter) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
	// Adapt to application service request type
	req := &domainTranslation.TranslateRequest{
		Source:      t.Source,
		Destination: t.Destination,
		Original:    t.Original,
	}

	res, err := a.svc.Translate(ctx, req)
	if err != nil {
		return entity.Translation{}, err
	}
	// Handle different possible response shapes
	if r, ok := res.(*domainTranslation.TranslateResponse); ok {
		return entity.Translation{
			Source:      t.Source,
			Destination: t.Destination,
			Original:    t.Original,
			Translation: r.Translation,
		}, nil
	}
	if tr, ok := res.(entity.Translation); ok {
		return tr, nil
	}
	return entity.Translation{}, fmt.Errorf("invalid translation response type")
}

func (a translationAdapter) History(ctx context.Context) (entity.TranslationHistory, error) {
	res, err := a.svc.GetHistory(ctx, 100, 0)
	if err != nil {
		return entity.TranslationHistory{}, err
	}
	// Handle application-layer type
	if h, ok := res.(*domainTranslation.TranslationHistory); ok {
		items := make([]entity.Translation, 0, len(h.History))
		for _, d := range h.History {
			items = append(items, entity.Translation{
				Source:      d.Source,
				Destination: d.Destination,
				Original:    d.Original,
				Translation: d.Translation,
			})
		}
		return entity.TranslationHistory{History: items}, nil
	}
	// Already entity type
	if h, ok := res.(entity.TranslationHistory); ok {
		return h, nil
	}
	// Slice of entity translations
	if list, ok := res.([]entity.Translation); ok {
		return entity.TranslationHistory{History: list}, nil
	}
	return entity.TranslationHistory{}, fmt.Errorf("invalid history type")
}

// Run creates objects via constructors using Wire dependency injection.
func Run(cfg *config.Config) {
	// Initialize app with all dependencies using Wire
	app, err := wire.InitializeApp(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to initialize app: %w", err))
	}

	// Ensure postgres connection is closed on exit
	if app.Postgres != nil {
		defer app.Postgres.Close()
	}

	// Setup routers
	if app.GRPCServer != nil && app.HTTPServer != nil && app.Logger != nil && app.TranslationUC != nil {
		grpcServer := app.GRPCServer.GetServer().(*grpc.Server)
		httpApp := app.HTTPServer.GetApp().(*fiber.App)
		logger := app.Logger.(logger.Interface)
		translationUC := usecase.Translation(translationAdapter{svc: app.TranslationUC})

		grpcController.NewRouter(grpcServer, translationUC, logger)
		http.NewRouter(httpApp, cfg, translationUC, logger)

		// Start servers
		app.GRPCServer.Start()
		app.HTTPServer.Start()
	}

	// Log successful initialization
	if app.Logger != nil {
		app.Logger.Info("app - Run - application initialized successfully")
		app.Logger.Info("app - Run - Swagger available at: http://localhost:8080/swagger/")
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		app.Logger.Info("app - Run - signal: %s", s.String())
	}

	// Shutdown
	if app.HTTPServer != nil {
		app.HTTPServer.Stop()
	}
	if app.GRPCServer != nil {
		app.GRPCServer.Stop()
	}
}
