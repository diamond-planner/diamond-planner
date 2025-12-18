<script lang="ts">
  import {browser} from "$app/environment";
  import LoginForm from "$lib/dp/auth/LoginForm.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.ts";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

{#if browser}
  {#if client.authStore.isValid}
    <h1 class="h1">Login Page</h1>
    <p>You are logged in as:
      <span class="mx-2 badge preset-filled-primary-500">
        {authRecord.first_name} {authRecord.last_name}
      </span>
      <br>
      You can log out by clicking your avatar in the top right.
    </p>
    <a class="btn preset-tonal-primary border border-primary-500" href="/account">Go to Account Page</a>
  {:else }
    <LoginForm/>
  {/if}
{/if}

<style>
    p {
        margin-block: 2em;
    }
</style>