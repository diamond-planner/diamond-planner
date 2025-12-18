<script lang="ts">
  import ClubLocationsSection from "$lib/dp/components/club/ClubLocationsSection.svelte";
  import type {ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
  import type {PageProps} from "../../../../../../../../.svelte-kit/types/src/routes";

  let {data}: PageProps = $props();

  let club: ExpandedClub = $derived(data.club);
</script>

<svelte:head>
  <title>Club Locations</title>
  <meta
          content="All configured locations for {club.name}, including baseball fields and other venues."
          name="description"
  />
</svelte:head>

<h1 class="h1">{club.name} - Locations</h1>

<section class="mt-8!">
  <ClubLocationsSection club={club} locations={data?.locations ?? []}/>
</section>