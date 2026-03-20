package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"interview/order/model"
	"interview/order/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStore struct {
	saveFunc func(model.Order) error
}

func (m *MockStore) SaveOrder(ctx context.Context, o model.Order) error {
	return m.saveFunc(o)
}

func (m *MockStore) GetOrder(ctx context.Context, id string) (model.Order, error) {
	return model.Order{ID: "123", Item: "bottle", Amount: 2000}, nil
}

func TestHandleOrder(t *testing.T) {
	tests := []struct {
		name           string
		payload        model.Order
		mockErr        error
		expectedStatus int
		endpoint       string
		method         string
		function       func(store.Repository) http.HandlerFunc
		pathID         string
	}{
		{

			name:           "Success Case",
			payload:        model.Order{ID: "123", Item: "Laptop", Amount: 1500},
			mockErr:        nil,
			expectedStatus: http.StatusAccepted,
			endpoint:       "/order",
			method:         http.MethodPost,
			function:       HandleOrder,
		},
		{
			name:           "Database Failure",
			payload:        model.Order{ID: "500", Item: "Error", Amount: 0},
			mockErr:        fmt.Errorf("db connection lost"),
			expectedStatus: http.StatusInternalServerError,
			endpoint:       "/order",
			method:         http.MethodPost,
			function:       HandleOrder,
		},
		{
			name:           "GET Success",
			payload:        model.Order{ID: "500", Item: "Error", Amount: 0},
			mockErr:        nil,
			expectedStatus: http.StatusOK,
			endpoint:       "/order/get/123",
			method:         http.MethodGet,
			function:       GetOrder,
			pathID:         "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStore{
				saveFunc: func(o model.Order) error { return tt.mockErr },
			}

			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(tt.method, tt.endpoint, bytes.NewBuffer(body))

			if tt.pathID != "" {
				req.SetPathValue("id", tt.pathID)
			}

			rr := httptest.NewRecorder()

			handler := tt.function(mock)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("%s: expected status %d, got %d. Body: %s", tt.name, tt.expectedStatus, rr.Code, rr.Body.String())
			}
		})
	}

}
