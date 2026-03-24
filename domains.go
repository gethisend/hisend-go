package hisend

import (
	"fmt"
)

type DomainsService struct {
	client *Client
}

func (s *DomainsService) List() ([]Domain, error) {
	var domains []Domain
	err := s.client.request("GET", "/domains", nil, &domains)
	return domains, err
}

func (s *DomainsService) Get(id int) (*Domain, error) {
	var domain Domain
	err := s.client.request("GET", fmt.Sprintf("/domains/%d", id), nil, &domain)
	return &domain, err
}

func (s *DomainsService) Verify(id int) (*Domain, error) {
	var domain Domain
	err := s.client.request("GET", fmt.Sprintf("/domains/%d/verify", id), nil, &domain)
	return &domain, err
}

func (s *DomainsService) Add(data AddDomainRequest) (*Domain, error) {
	var domain Domain
	err := s.client.request("POST", "/domains", data, &domain)
	return &domain, err
}

func (s *DomainsService) Delete(id int) error {
	return s.client.request("DELETE", fmt.Sprintf("/domains/%d", id), nil, nil)
}
