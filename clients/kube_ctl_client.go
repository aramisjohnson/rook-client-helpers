package clients


type kubeCtlClient struct{

}

func CreateNewKubeCtlClient() *kubeCtlClient	{
	return &kubeCtlClient{}
}

func (k *kubeCtlClient) Connect() error	{
	return nil
}

func (k *kubeCtlClient)  Disconnect() error	{
	return nil
}

func (k *kubeCtlClient) Create () error {
	return nil
}

func (k *kubeCtlClient) Delete() error {
	return nil
}

func (k *kubeCtlClient) GetPods() error	{
	return nil
}

func (k *kubeCtlClient) GetPodStatus() error	{
	return nil
}
