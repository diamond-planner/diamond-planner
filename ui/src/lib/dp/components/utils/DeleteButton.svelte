<script lang="ts">
  import {Trash} from "lucide-svelte";
  import type {HTMLButtonAttributes} from "svelte/elements";
  import {invalidate} from "$app/navigation";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {Toast} from "$lib/dp/types/Toast.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  interface Props extends HTMLButtonAttributes {
    id: string;
    modelName: string;
    action: (id: string) => void;
    classes?: string;
    buttonText?: string;
    iconSize?: number | string;
  }

  const {
    id,
    modelName,
    action,
    classes = "btn-sm btn-icon preset-tonal-error border border-error-500",
    buttonText = "",
    iconSize = 24,
    ...restProps
  }: Props = $props();

  function deleteItem() {
    try {
      action(id);

      const toastSettingsDeletions: Toast = {
        message: `${modelName} deleted successfully.`,
        background: "preset-filled-success-500",
      };

      toastController.trigger(toastSettingsDeletions);
      invalidate("nav:load");
    } catch {
      const toastSettingsDeletions: Toast = {
        message: `Error deleting ${modelName}.`,
        background: "preset-filled-error-500",
      };

      toastController.trigger(toastSettingsDeletions);
    }
    closeModal();
  }
</script>

<Dialog closeButtonClasses="sr-only" triggerClasses="btn {classes}" triggerProps={restProps}>
  {#snippet triggerContent()}
    <Trash size={iconSize}/>
    {#if buttonText}
      <span>{buttonText}</span>
    {/if}
  {/snippet}

  {#snippet title()}
    <span>Please Confirm</span>
  {/snippet}

  {#snippet description()}
    <span>Are you sure you wish to delete this {modelName}? In most cases, data cannot be recovered after deletion. Please be certain.</span>
  {/snippet}

  <div class="flex justify-end gap-2 mt-1">
    <button class="btn preset-tonal-surface border border-surface-500" data-test-role="modal-cancel"
            onclick={closeModal}
            type="button">Cancel
    </button>
    <button class="btn preset-filled" data-test-role="modal-confirm" onclick={deleteItem} type="button">Confirm</button>
  </div>
</Dialog>
