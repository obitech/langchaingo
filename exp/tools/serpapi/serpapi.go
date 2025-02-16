package serpapi

import (
	"errors"
	"os"
	"strings"

	"github.com/tmc/langchaingo/exp/tools"
	"github.com/tmc/langchaingo/exp/tools/serpapi/internal"
)

var ErrMissingToken = errors.New("missing the serpapi API key, set it in the SERPAPI_API_KEY environment variable")

type Tool struct {
	client *internal.SerpapiClient
}

var _ tools.Tool = Tool{}

// New creates a new serpapi tool to search on internet.
func New() (*Tool, error) {
	apiKey := os.Getenv("SERPAPI_API_KEY")
	if apiKey == "" {
		return nil, ErrMissingToken
	}

	return &Tool{
		client: internal.New(apiKey),
	}, nil
}

func (t Tool) Name() string {
	return "Google Search"
}

func (t Tool) Description() string {
	return `
	"A wrapper around Google Search. "
	"Useful for when you need to answer questions about current events. "
	"Always one of the first options when you need to find information on internet"
	"Input should be a search query."`
}

func (t Tool) Call(input string) (string, error) {
	result, err := t.client.Search(input)
	if err != nil {
		if errors.Is(err, internal.ErrNoGoodResult) {
			return "No good Google Search Result was found", nil
		}

		return "", err
	}

	return strings.Join(strings.Fields(result), " "), nil
}
