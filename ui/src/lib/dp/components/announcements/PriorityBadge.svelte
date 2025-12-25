<script lang="ts">
  import type {AnnouncementsResponse} from "$lib/dp/types/pb-types.ts";

  type Priority = AnnouncementsResponse["priority"]

  interface Props {
    priority: Priority;
  }

  let {priority}: Props = $props();

  const badgeClass = $derived.by(() => {
    switch (priority) {
      case "info":
        return "preset-outlined-secondary-500 badge-info";
      case "warning":
        return "preset-outlined-warning-500";
      case "danger":
        return "preset-outlined-error-500";
    }
  });
</script>

<div class="priority-badge-container">
  <span class="badge {badgeClass}">
    {#if priority === "info"}
      Info
    {:else if priority === "warning"}
      Important
    {:else if priority === "danger"}
      Critical
    {/if}
  </span>
</div>

<style>
  .priority-badge-container {
    justify-self: flex-end;
    grid-column: span 1 / span 1;
  }

  :global {
    .priority-badge-container {
      .badge.badge-info {
        border-color: white;
      }
    }
  }
</style>
