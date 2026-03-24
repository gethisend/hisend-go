package hisend

import (
	"fmt"
)

type EmailsService struct {
	client *Client
}

func (s *EmailsService) List() ([]Email, error) {
	var emails []Email
	err := s.client.request("GET", "/emails", nil, &emails)
	return emails, err
}

func (s *EmailsService) Get(id int) (*Email, error) {
	var email Email
	err := s.client.request("GET", fmt.Sprintf("/emails/%d", id), nil, &email)
	return &email, err
}

func (s *EmailsService) Send(data SendEmailRequest) (*Email, error) {
	var email Email
	err := s.client.request("POST", "/emails", data, &email)
	return &email, err
}

func (s *EmailsService) SendBatch(data SendEmailBatchRequest) (*SendEmailBatchResponse, error) {
	var res SendEmailBatchResponse
	err := s.client.request("POST", "/emails/batch", data, &res)
	return &res, err
}
