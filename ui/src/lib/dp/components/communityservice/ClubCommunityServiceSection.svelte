<script lang="ts">
  import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";
  import type {ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
  import CommunityServiceDatatable from "$lib/dp/components/communityservice/CommunityServiceDatatable.svelte";

  interface Props {
    rows: ClubCommunityServiceRow[];
    club: ExpandedClub;
  }

  let {rows, club}: Props = $props();
</script>

<div class="wrapper">
  <CommunityServiceDatatable data={rows} {club}/>
</div>

<style>
  .wrapper {
    margin-block: calc(var(--spacing) * 6);
  }
</style>
