<script lang="ts">
  import type {Snippet} from "svelte";

  interface Props {
    navLinks?: Snippet;
  }

  let {navLinks}: Props = $props();
</script>

<nav
  class="bottom-nav"
>
  <ul>
    {@render navLinks?.()}
  </ul>
</nav>

<style>
  ul {
    display: flex;
  }

  .bottom-nav {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    padding-bottom: env(safe-area-inset-bottom);
    background-color: var(--nav-item-background);
    z-index: 15;
    border-top: 1px solid light-dark(var(--color-surface-500), var(--color-surface-100));
    border-bottom: 1px solid light-dark(var(--color-surface-500), var(--color-surface-100));

    @media (min-width: 48rem) {
      display: none;
    }
  }
</style>
