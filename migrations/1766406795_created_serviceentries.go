package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": "@request.auth.id != \"\" \n&& member.club.id ?= club.id // you can only do services for your own club\n&& (club.admins.id ?= @request.auth.id // club admins may create entries for their club\n  || member.id = @request.auth.id) // or you can create for yourself",
			"deleteRule": "@request.auth.id != \"\" \n&& (club.admins.id ?= @request.auth.id // club admins may delete\n  || member.id = @request.auth.id) // or you can delete your own",
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "date1692238742",
					"max": "",
					"min": "",
					"name": "service_date",
					"presentable": true,
					"required": true,
					"system": false,
					"type": "date"
				},
				{
					"hidden": false,
					"id": "number2691386019",
					"max": null,
					"min": 1,
					"name": "minutes",
					"onlyInt": true,
					"presentable": true,
					"required": true,
					"system": false,
					"type": "number"
				},
				{
					"cascadeDelete": false,
					"collectionId": "_pb_users_auth_",
					"hidden": false,
					"id": "relation1894054520",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "member",
					"presentable": true,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text724990059",
					"max": 255,
					"min": 0,
					"name": "title",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1843675174",
					"max": 1024,
					"min": 0,
					"name": "description",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1346949349",
					"max": 1024,
					"min": 0,
					"name": "admin_comment",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"cascadeDelete": false,
					"collectionId": "_pb_users_auth_",
					"hidden": false,
					"id": "relation3707420383",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "board_member",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": true,
					"collectionId": "z806hmy5um3qy5x",
					"hidden": false,
					"id": "relation3102619762",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "club",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_3454477543",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_KrfoEt3R3A` + "`" + ` ON ` + "`" + `serviceentries` + "`" + ` (` + "`" + `service_date` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_YisjeC219R` + "`" + ` ON ` + "`" + `serviceentries` + "`" + ` (` + "`" + `member` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_aRV5YySqjm` + "`" + ` ON ` + "`" + `serviceentries` + "`" + ` (` + "`" + `club` + "`" + `)"
			],
			"listRule": "@request.auth.id != \"\" \n&& (club.admins.id ?= @request.auth.id // club admins may view\n  || member.id = @request.auth.id) // or you can view your own",
			"name": "serviceentries",
			"system": false,
			"type": "base",
			"updateRule": "@request.auth.id != \"\" \n&& (club.admins.id ?= @request.auth.id // club admins may edit\n  || member.id = @request.auth.id) // or you can edit your own",
			"viewRule": "@request.auth.id != \"\" \n&& (club.admins.id ?= @request.auth.id // club admins may view\n  || member.id = @request.auth.id) // or you can view your own"
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3454477543")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
