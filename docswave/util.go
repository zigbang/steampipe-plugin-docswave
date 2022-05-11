package docswave

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*http.Client, string, error) {
	token := os.Getenv("DOCSWAVE_TOKEN")

	docswaveConfig := GetConfig(d.Connection)
	if docswaveConfig.Token != nil {
		token = *docswaveConfig.Token
	}

	if token == "" {
		return nil, "", errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := &http.Client{}

	return client, token, nil
}
