<script lang="ts">
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  const authRecord = authSettings.record as CustomAuthModel;

  const form = $state({
    email: authRecord.email,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    const sent = await client.collection("users").requestEmailChange(form.email);

    if (sent) {
      closeModal();
      toastController.trigger({
        message: "A verification email has been sent to your new address.",
        background: "preset-filled-success-500",
      });
    } else {
      toastController.trigger({
        message: "An error occurred while sending your email change request.",
        background: "preset-filled-error-500",
      });
    }
    client.authStore.clear();
  }
</script>

<form class="mt-4 space-y-3" onsubmit={submitForm}>

  <label class="label">
    E-Mail
    <input
            autocomplete="email"
            bind:value={form.email}
            class="input"
            name="email"
            required
            type="email"
    />
    <span class="mt-3 font-light text-sm">
                You will be logged out and instructed to verify your new email address.
            </span>
  </label>

  <div class="flex justify-end gap-3 mt-3">
    <button class="mt-2 btn preset-tonal-primary border border-primary-500" disabled={form.email === authRecord.email}
            type="submit">
      Confirm
    </button>
  </div>
</form>
