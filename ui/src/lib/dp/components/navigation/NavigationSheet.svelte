<script lang="ts">
  import SidebarNavigation from "$lib/dp/components/meta/SidebarNavigation.svelte";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import type {ExpandedClub, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    clubs: ExpandedClub[];
    teams: ExpandedTeam[];
  }

  let {clubs, teams}: Props = $props();

  let open = $state(false);
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger aria-label="open navigation" class="nav-trigger btn btn-sm">
    <svg viewBox="0 0 100 80">
      <rect height="20" width="100"/>
      <rect height="20" width="100" y="30"/>
      <rect height="20" width="100" y="60"/>
    </svg>
  </Sheet.Trigger>

  <Sheet.Content class="navigation-sheet-content" side="left">
    <Sheet.Header></Sheet.Header>

    <a aria-label="to home page" class="anchor" href="/" onclick={() => open = false}>
      <img alt="Diamond Planner Logo" src="/icon_dp.svg">

      <h2>Home Page</h2>
    </a>

    <hr/>

    <SidebarNavigation bind:sheetOpen={open} clubs={clubs} teams={teams}/>

  </Sheet.Content>
</Sheet.Root>

<style>
    :global {
        .nav-trigger {
            margin-inline-end: 2em;
            @media (min-width: 48rem) {
                display: none;
            }
        }
    }

    a {
        display: flex;
        justify-content: space-around;
        padding: 0.5em;

        img {
            max-width: calc(var(--spacing) * 14);
        }

        h2 {
            padding: 1em;
            -webkit-font-smoothing: antialiased;
            -moz-osx-font-smoothing: grayscale;
        }
    }

    hr {
        margin-block-end: 0.5em;
    }

    svg {
        --svg-size: calc(var(--spacing) * 4);
        width: var(--svg-size);
        height: var(--svg-size);
        fill: var(--color-surface-950-50);
    }
</style>
