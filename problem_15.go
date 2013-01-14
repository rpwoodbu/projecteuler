package projecteuler

/*
  http://projecteuler.net/problem=15

  Starting in the top left corner of a 2x2 grid, there are 6 routes (without backtracking) to the bottom right corner.

  [See link above for depictive graphic.]

  How many routes are there through a 20x20 grid?
*/

func GridRouteCounter(x, y int) int64 {
  // Add one because we're interested in intersections, not cells.
  cache := make([][]int64, x+1)
  for i := range cache {
	cache[i] = make([]int64, y+1)
  }

  // Seed the cache with an initial bottom-right value:
  cache[x][y] = 1

  return GridRouteLookup(0, 0, cache)
}

func GridRouteLookup(x, y int, cache [][]int64) int64 {
  routes := cache[x][y]
  if routes != 0 {
	return routes
  }

  sizeX := len(cache)
  sizeY := len(cache[0])

  if x+1 < sizeX {
	routes += GridRouteLookup(x+1, y, cache)
  }

  if y+1 < sizeY {
	routes += GridRouteLookup(x, y+1, cache)
  }

  if routes == 0 {
	panic("Found point with no routes.")
  }

  cache[x][y] = routes
  return routes
}

func Problem15() int64 {
  return GridRouteCounter(20, 20)
}
