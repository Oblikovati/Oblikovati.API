// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// PlanMove describes a policy-aware sketch move without changing the document.
// Example: plan, err := c.Sketch().PlanMove(args).
// mcp:tool plan_sketch_move
// mcp:summary Plan a constraint-aware sketch move without changing the document.
func (s Sketch) PlanMove(args wire.PlanSketchMoveArgs) (wire.PlanSketchMoveResult, error) {
	return call[wire.PlanSketchMoveResult](s.c, wire.MethodSketchPlanMove, args)
}

// CommitMove atomically applies an unchanged plan after any required approvals.
// A caller must also check that result.Committed is true before reporting success.
// Example: result, err := c.Sketch().CommitMove(args).
// mcp:tool commit_sketch_move
// mcp:summary Commit an unchanged sketch move plan with explicit approvals.
func (s Sketch) CommitMove(args wire.CommitSketchMoveArgs) (wire.CommitSketchMoveResult, error) {
	return call[wire.CommitSketchMoveResult](s.c, wire.MethodSketchCommitMove, args)
}
