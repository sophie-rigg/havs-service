package exposure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sophie-rigg/havs-service/models"
	"github.com/sophie-rigg/havs-service/models/test_matchers"
	"github.com/sophie-rigg/havs-service/storage/mocks"
	"io"
	"net/http/httptest"
	"testing"
)

func Test_handler_Post(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		body           string
		storage        func(client *mocks.MockClient)
		want           *models.Exposure
		wantStatusCode int
	}{
		{
			name:   "Post exposure",
			url:    "/exposure",
			method: "POST",
			body:   `{"equipment_id": "123", "duration": 60, "user_id": "123"}`,
			storage: func(client *mocks.MockClient) {
				client.EXPECT().GetEquipmentItem("123").Return(&models.EquipmentItem{
					ID:                   "123",
					Name:                 "Test Equipment",
					VibrationalMagnitude: 2.6,
				}, nil)

				client.EXPECT().GetUser("123").Return(&models.User{
					ID:   "123",
					Name: "Test User",
				}, nil)

				client.EXPECT().InsertExposure(test_matchers.NewExposureMatcher(&models.Exposure{
					Equipment: &models.EquipmentItem{
						ID:                   "123",
						Name:                 "Test Equipment",
						VibrationalMagnitude: 2.6,
					},
					Duration: 60,
					A8:       0.9192388155425119,
					Points:   13.520000000000001,
					User: &models.User{
						ID:   "123",
						Name: "Test User",
					},
				})).Return(nil)
			},
			want: &models.Exposure{
				Equipment: &models.EquipmentItem{
					ID:                   "123",
					Name:                 "Test Equipment",
					VibrationalMagnitude: 2.6,
				},
				Duration: 60,
				A8:       0.9192388155425119,
				Points:   13.520000000000001,
				User: &models.User{
					ID:   "123",
					Name: "Test User",
				},
			},
			wantStatusCode: 200,
		},
		{
			name:   "Post exposure with invalid request",
			url:    "/exposure",
			method: "POST",
			body:   `{"equipment_id": "123", "duration": 60, "user_id": ""}`,
			storage: func(client *mocks.MockClient) {
				client.EXPECT().GetEquipmentItem("123").Return(&models.EquipmentItem{
					ID:                   "123",
					Name:                 "Test Equipment",
					VibrationalMagnitude: 2.6,
				}, nil)

				client.EXPECT().InsertExposure(test_matchers.NewExposureMatcher(&models.Exposure{
					Equipment: &models.EquipmentItem{
						ID:                   "123",
						Name:                 "Test Equipment",
						VibrationalMagnitude: 2.6,
					},
					Duration: 60,
					A8:       0.9192388155425119,
					Points:   13.520000000000001,
					User:     &models.User{},
				})).Return(nil)
			},
			want: &models.Exposure{
				Equipment: &models.EquipmentItem{
					ID:                   "123",
					Name:                 "Test Equipment",
					VibrationalMagnitude: 2.6,
				},
				Duration: 60,
				A8:       0.9192388155425119,
				Points:   13.520000000000001,
				User:     &models.User{},
			},
			wantStatusCode: 400,
		},
		{
			name:   "Post exposure with invalid equipment",
			url:    "/exposure",
			method: "POST",
			body:   `{"equipment_id": "123", "duration": 60, "user": "123"}`,
			storage: func(client *mocks.MockClient) {
				client.EXPECT().GetEquipmentItem("123").Return(nil, fmt.Errorf("equipment not found"))
			},
			want:           nil,
			wantStatusCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestBody bytes.Buffer
			if tt.body != "" {
				_, err := requestBody.WriteString(tt.body)
				if err != nil {
					t.Fatal(err)
				}
			}
			req := httptest.NewRequest(tt.method, tt.url, &requestBody)
			w := httptest.NewRecorder()

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
				var resultExposure models.Exposure
				err = json.Unmarshal(body, &resultExposure)
				if err != nil {
					t.Fatal(err)
				}

				if !(test_matchers.NewExposureMatcher(tt.want)).Matches(&resultExposure) {
					t.Fatal(fmt.Sprintf("expected %v, got: %v", tt.want, resultExposure))
				}
			}

		})
	}
}
