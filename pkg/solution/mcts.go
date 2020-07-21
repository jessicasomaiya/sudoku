package solution

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// // MCTS manages rounds and maintains pointers to the tree
// type MCTS struct {
// 	Root     *Node // Keep a reference to the root node to start gambling agian
// 	MaxDepth int
// 	Rand     *rand.Rand
// 	Fixtures *Fixtures // Reusable fixtures
// 	Evaluate bool
// }

// func NewMCTS(root *Node) *MCTS {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	m := &MCTS{}
// 	m.Root = root
// 	return m
// }

// // Gamble finds the next solution and prints it when one is found
// func (m *MCTS) Gamble(loops int, fp *os.File, exportFile *os.File) {
// 	// start with the root node
// 	n := m.Root
// 	deepest := m.Root

// 	i := 0
// 	for i = 0; i < loops && m.Root.Status != DEAD; i++ {
// 		if i%1000000 == 1 {
// 			fp.WriteString(time.Now().String() + " Deepest level reached " + strconv.Itoa(FillFixtures(m.Fixtures, deepest).Completed()) + "\n")
// 			fp.WriteString(time.Now().String() + " Progress node at loop " + strconv.Itoa(i) + "\n")
// 			FillFixtures(m.Fixtures, n).Print(fp)
// 		}
// 		best := m.BestChild(n, m.Clubs)
// 		// dead end, kill nodes
// 		if best == nil {
// 			n.RecurseLoss()
// 			n.Kill()
// 			// go back to root and gamble again
// 			n = m.Root
// 			continue
// 		}

// 		if best.Depth == (m.MaxDepth)-1 {
// 			fmt.Fprintln(fp, best.ReverseValues())
// 			fp.WriteString(time.Now().String() + " Complete node at loop " + strconv.Itoa(i) + "\n")
// 			// FillFixtures(m.Fixtures, best).Print(fp)
// 			// Print Complete with more info
// 			FillFixtures(m.Fixtures, best).PrintComplete(fp)
// 			// only export solutions where all rows are valid true's.
// 			if !m.hasCompleteFalse() {
// 				m.Fixtures.PrintRaw(exportFile)
// 			}

// 			best.RecurseLoss()
// 			best.Kill()
// 		}
// 		// we have a next node, gamble with it
// 		n = best
// 		if n.Depth > deepest.Depth {
// 			deepest = n
// 		}
// 	}

// 	if deepest == nil {
// 		fp.WriteString("Did not find anything after " + strconv.Itoa(i) + " loops...\n")
// 		return
// 	}

// 	// depth has been reached
// 	fp.WriteString("Completed " + strconv.Itoa(i) + " loops\n")
// 	fp.WriteString("Reached depth " + strconv.Itoa(deepest.Depth) + "\n")
// }

// // hasCompleteFalse returns true if a row has been tagged valid false due to the Gap rule.
// func (m *MCTS) hasCompleteFalse() bool {
// 	var valid bool
// 	var overall bool

// 	for i := 0; i < len(m.Fixtures.f); i += m.Fixtures.weeks {
// 		valid = m.Fixtures.validator.ValidateGap(m.Fixtures.f[i:i+m.Fixtures.weeks], 0)
// 		if valid == false {
// 			overall = true
// 		}
// 	}
// 	return overall
// }

// // BestChild returns the child node with the highest chance of success
// func (m *MCTS) BestChild(n *Node, clubs int) *Node {
// 	var best *Node
// 	tries := n.Tries()
// 	bestOdds, bestObj := 0.0, 0.0
// 	// bestOdds := 0.0
// 	// m.Rand.Shuffle(len(n.Children), func(i, j int) {
// 	// 	n.Children[i], n.Children[j] = n.Children[j], n.Children[i]
// 	// })
// 	for _, child := range n.Children {
// 		if child.Status == DEAD {
// 			continue
// 		}
// 		odds := m.Odds(tries, child)
// 		switch {
// 		// no bestchild yet, bestOdds = 0, set bestchild = newchild
// 		case best == nil || odds > bestOdds:
// 			bestOdds = odds
// 			best = child
// 		// Odds are same and obj of newchild is higher than current bestchild, choose newchild
// 		case odds != 0 && odds == bestOdds && child.ObjFunc > bestObj:
// 			bestOdds = odds
// 			bestObj = child.ObjFunc
// 			best = child
// 		// // if odds are same and obj are same, pick the given child
// 		case odds != 0 && odds == bestOdds && child.ObjFunc == bestObj:
// 			bestOdds = odds
// 			bestObj = child.ObjFunc
// 			best = child
// 			// case odds != 0 && odds == bestOdds && child.ObjFunc == bestObj:
// 			// 	if r := m.Rand.Intn(2); r == 1 {
// 			// 		bestOdds = odds
// 			// 		bestObj = child.ObjFunc
// 			// 		best = child
// 			// 	}
// 		}

// 	}

// 	switch {
// 	case best == nil && n.Status == FULL:
// 		return nil
// 	case n.Status == FULL:
// 		return best
// 	case best == nil:
// 		best, _ = n.GenerateChild(clubs, m.BestValue)
// 		if best != nil {
// 			n.RecurseWin()
// 		}
// 	default:
// 		// we have a best but try new node to see if it would be better
// 		odds := m.Odds(-1, n)
// 		if bestOdds < odds {
// 			child, err := n.GenerateChild(clubs, m.BestValue)
// 			if err == nil {
// 				n.RecurseWin()
// 				return child
// 			}
// 			n.Status = FULL
// 		}
// 	}
// 	return best
// }

// // BestValue will return the best valid value based on an objective function for the next
// // unscheduled slot
// func (m *MCTS) BestValue(n *Node, clubs int) (int, float64, error) {
// 	// create fixtures from the node
// 	return FillFixtures(m.Fixtures, n).NextValue(m.Evaluate)
// }

// // Odds takes the wins and losses of the node and the total wins and
// // losses of the root node and tries to calculate the odds
// func (m *MCTS) Odds(total int, n *Node) float64 {
// 	if total == -1 {
// 		maxdepth := float64(m.Clubs * (m.Clubs - 1))
// 		return float64(n.Depth) / (maxdepth - 10.0)
// 	}

// 	if n.ObjFunc == 0 {
// 		return 0
// 	}

// 	d := float64(n.Win + n.Loss)
// 	t := float64(n.Tries())

// 	if d == 0 {
// 		return 0.8
// 	}

// 	return float64(n.Win)/d + math.Sqrt(2*(math.Log(t))/d)
// }

// // UCBOne takes in a child node and retuns the Upper Confidence Bound, ie the odds of picking this child
// func (m *MCTS) UCBOne(n *Node) float64 {
// 	t := float64(n.parent.Win + n.parent.Loss)
// 	d := float64(n.Win + n.Loss)

// 	if d == 0 {
// 		return n.ObjFunc
// 	}

// 	exploit := float64(n.Win) / d
// 	explore := math.Sqrt(2 * (math.Log(t) / d))

// 	return explore + exploit
// }
