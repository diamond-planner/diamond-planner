<script lang="ts">
  import Cap from "$lib/dp/components/icons/Cap.svelte";
  import Pants from "$lib/dp/components/icons/Pants.svelte";
  import Shirt from "$lib/dp/components/icons/Shirt.svelte";
  import type {UniformsetsResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    attire: UniformsetsResponse;
  }

  let {attire}: Props = $props();
</script>

<h2 class="h2">Uniform</h2>
<div class="p-4 flex gap-2 md:gap-3 justify-evenly">
  <Cap classes="w-16" identifier={attire.cap}/>
  <Shirt classes="w-12 h-12" fillColor={attire.jersey}/>
  <Pants classes="w-12 h-12" fillColor={attire.pants}/>
</div>