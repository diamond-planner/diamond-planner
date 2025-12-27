<script lang="ts">
interface Option<T extends string> {
  label?: string
  value: T
}

interface Props<T extends string> {
  label?: string
  name: string
  options: Option<T>[] | T[]
  value: T
  class?: string
  listClass?: string
  triggerClass?: string
}

// Usage:
// <TabsRadioGroup name="bats" options={["left","right","switch"]} bind:value={form.bats} label="Bats" />
let {
  label,
  name,
  options,
  value = $bindable(),
  class: klass = "",
  listClass = "md:col-span-2 tabs-list input",
  triggerClass = "tabs-trigger btn",
}: Props<string> = $props()

function toOptions(arr: Option<string>[] | string[]): Option<string>[] {
  return (arr as any[]).map((o) =>
    typeof o === "string" ? { value: o, label: o.charAt(0).toUpperCase() + o.slice(1) } : o
  )
}

const opts: Option<string>[] = toOptions(options)
</script>

{#if label}
  <span class="block">{label}</span>
{/if}
<span class={listClass + (klass ? ` ${klass}` : "")}>
  {#each opts as opt}
    <label class={[triggerClass, value === opt.value && "preset-filled"]}>
      <input
              type="radio"
              class="hidden"
              name={name}
              bind:group={value}
              value={opt.value}
      />
      {opt.label}
    </label>
  {/each}
</span>
