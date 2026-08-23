package observatory

import "context"

type SummaryService struct{ query *QueryService }

func NewSummaryService(q *QueryService) *SummaryService { return &SummaryService{query: q} }
func (s *SummaryService) Build(ctx context.Context, id string) (Statistics, error) {
	layers, err := s.query.List(ctx, id)
	if err != nil {
		return Statistics{}, err
	}
	return ComputeStatistics(layers), nil
}
func (s *SummaryService) Names(ctx context.Context, id string) ([]string, error) {
	layers, err := s.query.List(ctx, id)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(layers))
	for _, l := range SortByName(layers) {
		out = append(out, l.Name)
	}
	return out, nil
}
