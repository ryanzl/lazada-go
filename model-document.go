package lazada

type GetDocumentRequest struct {
	DocType                string   `json:"doc_type,omitempty"`
	OrderItemIds   string   `json:"order_item_ids,omitempty"`
}

type GetDocumentResponse struct {
	Code                 string   `json:"code,omitempty"`
	Data                 *GetDocumentData  `json:"data,omitempty"`
	RequestId            string   `json:"request_id,omitempty"`
}
type GetDocumentData struct {
	Document                GetDocument   `json:"document,omitempty"`
}
type GetDocument struct {
	File                string   `json:"file,omitempty"`
	MimeType   string   `json:"mime_type,omitempty"`
}