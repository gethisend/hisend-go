package hisend

import (
	"fmt"
)

type ThreadsService struct {
	client *Client
}

func (s *ThreadsService) List() ([]Thread, error) {
	var threads []Thread
	err := s.client.request("GET", "/threads", nil, &threads)
	return threads, err
}

func (s *ThreadsService) GetEmails(id int) ([]Email, error) {
	var emails []Email
	err := s.client.request("GET", fmt.Sprintf("/threads/%d/emails", id), nil, &emails)
	return emails, err
}
