<script lang="ts">
  // for some reason this import is falsely detected as unused
  //@ts-expect-error
  import {Popover} from "bits-ui";
  import {MessageCircleMore} from "lucide-svelte";
  import ParticipationForm from "$lib/dp/components/forms/ParticipationForm.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedParticipation} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    participation: ExpandedParticipation;
    isAdmin: boolean;
    classes?: string;
  }

  const {participation, isAdmin, classes = ""}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const canEdit = $derived(participation.user === authRecord.id || isAdmin);
  let isOpen = $state(false);

  const createdDate =
      participation.created !== "" ? new Date(participation.created).toLocaleString() : "---";
  const updatedDate =
      participation.updated !== "" ? new Date(participation.updated).toLocaleString() : "---";
</script>

<Popover.Root bind:open={isOpen}>
  <Popover.Trigger>
    {#snippet child({props})}
      <button {...props} class="{classes} relative inline-block">
        {#if participation.comment}
          <span class="badge-icon preset-filled-surface-400-600">
            <MessageCircleMore size="13"/>
          </span>
        {/if}
        {participation?.expand?.user?.first_name}
      </button>
    {/snippet}
  </Popover.Trigger>

  <Popover.Content class="z-15">
    <div class="popover-content card shadow-2xl">

      <div class="item-container">
        <p class="font-light text-xs">Created:</p>
        <p>{createdDate}</p>
      </div>

      <div class="item-container">
        <p class="font-light text-xs">Last Updated:</p>
        <p>{updatedDate}</p>
      </div>

      <div class="item-container">
        <p class="font-light text-xs">Comment/Reason:</p>
        <p>{participation.comment}</p>
      </div>

      {#if canEdit}
        <Dialog triggerClasses="btn btn-sm preset-tonal-primary border border-primary-500 mt-2">

          {#snippet triggerContent()}
            Edit
          {/snippet}

          {#snippet title()}
            <header>
              <h2 class="h3">
                Edit participation data
                for {participation?.expand?.user?.first_name} {participation?.expand?.user?.last_name}
              </h2>
            </header>
          {/snippet}

          <ParticipationForm {participation}/>
        </Dialog>
      {/if}
    </div>
    <Popover.Close/>
    <Popover.Arrow/>
  </Popover.Content>

</Popover.Root>

<style lang="postcss">
    .popover-content {
        border: 1px solid;
        color: light-dark(black, white);
        background-color: light-dark(var(--color-surface-50), var(--color-surface-800));
        padding: calc(var(--spacing) * 4);
        font-size: var(--text-sm);
        line-height: var(--tw-leading, var(--text-sm--line-height));
        max-width: calc(var(--spacing) * 80);
    }

    .item-container {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        justify-content: space-between;
        margin-block: 0.3em;
    }

    .badge-icon {
        color: white;
        top: -15px;
        right: -11px;
        z-index: 10;
        position: absolute;
    }
</style>