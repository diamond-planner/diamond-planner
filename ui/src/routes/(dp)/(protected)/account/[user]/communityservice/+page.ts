import type {PageLoad} from './$types';
import {authSettings, client} from "$lib/dp/client.svelte.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";
import type {CustomAuthModel, ExpandedServiceEntry} from "$lib/dp/types/ExpandedResponse.ts";
import {error} from "@sveltejs/kit";

export const load: PageLoad = (async ({params, fetch}) => {
  let user: CustomAuthModel;
  if (params.user === authSettings?.record?.id) {
    user = authSettings.record as CustomAuthModel;
  } else {
    user = await client.collection(Collection.Users).getOne<CustomAuthModel>(params.user);
  }

  if (!user) {
    error(404, "User not found")
  }

  const ret: Record<string, Promise<ExpandedServiceEntry[]>> = {};

  user.club.forEach((clubID) => {
    ret.clubID = client.collection(Collection.ServiceEntries).getFullList<ExpandedServiceEntry>({
      expand: "club",
      fetch: fetch,
      filter: `club.id = "${clubID}"`,
      sort: "-service_date",
      requestKey: `${params.user}-${clubID}-serviceentries`,
    })
  })

  return {
    user: user,
    serviceEntries: ret,
  };
});
