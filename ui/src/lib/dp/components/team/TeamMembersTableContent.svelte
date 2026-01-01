<script lang="ts">
  import type {TableHandler} from "@vincjo/datatables";
  import {CircleCheck, CircleMinus, Lock, LockOpen, Trash} from "lucide-svelte";
  import {invalidateAll} from "$app/navigation";
  import Avatar from "$lib/dp/components/utils/Avatar.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {TeamsUpdate, UsersUpdate} from "$lib/dp/types/pb-types.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import LocalDate from "../utils/LocalDate.svelte";
  import {positionKeysToEnumStringValues} from "$lib/dp/types/BaseballPosition.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    handler: TableHandler<CustomAuthModel>;
    team: ExpandedTeam;
    showAdminSection?: boolean;
  }

  const {handler, team, showAdminSection = false}: Props = $props();

  async function makeUserAdmin(model: CustomAuthModel) {
    try {
      await client.collection(Collection.Teams).update<TeamsUpdate>(team.id, {
        "admins+": model.id,
      });
      toastController.trigger({
        message: `User "${model.first_name + " " + model.last_name}" has been added as an admin for team "${team.name}"`,
        background: "preset-filled-success-500",
      });
      await invalidateAll();
    } catch (error) {
      console.error(error);
      toastController.trigger({
        message: "An error occurred while adding user as admin",
        background: "preset-filled-error-500",
      });
    }
  }

  async function removeUserAsAdmin(model: CustomAuthModel) {
    try {
      await client.collection(Collection.Teams).update<TeamsUpdate>(team.id, {
        "admins-": model.id,
      });
      toastController.trigger({
        message: `User "${model.first_name + " " + model.last_name}" has been removed as admin for team "${team.name}"`,
        background: "preset-filled-success-500",
      });
      await invalidateAll();
    } catch (error) {
      console.error(error);
      toastController.trigger({
        message: "An error occurred while removing user as admin",
        background: "preset-filled-error-500",
      });
    }
  }

  async function deleteUserFromTeam(model: CustomAuthModel) {
    try {
      await client.collection(Collection.Users).update<UsersUpdate>(model.id, {
        "teams-": team.id,
      });
      toastController.trigger({
        message: `User "${model.first_name + " " + model.last_name}" has been removed as member from team "${team.name}"`,
        background: "preset-filled-success-500",
      });
      await invalidateAll();
    } catch (error) {
      console.error(error);
      toastController.trigger({
        message: "An error occurred while removing user as team member",
        background: "preset-filled-error-500",
      });
    }
    closeModal();
  }

  const rows = $derived(handler.rows);
</script>

{#each rows as row}
  <tr class="content-row">
    <td>
      <div class="avatar-container">
        <Avatar
          src={client.files.getURL(row, row.avatar)}
          fallback={`${row.first_name.charAt(0)?.toUpperCase()}${row.last_name.charAt(0)?.toUpperCase()}`}
          --size="2.75rem"
        />
        <div>
          <span>{row.first_name} {row.last_name}</span>
          <div>
            <a class="font-extralight" href="mailto:{row.email}"
            >{row.email}</a
            >
          </div>
        </div>
      </div>
    </td>

    <td>
      <div class="checkmark-container">
        {#if row.verified}
          <div>
            <CircleCheck color="var(--color-success-500)"/>
          </div>
          <div>Verified</div>
        {:else}
          <div>
            <CircleMinus size="22" color="var(--color-error-500)"/>
          </div>
          <div>Unverified</div>
        {/if}
      </div>
    </td>

    <td>
      {#if row.number}
        <span class="badge-icon jersey-number preset-filled-primary-500">
          {row.number}
        </span>
      {:else}
        <span>None</span>
      {/if}
    </td>

    <td>
      {#if row.bsm_id}
        <code class="code">{row.bsm_id}</code>
      {:else}
        <span>None</span>
      {/if}
    </td>

    <td>{positionKeysToEnumStringValues(row.position)}</td>

    <td>
      <LocalDate date={row.last_login}/>
    </td>

    {#if showAdminSection}
      <td>
        <div class="flex gap-1 lg:gap-2 justify-end">

          {#if team.admins.includes(row.id)}
            <button class="badge preset-tonal-warning border border-warning-500" onclick={() => removeUserAsAdmin(row)}>
              <Lock class="m-0.5" size="18"/>
              Revoke Admin
            </button>
          {:else }
            <button class="badge preset-tonal-tertiary border border-tertiary-500" onclick={() => makeUserAdmin(row)}>
              <LockOpen class="m-0.5" size="18"/>
              Make Admin
            </button>
          {/if}

          <Dialog triggerClasses="btn btn-sm preset-tonal-error border border-error-500 gap-0!"
                  closeButtonClasses="sr-only">
            {#snippet triggerContent()}
              <Trash class="m-0.5" size="18" aria-label="Remove User from Team"/>
            {/snippet}

            {#snippet title()}
              <span>Please Confirm</span>
            {/snippet}

            {#snippet description()}
              <span>Are you sure you wish to remove "{row.first_name + " " + row.last_name}" from "{team.name}"?</span>
            {/snippet}

            <div class="button-container">
              <button class="btn preset-tonal-surface border border-surface-500" type="button" onclick={closeModal}>
                Cancel
              </button>
              <button class="btn preset-filled" type="button" onclick={() => deleteUserFromTeam(row)}>Confirm</button>
            </div>
          </Dialog>
        </div>
      </td>
    {/if}
  </tr>
{/each}

{#if rows.length === 0}
  <tr>
    <th class="py-4" colspan="8">No team members to display.</th>
  </tr>
{/if}

<style>
  .checkmark-container {
    display: flex;
    align-items: center;
    gap: var(--spacing);
  }

  .button-container {
    display: flex;
    justify-content: flex-end;
    gap: calc(var(--spacing) * 2);
    margin-block-start: var(--spacing);
  }

  .content-row {
    vertical-align: middle;
  }

  .avatar-container {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
  }

  .jersey-number {
    color: white;
    font-size: var(--text-lg);
    padding: calc(var(--spacing) * 3);
  }
</style>
