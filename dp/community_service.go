package dp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	ServiceEntryCollection = "serviceentries"
)

var _ core.RecordProxy = (*ServiceEntry)(nil)

// ServiceEntry RecordProxy for collection `serviceentries`.
// Provides type-safe struct access to service entry records.
type ServiceEntry struct {
	core.BaseRecordProxy
}

func (s *ServiceEntry) ID() string {
	return s.Id
}

func (s *ServiceEntry) ServiceDate() types.DateTime {
	return s.GetDateTime("service_date")
}

func (s *ServiceEntry) SetServiceDate(dt types.DateTime) {
	s.Set("service_date", dt)
}

func (s *ServiceEntry) Minutes() int {
	return s.GetInt("minutes")
}

func (s *ServiceEntry) SetMinutes(minutes int) {
	s.Set("minutes", minutes)
}

func (s *ServiceEntry) Member() string {
	return s.GetString("member")
}

func (s *ServiceEntry) SetMember(id string) {
	s.Set("member", id)
}

func (s *ServiceEntry) Title() string {
	return s.GetString("title")
}

func (s *ServiceEntry) SetTitle(title string) {
	s.Set("title", title)
}

func (s *ServiceEntry) Description() string {
	return s.GetString("description")
}

func (s *ServiceEntry) SetDescription(description string) {
	s.Set("description", description)
}

func (s *ServiceEntry) AdminComment() string {
	return s.GetString("admin_comment")
}

func (s *ServiceEntry) SetAdminComment(comment string) {
	s.Set("admin_comment", comment)
}

func (s *ServiceEntry) BoardMember() string {
	return s.GetString("board_member")
}

func (s *ServiceEntry) SetBoardMember(id string) {
	s.Set("board_member", id)
}

func (s *ServiceEntry) Club() string {
	return s.GetString("club")
}

func (s *ServiceEntry) SetClub(id string) {
	s.Set("club", id)
}
