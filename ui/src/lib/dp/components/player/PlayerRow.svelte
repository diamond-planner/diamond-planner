<script lang="ts">
  import PlayerNumberGraphic from "$lib/dp/components/player/PlayerNumberGraphic.svelte";
  import type {Player} from "$lib/dp/types/Player.ts";

  interface Props {
    player: Player;
    visual?: "image" | "number";
  }

  let {player, visual = "number"}: Props = $props();
</script>

{#if player.bsm_id > 0}
  <a class="player-container flex items-center gap-2" href="/players/{player.bsm_id}">
    {@render rowContent()}
  </a>
{:else }
  <div class="player-container flex items-center gap-2">
    {@render rowContent()}
  </div>
{/if}

{#snippet rowContent()}
  {#if visual === "number"}
    <PlayerNumberGraphic content={player.number} classes="preset-filled-primary-500 scale-75"/>
  {:else }

    {#if player.media?.length > 0}
      <img class="rounded-lg max-h-20" src="{player.media.at(0)?.url}" alt="{player.media.at(0)?.alt}" loading="lazy">
    {:else}
      <img class="rounded-lg max-h-20" src="/team-placeholder-white.jpg" alt="Player fallback" loading="lazy">
    {/if}
  {/if}

  <div class="flex flex-col gap-0.5">
    <p>{player.fullname}</p>
    <p class="text-sm font-light">{player.positions.join(", ")}</p>
  </div>
{/snippet}
