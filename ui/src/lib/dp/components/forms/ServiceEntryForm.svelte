<script lang="ts">
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedServiceEntry} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import type {
    ClubsResponse,
    ServiceentriesResponse,
    ServiceentriesUpdate,
    UsersResponse
  } from "$lib/dp/types/pb-types.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {invalidate} from "$app/navigation";
  import Flatpickr from "$lib/dp/components/formElements/Flatpickr.svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";

  const authRecord = authSettings.record as CustomAuthModel;

  interface Props {
    serviceEntry: ExpandedServiceEntry | null;
    club: ClubsResponse;
  }

  let {serviceEntry, club}: Props = $props();

  type MinuteSuggestion = {
    value: number,
    label: string,
  }
  const minutesSuggestions: MinuteSuggestion[] = [
    {value: 30, label: "0.5h"},
    {value: 60, label: "1h"},
    {value: 90, label: "1.5h"},
    {value: 120, label: "2h"},
    {value: 150, label: "2.5h"},
    {value: 180, label: "3h"},
    {value: 210, label: "3.5h"},
    {value: 240, label: "4h"},
  ];

  function setSuggestion(suggestion: MinuteSuggestion) {
    form.minutes = suggestion.value;
  }

  const form: ServiceentriesUpdate = $state(serviceEntry ?? {
    id: "",
    service_date: "",
    minutes: 0,
    member: authRecord.id,
    club: club.id,
    title: "",
    description: "",
    board_member: "",
  });

  const isAdmin = $derived(club?.admins.includes(authRecord.id));

  const possibleClubs = client.collection(Collection.Clubs).getFullList<ClubsResponse>({
    requestKey: null,
  });

  const allClubMembers = client.collection(Collection.Users).getFullList<UsersResponse>({
    filter: `club ?~ '${club.id}'`,
    requestKey: null,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ServiceentriesResponse | null = null;

    try {
      if (form.id) {
        result = await client.collection(Collection.ServiceEntries).update(form.id, form);
      } else {
        result = await client.collection(Collection.ServiceEntries).create(form);
      }
    } catch (e) {
      toastController.triggerGenericFormErrorMessage("Community Service Entry");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Community Service Entry");
      closeModal();
      await invalidate("communityservice:list");
    }
  }
</script>

<form class="edit-form" onsubmit={submitForm}>

  <label class="label">
    Title
    <input
      bind:value={form.title}
      class="input"
      name="title"
      required
      type="text"
    />
  </label>

  <label class="label">
    Date
    <Flatpickr
      bind:value={form.service_date}
      options={Object.assign(DateTimeUtility.datePickerOptionsNoTime, {
                      static: true, // render the picker as a child element to the form to work in a sheet portal context
                  })}
    />
  </label>

  <fieldset class="rounded-base">
    <legend class="legend">Duration</legend>

    <label class="label">
      Minutes
      <input
        autocomplete="off"
        bind:value={form.minutes}
        class="input"
        name="number"
        type="number"
        required
      />
    </label>

    <div class="suggestion-buttons">
      {#each minutesSuggestions as suggestion}
        <button
          type="button"
          class="chip preset-tonal-primary"
          onclick={() => setSuggestion(suggestion)}
        >
          {suggestion.label}
        </button>
      {/each}
    </div>
  </fieldset>

  <label class="label">
    Description
    <textarea
      bind:value={form.description}
      class="textarea"
      name="desc"
      rows="2"
    ></textarea>
  </label>

  {#if !form.id && isAdmin}
    <label class="label">
      Club
      <select
        bind:value={form.club}
        class="select"
        name="club"
        required
      >
        {#await possibleClubs then clubs}
          {#each clubs as club}
            <option value={club.id}>{club.name}</option>
          {/each}
        {/await}
      </select>
    </label>
  {/if}

  <label class="label">
    Board Member
    <select
      bind:value={form.board_member}
      class="select"
      name="board_member"
      required
    >
      {#await allClubMembers then members}
        {#each members.filter((element) => club?.admins.includes(element.id)) as member}
          <option value={member.id}>{member.first_name} {member.last_name}</option>
        {/each}
      {/await}
    </select>
  </label>

  {#if isAdmin}
    <label class="label">
      Member
      <select
        bind:value={form.member}
        class="select"
        name="member"
        required
      >
        {#await allClubMembers then members}
          {#each members as member}
            <option value={member.id}>{member.first_name} {member.last_name}</option>
          {/each}
        {/await}
      </select>
    </label>
  {/if}

  <div class="submit-wrapper">
    <button
      class="btn preset-tonal-primary"
      type="submit"
    >Confirm
    </button>
  </div>
</form>

<style>
  .edit-form {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 2);
  }

  fieldset {
    border: 1px solid var(--color-surface-200-800);
    padding: calc(var(--spacing) * 3);
  }

  .suggestion-buttons {
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 2);
    margin-block-start: calc(var(--spacing) * 3);
  }

  .submit-wrapper {
    display: flex;
    justify-content: flex-end;
    gap: calc(var(--spacing) * 3);
    margin-block-start: calc(var(--spacing) * 3);
  }
</style>
