package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eastrocky/thank-you-notes-service/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestThanksGetOK(t *testing.T) {
	var (
		router   = setupRouter()
		expected = []gin.H{}
		actual   []gin.H
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/thanks/You", nil)
	router.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
}

func TestThanksPostOKAnonymous(t *testing.T) {
	var (
		router      = setupRouter()
		thankyou, _ = json.Marshal(model.ThankYou{
			From: "Anonymous",
			To:   "Jimmy Fallon",
			For:  "Thank you for all the 'Thank You Notes'",
		})
		expected = gin.H{
			"from": "Anonymous",
			"to":   "Jimmy Fallon",
			"for":  "Thank you for all the 'Thank You Notes'",
		}
		actual gin.H
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/thanks", bytes.NewReader(thankyou))
	router.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expected, actual)
}

func TestThanksPostBadRequest(t *testing.T) {
	var (
		router      = setupRouter()
		thankyou, _ = json.Marshal(model.ThankYou{
			From: "",
			To:   "",
			For:  "Thank you for all the 'Thank You Notes'",
		})
		expected = gin.H{
			"code":    http.StatusText(http.StatusBadRequest),
			"message": "Key: 'ThankYou.From' Error:Field validation for 'From' failed on the 'required' tag\nKey: 'ThankYou.To' Error:Field validation for 'To' failed on the 'required' tag",
		}
		actual gin.H
	)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/thanks", bytes.NewReader(thankyou))
	router.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}
