<script lang="ts">
  import {Plus} from "lucide-svelte";
  import TeamMembersForm from "$lib/dp/components/forms/TeamMembersForm.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import type {ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    club: ClubsResponse;
    team: ExpandedTeam;
  }

  let {club, team}: Props = $props();
</script>

<Dialog triggerClasses="btn preset-filled-primary-500">

  {#snippet triggerContent()}
    <Plus/>
    <span>Add Members</span>
  {/snippet}

  {#snippet title()}
    Add Members to team "{team.name}"
  {/snippet}

  {#snippet description()}
    <div class="text-sm">
      Select all {club.name} members that should be added to the team.
    </div>
  {/snippet}

  <TeamMembersForm {club} {team}/>
</Dialog>