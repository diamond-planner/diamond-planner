<script lang="ts">
  import {invalidate} from "$app/navigation";
  import LocationForm from "$lib/dp/components/forms/LocationForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
  import type {LocationsResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    club: ExpandedClub;
    locations: LocationsResponse[];
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const {club, locations}: Props = $props();

  function locationDeleteAction(id: string) {
    client.collection("locations").delete(id);
    invalidate("club:locations");
  }
</script>

<ul class="preset-tonal-surface rounded-base p-4 shadow-xl text-sm">
  {#each locations as location, index (location.id)}
    <li class="location-grid">
      <h3 class="font-bold overflow-hidden">
        {location.address_addon
            ? location.address_addon
            : location.internal_name} ({location.name})
      </h3>
      <p>
        {location?.street}, {location.postal_code}
        {location.city}
      </p>

      <div class="justify-self-end flex flex-col md:flex-row gap-1 items-center">
        {#if club?.admins.includes(authRecord.id)}
          <LocationForm
                  {club}
                  {location}
                  buttonClasses="btn btn-icon preset-tonal-tertiary border border-tertiary-500"
          />

          <DeleteButton
                  id={location.id}
                  modelName="Location"
                  action={locationDeleteAction}
                  classes="preset-tonal-error border border-error-500 btn-icon ms-1"
          />
        {/if}
      </div>
    </li>

    {#if index < locations.length - 1}
      <hr class="my-2"/>
    {/if}

  {:else }
    <p>This club doesn't have saved locations yet.</p>
  {/each}
</ul>

{#if club?.admins.includes(authRecord.id)}
  <LocationForm
          {club}
          location={null}
          buttonClasses="btn preset-tonal-primary border border-primary-500 my-3"
  />
{/if}

<style lang="postcss">
    .location-grid {
        display: grid;
        grid-template-columns: repeat(3, minmax(0, 1fr));
        align-items: center;
        gap: 1rem;
    }
</style>
