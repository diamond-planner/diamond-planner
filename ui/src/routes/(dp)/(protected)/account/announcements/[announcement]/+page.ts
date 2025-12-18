import {watchSingleRecord} from "$lib/dp/records/RecordOperations.ts";
import type {ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
import type {PageLoad} from "../../../../../../../.svelte-kit/types/src/routes";

export const load = (async ({params, fetch, depends}) => {
  const announcement = await watchSingleRecord<ExpandedAnnouncement>(
      "announcements",
      params.announcement,
      {
        fetch: fetch,
        expand: "author,club,team.club,comments_via_announcement.user",
        requestKey: `announcement-single-${params.announcement}`,
      }
  );
  depends("announcement:single", "comments:list");

  return {
    announcement: announcement,
  };
}) satisfies PageLoad;
