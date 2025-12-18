import type {ListResult, RecordListOptions, RecordModel, UnsubscribeFunc} from "pocketbase";
import {type Readable, readable, type Subscriber} from "svelte/store";
import {browser} from "$app/environment";
import {client} from "../client.svelte.js";
import type {PageStore} from "./PageStore.ts";

/*
 * Save (create/update) a record (a plain object). Automatically converts to
 * FormData if needed.
 */
export async function save<T>(collection: string, record: any, create = false) {
  // convert obj to FormData in case one of the fields is instanceof FileList
  const data = object2formdata(record);
  if (record.id && !create) {
    // "create" flag overrides update
    return await client.collection(collection).update<T>(record.id, data);
  } else {
    return await client.collection(collection).create<T>(data);
  }
}

// convert obj to FormData in case one of the fields is instanceof FileList
function object2formdata(obj: {}) {
  // check if any field's value is an instanceof FileList
  if (!Object.values(obj).some((val) => val instanceof FileList || val instanceof File)) {
    // if not, just return the original object
    return obj;
  }
  // otherwise, build FormData (multipart/form-data) from obj
  const fd = new FormData();
  for (const [key, val] of Object.entries(obj)) {
    if (val instanceof FileList) {
      for (const file of val) {
        fd.append(key, file);
      }
    } else if (val instanceof File) {
      // handle File before "object" so that it doesn't get serialized as JSON
      fd.append(key, val);
    } else if (Array.isArray(val)) {
      // for some reason, multipart/form-data wants arrays to be comma-separated strings
      fd.append(key, val.join(","));
    } else if (typeof val === "object") {
      fd.append(key, JSON.stringify(val));
    } else {
      fd.append(key, val as any);
    }
  }
  return fd;
}

export async function watchSingleRecord<T extends RecordModel>(
    idOrName: string,
    recordID: string,
    queryParams = {} as RecordListOptions,
    realtime = browser
): Promise<Readable<T>> {
  const collection = client.collection(idOrName);
  let result = await collection.getOne<T>(recordID, queryParams);

  let set: Subscriber<T>;
  let unsubRealtime: UnsubscribeFunc | undefined;

  const store = readable<T>(result, (_set) => {
    set = _set;
    // watch for changes (only if you're in the browser)
    if (realtime)
      collection
          .subscribe<T>(
              recordID,
              ({action, record}) => {
                ;(async (action: string) => {
                  if (action === "update") {
                    result = record;
                  }
                  return result;
                })(action).then((record) => set((result = record)));
              },
              queryParams
          )
          // remember for later
          .then((unsub) => (unsubRealtime = unsub));
  });

  return {
    ...store,
    subscribe(run, invalidate) {
      const unsubStore = store.subscribe(run, invalidate);
      return async () => {
        unsubStore();
        // ISSUE: Technically, we should AWAIT here, but that will slow down navigation UX.
        if (unsubRealtime) /* await */ unsubRealtime();
      };
    },
  };
}

export async function watchWithPagination<T extends RecordModel>(
    idOrName: string,
    queryParams = {} as RecordListOptions,
    page = 1,
    perPage = 20,
    realtime = browser
): Promise<PageStore<T>> {
  const collection = client.collection(idOrName);
  let result = await collection.getList<T>(page, perPage, queryParams);
  let set: Subscriber<ListResult<T>>;
  let unsubRealtime: UnsubscribeFunc | undefined;

  // fetch first page
  const store = readable<ListResult<T>>(result, (_set) => {
    set = _set;
    // watch for changes (only if you're in the browser)
    if (realtime)
      collection
          .subscribe<T>(
              "*",
              ({action, record}) => {
                ;(async (action: string) => {
                  // see https://github.com/pocketbase/pocketbase/discussions/505
                  switch (action) {
                      // ISSUE: no subscribe event when a record is modified and no longer fits the "filter"
                      // @see https://github.com/pocketbase/pocketbase/issues/4717
                    case "update":
                    case "create": {
                      // record = await expand(queryParams.expand, record);
                      const index = result.items.findIndex((r) => r.id === record.id);
                      // replace existing if found, otherwise append
                      if (index >= 0) {
                        result.items[index] = record;
                        return result.items;
                      } else {
                        return [...result.items, record];
                      }
                    }
                    case "delete":
                      return result.items.filter((item) => item.id !== record.id);
                  }
                  return result.items;
                })(action).then((items) => set((result = {...result, items})));
              },
              queryParams
          )
          // remember for later
          .then((unsub) => (unsubRealtime = unsub));
  });

  async function setPage(newpage: number) {
    const {page, totalPages, perPage} = result;
    if (page > 0 && page <= totalPages) {
      set((result = await collection.getList(newpage, perPage, queryParams)));
    }
  }

  return {
    ...store,
    subscribe(run, invalidate) {
      const unsubStore = store.subscribe(run, invalidate);
      return async () => {
        unsubStore();
        // ISSUE: Technically, we should AWAIT here, but that will slow down navigation UX.
        if (unsubRealtime) /* await */ unsubRealtime();
      };
    },
    setPage,
    async next() {
      setPage(result.page + 1);
    },
    async prev() {
      setPage(result.page - 1);
    },
  };
}
