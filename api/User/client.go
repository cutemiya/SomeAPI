package User

import (
	"api/lib/rateclient"
	serviceModel "api/model"
	"api/user/model"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"time"
)

const (
	rateEvery = 1 * time.Second
	rateBurst = 3
)

const (
	host = `https://randomuser.me/api`
)

type Client struct {
	logger *zap.SugaredLogger
	cli    rateclient.RLHTTPClient
}

func NewClient(logger *zap.SugaredLogger, cli *http.Client) Client {
	return Client{
		logger: logger,
		cli:    rateclient.NewClient(cli, rate.NewLimiter(rate.Every(rateEvery), rateBurst)),
	}
}

func (c Client) GetInformation(ctx context.Context) (
	serviceModel.User,
	serviceModel.Location,
	serviceModel.Login,
	serviceModel.Picture,
	error) {

	var data model.User

	err := c.do(ctx, http.MethodGet, host, nil, &data)
	if err != nil {
		return serviceModel.User{},
			serviceModel.Location{},
			serviceModel.Login{},
			serviceModel.Picture{},
			err
	}

	return MapClientToServiceUserModel(data),
		MapClientToServiceLocationModel(data),
		MapClientToServiceLoginModel(data),
		MapClientToServicePictureModel(data),
		nil
}

func (c Client) do(ctx context.Context, method string, url string, body io.Reader, output interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}

	res, err := c.cli.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d; body: %v", res.StatusCode, string(bodyBytes))
	}

	var data map[string]interface{}
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return err
	}

	marshal, err := json.Marshal(data["results"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return err
	}

	if output != nil {
		err = json.Unmarshal(marshal, &output)
		if err != nil {
			return err
		}
	}

	return nil
}
