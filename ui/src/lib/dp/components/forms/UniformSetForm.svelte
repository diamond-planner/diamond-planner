<script lang="ts">
  import {invalidate} from "$app/navigation";
  import {save} from "$lib/dp/records/RecordOperations.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {UniformsetsCreate, UniformsetsResponse} from "$lib/dp/types/pb-types.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  interface Props {
    uniformSet?: UniformsetsResponse | null;
    clubID: string;
  }

  const {uniformSet, clubID}: Props = $props();

  let form: UniformsetsCreate = $state(
      uniformSet ?? {
        id: "",
        name: "",
        cap: "",
        jersey: "",
        pants: "",
        club: "",
      }
  );

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();
    form.club = clubID;

    let result: UniformsetsResponse | null = null;

    try {
      result = await save<UniformsetsResponse>("uniformsets", form);
    } catch {
      toastController.triggerGenericFormErrorMessage("Uniform Set");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Uniform Set");
      await invalidate("club:single");
      closeModal();
    }
  }
</script>

<form class="mt-4 space-y-3" onsubmit={submitForm}>
  <input
          autocomplete="off"
          bind:value={form.id}
          class="input"
          name="id"
          readonly
          type="hidden"
  />

  <label class="label">
    <span>Name</span>
    <input
            bind:value={form.name}
            class="input"
            name="name"
            required
            type="text"
    />
  </label>

  <label class="label">
    <span>Cap Color</span>
    <span class="grid grid-cols-[auto_1fr] gap-2">
                <input bind:value={form.cap} class="input" type="color"/>
                <input bind:value={form.cap} class="input" required tabindex="-1" type="text"/>
            </span>
  </label>

  <div class="text-sm mt-4!">
    You can use the special identifiers <code class="code">B</code> and <code class="code">Bird</code> to get
    the Skylarks Home and Road Cap instead of a
    generic colored one.
  </div>

  <label class="label">
    <span>Jersey Color</span>
    <span class="grid grid-cols-[auto_1fr] gap-2">
                <input bind:value={form.jersey} class="input" type="color"/>
                <input bind:value={form.jersey} class="input" required tabindex="-1" type="text"/>
            </span>
  </label>

  <label class="label">
    <span>Pants Color</span>
    <span class="grid grid-cols-[auto_1fr] gap-2">
                <input bind:value={form.pants} class="input" type="color"/>
                <input bind:value={form.pants} class="input" required tabindex="-1" type="text"/>
            </span>
  </label>

  <div class="text-sm mt-4!">
    Both hexadecimal color values (6 characters with leading '#') and
    <a class="anchor" href="https://developer.mozilla.org/en-US/docs/Web/CSS/named-color" target="_blank">
      named CSS colors
    </a>
    can be used.
  </div>

  <div class="flex justify-end gap-3 mt-3">
    <button class="mt-2 btn preset-tonal-primary border border-primary-500" type="submit">Submit</button>
  </div>
</form>
