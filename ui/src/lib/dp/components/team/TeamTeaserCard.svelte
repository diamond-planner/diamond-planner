<script lang="ts">
  import {File, IdCard, ShieldHalf, Users} from "lucide-svelte";
  import type {ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";

  /**
   * Used on Team Detail page.
   */

  interface props {
    team: ExpandedTeam;
    link: boolean;
  }

  let {team, link = false}: props = $props();
</script>

<article class="card block preset-tonal-surface p-3" class:card-hover={link}>

  <div class="flex items-center gap-3 self-end">
    <File/>
    <div>
      <p>{team.name}</p>
      <p class="text-sm font-light">Name</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <IdCard/>
    <div>
      <p>{team.bsm_league_group}</p>
      <p class="text-sm font-light">BSM-Liga (für aktuelle Saison)</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <ShieldHalf/>
    <div>
      <p>{team?.expand?.club.name}</p>
      <p class="text-sm font-light">Club</p>
    </div>
  </div>

  <hr class="my-2">

  <div class="flex items-center gap-3">
    <Users/>
    <dl>
      <dd class="capitalize">{team.age_group}</dd>
      <dt class="text-sm font-light">Age Group</dt>
    </dl>
  </div>

</article>
