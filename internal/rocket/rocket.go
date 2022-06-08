//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/QQuinn03/go_grpc/internal/rocket Store

package rocket

import "context"

// Rocket --should define the definition of our // Rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we need to satisfy for our // service to work correctly
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service --our service, responsible for //update the rocket inventory
type Service struct {
	Store Store
}

// New --returns a new instant of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketById -- Retrieves a rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// AddRocket - Adds a rocket to our store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket - most likely rapid // unscheduled disassembly
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
