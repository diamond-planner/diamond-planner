<script lang="ts">
  import type {Match} from "bsm.js";
  import GameResultIndicator from "$lib/dp/components/event/match/GameResultIndicator.svelte";
  import MatchTeaserCardRow from "$lib/dp/components/event/match/MatchTeaserCardRow.svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {GameWinner} from "$lib/dp/enum/GameWinner.ts";
  import {LogoUtility} from "$lib/dp/utility/LogoUtility.ts";
  import {MatchDecorator} from "$lib/dp/service/MatchDecorator.ts";

  interface Props {
    match: Match;
    teamName?: string;
  }

  const {match, teamName}: Props = $props();

  const awayLogo = $derived(LogoUtility.getLogoPathForTeamName(match?.away_league_entry?.team));
  const homeLogo = $derived(LogoUtility.getLogoPathForTeamName(match?.home_league_entry?.team));

  const matchDecorator = $derived(new MatchDecorator(match));
  const winner = $derived(matchDecorator.getWinnerForMatch());
  const matchDate = $derived(DateTimeUtility.parseDateFromBSMString(match.time));
</script>

<a
  class=" block card card-hover text-sm {matchDecorator.isPlayoffGame() ===
    true
        ? 'preset-tonal-primary border border-primary-500'
        : 'preset-tonal-surface border border-surface-500'}"
  href="/gamecenter/game-detail/{match.id}"
>
  <header class="card-header">
    <div class="flex justify-between items-center">
      <div>
        <h3 class="font-medium">{match.league?.name}</h3>
        <p class="font-light">
          {DateTimeUtility.dateTimeFormatMedium.format(matchDate)}
        </p>
      </div>
      <GameResultIndicator {match} {teamName}/>
    </div>
  </header>

  <hr class="mt-2 mx-3"/>

  <section>
    <MatchTeaserCardRow
      logo={awayLogo}
      score={match.away_runs}
      subduedScore={winner === GameWinner.home}
      teamName={match.away_team_name}
    />
    <MatchTeaserCardRow
      logo={homeLogo}
      score={match.home_runs}
      subduedScore={winner === GameWinner.away}
      teamName={match.home_team_name}
    />
  </section>
</a>

<style>
  section {
    padding: calc(var(--spacing) * 4);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 2);

  }
</style>
