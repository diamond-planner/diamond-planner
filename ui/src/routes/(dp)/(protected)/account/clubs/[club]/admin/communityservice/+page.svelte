<script lang="ts">
  import type {PageProps} from "./$types";
  import ProgressRing from "$lib/dp/components/utils/ProgressRing.svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">
  Community Service Admin Dashboard
</h1>
<p>for {data.club.name}</p>

{#await data.rows}
  <ProgressRing/>
{:then rows}
  <div>{JSON.stringify(rows)}</div>
{:catch error}
  <p>error loading: {error.message}</p>
{/await}
