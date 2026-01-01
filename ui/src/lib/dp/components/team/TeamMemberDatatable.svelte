<script lang="ts">
  import {TableHandler} from "@vincjo/datatables";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import TeamMembersTableContent from "./TeamMembersTableContent.svelte";
  import GenericDatatable from "$lib/dp/components/datatable/GenericDatatable.svelte";
  import type {FilterSortOptions} from "$lib/dp/types/FilterSortOptions.ts";

  interface Props {
    data: CustomAuthModel[];
    team: ExpandedTeam;
    rowsPerPage?: number;
    showAdminSection?: boolean;
  }

  const filterSortOptions: FilterSortOptions = {
    last_name: {displayName: "Name", value: ""},
    verified: {displayName: "Status", value: ""},
    number: {displayName: "Number", value: ""},
    bsm_id: {displayName: "BSM ID", value: ""},
    position: {displayName: "Position", value: ""},
    last_login: {displayName: "Last Login", value: ""},
  };

  let {
    data,
    team,
    rowsPerPage = 25,
    showAdminSection = false
  }: Props = $props();

  const handler = $derived(
    new TableHandler<CustomAuthModel>(data, {
      rowsPerPage: rowsPerPage,
    })
  );
</script>

<GenericDatatable {handler} {filterSortOptions} showInHeader="sort">
  {#snippet additionalSort()}
    {#if showAdminSection}
      <th>Actions</th>
    {/if}
  {/snippet}

  {#snippet contentRows()}
    <TeamMembersTableContent {handler} {showAdminSection} {team}/>
  {/snippet}
</GenericDatatable>
