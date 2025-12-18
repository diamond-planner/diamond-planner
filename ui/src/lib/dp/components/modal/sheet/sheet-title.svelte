<script lang="ts">
import { Dialog as SheetPrimitive } from "bits-ui"

let { ref = $bindable(null), class: className, ...restProps }: SheetPrimitive.TitleProps = $props()
</script>

<SheetPrimitive.Title
  bind:ref
  class={["text-foreground text-lg font-semibold", className]}
  {...restProps}
/>
