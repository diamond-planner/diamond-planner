<script lang="ts">
  import {SquarePen} from "lucide-svelte";
  import TabsRadioGroup from "$lib/dp/components/formElements/TabsRadioGroup.svelte";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import TagsInput from "$lib/dp/components/formElements/TagsInput.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {
    getAllBaseballPositionStringValues,
    positionEnumStringValuesToKeys,
    positionKeysToEnumStringValues,
  } from "$lib/dp/types/BaseballPosition.ts";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";

  interface ValidateArgs {
    inputValue: string;
    value: string[];
  }

  interface Props {
    buttonClasses?: string;
  }

  const {buttonClasses = ""}: Props = $props();

  // Mark: here we do not need to be reactive as we're editing
  const authRecord = authSettings.record as CustomAuthModel;
  let form = $state({
    id: authRecord.id ?? "",
    number: authRecord.number ?? "",
    bats: authRecord.bats ?? "",
    position: [""],
    throws: authRecord.throws ?? "",
    image: authRecord.image ?? "",
    umpire: authRecord.umpire ?? "0",
    scorer: authRecord.scorer ?? "0",
    bsm_id: authRecord.bsm_id ?? 0,
    display_on_website: authRecord.display_on_website ?? false,
  });

  let open = $state(false);

  const possiblePositionValues = getAllBaseballPositionStringValues();
  let selectedPositions: string[] = $state(positionKeysToEnumStringValues(authRecord.position));

  function validatePositionValue(details: ValidateArgs): boolean {
    return possiblePositionValues.includes(details.inputValue);
  }

  function addPosition(value: string) {
    if (!selectedPositions.includes(value)) {
      selectedPositions.push(value);
    } else {
      toastController.trigger({
        message: "Position is already selected.",
        background: "preset-filled-error-500",
      });
    }
  }

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    form.position = positionEnumStringValuesToKeys(selectedPositions);

    let result: CustomAuthModel | null = null;

    try {
      result = await client.collection("users").update<CustomAuthModel>(form.id, form);
    } catch {
      toastController.triggerGenericFormErrorMessage("Player data");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Player data");
      open = false;
    }
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    <SquarePen/>
    <span>Edit Player Data</span>
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      <h2 class="h3">
        Edit Player Data for "{`${authRecord.first_name} ${authRecord.last_name}`}"
      </h2>
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

        <label class="label">
          Jersey Number
          <input
                  autocomplete="off"
                  bind:value={form.number}
                  class="input"
                  name="number"
                  type="number"
          />
        </label>

        <label class="label">
          BSM ID (Player Pass Number)
          <input
                  autocomplete="off"
                  bind:value={form.bsm_id}
                  class="input"
                  name="bsm_id"
                  readonly
                  type="text"
          />
        </label>

        <span class="label-wide">
          <Switch
                  bind:checked={form.display_on_website}
                  name="display_on_website"
          >
            Display this data publicly on team website
          </Switch>
        </span>

        <label class="label label-wide flex flex-col gap-1 md:gap-2">
          Positions
          <TagsInput
                  name="positions"
                  onValueChange={(e) => (selectedPositions = e.value)}
                  placeholder="Enter positions or click buttons..."
                  validate={validatePositionValue}
                  value={selectedPositions}
          />
          <span class="flex flex-wrap gap-2 label-wide">
          {#each possiblePositionValues as value}
            <button
                    type="button"
                    class="btn btn-sm preset-outlined-primary-500"
                    onclick={() => addPosition(value)}
            >
              {value}
            </button>
          {/each}
        </span>
        </label>

        <TabsRadioGroup
                bind:value={form.bats}
                label="Bats"
                name="bats"
                options={["left", "right", "switch"]}
        />

        <TabsRadioGroup
                bind:value={form.throws}
                label="Throws"
                name="throws"
                options={["left", "right", "switch"]}
        />

        <TabsRadioGroup
                bind:value={form.umpire}
                label="Umpire License"
                name="umpire"
                options={["none", "A", "B", "C", "D"]}
        />

        <TabsRadioGroup
                bind:value={form.scorer}
                label="Scorer License"
                name="scorer"
                options={["none", "A", "B", "C", "D"]}
        />
      </div>
      <hr class="my-5!"/>

      <div class="flex justify-center gap-3">
        <button class="mt-2 btn preset-tonal-primary border border-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>

<style>
    .label-wide {
        grid-column: span 2 / span 2;
        gap: calc(var(--spacing) * 2);
    }
</style>
