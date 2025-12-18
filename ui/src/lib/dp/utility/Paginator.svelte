<script lang="ts">
  import {ChevronLeft, ChevronRight} from "lucide-svelte";
  import type {PageStore} from "$lib/dp/records/PageStore.ts";

  const {
    store,
    showIfSinglePage = false,
  }: {
    store: PageStore
    showIfSinglePage?: boolean
  } = $props();
</script>

{#if showIfSinglePage || $store.totalPages > 1}
  <div
          class="paginator flex flex-col md:flex-row items-center space-y-4 md:space-y-0 md:space-x-4 justify-between"
  >
    <div class="flex rounded-container preset-tonal-surface p-1">
      <button
              class="fill-current p-2 rounded-container"
              type="button"
              onclick={() => store.prev()}
              disabled={$store.page <= 1}
      >
        <ChevronLeft/>
      </button>

      <span class="text-sm! p-2">{$store.page}/{$store.totalPages}</span>

      <button
              class="fill-current p-2 rounded-container"
              type="button"
              onclick={() => store.next()}
              disabled={$store.page >= $store.totalPages}
      >
        <ChevronRight/>
      </button>
    </div>
  </div>
{/if}

<style>
    button:hover:not(:disabled), button:focus:not(:disabled) {
        background-color: var(--color-primary-50-950);
        color: var(--color-primary-950-50);
    }
</style>