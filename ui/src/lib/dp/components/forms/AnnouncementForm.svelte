<script lang="ts">
  import {Plus, SquarePen} from "lucide-svelte";
  import {invalidateAll} from "$app/navigation";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
  import type {AnnouncementsResponse, ClubsResponse, TeamsResponse} from "$lib/dp/types/pb-types.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  interface Props {
    announcement: ExpandedAnnouncement | null;
    club: ClubsResponse | null;
    team: TeamsResponse | null;
    buttonClasses?: string;
    showLabel?: boolean;
  }

  const {
    announcement = null,
    club = null,
    team = null,
    buttonClasses = "",
    showLabel = true,
  }: Props = $props();

  // svelte-ignore state_referenced_locally - we want just the initial value here
  let form: Partial<ExpandedAnnouncement> = $state(
      announcement ?? {
        title: "",
        bodytext: "",
        link: "",
        link_text: "",
        author: authRecord?.id,
        club: club?.id,
        team: team?.id,
      }
  );

  let open = $state(false);

  const isEditing = $derived(announcement !== null);

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: AnnouncementsResponse | null = null;

    try {
      if (form.id) {
        result = await client.collection("announcements").update<AnnouncementsResponse>(form.id, form);
      } else {
        result = await client.collection("announcements").create<AnnouncementsResponse>(form);
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Announcement");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Announcement");
      open = false;
    }
    await invalidateAll();
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses} data-testid="announcement-form-trigger-{isEditing ? 'edit' : 'create'}">
    {#if form.id}
      <SquarePen/>
      {#if showLabel}
        <span>Edit Announcement</span>
      {/if}
    {:else}
      <Plus/>
      {#if showLabel}
        <span>Create new</span>
      {/if}
    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Announcement "{form?.title}"</h2>
      {:else}
        <h2 class="h3">Create new Announcement</h2>
      {/if}
    </header>

    <form class="mt-4 space-y-3" onsubmit={submitForm}>
      <div class="edit-form-grid">
        <input
                autocomplete="off"
                bind:value={form.id}
                class="input"
                name="id"
                readonly
                type="hidden"
        />

        <input
                autocomplete="off"
                class="input"
                name="author"
                readonly
                type="hidden"
                value={form.author}
        />

        {#if club}
          <input
                  autocomplete="off"
                  value={club.id}
                  class="input"
                  name="club"
                  readonly
                  type="hidden"
          />
        {/if}

        {#if team}
          <input
                  autocomplete="off"
                  value={team.id}
                  class="input"
                  name="team"
                  readonly
                  type="hidden"
          />
        {/if}

        <label class="label field-wide">
          <span>Title</span>
          <input
                  bind:value={form.title}
                  class="input"
                  name="title"
                  required
                  type="text"
          />
        </label>

        <label class="label field-wide">
          Announcement Text
          <textarea bind:value={form.bodytext} class="textarea" name="desc" required rows="10"
          ></textarea>
        </label>

        <fieldset class="field-wide border border-surface-200-800 p-3 rounded-base">
          <legend class="legend mb-3">Priority</legend>
          {#each ["info", "warning", "danger"] as prio}
            <label class="label priority-radio-label flex items-center gap-2 my-1">
              <input
                      class="radio"
                      type="radio"
                      name="priority"
                      value={prio}
                      required
                      checked={prio === "info"}
                      bind:group={form.priority}
              />
              {prio}
            </label>
          {/each}
        </fieldset>

        <fieldset class="field-wide border border-surface-200-800 p-3 rounded-base">
          <legend class="legend mb-3">Link Settings</legend>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
            <label class="label">
              <span>Link</span>
              <input
                      bind:value={form.link}
                      class="input"
                      name="link"
                      pattern="https?://.+"
                      placeholder="https://example.com"
                      title="Please enter a valid URL"
                      type="url"
              />
              <span class="text-sm">Single link in case the announcement is used as a call to action.</span>
            </label>

            <label class="label">
              <span>Link Text</span>
              <input
                      bind:value={form.link_text}
                      class="input"
                      name="link_text"
                      placeholder="Click here"
                      type="text"
              />
              <span class="text-sm">If not set, the link itself will be used as a the text.</span>
            </label>
          </div>
        </fieldset>
      </div>

      <hr class="my-5!"/>

      <div class="flex justify-center gap-3">
        <button class="mt-2 btn preset-filled-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>

  </Sheet.Content>
</Sheet.Root>

<style>
    .priority-radio-label {
        text-transform: capitalize;
    }
</style>
