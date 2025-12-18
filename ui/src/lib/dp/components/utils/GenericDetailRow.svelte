<script lang="ts">
import type { Snippet } from "svelte"

interface Props {
  icon: Snippet
  categoryName: string
  rowValue: string | number
}

let { icon, categoryName, rowValue }: Props = $props()
</script>

<div class="container">
  {@render icon()}
  <dl class="container">
    <dt class="font-light">{categoryName}:</dt>
    <dd class="font-bold">{rowValue}</dd>
  </dl>

</div>

<style lang="postcss">
    .container {
        display: flex;
        align-items: center;
        gap: calc(var(--spacing) * 4)
    }
</style>