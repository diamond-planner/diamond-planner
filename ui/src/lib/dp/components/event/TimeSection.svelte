<script lang="ts">
  import type {Snippet} from "svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";

  interface Props {
    timeValue: string;
    displayText: string;
    icon: Snippet;
    classes?: string;
  }

  const {timeValue, displayText, icon, classes = ""}: Props = $props();

  const displayedTime = $derived(new Date(timeValue));
</script>

<div class="section-container {classes}">
  {@render icon()}
  <p>
    {displayText}
    {#if timeValue}
      <time datetime="{timeValue}" class="font-bold"
      >{displayedTime?.toLocaleTimeString(
          "de-DE",
          DateTimeUtility.eventTimeFormat,
      )}</time
      >
    {:else}
      <span class="font-medium">---</span>
    {/if}
  </p>
</div>

<style>
    .section-container {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }
</style>