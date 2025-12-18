<script lang="ts">
  import {AlertCircle, Info} from "lucide-svelte";
  import type {AnnouncementsResponse} from "$lib/dp/types/pb-types.ts";

  type Priority = AnnouncementsResponse["priority"]

  interface Props {
    priority: Priority;
  }

  let {priority}: Props = $props();

  const badgeClass = $derived.by(() => {
    switch (priority) {
      case "info":
        return "preset-outlined-secondary-500 dark:border-white";
      case "warning":
        return "preset-outlined-warning-500";
      case "danger":
        return "preset-outlined-error-500";
    }
  });
</script>

<div class="col-span-1 justify-self-end">
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
