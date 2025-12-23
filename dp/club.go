package dp

import "github.com/pocketbase/pocketbase/core"

const (
	ClubsCollection = "clubs"
)

var _ core.RecordProxy = (*Club)(nil)

// Club RecordProxy for collection `clubs`.
// Provides type-safe struct access to club records.
type Club struct {
	core.BaseRecordProxy
}

func (c *Club) ID() string {
	return c.Id
}

func (c *Club) Name() string {
	return c.GetString("name")
}

func (c *Club) SetName(name string) {
	c.Set("name", name)
}

func (c *Club) BSMID() int {
	return c.GetInt("bsm_id")
}

func (c *Club) SetBSMID(id int) {
	c.Set("bsm_id", id)
}

func (c *Club) BSMAPIKey() string {
	return c.GetString("bsm_api_key")
}

func (c *Club) SetBSMAPIKey(key string) {
	c.Set("bsm_api_key", key)
}

func (c *Club) Acronym() string {
	return c.GetString("acronym")
}

func (c *Club) SetAcronym(acronym string) {
	c.Set("acronym", acronym)
}

func (c *Club) Admins() []string {
	return c.GetStringSlice("admins")
}

func (c *Club) SetAdmins(adminIDs []string) {
	c.Set("admins", adminIDs)
}

func (c *Club) ServiceRequirement() int {
	return c.GetInt("service_requirement")
}

func (c *Club) SetServiceRequirement(req int) {
	c.Set("service_requirement", req)
}
