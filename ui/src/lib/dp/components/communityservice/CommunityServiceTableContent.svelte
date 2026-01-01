<script lang="ts">
  import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";
  import {CircleCheck, CircleMinus} from "lucide-svelte";
  import type {TableHandler} from "@vincjo/datatables";
  import {toHours} from "$lib/dp/utility/toHours.ts";
  import TargetVisualizer from "$lib/dp/components/communityservice/TargetVisualizer.svelte";
  import {page} from "$app/state";

  interface Props {
    handler: TableHandler<ClubCommunityServiceRow>;
  }

  let {handler}: Props = $props();
  const rows = $derived(handler.rows);

  const season = $derived(page.url.searchParams.get("season"));
</script>

{#each rows as row(row.id)}
  <tr>
    <td>
      {row.email}
    </td>

    <td>
      <a class="anchor" href="/account/{row.id}/communityservice?season={season}"
         title="Go to community service page for user {row.name}">
        {row.name}
      </a>
    </td>

    <td>
      <TargetVisualizer
        current={toHours(row.total_minutes)}
        target={toHours(row.target)}
        showTarget={true}
        --visualizer-spacing="1"
      />
    </td>

    <td>
      <div class="checkmark-container">
        {#if row.target_met}
          <div>
            <CircleCheck color="var(--color-success-500)"/>
          </div>
          <div>Met</div>
        {:else}
          <div>
            <CircleMinus color="var(--color-error-500)"/>
          </div>
          <div>Not Met</div>
        {/if}
      </div>
    </td>
  </tr>
{/each}

<style>
  .checkmark-container {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
  }
</style>
