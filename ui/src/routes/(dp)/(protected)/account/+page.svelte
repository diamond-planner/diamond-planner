<script lang="ts">
  import TeamListTeaser from "$lib/dp/components/team/TeamListTeaser.svelte";
  import PlayerProfileClubsSection from "$lib/dp/components/user/PlayerProfileClubsSection.svelte";
  import UserDataCard from "$lib/dp/components/user/UserDataCard.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const {data} = $props();
</script>

<h1 class="h1">My Dashboard</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-6">
  <UserDataCard model={authRecord}/>

  <PlayerProfileClubsSection clubs={data.clubs}/>
</div>

<h2 class="h2">My Teams</h2>

<div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-6">
  {#each data.teams as team}
    <TeamListTeaser {team} link={true}/>
  {/each}

  {#if data.teams?.length === 0}
    <p>You are not a member of any teams yet.</p>
  {/if}
</div>

<style>
    h2 {
        margin-block: calc(var(--spacing) * 4);
    }
</style>