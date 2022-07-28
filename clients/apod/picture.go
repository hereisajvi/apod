package apod

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Picture struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDURL          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func (c Client) GetPictureOfTheDay(ctx context.Context, date string) (Picture, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s?date=%s&api_key=%s", baseURL, date, c.apiKey), http.NoBody)
	if err != nil {
		return Picture{}, errors.Wrap(err, "could not build http request")
	}

	response, err := c.client.Do(request)
	if err != nil {
		return Picture{}, errors.Wrap(err, "could not send http request")
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var picture Picture

	err = decoder.Decode(&picture)
	if err != nil {
		return Picture{}, errors.Wrap(err, "could not decode response body")
	}

	return picture, nil
}
