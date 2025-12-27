<script lang="ts">
  import type {CommunityServiceData} from "$lib/dp/types/CommunityServiceData.ts";
  import PersonalCommunityServiceCard from "$lib/dp/components/communityservice/PersonalCommunityServiceCard.svelte";
  import TargetVisualizer from "$lib/dp/components/communityservice/TargetVisualizer.svelte";
  import {toHours} from "$lib/dp/utility/toHours.ts";

  interface Props {
    serviceEntryData: CommunityServiceData;
  }

  let {serviceEntryData}: Props = $props();
</script>

{#each serviceEntryData.items as item(item.club.id)}
  <div class="club-services">
    <h2 class="h3">{item.club.name}</h2>

    <TargetVisualizer
      current={toHours(item.current_minutes)}
      target={toHours(item.club.service_requirement)}
      showTarget={true}
    />

    <div class="entry-cards">
      {#each item?.entries as entry(entry.id)}
        <PersonalCommunityServiceCard {entry}/>
      {/each}
    </div>
  </div>
{/each}

<style>
  .entry-cards {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
  }
</style>
