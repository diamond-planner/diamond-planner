<script lang="ts">
  import {TableHandler} from "@vincjo/datatables";
  import type {ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
  import GenericDatatable from "$lib/dp/components/datatable/GenericDatatable.svelte";
  import type {FilterSortOptions} from "$lib/dp/types/FilterSortOptions.ts";
  import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";
  import CommunityServiceTableContent from "$lib/dp/components/communityservice/CommunityServiceTableContent.svelte";

  interface Props {
    data: ClubCommunityServiceRow[];
    club: ExpandedClub;
    rowsPerPage?: number;
  }

  const filterSortOptions: FilterSortOptions = {
    email: {displayName: "E-Mail", value: ""},
    name: {displayName: "Name", value: ""},
    total_minutes: {displayName: "Status", value: ""},
    target_met: {displayName: "Target Met", value: ""},
  };

  let {
    data,
    club,
    rowsPerPage = 25,
  }: Props = $props();

  const handler = $derived(
    new TableHandler<ClubCommunityServiceRow>(data, {
      rowsPerPage: rowsPerPage,
    })
  );
</script>

<GenericDatatable
  {handler}
  {filterSortOptions}
  showInHeader="all"
  caption="Community Service Data for {club.name}"
>
  {#snippet contentRows()}
    <CommunityServiceTableContent {handler}/>
  {/snippet}
</GenericDatatable>
