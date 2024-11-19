package role

type Role interface {
	Execute() error
	GracefulShutdown() error
}
