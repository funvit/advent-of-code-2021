package aoc

type (
	Registry struct {
		days map[uint]DaySolution
	}
	DaySolution struct {
		Part1 Solver
		Part2 Solver
	}
	Solver func(lines []string) (result interface{})
)

func (r *Registry) GetSolver(day uint, part uint) (solver Solver, ok bool) {
	s, ok := r.days[day]
	if !ok {
		return nil, false
	}

	switch part {
	case 1:
		return s.Part1, true
	case 2:
		return s.Part2, true
	default:
		return nil, false
	}
}

var defaultRegistry *Registry

func init() {
	defaultRegistry = NewRegistry()
}

func DefaultRegistry() *Registry {
	return defaultRegistry
}

func NewRegistry() *Registry {
	return &Registry{
		days: map[uint]DaySolution{},
	}
}

func RegisterSolutionWithRegistry(day uint, solution DaySolution, registry *Registry) {
	registry.days[day] = solution
}

func RegisterSolution(day uint, solution DaySolution) {
	RegisterSolutionWithRegistry(day, solution, defaultRegistry)
}
