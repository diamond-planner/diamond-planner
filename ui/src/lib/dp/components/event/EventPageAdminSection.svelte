<script lang="ts">
  import {CalendarArrowDown, CalendarPlus, Edit, Info} from "lucide-svelte";
  import {goto} from "$app/navigation";
  import EventForm from "$lib/dp/components/forms/EventForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import type {ExpandedEvent} from "$lib/dp/types/ExpandedResponse.ts";
  import type {EventsUpdate} from "$lib/dp/types/pb-types.ts";

  interface Props {
    event: ExpandedEvent;
  }

  const {event}: Props = $props();

  const guestPlayerForm = $state({
    name: "",
  });

  async function submitNewGuestPlayer(e: SubmitEvent) {
    e.preventDefault();

    await client.collection("events").update<EventsUpdate>(event.id, {
      guests:
          event.guests.length === 0 ? guestPlayerForm.name : event.guests + "," + guestPlayerForm.name,
    });
    guestPlayerForm.name = "";
  }

  async function deleteEvent(id: string) {
    await client.collection("events").delete(id);
    await goto(`/account/team/${event.team}`);
  }
</script>

<hr class="mt-6"/>

<h2 class="h2 my-4">Admin Section</h2>

<div
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2 lg:gap-3"
>
  <div class="card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="h4 font-semibold">Event Details</h3>
    </header>

    <section class="p-4 space-y-2">
      <div class="flex items-center gap-3">
        <Info/>
        <div>
          <p class="text-sm font-light">BSM event ID</p>
          <p>{event.bsm_id}</p>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <CalendarPlus/>
        <div>
          <p class="text-sm font-light">Created</p>
          <p>{new Date(event.created).toLocaleString()}</p>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <CalendarArrowDown/>
        <div>
          <p class="text-sm font-light">Last Updated</p>
          <p>{new Date(event.updated).toLocaleString()}</p>
        </div>
      </div>
    </section>
  </div>

  <div class="card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="h4 font-semibold">Add Guest player</h3>
    </header>

    <section class="p-4 space-y-2">
      <form
              onsubmit={submitNewGuestPlayer}
      >
        <label class="label">
          Name
          <input
                  bind:value={guestPlayerForm.name}
                  class="input mt-1"
                  placeholder="Enter the guest player's name"
                  type="text"
          />
        </label>
        <button class="mt-3 btn preset-tonal border border-surface-500" type="submit">Submit
        </button
        >
      </form>
    </section>
  </div>

  <div class="card preset-outlined-surface-500">
    <header class="card-header">
      <h3 class="h4 font-semibold">Actions</h3>
    </header>

    <section class="p-4 space-y-3">
      <div class="flex flex-col gap-2 lg:gap-3">

        <EventForm
                buttonClasses="btn preset-tonal-surface border border-surface-500"
                clubID={event?.expand?.team?.club ?? ""}
                event={event}
                teamID={event.team}
        >
          {#snippet triggerContent()}
            <Edit/>
            <span>Edit Event</span>
          {/snippet}
        </EventForm>

        <DeleteButton
                action={deleteEvent}
                buttonText="Delete Event"
                classes="preset-tonal-error border border-error-500"
                id={event.id}
                modelName="Event"
        />
      </div>
    </section>
  </div>
</div>