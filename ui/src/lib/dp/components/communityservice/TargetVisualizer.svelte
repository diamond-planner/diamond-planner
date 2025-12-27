<script lang="ts">
  interface Props {
    current: number,
    target: number,
    showTarget?: boolean
  }

  let {current, target, showTarget = false}: Props = $props();
  let percentage = $derived(Math.min(Math.max((current / target) * 100, 0), 100));
</script>

<div style="--progress: {percentage};">
  {#if showTarget}
    <dl>
      <dt>Target:</dt>
      <dd class="visualized chip preset-tonal">{current} / {target}</dd>
    </dl>
  {:else }
    <p class="visualized chip preset-tonal">{current}</p>
  {/if}
</div>

<style>
  dl {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    margin-block: calc(var(--spacing) * 4);
  }

  .visualized {
    font-size: var(--text-lg);
    font-weight: var(--font-weight-bold);

    /* Theoretical Color Logic:
       Hue: 0 (Red) to 120 (Green).
    */
    --hue: calc(var(--progress) * 1.2);
    --visualizer-background: hsl(var(--hue), 70%, 45%);
    background-color: var(--visualizer-background);
    color: contrast-color(var(--visualizer-background));
  }
</style>
