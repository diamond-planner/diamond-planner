<script lang="ts">
  import type {Match} from "bsm.js";
  import {env} from "$env/dynamic/public";
  import {MatchState} from "$lib/dp/enum/MatchState.ts";
  import {MatchDecorator} from "$lib/dp/service/MatchDecorator.ts";

  interface Props {
    match: Match;
    teamName?: string;
  }

  const {match, teamName}: Props = $props();

  const matchDecorator = $derived(new MatchDecorator(match));
  const matchState = $derived(matchDecorator.getMatchState(teamName ?? env.PUBLIC_TEAM_NAME));
</script>

<div class="text-lg font-semibold">
  {#if matchState === MatchState.won}
    <span class="text-success-700 dark:text-success-500">W</span>
  {:else if matchState === MatchState.lost}
    <span class="text-error-600 dark:text-error-400">L</span>
  {:else if matchState === MatchState.derby}
    <span class="text-primary-600-400 text-2xl">♥</span>
  {:else if matchState === MatchState.final}
    <span>F</span>
  {:else if matchState === MatchState.notYetPlayed}
    <span class="text-md">TBD</span>
  {:else if matchState === MatchState.cancelled}
    <span class="text-md">PPD</span>
  {/if}
</div>
