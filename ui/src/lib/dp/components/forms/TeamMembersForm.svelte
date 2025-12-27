<script lang="ts">
  import {invalidateAll} from "$app/navigation";
  import MultiSelectCombobox from "$lib/dp/components/formElements/MultiSelectCombobox.svelte";
  import {client, manualAuthRefresh} from "$lib/dp/client.svelte.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse, UsersResponse, UsersUpdate} from "$lib/dp/types/pb-types.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  interface Props {
    club: ClubsResponse;
    team: ExpandedTeam;
  }

  const {club, team}: Props = $props();

  let form = $state({
    id: team.id,
  });

  let selectedUsers: UsersResponse[] = $state([]);

  const allUsersForClub = client.collection("users").getFullList<UsersResponse>({
    filter: `club ?~ '${club.id}' && teams !~ '${team.id}'`, // all users that are club members, but not members of this team
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    try {
      for (const user of selectedUsers) {
        await client.collection("users").update<UsersUpdate>(user.id, {
          "teams+": team.id,
        });
      }
      toastController.triggerGenericFormSuccessMessage("Team members");

      // set manually in case we have updated the currently logged-in user
      await manualAuthRefresh();
    } catch (error) {
      console.error(error);
      toastController.triggerGenericFormErrorMessage("Team members");
    }
    await invalidateAll();
    closeModal();
  }
</script>

<form class="mt-4 space-y-3" onsubmit={submitForm}>
  <div class="edit-form-grid">
    <input
            autocomplete="off"
            bind:value={form.id}
            class="input"
            name="id"
            readonly
            type="hidden"
    />

    <label class="label space-y-3 field-wide">
      <span>Selected Users:</span><br>

      {#if selectedUsers.length === 0}
        <span>---</span>
      {/if}

      {#await allUsersForClub then users}
        <MultiSelectCombobox
                itemName="Member"
                bind:selectedItems={selectedUsers}
                allItems={users}
                allowDeletionOfLastItem={true}
        />
      {/await}
    </label>
  </div>

  <hr class="my-5!"/>

  <div class="flex justify-center gap-3">
    <button class="mt-2 btn preset-tonal-primary border border-primary-500" type="submit">
      Submit
    </button>
  </div>
</form>
