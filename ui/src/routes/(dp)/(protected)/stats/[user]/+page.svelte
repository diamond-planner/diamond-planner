<script lang="ts">
  import {goto} from "$app/navigation";
  import AttendanceTotalStatsBlock from "$lib/dp/components/stats/AttendanceTotalStatsBlock.svelte";
  import StatsByTypePieChart from "$lib/dp/components/stats/StatsByTypePieChart.svelte";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";
  import type {PersonalAttendanceStatsItem} from "$lib/dp/types/PersonalAttendanceStats.ts";
  import {range} from "$lib/dp/utility/range.ts";
  import SeasonSelector from "$lib/dp/components/utils/SeasonSelector.svelte";

  let {data} = $props();

  const user = $derived(data.user);

  const currentYear = new Date().getFullYear();
  const seasonOptions = range(2023, currentYear);

  let selectedSeason = $state(currentYear);

  const reloadWithQuery = () => {
    let queryString = `?season=${selectedSeason}`;

    goto(queryString, {
      noScroll: true,
    });
  };
</script>

<h1 class="h1">Stats for {user?.first_name} {user?.last_name}</h1>

<SeasonSelector bind:selectedSeason={selectedSeason} onChangeCallback={reloadWithQuery} seasonOptions={seasonOptions}/>

{#snippet statsSection(statsItem: PersonalAttendanceStatsItem)}
  <h2 class="h2 my-4">{statsItem.teamName}</h2>
  <section>
    <StatsByTypePieChart statsItem={statsItem}/>
  </section>

  <section>
    {#each statsItem.attendanceTotals as attendance}
      <AttendanceTotalStatsBlock
              season={statsItem.season}
              attendance={attendance}
      />
    {/each}
  </section>
{/snippet}

{#await data.statsItem}
  <ProgressRing/>
{:then statsItem}

  {#if statsItem}
    {@render statsSection(statsItem)}
  {/if}

{:catch error}
  <p>error loading: {error.message}</p>
{/await}

{#await Promise.all(data.teamStatsItems)}
  <ProgressRing/>
{:then teamStatsItems}

  {#each teamStatsItems as teamStatsItem}
    <div class="stats-section-block">
      {@render statsSection(teamStatsItem)}
    </div>
  {/each}

{:catch error}
  <p>error loading: {error.message}</p>
{/await}

<style>
  section {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block-end: calc(var(--spacing) * 3);

    @media (min-width: 64rem) {
      grid-template-columns: repeat(2, 1fr);
    }

    @media (min-width: 80rem) {
      grid-template-columns: repeat(3, 1fr);
    }
  }

  .stats-section-block {
    margin-block-start: calc(var(--spacing) * 8) !important;
  }
</style>
