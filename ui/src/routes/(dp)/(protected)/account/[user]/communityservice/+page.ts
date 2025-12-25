import type {PageLoad} from './$types';
import {client} from "$lib/dp/client.svelte.ts";

export const load: PageLoad = (async ({params, fetch}) => {
  const communityServiceData = client.send(`/api/communityservice/${params.user}/${2025}`, {
    fetch: fetch,
    requestKey: `${params.user}-serviceentries`,
  });

  return {
    set: communityServiceData,
  };
});
