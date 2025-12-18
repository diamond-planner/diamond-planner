<script lang="ts">
  import type {SingleStatElement} from "$lib/dp/types/SingleStatElement.ts";

  interface props {
    block: SingleStatElement;
    classes?: string;
  }

  const {block, classes = ""}: props = $props();
</script>

<div class="stat {classes}">
  <div class="stat-title">{block.title}</div>
  <div class="stat-value">{block.value}</div>
  <div class="stat-desc text-wrap">{block.desc}</div>
</div>

<style lang="postcss">
  .stat-desc {
    overflow: hidden;
  }
</style>
