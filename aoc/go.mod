module example.com/aoc

go 1.21.4

replace example.com/day => ../day

require (
	example.com/day v0.0.0-00010101000000-000000000000
	example.com/utility v0.0.0-00010101000000-000000000000
)

replace example.com/utility => ../utility
