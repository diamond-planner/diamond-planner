import {client} from "$lib/dp/client.svelte.js";
import type {
  CustomAuthModel,
  ExpandedAnnouncement,
  ExpandedClub,
  ExpandedTeam
} from "$lib/dp/types/ExpandedResponse.ts";
import type {LayoutLoad} from "./$types";
import {Collection} from "$lib/dp/enum/Collection.ts";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";

export const load = (async ({fetch, depends, url, parent}) => {
  const model = client.authStore.record as CustomAuthModel;
  const data = await parent();

  if (!client.authStore.isValid) {
    return;
  }

  let clubs = data.clubs;
  let teams = data.teams;

  if (!teams) {
    teams = await client.collection(Collection.Teams).getFullList<ExpandedTeam>({
      filter: `"${model?.teams}" ?~ id`,
      expand: "club,admins",
      fetch: fetch,
      sort: "+name",
    });
  }

  if (!clubs) {
    clubs = await client.collection(Collection.Clubs).getFullList<ExpandedClub>({
      filter: `"${model?.club}" ?~ id`,
      fetch: fetch,
      expand: "admins",
    });
  }

  const pageQuery = url.searchParams.get("page") ?? "1";
  const page = Number(pageQuery);

  const announcements = await watchWithPagination<ExpandedAnnouncement>(
    Collection.Announcements,
    {
      sort: "-updated",
      fetch: fetch,
      expand: "author,club,team,comments_via_announcement.user",
      requestKey: `${model?.id}-announcements`,
    },
    page,
    3
  );

  depends("teams:list");

  return {
    clubs: clubs,
    teams: teams,
    announcementStore: announcements,
  };
}) satisfies LayoutLoad;
