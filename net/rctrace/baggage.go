// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rctrace

import (
	"context"

	"go.opentelemetry.io/otel/baggage"

	"github.com/focela/ratcatcher/container/rcmap"
	"github.com/focela/ratcatcher/container/rcvar"
	"github.com/focela/ratcatcher/utils/rcconv"
)

// Baggage holds the data through all tracing spans.
type Baggage struct {
	ctx context.Context
}

// NewBaggage creates and returns a new Baggage object from given tracing context.
func NewBaggage(ctx context.Context) *Baggage {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Baggage{
		ctx: ctx,
	}
}

// Ctx returns the context that Baggage holds.
func (b *Baggage) Ctx() context.Context {
	return b.ctx
}

// SetValue is a convenient function for adding one key-value pair to baggage.
// Note that it uses attribute.Any to set the key-value pair.
func (b *Baggage) SetValue(key string, value interface{}) context.Context {
	member, _ := baggage.NewMember(key, rcconv.String(value))
	bag, _ := baggage.New(member)
	b.ctx = baggage.ContextWithBaggage(b.ctx, bag)
	return b.ctx
}

// SetMap is a convenient function for adding map key-value pairs to baggage.
// Note that it uses attribute.Any to set the key-value pair.
func (b *Baggage) SetMap(data map[string]interface{}) context.Context {
	members := make([]baggage.Member, 0)
	for k, v := range data {
		member, _ := baggage.NewMember(k, rcconv.String(v))
		members = append(members, member)
	}
	bag, _ := baggage.New(members...)
	b.ctx = baggage.ContextWithBaggage(b.ctx, bag)
	return b.ctx
}

// GetMap retrieves and returns the baggage values as map.
func (b *Baggage) GetMap() *rcmap.StrAnyMap {
	m := rcmap.NewStrAnyMap()
	members := baggage.FromContext(b.ctx).Members()
	for i := range members {
		m.Set(members[i].Key(), members[i].Value())
	}
	return m
}

// GetVar retrieves value and returns a *rcvar.Var for specified key from baggage.
func (b *Baggage) GetVar(key string) *rcvar.Var {
	value := baggage.FromContext(b.ctx).Member(key).Value()
	return rcvar.New(value)
}
