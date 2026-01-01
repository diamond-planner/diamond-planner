<script lang="ts">
  import type {CommunityServiceData} from "$lib/dp/types/CommunityServiceData.ts";
  import CommunityServiceItemCard from "$lib/dp/components/communityservice/CommunityServiceItemCard.svelte";
  import TargetVisualizer from "$lib/dp/components/communityservice/TargetVisualizer.svelte";
  import {toHours} from "$lib/dp/utility/toHours.ts";
  import {Plus} from "lucide-svelte";
  import ServiceEntryForm from "$lib/dp/components/forms/ServiceEntryForm.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";

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
      --visualizer-spacing="4"
    />

    <div class="entry-cards">
      {#each item?.entries as entry(entry.id)}
        <CommunityServiceItemCard {entry}/>
      {/each}
    </div>

    <Dialog triggerClasses="btn preset-tonal-secondary border">

      {#snippet triggerContent()}
        <Plus/>
        Create new
      {/snippet}

      {#snippet title()}
        <header>
          <h2 class="h4">Create new Community Service Entry</h2>
        </header>
      {/snippet}

      <ServiceEntryForm serviceEntry={null} club={item.club}/>
    </Dialog>
  </div>
{/each}

<style>
  .club-services {
    margin-block: calc(var(--spacing) * 3);
  }

  .entry-cards {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 3);
    margin-block-end: calc(var(--spacing) * 4);

    @media (min-width: 48rem) {
      grid-template-columns: repeat(2, 1fr);
    }

    @media (min-width: 80rem) {
      grid-template-columns: repeat(3, 1fr);
    }
  }
</style>
