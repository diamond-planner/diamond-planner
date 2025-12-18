<script lang="ts">
  import {Plus, SquareArrowOutUpRight} from "lucide-svelte";
  import AnnouncementSectionContent from "$lib/dp/components/announcements/AnnouncementSectionContent.svelte";
  import ClubDetailCard from "$lib/dp/components/club/ClubDetailCard.svelte";
  import TeamListTeaser from "$lib/dp/components/team/TeamListTeaser.svelte";
  import UniformSetInfoCard from "$lib/dp/components/uniformset/UniformSetInfoCard.svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import TeamForm from "$lib/dp/components/forms/TeamForm.svelte";
  import UniformSetForm from "$lib/dp/components/forms/UniformSetForm.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {
    CustomAuthModel,
    ExpandedClub,
    ExpandedTeam,
    ExpandedUniformSet,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type {PageProps} from "./$types";

  const {data}: PageProps = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const club: ExpandedClub = $derived(data.club);
  const teams: ExpandedTeam[] = $derived(data.teams);
  const uniformSets: ExpandedUniformSet[] = $derived(data.uniformSets);
  const announcementStore = $derived(data.announcementStore);
</script>

<svelte:head>
  <title>Club Details</title>
  <meta
    content="Club overview page for the {club.name} with info about teams and uniform sets."
    name="description"
  />
</svelte:head>

<h1 class="h1">{club.name}</h1>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
  <ClubDetailCard {club}/>
</section>

<section class="mt-8! space-y-4">
  <header>
    <h2 class="h2 mb-3">Announcements</h2>
  </header>

  <AnnouncementSectionContent store={announcementStore}/>

  {#if club?.admins.includes(authRecord.id)}
    <AnnouncementForm
      announcement={null}
      team={null}
      {club}
      buttonClasses="btn preset-filled-primary-500"
      showLabel={true}
    />
  {/if}
</section>

<section class="mt-8!">
  <header>
    <h2 class="h2 mb-3">Club Teams</h2>
  </header>

  <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-3">
    {#each teams as team (team.id)}
      <TeamListTeaser {team} link={true}/>
    {/each}

    {#if teams.length === 0}
      <span>This club does not have any teams yet.</span>
    {/if}
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <TeamForm
      team={null}
      {club}
      buttonClasses="btn preset-filled-primary-500"
      showLabel={true}
    />
  {/if}
</section>

<section class="mt-8!">
  <header>
    <h2 class="h2 mb-3">Uniform Sets</h2>
  </header>

  <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-3">
    {#each uniformSets as uniformSet}
      <UniformSetInfoCard {uniformSet}/>
    {/each}
  </div>

  {#if club?.admins.includes(authRecord.id)}
    <Dialog triggerClasses="btn preset-filled-primary-500">
      {#snippet triggerContent()}
        <Plus/>
        <span>Create new</span>
      {/snippet}

      {#snippet title()}
        <header>
          <h2>Create new Uniform Set</h2>
        </header>
      {/snippet}

      <UniformSetForm uniformSet={null} clubID={club.id}/>
    </Dialog>
  {/if}
</section>

<section class="mt-10!">
  <header>
    <h2 class="h2 mb-3">Team Locations</h2>
  </header>

  <a
    class="btn preset-filled-primary-500"
    href="/account/clubs/{club.id}/locations"
  >
    <span>Locations Page</span>
    <SquareArrowOutUpRight size="20"/>
  </a>
</section>

{#if club?.admins.includes(authRecord.id)}
  <hr class="mt-10 mb-6"/>

  <section>
    <header>
      <h2 class="h3 mb-4">Admin Section</h2>
    </header>

    <div
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-2 lg:gap-3"
    >
      <article class="card admin-card preset-outlined-surface-500">
        <header class="card-header">
          <h3 class="h4 font-semibold">Club deletion</h3>
        </header>

        <section class="card-section">
          <p>
            Deleting a club will delete all team, event and participation data.
            For safety reasons, a club can therefore only be deleted by a
            superadmin.
          </p>
          <p>Please contact your club's administration.</p>
        </section>

        <!--        TODO: add email link-->
        <!--        <footer class="card-footer">-->
        <!--          <a-->
        <!--                  class="btn preset-tonal-secondary border border-secondary-500 dark:preset-filled-secondary-500 dark:border grow"-->
        <!--                  href=""-->
        <!--          >-->
        <!--            <Mail/>-->
        <!--            <span class="ms-2">Contact</span>-->
        <!--          </a>-->
        <!--        </footer>-->
      </article>
    </div>
  </section>
{/if}

<style>
  .card-section {
    padding: calc(var(--spacing) * 4);

    p:first-child {
      margin-block-end: calc(var(--spacing) * 2);
    }
  }
</style>
