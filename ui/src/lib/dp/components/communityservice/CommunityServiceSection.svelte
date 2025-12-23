<script lang="ts">
  import type {ExpandedServiceEntry} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    serviceEntries: ExpandedServiceEntry[]
  }

  let {serviceEntries}: Props = $props();
</script>

{JSON.stringify(serviceEntries)}
