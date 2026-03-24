package hisend

import (
	"fmt"
)

type RoutingService struct {
	client *Client
}

func (s *RoutingService) List(domainID int) ([]Routing, error) {
	var routing []Routing
	err := s.client.request("GET", fmt.Sprintf("/domains/%d/routing", domainID), nil, &routing)
	return routing, err
}

func (s *RoutingService) Create(domainID int, data CreateRoutingRequest) (*Routing, error) {
	var routing Routing
	err := s.client.request("POST", fmt.Sprintf("/domains/%d/routing", domainID), data, &routing)
	return &routing, err
}

func (s *RoutingService) Update(domainID, id int, data UpdateRoutingRequest) (*Routing, error) {
	var routing Routing
	err := s.client.request("PUT", fmt.Sprintf("/domains/%d/routing/%d", domainID, id), data, &routing)
	return &routing, err
}

func (s *RoutingService) Get(domainID, id int) (*Routing, error) {
	var routing Routing
	err := s.client.request("GET", fmt.Sprintf("/domains/%d/routing/%d", domainID, id), nil, &routing)
	return &routing, err
}

func (s *RoutingService) Delete(domainID, id int) error {
	return s.client.request("DELETE", fmt.Sprintf("/domains/%d/routing/%d", domainID, id), nil, nil)
}
