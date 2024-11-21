package common

import (
	"io"
	"log"

	"google.golang.org/grpc"
)

func ExecuteOnceWithBidiClient[Req any, Res any](stream grpc.BidiStreamingClient[Req, Res], req *Req) (res *Res, err error) {
	err = stream.Send(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := stream.CloseSend(); err != nil {
			log.Printf("failed to close send: %v", err)
		}
	}()

	in, err := stream.Recv()
	if err != nil {
		return nil, err
	}

	return in, nil
}

func ListenBidiServer[Req any, Res any](stream grpc.BidiStreamingServer[Req, Res], f func(req *Req) error) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err = f(req); err != nil {
			return err
		}
	}

	return nil
}

func ListenBidiClient[Req any, Res any](stream grpc.BidiStreamingClient[Req, Res], f func(res *Res) error) error {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err = f(res); err != nil {
			return err
		}
	}

	return nil
}
