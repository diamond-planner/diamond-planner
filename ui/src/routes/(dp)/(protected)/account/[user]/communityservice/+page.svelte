<script lang="ts">
  import type {PageProps} from "./$types";
  import CommunityServiceSection from "$lib/dp/components/communityservice/CommunityServiceSection.svelte";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Community Service</h1>

{#await data.set}
  <ProgressRing/>
{:then set}
  <p class="h5">for {set.user.first_name} {set.user.last_name}</p>
  <section>
    <CommunityServiceSection serviceEntryData={set}/>
  </section>


{:catch error}
  <p>error loading: {error.message}</p>
{/await}

