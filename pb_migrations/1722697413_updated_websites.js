/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("u0bpx7lw5zag2vc")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "az7yvmi2",
    "name": "authorId",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("u0bpx7lw5zag2vc")

  // remove
  collection.schema.removeField("az7yvmi2")

  return dao.saveCollection(collection)
})
