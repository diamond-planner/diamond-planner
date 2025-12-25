<script lang="ts">
  // @ts-expect-error
  import {Tabs} from "bits-ui";
  import PlayerRow from "$lib/dp/components/player/PlayerRow.svelte";
  import type {Player} from "$lib/dp/types/Player.ts";

  interface Props {
    players: Player[];
  }

  let {players}: Props = $props();
  let visual: "number" | "image" = $state("number");
</script>

<label class="label flex flex-col gap-1 my-4 lg:max-w-[50%]">
  Visualization
  <Tabs.Root bind:value={visual}>
    <Tabs.List class="tabs-list input">
      <Tabs.Trigger class="tabs-trigger btn flex-grow" value="number">Number</Tabs.Trigger>
      <Tabs.Trigger class="tabs-trigger btn flex-grow" value="image">Image</Tabs.Trigger>
    </Tabs.List>
  </Tabs.Root>
</label>

<div class="card p-3 preset-tonal shadow-xl">
  {#each players as player (player.uid)}
    <PlayerRow {player} {visual}/>

    <hr class="my-3">
  {/each}
</div>
