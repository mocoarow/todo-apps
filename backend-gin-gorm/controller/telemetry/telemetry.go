package telemetry

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/baggage"
)

func AddBaggageMembers(ctx context.Context, values map[string]string) (context.Context, error) {
	bag := baggage.FromContext(ctx)
	for key, value := range values {
		member, err := baggage.NewMember(key, value)
		if err != nil {
			return nil, fmt.Errorf("baggage.NewMember: %w", err)
		}
		if newBag, err := bag.SetMember(member); err == nil {
			bag = newBag
		} else {
			return nil, fmt.Errorf("baggage.SetMember: %w", err)
		}
	}
	ctx = baggage.ContextWithBaggage(ctx, bag)

	return ctx, nil
}
