import type {ListResult} from "pocketbase";
import type {Readable} from "svelte/store";

export interface PageStore<T = any> extends Readable<ListResult<T>> {
  setPage(newpage: number): Promise<void>;

  next(): Promise<void>;

  prev(): Promise<void>;
}
