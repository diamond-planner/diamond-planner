<script lang="ts">
  import {CalendarPlus} from "lucide-svelte";
  import {slide} from "svelte/transition";
  import {invalidateAll} from "$app/navigation";
  import EventSeriesListItem from "$lib/dp/components/eventseries/EventSeriesListItem.svelte";
  import EventSeriesForm from "$lib/dp/components/forms/EventSeriesForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {client} from "$lib/dp/client.svelte.ts";
  import type {EventSeriesCreationData, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {EventseriesResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    team: ExpandedTeam;
    eventSeries: EventseriesResponse[];
    buttonClasses?: string;
  }

  const {team, eventSeries, buttonClasses = ""}: Props = $props();

  let open = $state(false);
  let showForm = $state(false);

  let selectedEventSeries: EventSeriesCreationData | null = $state(null);
  let eventSeriesFormContainer: HTMLDivElement;

  function setupAndShowForm(eventSeries: EventSeriesCreationData | null) {
    selectedEventSeries = eventSeries;
    showForm = true;

    // wait for Svelte transition to finish
    setTimeout(() => {
      eventSeriesFormContainer?.scrollIntoView({
        behavior: "smooth",
        block: "start",
        inline: "nearest",
      });
    }, 230);
  }

  async function deleteEventSeries(id: string) {
    const success = await client.collection("eventseries").delete(id);

    if (success) {
      await invalidateAll();
      open = false;
      showForm = false;
    }
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    <CalendarPlus/>
    <span>Manage Event Series</span>
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      <h2 class="h3">Manage Event Series for {team?.name}</h2>
    </header>

    <div class="space-y-4">

      <h3 class="h4 mt-6!">Active Event Series</h3>

      {#each eventSeries as series (series.id)}
        <EventSeriesListItem eventSeries={series}/>

        <div class="flex gap-2">
          <button class="btn preset-tonal border border-surface-500" onclick={() => setupAndShowForm(series)}>Edit
          </button>

          <DeleteButton
                  id={series.id}
                  modelName="Event Series"
                  action={deleteEventSeries}
                  classes="preset-tonal-error border border-error-500"
                  buttonText="Delete"
          />
        </div>
      {/each}
    </div>

    {#if eventSeries.length === 0}
      <section class="font-light space-y-4">
        <div>No event series have been set up for this team yet.</div>
      </section>
    {/if}
    <button class="btn preset-filled-primary-500 mt-6!" onclick={() => setupAndShowForm(null)}>
      Add new Event Series
    </button>

    <hr class="my-6!">

    <div bind:this={eventSeriesFormContainer}>
      {#if showForm}
        <div transition:slide={{duration: 230}}>
          <EventSeriesForm team={team} eventSeries={selectedEventSeries} bind:showForm={showForm}/>
        </div>
      {/if}
    </div>
  </Sheet.Content>
</Sheet.Root>