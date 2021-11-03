package elasticsearch

import (
	"context"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

type Config struct {
	URLs string `envconfig:"ES_CONNECTION_URLS" mapstructure:"es_connection_urls" required:"true"`
}

func InitElasticSearchClient(conf Config) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(conf.URLs),
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create elasticsearch client")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	_, _, err = client.Ping(conf.URLs).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to ping elasticsearch")
	}
	return client, err
}
