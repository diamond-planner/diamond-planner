<script lang="ts">
  import type {ServiceentriesResponse} from "$lib/dp/types/pb-types.ts";
  import {toHours} from "$lib/dp/utility/toHours.ts";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {appLocale} from "$lib/dp/locale.svelte.ts";
  import {CalendarDaysIcon} from "lucide-svelte";

  interface Props {
    entry: ServiceentriesResponse;
  }

  let {entry}: Props = $props();
  let serviceDate = $derived(new Date(entry.service_date));
</script>

<article class="card preset-tonal-surface shadow-lg">
  <div class="entry-card-header">
    <h3 class="h5">{entry.title}</h3>
    <p class="entry-minutes chip preset-tonal-primary">{toHours(entry.minutes)}</p>
  </div>

  <time class="service-date" datetime="{entry.service_date}">
    <CalendarDaysIcon size="16"/>
    <span>
      {serviceDate.toLocaleDateString(appLocale.current, DateTimeUtility.eventSeriesDateFormat)}
    </span>
  </time>

  <details>
    <summary>Details</summary>
    <p>{entry.description}</p>
  </details>
</article>

<style>
  .card {
    padding: calc(var(--spacing) * 3);
  }

  .entry-card-header {
    display: flex;
    align-items: center;
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
</style>
