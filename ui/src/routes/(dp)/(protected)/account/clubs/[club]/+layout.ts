import {client} from "$lib/dp/client.svelte.js";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";
import type {
  ExpandedAnnouncement,
  ExpandedClub,
  ExpandedTeam,
  ExpandedUniformSet,
} from "$lib/dp/types/ExpandedResponse.ts";
import type {LayoutLoad} from "./$types";
import {Collection} from "$lib/dp/enum/Collection.ts";

export const load = (async ({fetch, params, depends, url}) => {
  const club = await client.collection(Collection.Clubs).getOne<ExpandedClub>(params.club, {
    expand: "admins",
    fetch: fetch,
  });

  const teams = await client.collection(Collection.Teams).getFullList<ExpandedTeam>({
    filter: `club.id = "${club.id}"`,
    fetch: fetch,
    expand: "club,admins",
    sort: "+name",
  });

  const uniformSets = await client.collection(Collection.UniformSets).getFullList<ExpandedUniformSet>({
    filter: `club.id = "${club.id}"`,
    fetch: fetch,
    expand: "club",
  });

  const pageQuery = url.searchParams.get("page") ?? "1";
  const page = Number(pageQuery);

  const announcements = await watchWithPagination<ExpandedAnnouncement>(
      Collection.Announcements,
      {
        filter: `club.id = "${club.id}"`,
        sort: "-updated",
        fetch: fetch,
        expand: "author,club,team,comments_via_announcement.user",
        requestKey: `club-${club.id}-announcements`,
      },
      page,
      3
  );

  depends("club:single");

  return {
    club: club,
    teams: teams,
    uniformSets: uniformSets,
    announcementStore: announcements,
  };
}) satisfies LayoutLoad;
