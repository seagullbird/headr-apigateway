package service

//go:generate mockgen -destination=./mock/mock_service.go -package=mock github.com/seagullbird/headr-contentmgr/service Service

import (
	"context"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-errors/errors"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-contentmgr/db"
	repoctlservice "github.com/seagullbird/headr-repoctl/service"
	"strings"
)

// Service describes a service that deals with content management operations (contentmgr).
type Service interface {
	// TODO: Add DeleteAllPosts service for sitemgr, do not need to be exposed by http. Just grpc
	NewPost(ctx context.Context, post db.Post) (uint, error)
	DeletePost(ctx context.Context, id uint) (uint, error)
	GetPost(ctx context.Context, id uint) (db.Post, error)
	PatchPost(ctx context.Context, post db.Post) error
	GetAllPosts(ctx context.Context) ([]uint, error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(repoctlsvc repoctlservice.Service, store db.Store, logger log.Logger) Service {
	var svc Service
	{
		svc = newBasicService(repoctlsvc, store)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	repoctlsvc repoctlservice.Service
	store      db.Store
}

func newBasicService(repoctlsvc repoctlservice.Service, store db.Store) basicService {
	return basicService{
		repoctlsvc: repoctlsvc,
		store:      store,
	}
}

var (
	// ErrPostNotFound indicates a post not found error
	// could be the post does not exist or does not belong to the user
	ErrPostNotFound = errors.New("post not found")
)

func (s basicService) NewPost(ctx context.Context, post db.Post) (uint, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	post.UserID = userID
	id, err := s.store.InsertPost(&post)
	if err != nil {
		return 0, err
	}
	filename := post.Filename + "." + post.Filetype
	filecontent := post.String()
	return id, s.repoctlsvc.WritePost(ctx, post.SiteID, filename, filecontent)
}

func (s basicService) DeletePost(ctx context.Context, id uint) (uint, error) {
	postptr, err := s.store.GetPost(id)
	if err != nil {
		return 0, ErrPostNotFound
	}
	// Post does not belong to the authenticated user
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	if postptr.UserID != userID {
		return 0, ErrPostNotFound
	}
	err = s.store.DeletePost(postptr)
	if err != nil {
		return 0, err
	}
	if err := s.repoctlsvc.RemovePost(ctx, postptr.SiteID, postptr.Filename+"."+postptr.Filetype); err != nil {
		return 0, err
	}
	return id, nil
}

func (s basicService) GetPost(ctx context.Context, id uint) (db.Post, error) {
	postptr, err := s.store.GetPost(id)
	post := *postptr
	if err != nil {
		return post, ErrPostNotFound
	}
	// Post does not belong to the authenticated user
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	if post.UserID != userID {
		return post, ErrPostNotFound
	}
	wholeContent, err := s.repoctlsvc.ReadPost(ctx, post.SiteID, post.Filename+"."+post.Filetype)
	if err != nil {
		return post, err
	}
	content := strings.TrimLeft(strings.Split(wholeContent, "<!--more-->")[1], "\n")
	post.Content = content
	return post, nil
}

func (s basicService) GetAllPosts(ctx context.Context) ([]uint, error) {
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	return s.store.GetAllPosts(userID)
}

func (s basicService) PatchPost(ctx context.Context, post db.Post) error {
	existingPost, err := s.store.GetPost(post.ID)
	if err != nil {
		return err
	}
	userID := ctx.Value(jwt.JWTClaimsContextKey).(stdjwt.MapClaims)["sub"].(string)
	if existingPost == nil || existingPost.UserID != userID {
		return ErrPostNotFound
	}
	// TODO: Add Validation for post
	currentPost, err := s.store.PatchPost(existingPost, &post)
	if err != nil {
		return err
	}
	if post.Content != "" {
		currentPost.Content = post.Content
	}
	filename := currentPost.Filename + "." + currentPost.Filetype
	filecontent := currentPost.String()
	return s.repoctlsvc.WritePost(ctx, currentPost.SiteID, filename, filecontent)
}
