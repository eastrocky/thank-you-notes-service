package repository

import "github.com/eastrocky/thank-you-notes-service/model"

// Repository defines methods for storing thank you notes
type Repository interface {
	Save(thankyou model.ThankYou) error
	Get(to string) ([]model.ThankYou, error)
}

// MemoryRepository is an in-memory repository for storing and fetching thank you notes
type MemoryRepository struct {
	ThankYous []model.ThankYou
}

// Save stores a thank you note
func (r *MemoryRepository) Save(thankyou model.ThankYou) error {
	r.ThankYous = append(r.ThankYous, thankyou)
	return nil
}

// Get gets a list of thank you notes
func (r *MemoryRepository) Get(to string) ([]model.ThankYou, error) {
	thanksTo := make([]model.ThankYou, 0, len(r.ThankYous))
	for _, thanks := range r.ThankYous {
		if thanks.To == to {
			thanksTo = append(thanksTo, thanks)
		}
	}
	return thanksTo, nil
}

// NewMemoryRepository initializes a new in-memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		ThankYous: []model.ThankYou{},
	}
}
