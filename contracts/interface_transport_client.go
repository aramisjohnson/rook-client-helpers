package contracts


type ITransportClient interface {
	Connect() error
	Disconnect() error
	Execute(cmd string) error
}
