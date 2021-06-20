package list

import (
	"context"
	"errors"

	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/logging"
	"github.com/enesanbar/workspace/golang/projects/acme-wire/internal/modules/data"
)

var (
	// error thrown when there are no people in the database
	errPeopleNotFound = errors.New("no people found")
)

// NewLister creates and initializes a Lister
func NewLister(cfg Config) *Lister {
	return &Lister{
		cfg: cfg,
	}
}

// Config is the config for Lister
type Config interface {
	Logger() logging.Logger
	DataDSN() string
}

// Lister will attempt to load all people in the database.
// It can return an error caused by the data layer
type Lister struct {
	cfg  Config
	data myLoader
}

// Do will load the people from the data layer
func (l *Lister) Do(ctx context.Context) ([]*data.Person, error) {
	// load all people
	people, err := l.load(ctx)
	if err != nil {
		return nil, err
	}

	if len(people) == 0 {
		// special processing for 0 people returned
		return nil, errPeopleNotFound
	}

	return people, nil
}

// load all people
func (l *Lister) load(ctx context.Context) ([]*data.Person, error) {
	people, err := l.getLoader().LoadAll(ctx)
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are encapsulating the implementation details from our users.
			return nil, errPeopleNotFound
		}
		return nil, err
	}

	return people, nil
}

func (l *Lister) getLoader() myLoader {
	if l.data == nil {
		l.data = data.NewDAO(l.cfg)

		// temporarily add a log tracker
		l.data.(*data.DAO).Tracker = data.NewLogTracker(l.cfg.Logger())
	}

	return l.data
}

//go:generate mockery -name=myLoader -case underscore -testonly -inpkg -note @generated
type myLoader interface {
	LoadAll(ctx context.Context) ([]*data.Person, error)
}
