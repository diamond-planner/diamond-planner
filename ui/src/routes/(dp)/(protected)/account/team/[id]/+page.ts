import {error} from "@sveltejs/kit";
import {dev} from "$app/environment";
import {client} from "$lib/dp/client.svelte.js";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";
import {EventService} from "$lib/dp/service/EventService.ts";
import type {ExpandedAnnouncement, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.js";
import type {EventseriesResponse} from "$lib/dp/types/pb-types.ts";
import type {PageLoad} from "../../../../../../../.svelte-kit/types/src/routes";

export const load = (async ({fetch, parent, params, url, depends}) => {
  const data = await parent();
  const teams: ExpandedTeam[] = data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    try {
      team = await client.collection("teams").getOne<ExpandedTeam>(params.id, {
        expand: "club,admins",
        fetch: fetch,
      });
    } catch (e) {
      if (dev) {
        console.error(e);
      }
      error(404, "There was an error loading the team.");
    }
  }
  if (!team) throw error(404, "Team not found");

  const eventService = new EventService();
  const events = await eventService.loadEventStore(team.id, url, fetch);

  const targetDate = new Date();
  targetDate.setFullYear(targetDate.getFullYear() - 1, 0, 1);
  const eventSeriesCutoff = targetDate.toISOString();

  const eventSeries = await client.collection("eventseries").getFullList<EventseriesResponse>({
    filter: `team = "${team.id}" && series_start >= "${eventSeriesCutoff}"`,
    fetch: fetch,
    sort: "+series_start",
  });

  const pageQuery = url.searchParams.get("page") ?? "1";
  const page = Number(pageQuery);

  const announcements = await watchWithPagination<ExpandedAnnouncement>(
      "announcements",
      {
        filter: `team.id = "${team.id}"`,
        sort: "-updated",
        fetch: fetch,
        expand: "author,club,team,comments_via_announcement.user",
        requestKey: `team-${team.id}-announcements`,
      },
      page,
      3
  );

  depends("event:list", "comments:list");

  return {
    team: team,
    events: events,
    eventSeries: eventSeries,
    announcementStore: announcements,
  };
}) satisfies PageLoad;
