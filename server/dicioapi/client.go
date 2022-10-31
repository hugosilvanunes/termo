package dicioapi

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hugosilvanunes/termo/config"
	"go.uber.org/zap"
)

type Client struct {
	cli      *resty.Client
	log      *zap.Logger
	dicioURL string
}

func NewClient(cfg config.Config) *Client {
	client := resty.New()
	client.SetDebug(cfg.Env.ENV == "dev")
	client.SetTimeout(5 * time.Second)

	return &Client{cli: client, dicioURL: cfg.Env.DicioURL}
}

func (c Client) GetRandomWord() (*DicioResponse, error) {
	resp, err := c.cli.R().SetResult(&DicioResponse{}).Get(c.dicioURL)
	if err != nil {
		c.log.Error("cannnot request to dicio api", zap.Error(err))
		return nil, err
	}

	resBody, ok := resp.Result().(*DicioResponse)
	if !ok {
		c.log.Error("cannot parse response body")
		return nil, errors.New("cannot parse response body")
	}

	return resBody, nil
}
