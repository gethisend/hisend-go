package hisend

type EmailAddress struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailAttachment struct {
	Filename    string `json:"filename"`
	S3Key       string `json:"s3_key,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size,omitempty"`
}

type Email struct {
	ID            int               `json:"id"`
	ProjectID     int               `json:"project_id"`
	MessageID     string            `json:"message_id"`
	From          string            `json:"from"`
	To            string            `json:"to"`
	SenderDetails EmailAddress      `json:"sender_details"`
	ToDetails     []EmailAddress    `json:"to_details"`
	Cc            []EmailAddress    `json:"cc"`
	Bcc           []EmailAddress    `json:"bcc"`
	Attachments   []EmailAttachment `json:"attachments"`
	Subject       string            `json:"subject"`
	Content       string            `json:"content"`
	HTMLBody      string            `json:"html_body"`
	TextBody      string            `json:"text_body"`
	Status        string            `json:"status"`
	Direction     string            `json:"direction"`
	ThreadID      *int              `json:"thread_id,omitempty"`
	InReplyTo     *string           `json:"in_reply_to,omitempty"`
	References    *string           `json:"references,omitempty"`
	CreatedAt     string            `json:"created_at"`
}

type Thread struct {
	ID            int    `json:"id"`
	ProjectID     int    `json:"project_id"`
	Subject       string `json:"subject"`
	MessageCount  int    `json:"message_count"`
	LatestSnippet string `json:"latest_snippet"`
	LastMessageAt string `json:"last_message_at"`
}

type DNSRecord struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Domain struct {
	ID                 int         `json:"id"`
	ProjectID          int         `json:"project_id"`
	Name               string      `json:"name"`
	VerificationStatus string      `json:"verification_status"`
	DKIMRecords        string      `json:"dkim_records"`
	DNSRecords         []DNSRecord `json:"dns_records,omitempty"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
}

type Endpoint struct {
	ID              int      `json:"id"`
	ProjectID       int      `json:"project_id"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	URL             *string  `json:"url,omitempty"`
	Email           *string  `json:"email,omitempty"`
	EmailList       []string `json:"email_list,omitempty"`
	TriggerCount    int      `json:"trigger_count"`
	LastTriggeredAt string   `json:"last_triggered_at"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

type Routing struct {
	ID           int        `json:"id"`
	DomainID     int        `json:"domain_id"`
	Type         string     `json:"type"`
	EmailAddress *string    `json:"email_address,omitempty"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	Endpoints    []Endpoint `json:"endpoints,omitempty"`
}

type AddDomainRequest struct {
	Name string `json:"name"`
}

type AttachmentReq struct {
	Filename    string `json:"filename"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}

type SendEmailRequest struct {
	From        string          `json:"from"`
	To          interface{}     `json:"to"` // Can be string or []string based on TS typing but handling in Go can be done with interface{} or specific structure
	Cc          []string        `json:"cc,omitempty"`
	Bcc         []string        `json:"bcc,omitempty"`
	Subject     *string         `json:"subject,omitempty"`
	HTML        *string         `json:"html,omitempty"`
	Text        *string         `json:"text,omitempty"`
	Attachments []AttachmentReq `json:"attachments,omitempty"`
	InReplyTo   *string         `json:"in_reply_to,omitempty"`
}

type SendEmailBatchRequest []SendEmailRequest

type SendEmailBatchResponse struct {
	Results []interface{} `json:"results"`
}

type CreateRoutingRequest struct {
	Type         string   `json:"type"`
	EmailAddress *string  `json:"email_address,omitempty"`
	EndpointIDs  []int    `json:"endpoint_ids"`
}

type UpdateRoutingRequest struct {
	EmailAddress *string `json:"email_address,omitempty"`
	EndpointIDs  []int   `json:"endpoint_ids,omitempty"`
}
