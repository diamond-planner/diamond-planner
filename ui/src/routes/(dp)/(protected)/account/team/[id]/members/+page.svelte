<script lang="ts">
  import TeamMemberDatatable from "$lib/dp/components/team/TeamMemberDatatable.svelte";
  import TeamAddMemberButton from "$lib/dp/components/team/TeamAddMemberButton.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";

  const {data} = $props();

  const players = $derived(data.players);

  const showAdminSection = $derived.by(() => {
    if (!authSettings?.record?.id) {
      return false;
    }
    if (data.team?.expand?.club?.admins.includes(authSettings?.record.id)) {
      return true;
    }
    return data.team.admins.includes(authSettings?.record?.id);
  });
</script>

<svelte:head>
  <title>Team Members</title>
  <meta content="Team member list for {data.team.name}." name="description"/>
</svelte:head>

<h1 class="h1">Team members for Team "{data.team.name}"</h1>

<div class="mt-8!">
  <TeamMemberDatatable data={$players.items} {showAdminSection} team={data.team}/>
</div>

{#if showAdminSection}
  <TeamAddMemberButton team={data.team} club={data.team?.expand?.club}/>
{/if}