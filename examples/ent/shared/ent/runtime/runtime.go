// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/go-saas/saas/examples/ent/shared/ent/post"
	"github.com/go-saas/saas/examples/ent/shared/ent/schema"
	"github.com/go-saas/saas/examples/ent/shared/ent/tenant"
	"github.com/go-saas/saas/examples/ent/shared/ent/tenantconn"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postMixin := schema.Post{}.Mixin()
	postMixinHooks0 := postMixin[0].Hooks()
	post.Hooks[0] = postMixinHooks0[0]
	postMixinInters0 := postMixin[0].Interceptors()
	post.Interceptors[0] = postMixinInters0[0]
	tenantMixin := schema.Tenant{}.Mixin()
	tenantMixinFields0 := tenantMixin[0].Fields()
	_ = tenantMixinFields0
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescCreateTime is the schema descriptor for create_time field.
	tenantDescCreateTime := tenantMixinFields0[0].Descriptor()
	// tenant.DefaultCreateTime holds the default value on creation for the create_time field.
	tenant.DefaultCreateTime = tenantDescCreateTime.Default.(func() time.Time)
	// tenantDescUpdateTime is the schema descriptor for update_time field.
	tenantDescUpdateTime := tenantMixinFields0[1].Descriptor()
	// tenant.DefaultUpdateTime holds the default value on creation for the update_time field.
	tenant.DefaultUpdateTime = tenantDescUpdateTime.Default.(func() time.Time)
	// tenant.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	tenant.UpdateDefaultUpdateTime = tenantDescUpdateTime.UpdateDefault.(func() time.Time)
	tenantconnMixin := schema.TenantConn{}.Mixin()
	tenantconnMixinFields0 := tenantconnMixin[0].Fields()
	_ = tenantconnMixinFields0
	tenantconnFields := schema.TenantConn{}.Fields()
	_ = tenantconnFields
	// tenantconnDescCreateTime is the schema descriptor for create_time field.
	tenantconnDescCreateTime := tenantconnMixinFields0[0].Descriptor()
	// tenantconn.DefaultCreateTime holds the default value on creation for the create_time field.
	tenantconn.DefaultCreateTime = tenantconnDescCreateTime.Default.(func() time.Time)
	// tenantconnDescUpdateTime is the schema descriptor for update_time field.
	tenantconnDescUpdateTime := tenantconnMixinFields0[1].Descriptor()
	// tenantconn.DefaultUpdateTime holds the default value on creation for the update_time field.
	tenantconn.DefaultUpdateTime = tenantconnDescUpdateTime.Default.(func() time.Time)
	// tenantconn.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	tenantconn.UpdateDefaultUpdateTime = tenantconnDescUpdateTime.UpdateDefault.(func() time.Time)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
