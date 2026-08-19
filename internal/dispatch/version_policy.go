package dispatch

import "riderguard/internal/domain"

func preferNewerRule(left, right *domain.Rule) bool {
	if left == nil || right == nil {
		return false
	}
	return left.SelectionVersion() > right.SelectionVersion()
}
