<script lang="ts">
  interface Props {
    selectedSeason: number
    seasonOptions: number[]
    onChangeCallback: () => void
  }

  let {selectedSeason = $bindable(), seasonOptions, onChangeCallback}: Props = $props()
</script>

<div class="container lg:gap-8 preset-tonal-surface rounded-base">
  <label class="label md:grow-0">
    Season
    <select bind:value={selectedSeason} class="select" onchange={onChangeCallback}>
      {#each seasonOptions as option}
        <option value="{option}">{option}</option>
      {/each}
    </select>
  </label>
</div>

<style>
  .container {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    gap: calc(var(--spacing) * 4);
    padding-inline: calc(var(--spacing) * 4);
    padding-block: calc(var(--spacing) * 3);
  }

  label {
    display: flex;
    align-items: center;
    flex-grow: 1;
    justify-content: space-between;
    gap: calc(var(--spacing) * 2);
  }
</style>
