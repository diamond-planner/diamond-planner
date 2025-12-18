import type {PageStore} from "$lib/dp/records/PageStore.ts";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";
import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
import type {Fetch} from "$lib/dp/utility/Fetch.ts";

export class EventService {
  /**
   * Takes a team ID and an URL, construct PocketBase filter string from query parameters
   * and loads events with realtime.
   */
  public async loadEventStore(
      teamID: string,
      url: URL,
      fetch: Fetch
  ): Promise<PageStore<ExpandedEvent>> {
    let filter = `team = "${teamID}"`;

    const timeframe = url.searchParams.get("timeframe");

    if (!timeframe || timeframe === "next") {
      filter = filter.concat(`&& starttime >= @todayStart`);
    } else if (timeframe === "past") {
      filter = filter.concat(`&& starttime <= @todayStart`);
    }

    const showTypes = url.searchParams.get("type") ?? "any";

    if (showTypes !== "any") {
      filter = filter.concat(`&& type = "${showTypes}"`);
    }

    // add sort parameter

    let sort = "+starttime";

    if (url.searchParams.get("sort") === "desc") {
      sort = "-starttime";
    }

    // check pagination info

    const pageNumber = Number(url.searchParams.get("page")) ?? 1;

    return await watchWithPagination<ExpandedEvent>(
        "events",
        {
          filter: filter,
          sort: sort,
          expand: "participations_via_event.user, attire, location",
          fetch: fetch,
        },
        pageNumber,
        6
    );
  }
}
