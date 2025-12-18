import {client} from "$lib/dp/client.svelte.js";
import type {ParticipationsCreate, ParticipationsResponse} from "$lib/dp/types/pb-types.ts";

export async function sendParticipationData(
    data: ParticipationsCreate
): Promise<ParticipationsResponse> {
  if (data.id) {
    return await client.collection("participations").update(data.id, data);
  } else {
    return await client.collection("participations").create(data);
  }
}
