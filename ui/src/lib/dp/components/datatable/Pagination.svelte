<script lang="ts">
  import type {TableHandler} from "@vincjo/datatables";

  interface Props {
    handler: TableHandler;
  }

  let {handler}: Props = $props();
  const pageNumber = $derived(handler.currentPage);
  const pageCount = $derived(handler.pageCount);
  const pages = $derived(handler.pagesWithEllipsis);
</script>

<!-- Desktop buttons -->
<section class="desktop-section preset-outlined-primary-500 rounded-container">
  <button
    type="button"
    class="hover:preset-tonal-primary"
    class:disabled={pageNumber === 1}
    onclick={() => handler.setPage('previous')}
  >
    ←
  </button>
  {#each pages as page}
    <button
      type="button"
      class="hover:preset-tonal-primary"
      class:active={pageNumber === page}
      class:ellipse={page === null}
      onclick={() => handler.setPage(page)}
    >
      {page ?? '...'}
    </button>
  {/each}
  <button
    type="button"
    class="hover:preset-tonal-primary"
    class:disabled={pageNumber === pageCount}
    onclick={() => handler.setPage('next')}
  >
    →
  </button>
</section>

<!-- Mobile buttons -->
<section class="mobile-section">
  <button
    type="button"
    class="btn preset-outlined-primary-500 mr-2"
    class:disabled={pageNumber === 1}
    onclick={() => handler.setPage('previous')}
  >
    ←
  </button>
  <button
    type="button"
    class="btn preset-outlined-primary-500"
    class:disabled={pageNumber === pageCount}
    onclick={() => handler.setPage('next')}
  >
    →
  </button>
</section>

<style>
  .desktop-section {
    display: none;
    align-items: center;
    height: calc(var(--spacing) * 10);

    @media (min-width: 64rem) {
      display: flex;
    }

    button {
      padding: calc(var(--spacing) * 2);
    }
  }

  @media (min-width: 64rem) {
    .mobile-section {
      display: none;

      button {
        margin-block-end: calc(var(--spacing) * 2);
      }
    }
  }
</style>
