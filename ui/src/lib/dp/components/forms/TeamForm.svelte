<script lang="ts">
  import {Plus, SquarePen} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import TabsRadioGroup from "$lib/dp/components/formElements/TabsRadioGroup.svelte";
  import MultiSelectCombobox from "$lib/dp/components/formElements/MultiSelectCombobox.svelte";
  //@ts-ignore
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse, UsersResponse} from "$lib/dp/types/pb-types.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    club: ClubsResponse;
    team: ExpandedTeam | null;
    buttonClasses?: string;
    showLabel?: boolean;
  }

  const {club, team, buttonClasses = "", showLabel = true}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let open = $state(false);

  let form: Partial<ExpandedTeam> & { age_group: string } = $state(
      team ?? {
        id: "",
        name: "",
        age_group: "adults",
        signup_key: "",
        club: club.id, // no binding, cannot be changed via this form
        description: "",
        admins: [],
      }
  );

  let selectedAdmins: UsersResponse[] = $state(form?.expand?.admins ?? []);

  const allTeamMembers = client.collection(Collection.Users).getFullList<UsersResponse>({
    filter: `teams ?~ '${team?.id}'`,
    requestKey: `allTeamMembers-${team?.id}`,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    let result: ExpandedTeam | null = null;

    form.admins = selectedAdmins.map((admin) => {
      return admin.id;
    });

    try {
      if (form.id) {
        result = await client.collection(Collection.Teams).update<ExpandedTeam>(form.id, form);
      } else {
        // a user creating a team becomes its first admin
        form.admins.push(authRecord?.id);

        result = await client.collection(Collection.Teams).create<ExpandedTeam>(form);
      }
    } catch {
      toastController.triggerGenericFormErrorMessage("Team");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Team");
      open = false;
      await invalidate("teams:list");
      await invalidate("club:single");
      await invalidate("nav:load");
    }
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    {#if form.id}

      <SquarePen/>
      {#if showLabel}
        <span>Edit Team</span>
      {/if}

    {:else}

      <Plus/>
      {#if showLabel}
        <span>Create new</span>
      {/if}

    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      {#if form.id}
        <h2 class="h3">Edit Team "{form?.name}"</h2>
      {:else}
        <h2 class="h3">Create new Team</h2>
      {/if}
    </header>

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

        <label class="label col-span-2 md:col-span-1">
          Name
          <input
                  bind:value={form.name}
                  class="input"
                  name="title"
                  required
                  type="text"
          />
        </label>

        <label class="label col-span-2 md:col-span-1">
          Club
          <input
                  autocomplete="off"
                  class="input"
                  name="id"
                  readonly
                  type="text"
                  value={club?.name}
          />
        </label>

        <label class="label col-span-2">
                <span>
                Signup Key
                </span>
          <input
                  bind:value={form.signup_key}
                  class="input"
                  minlength="8"
                  name="signup_key"
                  placeholder="minimum 8 characters"
                  required
                  type="text"
          />
          <span class="text-sm">
                    A valid signup key needs to be entered upon user account creation.
                    New users are automatically added as members to the team corresponding to the signup key used.
                </span>
        </label>

        <TabsRadioGroup
                bind:value={form.age_group}
                label="Age Group"
                name="age_group"
                options={["adults", "minors"]}
                listClass="tabs-list input col-span-2"
        />

        <label class="label col-span-2">
          Description
          <textarea
                  bind:value={form.description}
                  class="textarea"
                  name="desc"
                  rows="8"
          ></textarea>
        </label>

        <label class="label space-y-3 field-wide">
          <span>Team Admins</span><br>

          {#await allTeamMembers then users}
            <MultiSelectCombobox itemName="Admin" bind:selectedItems={selectedAdmins} allItems={users}/>
          {/await}
        </label>
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
