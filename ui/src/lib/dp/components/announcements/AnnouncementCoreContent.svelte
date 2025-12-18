<script lang="ts">
  import {User} from "lucide-svelte";
  import type {ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
  import PriorityBadge from "./PriorityBadge.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
    textClasses?: string;
  }

  let {announcement, textClasses = ""}: Props = $props();
</script>

<header class="flex justify-between hyphens-auto">
  <h3 class="h4">{announcement.title}</h3>

  <PriorityBadge priority={announcement.priority}/>
</header>

<div class="mt-4">
  <p class="my-4 flex gap-3 {textClasses}">
    <User/>
    <span>
    {announcement.expand?.author.first_name}
      {announcement.expand?.author.last_name}
  </span>
  </p>

  <p class={textClasses}>
    {@html announcement.bodytext}
  </p>
</div>
