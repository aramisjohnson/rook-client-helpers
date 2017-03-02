package clients

import (
	"github.com/aramisjohnson/rook-client-helpers/contracts"
	"errors"
)

type rookBlockClient struct {
	transportClient contracts.ITransportClient
}


func CreateRookBlockClient (client contracts.ITransportClient) *rookBlockClient	{
	client.Connect()

	return &rookBlockClient{transportClient: client};
}


func (r *rookBlockClient) Mount() error	{
	return errors.New("Not Implemented")
}

func (r *rookBlockClient) UnMount() error	{
	return errors.New("Not Implemented")
}

func (r *rookBlockClient) Create() error	{
	return errors.New("Not Implemented")
}

func (r *rookBlockClient) Dispose() error	{
	r.transportClient.Disconnect()
}