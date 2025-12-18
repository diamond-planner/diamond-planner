<script lang="ts">
  import {Check, CircleQuestionMark, X} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {EventParticipationState} from "$lib/dp/types/EventParticipationState.ts";
  import type {CustomAuthModel, ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ParticipationsCreate} from "$lib/dp/types/pb-types.ts";
  import {sendParticipationData} from "$lib/dp/utility/sendParticipationData.ts";

  interface props {
    event: ExpandedEvent;
    chipClasses?: string;
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const {event, chipClasses = ""}: props = $props();

  const userParticipation: ParticipationsCreate = $derived(
      event.userParticipation ?? {
        id: "",
        user: authRecord?.id,
        event: event.id,
        state: "",
        comment: "",
      }
  );

  // JS: splitting an empty string by comma returns length `1`
  const guestCount = $derived(event.guests === "" ? 0 : event.guests.split(",").length);

  let submitting = $state(false);

  async function updateParticipationStatus(state: EventParticipationState): Promise<void> {
    if (submitting) {
      return;
    }
    submitting = true;
    try {
      userParticipation.state = state;
      await sendParticipationData(userParticipation);
      await invalidate("event:list");
    } finally {
      submitting = false;
    }
  }
</script>

<section class="flex justify-end items-end gap-2 flex-wrap">
  <button
          aria-busy={submitting}
          class="chip {chipClasses} preset-outlined-success-500 participation-button"
          class:preset-filled-success-500={userParticipation.state === "in"}
          disabled={submitting}
          onclick={() => updateParticipationStatus("in")}
  >
    <span><Check size="14"/></span>
    <span>{(event.participations.in.length ?? 0) + guestCount}</span>
  </button>

  <button
          aria-busy={submitting}
          class="chip {chipClasses} preset-outlined-warning-500 participation-button"
          class:preset-filled-warning-500={userParticipation.state === "maybe"}
          disabled={submitting}
          onclick={() => updateParticipationStatus("maybe")}
  >
    <span><CircleQuestionMark size="14"/></span>
    <span>{event.participations.maybe.length ?? 0}</span>
  </button>

  <button
          aria-busy={submitting}
          class="chip {chipClasses} preset-outlined-error-500 participation-button"
          class:preset-filled-error-500={userParticipation.state === "out"}
          disabled={submitting}
          onclick={() => updateParticipationStatus("out")}
  >
    <span><X size="14"/></span>
    <span>{event.participations.out.length ?? 0}</span>
  </button>
</section>