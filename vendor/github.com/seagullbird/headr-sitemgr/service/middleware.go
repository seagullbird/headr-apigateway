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

func (mw loggingMiddleware) NewSite(ctx context.Context, sitename string) (id uint, err error) {
	id, err = mw.next.NewSite(ctx, sitename)
	mw.logger.Log("method", "NewSite", "sitename", sitename, "siteID", id, "err", err)
	return
}

func (mw loggingMiddleware) DeleteSite(ctx context.Context, siteID uint) (err error) {
	err = mw.next.DeleteSite(ctx, siteID)
	mw.logger.Log("method", "DeleteSite", "siteID", siteID, "err", err)
	return
}

func (mw loggingMiddleware) CheckSitenameExists(ctx context.Context, sitename string) (bool, error) {
	exists, err := mw.next.CheckSitenameExists(ctx, sitename)
	mw.logger.Log("method", "CheckSitenameExists", "sitename", sitename, "exists", exists, "err", err)
	return exists, err
}

func (mw loggingMiddleware) GetSiteIDByUserID(ctx context.Context) (uint, error) {
	siteID, err := mw.next.GetSiteIDByUserID(ctx)
	mw.logger.Log("method", "GetSiteIDByUserID", "siteID", siteID, "err", err)
	return siteID, err
}

func (mw loggingMiddleware) GetConfig(ctx context.Context, siteID uint) (string, error) {
	config, err := mw.next.GetConfig(ctx, siteID)
	mw.logger.Log("method", "GetConfig", "siteID", siteID)
	return config, err
}

func (mw loggingMiddleware) UpdateConfig(ctx context.Context, siteID uint, config string) error {
	err := mw.next.UpdateConfig(ctx, siteID, config)
	mw.logger.Log("method", "UpdateConfig", "siteID", siteID, "err", err)
	return err
}

func (mw loggingMiddleware) GetThemes(ctx context.Context, siteID uint) (string, error) {
	s, err := mw.next.GetThemes(ctx, siteID)
	mw.logger.Log("method", "GetThemes", "siteID", siteID, "err", err)
	return s, err
}

func (mw loggingMiddleware) UpdateSiteTheme(ctx context.Context, siteID uint, theme string) error {
	err := mw.next.UpdateSiteTheme(ctx, siteID, theme)
	mw.logger.Log("method", "UpdateSiteTheme", "siteID", siteID, "theme", theme, "err", err)
	return err
}

func (mw loggingMiddleware) PostAbout(ctx context.Context, siteID uint, content string) error {
	err := mw.next.PostAbout(ctx, siteID, content)
	mw.logger.Log("method", "PostAbout", "siteID", siteID, "err", err)
	return err
}

func (mw loggingMiddleware) GetAbout(ctx context.Context, siteID uint) (string, error) {
	content, err := mw.next.GetAbout(ctx, siteID)
	mw.logger.Log("method", "GetAbout", "siteID", siteID, "err", err)
	return content, err
}
