<script lang="ts">
  import PlayerNumberGraphic from "$lib/dp/components/player/PlayerNumberGraphic.svelte";
  import type {Player} from "$lib/dp/types/Player.ts";

  interface Props {
    player: Player;
  }

  let {player}: Props = $props();
</script>

<div class="card preset-tonal-surface shadow-xl p-3">
  <div class="flex justify-center">
    {#if player.media?.length > 0}
      <img class="rounded-lg max-h-72" src="{player.media.at(0)?.url}" alt="{player.media.at(0)?.alt}" loading="lazy">
    {:else}
      <img class="rounded-lg max-h-72" src="/team-placeholder-white.jpg" alt="Player fallback" loading="lazy">
    {/if}
  </div>

  <div class="flex items-center gap-4 mt-3">
    <PlayerNumberGraphic content={player.number}/>
    <p class="h3">{player.fullname}</p>
  </div>

</div>
