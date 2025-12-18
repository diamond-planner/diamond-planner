<script lang="ts">
  import StatsBlockContent from "$lib/dp/components/stats/StatsBlockContent.svelte";
  import type {AttendanceTotal} from "$lib/dp/types/PersonalAttendanceStats.ts";
  import {capitalize} from "$lib/dp/utility/capitalize.ts";
  import type {SingleStatElement} from "$lib/dp/types/SingleStatElement.ts";

  interface Props {
    season: number | "All Time";
    attendance: AttendanceTotal;
  }

  const {attendance, season}: Props = $props();

  const title = capitalize(attendance.type);

  const block: SingleStatElement = {
    title: title,
    value: `${attendance.attended} / ${attendance.total}`,
    desc: `All possible events of this type for ${season}`,
  };
</script>

<div class="card preset-tonal-surface shadow-lg">
  <header class="card-header">
    <h2 class="h4 font-semibold">Totals - {title}</h2>
  </header>

  <section class="p-4 flex justify-center">
    <StatsBlockContent {block} classes="place-items-center gap-3"/>
  </section>
</div>
