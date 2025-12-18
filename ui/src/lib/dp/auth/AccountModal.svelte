<script lang="ts">
  import {goto} from "$app/navigation";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import {authSettings, client} from "../client.svelte.js";

  async function logout() {
    closeModal();
    authSettings.record = null;
    client.authStore.clear();
    await goto("/login");
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

<article class="py-1">
  <div class="flex mt-1 mb-5 justify-between items-center">
    <p>Logged in as:</p>

    <a class="badge preset-filled-primary-500" href="/account" onclick={closeModal}>
      <samp>{`${authRecord?.first_name ?? ""} ${authRecord?.last_name ?? ""}`}</samp>
    </a>
  </div>

  <div class="flex mt-1 mb-5 justify-between items-start">
    <p>Club:</p>

    <div class="flex flex-wrap gap-2 justify-end">
      {#each authRecord?.expand?.club as club}
        <a href="/account/clubs/{club.id}" class="badge preset-filled-tertiary-500" onclick={closeModal}>
          {club.name}
        </a>
      {/each}
    </div>
  </div>

  <button class="btn preset-tonal-primary border border-primary-500 mb-1" onclick={logout}>Log out</button>
</article>
