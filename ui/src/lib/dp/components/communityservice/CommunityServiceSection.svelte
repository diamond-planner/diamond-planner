<script lang="ts">
  import type {CommunityServiceData} from "$lib/dp/types/CommunityServiceData.ts";

  interface Props {
    serviceEntryData: CommunityServiceData;
  }

  let {serviceEntryData}: Props = $props();
</script>

{JSON.stringify(serviceEntryData)}
