<script lang="ts">
  import type {Match} from "bsm.js";
  import {Ban, Clock} from "lucide-svelte";
  import CommentSection from "$lib/dp/components/comments/CommentSection.svelte";
  import EventAttireSection from "$lib/dp/components/event/EventAttireSection.svelte";
  import EventCoreInfo from "$lib/dp/components/event/EventCoreInfo.svelte";
  import EventPageAdminSection from "$lib/dp/components/event/EventPageAdminSection.svelte";
  import EventParticipantsOverviewSection from "$lib/dp/components/event/EventParticipantsOverviewSection.svelte";
  import EventParticipationSection from "$lib/dp/components/event/EventParticipationSection.svelte";
  import EventTypeBadge from "$lib/dp/components/event/EventTypeBadge.svelte";
  import TimeSection from "$lib/dp/components/event/TimeSection.svelte";
  import MatchDetailLocationCard from "$lib/dp/components/event/match/MatchDetailLocationCard.svelte";
  import MatchTeaserCard from "$lib/dp/components/event/match/MatchTeaserCard.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse} from "$lib/dp/types/pb-types.ts";

  const {data} = $props();

  const event = $derived(data.event);

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canParticipate = $derived(authRecord.teams.includes($event.team));

  //@ts-expect-error - the multi-level expanding trips the typedef up
  const club = $derived($event?.expand?.club) as ClubsResponse;

  const matchJSON = $derived($event?.match_json) as unknown as Match;
</script>

<div class="space-y-4 lg:space-y-6 xl:space-y-7">
  <div class="flex items-center gap-3">
    <h1 class="h1" class:line-through={$event.cancelled}>{$event.title}</h1>
    <div>
      <EventTypeBadge type={$event.type}/>
    </div>
  </div>

  {#if $event.cancelled}
    <span class="badge preset-filled-error-500 gap-1">
      <Ban/>
      Cancelled
    </span>
  {/if}

  <article class="mb-8!" class:line-through={$event.cancelled}>
    <section>
      <p>{$event.desc}</p>
    </section>
  </article>

  <div class="space-y-6" class:line-through={$event.cancelled}>
    <EventCoreInfo event={$event}>
      {#snippet additionalTimeSection()}
        <TimeSection timeValue={$event.endtime} displayText="End:" classes="col-span-2">
          {#snippet icon()}
            <Clock size="18"/>
          {/snippet}
        </TimeSection>
      {/snippet}
    </EventCoreInfo>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2">
    {#if $event.expand.location}
      <section class="mt-3 lg:mt-5">
        <h2 class="h2">Location Details</h2>

        <MatchDetailLocationCard
          field={$event.expand.location}
          classes="my-4 space-y-5"
          showDividers={false}
        />
      </section>
    {/if}

    {#if $event.expand.attire}
      <section class="mt-3 lg:mt-5">
        <EventAttireSection attire={$event.expand.attire}/>
      </section>
    {/if}
  </div>

  {#if !$event.cancelled}
    <hr class="my-8!"/>

    <div class="md:flex md:justify-between md:items-center space-y-2.5 md:space-y-0">
      <h2 class="h4">My Participation</h2>

      {#if canParticipate}
        <EventParticipationSection event={$event} chipClasses="flex-grow"/>
      {:else}
        <div class="flex justify-end">
          <p>Only team members can participate in events.</p>
        </div>
      {/if}
    </div>

    <hr class="my-8"/>

    <EventParticipantsOverviewSection event={$event}/>
  {/if}

  {#if $event.match_json}
    <section>
      <h2 class="h2 mb-3">Game Data</h2>
      <div class="grid grid-cols-1 md:grid-cols-2">
        <MatchTeaserCard match={matchJSON} teamName={$event?.expand?.team?.name}/>
      </div>
    </section>
  {/if}

  <section class="my-6">
    <div class="mt-4 p-3 md:p-4 border border-surface-900-100 rounded-base max-w-[65ch] mx-auto">
      <header>
        <h2 class="h3">Comments</h2>
      </header>
      <CommentSection
        club={club}
        comments={$event?.expand?.comments_via_event ?? []}
        targetID={$event.id}
        targetType="event"
      />
    </div>
  </section>

  {#if $event.expand?.team?.admins.includes(authRecord.id) || $event?.expand?.team?.expand?.club?.admins.includes(authRecord.id)}
    <EventPageAdminSection event={$event}/>
  {/if}
</div>
