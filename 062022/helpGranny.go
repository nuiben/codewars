package kata

import "math"

func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
	var dist float64 = 0.0

	//process needs to find only the friends which are in ftwns
	//this could be a function, but for now we use this nested for-loop to create a slice
	//which only contains the friends we know the locations of
	var friendsInMap []string
	var searchString = ""

	for _, j := range arrFriends {
		searchString = j
		for i, _ := range ftwns {
			if i == searchString {
				friendsInMap = append(friendsInMap, i)
			}
		}
	}

	//here we account for the first and last legs of the journey which dont require the Pythagorean theorem
	var prev string = arrFriends[0]
	var last string = arrFriends[len(friendsInMap)-1]
	dist += h[ftwns[prev]]
	dist += h[ftwns[last]]

	//Here we calculate the opposite sides of the right triangles passing the adjacent (prev) and hypotenuse (i)
	//to a func called findOpp, adding each side to our total distance and updating 'prev' variable to i.
	// This makes the hypotenuse of the last triangle become the adjacent of the next triangle.

	for _, i := range friendsInMap[1:] {
		dist += findOpp(h[ftwns[prev]], h[ftwns[i]])
		prev = i
	}

	return int(dist)
}

func findOpp(adj, hyp float64) float64 {

	//we have adjacent (b) and hypotenuse (c) but need to findOpp (a)
	//if a^2 + b^2 = c^2
	//then c^2 - b^2 = a^2
	var opp float64 = 0.0
	opp = math.Sqrt((hyp * hyp) - (adj * adj))
	return opp
}
