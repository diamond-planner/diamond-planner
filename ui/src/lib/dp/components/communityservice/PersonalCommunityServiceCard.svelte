<script lang="ts">
  import {toHours} from "$lib/dp/utility/toHours.ts";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {appLocale} from "$lib/dp/locale.svelte.ts";
  import {CalendarDaysIcon, SquarePen} from "lucide-svelte";
  import ServiceEntryForm from "$lib/dp/components/forms/ServiceEntryForm.svelte";
  import type {ExpandedServiceEntry} from "$lib/dp/types/ExpandedResponse.ts";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {client} from "$lib/dp/client.svelte.ts";
  import {invalidate} from "$app/navigation";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    entry: ExpandedServiceEntry;
  }

  async function deleteEntry(id: string) {
    const result = await client.collection(Collection.ServiceEntries).delete(id);

    if (result) {
      await invalidate("communityservice:list");
    }
  }

  let {entry}: Props = $props();
  let serviceDate = $derived(new Date(entry.service_date));
</script>

<article class="card preset-tonal-surface shadow-lg">
  <div class="entry-card-header">
    <h3 class="h5">{entry.title}</h3>
    <p class="entry-minutes chip preset-tonal-primary">{toHours(entry.minutes).toFixed(1)}</p>
  </div>

  <time class="service-date" datetime="{entry.service_date}">
    <CalendarDaysIcon size="16"/>
    <span>
      {serviceDate.toLocaleDateString(appLocale.current, DateTimeUtility.eventSeriesDateFormat)}
    </span>
  </time>

  <details>
    <summary>Details</summary>
    <div class="details-content-wrapper">
      <p>{entry.description}</p>

      <div>
        <Dialog triggerClasses="btn btn-sm preset-tonal-secondary border">

          {#snippet triggerContent()}
            <SquarePen size="14"/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2 class="h4">Edit Community Service Entry "{entry.title}"</h2>
            </header>
          {/snippet}

          <ServiceEntryForm serviceEntry={entry} club={entry?.expand?.club}/>
        </Dialog>

        <DeleteButton
          action={deleteEntry}
          classes="btn-sm preset-tonal-error border border-error-500"
          id={entry.id}
          modelName="Community Service Entry"
          iconSize={14}
        />
      </div>
    </div>
  </details>
</article>

<style>
  .card {
    padding: calc(var(--spacing) * 3);
  }

  .entry-card-header {
    display: flex;
    align-items: start;
    justify-content: space-between;
  }

  .service-date {
    font-weight: var(--font-weight-light);
    font-size: var(--text-xs);
    display: flex;
    align-items: center;
    gap: var(--spacing);
    padding-block: var(--spacing);
  }

  details {
    font-size: var(--text-sm);
    padding-block-start: 0.25rem;
  }

  .entry-minutes {
    font-size: var(--text-lg);
    font-weight: var(--font-weight-bold);
  }

  .details-content-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: end;
  }
</style>
