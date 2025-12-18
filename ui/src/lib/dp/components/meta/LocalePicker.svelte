<script lang="ts">
  import {goto} from "$app/navigation";
  import {page} from "$app/state";
  import {appLocale} from "$lib/dp/locale.svelte.ts";

  async function reload() {
    await goto(page.url, {
      noScroll: true,
      keepFocus: true,
      invalidateAll: true,
    });
    const htmlElement = document.querySelector("html");

    if (htmlElement) {
      htmlElement.lang = appLocale.current;
    }
  }
</script>

<select bind:value={appLocale.current} class="select" onchange={reload}>
  <option value="en">🇺🇸</option>
  <option value="de">🇩🇪</option>
  <option value="fr">🇫🇷</option>
  <option value="es">🇪🇸</option>
  <option value="pl">🇵🇱</option>
  <option value="ru">🇷🇺</option>
</select>


<style>
  select {
    font-size: 1.5rem;
  }

  option {
    font-size: 2rem;
  }
</style>
