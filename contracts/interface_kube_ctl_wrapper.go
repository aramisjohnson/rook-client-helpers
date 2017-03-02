package contracts


type IKubeCtlClient interface {
	Create () error
	Delete() error
	GetPods() error
	GetPodStatus() error
}



