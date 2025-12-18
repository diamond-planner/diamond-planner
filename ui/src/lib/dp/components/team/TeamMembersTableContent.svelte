<script lang="ts">
  import type {DataHandler} from "@vincjo/datatables";
  import {CircleCheck, Lock, LockOpen, Trash} from "lucide-svelte";
  import {invalidateAll} from "$app/navigation";
  import Avatar from "$lib/dp/components/utils/Avatar.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {TeamsUpdate, UsersUpdate} from "$lib/dp/types/pb-types.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import LocalDate from "../utils/LocalDate.svelte";
  import ThSort from "$lib/dp/components/datatable/ThSort.svelte";

  interface Props {
    handler: DataHandler<CustomAuthModel>;
    team: ExpandedTeam;
    showAdminSection?: boolean;
  }

  const {handler, team, showAdminSection = false}: Props = $props();

  async function makeUserAdmin(model: CustomAuthModel) {
    try {
      await client.collection("teams").update<TeamsUpdate>(team.id, {
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
      await client.collection("teams").update<TeamsUpdate>(team.id, {
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
      await client.collection("users").update<UsersUpdate>(model.id, {
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

  const rows = $derived(handler.getRows());
</script>

<thead>
<tr class="sticky preset-tonal-surface dark:preset-filled-surface-300-700">
  <ThSort {handler} orderBy="last_name">Name</ThSort>
  <ThSort {handler} orderBy="verified">Status</ThSort>
  <ThSort {handler} orderBy="number">Number</ThSort>
  <ThSort {handler} orderBy="bsm_id">BSM ID</ThSort>
  <ThSort {handler} orderBy="position">Position</ThSort>
  <ThSort {handler} orderBy="last_login">Last Login</ThSort>

  {#if showAdminSection}
    <th>Actions</th>
  {/if}
</tr>
</thead>

<tbody>
{#each $rows as row}
  <tr class="align-middle">
    <td>
      <div class="flex items-center gap-2">
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
      <div class="flex gap-1">
        {#if row.verified}
          <div>
            <CircleCheck
                    class="text-success-600 dark:text-success-500"
            />
          </div>
          <div>Verified</div>
        {:else}
          <div>
            <!-- Icon for some reason not in lib -->
            <svg
                    class="w-5 h-5 text-error-600-400"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    fill="currentColor"
                    viewBox="0 0 24 24"
            >
              <path
                      fill-rule="evenodd"
                      d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm5.757-1a1 1 0 1 0 0 2h8.486a1 1 0 1 0 0-2H7.757Z"
                      clip-rule="evenodd"
              />
            </svg>
          </div>
          <div>Unverified</div>
        {/if}
      </div>
    </td>

    <td>
      {#if row.number}
                    <span class="badge-icon p-3 preset-filled-primary-500 text-white text-lg"
                    >{row.number}</span
                    >
      {:else}
        <span>None</span>
      {/if}
    </td>

    <td>
      {#if row.bsm_id}
        {row.bsm_id}
      {:else}
        <span>None</span>
      {/if}
    </td>

    <td>{row.position}</td>

    <td>
      <LocalDate date={row.last_login}/>
    </td>

    {#if showAdminSection}
      <td>
        <div class="flex gap-1 lg:gap-2 justify-end">
          <!--        <button class="btn btn-sm variant-ghost-primary">-->
          <!--          <EditOutline/>-->
          <!--          Edit-->
          <!--        </button>-->

          {#if team.admins.includes(row.id)}
            <button class="badge preset-tonal-warning border border-warning-500" onclick={() => removeUserAsAdmin(row)}>
              <Lock class="m-0.5" size="18"/>
              Revoke Admin Access
            </button>
          {:else }
            <button class="badge preset-tonal-tertiary border border-tertiary-500" onclick={() => makeUserAdmin(row)}>
              <LockOpen class="m-0.5" size="18"/>
              Make Admin
            </button>
          {/if}

          <Dialog triggerClasses="badge preset-tonal-error border border-error-500 gap-0!" closeButtonClasses="sr-only">
            {#snippet triggerContent()}
              <Trash class="m-0.5" size="18"/>
              Remove
            {/snippet}

            {#snippet title()}
              <span>Please Confirm</span>
            {/snippet}

            {#snippet description()}
              <span>Are you sure you wish to remove "{row.first_name + " " + row.last_name}" from "{team.name}"?</span>
            {/snippet}

            <div class="flex justify-end gap-2 mt-1">
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

{#if $rows.length === 0}
  <tr>
    <th class="py-4" colspan="8">No team members to display.</th>
  </tr>
{/if}
</tbody>
