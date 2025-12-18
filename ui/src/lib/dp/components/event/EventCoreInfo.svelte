<script lang="ts">
  import {Calendar, Clock, MapPin} from "lucide-svelte";
  import type {Snippet} from "svelte";
  import TimeSection from "$lib/dp/components/event/TimeSection.svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";

  interface props {
    event: ExpandedEvent;
    classes?: string;
    additionalTimeSection?: Snippet | undefined;
  }

  const {event, classes = "", additionalTimeSection = undefined}: props = $props();

  const startTime = $derived(new Date(event.starttime));
</script>

<section class={classes}>
  <div class="grid grid-cols-6 gap-4">
    <div class="flex items-center gap-2 col-span-6">
      <Calendar size="18"/>
      <time class="font-bold" datetime="{event.starttime}">
        {startTime.toLocaleDateString("de-DE", DateTimeUtility.eventDateFormat)}
      </time>
    </div>

    <TimeSection
            classes={additionalTimeSection ? "col-span-2" : "col-span-3"}
            displayText="Meet:"
            timeValue={event.meetingtime ? event.meetingtime : event.starttime}
    >
      {#snippet icon()}
        <Clock size="18"/>
      {/snippet}
    </TimeSection>

    <TimeSection
            classes={additionalTimeSection ? "col-span-2" : "col-span-3"}
            displayText="Start:"
            timeValue={event.starttime}
    >
      {#snippet icon()}
        <Clock size="18"/>
      {/snippet}
    </TimeSection>

    {@render additionalTimeSection?.()}

    <div class="flex items-center gap-2 col-span-6">
      <MapPin size="18"/>
      {#if event?.expand?.location}
        <p>
          {event?.expand?.location.address_addon}, {event?.expand?.location
            .city}
        </p>
      {:else}
        <p>No location provided.</p>
      {/if}
    </div>
  </div>
</section>