<script lang="ts">
import type { Snippet } from "svelte"

interface Props {
  lead?: Snippet
  children?: Snippet
  trail?: Snippet
}

let { lead, children, trail }: Props = $props()
</script>

<header data-testid="app-bar">
  <section data-testid="app-bar-toolbar">

    <div class="lead">
      {@render lead?.()}
    </div>

    <div class="center">
      {@render children?.()}
    </div>

    <div class="trail">
      {@render trail?.()}
    </div>

  </section>
</header>

<style>
    header {
        width: 100%;
        display: flex;
        flex-direction: column;
        background-color: var(--nav-item-background);
        padding: calc(var(--spacing) * 3);
    }

    section {
        display: flex;
        justify-content: space-between;
        gap: calc(var(--spacing) * 4);
    }

    .lead, .trail {
        display: flex;
    }

    .center {
        flex-grow: 1;
        place-self: center;
    }
</style>