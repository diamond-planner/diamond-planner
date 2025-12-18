<script lang="ts">
  import {
    ChartColumnStacked,
    ChartLine,
    CircleUserRound,
    IdCard,
    LockKeyhole,
    Shield,
    SquareUserRound,
    Users,
    UsersRound,
  } from "lucide-svelte";
  import AccordionItem from "$lib/dp/components/modal/AccordionItem.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedClub, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    clubs: ExpandedClub[];
    teams: ExpandedTeam[];
    sheetOpen?: boolean;
  }

  let {clubs, teams, sheetOpen = $bindable()}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const isUserAuthenticated = $derived(!!authSettings.record);

  function sheetClose(): void {
    if (sheetOpen) {
      sheetOpen = false;
    }
  }

  const value = $state(["account", "teamClub", "admin"]);
</script>

<nav class="list-nav py-1 px-1 lg:px-4">
  {#if isUserAuthenticated}
    <AccordionItem>
      {#snippet lead()}
        <SquareUserRound/>
      {/snippet}

      {#snippet control()}
        <span>My Account</span>
      {/snippet}

      {#snippet panel()}
        <a href="/account" onclick={sheetClose}>
          <CircleUserRound/>
          <span>Dashboard</span>
        </a>

        <a href="/stats/{authRecord?.id}" onclick={sheetClose}>
          <ChartLine/>
          <span>Personal Stats</span>
        </a>

        <a href="/account/playerprofile" onclick={sheetClose}>
          <IdCard/>
          <span>Player Profile</span>
        </a>
      {/snippet}
    </AccordionItem>

    <hr class="hr"/>

    <AccordionItem panelPadding="py-0 px-4">
      {#snippet lead()}
        <UsersRound/>
      {/snippet}

      {#snippet control()}
        <span>My Clubs & Teams</span>
      {/snippet}

      {#snippet panel()}
        {#each clubs as club (club.id)}
          <a href="/account/clubs/{club.id}" onclick={sheetClose}>
            <Shield/>
            <span>{club.name} ({club.acronym})</span>
          </a>
        {/each}

        {#if teams?.length > 0}
          <hr class="hr"/>

          {#each teams as team (team.id)}
            <a href="/account/team/{team.id}" onclick={sheetClose}>
              <Users/>
              <span>{team.name} ({team?.expand?.club?.acronym}) </span>
            </a>
          {/each}
        {/if}
      {/snippet}
    </AccordionItem>

    <hr class="hr"/>

    <AccordionItem panelPadding="py-0 px-4">
      {#snippet lead()}
        <LockKeyhole/>
      {/snippet}

      {#snippet control()}
        Administration
      {/snippet}

      {#snippet panel()}
        <a href="/stats/admin" onclick={sheetClose}>
          <ChartColumnStacked/>
          <span>Admin Dashboard</span>
        </a>
      {/snippet}
    </AccordionItem>
  {/if}
</nav>

<style lang="postcss">
    a {
        display: flex;
        align-items: center;
        margin: 0.3rem;
        padding: 0.4rem;
    }

    a:hover,
    a:focus {
        background-color: var(--color-primary-50-950);
        color: var(--color-primary-950-50);
        border-radius: var(--radius-container);
    }

    span {
        margin-left: 0.5rem;
    }
</style>
