package mud

// TrendService is a service for supervised trend forecasting
type TrendService service

// Wait for learn
func (s *TrendService) Wait() {
	s.graph.Wait()
}