<script lang="ts">
  import {Edit, Plus} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import LeafletMapCoordinatePicker from "$lib/dp/components/map/LeafletMapCoordinatePicker.svelte";
  //@ts-expect-error
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import {save} from "$lib/dp/records/RecordOperations.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
  import type {LocationsResponse} from "$lib/dp/types/pb-types.ts";

  // Gail S. Halvorsen Park coordinates
  const DEFAULT_LATITUDE = 52.482762;
  const DEFAULT_LONGITUDE = 13.407932;

  interface Props {
    club: ExpandedClub;
    location?: LocationsResponse | null;
    buttonClasses?: string;
  }

  const {location = null, club, buttonClasses = ""}: Props = $props();

  let form = $state(
      location ?? {
        id: "",
        name: "",
        street: "",
        postal_code: "",
        city: "",
        address_addon: "",
        country: "",
        human_country: "",
        internal_name: "",
        bsm_id: 0,
        latitude: DEFAULT_LATITUDE,
        longitude: DEFAULT_LONGITUDE,
        club: club.id,
      }
  );

  let open = $state(false);

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();
    form.club = club.id;

    let result: LocationsResponse | null = null;

    try {
      result = await save<LocationsResponse>("locations", form);
    } catch {
      toastController.triggerGenericFormErrorMessage("Location");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Location");
      open = false;
    }
    await invalidate("club:locations");
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    {#if form.id}
      <Edit/>
    {:else}
      <Plus/>
      <span>Create new</span>
    {/if}
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="my-3">
      <h2 class="h2">
        {#if form.id}
          Edit Location "{form?.address_addon ?? form?.internal_name}"
        {:else}
          Add New Location
        {/if}
      </h2>
    </header>

    <form onsubmit={submitForm}>
      <div class="edit-form-grid">
        <div>
          <label class="" for="address_addon">Address Addon*</label>
          <input
                  bind:value={form.address_addon}
                  class="input"
                  id="address_addon"
                  name="address_addon"
                  required
                  type="text"
          />
          <span class="font-light text-sm"
          >Often used in BSM data for the field name.</span
          >
        </div>

        <div>
          <label class="label" for="name">Name (Sport)*</label>
          <input
                  bind:value={form.name}
                  class="input"
                  id="name"
                  name="name"
                  required
                  type="text"
          />
        </div>

        <div>
          <label class="label" for="street">Street</label>
          <input
                  bind:value={form.street}
                  class="input"
                  id="street"
                  name="street"
                  type="text"
          />
        </div>

        <div>
          <label class="label" for="postal_code">Postal Code</label>
          <input
                  bind:value={form.postal_code}
                  class="input"
                  id="postal_code"
                  name="postal_code"
                  type="text"
          />
        </div>

        <div>
          <label class="label" for="city">City</label>
          <input
                  bind:value={form.city}
                  class="input"
                  id="city"
                  name="city"
                  type="text"
          />
        </div>

        <div>
          <label class="label" for="bsm_id">BSM ID</label>
          <input
                  autocomplete="off"
                  bind:value={form.bsm_id}
                  class="input"
                  id="bsm_id"
                  name="bsm_id"
                  readonly
                  type="number"
          />
        </div>

        <div>
          <label class="label" for="country">Country Code</label>
          <input
                  bind:value={form.country}
                  class="input"
                  id="country"
                  maxlength="2"
                  name="country"
                  type="text"
          />
        </div>

        <div>
          <label class="label" for="human_country">Country Name</label>
          <input
                  bind:value={form.human_country}
                  class="input"
                  id="human_country"
                  name="human_country"
                  type="text"
          />
        </div>

        <div class="field-wide">
          <label class="label" for="internal_name">Internal Name</label>
          <input
                  bind:value={form.internal_name}
                  class="input"
                  id="internal_name"
                  name="internal_name"
                  type="text"
          />
          <span class="font-light text-sm"
          >Custom name field that will not be overwritten by automatic
            imports.</span
          >
        </div>

        <div>
          <label class="label" for="latitude">Latitude*</label>
          <input
                  bind:value={form.latitude}
                  class="input"
                  id="latitude"
                  name="latitude"
                  required
                  step="any"
                  type="number"
          />
        </div>

        <div>
          <label class="label" for="longitude">Longitude*</label>
          <input
                  bind:value={form.longitude}
                  class="input"
                  id="longitude"
                  name="longitude"
                  required
                  step="any"
                  type="number"
          />
        </div>

        <label class="field-wide">
          Coordinate Map
          <span class="text-sm font-light block">Select a location on the map to set coordinates automatically</span>
          <LeafletMapCoordinatePicker bind:latitude={form.latitude} bind:longitude={form.longitude}/>
        </label>

      </div>

      <div class="mt-4 flex justify-between items-center">
        <button class="btn preset-tonal-primary border border-primary-500" type="submit">Submit</button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>
