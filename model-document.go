package lazada

type GetDocumentRequest struct {
	DocType                string   `json:"doc_type,omitempty"`
	OrderItemIds   string   `json:"order_item_ids,omitempty"`
}

type GetDocumentResponse struct {
	Code                 string   `json:"code,omitempty"`
	Data                 *GetDocument  `json:"data,omitempty"`
	RequestId            string   `json:"request_id,omitempty"`
}

type GetDocument struct {
	File                string   `json:"fileC,omitempty"`
	MimeType   string   `json:"mime_type,omitempty"`
}