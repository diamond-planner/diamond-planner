import {error} from "@sveltejs/kit";
import {client} from "$lib/dp/client.svelte.js";
import {watchWithPagination} from "$lib/dp/records/RecordOperations.ts";
import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
import type {PageLoad} from "../../../../../../../../.svelte-kit/types/src/routes";

export const load = (async ({fetch, parent, params, depends}) => {
  const data = await parent();
  const teams: ExpandedTeam[] = data.teams;

  let team = teams.find((team) => team.id === params.id);

  if (!team) {
    team = await client.collection("teams").getOne<ExpandedTeam>(params.id, {
      expand: "club,admins",
      fetch: fetch,
    });
  }
  if (!team) throw error(404, "Team not found");

  const players = await watchWithPagination<CustomAuthModel>(
      "users",
      {
        filter: `teams ?~ "${params.id}"`,
        fetch: fetch,
      },
      1,
      99
  );

  depends("members:list");

  return {
    players: players,
    team: team,
  };
}) satisfies PageLoad;
