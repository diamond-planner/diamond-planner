<script lang="ts">
  import type {AuthProviderInfo, RecordAuthResponse, RecordModel, RecordService} from "pocketbase";
  import {goto} from "$app/navigation";
  import {providerLogin} from "$lib/dp/auth/Auth.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";

  interface Props {
    authProvider: AuthProviderInfo;
    collection: RecordService<RecordModel>;
    signup_key?: string;
    disabled: boolean;
  }

  const {authProvider, collection, signup_key = "", disabled}: Props = $props();

  async function submitOAuthRequest(provider: AuthProviderInfo) {
    let authResponse: RecordAuthResponse<RecordModel> | undefined;
    try {
      authResponse = await providerLogin(provider, collection, signup_key);
    } catch (error) {
      console.error(error);
      toastController.trigger({
        message:
            "There was an error processing your authentication request via external provider. Please double-check your signup key",
        background: "preset-filled-error-500",
      });
    }

    if (authResponse) {
      await goto("/account", {invalidateAll: true});
    }
  }

  const isGenericProvider = $derived(
      authProvider.name !== "google" &&
      authProvider.name !== "apple" &&
      authProvider.name !== "github" &&
      authProvider.name !== "discord"
  );
</script>

<!--
@component Shows a button to log in with an external service. Uses branded image if available,
falls back to simple button if not.
-->

<button
        aria-label="sign in with {authProvider.displayName}"
        class={['oauth-button', isGenericProvider && 'btn preset-tonal-tertiary border border-tertiary-500']}
        disabled={disabled}
        onclick={() => submitOAuthRequest(authProvider)}
        type="button"
>
  {#if (authProvider.name === "google")}
    <img class="light-logo" src="/providers/Google_light.svg" alt="Google logo" width="40">
    <img class="dark-logo" src="/providers/Google_dark.svg" alt="Google logo" width="40">

  {:else if authProvider.name === "apple"}
    <img class="light-logo" src="/providers/siwa_apple_light.svg" alt="Apple logo" width="40">
    <img class="dark-logo" src="/providers/siwa_apple_dark.svg" alt="Apple logo" width="40">

  {:else if authProvider.name === "github"}
    <img class="light-logo" src="/providers/github-mark-white.svg" alt="GitHub logo" width="40">
    <img class="dark-logo" src="/providers/github-mark.svg" alt="GitHub logo" width="40">

  {:else if authProvider.name === "discord"}
    <img src="/providers/discord_white_in_blue_circle.svg" alt="Discord logo" width="40">

  {:else }
    <span>{authProvider.displayName}</span>
  {/if}
</button>

<!-- svelte-ignore css_unused_selector -->
<style>
    .oauth-button[disabled] {
        cursor: not-allowed;
        opacity: 0.8;
        filter: grayscale(1);
    }
</style>
