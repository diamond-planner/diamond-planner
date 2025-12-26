<script lang="ts">
  import {CalendarPlus} from "lucide-svelte";
  import {goto, invalidateAll} from "$app/navigation";
  import EventSeriesView from "$lib/dp/components/event/EventSeriesView.svelte";
  import EventForm from "$lib/dp/components/forms/EventForm.svelte";
  import TeamForm from "$lib/dp/components/forms/TeamForm.svelte";
  import TeamGamesModal from "$lib/dp/components/forms/TeamGamesModal.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {EventseriesResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    team: ExpandedTeam;
    eventSeries: EventseriesResponse[];
  }

  const {team, eventSeries}: Props = $props();

  function teamDeleteAction(id: string) {
    goto(`/account`);
    client.collection("teams").delete(id);
    invalidateAll();
  }

  const model = $derived(authSettings.record as CustomAuthModel);
</script>

<h2 class="h3">Admin Section</h2>

<div class="outer-grid">
  <div class="card admin-card preset-outlined-surface-500">
    <div>
      <header class="card-header">
        <h3 class="h4 font-semibold">Games</h3>
      </header>

      <section class="p-4 space-y-3">
        <p class="font-light">
          Events of type "Game" don't need to be created manually.
          All data can be imported from BSM and kept in sync with
          the current state.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="flex flex-col gap-2 lg:gap-3">

        <Dialog triggerClasses="btn preset-tonal-primary border border-primary-500">

          {#snippet triggerContent()}
            <CalendarPlus/>
            <span>Setup Games Import</span>
          {/snippet}

          {#snippet title()}
            <header>
              <h2 class="h3">Games Import Setup for {team.name}</h2>
            </header>
          {/snippet}

          <TeamGamesModal {team}/>
        </Dialog>
      </div>
    </footer>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <div>
      <header class="card-header">
        <h3 class="h4 font-semibold">Event Series</h3>
      </header>

      <section class="p-4 space-y-3">
        <p class="font-light">
          Manage all recurring series of events (mostly of type
          "Practice") that take place on the same day every week.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="flex flex-col gap-2 lg:gap-3">
        <EventSeriesView
          buttonClasses="btn preset-tonal-primary"
          {eventSeries}
          {team}
        />
      </div>
    </footer>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <div>
      <header class="card-header">
        <h3 class="h4 font-semibold">Single Events</h3>
      </header>

      <section class="p-4 space-y-3">
        <p class="font-light">
          Single Events are not connected to any other events and
          don't contain additional logic. Well suited for one-off
          occurrences.
        </p>
      </section>
    </div>
    <footer class="card-footer">
      <div class="flex flex-col gap-2 lg:gap-3">

        <EventForm
          buttonClasses="btn preset-tonal-tertiary border border-tertiary-500"
          clubID={team?.club ?? ""}
          event={null}
          teamID={team.id}
        >
          {#snippet triggerContent()}
            <CalendarPlus/>
            <span>Create Single Event</span>
          {/snippet}
        </EventForm>

      </div>
    </footer>
  </div>

  <div class="card admin-card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="h4 font-semibold">Team Settings</h3>
    </header>

    <section class="p-4 space-y-3">
      <div class="flex flex-col gap-2 lg:gap-3">
        <TeamForm buttonClasses="btn preset-tonal-surface border border-surface-500" club={team?.expand?.club}
                  team={team}/>
      </div>
    </section>

    {#if team?.expand?.club?.admins.includes(model.id)}
      <footer class="card-footer mt-2">
        <div class="flex flex-col gap-2 lg:gap-3 rounded-base preset-outlined-error-500 px-2 py-3">
          <header class="mx-2">Danger Zone</header>

          <DeleteButton
            id={team.id}
            modelName="Team"
            action={teamDeleteAction}
            classes="preset-tonal-error border border-error-500 mx-1"
            buttonText="Delete Team"
          />
        </div>
      </footer>
    {/if}
  </div>
</div>

<style lang="postcss">
  .admin-card {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .outer-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 2);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 48rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr 1fr;
      gap: calc(var(--spacing) * 3);
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr 1fr;
    }
  }
</style>
