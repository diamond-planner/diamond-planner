<script lang="ts">
  import {DataHandler} from "@vincjo/datatables";
  import Pagination from "$lib/dp/components/datatable/Pagination.svelte";
  import RowCount from "$lib/dp/components/datatable/RowCount.svelte";
  import RowsPerPage from "$lib/dp/components/datatable/RowsPerPage.svelte";
  import Search from "$lib/dp/components/datatable/Search.svelte";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import TeamMembersTableContent from "./TeamMembersTableContent.svelte";

  interface Props {
    data: CustomAuthModel[];
    team: ExpandedTeam;
    rowsPerPage?: number;
    showAdminSection?: boolean;
  }

  let {data, team, rowsPerPage = 25, showAdminSection = false}: Props = $props();

  const handler = $derived(
    new DataHandler<CustomAuthModel>(data, {
      rowsPerPage: rowsPerPage,
    })
  );
</script>

<div class="table-wrap">
  <!-- Header -->
  <header class="table-header">
    <Search {handler}/>

    <RowsPerPage {handler}/>
  </header>

  <!-- Table -->
  <table class="table table-compact">
    <TeamMembersTableContent {handler} {showAdminSection} {team}/>

    <tfoot></tfoot>
  </table>

  <!-- Footer -->
  <footer class="table-footer">
    <RowCount {handler}/>
    <Pagination {handler}/>
  </footer>
</div>
