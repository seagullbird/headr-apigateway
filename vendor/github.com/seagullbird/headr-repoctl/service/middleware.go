package service

import (
	"context"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			logger,
			next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) NewSite(ctx context.Context, siteID uint, theme string) (err error) {
	err = mw.next.NewSite(ctx, siteID, theme)
	mw.logger.Log("method", "NewSite", "siteID", siteID, "err", err)
	return
}

func (mw loggingMiddleware) DeleteSite(ctx context.Context, siteID uint) (err error) {
	err = mw.next.DeleteSite(ctx, siteID)
	mw.logger.Log("method", "DeleteSite", "siteID", siteID, "err", err)
	return
}

func (mw loggingMiddleware) WritePost(ctx context.Context, siteID uint, filename, content string) (err error) {
	err = mw.next.WritePost(ctx, siteID, filename, content)
	mw.logger.Log("method", "WritePost", "siteID", siteID, "filename", filename, "err", err)
	return
}

func (mw loggingMiddleware) RemovePost(ctx context.Context, siteID uint, filename string) (err error) {
	err = mw.next.RemovePost(ctx, siteID, filename)
	mw.logger.Log("method", "RemovePost", "siteID", siteID, "filename", filename, "err", err)
	return
}

func (mw loggingMiddleware) ReadPost(ctx context.Context, siteID uint, filename string) (content string, err error) {
	content, err = mw.next.ReadPost(ctx, siteID, filename)
	mw.logger.Log("method", "ReadPost", "siteID", siteID, "filename", filename, "err", err)
	return
}

func (mw loggingMiddleware) WriteConfig(ctx context.Context, siteID uint, config string) error {
	err := mw.next.WriteConfig(ctx, siteID, config)
	mw.logger.Log("method", "WriteConfig", "siteID", siteID, "err", err)
	return err
}

func (mw loggingMiddleware) ReadConfig(ctx context.Context, siteID uint) (string, error) {
	config, err := mw.next.ReadConfig(ctx, siteID)
	mw.logger.Log("method", "ReadConfig", "siteID", siteID, "err", err)
	return config, err
}

func (mw loggingMiddleware) UpdateAbout(ctx context.Context, siteID uint, content string) error {
	err := mw.next.UpdateAbout(ctx, siteID, content)
	mw.logger.Log("method", "UpdateAbout", "siteID", siteID, "err", err)
	return err
}

func (mw loggingMiddleware) ReadAbout(ctx context.Context, siteID uint) (string, error) {
	content, err := mw.next.ReadAbout(ctx, siteID)
	mw.logger.Log("method", "ReadAbout", "siteID", siteID, "err", err)
	return content, err
}

func (mw loggingMiddleware) ChangeDefaultConfig(ctx context.Context, siteID uint, theme string) error {
	err := mw.next.ChangeDefaultConfig(ctx, siteID, theme)
	mw.logger.Log("method", "ChangeDefaultConfig", "siteID", siteID, "theme", theme, "err", err)
	return err
}
