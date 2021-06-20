package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"

	mocks2 "github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers/mocks"

	"github.com/stretchr/testify/require"

	"github.com/alexedwards/scs/v2/memstore"
	"github.com/enesanbar/workspace/golang/projects/vigilate/config"

	"github.com/alexedwards/scs/v2"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/helpers"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/services"
	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/session"
)

func TestHandlerLoginGet_ServeHTTP(t *testing.T) {
	type fields struct {
		Session *session.Session
	}
	type args struct {
		mock func() *mocks2.PrefRepository
		r    func() *http.Request
	}

	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedStatus int
	}{
		{
			name: "login screen",
			args: args{
				mock: func() *mocks2.PrefRepository {
					mockPrefRepo := &mocks2.PrefRepository{}
					mockPrefRepo.On("AllPreferences").Return([]models.Preference{}, nil).Once()
					return mockPrefRepo
				},
				r: func() *http.Request {
					req, err := http.NewRequest("GET", "/person/list", nil)
					require.NoError(t, err)

					return req
				},
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memSession := session.NewSession(scs.New(), memstore.New(), &config.Config{InProduction: false})

			h := &HandlerLoginGet{
				Session:  memSession,
				Renderer: services.NewRenderer(memSession, helpers.NewPreferences(tt.args.mock())),
			}

			response := httptest.NewRecorder()

			//memSession.Manager.LoadAndSave()
			h.ServeHTTP(response, tt.args.r())

			// validate outputs
			require.Equal(t, tt.expectedStatus, response.Code)

		})
	}
}
