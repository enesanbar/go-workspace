package list

import (
	"context"
	"errors"
	"testing"

	"github.com/enesanbar/workspace/golang/projects/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLister_Do_happyPath(t *testing.T) {
	// configure the mock loader
	mockResult := []*data.Person{
		{
			ID:       1234,
			FullName: "Sally",
		},
		{
			ID:       5678,
			FullName: "Jane",
		},
	}

	mockLoader := &mockMyLoader{}
	mockLoader.On("LoadAll", mock.Anything).Return(mockResult, nil).Once()

	// call method
	lister := &Lister{
		data: mockLoader,
	}
	persons, err := lister.load(context.Background())

	// validate expectations
	require.NoError(t, err)
	assert.Equal(t, 2, len(persons))
	assert.True(t, mockLoader.AssertExpectations(t))
}

func TestLister_Do_noResults(t *testing.T) {
	// configure the mock loader
	mockLoader := &mockMyLoader{}
	mockLoader.On("LoadAll", mock.Anything).Return(nil, data.ErrNotFound).Once()

	// call method
	lister := &Lister{
		data: mockLoader,
	}
	persons, err := lister.load(context.Background())

	// validate expectations
	require.Equal(t, errPeopleNotFound, err)
	assert.Equal(t, 0, len(persons))
	assert.True(t, mockLoader.AssertExpectations(t))
}

func TestLister_Do_error(t *testing.T) {
	// configure the mock loader
	mockLoader := &mockMyLoader{}
	mockLoader.On("LoadAll", mock.Anything).Return(nil, errors.New("something failed")).Once()

	// call method
	lister := &Lister{
		data: mockLoader,
	}
	persons, err := lister.load(context.Background())

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, len(persons))
	assert.True(t, mockLoader.AssertExpectations(t))
}
