<script lang="ts">
  import {
    BookHeart,
    ChartColumnStacked,
    ChartLine,
    CircleUserRound,
    House,
    IdCard,
    LockKeyhole,
    Shield,
    SquareUserRound,
    ToolCase,
    Users,
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
</script>

<nav class="list-nav">
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

        <a href="/account/{authRecord.id}/communityservice" onclick={sheetClose}>
          <ToolCase/>
          <span>Community Service</span>
        </a>
      {/snippet}
    </AccordionItem>

    {#each clubs as club (club.id)}
      <hr class="hr"/>
      <AccordionItem panelPadding="py-0 px-4">
        {#snippet lead()}
          <Shield/>
        {/snippet}

        {#snippet control()}
          <span>{club.name} ({club.acronym})</span>
        {/snippet}

        {#snippet panel()}
          <a href="/account/clubs/{club.id}" onclick={sheetClose}>
            <House/>
            <span>Club Main Page</span>
          </a>

          {#if teams?.length > 0}

            {#each teams.filter(team => team.club === club.id) as team (team.id)}
              <a href="/account/team/{team.id}" onclick={sheetClose}>
                <Users/>
                <span>{team.name} ({team?.expand?.club?.acronym}) </span>
              </a>
            {/each}
          {/if}
        {/snippet}
      </AccordionItem>

      {#if club.admins.includes(authRecord?.id)}
        <hr class="hr"/>
        <AccordionItem panelPadding="py-0 px-4">
          {#snippet lead()}
            <LockKeyhole/>
          {/snippet}

          {#snippet control()}
            Club Administration ({club.acronym})
          {/snippet}

          {#snippet panel()}
            <a href="/account/clubs/{club.id}/admin/attendance" onclick={sheetClose}>
              <ChartColumnStacked/>
              <span>Attendance Dashboard</span>
            </a>

            <a href="/account/clubs/{club.id}/admin/communityservice" onclick={sheetClose}>
              <BookHeart/>
              <span>Community Service</span>
            </a>
          {/snippet}
        </AccordionItem>
      {/if}
    {/each}
  {/if}
</nav>

<style>
  .list-nav {
    padding-inline: var(--spacing);
    padding-block: var(--spacing);

    @media (min-width: 64rem) {
      padding-inline: calc(var(--spacing) * 4);
    }
  }

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

  :global {
    .list-nav {
      .lucide-icon {
        min-width: 24px;
      }
    }
  }
</style>
