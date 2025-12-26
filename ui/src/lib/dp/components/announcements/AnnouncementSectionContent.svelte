<script lang="ts">
  import AnnouncementCard from "$lib/dp/components/announcements/AnnouncementCard.svelte";
  import type {PageStore} from "$lib/dp/records/PageStore.ts";
  import type {ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
  import Paginator from "$lib/dp/utility/Paginator.svelte";

  interface Props {
    store: PageStore<ExpandedAnnouncement>;
  }

  const {store}: Props = $props();
</script>

<div>
  {#if $store.totalItems > 0}
    <ul>
      {#each $store.items as announcement (announcement.id)}
        <li>
          <AnnouncementCard {announcement}/>
        </li>
      {/each}
    </ul>

    <Paginator {store} showIfSinglePage={false}/>
  {:else}
    <p>No announcements available.</p>
  {/if}
</div>

<style>
  ul {
    display: grid;
    grid-template-columns: 1fr;
    gap: calc(var(--spacing) * 5);
    margin-block-end: calc(var(--spacing) * 6);

    @media (min-width: 64rem) {
      grid-template-columns: 1fr 1fr;
    }

    @media (min-width: 80rem) {
      grid-template-columns: 1fr 1fr 1fr;
    }
  }
</style>
