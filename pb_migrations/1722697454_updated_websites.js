/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("u0bpx7lw5zag2vc")

  collection.listRule = "@request.auth.id = authorId"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("u0bpx7lw5zag2vc")

  collection.listRule = "@request.auth.email = 'alahmadrosid@gmail.com'"

  return dao.saveCollection(collection)
})
