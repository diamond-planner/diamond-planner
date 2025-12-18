<script lang="ts">
  import {ClipboardList, Dumbbell, ShieldHalf} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import TeamForm from "$lib/dp/components/forms/TeamForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";

  /**
   * Used for Teams in List Teasers and on Club Page.
   */

  interface props {
    team: ExpandedTeam;
    link: boolean;
  }

  const {team, link = false}: props = $props();

  const model = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("teams").delete(id);
    invalidate("club:single");
  }
</script>

<article class="card preset-tonal-surface shadow-lg" class:card-hover={link}>
  <a href="/account/team/{team.id}">
    <header class="card-header">
      <h3 class="h4 font-semibold">{team.name}</h3>
    </header>

    <section class="p-4">

      <div class="space-y-2">

        <div class="item-row">
          <ShieldHalf/>
          <div>
            <dd>{team?.expand?.club.name}</dd>
            <dt class="text-sm font-light">Club</dt>
          </div>
        </div>

        <div class="item-row">
          <Dumbbell/>
          <div>
            <dd>{team.bsm_league_group}</dd>
            <dt class="text-sm font-light">BSM-Liga (für aktuelle Saison)</dt>
          </div>
        </div>

        <div class="item-row">
          <ClipboardList/>
          <dl>
            <dd class="capitalize">{team.age_group}</dd>
            <dt class="text-sm font-light">Age Group</dt>
          </dl>
        </div>
      </div>
    </section>
  </a>

  <footer class="card-footer">
    <hr class="my-2">
    <div class="flex justify-end gap-2">

      {#if team?.expand?.club?.admins.includes(model.id) || team?.admins.includes(model.id)}
        <TeamForm
                club={team.expand?.club}
                team={team}
                buttonClasses="btn btn-sm btn-icon preset-tonal-tertiary border border-tertiary-500"
                showLabel={false}
        />
      {/if}

      {#if team?.expand?.club?.admins.includes(model.id)}
        <DeleteButton
                id={team.id}
                modelName="Team"
                action={deleteAction}
        />
      {/if}

    </div>
  </footer>
</article>

<style>
    .card {
        display: block;
    }

    .item-row {
        display: flex;
        align-items: center;
        gap: calc(var(--spacing) * 4);
    }
</style>