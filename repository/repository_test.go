package repository

import (
	"testing"

	"github.com/eastrocky/thank-you-notes-service/model"
	"github.com/stretchr/testify/assert"
)

func TestMemoryRepositorySave(t *testing.T) {
	var (
		repo     = NewMemoryRepository()
		thankyou = model.ThankYou{
			From: "Me",
			To:   "You",
			For:  "Great things!",
		}
	)
	repo.Save(thankyou)
	assert.Equal(t, 1, len(repo.ThankYous))
	assert.EqualValues(t, thankyou, repo.ThankYous[0])
}

func TestMemoryRepositoryGet(t *testing.T) {
	var (
		repo    = NewMemoryRepository()
		fromYou = model.ThankYou{
			From: "You",
			To:   "Me",
			For:  "Great things!",
		}
		fromMe = model.ThankYou{
			From: "Me",
			To:   "You",
			For:  "Greater things!",
		}
	)
	repo.ThankYous = []model.ThankYou{
		fromYou,
		fromMe,
	}
	toYou, _ := repo.Get("You")
	assert.Equal(t, 1, len(toYou))
	assert.EqualValues(t, fromMe, toYou[0])
}
