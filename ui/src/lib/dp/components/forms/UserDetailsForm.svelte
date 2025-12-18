<script lang="ts">
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  const authRecord = authSettings.record as CustomAuthModel;

  let files: FileList | null = $state(null);

  let form = $state({
    id: authRecord.id,
    firstName: authRecord.first_name,
    lastName: authRecord.last_name,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    const formData = new FormData();

    formData.append("id", form.id);

    if (form.firstName) {
      formData.append("first_name", form.firstName);
    }

    if (form.lastName) {
      formData.append("last_name", form.lastName);
    }

    if (files) {
      for (const file of files) {
        formData.append("avatar", file);
      }
    }

    client.collection("users").update(form.id, formData, {expand: "club"}).then(closeModal);
  }
</script>

<p class="mt-2 font-light">
  This data is only used for internal purposes and never shown anywhere publicly.
  (Your team and club administrators have access).
</p>

<form class="mt-4 space-y-3" onsubmit={submitForm}>
  <input
          autocomplete="off"
          bind:value={form.id}
          class="input"
          name="id"
          readonly
          type="hidden"
  />

  <label class="label">
    First Name(s)
    <input autocomplete="given-name" bind:value={form.firstName} class="input" name="firstname"/>
  </label>

  <label class="label">
    Last Name
    <input autocomplete="family-name" bind:value={form.lastName} class="input" name="lastname"/>
  </label>

  <label class="label">
    Profile Image
    <input
            accept="image/*"
            bind:files
            class="input"
            name="avatar"
            type="file"
    />
  </label>

  {#if files}
    <h2>Selected file:</h2>
    {#each Array.from(files) as file}
      <p>{file.name} ({file.size} bytes)</p>
    {/each}
  {/if}

  <p class="font-light">
    This is your app profile picture. It will not be used for your player profile image.
  </p>

  <div class="flex justify-end gap-3 mt-3">
    <button class="mt-2 btn preset-tonal-primary border border-primary-500" type="submit">Confirm</button>
  </div>
</form>
