<script lang="ts">
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";
  import {EventSeriesUtility} from "$lib/dp/service/EventSeriesUtility.ts";
  import type {EventseriesResponse} from "$lib/dp/types/pb-types.ts";
  import {appLocale} from "$lib/dp/locale.svelte.ts";

  interface Props {
    eventSeries: EventseriesResponse;
  }

  const {eventSeries}: Props = $props();

  const startDate = $derived(new Date(eventSeries.series_start));
  const endDate = $derived(new Date(eventSeries.series_end));

  const options: { weekday: "long" | "short" | "narrow" | undefined } = {weekday: "long"};
  const seriesState = $derived(EventSeriesUtility.getSeriesState(startDate, endDate));
</script>

<article class="p-4 preset-outlined-surface rounded-base shadow-md">
  <div class="grid grid-cols-1 md:grid-cols-3 md:gap-2 lg:gap-4">

    <div class="flex flex-col justify-between mb-3 md:mb-0">
      <h4 class="h5 font-bold">{eventSeries.title}</h4>
      <div>every {eventSeries.interval} days</div>
    </div>

    <div class="grid grid-cols-2 gap-4 md:col-span-2">
      <!-- Start Date -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Series First Occurrence:</span>
        <time class="text-sm md:text-base" datetime="{eventSeries.series_start}">
          {startDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateTimeFormat)}
        </time>
      </div>
      <!-- End Date -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Series End After:</span>
        <time class="text-sm md:text-base" datetime="{eventSeries.series_end}">
          {endDate.toLocaleDateString("de-DE", DateTimeUtility.eventSeriesDateFormat)}
        </time>
      </div>
      <!-- Start Time -->
      <div class="flex flex-col items-start">
        <span class="text-sm font-medium">Event Duration:</span>
        <time class="text-sm md:text-base" datetime="P{eventSeries.duration}M">{eventSeries.duration} Minutes</time>
      </div>

      <div class="justify-self-start self-center flex gap-2 flex-wrap">
        <span class="badge block preset-outlined">
          {new Intl.DateTimeFormat(appLocale.current, options).format(startDate)}
        </span>

        {#if seriesState === "ongoing"}
        <span class="badge block preset-tonal-primary border border-primary-500">
          Ongoing
        </span>
        {:else if seriesState === "past"}
          <span class="badge block preset-outlined">
          Past
        </span>
        {:else if seriesState === "future"}
          <span class="badge block preset-tonal-tertiary border border-tertiary-500">
          Upcoming
        </span>
        {/if}
      </div>
    </div>
  </div>
</article>
