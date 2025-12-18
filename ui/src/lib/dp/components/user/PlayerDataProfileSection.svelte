<script lang="ts">
  import {Link} from "lucide-svelte";
  import PlayerDataForm from "$lib/dp/components/forms/PlayerDataForm.svelte";
  import PlayerDataCard from "$lib/dp/components/player/PlayerDataCard.svelte";
  import PlayerHeaderSection from "$lib/dp/components/player/PlayerHeaderSection.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {positionKeysToEnumStringValues} from "$lib/dp/types/BaseballPosition.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import type {Player} from "$lib/dp/types/Player.ts";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  // suboptimal adapter until data is properly fetched from PocketBase
  //@ts-expect-error
  const playerObject: Player = $derived({
    firstname: authRecord.first_name,
    lastname: authRecord.last_name,
    fullname: authRecord.first_name + " " + authRecord.last_name,
    birthday: 0,
    admission: new Date(authRecord.created).toLocaleDateString(),
    number: authRecord.number,
    throwing: authRecord.throws,
    batting: authRecord.bats,
    bsm_id: authRecord.bsm_id,
    positions: positionKeysToEnumStringValues(authRecord.position),
    teams: [],
    coach: "",
    media: [
      {
        url: client.files.getURL(authRecord, authRecord.avatar),
        alt: "Player Profile Avatar",
      },
    ],
  });
</script>

<div class="card preset-tonal-primary shadow-lg flex flex-col justify-between">
  <header class="card-header">
    <h2 class="h4 font-semibold">Public Profile</h2>
  </header>

  <section class="p-4 font-light">
    <p>
      You have the option to display your player data in your public
      profile on the website and apps.
    </p>
    <p class="mt-2">
      Please note that all statistics and basic info such as your name and
      team affiliation is considered public information when participating
      in official DBV and BSVBB competitions.
    </p>
  </section>

  <footer class="card-footer">
    <div class="flex flex-col gap-2 lg:gap-3">
      <PlayerDataForm buttonClasses="btn preset-tonal-primary border border-primary-500"></PlayerDataForm>

      {#if authRecord.bsm_id}
        <a
          href="/players/{authRecord?.bsm_id}"
          class="btn preset-tonal-secondary border border-secondary-500"
        >
          <Link/>
          <span>Go to Profile Page</span>
        </a>
      {/if}
    </div>
  </footer>
</div>

<PlayerHeaderSection player={playerObject}/>

<PlayerDataCard player={playerObject} showTeams={false}/>
