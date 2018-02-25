package transport

import (
	"context"
	"github.com/go-errors/errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	"github.com/seagullbird/headr-contentmgr/pb"
	"github.com/seagullbird/headr-contentmgr/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	newpost grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) pb.ContentmgrServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		newpost: grpctransport.NewServer(
			endpoints.NewPostEndpoint,
			decodeGRPCNewPostRequest,
			encodeGRPCNewPostResponse,
			options...,
		),
	}
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
	var newpostEndpoint kitendpoint.Endpoint
	{
		newpostEndpoint = grpctransport.NewClient(
			conn,
			"pb.Contentmgr",
			"NewPost",
			encodeGRPCNewPostRequest,
			decodeGRPCNewPostResponse,
			pb.CreateNewPostReply{},
		).Endpoint()
	}

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return endpoint.Set{
		NewPostEndpoint: newpostEndpoint,
	}
}

func (s *grpcServer) NewPost(ctx context.Context, req *pb.CreateNewPostRequest) (*pb.CreateNewPostReply, error) {
	_, rep, err := s.newpost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateNewPostReply), nil
}

func encodeGRPCNewPostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.NewPostRequest)
	return &pb.CreateNewPostRequest{
		Title:    req.Title,
		Summary:  req.Summary,
		Content:  req.Content,
		Tags:     req.Tags,
		Author:   req.Author,
		Sitename: req.Sitename,
		Date:     req.Date,
	}, nil
}

func decodeGRPCNewPostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateNewPostRequest)
	return endpoint.NewPostRequest{
		Title:    req.Title,
		Summary:  req.Summary,
		Content:  req.Content,
		Tags:     req.Tags,
		Author:   req.Author,
		Sitename: req.Sitename,
		Date:     req.Date,
	}, nil
}

func encodeGRPCNewPostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.NewPostResponse)
	return &pb.CreateNewPostReply{
		Err: err2str(resp.Err),
	}, nil
}

func decodeGRPCNewPostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateNewPostReply)
	return endpoint.NewPostResponse{Err: str2err(reply.Err)}, nil
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}
