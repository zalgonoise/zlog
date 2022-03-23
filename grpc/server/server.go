package server

import (
	"errors"
	"net"

	"github.com/zalgonoise/zlog/config"
	"github.com/zalgonoise/zlog/log"
	pb "github.com/zalgonoise/zlog/proto/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ErrMessageParse error = errors.New("failed to parse message")
	ErrAddrListen   error = errors.New("failed to listen to input address")
)

type GRPCLogServer struct {
	Addr      string
	opts      []grpc.ServerOption
	Logger    log.Logger
	SvcLogger log.Logger
	ErrCh     chan error
	LogSv     *pb.LogServer
	Server    *grpc.Server
}

type GRPCLogServerBuilder struct {
	conf *config.Configs
}

func newGRPCLogServer(confs ...config.Config) *GRPCLogServerBuilder {
	// enforce defaults on builder
	builder := &GRPCLogServerBuilder{
		conf: gRPCLogServerDefaults(),
	}

	// apply input configs
	for _, config := range confs {
		if config.Is(gRPCLogServerBuilderType) {
			config.Apply(builder.conf)
		}
	}

	return builder
}

// factory
func New(confs ...config.Config) *GRPCLogServer {
	builder := newGRPCLogServer(confs...)

	// grpc options (already initialized)
	opt := builder.conf.Get(LSOpts.String()).([]grpc.ServerOption)
	tls := builder.conf.Get(LSTLS.String()).([]grpc.ServerOption)

	var grpcOpts []grpc.ServerOption
	grpcOpts = append(grpcOpts, opt...)
	grpcOpts = append(grpcOpts, tls...)

	server := &GRPCLogServer{
		Addr:      builder.conf.Get(LSAddress.String()).(string),
		opts:      grpcOpts,
		Logger:    builder.conf.Get(LSLogger.String()).(log.Logger),
		SvcLogger: builder.conf.Get(LSSvcLogger.String()).(log.Logger),
		ErrCh:     make(chan error),
		LogSv:     pb.NewLogServer(),
	}

	go server.registerComms()

	return server

}

func (s GRPCLogServer) registerComms() {
	for {
		msg, ok := <-s.LogSv.Comm
		if !ok {
			s.SvcLogger.Log(log.NewMessage().Level(log.LLWarn).Prefix("gRPC").Sub("LogServer.Comm").Message("couldn't parse message from LogServer").Metadata(log.Field{"error": ErrMessageParse.Error()}).Build())
			continue
		}

		s.SvcLogger.Log(log.NewMessage().FromProto(msg).Build())
	}
}

func (s GRPCLogServer) listen() net.Listener {
	lis, err := net.Listen("tcp", s.Addr)

	if err != nil {
		s.ErrCh <- err

		s.SvcLogger.Log(log.NewMessage().Level(log.LLFatal).Prefix("gRPC").Sub("listen").
			Message("couldn't listen to input address").Metadata(log.Field{
			"error": err.Error(),
			"addr":  s.Addr,
		}).Build())

		return nil
	}

	s.SvcLogger.Log(log.NewMessage().Level(log.LLInfo).Prefix("gRPC").Sub("listen").
		Message("gRPC server is listening to connections").Metadata(log.Field{
		"addr": s.Addr,
	}).Build())

	return lis
}

func (s GRPCLogServer) handleMessages() {
	s.SvcLogger.Log(log.NewMessage().Level(log.LLDebug).Prefix("gRPC").Sub("handler").Message("message handler is running").Build())

	for {
		select {
		case msg := <-s.LogSv.MsgCh:
			logmsg := log.NewMessage().FromProto(msg).Build()
			s.Logger.Log(logmsg)

			s.SvcLogger.Log(log.NewMessage().Level(log.LLDebug).Prefix("gRPC").Sub("handler").Message("input log message parsed and registered").Build())

		case <-s.LogSv.Done:
			s.SvcLogger.Log(log.NewMessage().Level(log.LLDebug).Prefix("gRPC").Sub("handler").Message("received done signal").Build())
			return
		}
	}
}

func (s GRPCLogServer) Serve() {
	lis := s.listen()
	if lis == nil {
		return
	}

	go s.handleMessages()

	s.Server = grpc.NewServer(s.opts...)
	pb.RegisterLogServiceServer(s.Server, s.LogSv)

	// gRPC reflection
	reflection.Register(s.Server)

	s.SvcLogger.Log(log.NewMessage().Level(log.LLDebug).Prefix("gRPC").Sub("serve").
		Message("gRPC server is running").Metadata(log.Field{
		"addr": s.Addr,
	}).Build())

	if err := s.Server.Serve(lis); err != nil {
		s.ErrCh <- err

		s.SvcLogger.Log(log.NewMessage().Level(log.LLDebug).Prefix("gRPC").Sub("serve").
			Message("gRPC server crashed with an error").Metadata(log.Field{
			"error": err.Error(),
			"addr":  s.Addr,
		}).Build())
		return
	}

}

func (s GRPCLogServer) Stop() {
	s.LogSv.Done <- struct{}{}
	s.Server.Stop()

	s.SvcLogger.Log(log.NewMessage().Level(log.LLDebug).Prefix("gRPC").Sub("stop").Message("received done signal").Build())
}
