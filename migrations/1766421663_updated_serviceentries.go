package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3454477543")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" \n&& (club.admins.id ?= board_member.id || @request.body.board_member:isset = false)\n&& member.club.id ?= club.id // you can only do services for your own club\n&& (club.admins.id ?= @request.auth.id // club admins may create entries for their club\n  || member.id = @request.auth.id) // or you can create for yourself",
			"updateRule": "@request.auth.id != \"\"\n&& (club.admins.id ?= board_member.id || @request.body.board_member:isset = false)\n&& (club.admins.id ?= @request.auth.id // club admins may edit\n  || member.id = @request.auth.id) // or you can edit your own"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3454477543")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" \n&& club.admins.id ?= board_member.id\n&& member.club.id ?= club.id // you can only do services for your own club\n&& (club.admins.id ?= @request.auth.id // club admins may create entries for their club\n  || member.id = @request.auth.id) // or you can create for yourself",
			"updateRule": "@request.auth.id != \"\"\n&& club.admins.id ?= board_member.id\n&& (club.admins.id ?= @request.auth.id // club admins may edit\n  || member.id = @request.auth.id) // or you can edit your own"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
