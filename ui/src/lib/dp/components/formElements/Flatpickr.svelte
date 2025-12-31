<script lang="ts">
  import {onMount} from "svelte";
  import flatpickr from "flatpickr";
  import "flatpickr/dist/flatpickr.css";
  import type {Options} from "flatpickr/dist/types/options";
  import {browser} from "$app/environment";

  if (browser && window.matchMedia("(prefers-color-scheme: dark)").matches) {
    import("flatpickr/dist/themes/dark.css");
  }

  interface Props {
    value: string; // flatpickr accepts timestamps, date objects and ISO strings, but we work with strings exclusively
    options: Options;
  }

  let {value = $bindable(), options = {}}: Props = $props();

  let datepicker: HTMLInputElement;

  onMount(() => {
    const fp = flatpickr(datepicker, {
      ...options,
      defaultDate: value ?? "",
      onChange: (selectedDates) => {
        value = selectedDates[0].toISOString();
      },
    });

    return () => {
      fp.destroy();
    };
  });
</script>

<input class="input" type="text" bind:this={datepicker} {value}/>

<!-- svelte-ignore css_unused_selector -->
<style lang="postcss">
  input {
    color: unset;
  }
</style>
