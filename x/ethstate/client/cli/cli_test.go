package cli

import (
	"context"
	"fmt"
	"testing"
	"time"

	//ethquerymoduletypes "github.com/ajansari95/cosmicether/x/ethquery/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	libclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
)

func newRPCClient(addr string, timeout time.Duration) (*rpchttp.HTTP, error) {
	httpClient, err := libclient.DefaultHTTPClient(addr)
	if err != nil {
		return nil, err
	}

	httpClient.Timeout = timeout
	rpcClient, err := rpchttp.NewWithClient(addr, "/websocket", httpClient)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}

func TestLogicForRewardMap(t *testing.T) {

	rpcClientC, _ := newRPCClient("http://localhost:26657", 1*time.Second)
	//height := int64(3657)

	//req := &ethquerymoduletypes.QueryRequestsRequest{}

	res, err := rpcClientC.ABCIQuery(context.Background(), "/cosmicether.ethquery.Query/Queries", []byte(""))
	fmt.Println(res.Response.Value)

	fmt.Println(err)

}
