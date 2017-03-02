package clients

import "errors"

type kubeCtlTransportClient struct	{

}


func CreateNewKubeCtlTransportClient() *kubeCtlTransportClient	{
	return &kubeCtlTransportClient{};
}

func (k *kubeCtlTransportClient) Connect() error	{
	return errors.New("Not Implemented")
}

func (k *kubeCtlTransportClient) Disconnect() error	{
	return errors.New("Not Implemented")
}

func (k *kubeCtlTransportClient)Execute(cmd string) error	{
	return errors.New("Not Implemented")
}
