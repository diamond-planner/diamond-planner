<script lang="ts">
import type { DataHandler } from "@vincjo/datatables"

interface Props {
  handler: DataHandler
}

let { handler }: Props = $props()
const pageNumber = handler.getPageNumber()
const pageCount = handler.getPageCount()
const pages = handler.getPages({ ellipsis: true })
</script>

<!-- Desktop buttons -->
<section class=" preset-outlined-primary-500 h-10 hidden lg:flex rounded-container items-center">
    <button
            type="button"
            class="hover:preset-tonal-primary p-2"
            class:disabled={$pageNumber === 1}
            onclick={() => handler.setPage('previous')}
    >
        ←
    </button>
    {#each $pages as page}
        <button
                type="button"
                class="hover:preset-tonal-primary p-2 m-1"
                class:active={$pageNumber === page}
                class:ellipse={page === null}
                onclick={() => handler.setPage(page)}
        >
            {page ?? '...'}
        </button>
    {/each}
    <button
            type="button"
            class="hover:preset-tonal-primary p-2"
            class:disabled={$pageNumber === $pageCount}
            onclick={() => handler.setPage('next')}
    >
        →
    </button>
</section>

<!-- Mobile buttons -->
<section class="lg:hidden">
    <button
            type="button"
            class="btn preset-outlined-primary-500 mr-2 mb-2"
            class:disabled={$pageNumber === 1}
            onclick={() => handler.setPage('previous')}
    >
        ←
    </button>
    <button
            type="button"
            class="btn preset-outlined-primary-500 mb-2"
            class:disabled={$pageNumber === $pageCount}
            onclick={() => handler.setPage('next')}
    >
        →
    </button>
</section>