<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {ExpandedParticipation} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import {sendParticipationData} from "$lib/dp/utility/sendParticipationData.ts";

  interface Props {
    participation: ExpandedParticipation;
  }

  let {participation}: Props = $props();

  let form = $state(participation);

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    await sendParticipationData(form).then(() => closeModal());
    await invalidate("event:list");
  }

  const createdDate =
      participation.created !== "" ? new Date(participation.created).toLocaleString() : "---";
  const updatedDate =
      participation.updated !== "" ? new Date(participation.updated).toLocaleString() : "---";
</script>

<form class="mt-4 space-y-3" onsubmit={submitForm}>
  <div class="grid grid-cols-2 gap-2 md:gap-3 lg:gap-4">

    <dl class="space-y-1">
      <dt class="text-sm font-light">Created</dt>
      <dd>{createdDate}</dd>
    </dl>

    <dl class="space-y-1">
      <dt class="text-sm font-light">Last Updated</dt>
      <dd>{updatedDate}</dd>
    </dl>

    <input
            autocomplete="off"
            bind:value={form.id}
            class="input"
            name="id"
            readonly
            type="hidden"
    />

    <label class="label col-span-2">
      <span class="block">Comment</span>
      <input
              autocomplete="off"
              bind:value={form.comment}
              class="input "
              name="id"
              placeholder="background info about your attendance"
              type="text"
      />
    </label>

    <div class="label col-span-2">
      State

      <span class="btn-group">
        <button
                class={["btn hover:preset-tonal-success flex-grow", form.state === "in" && "preset-filled-success-500 text-black"]}
                onclick={() => form.state = "in"}
                type="button"
        >
          In
        </button>
        <button
                class={["btn hover:preset-tonal-warning flex-grow", form.state === "maybe" && "preset-filled-warning-500 text-black"]}
                onclick={() => form.state = "maybe"}
                type="button"
        >
          Maybe
        </button>
        <button
                class={["btn hover:preset-tonal-error flex-grow", form.state === "out" && "preset-filled-error-500 text-white"]}
                onclick={() => form.state = "out"}
                type="button"
        >
          Out
        </button>
      </span>
    </div>

    <div class="flex justify-center col-span-2">
      <button class="mt-2 btn preset-filled-primary-500" type="submit">
        Submit
      </button>
    </div>
  </div>
</form>

<style>
    .btn-group {
        display: flex;
        justify-items: stretch;
        border: 1px solid;
    }
</style>