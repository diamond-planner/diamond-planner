<script lang="ts">
  interface ConicStop {
    /** The legend label value. */
    label?: string;
    /** The desired color value. */
    color: string;
    /** Starting stop value (0-100) */
    start: number;
    /** Ending stop value (0-100) */
    end: number;
  }

  const {
    stops = [{color: "neutral", start: 0, end: 100}],
    legend = false,
    spin = false,
    width = "w-24",
    hover = "bg-primary-hover-token",
    digits = 0,
    regionCaption = "",
    regionCone = "",
    regionLegend = "",
  } = $props<{
    /** Provide a data set of color stops and labels. */
    stops?: ConicStop[]
    /** Enable a contextual legend. */
    legend?: boolean
    /** When enabled, the conic gradient will spin. */
    spin?: boolean
    /** Style the conic gradient width. */
    width?: string
    /** Style the legend hover effect. */
    hover?: string
    /** Set the number of digits on the legend values. */
    digits?: number
    /** Style the caption region above the gradient. */
    regionCaption?: string
    /** Style the conic gradient region. */
    regionCone?: string
    /** Style the legend region below the gradient. */
    regionLegend?: string
  }>();

  // Local variables
  let cone: string = $state("");
  let generatedLegendList: any[] = $state([]);

  // Styles
  const cBase = "flex flex-col items-center space-y-4 w-";
  const cCaption = "text-center";
  const cCone = "block aspect-square rounded-full";
  const cLegend = "text-sm w-full";
  const cSwatch = "block aspect-square bg-black w-5 rounded-full mr-2";

  // Generate Conic Gradient style
  function genConicGradient(): void {
    let d: any = stops?.map((v: ConicStop) => `${v.color} ${v.start}% ${v.end}%`);
    cone = `conic-gradient(${d.join(", ")})`;
  }

  // Generate Legend
  function genLegend(): any {
    if (!legend) return;
    generatedLegendList = stops?.map((v: ConicStop) => {
      return {
        label: v.label,
        color: v.color,
        value: (v.end - v.start).toFixed(digits),
      };
    });
  }

  $effect(() => {
    genConicGradient();
    genLegend();
  });

  const classesBase = $derived(`${cBase}`);
  const classesCaption = $derived(`${cCaption} ${regionCaption}`);
  const classesCone = $derived(`${cCone} ${width} ${regionCone}`);
  const classesLegend = $derived(`${cLegend} ${regionLegend}`);
</script>

<!--
@component Converted from Skeleton v2
-->

<figure class="conic-gradient {classesBase}" data-testid="conic-gradient">
  <!-- Label -->
  <figcaption class="conic-caption {classesCaption}"></figcaption>
  <!-- Conic Gradient -->
  {#if cone}
    <div class="conic-cone {classesCone}" class:animate-spin={spin} style:background={cone}></div>
  {/if}
  <!-- Legend -->
  {#if legend && generatedLegendList}
    <ul class="conic-list list {classesLegend} space-y-2">
      {#each generatedLegendList as {color, label, value}}
        <li class="conic-item text-xs flex {hover}">
          <span class="conic-swatch {cSwatch}" style:background={color}></span>
          <span class="conic-label flex-auto">{label}</span>
          <strong class="conic-value">{value}%</strong>
        </li>
      {/each}
    </ul>
  {/if}
</figure>