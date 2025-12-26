<script lang="ts">
  import type {PageProps} from "./$types";
  import CommunityServiceSection from "$lib/dp/components/communityservice/CommunityServiceSection.svelte";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";
  import {range} from "$lib/dp/utility/range.ts";
  import {goto} from "$app/navigation";
  import SeasonSelector from "$lib/dp/components/utils/SeasonSelector.svelte";

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

<h1 class="h1">Community Service</h1>

{#await data.set}
  <ProgressRing/>
{:then set}
  <p class="h5">for {set.user.first_name} {set.user.last_name} ({set.season})</p>
  <section class="service-section">
    <SeasonSelector bind:selectedSeason={selectedSeason} onChangeCallback={reloadWithQuery} seasonOptions={seasonOptions}/>
    <CommunityServiceSection serviceEntryData={set}/>
  </section>


{:catch error}
  <p>error loading: {error.message}</p>
{/await}

<style>
  .service-section {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 3);
  }

  .h5 {
    margin-block-end: calc(var(--spacing) * 4);
  }
</style>
