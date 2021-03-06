package crunchyroll

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("crunchyroll", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Crunchyroll.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.crunchyroll.com/search?q=%s", url.QueryEscape(q))
}
