package crowi

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/context"
)

// SearchService handles communication with the Search related
// methods of the Crowi API.
type SearchService service

func (s *SearchService) List(ctx context.Context, path string, query string, opt *SearchListOptions) (*Pages, error) {
	var pages Pages

	params := url.Values{}
	params.Set("access_token", s.client.config.Token)
	params.Set("tree", path)
	params.Set("q", query)

	err := s.client.newRequest(ctx, http.MethodGet, "/_api/search", params, &pages)
	if err != nil {
		return nil, err
	}

	if opt != nil && opt.ListOptions.Pagenation {
		offset := 0
		var p []PageInfo
		for {
			params.Set("offset", fmt.Sprintf("%d", offset))
			err := s.client.newRequest(ctx, http.MethodGet, "/_api/search", params, &pages)
			if err != nil {
				break
			}

			p = append(p, pages.Pages...)
			offset += 50
		}
		pages.Pages = p
	}
	return &pages, nil
}

// SearchListOptions specifies the optional parameters to the
// SearchService.List methods.
type SearchListOptions struct {
	ListOptions
}
