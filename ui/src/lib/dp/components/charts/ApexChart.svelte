<script lang="ts">
import { onDestroy, onMount, type Snippet } from "svelte"
import { browser } from "$app/environment"

interface Props {
  options: ApexCharts.ApexOptions
  beforeContent?: Snippet
}

let { options, beforeContent }: Props = $props()

let chartElement: HTMLElement
let chart: ApexCharts

onMount(async () => {
  if (browser) {
    const ApexCharts = (await import("apexcharts")).default
    chart = new ApexCharts(chartElement, options)
    await chart.render()

    const darkModeQuery = window.matchMedia("(prefers-color-scheme: dark)")

    function changeTheme(event: MediaQueryListEvent | MediaQueryList) {
      if (event.matches) {
        chart.updateOptions({
          theme: {
            mode: "dark",
          },
        })
      } else {
        chart.updateOptions({
          theme: {
            mode: "light",
          },
        })
      }
    }

    changeTheme(darkModeQuery)
    darkModeQuery.addEventListener("change", changeTheme)
  }
})

onDestroy(() => {
  if (chart) {
    chart.destroy()
  }
})
</script>

<div class="chart-wrapper">
  {@render (beforeContent?.())}
  <div bind:this={chartElement} class="chart-container"></div>
</div>