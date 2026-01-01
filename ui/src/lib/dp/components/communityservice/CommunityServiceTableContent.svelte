<script lang="ts">
  import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";

  interface Props {
    rows: ClubCommunityServiceRow[];
  }

  let {rows}: Props = $props();
</script>

{#each rows as row(row.id)}
  <tr>
    <td>
      {row.email}
    </td>

    <td>
      {row.name}
    </td>

    <td>
      {row.total_minutes}
    </td>

    <td>
      {row.target_met}
    </td>
  </tr>
{/each}

<style>
  .checkmark-container {
    display: flex;
    align-items: center;
    gap: var(--spacing);
  }
</style>
