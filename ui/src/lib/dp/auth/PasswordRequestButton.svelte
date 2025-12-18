<script lang="ts">
  import type {Snippet} from "svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.js";

  interface Props {
    email: string;
    classes?: string;
    disabled?: boolean;
    additionalAction?: () => void;
    children?: Snippet;
  }

  const {
    email,
    disabled = false,
    additionalAction = () => {
    },
    classes = "",
    children,
  }: Props = $props();

  async function triggerPasswordChange() {
    additionalAction();

    try {
      const success = await client.collection("users").requestPasswordReset(email);

      if (success) {
        toastController.trigger({
          message: `If we have your email (${email}) on file, you will receive a reset password request in your inbox.`,
          background: "preset-filled-success-500",
        });
      }

      if (client.authStore.isValid) {
        client.authStore.clear();
      }
    } catch {
      toastController.trigger({
        message: `Error sending reset password request.`,
        background: "preset-filled-error-500",
      });
    }
    closeModal();
  }
</script>

<Dialog
        closeButtonClasses="sr-only"
        disabled={disabled}
        triggerClasses={classes}
>
  {#snippet triggerContent()}
    {@render children?.()}
  {/snippet}

  {#snippet title()}
    <span>Please Confirm</span>
  {/snippet}

  {#snippet description()}
    <span>
      Send change password request for <code class="code">{email ? email : "---"}</code>?
      You will be logged out and an email with instructions will be sent to your address.
    </span>
  {/snippet}

  <div class="flex justify-end gap-2 mt-1">
    <button class="btn preset-tonal-surface border border-surface-500" onclick={closeModal} type="button">Cancel
    </button>
    <button class="btn preset-filled" onclick={triggerPasswordChange} type="button">Confirm</button>
  </div>
</Dialog>
