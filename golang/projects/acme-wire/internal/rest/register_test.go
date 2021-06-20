package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/rest/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterHandler_ServeHTTP(t *testing.T) {
	scenarios := []struct {
		desc           string
		inRequest      func() *http.Request
		inMockModel    func() *mocks.RegisterModel
		expectedStatus int
		expectedHeader string
	}{
		{
			desc: "Happy path",
			inRequest: func() *http.Request {
				validRequest := buildValidRequest()
				request, err := http.NewRequest(http.MethodPost, "/person/register", validRequest)
				require.NoError(t, err)
				return request
			},
			inMockModel: func() *mocks.RegisterModel {
				// valid downstream configuration
				resultID := 1234
				var resultErr error

				mockRegisterModel := &mocks.RegisterModel{}
				mockRegisterModel.On("Do", mock.Anything, mock.Anything).Return(resultID, resultErr).Once()

				return mockRegisterModel
			},
			expectedStatus: http.StatusCreated,
			expectedHeader: "/person/1234/",
		},
		{
			desc: "Bad input / User error",
			inRequest: func() *http.Request {
				invalidRequest := bytes.NewBufferString(`this is not a valid JSON`)
				request, err := http.NewRequest(http.MethodPost, "/person/register", invalidRequest)
				require.NoError(t, err)

				return request
			},
			inMockModel: func() *mocks.RegisterModel {
				// Dependency should not be called
				mockRegisterModel := &mocks.RegisterModel{}
				return mockRegisterModel
			},
			expectedStatus: http.StatusBadRequest,
			expectedHeader: "",
		},
		{
			desc: "Dependency error",
			inRequest: func() *http.Request {
				validRequest := buildValidRequest()
				request, err := http.NewRequest(http.MethodPost, "/person/register", validRequest)
				require.NoError(t, err)

				return request
			},
			inMockModel: func() *mocks.RegisterModel {
				// call to the dependency failed
				resultErr := errors.New("something failed")

				mockRegisterModel := &mocks.RegisterModel{}
				mockRegisterModel.On("Do", mock.Anything, mock.Anything).Return(0, resultErr).Once()

				return mockRegisterModel
			},
			expectedStatus: http.StatusBadRequest,
			expectedHeader: "",
		},
	}
	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			// define model layer mock
			mockRegisterModel := s.inMockModel()

			// build handler
			handler := NewRegisterHandler(mockRegisterModel)

			// perform request
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, s.inRequest())

			// validate outputs
			require.Equal(t, s.expectedStatus, response.Code)

			// call should output the location to a new person
			resultHeader := response.Header().Get("Location")
			assert.Equal(t, s.expectedHeader, resultHeader)

			// validate the mock was used as we expected
			assert.True(t, mockRegisterModel.AssertExpectations(t))
		})
	}
}

func buildValidRequest() io.Reader {
	requestData := &registerRequest{
		FullName: "Joan Smith",
		Currency: "AUD",
		Phone:    "01234567890",
	}

	data, _ := json.Marshal(requestData)
	return bytes.NewBuffer(data)
}
