package domain

//Review book model
type Review struct {
	ID        string `json:"id"`
	EntityID  string `json:"entityID"`
	EntityKey string `json:"entityKey"`
	Content   string `json:"content"`
}

//NewReview new review entity
func NewReview(id string, content string, entityID string, entityKey string) *Review {
	return &Review{
		ID:        id,
		Content:   content,
		EntityID:  entityID,
		EntityKey: entityKey,
	}
}
