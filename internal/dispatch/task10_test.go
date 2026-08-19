package dispatch

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"riderguard/internal/domain"
)

func TestSamePriorityChoosesNewestVersion(t *testing.T) {
	when := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	clk := clock.NewMock()
	clk.Set(when)
	item := &domain.RightsCase{ID: "case-10", Category: "劳动报酬", RegisteredAt: when}
	rules := []*domain.Rule{
		{Version: 1, Name: "old", MatchCategory: "劳动报酬", LeadDepartment: "旧组", Priority: 5, EffectiveFrom: when.Add(-time.Hour), Status: domain.RuleStatusActive},
		{Version: 2, Name: "new", MatchCategory: "劳动报酬", LeadDepartment: "新组", Priority: 5, EffectiveFrom: when.Add(-time.Hour), Status: domain.RuleStatusActive},
	}
	ref, err := NewAdjudicator(clk).Adjudicate(context.Background(), item, rules)
	if err != nil || ref.RuleVersion != 2 {
		t.Fatalf("expected newest rule, ref=%+v err=%v", ref, err)
	}
}
