<script lang="ts">
  import type {PageProps} from "./$types";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";
  import {range} from "$lib/dp/utility/range.ts";
  import {goto} from "$app/navigation";
  import SeasonSelector from "$lib/dp/components/utils/SeasonSelector.svelte";
  import ClubCommunityServiceSection from "$lib/dp/components/communityservice/ClubCommunityServiceSection.svelte";

  let {data}: PageProps = $props();

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

<h1 class="h1">
  Community Service Admin Dashboard
</h1>
<p>for <strong>{data.club.name}</strong></p>

{#await data.rows}
  <ProgressRing/>
{:then rows}
  <section>
    <SeasonSelector
      bind:selectedSeason={selectedSeason}
      onChangeCallback={reloadWithQuery}
      seasonOptions={seasonOptions}
    />
    <ClubCommunityServiceSection {rows} club={data.club}/>
  </section>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}

<style>
  p {
    margin-block-end: calc(var(--spacing) * 4);
  }
</style>
