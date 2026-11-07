package auth

import (
	"context"
	"slices"
	"strings"

	"github.com/google/uuid"
)

// Scope is where a grant applies. All wins over Sites.
type Scope struct {
	All   bool
	Sites []uuid.UUID
}

// Grant is a role held at a scope, with its complex already expanded to sites.
type Grant struct {
	RoleKey     string
	RoleName    string
	Permissions []string
	ScopeType   string
	ScopeID     *uuid.UUID
	Scope       Scope
}

// Principal is the signed-in user.
type Principal struct {
	UserID     uuid.UUID
	EmployeeID uuid.UUID
	SiteID     uuid.UUID
	SessionID  uuid.UUID
	Channel    string
	DeviceID   *uuid.UUID
	IP         string
	Limit      Limit
	Grants     []Grant
	// Active is the scope of the permission the current operation needs.
	Active Scope
}

// Matches reports whether pattern covers key. A "*" stands for one segment,
// and as the last segment for all the rest, so "cash.*", "org.*.read" and "*"
// all work. "*.read" means that action in any module, at any depth.
func Matches(pattern, key string) bool {
	if pattern == "*" || pattern == key {
		return true
	}
	if action, ok := strings.CutPrefix(pattern, "*."); ok {
		return strings.HasSuffix(key, "."+action)
	}
	want, have := strings.Split(pattern, "."), strings.Split(key, ".")
	for i, part := range want {
		last := i == len(want)-1
		if last && part == "*" {
			return len(have) > i
		}
		if i >= len(have) || (part != "*" && part != have[i]) {
			return false
		}
	}
	return len(want) == len(have)
}

// ScopeFor returns the union of scopes granting permission, and false if none.
func (p *Principal) ScopeFor(permission string) (Scope, bool) {
	var out Scope
	found := false
	for _, g := range p.Grants {
		if !slices.ContainsFunc(g.Permissions, func(k string) bool { return Matches(k, permission) }) {
			continue
		}
		found = true
		out.All = out.All || g.Scope.All
		out.Sites = append(out.Sites, g.Scope.Sites...)
	}
	return out, found
}

// Can reports whether the principal holds permission anywhere.
func (p *Principal) Can(permission string) bool {
	_, ok := p.ScopeFor(permission)
	return ok
}

type principalKey struct{}

// With returns ctx carrying p.
func With(ctx context.Context, p *Principal) context.Context {
	return context.WithValue(ctx, principalKey{}, p)
}

// From returns the principal in ctx, or nil for system work.
func From(ctx context.Context) *Principal {
	p, _ := ctx.Value(principalKey{}).(*Principal)
	return p
}
