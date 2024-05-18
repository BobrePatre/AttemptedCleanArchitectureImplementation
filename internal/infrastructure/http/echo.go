package http

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"log/slog"
	"net"
	"net/http"
	"time"
)

var httpServerTag = slog.String("server", "http_server")

const (
	serverPort = "8080"
	serverHost = "0.0.0.0"
)

func NewHttpServer(logger *slog.Logger) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError:      true,
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       true,
		LogRoutePath:     true,
		LogRequestID:     true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				"uri", v.URI,
				"status", v.Status,
				"method", v.Method,
				"host", v.Host,
				"user_agent", v.UserAgent,
				"error", v.Error,
				"latency", v.Latency,
				"request_id", v.RequestID,
				"protocol", v.Protocol,
			)
			return nil
		},
	}))
	return e
}

func RunHttpServer(lc fx.Lifecycle, e *echo.Echo, logger *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("starting on port", httpServerTag)
				if err := e.Start(net.JoinHostPort(serverHost, serverPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("shutting down", httpServerTag)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutting down", httpServerTag)
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			return e.Shutdown(ctx)
		},
	})
}
