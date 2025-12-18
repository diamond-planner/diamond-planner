<script lang="ts">
  // @ts-expect-error
  import {Tabs} from "bits-ui";
  import {Users} from "lucide-svelte";
  import {goto} from "$app/navigation";
  import AnnouncementSectionContent from "$lib/dp/components/announcements/AnnouncementSectionContent.svelte";
  import EventTeaser from "$lib/dp/components/event/EventTeaser.svelte";
  import TeamAdminSection from "$lib/dp/components/team/TeamAdminSection.svelte";
  import TeamTeaserCard from "$lib/dp/components/team/TeamTeaserCard.svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, EventType} from "$lib/dp/types/ExpandedResponse.js";
  import Paginator from "$lib/dp/utility/Paginator.svelte";
  import type {PageProps} from "./$types";

  const {data}: PageProps = $props();
  const events = $derived(data.events);
  const currentPage = $derived($events.page);
  const announcementStore = $derived(data.announcementStore);

  let showEvents = $state("next");
  let sorting: "asc" | "desc" | string = $state("asc");
  let showTypes: EventType | "any" | string = $state("any");

  const reloadWithQuery = () => {
    const queryString = `?timeframe=${showEvents}&page=${currentPage}&sort=${sorting}&type=${showTypes}`;

    goto(queryString, {
      noScroll: true,
    });
  };

  $effect.pre(() => {
    console.log(showEvents);
    reloadWithQuery();
  });

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(
    data.team.admins.includes(authRecord?.id) ||
    data.team?.expand?.club?.admins.includes(authRecord?.id)
  );
</script>

<h1 class="h1 my-3!">{data.team.name} ({data.team?.expand?.club.name})</h1>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 my-6!">
  <article class="card preset-tonal-surface lg:col-span-2">
    <header class="card-header">
      <h2 class="h4 font-medium">Team Description</h2>
    </header>
    <section class="p-4">{@html data.team.description}</section>
  </article>

  <TeamTeaserCard link={false} team={data.team}/>
</div>

<section class="my-8! space-y-4">
  <header>
    <h2 class="h2 mb-3">Announcements</h2>
  </header>

  <AnnouncementSectionContent store={announcementStore}/>

  {#if canEdit}
    <AnnouncementForm
      announcement={null}
      team={data.team}
      club={null}
      buttonClasses="btn preset-filled-primary-500"
      showLabel={true}
    />
  {/if}
</section>

<section>
  <header>
    <h2 class="h2">Team Events</h2>
  </header>

  <div
    class="flex flex-wrap gap-4 preset-tonal-surface justify-between px-4 py-3 rounded-base text-sm lg:text-base"
  >
    <label
      class="flex items-center gap-2 grow justify-between xl:justify-start md:grow-0"
    >
      <span>Timeframe</span>
      <Tabs.Root bind:value={showEvents}>
        <Tabs.List class="tabs-list event-segment-container p-1!">
          <Tabs.Trigger class="tabs-trigger btn active:preset-filled-error-300-700" data-testid="segment-item"
                        value="next">Next
          </Tabs.Trigger>
          <Tabs.Trigger class="tabs-trigger btn active:preset-filled" data-testid="segment-item" value="past">Past
          </Tabs.Trigger>
        </Tabs.List>
      </Tabs.Root>
    </label>

    <label class="flex items-center gap-2 justify-between xl:justify-start grow md:grow-0">
      <span>Sort</span>
      <Tabs.Root bind:value={sorting}>
        <Tabs.List class="tabs-list flex-wrap event-segment-container p-1!">
          <Tabs.Trigger class="tabs-trigger btn" data-testid="segment-item" value="asc">Ascending</Tabs.Trigger>
          <Tabs.Trigger class="tabs-trigger btn" data-testid="segment-item" value="desc">Descending</Tabs.Trigger>
        </Tabs.List>
      </Tabs.Root>
    </label>

    <label class="flex items-center gap-2 justify-between xl:justify-start grow md:grow-0">
      <span>Type</span>
      <Tabs.Root bind:value={showTypes}>
        <Tabs.List class="tabs-list flex-wrap event-segment-container p-1! gap-1">
          <Tabs.Trigger class="tabs-trigger btn" data-testid="segment-item" value="any">All</Tabs.Trigger>
          <Tabs.Trigger class="tabs-trigger btn" data-testid="segment-item" value="game">Game</Tabs.Trigger>
          <Tabs.Trigger class="tabs-trigger btn" data-testid="segment-item" value="practice">Practice</Tabs.Trigger>
          <Tabs.Trigger class="tabs-trigger btn" data-testid="segment-item" value="misc">Other</Tabs.Trigger>
        </Tabs.List>
      </Tabs.Root>
    </label>
  </div>

  <div class="events-grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3 lg:gap-4 2xl:gap-5">
    {#each $events.items as event (event.id)}
      <div class="">
        <EventTeaser {event} link={true}/>
      </div>
    {:else}
      <p>No events available with the current filters.</p>
    {/each}
  </div>
</section>

<Paginator showIfSinglePage={false} store={events}/>

<hr class="my-8!"/>
<section class="space-y-2 lg:space-y-4">
  <header>
    <h2 class="h3">Links</h2>
  </header>

  <div class="flex flex-wrap items-center gap-2 lg:gap-3">
    <a
      class="btn preset-tonal-tertiary border border-tertiary-500"
      href="/account/team/{data.team.id}/members"
    >
      <Users/>
      <span>Player List</span>
    </a>
  </div>
</section>

{#if canEdit}
  <hr class="my-2"/>

  <TeamAdminSection team={data.team} eventSeries={data.eventSeries}/>
{/if}

<style>
  :global {
    .event-segment-container {
      border: 1px solid var(--color-surface-600-400);

      .tabs-trigger {
        padding: 0.25rem 0.6rem;
      }
    }
  }

  .events-grid {
    margin-block: calc(var(--spacing) * 4);
    display: grid;
  }

  header {
    margin-block: calc(var(--spacing) * 3);
  }

  section {
    margin-block: calc(var(--spacing) * 4);
  }
</style>
