<script lang="ts">
  import {flip} from "svelte/animate";
  import {fly} from "svelte/transition";
  import {toastController} from "$lib/dp/service/ToastController.svelte.js";
</script>

<section class="gui-toast-group">
  {#each toastController.toastQueue as toast (toast.id)}
    <output
            transition:fly={{ y: 200, duration: 200 }}
            animate:flip={{duration: 400}}
            class="gui-toast {toast.background} rounded-base p-4"
    >
      {toast.message}
    </output>
  {/each}
</section>

<style>
    .gui-toast-group {
        position: fixed;
        z-index: 999999;
        inset-block-end: 0;
        inset-inline: 0;
        padding-block-end: 5vh;
        display: grid;
        justify-items: center;
        justify-content: center;
        gap: 1.5vh;
        pointer-events: none;
    }
</style>