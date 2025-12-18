<script lang="ts">
  import {LucideLogIn} from "lucide-svelte";
  import Avatar from "$lib/dp/components/utils/Avatar.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import AccountModal from "$lib/dp/auth/AccountModal.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import {client} from "../client.svelte.js";

  const {signupAllowed = true} = $props();
</script>

{#if authSettings.record && client.authStore.isValid}
  <Dialog triggerClasses="badge">
    {#snippet triggerContent()}
      <Avatar
              fallback={authSettings.record?.first_name.charAt(0).toUpperCase() +
          authSettings.record?.last_name.charAt(0).toUpperCase()}
              --size="3.5rem"
              src={client.files.getURL(authSettings.record ?? {}, authSettings.record?.avatar)}
      />
    {/snippet}

    {#snippet title()}
      <h2>Account Overview</h2>
    {/snippet}

    <AccountModal/>
  </Dialog>
{:else}
  <a
          class="btn preset-tonal-primary border border-primary-500 flex items-center gap-2"
          href="/login"
          title={signupAllowed ? "Login / Register" : "Login"}
  >
    <LucideLogIn/>
    Login
  </a>
{/if}
