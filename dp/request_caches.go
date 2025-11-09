package dp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	RequestCacheCollection = "requestcaches"
)

var _ core.RecordProxy = (*RequestCache)(nil)

// RequestCache is a RecordProxy for the `requestcaches` collection.
// It offers type-safe getters and setters for every field.
type RequestCache struct {
	core.BaseRecordProxy
}

// ID returns the record id.
func (r *RequestCache) ID() string {
	return r.Id
}

// Hash returns the cache hash.
func (r *RequestCache) Hash() string {
	return r.GetString("hash")
}
func (r *RequestCache) SetHash(hash string) {
	r.Set("hash", hash)
}

// ResponseBody returns the cached response body.
func (r *RequestCache) ResponseBody() string {
	return r.GetString("responseBody")
}
func (r *RequestCache) SetResponseBody(responseBody string) {
	r.Set("responseBody", responseBody)
}

// URL returns the request URL.
func (r *RequestCache) URL() string {
	return r.GetString("url")
}
func (r *RequestCache) SetURL(url string) {
	r.Set("url", url)
}

// Identifier returns the cache identifier.
func (r *RequestCache) Identifier() string {
	return r.GetString("identifier")
}
func (r *RequestCache) SetIdentifier(identifier string) {
	r.Set("identifier", identifier)
}

// Created returns the creation timestamp.
func (r *RequestCache) Created() types.DateTime {
	return r.GetDateTime("created")
}
func (r *RequestCache) SetCreated(created types.DateTime) {
	r.Set("created", created)
}

// Updated returns the last update DateTime.
func (r *RequestCache) Updated() types.DateTime {
	return r.GetDateTime("updated")
}
func (r *RequestCache) SetUpdated(updated types.DateTime) {
	r.Set("updated", updated)
}

// SaveDataToCache saves a given string `data` with the provided identifier
func SaveDataToCache(app core.App, identifier string, data string) error {
	requestCache := &RequestCache{}
	collection, err := app.FindCollectionByNameOrId(RequestCacheCollection)
	if err != nil {
		return err
	}
	record := core.NewRecord(collection)
	requestCache.SetProxyRecord(record)

	requestCache.SetIdentifier(identifier)
	requestCache.SetHash(GetMD5Hash(identifier))
	requestCache.SetResponseBody(data)
	if err := app.Save(record); err != nil {
		return err
	}
	return nil
}

// GetCacheEntryByIdentifier returns the cache object for a given identifier or an error if not found.
func GetCacheEntryByIdentifier(app core.App, identifier string) (*core.Record, error) {
	record, err := app.FindFirstRecordByData(RequestCacheCollection, "identifier", identifier)
	if err != nil {
		return nil, err
	}

	return record, nil
}
