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
    <ul class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-5 mb-6">
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