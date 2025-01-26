package exposuresummary

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sophie-rigg/havs-service/models"
	"github.com/sophie-rigg/havs-service/storage/mocks"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_handler_Get(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		vars           map[string]string
		storage        func(client *mocks.MockClient)
		want           *models.ExposureSummary
		wantStatusCode int
	}{
		{
			name:   "success",
			url:    "/users/1/exposure-summary?starting_at=2021-01-01T00:00:00Z&ending_at=2021-01-02T00:00:00Z",
			method: "GET",
			vars:   map[string]string{"userId": "1"},
			storage: func(s *mocks.MockClient) {
				startTime, _ := time.Parse(timeFormat, "2021-01-01T00:00:00Z")
				endTime, _ := time.Parse(timeFormat, "2021-01-02T00:00:00Z")
				s.EXPECT().GetExposuresByUserID("1", startTime, endTime).Return([]*models.Exposure{
					{
						Equipment: &models.EquipmentItem{
							ID:                   "123",
							Name:                 "Test Equipment",
							VibrationalMagnitude: 2.6,
						},
						Duration: 60,
						A8:       0.9192388155425119,
						Points:   13.520000000000001,
						User: &models.User{
							ID:   "1",
							Name: "Test User",
						},
					},
					{
						Equipment: &models.EquipmentItem{
							ID:                   "123",
							Name:                 "Test Equipment",
							VibrationalMagnitude: 2.6,
						},
						Duration: 60,
						A8:       0.8192388155425119,
						Points:   15.520000000000001,
						User: &models.User{
							ID:   "1",
							Name: "Test User",
						},
					},
				}, nil)
			},
			want: &models.ExposureSummary{
				A8:     1.7384776310850238,
				Points: 29.040000000000003,
				User: &models.User{
					ID:   "1",
					Name: "Test User",
				},
			},
			wantStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.vars)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := mocks.NewMockClient(ctrl)
			tt.storage(s)

			NewHandler(s).ServeHTTP(w, req)

			resp := w.Result()

			if resp.StatusCode != tt.wantStatusCode {
				t.Fatal(fmt.Sprintf("expected %d status code, got: %d", tt.wantStatusCode, resp.StatusCode))
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if tt.wantStatusCode == 200 {
				var resultExposure models.ExposureSummary
				err = json.Unmarshal(body, &resultExposure)
				if err != nil {
					t.Fatal(err)
				}

				if !assert.Equal(t, tt.want, &resultExposure) {
					t.Fatal("expected exposure does not match result")
				}
			}
		})
	}
}
