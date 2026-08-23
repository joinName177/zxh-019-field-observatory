package observatory

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Observation struct {
	ID         string            `json:"id"`
	Station    string            `json:"station"`
	Measured   float64           `json:"measured"`
	Unit       string            `json:"unit"`
	Observed   time.Time         `json:"observed"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

func NewObservation(id, station, unit string, measured float64, observed time.Time) Observation {
	return Observation{ID: id, Station: station, Unit: unit, Measured: measured, Observed: observed.UTC(), Attributes: map[string]string{}}
}

func (o Observation) Validate() error {
	if strings.TrimSpace(o.ID) == "" || strings.TrimSpace(o.Station) == "" {
		return fmt.Errorf("observation id and station are required")
	}
	if strings.TrimSpace(o.Unit) == "" {
		return fmt.Errorf("observation unit is required")
	}
	if o.Observed.IsZero() {
		return fmt.Errorf("observation timestamp is required")
	}
	return nil
}

func CloneObservation(o Observation) Observation {
	o.Attributes = cloneMap(o.Attributes)
	return o
}

func cloneMap(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

type ObservationStore struct {
	items map[string]Observation
}

func NewObservationStore() *ObservationStore {
	return &ObservationStore{items: map[string]Observation{}}
}

func (s *ObservationStore) Put(ctx context.Context, observation Observation) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := observation.Validate(); err != nil {
		return err
	}
	s.items[observation.ID] = CloneObservation(observation)
	return nil
}

func (s *ObservationStore) Get(ctx context.Context, id string) (Observation, error) {
	if err := ctx.Err(); err != nil {
		return Observation{}, err
	}
	item, ok := s.items[id]
	if !ok {
		return Observation{}, ErrObservationNotFound
	}
	return CloneObservation(item), nil
}

func (s *ObservationStore) List(ctx context.Context, station string) ([]Observation, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	items := make([]Observation, 0, len(s.items))
	for _, item := range s.items {
		if station == "" || item.Station == station {
			items = append(items, CloneObservation(item))
		}
	}
	return items, nil
}
