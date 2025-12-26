<script lang="ts">
  import type {CommunityServiceData} from "$lib/dp/types/CommunityServiceData.ts";
  import {toHours} from "$lib/dp/utility/toHours.ts";
  import PersonalCommunityServiceCard from "$lib/dp/components/communityservice/PersonalCommunityServiceCard.svelte";

  interface Props {
    serviceEntryData: CommunityServiceData;
  }

  let {serviceEntryData}: Props = $props();
</script>

{#each serviceEntryData.items as item(item.club.id)}
  <div class="club-services">
    <h2 class="h3">{item.club.name}</h2>

    <dl>
      <dt>Target:</dt>
      <dd>{toHours(item.current_minutes)} / {toHours(item.club.service_requirement)}</dd>
    </dl>

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

  dl {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    margin-block: calc(var(--spacing) * 4);
  }
</style>
