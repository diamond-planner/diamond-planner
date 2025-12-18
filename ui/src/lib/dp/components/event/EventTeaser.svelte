<script lang="ts">
  import {Ban} from "lucide-svelte";
  import EventTypeBadge from "$lib/dp/components/event/EventTypeBadge.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import EventCoreInfo from "./EventCoreInfo.svelte";
  import EventParticipationSection from "./EventParticipationSection.svelte";

  interface props {
    event: ExpandedEvent;
    link: boolean;
  }

  const {event, link}: props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  // club admins can see the team and change settings, but not participate
  const canParticipate = $derived(authRecord.teams.includes(event.team));
</script>

<article
        class="card preset-tonal-surface shadow-xl text-sm h-full"
        class:card-hover={link}
>
  <a class:line-through={event.cancelled} href="/account/event/{event.id}">
    <header class="card-header">
      <h2>
        <EventTypeBadge type={event.type}/>
        <span class="ms-1 font-bold h5">{event?.title}</span>
      </h2>
    </header>

    <EventCoreInfo classes={"p-4"} {event}/>
  </a>

  <footer class="card-footer">
    <hr class="my-2"/>
    {#if event.cancelled}
      <div class="flex justify-end">
                <span class="badge preset-filled-error-500 gap-1">
                    <Ban size="18"/>
                    Cancelled
                </span>
      </div>
    {:else}
      {#if canParticipate}
        <EventParticipationSection {event}/>
      {:else}
        <div class="flex justify-end">
          <p>Only team members can participate in events.</p>
        </div>
      {/if}
    {/if}
  </footer>
</article>
