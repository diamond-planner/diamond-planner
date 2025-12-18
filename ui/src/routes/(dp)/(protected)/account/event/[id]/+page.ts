import {watchSingleRecord} from "$lib/dp/records/RecordOperations.ts";
import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.js";
import type {PageLoad} from "./$types";

export const load = (async ({fetch, params, depends}) => {
  const event = await watchSingleRecord<ExpandedEvent>("events", params.id, {
    expand:
      "participations_via_event.user, attire, location, team.admins, team.club, comments_via_event.user",
    fetch: fetch,
    requestKey: `event-single-${params.id}`,
  });

  depends("event:list", "comments:list");

  return {
    event: event,
  };
}) satisfies PageLoad;
