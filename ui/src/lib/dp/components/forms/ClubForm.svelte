<script lang="ts">
  import {ClipboardEdit, Plus} from "lucide-svelte";
  import {invalidateAll} from "$app/navigation";
  import MultiSelectCombobox from "$lib/dp/components/form/MultiSelectCombobox.svelte";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse, UsersResponse, UsersUpdate} from "$lib/dp/types/pb-types.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  interface Props {
    club: ExpandedClub | null;
    buttonClasses?: string;
  }

  const {club, buttonClasses = ""}: Props = $props();

  let form: Partial<ExpandedClub> = $state(
      club ?? {
        id: "",
        name: "",
        bsm_id: 0,
        bsm_api_key: "",
        acronym: "",
        admins: [],
      }
  );

  let open = $state(false);

  let selectedAdmins: UsersResponse[] = $state(form?.expand?.admins ?? []);

  const allUsersForClub = client.collection("users").getFullList<UsersResponse>({
    filter: `club ?~ '${club?.id}'`,
    requestKey: `users-for-club-${club?.id}`,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ClubsResponse | null = null;

    try {
      if (form.id) {
        form.admins = selectedAdmins.map((admin) => {
          return admin.id;
        });

        result = await client.collection("clubs").update<ClubsResponse>(form.id, form);
      } else {
        // a user creating a club becomes its first admin
        form?.admins?.push(authRecord.id);

        result = await client.collection("clubs").create<ClubsResponse>(form);

        // a user needs to become a member of the new club
        await client.collection("users").update<UsersUpdate>(authRecord.id, {
          "club+": result.id,
        });
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Club");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Club");
      open = false;
    }
    await invalidateAll();
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    {#if form.id}
      <ClipboardEdit/>
      <span>Edit Club</span>
    {:else}
      <Plus/>
      <span>Create a Club</span>
    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Club "{form?.name}"</h2>
      {:else}
        <h2 class="h3">Create new Club</h2>
      {/if}
    </header>

    <form class="mt-4 space-y-3" onsubmit={submitForm}>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-2 md:gap-3 xl:gap-4">
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
                <span>
                Acronym
                </span>
          <input
                  bind:value={form.acronym}
                  class="input"
                  name="acronym"
                  type="text"
          />
        </label>

        <label class="label">
                <span>
                BSM Club ID
                </span>
          <input
                  bind:value={form.bsm_id}
                  class="input"
                  name="bsm_id"
                  required
                  type="number"
          >
          <span class="text-sm">Can be found in the BSM address bar while editing
                    (e.g. <span class="italic">https://bsm.baseball-softball.de/clubs/xxx/edit</span>).
                    Needs to be set for the app to function properly.
                </span>
        </label>

        <label class="label">
                <span>
                BSM API Key
                </span>
          <input
                  bind:value={form.bsm_api_key}
                  class="input"
                  name="bsm_api_key"
                  placeholder="current value not shown for security"
                  type="text"
          />
          <span class="text-sm">
                    Must be created in BSM in a user's account that has the role "Team Administration".
                    If set, all game events for the club can be automatically imported.
                </span>
        </label>

        {#if form.id /* we are editing and might want to change admins */}
          <label class="label space-y-3 md:col-span-2">
            <span>Club Admins</span><br>

            {#await allUsersForClub then users}
              <MultiSelectCombobox itemName="Admin" bind:selectedItems={selectedAdmins} allItems={users}/>
            {/await}

            <span class="text-sm">
                        It is not possible to remove the last admin from a club.
                        <br>
                        Caution: You can remove yourself as admin, losing all access rights!
                    </span>
          </label>
        {/if}
      </div>

      <hr class="my-5!"/>

      <div class="flex justify-center gap-3">
        <button class="mt-2 btn preset-filled-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>

  </Sheet.Content>
</Sheet.Root>
